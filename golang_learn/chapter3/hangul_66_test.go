package chapter3

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	// Output:
	// .

}

/*

메인 함수에 쓸 법한 내용을 이름만 Example...로 바꾸어서 만들었다. 밑에 // Output:이라고 써두었고,
함수 이름은 반드시 Example로 시작해야 한다.

여기서 go test hangul_test.go hangul.go를 수행한다. go test는 go run과 유사하지만 테스트를 수행한다.
혹은 주소를 적어도 된다.const


*/
