package chapter7

/*
			7.4 경쟁 상태

	고루틴들을 다루다 보면 여러 가지 복잡한 버그를 만들 수 있다. 모든 고루틴들이 막혀서 교착 상태(deadlock)가 발생하는 경우라면
	더 이상 진행이 되지 않기 때문에 프로그램이 오류를 출력하겠지만 쉽게 발겨하지 못하는 버그들이 있을 수 있다. 그중 하나가
	경쟁 상태(race condition)이다.

	경쟁 상태는 어떤 공유된 자원에 둘 이상의 프로세스가 동시에 접근하여 잘못된 결과가 나올 수 있는 상태를 말한다. 타이밍에 따라서
	결과가 달라질 수 있기에 고치기가 번거로운 버그들을 만들어 낸다. 이번에는 경쟁 상태에 빠지지 않게 하기 위한 기술들을 습득해 본다.

	채널을 잘 활용하면 이미 경쟁 상태에 문제를 많이 해결할 수 있다. 그러나 몇 가지 경우에는 채널을 이용하는 것 보다 sync 라이브러리를
	활용하는 것이 더 간단하며 라이브러리를 직접 개발할 대, 특히 atomic 라이브러리를 활용해야 하는 경우가 생긴다. 여기서는 atomic 라이브러리와
	채널을 활용하여 경쟁상태를 해결하고, sync 라이브러리에 있는 기능들을 활용하는 방법을 알아본다.





		7.4.1 동시성 디버그


	고루틴들을 다루다 보면 여러 가지 버그를 만들 수도 있다. 모든 고루틴들이 막혀서 교착 상태(deadlock)가 발생하는 경우라는 더 이상 진행이 안되기
	때문에 프로그램이 오류를 출력하겠지만 쉽게 발견하지 못하는 버그들이 있다. 그 중 하나가 경쟁 상태(race condition) 이다.


	경쟁 상태는 어떤 공유된 자원에 둘 이상의 프로세스가 동시에 접근하여 잘못된 결과가 나올 수 있는 상태를 의미한다. 타이밍에 따라서 결과가 달라질 수
	있기 때문에 고치기가 번거로운 버그를 만들어 낸다. 이번에는 경쟁 상태에 빠지지 않게 하기 위한 기술들을 습득해본다.


	채널을 잘 활용하면 이미 경쟁 상태 문제를 많이 해결할 수 있다. 그러나 몇 가지 경우에는 채널을 이용하는 것보다 sync 라이브러리를 활용하는 것이 더
	간단하며 이런 라이브러리를 직접 개발할 때, 특히 atomic 라이브러리를 활용해야 하는 경우가 생긴다. 여기서는 atomic 라이브러리와 채널을 경쟁 상태를
	해결하고 , sync 라이브러리에 있는 기능을 활용하는 방법을 알아본다.




		7.4.2 동시성 디버그


	다행히도 경쟁 상태 탐지 기능이 있다. 다음과 같이 go 도구를 이용할 때 -race  옵션을 주면 된다.

	$ go test -race mypkg		// to test the package
	$ go run -race mysrc.go		// to run the source file
	$ go build -race mycmd		// to build the command
	$ go install -race mypkg	// to instal the package




	동적으로 고루틴 이어붙이는 예제를 작성할 때, 고루틴들이 막혀서 메모리 누수가 일어나고 있는지 궁금할 것이다.
	현재 작동하는 고루틴의 수를 보고 싶으면 runtime.NumGoroutine()을 호출하면 된다. runtime에는 유용한 함수들도
	더 있다.

	runtime.NumCPU()와 runtime.GOMAXPROCS()를 이용하여 현재 사용 가능한 CPU의 수와 얼마나 많은 CPU를 이용할 것인지
	통제할 수도 있다.


	만약에 고루틴 수가 점점 더 늘어나면 어디선가 고루틴이 막혀 있을 가능성이 높다. 이럴 때는 적당한 시점에서 panic을
	발생시키면 고루틴의 스택 추적이 출력된다. 이것을 보면 몇 번째 라인에서 고루틴이 멈춰 있는지 확인할 수 있다.




	7.4.2 atomic과 sync.WaitGroup


	package main

import (
	"fmt"
	"time"
)

func main() {

	cnt := int64(10)

	for i := 0; i < 10; i++ {
		go func() {
			// do something
			cnt--
		}()
	}

	for cnt > 0 {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println(cnt)
}


고루틴 10개를 생성하였고 각각의 고루틴에서 원래 10으로 되어 있는 카운트를 감소시킨다.


마지막에 cnt를 출력하는 줄에 와서는 0을 출력하게 된다. 실제로 실행해보아도 0이 나온다.

그러면 된 걸까? 사실 이 코드에는 경쟁상태가 있다. 실행할 때 -race 옵션을 주면 경쟁 상태가
있다는 오류를 발생시킨다. 바로 cnt--부분과 cnt>0부분이다. 한줄로 되어 있는 이 코드에 무슨
경쟁 상태가 있다는 것일까?

Go코드는 한줄이지만 이 연산이 반드시 원자성(atomicity)을 띠지는 않는다. 메모리에서 값을 읽은
다음에 숫자를 하나 감소시키고 그 값을 다시 메모리에 저장하여야 하기 때문에 이것은 한 줄이지만
여러 연산을 해야 한다. 메모리에서 값을 일고 나서 다른 고루틴이 메모리를 읽은 다음에 값을 감소
시켜서 저장하고, 다시 이전의 고루틴으로 돌아와서 값을 저장한다면 2 감소되어야 할 값이 1감소될
수 있는 것이다.


그래서 운이 나쁘면 이 코드는 끝나지 않는 상태에 빠질 수 있다. for 반복문을 돌면서 계속 Sleep하기
때문에 완전히 고루틴이 막히지도 않아서 교착 상태가 쉽게 감지되지도 않는다.


sync/atomic 패키지에는 이런 경우를 대비하기위한 함수들이 있다. cnt--를 atomic.AddInt64(&cnt,-1)로
바꾸어주고 cnt>0을 atomic.LoadInt64(&cnt)>0로 바꾸어 주면 된다. 이제 -race 옵션을 줘도 아무 불평이
없어졌다.

atomic 패키지에는 이 외에도 다른 함수들이 있다. 한번 살펴보는 것이 좋다.


다행히 -race 옵션으로 검사할 수 있지만 상당히 복잡하다. 지금까지 채널을 배웠는데 채널을 이용하면 이런 복잡한
문제들을 아주 쉽게 해결할 수 있다. 채널을 이용하면 atomic도 없고, 락도 없는 코드에서 많은 동시성 문제들이 해결된다.



package main

import (
	"fmt"
)

func main() {

	req, resp := make(chan struct{}), make(chan int64)

	cnt := int64(10)

	go func(cnt int64) {

		defer close(resp)

		for _ = range req {
			cnt--
			resp <- cnt
		}

	}(cnt)

	for i := 0; i < 10; i++ {
		go func() {
			// do something
			req <- struct{}{}
		}()
	}

	for cnt = <-resp; cnt > 0; cnt = <-resp {
		close(req)
		fmt.Println(cnt)
	}
}


여기에는 atomic이 없다. 그럼에도 불구하고 경쟁 상태가 생기지 않는데, 채널이 싱크를 맞춰주기 때문이다.
따라서 채널을 이용하면 동시성이 있는 코드에서 발생하는 문제들을 많이 줄여줄 수 있기 때문이다. 그러면
이 코드면 될까?

채널 만으로 문제를 풀 수 있지만, 이 코드를 다시 보면 사실은 WaitGroup을 쓰고 싶었다는 것을 알 수 있다.
WaitGroup을 쓰지 않고 채널을 쓸 수 있지만, 읽는 사람이 정확히 무엇을 하고자 하는지 바로 파악하기
어려울 수 있다. 따라서 이 경우는 WaitGroup을 써 준다.


func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}

	wg.Wait()
}


WaitGroup을 썼기 때문에 의도를 명확히 파악알 수 있었따.






	7.4.3 sync.Once



	package main

import (
	"fmt"
	"sync"
)

func main() {

	done := make(chan struct{})

	go func() {
		defer close(done)
		fmt.Println("Initialized")
	}()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			<-done
			fmt.Println("Goroutine: ", i)
		}(i)

	}
	wg.Wait()
}


3번째 줄에 있는 함수는 초기화 함수이다. 그리고 8번째 있는 3번 반복하는 반복문의 시작이고,
10번째 줄부터 있는 코드는 고루틴이지만, <-done으로 먼저 기다린 뒤에 화면에 출력한다. 결과는
다음과 같이 나오거나 조금 다른 순서로 나올 것이다. 중요한 것은 어느 고루틴의 출력 부분이
수행되기 전에 초기화 코드가 수행된다는 점이다.



이렇게 한 번만 어떤 코드를 수행하고자 할 때 쓸 수 있는 것이 sync.Once이다. 주로 분산 처리를 할 때
, 초기화 코드에서 이용할 수 있다. 초기화는 한번만 하고 싶고, 모든 분산 처리 고루틴이 이 초기화를
하고 나서 수행되어야 하는 경우에 이용할 수 있다. 분산 처림 고루틴을 만들기 전에 초기화하면 되는데,
경우에 따라서는 기존의 라이브러리 등을 이용하면서 이렇게 하기 불가능한 경우가 있다. 이때 각각의
고루틴에서 이것을 활용하면 된다.



물론 채널만 이용하여 같은 효과를 낼 수 있지만 그렇게 하는 것보다 sync.Once 이용을 더 권장한다. 이유는 코드를
읽을 때 더 분명히 무엇을 하는지 파악할 수 있기 때문이다.

package main

import (
	"fmt"
	"sync"
)

func main() {

	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("Initialized")
			})
			fmt.Println("Goroutine:", i)
		}(i)
	}
	wg.Wait()
}


*/
