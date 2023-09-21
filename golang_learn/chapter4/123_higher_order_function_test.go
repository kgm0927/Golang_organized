package chapter4

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*

		4.2.2 고계 함수 (higher-order function)

	고계 함수는 곡예하듯이 함수를 주고 받는 것을 의미한다. 그것 뿐이다. 이름이
	고계 함수인 이유는 함수의 함수, 즉 함수를 넘기고 받는 함수이기 때문에 더 고차원적인
	함수라는 의미에서 붙인 것이다.


	여기서 ReadFrom 함수를 살펴본다. 이 함수는 파일에서 한 줄씩 읽어서 슬라이스에서 추가하고 있다.
	그런데 경우에 따라서 슬라이스에 추가할 것이 아니라 맵에서 추가하고 싶을 수도 있고, 네트워크를
	통하여 데이터를 보내고 싶을 수도 있을 것이다. 이때 고계 함수를 사용한다.

*/

func ReadFrom(r io.Reader, f func(line string)) error {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {

		return err
	}

	return nil

}

/*
	위의 함수는 이제 r에서 한 줄씩 읽어서 매 줄마다 f 함수를 호출한다. 호출자가 자유롭게 원하는 동작을
	지시할 수 있게 되었다. 아래 코드와 같이 ReadFrom을 호출할 수 있게 되었다.

*/

func Example_readFrom_Print() {
	r := strings.NewReader("bill\ntom\njane\n")

	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})

	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// ( bill )
	// ( tom )
	// ( jane )
}

/*
	두 번째 인자에서 함수 리터럴을 넣어서 호출하였다. 이 함수 리터럴은 양쪽에 괄호를 붙여서 출력하는 코드이다.
	그러면 원래 동작과 같이 슬라이스에 추가해 넣고 싶으면 어떻게 할까? 클로저를 살펴본다.

*/
