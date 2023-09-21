package chapter7

/*
	지금까지는 프로그램이 하나으ㅢ 순차적인 흐름으로 수행되었다. 이번에는 동시성에 대한 개념을
	구현할 예정이다.


	7.1 고루틴

	고루틴은 가변운 스레드와 같은 것으로 현재 수행 흐름과 별개의 흐름을 만들어준다. 고루틴을 생성하는
	방법은 간단하다. 함수를 다음과 같이 호출하는데

	f(x,y,z)

	이것을 go 키워드를 붙여준다.


	go f(x,y,z)

	앞에서 go를 붙여서 이와 같이 함수를 호출하게 되면 f(x,y,z) 호출과 현재 함수의 흐름은 메모리를 공유하는 논리적으로
	별개의 흐름이 된다. 여기서 논리적으로 별개의 흐름이라고 한 이유는 물리적으로 별개의 흐름이 되는 것과는 구분되기
	때문이다.

*/