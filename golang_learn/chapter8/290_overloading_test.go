package chapter8

import "fmt"

/*
	8.1 오버로딩


	오버로딩은 같은 이름의 함수 및 메서드를 여러 개 둘 수 있는 기능이다. 인자의
	자료형에 따라 혹은 개수에 따라서 다른 함수 및 메서드가 호출되게 할 수 있다. 그러나
	Go언어에서는 지원되지 않는다.


	Go언어에서 오버로딩을 어떻게 흉내내는지가 중요한 것이 아니라 어떤 문제를 풀기 위해서 오버로딩이
	필요한지 생각하는 것이 중요하다. 몇 가지 유형으로 나누어서 살펴볼 수 있다.


	* 자료형에 따라 다른 이름 붙이기: 오버로딩을 반드시 하지 않아도 되는 경우가 많다. 이 경우에는 자료형에 따라서
	다른 함수의 이름을 붙이자.


	* 동일한 자료형의 자료 개수에 따른 오버로딩: 예를 들어 max(a,b)와 max(a,b,c)를 지원하는 함수를 만드는 경우다.
	이 경우에는 가변 인자를 사용하면 된다. 가변 인자는 4장을 본다.

	* 자료형 스위치 활용하기: 오버로딩을 반드시 해야 하는 경우는 인터페이스로 인자를 받고, 메서드 내에서 자료형 스위치로
	다른 자료형에 맞추어 다른 코드가 수행되게 할 수 있다. 자료형 스위치는 5장에 있다.


	* 다양한 인자 넘기기: 편의를 위하여 오버로딩을 하는 경우가 있다. 서로 다른 이름을 붙여도 상관이 없지만 기본값을 표함한 여러
	설정을 넘기는 경우는 이들을 모두 묶은 구조체를 넘기는 것을 고려하자. 오버로딩을 쓰는 경우보다 코드가 더 깔끔해지는 경우가 많다.




	오버로딩을 하기 보다 다른 이름을 붙이는 것이 더 나은 경우를 보겠다. 이것은 C++ 예시이다.

	int volume(int s){
		return s*s*s;
	}

	// volume of a cylinder
	double volume(double r, int h){
		return 3.14*r*r*static_case<double>();
	}



	위의 것은 오버로딩의 좋은 예로 올려 놓았지만, 같은 volume으로 같이 두어서 오히려 혼란스럽다.


	이번에는 편의를 위하여 다양한 인자가 오버로딩되는 경우를 보겠다. 아래는 C++ 혹은 자바 코드 예제이다.


	Element getElement(int idx){
		return getElement(idx,DEFAULT);
	}

	Element getElement(int idx,Language lang){
		return getElement(idx,lang,false)
	}

	Element getElement(int idx,Language lang, bool excludeEmpty){
			// ... implementation ...
	}



*/

/*	이런 경우(다양한 인자가 오버로딩 되는 경우)는 구조체를 넘기는 것이 좋다.
 */
type Language string

type Option struct {
	Idx          int
	Lang         Language
	excludeEmpty bool
}

func GetElement(opt Option) /**Element */ {
	// ...
}

/*
	인터페이스를 활용하는 것이 나을 수도 있다. 다음과 같이 이름이 같고 인자의 자료형이 다른 String() 함수를 둘 이상 정의할 수 없다.

	func String(int i){...}
	func String(double d) {...}

*/

/*
	이것은 함수 이름을 IntToString, DoubleToString으로 각기 따로 지어도 되지만 인터페이스를 활용하는 것이 더 나은 방법이다.
	아래 예제는 Int와 Double 자료형에 String() 메서드를 정의했다.


*/
type Stringer interface {
	String() string
}

type Int int
type Double float64

func (i Int) String() string {
	I := int(i)
	return fmt.Sprint(I)
}
func (d Double) String() string {
	D := float64(d)
	return fmt.Sprint(D)
}


/*
	이제 다음과 같이 String()을 지원하는 것들에 대하여 String() 호출을 할 수 있다.

*/
func Example_string() {
	fmt.Println(Int(5).String(), Double(3.7).String())
	// Output:
	// .

}


/*
	메서드가 아닌 일반 함수로도 쉽게 변환이 가능하다.
*/


func ToString(s Stringer)string  {
	return s.String()
}

/*
그러나 다시 한 번 생각하여, 어떤 문제를 해결하고자 하는지를 생각한다면 굳이 억지로 오버로딩을 할 경우는 줄어든다.
*/