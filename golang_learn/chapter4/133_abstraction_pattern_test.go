package chapter4

/*
	4.2.8 패턴의 추상화

	변수라는 것이 있어서 어떤 값이나 연산의 결과에 대하여 이름을 붙이고 추상화할 수 있다.
	그리고 함수라는 것이 있어서 코드를 값의 입출력으로 추상화할 수 있다. 그리고 고계 함수를
	이용하염 좀 더 높은 수준의 추상활르 이룰 수 있다.

	chapter3의 한글 받침의 코드를 확인해 보자.

*/
var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {

		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

/*
	7번째 줄을 보면 r에서 start를 뺀 결과가 한글의 색인 번호, 즉 '가'는 0, '각'은 1, ... 이렇게 이어지는 번호라는 것을
	index라는 이름을 주어 의미를 부여했다. 그 아래에 index는 단 한번 쓰였지만 int(r-start)를 매번 나열하지 않고 index를
	여러 번 사용할 수 있다. 가장 단순하고 낮은 수준의 추상화이다.

	HasConsonantSuffix 함수는 문자열을 넘기면 bool형으로 된 결과가 돌아오는 한글 받침이 마지막에 어디에 있는지 알아보는 함수로
	추상화되어 있다. 아이디어는 코드로 표현되고 이 아이디어를 인자로 넘어오는 모든 문자열에 대하여 적용할 수 있다. 변수에
	단순한 값을 담는 것보다 더 높은 수준의 추상화이다.




	아까 전에 살펴보았던 생성기 예제를 확인해 본다.

	125_generator_test.go

	func NewIntGenerator() func() int {
	var next int

	return func() int {
		next++
		return next
	}
}
	위와 같은 생성기를 만들었는데 NewVertexIDGenerator, NewEdgeIDGenerator와 같이 동일한 패턴의
	생성기를 계속 만들었다. 만들면서 별 죄책감 없이 코드를 복사&붙여넣기 하고는 했는데, 반복되는
	패턴이 있다는 것은 추상화를 할 수 있다는 것이다.


	func NewVertexIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}
	이 코드가 위의 코드와 무엇이 다른가? 결과 반환하기 전에 VertexID로 형변환하는 것 밖에는 없다.
	그럼 재사용이 가능하지 않을까?


	func NewVertexIDGenerator()func ()VertexID{
		gen:=NewIntGenerator()
		return func()VertexID{
			return VertexID(gen())
		}
	}


	이와 같은 패턴의 추상화에 대한 예제는 'Structure and Interpreatation of Computer Programs' (SICP)의 1장 3절에
	나온다.




	위의 코드는 그 책의 1장 3절을 거의 옮긴 것이다.

package main

import (
	"fmt"
	"math"
)

type Func func(float64) float64
type Transform func(Func) Func

const tolerance = 0.00001
const dx = 0.00001

func Square(x float64) float64 {
	return x * x
}

func FixedPoint(f Func, firstGuess float64) float64 {
	closeEnough := func(v1, v2 float64) bool {
		return math.Abs(v1-v2) < tolerance
	}
	var try Func

	try = func(guess float64) float64 {

		next := f(guess)
		if closeEnough(guess, next) {
			return next
		} else {
			return try(next)
		}

	}
	return try(firstGuess)
}

func FixedPointOfTransform(g Func, transform Transform, guess float64) float64 {
	return FixedPoint(transform(g), guess)
}

func Deriv(g Func) Func {
	return func(x float64) float64 {
		return (g(x+dx) - g(x)) / dx
	}
}

func NewtonTransform(g Func) Func {
	return func(x float64) float64 {
		return x - (g(x) / Deriv(g)(x))
	}
}

func Sqrt(x float64) float64 {
	return FixedPointOfTransform(func(y float64) float64 {
		return Square(y) - x
	}, NewtonTransform, 1.0)
}

func main() {
	fmt.Println(Sqrt(2))
}


(기타 설명은 생략한다.)


*/
