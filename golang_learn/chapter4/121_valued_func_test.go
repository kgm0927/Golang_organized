package chapter4

import "fmt"

/*

	4.2 값으로 취급되는 함수

	Go언어에서 함수는 일급 시민(First-class citizen)으로 구분이 된다. 함수 역시 값으로 변수에
	담길 수 있고 다른 함수로 옮기거나 돌려받을 수 있다는 것이다. 어떤 언어에서는 함수가 이급 취급을
	받아서, 일급취급을 받는 정수값과 같은 것들을 변수에 담고 함수로 넘기로 받을 수 있지만 함수는
	이렇게 할 수 없다.



	4.2.1 함수 리터럴


	지금까지 함수를 선언할 때 함수 이름, 그 함수의 자료형, 그리고 코드를 작성했다. 마치 함수의 이름은
	함수의 값을 담는 변수와 같이 보인다.


*/

func add(a, b int) int {
	return a + b
}

/*

	그러면 순수하게 함수의 값만 표현하려면 어떻게 해야 하는가? 그냥 이름을 빼면 되지 않을 까?


	func () int {
		return a+b
	}

	그렇다. 이제 이 함수는 add라는 이름이 없어진 순수한 값이 되었다. 이것을 함수 리터럴(Functional literal)이라고
	부르고, 익명 함수라고 부르고, 익명 함수라고 부를 수도 있다고 익히 말했다. 함수형 언어에서 람다 함수와 동일한 방법으로 사용할
	수 있다. 이것을 변수에 담고 다른 함수에 넘겨 주고, 돌려받고 할 수 있다.

	아래의 함수는 인자와 반환값이 없는 함수이다.
*/

func printtHello() {
	fmt.Println("Hello!")
}

// 이제 여기서 이 함수의 이름을 없애도 함수 리터럴만들 남긴 채 다음에 호출한다.

func Example_funcLiteral() {
	func() {
		fmt.Println("Hello")
	}()
	// Output:
	// Hello!

}

/*
	func 안에 들어 있는 것이 '함수 리터럴'이라고 한다. 그 뒤에 괄호를 여닫은 ()이 해당 함수를
	호출한 것이다. 다음은 이 함수를 변수에 담은 다음 이용해본다.

*/

func Example_funcLiteralVar() {
	printtHello := func() {
		fmt.Println("Hello")
	}
	printtHello()
	// Output:
	// Hello!
}

/*
	함수 리터럴을 printHello에 담았다. 그리고 5번째 줄에서 printHello 변수에 담긴 함수를 그대로 호출하였다.
	70번째 줄에 printHello가 있는 위치에 그것이 담겨 있는 함수값을 그대로 갖다 붙이면 아까 전에 본 코드의 모양이
	된다.

	다른 함수형 언어의 '람다'와 비슷하다.

*/
