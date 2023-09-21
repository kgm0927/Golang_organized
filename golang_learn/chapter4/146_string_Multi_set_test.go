package chapter4

import "strings"

/*
	4.3.2 문자열 다중 집합

	3장의 연습문제 중에 MultiSet을 구현하는 것이 있었다. 메서드 때문에
	이 MultiSet을 좀 더 우아하게 구현할 수 있게 되었다.

	먼저 자료형에 이름을 붙여야 한다. 그래야 메서드를 정의할 수 있으며, 문자열 다중 집합을
	만들 것이다. 따라서 맵의 키는 집합의 원소인 '문자열'이고, 값은 해당 원소가 몇 번 반복되는지를
	표시하는 '정수'이다.

type MultiSet map[string]int
(132_argument_fix_test.go에서 이미 작성이 되어 있으므로 따로 만들지 않음.)

*/

func (m MultiSet) Insert(val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {

	s := "{"
	for val, count := range m {

		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}

/*
	훨씬 깔끔해졌다.

	단지 정수에 불과했던 자료가 ID가 되고 단지 map[string]int에 지나지 않던 자료형에 이름을 붙이자
	MultiSet이라는 추상 자료형이 된 것이다. 철저히 MultiSet에 집중하여 이 자료를 다루는 메소드들을
	정의한 것이다.

	프로그램 내에서 map[string]int로 표현되는 자료형은 다른 곳에서도 쓰일 것이다. 그러나 이름을 붙임으로써
	다른 것들돠 구별되게 되었다. map[string]int로 표현되는 명명된 자료형이 여러 개일 수 있으며 이들은 모두 다른
	메서드들을 가질 수 있다.

*/
