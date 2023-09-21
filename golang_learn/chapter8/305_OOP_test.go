package chapter8

import "fmt"

/*
	8.3 객체지향

	Go는 객체지향을 완전히 지향하지 않는다. 여기서 다루는 부분은 객체지향의 범위 밖에 있는 것들도 있다.
	그러나 객체지향과 같이 다루어지는 경우가 많으므로 여기서 다루어본다.


	8.3.1 다형성


	다형성은 객체지향의 꽃이다. 객체에 메서드를 호출했을 때, 그 객체가 메서드에 대한 다양한 구현을 갖고
	있을 수 있다.


	다형성은 메서드가 호출되었을 때, 어떤 자료형이냐에 따라 다른 구현을 할 수 있게 하는 것이다.

	이것은 Go의 인터페이스로 쉽게 구현이 가능하다.

*/

type Shape interface {
	Area() float32
}

type Square struct {
	Size float32
}

func (s Square) Area() float32 {
	return s.Size * s.Size
}

type Rectangle struct {
	width, Height float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.Height
}

type Triangle struct {
	Width, Height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Width * t.Height
}

func TotalArea(shapes []Shape) float32 {
	var total float32

	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func ExampleTotalArea() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
	}))
}

/*
	Shape는 Area라는 함수가 있는 스페이스이다. 이것은 추상 클래스의 극단적 형태로 아무것도 들어있지 않다.

	TotalArea 함수에 넘기는 슬라이스는 Area 메서드만 구현하고 있으면 어떤 도형들도 담아서 넘겨 줄 수 있다.
	다형성을 주로 하는 커맨드 패턴과 같은 것들도 Go에서 이렇게 쉽게 구현이 가능하다.


	8.3.2 인터페이스

	자바 등에서 쓰는 인터페이스는 Go의 인터페이스 구현이 가능하다. 그러나 중요한 점은 Go언어에는 인터페이스 내의
	메서드를 구현하기만 하면 그 인터페이스를 구현하는 것이 된다. 이는 아주 중요한 특성이다.



*/
