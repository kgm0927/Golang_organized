package chapter8

import (
	"fmt"
)

/*
	8.3.3 상속


	객체지향에서 상속은 어떤 클래스의 구현들을 재사용하기 위하여 사용된다. IsA 관계인 경우와 HasA 관계가 성립한다.

	일단 HasA 관계의 경우에는 전통적인 객체지향에서도 상속보다는 객체 구성 (object composition)이 더 맞다. 이 경우에는
	재사용하고자 하는 구형의 자료형의 변수를 struct에 내장하면 된다. 그리고 그 필드를 사용하면 훌륭히 재사용하는 것이다.
	HasA 관계에서는 말 그대로 그 구현을 필드로 가지고 있으면 된다.

	이제 IsA 관계에 대해서 살펴본다. IsA 관계의 상속에서는 많은 경우에 추상 클래스를 상속한다. 추상 클래스 구현이 없는
	경우 Go언어의 인터페이스를 이용하면 된다. 이 자료형이 어떤 추상 클래스로부터 상속받고 있다는 것을 명시적으로 써줄 필요가
	없다는 점이 다르다.

	왜 객체지향 프로그래밍에서는 추상 클래스를 이용할까? 구현에 구애받지 않고 이용하고 필요한 경우에 서로 다른 구현으로 다형성을
	활용하기 위한 것이다. 이것은 인터페이스를 이용한 다형성 예제에서 보았듯이 Go 언어로 충분히 가능하다. 조금 익숙해지면 오히려
	명시적으로 상속을 써넣지 않아도 되기 때문에 더 유연하고 편리하다.


	그러면 이제 추상 클래스가 아닌 클래스를 상속받는 경우를 생각해보자. 이 경우는 인터페이스와 구현을 함께 상속한다. 인터페이스를
	상속함으로써 상위 클래스의 일종이라는 것을 이용한 다형성 코드가 가능해지고, 구현을 상속함으로써 반복을 피하는 것이다.
	이것은 인터페이스와 구조체 내장을 동시에 사용하면 가능하다.


	상속은 어떤 문제를 푸는 것일까? 밑에서 알아본다.





			메서드 추가



	기존에 있는 코드를 재사용하면서 기능 추가를 하고 싶은 경우는 상속할 수 있다. 메서드 추가를 통하여 기능을 추가한다.
	그러나 이런 용도로 상속이라는 복잡한 개념을 이용할 필요가 있을지 모른다. Go는 상속을 지원하지 않지만 여기서 풀고자
	하는 문제는 풀 수 있을 것이다.


	type Rectangle struct{
	width, Height float32
}

	func (t Triangle) Area()float32  {
	return 0.5*t.Width*t.Height
}

	여기 이렇게 Area()가 구현되어 있는데, 여기에 메서드를 추가하고 싶다. 내가 만든 패키지를 수정하는 경우라면 그저 메서드를
	추가하면 해결된다. 그러나 원래 구현은 다른 패키지에 있고 직접적으로 수정하기 어렵다는 가정이 있으면 상속이 유용해지게 된다.


	이를 위하여 구조체 내장(5장 참고)을 한다. 구조체 안에 다른 구조체를 넣어서 필드 참조나 메서드 호출을 위해 불필요한 코드를 작성하는 것을
	피할 수 있다.

*/

type RectangleCircum struct{ Rectangle }

func (r RectangleCircum) Circum() float32 {
	return 2 * (r.width + r.Height)
}

/* 이제 다음과 같이 초기화하고 기존에 있던 기능과 추가된 기능 모두 사용할 수 있다.*/

func ExampleRectangleCircum() {
	r := RectangleCircum{Rectangle{3, 4}}
	fmt.Println(r.Area())
	fmt.Println(r.Circum())
	// Output:
	// 12
	// 14

}

/* 필요하다면 상속과 함께 생성자도 만들어줄 수 있다.*/

func NewRectangleCircum(width, height float32) *RectangleCircum {
	return &RectangleCircum{Rectangle{width, height}}
}

/*
	여러 단계로 내장이 이루어지는 경우를 생각해본다. 다시 RectangleCircum을 내장하여 이 구현에 다른 구현을
	더 추가하고 싶으면 이 생성자 함수를 호출하여 내장되는 부분의 생성 및 초기화를 수행하면 된다.

*/

