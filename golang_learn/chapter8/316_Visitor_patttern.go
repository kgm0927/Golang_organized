package chapter8

import "fmt"

/*
	비지터 패턴(Visitor Pattern)*은 알고리즘을 객체 구조에서 분리시키기 위한 디자인 패턴이다.
	인터페이스를 이용하여 구현하면 되므로 별로 어려움이 없이 구현할 수 있다.
	아래는 비지터 패턴의 예이다.

*/

type CarElementVisitor interface {
	VisitWheel(wheel Wheel)
	VisitEngine(engine Engine)
	VisitBody(bod Body)
	VisitCar(car Car)
}

type Acceptor interface {
	Accept(visit CarElementVisitor)
}

type Wheel string

func (w Wheel) Name() string {
	return string(w)
}

func (w Wheel) Accept(visitor CarElementVisitor) {
	visitor.VisitWheel(w)
}

type Engine struct{}

func (e Engine) Accept(visitor CarElementVisitor) {
	visitor.VisitEngine(e)
}

type Body struct{}

func (b Body) Accept(visitor CarElementVisitor) {

	visitor.VisitBody(b)
}

type Car []Acceptor

func (c Car) Accept(visitor CarElementVisitor) {
	for _, e := range c {
		e.Accept(visitor)
	}
	visitor.VisitCar(c)
}

type CarElementPrintVisitor struct{}

func (CarElementPrintVisitor) VisitWheel(wheel Wheel) {
	fmt.Println("Visiting" + wheel.Name() + " wheel.")
}

func (CarElementPrintVisitor) VisitEngine(engine Engine) {
	fmt.Println("Visiting engine")
}

func (CarElementPrintVisitor) VisitBody(body Body) {
	fmt.Println("Visiting body")
}

func (CarElementPrintVisitor) VisitCar(car Car) {
	fmt.Println("Visiting car")
}

type CarElementDoVisitor struct{}

func (CarElementDoVisitor) VisitWheel(wheel Wheel) {
	fmt.Println("Kicking my" + wheel.Name() + " wheel.")
}

func (CarElementDoVisitor) VisitEngine(engine Engine) {
	fmt.Println("Starting my engine")
}

func (CarElementDoVisitor) VisitBody(body Body) {
	fmt.Println("Moving my body")
}

func (CarElementDoVisitor) VisitCar(car Car) {
	fmt.Println("Starting my car")
}

/*
package main

import "github.com/golang_learn/chapter8"

func main() {
	car := chapter8.Car{
		chapter8.Wheel("front left"),
		chapter8.Wheel("front right"),
		chapter8.Wheel("back left"),
		chapter8.Wheel("back right"),
		chapter8.Body{},
		chapter8.Engine{},
	}

	car.Accept(chapter8.CarElementPrintVisitor{})
	car.Accept(chapter8.CarElementDoVisitor{})
}


CarElementPrintVisitor와 CarElementDoVisitor처럼 메서드를 구현하기만 하면
기능을 추가할 수 있다.

*/
