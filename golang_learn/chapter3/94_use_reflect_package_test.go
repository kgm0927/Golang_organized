package chapter3

import (
	"reflect"
	"testing"
)

/*

	안타깝지만 맵끼리 ==로 비교한다고 해서 내용을 비교해주지 않고 단순히 fmt.Println으로 출력하더라도
	매번 다른 내용을 출력되기 때문에 정형화된 패턴 없이 테스트하는 경우가 많아서 실수하는 경우가 매우 많다.
*/

func TestCount(t *testing.T) {
	codeCount := map[rune]int{}

	Count("가나다나", codeCount)
	if !reflect.DeepEqual(map[rune]int{'가': 1, '나': 2, '다': 1}, codeCount) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

/*

위의 값은 ok가 나옴

	다른 방법은 다음과 같이 테스트를 작성하는 방법이다. 맵의 크기와 각각의 키와 값들을 모두 비교했다.

	if codeCount['가'] != 1 || codeCount['나'] != 2 || codeCount['다'] != 1 {
		t.Error("codeCount mismatch:", codeCount)
	}
*/
