package chapter5

import (
	"fmt"
)

/*
		5.3.6 인터페이스 변환 스위치

	인터페이스들이 지정하는 범위는 다양할 수 있다. 예를 들어서 io.WriterAt, io.WriteSeeker, io.WriteCloser는
	모두 io.Writer를 내장하고 있는 인터페이스이다. 즉 이 세 인터페이스를 구현하고 있는 것은 io.Writer도 역시 구현하고
	있다는 것이다.

	그러면 이 중에서 가장 포괄적인 혹은 넓은 범위의 인터페이스는 io.Writer일 것이다. 궁극적으로는 비어 있는 인터페이스인
	interface{}는 모든 자료형을 포함하는 가장 넓은 인터페이스이다.

	인터페이스를 이용하여 해당 메서드들을 구현하기만 하면 받아들여서 사용하는 것을 했지만, 이젠 포괄적으로 인터페이스를 받아서 특정
	자료형일 때, 혹은 좀 더 좁은 범위의 인터페이스를 구현할 때 경우에 따라서 구현을 달리하고 싶다. 지료형 스위치(type switch)를
	이용하면 이것을 쉽게 할 수 있다.


	String.Join을 확장해서 구현해보자. 문자열 슬라이스를 받아서 구분자를 사이에 두고 각 문자열들을 연결시켜 준다.
	다음과 같이 동작하는 함수를 만들어 본다.


*/

func ExampleJoin() {
	t := Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: nil,
	}
	fmt.Println(Join(",", 1, "two", 3, t))
	// Output:
	// .
}

/*
	정수와 1과 3 문자열인 "two", 그리고 Task 모두를 연결시킬 수 있다. Join은 첫 번째 인자로 구분자를 받는다.
	위의 예제에서는 쉽표이다. 그 뒤에는 연결시키고 싶은 자료들을 나열할 수 있는데, 정수형, 문자열 그리고 String()
	메서드를 제공하는 다른 자료형들이 붙을 수 있다. 물론 fmt.Sprint와 같은 함수를 이용하면 이렇게 하지 않아도
	쉽게 구현할 수 있다. 하지만 이런 함수들을 이용하지 않고 구현하는 것을 목표로 해보자.


*/
