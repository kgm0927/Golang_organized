package chapter7

import "fmt"

/*
	7.3 동시성 패턴


	두 고루틴에서 한쪽에서는 보내고 다른 쪽에서는 받기만 하는 패턴을 이미 배웠다. 이제 다른
	패턴들을 만나보겠다.const

	7.3.1 파이프라인 패턴

	파이프라인은 한 단계의 출력이 다음 단계의 입력으로 이어지는 구조이다. 서로 다른 단계는 동시에 이루어질 수
	있기 때문에 성능상의 장점이 있을 수 있다.

	각자 하는 일을 반복함으로써 전문성이 생기고 분업의 효과가 있다.

	컴퓨터 시스템에서는 특히 서로 다른 종류의 하드웨어들이 어떤 일을 해야 할 때 파이프라인이 큰 효과가 있다.
	소프트웨어에서는 들어오는 데이터와 나가는 데이터에 집중하여 문제를 풀고자 할 때 장점이 있고, 버퍼를 활용하면
	경우에 따라 성능상의 장점도 얻을 수 있다.


	파이프라인 패턴은 생성기 패턴의 일종이다. 생성기 패턴과 동일하고 받기 전용 채널을 반환한다. 그러나 받기 전용 채널을
	넘겨받아서 입력으로 활용한다는 점에서 차이가 있다. 반환된 받기 전용 채널을 다른 파이프라인의 입력으로 넘겨줄 수 있기
	때문에 매우 자연스럽게 출력을 입력으로 연결하여 일직선으로 사슬처럼 연결된 파이프라인을 구성할 수 있다.

*/

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // 채널을 닫음
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)

	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

/*
	PlusOne은 받기 전용 채널을 받아서 다른 받기 전용 채널을 돌려주는 함수이다. 받느 채널에서
	숫자를 하나 증가시켜서 보내준다.

	위의 예에서는 PlusOne을 연달아 두 번 사용하자 집어늫은 수에서 2가 더해진 숫자들이 결과로
	나왔다.

	같은 함수뿐만 아니라 형태만 같다면 서로 다른 함수들도 이렇게 이어붙일 수 있다. 아래와 같이
	일직선 파이프라인을 구성할 수 있다. IntPipe라고 이름을 붙이겠다. 함수 형태에 이름을 붙이면 간단
	하게 할 수 있다.



*/

type IntPipe func(<-chan int) <-chan int

/*
	생성기 패턴과 마찬가지로 데이터를 보내는 쪽에서 채널 닫아야 한다는 사실을 명심해야 한다.
	그렇게 하지 않으면 파이프라인이 꼬여 버린다. 패턴을 외워두는 것이 좋다.
*/

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}

}

/*
	이제 Chain을 이용하여 PlusOne(PlusOne(c))를 Chain(PlusOne,PlusOne)(c)와 같이 표현할 수 있다.
	서로 다른 함수로 구분하여 본다면 Chain(A,B)(c)는 B(A(c))와 같다. 채널 c에 보내진 자료들을 A와
	B의 순서로 넘겨받게 된다.

	PlusTwo:=Chain(PlusOne,PlusOne)

	바로 이용하지 않고 다른 곳에 넘겨야 하는 경우에 Chain 고계 함수를 이용하면 편리하다.






*/
