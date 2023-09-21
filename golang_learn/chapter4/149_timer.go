package chapter4

import (
	"fmt"
	"time"
)

func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

/*
package main

import (
	"fmt"

	"github.com/golang_learn/chapter4"
)

func main() {
	fmt.Println("Ladies and getleman!")

	chapter4.CountDown(5)
}



*/

/*
	만약 time.Sleep이 없다면 순식간에 카운트 다운이 일어날 것이다.
	이 함수를 호출해보면 1초에 한 줄씩 출력이 되어 카운트 다운의 느낌이
	난다.


	이와 같은 타이머를 블로킹(blocking) 타이머라고 한다. 1초라는 시간을 기다리는 동안
	프로그램은 잠시 수행을 멈춘다.

	그러나 우리가 타이머를 이용할 때는 기다리는 경우에만 이용하는 것이 아니다. 찌개를 끓이는
	도중에 타이머를 맞춰 두고 다른 일을 한다. 이런 타이머를 넌 블로킹(non-blocking) 타이머라고
	한다. 그리고 블로킹을 동기(synchronous), 넌 블로킹을 비동기(asynchronous)라고 한다.

	time이라는 모듈 안에 있는 Timer를 사용하면 '넌 블로킹 타이머'를 이용할 수 있다. 넌 블로킹
	타이머가 이용되는 한 가지 예는 사용자 인터페이스에서 찾을 수 있다. 저장 버튼을 눌렀을 때
	화면 구석에서 저장되었다는 메시지가 나타났다가 몇 초 후에 사라지는 것을 많이 보았을 것이다.
	그 메시지가 나타나 있는 몇 초간 프로그램은 계속 수행이 되고 사용자의 반응에 응답해야 한다.


	이런 비동기적 상황에서 사용되는 테크닉이 '콜백'이다. 콜백은 어떤 조건이 만족될 때 호출해달라고
	요청하는 것으로, 이미 이번 장에서 어떻게 고계 함수를 이용하여 콜백을 이용하는지 알아보았다.

	time.AfterFunc(5*time.Second, func(){
		// 메세지를 없애는 코드
	})

	이렇게 되면 5초 뒤에 수행되는 코드를 수행하지 않고, 바로 그다음 코드들이 수행되다가 5초가 지난 뒤에
	메시지를 없애는 코드가 수행이 된다. 아까 전의 카운트 다운 예제에서 메인 함수의 맨 처음에 아래의 세 줄을
	추가해 준다.


	time.AfterFunc(5*time.Second, func(){
		fmt.Println("I am so excited")

	})


	카운트 다운이 되는 도중에 흥분되는 나머지 누군가가 보낸 메시지가 출력될 것이다. 타이머 코드가 수행되기 전에
	중단할 수도 있다. time.AfterFunc가 호출이 되면 반환값으로 타이머를 돌려준다.이 타이머를 스탑하면 된다.
*/
