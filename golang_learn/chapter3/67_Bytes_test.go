package chapter3

import "fmt"

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

/*
문자열을 어떻게 반복문에서 사용하는지에 따라서 유니코드 문자 단위로 동작하기도 하고, 바이트 단위로 동작하기도 한다.


여기서 새롭게 등장한 것이 바로 fmt.Printf()이다. Printf의 마지막 글자 f는 formatted의 머리글자이다.
%x는 16진수의 숫자 형식으로 출력한다.
*/
