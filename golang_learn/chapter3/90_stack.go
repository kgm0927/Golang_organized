package chapter3

import (
	"strconv"
	"strings"
)

func Eval(expr string) int {

	var ops []string
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string) {

		for len(ops) > 0 {

			op := ops[len(ops)-1]

			if strings.Index(higher, op) < 0 {
				// 목록에 없는 연산자이므로 종료
				return
			}

			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}

			b, a := pop(), pop()

			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}

		}
	}

	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)

		case "+", "-":
			// 덧셈과 뺄셈 이상의 우선순위를 가진 사칙연산 사용
			reduce("+-*/")
			ops = append(ops, token)

		case "*", "/":
			// 곱셈과 나눗셈 이상의 우선순위를 가진 것은 둘뿐
			reduce("*/")
			ops = append(ops, token)

		case ")":
			// 닫는 괄호는 여는 괄호까지 계산하고 제거
			reduce("+-*/(")

		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)

		}
	}
	reduce("+-*/")
	return nums[0]

}

/*
	3.2.8 스택

	스택은 후입선출(LIFO: Last In First Out) 구조를 갖고 있는 자료구조이다.


	Go 표준 라이브러리에 따로 스택 자료구조가 들어 있지는 않는다. 이미 슬라이스로도 충분히
	구현이 가능하기 때문이다. 스택에 집어 넣는 것은 append를 이용하면 되고, 스택에서 빼내는 것은
	마지막에 있는 것을 제거하는 것으로 가능하다.

	스택을 이용하여 사직 연산과 괄호를 이용한 간단한 정수 수식 계산기를 구현해본다. 한 개의 문자를
	구분하기 위해 strings.Split()을 이용하여 토크나이저를 구현할 수 있다.

	간단하게 구현하기 위해 함수 내부에 2개의 '함수 리터럴'을 정의하여 여러 번 호출하겠다. 함수 리터럴은
	처음 보는데, 이름을 안 쓰면 말 그대로 익명 함수가 된다. 이 함수를 변수에 보관하였다가 호출할 수 있는
	것이다. 이 함수들은 자신의 밖에 변수들에 접근할 수 있다. 이것은 '클로저'를 배울 때 다시 알아본다.

	한 함수는 pop()으로 숫자 스택에서 숫자를 하나 꺼내어 반환하는 일을 한다. 다른 함수 reduce() 함수는
	우선순위가 같거나 더 높은 연산자들의 목록을 받아서 해당 연산자들에 대하여 연산을 적용해주는 역할을 한다.

*/
