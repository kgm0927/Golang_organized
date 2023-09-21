package chapter3

/*
	맵을 사용하는 방법은 간단하다. 문자열 안에 있는 각 문자 수를 세는 함수를 작성해본다.


*/

func Count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

/*

	슬라이스와 다른 점은 맵을 이용할 때, 맵 변수 자체에 다시 할당하지 않으므로 포인터를 취하지 않아도 맵을
	변경할 수 있다. 그래서 'codeCount* map[rune]int'로 쓸 필요가 없다.

	물론 맵을 다른 맵으로 바꿔치기하고 싶으면 포인터를 넘겨야 한다.
*/
