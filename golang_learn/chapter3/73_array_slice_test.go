package chapter3

/*
3.2 배열과 슬라이스

3.2.1 배열



배열과 슬라이스에 대하여 알아본다. 배열과 슬라이스 모두 연속된 메모리 공간을 순차적으로
	이용하는 자료 구조인데, 배열이 직접 사용되는 경우도 있지만 주로 슬라이스를 사용하여 간접적으로
	배열을 이용하는 경우가 많다.*/

import "fmt"

func Example_array() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
	}
	// Output:
	// .
}

/*
위의 설명은 그냥 생략하겠다.

컴파일러가 배열의 개수를 알아내서 넣게 만들고 싶으면 [3] 대신에 [...]을 이용해도 좋다.

fruits=[...]string{"사과","바나나","토마토"}


*/
