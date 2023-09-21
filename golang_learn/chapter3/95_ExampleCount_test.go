package chapter3

import "fmt"

func ExampleCount() {
	codeCount := map[rune]int{}

	Count("가나다나", codeCount)

	for _, key := range []rune{'가', '나', '다'} {
		fmt.Println(string(key), codeCount[key])
	}
	// Output:
	// .
}

/*
	위의 방법은 지정한 키 외의 다른 키가 있는 경우 놓치게 된다.
	지정한 키가 맵에 없는 경우는 기본값으로 읽으므로 정수값인 위의
	경우에는 0이 나오게 해서 큰 문제는 없지만 반대의 경우 문제가 생긴다.
	즉 맵에서 '라'가 들어오면 테스트에는 잡아 낼 수 없다는 것이다.
	그리고 테스트에 실패했을 시 어떻게 잘못했는지 잘못을 판단할 수
	없다.




*/
