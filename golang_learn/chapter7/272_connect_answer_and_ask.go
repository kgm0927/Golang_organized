package chapter7

/*
	7.3.8 요청과 응답 짝짓기

	요청을 한 채널에 보내고 응답을 다른 채널로 받는 방식으로 파이프라인을 동작시킬 수 있을
	것이다. PlusOne의 경우에는 넘겨준 채널에 1을 넣으면 반환받은 채널에서 2를 받을 수 있다.

	그런데 고민이 될 수 있는 점이 있다. 응답을 받았을 때 이것이 어느 요청에 의한 응답인지 알아야
	하는 경우가 있다. 특히 분산 처리 되면 어느 것이 먼저 나올지 알 수가 없다. 알 필요가 없는 경우도
	있지만 알아야 하는 경우도 있다.

	한가지 방법은 채널로 자료를 넘겨주고 받을 때 ID 번호를 같이 넘겨서 ID 번호를 확인해보는 것이다.
	지금까지의 예제에서는 편의상 chan int 자료형으로 채널을 만들고 있는데 실무에서는 정수를 주고받기
	보다는 chan Msg와 같이 직접 정의한 자료형을 넘겨주고 받는 경우가 더 많을 것이다. 그때 그것이
	구조체라면 그 안에 고유한 요청 ID를 포함시켜두면 된다.

	그러나 보내는 쪽에서 요청 ID를 보관하고 있지만 이 요청에 대한 응답을 다른 고루틴이 받아 갈 수 있다면
	이것 역시 골치아파진다. 따라서 내가 보낸 요청에 대한 응답을 확실히 받기 위해서는 약간의 테크닉이 필요하다.
	요청을 보낼 때 결과를 받고 싶은 채널을 함께 실어 보내는 방법이 유용하다. 요청을 보내는 메시지에 채널도
	넣어서 보내면 된다.


*/
type Request struct {
	Num  int
	Resp chan Response
}

type Response struct {
	Num       int
	WorkkerID int
}

func PlusOneService(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}

/*

package main

import (
	"fmt"
	"sync"

	"github.com/chapter7"
)

func main() {

	reqs := make(chan chapter7.Request)
	defer close(reqs)

	for i := 0; i < 3; i++ {
		go chapter7.PlusOneService(reqs, i)
	}

	var wg sync.WaitGroup

	for i := 3; i < 53; i += 10 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			resps := make(chan chapter7.Response)

			reqs <- chapter7.Request{i, resps}
			fmt.Println(i, "=>", <-resps)

		}(i)
	}

	wg.Wait()
}


보내는 Request에 결과를 받을 채널도 함께 실어서 보냈다.

23 => {24 1}
3 => {4 0}
13 => {14 1}
33 => {34 0}
43 => {44 0}

(결과는 약간 다를 수도 있다.)

각각의 번호가 0,1,2로 붙어 있는 3개의 PlusOneService 고루틴으로 분산처리하고 있으며 총 5개의 요청을 서로 다른 고루틴에서
보냈다. 그리고 보낸 고루틴에서 자신이 보낸 요청에 대한 결과를 받을 수 있다.


fmt.Println(i,"=>",<-resps)


여기서는 요청 및 반응이 1대 1 대응이지만, 받는 부분을 위와 같이 처리하지 않고 다음과 같이 처리하면 요청 하나당 0개, 1개 혹은
여러 개의 응답을 받을 수 있다.


for resp:=range resps{
	fmt.Println(i,"=>",resp)
}

*/
