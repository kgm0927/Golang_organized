package chapter7

import "fmt"

/*
	이미 채널을 닫아 보았다. 채널이 닫히면 for range를 통하여 반복이 종료가 된다. 그 외 채널이 닫히면
	어떻게 되는지 더 자세히 알아본다.

	채널에서 값을 받을 때, <-c와 같은 형태를 이용한다는 것을 알고 있을 것이다. 아래에는 val에는 채널에서
	받은 값이 들어온다.

	val:=<-c

	만약 받는 쪽의 변수가 둘이라면 두 번째 변수에 채널이 열려 있는지 여부가 들어온다.

	val,ok:=<-c

	채널이 열린 상태라면 val에는 채널에서 받은 값이, ok에는 true 값이 넘어온다.

	닫혀 있을 경우 비어 있으며, ok에는 false 값이 들어온다.
*/

func Example_closedChannel() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 0
	// 0
	// 0

}

/*
	이미 닫은 채널을 또 닫으면 어떻게 될까? 이 경우에는 패닉이 발생한다.
	이런 닫힌 채널의 성질을 잘 알고 있으면 여러 패턴에 사용할 수 있다. 이런
	패턴들은 뒤에 알아볼 것이다.
*/
