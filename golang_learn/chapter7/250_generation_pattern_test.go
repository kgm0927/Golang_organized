package chapter7

import "fmt"

/*
	4장에서 배운 생성기를 채널을 이용하여 만들 수 있다. 아래에서 Fibonacci 함수를 방금배운 패턴대로 받아가기만 할 수 있는
	채널을 반환한다.

*/

// 피보나치 수열을 max까지 생성한다.

func Fibonacci(max int) <-chan int { // 받기 전용
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ", ")
	}
	// Output:
	// .
}

/*
	같은 생성기를 클로저를 이용하여 만들었다면 다음과 같은 모양이 된다.
*/

func FibonacciGenerator(max int) func() int {

	next, a, b := 0, 0, 1

	return func() int {
		next, a, b = a, b, a+b

		if next > max {
			return -1
		}
		return next
	}

}

func ExampleFibonacciGenerator() {
	fib := FibonacciGenerator(15)
	for n := fib(); n >= 0; n = fib() {
		fmt.Println(n, ",")
	}
	// Output: 0,1,1,2,3,5,8,13
}

/*
	비슷해보이지만 채널을 이용하는 방법에는 몇 가지 장점이 있다.

	1. 생성하는 쪽에서는 상태 저장 방법을 복잡하게 고민할 필요가 없다.
	2. 받는 쪽에서는 for의 range를 이용할 수 있다.
	3. 채널 버퍼를 이용하면 멀티 코어를 이용하거나 입출력 성능상의 장점을 이용할 수 있다.

*/

/*
 첫 번째 글자와 두 번째 글자의 후보들을 주면 가능한 모든 경우의 수를 생성하는 프로그램이다.
*/

func BabyNames(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func ExampleBabyNames() {
	for n := range BabyNames("성저명재경", "준호우훈진") {
		fmt.Println(n, ", ")
	}
	// Output: ...
}

/*
	이것을 클로저를 이용한 생성기로 구현한다면 조금 까다롭게 된다. 그러나 콜백 함수를 받아서 동작하게
	만든다면 위의 방법과 거의 동일하게 구현이 가능하다.

	생성기 패턴은 앞으로 자주 이용이 된다. 받기 전용 채널을 반환하는 것으로 이용하는 입장에서 깔끔하기
	때문에 앞으로도 계속 같은 패턴을 사용할 것이다.
*/
