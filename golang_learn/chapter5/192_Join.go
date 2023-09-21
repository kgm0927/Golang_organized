package chapter5

import (
	"fmt"
	"strconv"
	"strings"
)

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}

	t := make([]string, len(a))

	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String() // 180_custom_printer_test.go 참고 (이 기능은 json의 역직렬화와 관련이 있다. 이걸 배우지 않았으므로 넘긴다.)
		}
	}
	return strings.Join(t, sep)
}

/*
	위와 같이 구현이 된다. 여기서 눈여겨볼 부분은 switch문이다. 마치 a[i]를 type으로 형 단언을 해서 x에
	할당한 모양새이다. 특수한 구문으로 이렇게 하면 case에서 자료형의 이름을 지정하여 각각의 경우에 대하여 구현을 다르게 해줄 수
	있다.

	이때 case 내부에서는 해당 자료형의 값으로 x가 지정이 된다. 코드는 t를 a와 같은 문자열 슬라이스로 만들어준 다음에 문자열로 변환하여
	복사한 뒤 strings.Join을 호출하는 일을 한다. 눈여겨볼 부분인 switch 내부에서는 문자열일 때는 그냥 t를 복사해주고, 정수형일 경우는
	strconv.Itoa를 호출하여 문자열로 바꾸어서 넣어주었다. fmt.Stringer 인터페이스를 구현하면 String 메서드를 호출하여 넣어준다.


	switch에서 case로 받는 부분에서 string과 string과 int와 같은 구체적인 자료형을 써주어도 괜찮고, fmt.Stringer와 같은 인터페이스를 써주어도 좋으며
	서로 이렇게 섞어 써도 괜찮다.

	사실 위의 switch문은 형 단언을 이용해도 구할 수 있다.


	if x,ok:=a[i].(string);ok{
		t[i]=x
	}
	else if x,ok:=a[i].(int);ok{
		t[i]=strconv.Itoa(x)
	}
	else if x,ok:=a[i].(fmt.Stringer);ok{
		t[i]=x.String()
	}


	위와 같이 작성하여도 switch 문과 동일한 효과를 낼 수 있다. 그러나 switch 문을 이용하는 것이 보기에 좋고 더
	편리하기 때문에 이것을 권장한다.


*/
