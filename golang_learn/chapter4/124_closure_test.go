package chapter4

import (
	"fmt"
	"strings"
)

/*
		4.2.3 클로저

	닫힘이라는 의미의 클로저(closure)이다.

	클로저는 외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드를 의미한다.


*/

func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string

	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}

/*
	바로 전 예제에 대한 답은 이것이다. 이렇게 해야 우리가 새로 만든 ReadFrom 함수에서 슬라이스에
	읽은 줄들을 첨가할 수 있다. 이제 비밀이 풀렸다.


	ReadFrom에서 넘겨받는 함수 f가 line string 하나만 받을 수 있는데, 이것으로 대체 무엇을 할 수 있는가?

	위에 보면 lines라는 문자열 슬라이스를 ReadFrom에 넘겨주는 함수 리터럴 안에서 사용하고 있다. ReadFrom에
	넘기는 함수는 그 함수가 이용하는 변수들도 함께 가지고 넘어간다.


*/
