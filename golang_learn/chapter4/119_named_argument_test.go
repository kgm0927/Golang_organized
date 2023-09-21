package chapter4

import (
	"fmt"
	"io"
)

/*
	4.1.4 명명된 결과 인자

 지금까지 본 함수들은 함수가 넘겨받는 인자에 이름이 붙어 있지만 돌려주는 값에 이름 없이
 자료형만 나열하고 있다. 그러나 Go에서는 돌려주는 값들 역시 넘겨받는 인자와 같은 형태로 쓸 수
 있다. 넘겨받는 인자들은 넘어온 값들로 초기값이 설정되는데, 돌려주는 인자들은 기본값으로 초기화된다.
 정수면 0, 문자열이면 빈 문자열로 초기값이 설정된다.

 반환할 때에는 기존의 방식대로 결과값들을 return 뒤에 쉼표로 구분하여 나열할 수도 있고, 생략하고
 return만 쓸 수 있다. 생략한 경우에는 돌려주는 인자들의 값들이 반환된다.

*/

func writeTo(w io.Writer, lines []string) (n int64, err error) {
	for _, line := range lines {

		var nw int

		nw, err = fmt.Fprintln(w, line)
		n += int64(nw)
		if err != nil {
			return
		}
	}
	return
}
