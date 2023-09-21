package chapter5

import (
	"fmt"
	"time"
)

/*
	5.1.4 구조체 내장


	Go언어의 구조체는 유행하는 다른 프로그래밍 언어들의 구조체나 클래스에 비하여 별로 볼 것이 없다.
	많은 다른 언어에서 클래스나 구조체는 특별한 의미가 있다. 오직 클래스와 구조체만이 메서드를 가질 수
	있다거나 인터페이스를 구현할 수 있다고나, 직렬화 역직렬화가 가능하다거나 하는 등이 있다.

	메서드는 4장에서 봤듯이 구조체가 아니더라도 이름을 붙인 자료형이면 붙을 수 있다. 다음에 볼 예정이지만
	, 이것 역시 구조체일 필요는 없다. 그럼 무엇이 구조체를 특별하게 할까?


	구조체는 여러 자료형의 필드를 가질 수 있다는 점이 중요하다. 그러면 구조체를 재사용하는 방법은 어떤 것이 있을까?
	객체지향의 개념을 배운 사람이라면 상속과 포함 개념에 익숙할 것이다. 쉽게 말해 상속은 Is a관계, 포함은 Has a 관계
	라고 한다.


	Dealine이라는 자료형을 하나 만들어본다. 다른 프로그래밍 언어를 배우신 분들은 마감 시간을 가지는 클래스나
	구조체를 만들 것이다. 그러나 마감 시간 필드 하나만 있는 구조체를 굳이 만들 필요는 없다. Go 언어에서는 다음과
	같이 명명된 자료형, 혹은 자료형에 이름을 붙이는 일을 하면 된다. 꼭 구조체가 아니더라도 메서드를 작성할 수 있기에
	Deadline이라는 자료형에 OverDue 라는 메서드를 정의했다.

*/

type Deadline time.Time

func (d Deadline) OverDue() bool {
	return time.Time(d).Before(time.Now())
}

func Example_deadline_OverDue() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())
	// Output:
	// true
	// false
}

/*
현재 시간에서 4시간 전이 마감인 것을 d1에, 4시간 이후가 마감인 것을 d2에 담았다. 4시간 전에
마감인 것은 마감이 지난 것이 맞으므로 d1.OverDue()는 true를 돌려주고, 4시간 이후가 마감인
이후가 마감인 것은 아직 마감이 4시간이 남았으므로 d2.OverDue()는 false를 돌려준다.

데드라인이 없는 경우도 구현할 수 있을까? OverDue_2라는 함수를 따로 만들어 리시버를 포인터로 바꾸면 다음과
같은 코드가 가능하다.
*/
func (d *Deadline) OverDue_2() bool {
	return d != nil && time.Time(*d).Before(time.Now())
}

/*
 이제 마감 시간을 Task 구조체에서 사용해보자.*/

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

/*
	그리고 Task에 OverDue라는 메서드를 만들었다. 이 메서드는 그저 Deadline에 자기가
	할 일을 위임할 뿐이다.

*/

func (t Task) OverDue() bool {

	return t.Deadline.OverDue_2()
}

/*
이렇게 하면 이전에 나온 status와 관련된 코드들도 모두 추가하면 다음과 같은 테스트를 해볼 수 있다.
*/

func Example_taskTestAll() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, &d1}
	t2 := Task{"4h later", TODO, &d2}
	t3 := Task{"no due", TODO, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
	// Output:
	// true
	// false
	// false
}

/*
	Task 구조체가 Deadline 자료형의 필드를 가지고 있어서 메서드도 이용을 할 수 있다.
	그러나 메서드마다 모두 같은 이름의 메서드를 호출하는 코드를 작성해야 한다.
	이런 귀찮은 일을 덜어주는 것이 '내장 기능'이다.

	type Task struct {
	Title    string
	Status   status
	 *Deadline
}


	이전의 Task와 같지만 Deadline에 필드 이름을 생략했다. 이렇게 하면 Task에 대하여 의미없는
	OverDue 메서드를 작성한 필요가 없다. Task가 내장하고 있는 *Deadline자료형은 자료형의 이름과
	같은 Deadline이라는 필드를 가지게 되고 정의되어 있는 메서드도 바로 호출할 수 있는 상태가 된다.

	따라서 위의 Example_taskTestAll() 테스트는 그대로 성공하게 된다. Task의 OverDue 메서드를
	삭제하여도 성공한다.


	몇 가지 장점이 있지만 여기서는 Deadline을 내장하지 않는다. 이는 Json을 공부하기 위해 그런 것인데,
	필드가 내장되어 있으면 내장된 필드가 구조체 전체의 직렬화 결과를 바꿔버리는 문제가 있다. 오히려 다음과
	같이 Deadline 자료형을 구조체로 만들어 time.Time을 내장한다. (이것은 Deadline_2로 표현하겠다.)
*/

type Deadline_2 struct {
	time.Time
}

func NEwDeadline(t time.Time) *Deadline_2 {
	return &Deadline_2{t}
}

/*
	구조체를 내장하게 되면 내장된 구조체에 들어 있는 필드들도 바로 접근이 가능하게 된다.
	따라서 구조체 내장을 이용하면 여러 구조체에 있는 필드들이 모두 합쳐진 구조체 같은 것을
	만들 수 있다.

	물론 실제 합쳐진 것은 아니고 위와 같은 과정과 같이 내부에 구조체를 필드로 갖고 있지만 필드에
	접근할 때 궅이 길게 쓸 필요가 없다는 것이다.

*/

type Address struct {
	City  string
	State string
}

type Telephone struct {
	Mobile string
	Direct string
}

type Contact struct {
	Address
	Telephone
}

func Example_contact() {
	var c Contact
	c.Mobile = "123-456-789"
	fmt.Println(c.Telephone.Mobile)
	c.Address.City = "San Franciso"
	c.State = "CA"
	c.Direct = "N/A"
	fmt.Println(c)
	// Output:
	// .
}

/*
	위의 코드에서 보시면 c.Moblie에 접근했지만 c.Telephone.Mobile에 접근한 것과 마찬가지 효과가
	나타났다. 상속과 같은 개념과는 달리 실제로는 내부에 필드로 내장하고 있으면서 편의를 제공하는 것
	뿐이라는 것 명심해야 한다.

	구조체 내장을 활용한 디자인 문제들은 인터페이스를 배운 뒤에 다시 살펴볼 것이다.
	(이 다음은 json인데 이는 넘기겠다.)


*/
