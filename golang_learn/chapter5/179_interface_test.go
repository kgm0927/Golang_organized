package chapter5

import "io"

/*

	5.3 인터페이스

	구조체가 자료들의 묶음이라면 인터페이스는 메서드들의 묶음이다. 의미상으로 인터페이스는
	무언가를 할 수 있는 것을 의미한다. io.reader는 read 메서드를 정의하는 자료형들을 받을 수 있다.
	예를 들어서 File은 Read 메서드가 있으므로 io.reader가 될 수 있다. 의미상으로는 읽을 수 있는
	것을 의미한다.

	인터페이스 이름을 붙일 때는 주로 인터페이스의 메서드 이름에 er를 붙인다. io.Reader Read 메서드를
	갖고 있는 인터페이스이다. 문자열로 표현할 수 있는 것들은 fmt.Stringer가 된다.

	내가 이름 붙인 자료형에 String 메서드만 만들어주면 문자열로 표현할 수 있는 것이 되고, fmt.Print와 같은
	함수에서 이 메서드를 사용하게 된다. 따라서 이것을 이용하면 손수 출력 형식을 정해줄 수 있다.

	Json 직렬화 예제에서 살펴본 바와 같이 MarshalJSON을 구현하기만 해도 json.Marshaler가 될 수 있기 때문에
	Json으로 직렬화할 수 있는 것이 되면서, 우리가 구현해준 방법으로 직렬화를 하게 된다. 인터페이스를 이해하면
	굉장히 많은 것이 가능하다.


	5.3.1 인터페이스의 정의


	인터페이스는 구조체와 매우 유사한 구조를 띄고 있다.

	interface {
	Methed1()
	Methed2(i int)error
	}

	이 인터페이스는 Method1과 Method2 모양을 갖고 있다. 두 메서드를 정의하고 있는 자료형은 이 인터페이스로
	사용할 수 있다.

	인터페이스 역시 이름을 붙여줄 수 있다.
*/

type Loader interface {
	Load(filename string) error
}

/* 구조체의 내장과 비슷한 형식으로 여러 인터페이스를 합칠 수 있다.*/

type ReadWriter struct {
	io.Reader
	io.Writer
}

type ReadWriter_i interface {
	io.Reader
	io.Writer
}

/*
	이렇게 되면 io.reader의 모든 메서드와 io.Writer의 모든 메서드를 구현하는 이름 붙인
	자료형은 ReadWriter가 된다.
*/
