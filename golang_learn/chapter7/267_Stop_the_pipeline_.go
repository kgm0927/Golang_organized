package chapter7

/*
	7.3.6 파이프라인 중단하기


	자금까지 파이프라인을 구성할 때, 받기만 할 뿐 그만 보내달라고 요청할 수 없었다.

	for n:=range Pipeline(Source()){
		if condition(n){
			// Done. I don`t need anymore.
			break
		}
	}


	위와 같이 채널에서 받다가 끝까지 받지 않고 중간에서 더 이상 자료가 필요 없어졌을 때 break문으로
	반복문을 빠져나오고 싶었을 것이다.  그런데 이렇게 되면 보내는 쪽에서 채널이 막히게 되고 이 고루틴은
	종료되지 않고 계속 막혀 있게 된다.

	싫더라도 채널이 닫힐 때까지 자료를 모두 빼 주어야 한다. 그래야 보내는 고루틴이 종료되면서 소멸된다.


	package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/chapter7"
)

func main() {
	c := make(chan int)
	go func() {
		defer close(c)

		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()

	nums := chapter7.PlusOne(chapter7.PlusOne(chapter7.PlusOne(chapter7.PlusOne(chapter7.PlusOne(c)))))

	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	for _ = range nums {
		// Consume all nums
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}


	코드에서는 3, 13, 23, 33, ... 순서로 숫자를 보낸다. 이전에 작성해두었던 PlusOne 파이프라인을 다섯 번 통과한다.
	그러니까 받는 숫자는 8, 18, 28 ... 순서로 나올 것이다. 받는 쪽에서는 이 숫자를 받아서 출력하다가 18이 나오면 더 이상
	볼 필요가 없어서 반복문을 빠져나와 버린다.

	그 상태에서 혹시나 다른 고루틴이 아직 빠져나오지 못했을 경우가 있으므로 약간의 여유를 위해 100밀리초 동안 Sleep을 하여 고루틴
	으로 제어를 넘겨준다. 그 뒤에 현재 고루틴의 수를 출력했다. 그 뒤에 강제로 nums에 남은 모든 자료를 받아준 다음에 같은 방식으로
	고루틴의 수를 한 번 더 측정했다.


	결과는 약간씩 달라질 수 있다.


			8
			18
			NumGoroutine:  7
			NumGoroutine:  1


	이러한 상황은 별로 좋은 상황이 아닐 것이다. 모두 자료를 소진시키지 않으면 해제되지 않은 고루틴들이 메모리에 남아 있음녀 메모리 누수가
	발생한다. 모든 자료들을 소진시킨다고 해도 좋지 않은 경우가 많은데 보내는 고루틴이 많은 네트워크 트래픽을 유발시키거나 배터리를 소모한다면
	계속해서 데이터를 받아오면 그만큼 더 많은 네트워크 트래픽과 배터리 소모가 발생할 것이기 때문이다.


	억지로 채널을 닫을 수도 없다. 그러면 닫힌 채널에 자료를 보내려고 하면서 패닉이 발생한다. 채널 닫기는 보내는 쪽에서만 하는 것으로 정형화해야 한다.


	이럴 때 유용한 패턴 중 하나가 done 채널을 하나 더 두는 것이다. 보내는 고루틴에서 이 채널로부터 신호가 감지되면 보내는 것을 중단하고
	채널을 닫으면 된다. 신호는 close(done)으로 주면 된다.

*/

func PlusOne_2(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		for num := range in {
			select {
			case out <- num + 1:
			case <-done:
				return
			}
		}
	}()
	return out
}

/*

package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/chapter7"
)

func main() {

	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()

	done := make(chan struct{})
	nums := chapter7.PlusOne_2(done, chapter7.PlusOne_2(done, chapter7.PlusOne_2(done, chapter7.PlusOne_2(done, chapter7.PlusOne_2(done, c)))))

	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}

	close(done) // 종료

	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	for _ = range nums {
		// Consume all nums
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

}


*/
