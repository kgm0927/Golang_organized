package chapter4

import "fmt"

/*
	4.3.1 단순 자료형 메서드

	OOP(객체지향 프로그래밍, Object Oriented Programming)를 지원하는 많은 언어가 메서드를 지원한다.
	그러나 몇몇 언어는 클래스로 정의된 자료형만 메서드를 지원하기도 한다. 모든 명명된 자료형에서 메서드를
	정의할 수 있다.
*/

func Example_vertexID_print() {
	i := VertexID(100)
	fmt.Println(i)
	// Output:
	// VertexID(100)
}

/*
	i는 정수형이 아니라 VertexID형이지만 화면에 출력하면 정수형과 마찬가지로 출력이 된다.
	그러나 VertexID는 서로 사칙연산 할 수 있는 정수형과는 다르므로 다른 형태로 출력하고 싶어지면
	어떻게 하면 될까?


*/

func (id VertexID) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}

/*
	이와 같은 간단한 코드로 예제 코드의 결과가 바뀌었다. 코드 부분은 동일하고 출력 부분만 다르다.

*/

func Example_vertexID_String() {
	i := VertexID(100)
	fmt.Println(i)
	// Output:
	// VertexID(100)
}

/*
	메서드의 정의를 살펴봅니다. 함수 정의와 비슷한 형태이지만 함수 이름 앞에 인자가 하나 더 붙어
	있다. (id VertexID)가 바로 그것이다. i가 VertexID 자료형이면 i.String()과 같이 메서드를 호출할
	수 있다.

	그러면 fmt.Println에 i.String을 넘겨 준 것도 아닌데, 왜 이런 결과가 나왔을까요? 이것은 Go의 인터페이스
	라는 기능 때문이다. 5장에서 인터페이스를 본격적으로 다룬다.

*/
