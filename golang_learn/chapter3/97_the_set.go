package chapter3

/*
	사실 Go 언어에서 상수 시간에 키의 존재를 확인할 수 있는 집합은 따로 있지 않다.

	가장 단순한 방법은 맵을 이용하면서 값을 bool형으로 주는 것이다. 이번에는 문자열에
	중복되어 들어가 있는 글자가 있는지 검사하는 함수를 만들어본다.

*/

func hasDupeRune(s string) bool {
	runeSet := map[rune]bool{}
	for _, r := range s {
		if runeSet[r] {
			return true
		}
		runeSet[r] = true
	}
	return false
}

/*

	위의 함수는 깔끔하지만 불필요한 bool 값 때문에 메모리를 많이 차지한다.
	이걸 해결하려면 빈 구조체를 값으로 사용하면 된다. 이렇게 하면 값 부분의
	메모리를 따로 차지하지 않는다. 오버헤드가 없다.



*/
