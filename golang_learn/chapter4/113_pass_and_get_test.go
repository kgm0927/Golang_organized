package chapter4

import "fmt"

/*

	코드를 덩어리를 만든 다음에 그것을 호출하고 귀환할 수 있는 구조를 '서브루틴'이라고 한다.
	큰 프로그램을 서브루틴으로 구분하면 코드를 재사용하여 중복된 코드를 줄일 수 있고, 서브루틴의
	내부와 외부를 분리하여 생각할 수 있어서 코드를 추상화하고 단순화할 수 있다.


	Go에서 이런 서브루틴을 함수라고 한다. 다른 많은 언어에서도 함수라고 부르는 경우가 많으며,
	서브프로그램, 프로시저, 메서드, 호출 가능 객체 등 여러 가지 이름으로 불린다.


	함수는 한줄 한줄 다 알고 있지 않아도 어떤 함수가 무엇을 하는지 알고 있다면 그 함수에 입력과
	출력만 알고 쓸 수 있다.


	내부적으로 서부루틴은 주로 스택으로 구현한다. 일반적을 호출이 이루어지면 스택에 현재의 프로그램 카운터(PC)
	와 넘겨줄 인자들을 넣은 뒤에 프로그램 카운터의 값을 변경하며 호출될 서브루틴으로 건너뛴다. 이런 구조기에
	재귀호출을 이용한 프로그래밍도 가능하다.


	Go언어의 값에 의한 호출(call by value)만을 지원한다. 함수 내에 넘겨받은 변수값을 변경하더라도 함수 밖의
	변수에는 영향을 주지 않는다. 변수에 담겨 있던 값만 넘어왔기 때문이다. 따라서 함수 밖에 변수의 값을 변경하려면
	해당 값이 들어 있는 주소값을 넘겨받아서 그 주소에 있는 값을 변경하여 참조에 의한 호출(Call by reference)과
	비슷한 효과를 낼 수 있다.


		4.1 값 넘겨주고 넘겨받기

	서브루틴을 호출할 때, 그 서브루틴에 값을 넘겨받거나 받을 수 있다. 이것이 가능하기 때문에 서브루틴을 다른 값에
	대하여 재사용할 수 있고, 결과값을 받아서 이용할 수 있다.


	4.1.1 값 넘겨주기


	3.4.4 '텍스트 리스트 읽고 쓰기'를 해 봤을 때 ReadFrom_2함수에서 *[]string 자료형으로 lines를 받은 것을 기억할 것이다.

	func ReadFrom_2(r io.reader,lines *[]string)error{}

	[]string 자료형이 아닌 *[]string 자료형으로 받은 이유는 이 ReadFrom 함수가 lines변수의 값을 변경하고자 하기 때문이다.
	앞에서 말한 것 처럼 슬라이스 배열에 대한 포인터, 길이, 그리고 용량 이렇게 세 값으로 이루어진다. 만일 *[]string을 쓰지 않고
	[]string을 이용하여 넘겼다면 이 세 값이 넘어가며, 이 세값을 담고 있는 변수와는 연관성이 없다.

	즉, lines[]string과 같이 받았다면 함수 내의 lines 변수가 품고 있는 세 값을 변경한다 해도, 그것은 바깥 세상과는 무관한 일이
	된다.


	만약 lines[]string으로 받았지만 거기에 있는 배열 포일터를 따라가서 거기 있는 값을 변경시켰다면 어떤 일이 발생하는가?
	슬라이스가 담고 있어서 넘어간 세 값은 그저 값일 뿐이지만, 그 세 값 중 첫 번째 값인 배열 포인터를 타고 가면 함수 바깥에서
	쓰던 것과 같은 배열을 가리키고 있기 때문에 여기에 대하여 변경이 일어나면 영향을 받게 된다.
*/

func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

func ExampleAddOne() {
	n := []int{1, 2, 3, 4}
	AddOne(n)
	fmt.Println(n)
	// Output:
	// .
}

/*
 ExampleAddOne에서 1,2,3,4 이렇게 정수 넷을 담고 있는 슬라이스를 하나 만들었다.
 그리고 AddOne이라는 함수에 값을 넘겨준 뒤에 그 값을 찍어 보았다. 분명히 값을
 넘겨주었는데 왜 호출 뒤 변경이 일어날까? 그것은 슬라이스가 배열에 대한 포인터,
 길이, 용량 이렇게 세 값으로 이루어진 것으로 그 세 값이 넘어간 것이기 때문이다.

 이제 다시 ReadFrom_2 함수를 확인해본다.

 func ReadFrom(r io.reader,lines *[]string)error{}


 이 함수는 왜 *[]string과 같이 포인터를 넘겨주고 있을까? 이유는 이 함수가 넘겨준 슬라이스의
 값을 변경해야 하기 때문이다. 슬라이스 값이라면 배열 포인터, 길이, 용량이다. 이 함수는 슬라이스에
 새로운 값을 추가하기 때문에 필연적으로 길이를 변경해야 한다. 그리고 용량이 부족한 경우는 더 크기가
 큰 새로운 배열을 만들어야 하기 때문에 배열 포인터와 용량 또한 변경이 일어날 수 있다.


 포인터로 넘어온 값은 *을 앞에 붙여서 값을 참조할 수 있다. 변수 앞에 &를 붙이면 해당 변수에 담겨 있는 값의 포인터 값을
 얻을 수 있다.

 무엇이 값인지가 중요하며 포인터 자료형으로 받더라도 그것은 주소값이 넘어와서 받는 포인터 변수가 담게 되는 것으로 이해하면 된다.
*/
