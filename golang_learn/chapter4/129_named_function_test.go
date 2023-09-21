package chapter4

import "fmt"

/*
	4.2.6 명명된 함수형

	Go언어에서는 함수는 일급 시민으로 분류된다고 했다. 그러면 당연히 함수의 자료형
	역시 사용자가 정의할 수 있다.

*/

type BinOp func(int, int) int

/*
	두 정수를 넘겨받아서 정수 하나를 반환하는 함수형을 BinOp형으로 정의하였다. 그러면 명명된
	함수형도 자료형 검사를 하나?

	한다. 다음과 같이 BinOp를 받는 함수가 있다고 한다.
*/

func OpThreeAndFour(f BinOp) {
	fmt.Println(f(3, 4))
}

/*
	이때 함수를 호출할 때, 이렇게 한다.

	OpThreeAndFour(func(a,b int)int{
		return a+b
	})


	OpThreeAndFour는 BinOp형을 인자로 받는데, 호출 시에는 그냥 정수 둘을 받고 정수 하나를 반환하는 함수를 넘겨줬었다.
	이는 컴파일 오류를 발생시키지 않는다.

	이는 func(a,b int)int 자료형이 명명되지 않는 자료형이기 때문이다. 양쪽 모두 명명된 자료형이 아니면 서로간에 호환이 된다.


	위의 명명된 함수형은 BinOp와 동일한 표현형으로 되어 있다. 그러나 BinOp가 두 정수를 이용하여 연산한 결과를 반환하는 함수인데
	반해, BinSub의 경우에는 두 정수를 받아서 어떤 연산을 수행하고 그 결과값으로는 해당 함수가 몇 번 호출했는지는 반환하는 함수로
	의도하였다고 하자.

	따라서 BinOp와는 다른 의미이기 때문에 이렇게 다른 이름을 붙여서 선언하였다.

	다시 말해 BinOp는 같은 두 수를 넘기면 같은 반환값이 돌아오는 순수 함수(Pure function)이지만 BinSub는 그렇지 않다.
*/

type BinSub func(int, int) int

func BinOpBinSub(f BinOp) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}

/*
	위에서 정의한 BinOpToBinSub는 BinOp 함수를 받아서 BinSub 함수를 반환하는 함수이다.
	다음과 같이 사용할 수 있다.
*/

func Example_binOpToBinSub() {
	sub := BinOpBinSub(func(i1, i2 int) int {
		return i1 + i2
	})

	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	// Output:
	// 12
	// 12
	// 12
	// count: 3
}

/*
	BinOpToBinSub을 호출하여 sub은 BinSub 자료형이 되었다. 만약에 다시 BinOpToBinSub(sub)와 같이
	호출하거나 위의 코드에서 BinOpToBinSub로 두 번 둘러싸면 어떻게 될까?

	sub:=BinOpToBinSub(BinOpToBinSub(func(a,b int)int{

	}))

	이런 식으로 할 시 컴파일하면 오류가 발생한다. 즉, 함수 리터럴과 명명된 함수형 사이에는 자동으로 형변환이
	일어나지만 '명명된 함수형 사이에는 자동으로 형변환이 일어나지 않는다.'
*/
