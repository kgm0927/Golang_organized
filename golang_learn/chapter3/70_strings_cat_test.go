package chapter3

import "fmt"

/*
문자열은 읽기 전용이다. 그러면 어떻게 문자열을 이어붙인다. 사실 문자열을 이어붙이는 것이 문자열을
수정하는 것이 아니다. 두 문자열을 이어붙인 문자열을 새로 만드는 것이다. 문자열 수정이 아니다.

문자열을 이어붙이려면 + 연산을 이용하면 된다. Go의 문자열은 사실상 문자열에 대한 포인터와 비슷하기 때문에
다음과 같은 역할을 한다.

*/

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

/*
ps는 s의 주소값을 취한 포인터 형식인 *string형이다. 그리고 다시 s에 다시 "def"를 뒤에 이어 붙였다. ps에서
값이 변경이 일어난다.


fmt패키지에 있는 S로 시작하는 함수들로도 이어붙이기를 할 수 있다. 즉, 위의 s += "def" 대신에 s=fmt.Sprint(s,"def")
혹은 s=fmt.Sprintf("%sdef",s)와 같이 이용하여 괜찮다는 것이다. 물론 s=strings.Join([]string{s,"def"},"")같은 코드를
이용할 수 있다.

(Sprint 계열은 반드시 문자열이 아닌 것들도 마치 다른 Print 계열처럼 이어붙일 수 있다.)

문자열 슬라이스나 배열이라면 strings.join을 이용할 수 있고, 특히 구분자를 이용하여 이어붙일 때는 이것을 이용하면 편리하다.


fmt, strings, bytes, strconv에 있는 함수들을 좀 살펴보면 재미있을 것이다. strings와 bytes는 각각 string형과
[]bytes형에 대한 함수들이 서로 대응되어 구현되어 있다.
*/
