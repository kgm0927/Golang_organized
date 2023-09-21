package chapter4

/*

	4.3.3 포인터 리시버

	위의 메서드들은 모두 포인터 리시버(Pointer receiver)가 아니다. 포인터 리시버는 자료형이 포인터형인
	리시버이다. 리시버 역시 함수의 다른 인자들과 같이 같이 값으로 전달된다.
	포인터로 전달해야 할 경우에는 포인터 리시버를 사용한다.

	3장에서 그래프의 인접 리스트를 파일에서 읽어오고 파일로 쓰는 것을 살펴보았다. 자료형이 이름을 붙이기
	전까지는 단지 [][]int로 표현되는 이중 슬라이스였다.


*/

type Graph [][]int

/*
	메서드를 정의하기에 앞서 지난 3장에서 각 함수들의 자료형을 본다.

	func WriteTo(w io.Writer, adjList [][]int) error
	func WriteTo(w io.Writer, adjList [][]int) error

	단, Go 언어의 관습상 리시버의 이름을 길게 하지 않는다.

*/
