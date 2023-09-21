package chapter4

/*
	명명된 자료형

	자료형을 설명하면서 의도적으로 설명하지 않은 부분이 있다. 자료형에 새로 이름을 붙일 수 있다.
	3장에서 문자열 설명할 때 rune 형이 사실은 int32의 별칭이라고 말했다. 이것은 다음과 같이 새로 이름
	붙이기가 가능하다.

	이러한 자료형을 명명된 자료형(Named Type)이라고 한다.

	type rune int32

	엄밀히 말해서 int32 역시 명명된 자료에 속한다. 그러면 명명되지 않은 자료형은 어떤 것이 있을까?
*/

type runes []rune
type MyFunc func() int

/*
	위와 같이 runes와 MyFunc는 이름 자체만으로 자료형을 지칭하는 것은 아니기 때문에 명명된 자료형이 되고,
	[]rune과 func()int는 이름만으로 자료형을 지칭하는 것이 아니기 때문에 명명되지 않은 자료형으로 구분된다.
	명명된 자료형과 명명되지 않은 자료형 모두에 type 예약어를 사용하여 새 이름을 붙여줄 수 있다.


	이렇게 자료형을 검사함으로서 프로그램을 직접 수행해보기 전에 컴파일 시점에서 버그를 어느 정도 예방할 수 있다.

	예를 들어 정점과 간선으로 이루어진 그래프를 다루는 코드를 작성했다고 했을 때, 각각의 정점과 간선은 정수형으로 된
	ID 값을 사용하여 접근하는 코드이다.
*/

func NewVertexIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

/*
	위의 함수에서 Vertex를 Edge로 바꾼 NewEdgeIDGenerator라고 하는 간선 ID 생성자도 만들었다.
	NewEdge라고 하는 함수는 간선의 ID를 받아서 새로운 간선을 만드는 함수라고 보자.
*/

func NewEdge(eid int) {

}

func NewEdgeIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

/*
func main(){
	gen:=NewVertexIDGenerator()
	gen2:=NewEdgeIDGenerator()

	...
	e:=NewEdge(gen())
}


위의 메인함수의 문제점은 간선의 ID를 새로 생성해서 넘겨야 하는데 새로 생성된 정점의 ID를 넘겨버렸다. 동적하지
않은 코드를 보면서 한참 찾을 수도 있다.

이 버그가 컴파일 시간에 잡히지 않은 이유는 NewEdge를 호출할 때, 간선의 ID를 넘겨야 하는데 정점의 ID를 넘긴 것에 대한
실수를 잡지 못했기 때문이다.

우리는 여기서 정점과 간선의 ID에 대한 자료형에 서로 다른 이름을 붙일 수 있다.

*/

type VertexID int
type EdgeID int

/*
	위와 같이 int에 VertexID와 EdgeID 두 이름을 붙였다. 둘 다 int형으로 표현된다. 그러나 이들 둘은 int형과
	다른 이름이 붙은 자료형이다.

*/

func NewVertexIDGenerator_2() func() VertexID {

	var next VertexID
	return func() VertexID {
		next++
		return next
	}

}

/*
	이제 위와 같이 NewVertexIDGenerator를 변경하였다. 이 함수가 반환하는 생성기는 호출될 때마다 VertexID를 반환한다.
	만약에 next 변수가 VertexID형이 아니라 int형이라면 바로 next를 반환할 수 없고 VertexID로 형변환해주어야 한다.
	이는 int, VertexID 둘 다 명명된 자료형이기 때문인데, 서로 다른 명명된 자료형끼리는 호환되지 않는다.


	이를 통해서 main 함수에 기입하면 코드에 컴파일 에러가 난다. 이는 미리 실수를 방지할 수 있다.


	그러나 위에서 본 []rune과 runes와 같이 명명되지 않은 자료형과 명명된 자료형 사이에는 표현이 같으면 호환이 된다.
	따라서 위와 같은 경우는 컴파일 오류가 나지 않는다.

	func main(){
		var a []rune=runes{65,66}
		fmt.Println(string(a))
	}

	명명된 자료형을 이용하면 자료형을 하드 코드하는 것에 비하여 나중에 일괄적으로 해당 자료형의 표현을 변경할 수 있다는 점이 장점이다.


	그러나 위에서 본 []rune과 runes와 같이 명명되지 않은 자료형과 명명된 자료형 사이에는 표현이 같으면 호환이 된다. 따라서 위의 경우에는
	컴파일 오류가 생기지 않는다.


	명명된 자료형을 이용하면 자료형을 하드 코드하는 것에 비하여 나중에 일괄적으로 해당 자료형의 표현을 변경할 수 있다는 점을 장점으로 보면
	된다.
*/
