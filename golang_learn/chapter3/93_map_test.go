package chapter3

/*
	3.3 맵

	Go 언어에서 map은 해시테이블로 구현된다. 해시맵은 키와 값으로 되어 있는데
	키를 이용해서 값을 상수 시간에 가져올 수 있다. 그 대신에 해시맵에서는
	순서가 없다.

*/

type KeyType int
type valueType int

var m map[KeyType]valueType

/*
	그러나 맵을 담을 수는 있지만 맵이 생성되어 있지는 않기에 이 상태로는 맵을 이용할 수 없다.
	정확히 말하면 빈 맵으로 취급하여 맵을 읽을 수는 있지만 변경할 수 는 없다.

	nil 값을 가지고 있는 슬라이스에 append로 덧붙일 수 있는 것과 달리 맵은 일단 생성이 되어 있어야
	추가할 수 있다.

	m:=make(map[keyType]ValueType)


	아니면 아래와 같은 방법으로 빈 맵을 초기화할 수 있다.

	m:=map[keyType]valueType{}



	맵에선 읽을 때에는 두 가지 방법이 있다. m[key]를 이용하면 맵의 값을 읽을 수 있다. 만약 해당 키가
	없다면 값의 자료형을 기본 값으로 반환한다. (값의 자료형이 정수이면 0, 문자열이면 빈 문자열이 돌아온다.)

	맵을 읽을 때 두 개의 변수로 받게 되면, 두 번째 변수에 키가 존재하는지 여부를 bool형으로 받을 수 있다.

	value,ok:=m[key]


	위와 같은 형태의 코드라면 key가 m에 들어 있으면 ok가 true 값이 된다.



	맵에 값을 쓸 때는 다음과 같이 자연스럽게 쓸 수 있다. 맵에 이미 키가 들어 있는 경우에는 기존에 있던 값이
	변경되고 없는 경우에 새로 생긴다.
*/
