package chapter4

/*
	함수의 인자를 고정하고 싶을 때가 있는데, 3장의 연습문제 5번에서 MultiSet를 구현했다. 여기서 Insert를 확인해본다.

*/

// func Insert(m map[string]int,val string)
/*
	일단 명명된 자료형을 이용하여 map[string]int를 모두 MultiSet 자료형으로 바꾼다.
*/

type MultiSet map[string]int

type SetOp func(m MultiSet, val string)

// Insert 함수는 집합에 val을 추가한다.

func Insert(m MultiSet, val string) {

}

/*
	앞에서 우리가 만든 고계 함수 ReadFrom에 이 함수를 사용할 수 있을까? 지정된 MultiSet에 각 줄을
	집어넣고 싶은 경우이다. 이 경우에는 다음과 같이 호출이 가능하다.

	m:=NewMultiSet()

	ReadFrom(r, func(line string){
		Insert(m,line)
	})

	이제 r에서 읽어서 매 줄을 MultiSet에 추가해주는 코드가 되었다. 그러나 이렇게 함수의 형태를 변환하는 것 또한
	추상화가 가능하다.

*/

func InsertFunc(m MultiSet) func(val string) {

	return func(val string) {
		Insert(m, val)
	}

}

/*
	그러면 이제 호출 부분을 다음과 같이 단순화할 수 있다.

	m:=NewMultiSet()
	ReadFrom(r,InsertFunc(m))


	물론 더 일반화시킬 수 있다.
*/

func BindMap(f SetOp, m MultiSet) func(val string) {

	return func(val string) {
		f(m, val)
	}

}

/*
	이렇게 되면 다음과 같이 호출할 수 있다.

	m:=NewMultiSet()
	ReadFrom(r,BindMap(Insert,m))

	마치 Insert 함수의 첫 인자인 m을 고정한 함수를 이용하는 것처럼 사용할 수 있다.
*/
