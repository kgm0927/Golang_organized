package chapter4

import (
	"fmt"
	"strconv"
	"strings"
)

/*
4.2.9 자료구조에 담은 함수

Go에서 함수가 일급 시민이라고 말했다. 일급 시민이라 하면 변수를 담을 수 있고 함수에
넘겨주고 반환받을 수 있다. 한가지 더 중요한 점은 바로 자료 구조에 마음대로 담을 수 있다는 점이다.

3장에서 스택을 배울 때, 계산기를 코드에 작성했다. 이미 그 코드에서 클로저까지 사용하고 있었다.

이번엔 이런 자료들을 모아서 자료구조에 함수를 담아서 활용하는 방법을 이용해본다. Eval이라는 함수는
사칙연산만 가능한다. 제곱, 나머지 연산도 추가하려고 한다. 그리고 사칙연산의 정의도 바꾸도 싶다.
이 연산들을 '맵'으로 넘겨 맵의 키로 연산자를 두고, '그 값을 함수'를 준다.
*/
type PrecMap map[string]StrSet

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {

	ops := []string{"("}
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]

			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				// 더 낮은 순위 연산자이므로 여기서 계산 종료
				return
			}

			ops = ops[:len(ops)-1]

			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}

			b, a := pop(), pop()

			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {

		if token == "(" {

			ops = append(ops, token)

		} else if _, ok := prec[token]; ok {

			reduce(token)
			ops = append(ops, token)

		} else if token == ")" {
			// 닫는 괄호는 여는 괄호까지 계산하고 제거
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}

	}
	reduce(")") // 초기의 여는 괄호까지 모두 계산
	return nums[0]

}

/*
		길어서 생략한 부분이 있지만, 원래 switch-case 문이 있는 곳에 opMap을 접근하여 함수를 가져와서
		적용하도록 변경되었다. 이 함수를 호출할 때눈 다음과 같이 정의된 opMap을 호출시에 넘겨주어야 한다.
		아직은 코드 구조상 사칙연산 이외에는 추가할 수 없다.

		이제 우선 순위 처리가 문제이다. 이전에는 사칙 연산만 있었기 때문에 하드코드했지만, 이제는 어렵게 된다.
		그리고 reduce 함수가 한 글자짜리 연산자만 처리할 수 있는 것도 문제다.
		연산자를 구조체를 처리하면 더 깔끔해지겠지만 지금은 배운 범위 내에서 맵을 하나 더 받게 해본다. 이 맵의
		키는 연산자 하나, 값은 자신보다 더 높은 우선순위의 연산자 집합으로 해본다.

	3장에서 맵을 공부하면서 집합을 처리하는 패턴을 배웠는데, 4장에서 가변인자를 배웠으므로 다음과 같이 집합을 생성하는
	함수를 만들어보자.
*/

// string set type
type StrSet map[string]struct{}

// Returns a new StrSet.

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

/*
	그러면 이제 NewStrSet("a","b","c")와 같이 편리하게 문자열 집합을 생성할 수 있게
	되었다. 이 문자열 집합을 이용하여 우선순위 맵의 자료형을 정의해본다.

	// Map keyed by operator to set of higher precedence operators
	type PrecMap map[string]StrSet
	(이 타입의 맵은 맨 위에 구현시켰다. 그래야 코드를 실행할 수 있다.)




	이제 다음과 같이 두 개의 맵을 받아야 한다.

	func Eval(opMap map[string]BinOp,prec PrecMap ,expr string) int {


	그렇지만 사용하기가 매우 번거로워 보이는데, 매번 Eval 함수를 호출할 때마다 넘겨주기 어려울 것
	같지만, 다음과 같은 함수를 정의해 주면 된다. 다음 함수는 두 개 받아서 계산기를 돌려주는 함수이다.
*/

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {

	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}

}

/*
	이미 배운 인자 고정 패턴을 이용하였다. 다음과 같이 eval를 생성하여 반복해서 사용할 수 있다.
*/

func Example_newEvaluator() {
	eval := NewEvaluator(map[string]BinOp{
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}

			if b < 0 {
				return 0
			}

			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}

			return r
		},
		"*":   func(a, b int) int { return a * b },
		"/":   func(a, b int) int { return a / b },
		"mod": func(a, b int) int { return a % b },
		"+":   func(a, b int) int { return a + b },
		"-":   func(a, b int) int { return a - b },
	}, PrecMap{
		"**":  NewStrSet(),
		"*":   NewStrSet("**", "*", "/", "mod"),
		"/":   NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+":   NewStrSet("**", "*", "/", "mod", "+", "-"),
		"-":   NewStrSet("**", "*", "/", "mod", "+", "-"),
	})

	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1 - 2 - 4"))

}

/*

	별표 둘이 붙어 있는 **연산자를 제곱 연산자로 정의하였다. mod를 나머지 연산자로 추가하였다.
	PrecMap을 보면 각각의 연산자에 대하여 자신보다 높은 순위의 연산자를 나열해두었다. 다른 연산자들이
	스스로를 자신보다 높은 우선 순위의 연산자로 설정하고 있지만, **의 경우에는 그렇지 않다. 이유는
	**연산자가 오른쪽 결합을 하기 때문이다. 마지막 예제의 2**2**3의 경우에 2**(2**3) 순서로 계산하여야 하고
	1-2-4의 경우에는 순서로 계산을 해야 하기 때문에 이런 차이가 생긴다.

*/
