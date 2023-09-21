package chapter5

import (
	"fmt"
	"sort"
	"strings"
)

/*
	5.3.3 정렬과 힙


	정렬은 자료들을 어떤 순서에 따라서 늘어놓는 것이다. 서로 같지 않은 두 자료를 비교했을때
	어떤 자료가 먼저 오고 어떤 자료가 나오는지를 판단할 수 있어야 하며, 순환구조가 있으면 안된다.

	자료 수를 셋으로 늘린다면 가위, 바위, 보 중에서 어떤 것이 가장 우세한지 알 수 없다는 것이다.
	정렬에서 가장 많이 이용되는 순서는 수의 대소에 의한 순서와 사전식 순서(lexicographic order)이다.

	많은 정렬 알고리즘들이 있고 각자 특징들이 있는 시간 복잡도만 해도 최선(best), 평균(average), 최악(worst)의
	경우에 대하여 각기 다르다. 자료의 특징에 따라서 비교 정렬보다 더 나은 정렬 방법을 이용할 수도 있다.

	비교 정렬(comparison sort)이란 두 자료를 비교하여 어느 것이 먼저 오는지 순서만 알아내서 정렬하는 방법이다.
	안정 정렬(stable sort)이란 서로 같은 키에 대하여 원래의 순서가 그대로 유지되는 정렬 방식이다.



	Go의 sort.Sort에서 이용하는 방법은 비교 정렬이자 불안정 정렬(unstable sort)이다. 두 자료를 비교하여 어느 자료가
	더 먼저 와야 하는지 결과를 돌려주는 부분만 작성하면 나머지 부분은 이미 나와 있는 정렬 알고리즘을 이용하여 정렬해 준다.
	그러나 임의의 자료형에 대하여 두 자료의 순서를 비교하는 함수는 제네릭(Generic)을 지원하지 않는 언어에서는 구현하기
	까다로울 것이다.


		정렬 인터페이스의 구현

	Go는 제네릭을 지원하지 않지만 인터페이스를 지원하기 때문에 다양한 형태의 정렬을 수행할 수 있다. sort 패키지를
	보면 sort.Interface라는 인터페이스를 정의하고 있고 이것에 따르기만 하면 정렬을 할 수 있다.

*/

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with
	// index i should sort before the element with index j.

	Less(i, j int) bool

	Swap(i, j int)
}

/*
	i번째 자료와 j번째 자료를 지칭할 때 해당 자료의 자료형은 정해져 있지 않지만, 인덱스 i와
	j는 항상 정수형으로 고정될 것이다. 이를 이용하여 이 인터페이스를 만들어낸 것이다.


	문자열을 대소문자 구분 없이 정렬해보자. 기본적으로 제공되는 문자열 정렬을 이용하면 AppStore,
	MacBook, iPad, iPhone 순서로 정렬될 것이다. 대소문자 구분 없이 정렬한다는 것은 AppStore, iPad,
	iPhone, MacBook 순서로 정렬하는 것이다. 일단 두 문자열을 모두 소문자로 변경을 한 뒤 비교하고,
	같은 경우에는 원래 문자열을 비교하면 된다. 그렇게 하면 IPAD가 iPad보다 항상 먼저 나오게 할 수
	있다.
*/

type CaseInsensitive []string

func (c CaseInsensitive) Len() int { return len(c) }

func (c CaseInsensitive) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func (c CaseInsensitive) Less(i, j int) bool {

	return strings.ToLower(c[i]) < strings.ToLower(c[j]) || (strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])

}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})

	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// .
}

/*
	이제는 이미 익숙하겠지만 구조체를 이용하지 않아도 인터페이스를 사용할 수 있다는 점이다.
	문자열 배열을 CaseInsensitive 자료형으로 이름 짓고 이것에 대하여 정렬 인터페이스를 구현하였다.
	물론 문자열 배열은 서로 다른 여러 가지 이름을 붙일 수 있다.

*/
