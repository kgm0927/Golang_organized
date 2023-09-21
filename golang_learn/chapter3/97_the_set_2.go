package chapter3

/*
	문제는 특정 원소가 있는지 없는지는 어떻게 알까? 키가 있거나 없거나 맵에서 읽으면 빈 구조체가
	나오는데 말이다. 맵을 읽을 때 값을 2개로 받을 수 있다. 그때 두 번째 반환값을 조사하면 있는지
	없는지 알 수 있다.

*/

func hasDupeRune_2(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false
}

/*
	일단 struct{}는 아무런 필드가 없는 구조체이므로 값이 아니라 자료형이다. 맵을 만들 때 맵의
	자료형은 map[rune]struct가 되는 것이 아니라 map[rune]struct{}가 되는 것이므로 이 맵에 아무것도 안 들어
	있는 상태로 초기화를 하려면 위와 같이 map[rune]struct{}{}를 해 주어야 한다.

	struct{}는 빈 구조체 자료형이므로 해당 구조체에 아무 내용이 없는 사태의 값을 표현하려면 struct{}{}를
	써서 runeSet=struct{}{}와 같이 써야 한다.

*/