/*
		오버라이딩

	기존에 있던 구현을 다른 구현으로 대체하고자 하는 경우에도 상속을 쓸 수 있다. 이 문제 역시
	구조체 내장으로 해결이 가능하다.

	넓이 구하는 방법을 일부러 틀리게 대체하고 싶다. 그러면서도 원래 구현에 있는 코드 역시 재사용하고
	싶다. 기존 넢이의 2배를 출력하게 바꾸어 보겠다.

*/

type WrongRectangle struct{ Rectangle }

func (r WrongRectangle) Area() float32 {
	return r.Rectangle.Area() * 2
}

func ExampleWrongRectangle() {
	r := WrongRectangle{Rectangle{3, 4}}
	fmt.Println(r.Area())
	// Output: 24
}

/*
	바로 앞에서 다룬 메서드 추가 예제와 거의 비슷하다. 같은 이름의 메서드를 정의하면 된다.
	이렇게 정의하였다고 하더라도 여전히 WrongRectangle이 내장하고 있는 Rectangle에 대한 메서드는
	그대로 남아 있다. 같은 이름의 메서드를 정의해주었기 때문에 WrongRectangle에 대하여 Area 메서드를
	호출했을 때, Rectangle의 Area로 넘어가지 않은 것이다.


*/

/*
		서브 타입


	기존 객체가 쓰이는 곳에서 상속받는 객체를 쓰고자 상속하기도 한다. 인터페이스와 구조체 내장을
	모두 사용하면 된다. 이전의 WrongRectangle 예제를 보겠다.

	이 새로 만든 WrongRectangle도 Shape로 취급이 될까? 된다. 이것은 Area 메서드를 구현하고 있다.
	그러면 ReactangleCircum은 어떤가?



*/

func Example_totalArea() {
	fmt.Println(TotalArea(
		[]Shape{
			Square{3},
			Rectangle{6, 7},
			RectangleCircum{Rectangle{8, 9}},
			WrongRectangle{Rectangle{1, 2}},
		}))

	// Output: 127
}

/*

	실행 결과를 보다시피 둘 다 Shape으로 취급된다. RectangleCircum은 직접 Area 메서드를
	구현하고 있진 않지만 Area()를 호출하면 내장하고 있는 구조체 Area()가 호출될 것이기 때문에
	인터페이스를 구현하고 있는 것이 된다. 따라서 구조체를 내장하면 서브 타입으로 풀고자 하는
	문제도 풀 수 있음을 알 수 있다.


	어떤 자료형이 주어진 인터페이스를 구현하고 있는지를 알아보려면 reflect.Type.Implements메서드를
	쓰면 된다. reflect.TypeOf를 이용해서 자료형을 알아낸 다음에 Implements 메서드를 호출하면 된다.



	내장된 구조체가 있는지는 구조체에서 내장된 구조체의 이름으로 필드를 찾은 다음에 Anonymous 필드를
	찾아보면 된다.


	impl:=reflect.TypeOf(RectangleCircum{}).Implements(
		reflect.TypeOf((*Shape)(nil)).Elem(),
	)
	field,ok:=reflect.TypeOf(RectangleCircum{}).FieldByName("Rectangle")
	emb:=ok && field.Anonymous && field.Type==reflect.TypeOf(Rectangle{})

	위에서 impl에는 RectangleCircum이 Shape 인터페이스를 구현하는지 여부가 기록되며,
	emb는 Rectangle이 RectangleCircum에 내장되어 있는지 여부가 기록된다. 인터페이스 타입을
	얻어낼 때는 위에서 보다시피 조금 조심해야 한다.


	Shape(nil)은 nil인터페이스이기 때문에 인터페이스의 포인터로 nil을 만들고 자료형을 얻어내면
	인터페이스 자료형이 되고 여기에 Elem()을 호출하면 인터페이스 자료형을 얻어낼 수 있다.


	RectangleCircum{}과 같이 빈 객체를 만드는 것이 싫으면 비슷한 방법을 이용할 수 있다. 이렇게
	첫 줄을 변경하면 다음과 같이 된다.


	impl:= reflect.TypeOf((*RectangleCircum)(nil)).Elem().Implements(reflect.TypeOf((*Shape)(nil)).Elem())

	이렇게 reflect로 서브 타입 검사를 할 수 있는 것은 좋지만, 자주 쓰게 되지는 않는다. 중요한 것은 다른 언어의
	객체지향을 흉내내는 것이 풀고자 하는 문제를 가장 풀 수 있는 방법을 찾는 것이기 때문이다.
*/
