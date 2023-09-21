package chapter5

import (
	"container/heap"
	"fmt"
)

/*

	5.3.5 빈 인터페이스와 형 단언

	빈 인터페이스는 매우 유용하게 사용될 수 있다. 나열된 메서드들을 정의하고 있는 자료형은
	인터페이스로 취급될 수 있다는 점을 생각해보면 빈 인터페이스는 아무 자료형이나 취급할 수
	있다는 뜻이 된다. map[string]interface{} 자료형은 키가 문자열이고 값은 아무 자료형을 이용할
	수 있다.const



	그러면 interface{} 타입을 원래의 자료형으로 변환하려면 어떻게 하면 될까? 지금까지 형변환은
	크게 세 가지가 있다.


	* 숫자형 사이에서 형변환, int, int32, int64 등 서로 다른 크기의 정수형끼리, float32, float64와 같은
	실수형끼리, 그리고 정수와 실수 사이의 변환

	* string과 []byte 사이의 변환

	* 사실상 같은 자료형으로 표현되지만 이름이 다른 경우, type ID int라고 한 경우와 int 사이의 변환



	모두 이 자료형을 쓰고 괄호를 열고 변환할 값을 쓰고 닫으면 된다. ID(30)과 같이 int를 ID로 변환했다.






	인터페이스는 위에서 본 형변환과는 다르기 때문에 다른 문법을 사용한다. 인터페이스는 실제로 자료형과
	값을 가지고 있는 구조체로 표현이 된다. 따라서 형변환을 할 때 자료형이 맞는지 실행 시간에 검사가 일어나야
	한다. 그래서 형 단언(type assertion)이라 표현한다.
*/

func ExampleCaseInsensitive_heapString() {
	apple := CaseInsensitive([]string{"iPhone", "iPad", "MacBook", "AppStore"})

	heap.Init(&apple)
	for apple.Len() > 0 {
		popped := heap.Pop(&apple)
		s := popped.(string)
		fmt.Println(s)
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}

/*
	조금 전의 예제와 거의 같다. 다른 점은 heap.Pop(&apple)을 한 뒤 그 값을 바로 출력하지 않고 string형으로
	받은 것이다. heap.Pop을 보면 interface{}형을 반환하고 있다. 따라서 popped는 interface{}형이다. 그런데
	이것에 .(string)을 붙여서 이것은 필히 문자열형이라고 '형 단언'을 한 것이다.

	그래서 for 반복문 안에 있는 s는 문자열형이다. 만일 실행 시간에 이것이 단언한 형이 아니라면 패닉이 발생한다.



	형 단언은 빈 인터페이스만 쓸 수 있는 것이 아니다. 인터페이스를 실제 자료형으로 받을 때 마찬가지로 형 단언을 사용한다.


	var r io.Reader=NewReader()
	f:=r.(os.File)



	r이 실제로 os.File인 경우에 f를 해당 자료형으로 이용할 수 있다.그러나 자료형이 맞지 않으면
	패닉이 발생하기 때문에 반드시 저 자료형이 아닌 경우이거나, 여러 가지 가능성이 있을 때 두 값으로 받으면
	두 번째 값으로 검사할 수 있다.

	f,ok:=r.(os.File)



	형 단언은 구체적인 자료형 뿐만 아니라 다른 인터페이스로도 가능하다.

	var r io.Reader=NewReader()
	f,ok:=r.(io.ReadCloser)

	자세한 예제는 곧 살펴본다.
*/
