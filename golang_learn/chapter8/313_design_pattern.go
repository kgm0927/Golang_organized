package chapter8

import "fmt"

/*
		8.4 디자인 패턴


		몇가지 유명한 디자인 패턴들이 풀고자 하는 문제를 Go에서 어떻게 풀 수 있을지 알아본다.

	8.4.1 반복자 패턴

	이미 여러가지 반복자 패턴을 알아보았다. 4 장에서는 클로저를 이용하여 호출하는 반복자, 콜백을
	넘겨주어서 이 함수가 모든 원소에 대하여 호출되게 하는 반복자, 5장에서 본 인터페이스를 이용한 반복자,
	7장에서 본 채널을 이용한 반복자 모두가 해당된다.


	채널을 이용한 반복자를 이용할 때, 중간에 중단할 수 있어야 하는 경우에는 반드시 done 채널 혹은
	context.Context를 받아서 처리하게 작성해야 한다.


	8.4.2 추상 팩토리 패턴

	추상 팩토리 패턴(abstract factory pattern)은 팩토리들을 여럿 묶어놓은 팩토리를 추상화하는 패턴이다.
	아래의 예제는 윈도우와 맥의 UI 구현을 다르게 하기 위하여 추상 팩토리를 이용하는 예를 보여준다.


	맥의 UI 팩토리를 생성하면 맥의 버튼과 레이블 팩토리가 묶여서 함께 생성되게 하는 것이 추상 팩토리이다.


*/

type Button interface {
	Paint()
	OnClick()
}

type Label interface {
	Paint()
}

// WinButton is a Button implemention for Windows
type WinButton struct{}

func (WinButton) Paint()   { fmt.Println("win button paint") }
func (WinButton) OnClick() { fmt.Println("win button click") }

// WinLabel is a Label implemention for Windows.
type WinLabel struct{}

func (WinLabel) Paint() { fmt.Println("win label paint") }

// WinButton is a Button implementation for Mac.

type MacButton struct{}

func (MacButton) Paint()   { fmt.Println("mac button paint") }
func (MacButton) OnClick() { fmt.Println("mac button click") }

// WinLabel is a Label implementation for Mac.

type MacLabel struct{}

func (MacLabel) Paint() { fmt.Println("mac label paint") }

// UI factory can create buttons and labels
type UIFactory interface {
	CreateButton() Button
	CreateLabel() Label
}

// WinFactory is a UI Factory that can create Windows UI elements.

type WinFactory struct{}

func (WinFactory) CreateButton() Button {
	return WinButton{}
}

func (WinFactory) CreateLabel() Label {
	return WinLabel{}
}

// MacFactory is a UI factory that can create Mac UI elements.

type MacFactory struct{}

func (MacFactory) CreateButton() Button {
	return MacButton{}
}

func (MacFactory) CreateLabel() Label {
	return MacLabel{}
}

// CreateFactory returns a UIFactory of the given os.

func CreateFactory(os string) UIFactory {
	if os == "win" {
		return WinFactory{}
	} else {
		return MacFactory{}
	}
}

func Run(f UIFactory) {
	button := f.CreateButton()
	button.Paint()
	button.OnClick()
	label := f.CreateLabel()
	label.Paint()
}
