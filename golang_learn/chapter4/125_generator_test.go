package chapter4

import "fmt"

func NewIntGenerator() func() int {
	var next int

	return func() int {
		next++
		return next
	}
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()

	fmt.Println(gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen())
	// Output:
	// .
}

// 그냥 NewIntGenerator를 사용할 경우 에러가 뜬다.

/*

	NewIntGenerator는 함수를 반환하는 고계함수이다. 반환값의 자료형을 살펴보면 func()int로 되어 있으니
	정수를 반환하는 함수를 반환하는 함수가 바로 NewIntGenerator이다. 게다가 이 함수가 반환하는 함수는 '클로저'
	이다.

	반환하는 함수 리터럴이 속해 있는 스코프 안에 있는 next 변수와 함께 세트로 묶인다. 만약 NewIntGenerator()를
	여러 번 호출하여 함수 여러 개를 갖고 있다면 각각의 함수가 갖고 있는 next는 따로 분리되어 있다.


*/

func ExampleNewIntGenerator_multiple() {
	gen1 := NewIntGenerator()
	gen2 := NewIntGenerator()

	fmt.Println(gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1(), gen1())
	// Output:
	// .
}

/*
	위의 코드에서 gen1과 gen2에 세트로 묶어 있는 next는 서로 다르다. 그렇기 때문에 gen1을 호출하여 증가시킨 숫자와
	gen2를 호출하여 증가시킨 숫자는 서로 다른다. gen1, gen2 두 함수의 상태가 분리되어 있는 것이다.

	같은 방식을 이용하여 느긋한 계산법(Lazy evaluatioin)을 구현하거나 무한한 크기의 자료구조를 만들 수 있다.

*/
