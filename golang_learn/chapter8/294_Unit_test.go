package chapter8

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

/*
		8.2.1 유닛 테스트


	JUnit을 비롯한 xUnit 스타일의 테스트에서는 assertEqual과 같은 함수를 이용하여 두 값이 서로 같은지 비교한다.
	이것을 대체하기 위한 방법은 어떤 것들이 있을까?


	가장 잘 떠오르는 것은 if를 이용한 것이다.


	if expected !=actual{
		 t. Error("Not qual")
	}

	이 코드의 문제는 assertEqual로는 한 줄로 간결하게 표현할 수 있는 것이 3줄로 늘어났다는 것이다.

	그래서 assertEqual을 직접 작성하려고 하는데, assertStringEqual와 같이 자료형을 한정시켜
	만드는 것이다.
*/

func assertEqualStirng(t *testing.T, excepted, actual string) {
	if excepted != actual {
		t.Errorf("%s!=%s", excepted, actual)
	}
}

/*
	다른 하나는 reflect.DeepEqual을 이용하여 범용적인 assertEqual을 작성하는 방법이다.
	reflect 패키지는 실행 시간에 얻을 수 있는 자료형과 값에 대한 정보를 얻을 수 있는 패키지이다.
	Go언어의 어느 형태든 이용할 수 있으므로 assertEqual 대신에 reflect.DeepEqual을 활용할 수 있다.

*/

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v!=%v", expected, actual)
	}

}

/*
이렇게 assertEqual을 만드는 방법을 탈피해서 테이블 기반 테스트를 만들 수 있다. 5장의 테이블 기반 테스트 부분은 참고하면 좋겠다.(Json와 관련된 것이므로 넘긴다.)
*/

/*
다른 방법은 Example 테스트를 이용하는 방법이다. 이미 여러 번 쓰인 것이다.

*/

func Example() {

}

// 설명은 생략한다.

/*
	이번에는 어떤 구조체가 어떤 필드를 가지고 있는지 확인하는 방법이다. 확인하기보다는 필드 이름을 모두 출력하는 예제를 보여드리는 것이 더 나을 것
	같다.

*/

func FieldNames(s interface{}) ([]string, error) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("FieldNames: s is not a struct")
	}

	names := []string{}
	n := t.NumField()

	for i := 0; i < n; i++ {
		names = append(names, t.Field(i).Name)
	}
	return names, nil
}

/*
	이제 이렇게 필드 이름들을 모두 둘러볼 수 있다. reflect.Value를 이용하면 특정 필드의
	값을 얻어낼 수 있다.



*/

func Example_print() {
	s := struct {
		id   int
		Name string
		Age  int
	}{}
	fmt.Println(FieldNames(s))
	// Output:
	// .
}

/*
reflect 패키지에 함수나 메서드도 다룰 수 있으므로 함수를 받아서 다른 함수형으로 변경하여 반환하는 것도
가능하다. 예를 들어 무엇도 반환하지 않는 함수지만 에러를 반드시 돌려줘야 하는 경우라면 nil 에러를 돌려주게
할 수 있는데, 서로 다른 자료형에 대하여 이 함수를 만들 수 있다.
*/
func AppendNilError(f interface{}, err error) (interface{}, error) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return nil, errors.New("AppendNilError: f is not a function")
	}

	in, out := []reflect.Type{}, []reflect.Type{}

	for i := 0; i < t.NumIn(); i++ {
		in = append(in, t.In(i))
	}
	for i := 0; i < t.NumOut(); i++ {
		out = append(out, t.Out(i))
	}

	out = append(out, reflect.TypeOf((*error)(nil)).Elem())
	funcType := reflect.FuncOf(in, out, t.IsVariadic())
	v := reflect.ValueOf(f)
	funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {

		results = v.Call(args)
		results = append(results, reflect.ValueOf(&err).Elem())
		return results

	})

	return funcValue.Interface(), nil
}

/*
 이 함수는 어떤 함수를 받아서 그 함수의 반환값으로 err 값을 맨 뒤에 추가한 함수를 반환하는 함수이다.
 다음과 같이 사용이 가능하다.

 f:=func(){
	fmt.Println("called")
 }

 f2,err:=AppendNilError(f,errors.New("test error"))
 fmt.Println("AppendNilError.err:",err)
 fmt.Println(f2.(func()error)())


 이 라이브러리를 이용하면 동적 자료형 언어에서 할 수 있는 많은 것을 할 수 있다. 그러나 reflect를 사용하면
 정적인 자료형 검사를 할 수 없으므로 꼭 필요한 경우에만 이용하는 것을 권장한다.
*/
