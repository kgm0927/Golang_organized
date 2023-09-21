package chapter5

import "fmt"

/*

	5.3.2 커스텀 프린터


	이름 붙인 자료형은 Print 계열의 함수들을 이용하여 출력할 때 나타나는 형식을 정의할 수 있다.
	결과적으로 Print는 문자열이 있어야 출력할 수 있다. 정수형이나 실수형도 어떻게 나타낼지는 문자열로
	나타내지 않으면 출력할 수 없다. 가령 17이라는 정수를 아라비아 숫자로 출력할 것인지 로마 숫자로
	출력할 것인지, 10진수로 출력할 것인지 8진수로 출력할 것인지 등에 따라서 다른 방식으로 출력해야 한다.
	구조체도 어떤 형식이로든 문자열로 변환하지 않으면 출력할 수 없다.

	Print 계열 함수들이 문자열이 아닌 자료들을 출력할 때 미리 정해진 방식을 이용하는데 이것을 설정해주고 싶으면
	String() 함수를 정의해주면 된다. 실제로 fmt.Stringer이라는 인터페이스는 func String() string 메서드를 가지고 있다.
*/

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

/*
	Stringer를 받는 간단한 프린트 함수는 다음과 같다.

*/

func PrintStringer(data fmt.Stringer) {
	fmt.Print(data.String())
}

/*
	이제 String() 메서드를 구현하고 있는 모든 이름 붙은 자료형들을 출력할 수 있다.
	PrintStringer에는 Task 자료형도 넘길 수 있고 다른 자료형도 넘길 수 없다.

	그러나 위와 같은 함수는 만들 필요가 없다. 이미 fmt.Print와 같은 함수들이 이것을 이용하고
	있기 때문에 다른 자료형들을 출력하는 것과 같은 방법으로 출력할 수 있다.

	이렇게 되는 이유는 fmt.Println 함수가 인터페이스를 검사해서 Stringer인 경우에는 String 메서드를
	호출하여 출력하기 때문이다.
*/

func Exampltask_String() {
	fmt.Println(Task{"Laundry", DONE, nil})
	// Output: [v] Laundry <nil>
}

/*
	자료형 변환을 이용하면 다른 구현의 출력을 하게 만들 수 있다. 이번에 우리는 하위 개념을 만들어보려고 한다.
	(이 개념은 나중에 다시 넘긴다.)

*/
