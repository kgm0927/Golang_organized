package chapter7

/*
	7.3.5 select


	select를 이용하면 동시에 여러 채널과 통신할 수 있다. select의 형태는 switch문과 비슷하지만
	동시성 프로그래밍에 사용되며 다음과 같은 특징이 있다.

	* 모든 case가 계산된다. 거기에 함수 호출등이 있다면 select를 수행할 때 모두 호출된다. 아래 예제에서
	c3 채널이 준비가 되어 있지 않더라도 f()는 반드시 호출이 된다.


	* 각 case는 채널에 입출력하는 형태가 되며 막히지 않고 입출력이 가능한 case가 있으면 그중에 하나가
	선택되어 입출력이 수행되고 해당 case의 코드만 수행된다.

	* default가 있다면 모든 case에 입출력이 불가능할 때 코드가 수행된다. default가 없고 모든 case에 입출력이
	불가능하면 어느 하나라도 가능해질 때까지 기다린다.


	select {
	case n:=<-c1:
		fmt.Println(n,"is from c1")

	case n:=<-c2:
		fmt.Println(n,"is from c2")

	case c3<-f():
		fmt.Println("No channel is ready")
	}


	이제 각각의 패턴을 보면서 익히겠다.



	팬인하기

	select를 이용하면 고루틴을 여러 개 이용하지 않고도 팬인ㅇ르 할 수 있다. 위의 예제에서 받은 부분만 떼어놓으면 된다.



	select{
	case n:=<-c1:c<-n
	case n:=<-c2:c<-n
	case n:=<-c3:c<-n
	}


	위의 코드는 c1, c2, c3 중 어느 채널이라도 자료가 준비되어 있으면 그것을 c로 보내는 코드이다. 저 select문을 for 반복문으로
	둘러싸면 팬인을 할 수 있다.


	만일 c1, c2, c3 중 어떤 채널이 닫혀 있다면 어떻게 될까? 닫혀 있는 채널은 막히지 않고 기본값을 계속해서 받아갈 수 있기 때문에
	이 경우는 닫힌 채널의 case가 선택되어 단힌 채널로부터 기본값이 받아질 가능성이 있다. 이것까지 처리하여 FanIn을 구현해보자. 기존
	의 FanIn함수는 생성기 행텨로 채널을 바로 돌려받는 형태이기 때문에 여기서도 그 패턴을 이용해본다.

	단 이전의 FanIn 함수는 정해지지 않는 수의 여러 채널을 팬인할 수 있었지만, 이번에는 그렇게 할 수 없기 때문에 3개로 제한해서 사용해보겠다.
	이름도 FanIn3로 지었다.


*/

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3

	closeChan := func(c *<-chan int) bool { // 안에 아무것도 없을 경우.
		*c = nil
		openCnt--
		return openCnt == 0
	}

	go func() {
		defer close(out)
		for {
			select {

			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}

			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}

			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}

			}
		}
	}()

	return out
}

/*
	결과는 매번 바뀔 수 있다.


	이 코드의 특징은 닫힌 채널을 nil로 바뀌주었다는 데에 있다. nil 채널에는 보내기 및 받기가 모두 막히게(blocking) 된다. 그렇기 때문에
	채널이 닫혔다는 것이 발견되면 이것을 영원히 막히는 채널로 바꿔준 것이다. 물론 채널 자체를 바꾼 것이 아니라 채널 변수를 nil로 바꾼 것이기에
	혹시나 이 예제에는 보이지 않는 다른 고루틴이 닫힌 채널에서 자료를 받아가고 있었다고 해도 그 쪽에서는 아무런 영향을 주지 않는다.


	결국 닫힌 채널들은 모두 막아버리면서 열려 있는 채널 숫자를 줄여서 이것이 0이 될 때 함수에서 반환해버리고 그것과 동시에 out 채널이 닫히게 되어
	결국 들어오는 3개의 채널이 모두 닫히면 나가는 채널을 닫는 로직이 구성된다. 이런 로직이 없다면 3개의 채널 중에서 닫힌 채널이 생기면 그 채널로부터
	0 값이 계속 들어오게 된다. 이때 ok 값을 참조하여 무시하고 넘어가는 방법도 있지만 깔끔하지 못한 방식이다.


*/

/*

	채널을 기다리지 않고 받기

	지금까지 채널에서 자료를 받을 때, 아직 채널에 값이 준비되지 않았으면 준비될 때까지 기다리는 방식으로만 프로그램을 작성했다.
	채널값이 있으면 받고, 없으면 그냥 스킵하는 흐름을 구성하려면 select를 사용하면 된다.


	select{
	case n:=<-c:
		fmt.Println(n)

	default:
		fmt.Println("Data is not ready. Skipping ...")
	}

	이렇게 default를 넣으면 동작한다.




	시간 제한

	채널과 통신을 기다리되 일정 시간동안만 기다리겠다면 time.After 함수를 이용할 수 있다. 이 함수는 채널을 돌려주는데 지정된 시간이 지나면
	이 채널로 현재 시간이 전달된다. case 하나로 이 채널에서 받아오면 시간이 넘었다는 것으로 return으로 이 함수 전체를 빠져나가게 구성하면
	된다.


	select{
	case n:=<-recv:
		fmt.Println(n)

	case send<-1:
		fmt.Println("sent 1")

	case <-time.After(5*time.Second):
		fmt.Println("No send and receive communication for 5 seconds")
		return
	}


	이 select가 for 반복문으로 둘러싸여 있을 때 타이머는 매번 반복될 때마다 새로 생성되므로, 한번의 채널 통신마다 5초의 제한 시간이 생긴다.(원래 select는 default
	가 없다면, 될 때까지 된다.)
	만약에 전체 제한시간을 걸고 싶을 경우 타이머 채널을 보관해두고 쓰면 된다.


	timeout:=time.After(5*time.Second)
	for{
		select{
		case n:=<-recv:
			fmt.Println(n)

		case send<-1:
			fmt.Println("sent 1")

		case <-timeout:
			fmt.Println("Communication wasn`t finished in 5 sec")
			return
		}
	}


	이런 코드가 작성이 되면 recv와 send에 빈번하게 자료가 반복적으로 오고 가더라도 5초 동안만 처리하게 된다.
*/
