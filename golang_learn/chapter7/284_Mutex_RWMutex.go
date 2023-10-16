package chapter7

import "sync"

/*
	7.4.4 Mutex와 RWMutex

	뮤텍스(Mutex)는 상호 배타 잠금 기능이 있다. 동시에 둘 이상의 고루틴의 코드의 흐름을
	제어할 수 있다. 채널을 잘 활용하면 많은 경우에 복잡한 뮤텍스를 이용할 필요가 없다.
	그러나 외부 자원에 접근하는 경우 등에 이것을 활용하면 효과적인 경우가 있다.

	뮤텍스를 가장 활용하는 방법 중 하나는 다음과 같은 접근하고자 하는 자원 포인터와 뮤텍스
	포인터를 하나의 구조체에 넣어두고 사용하는 방법이다.


*/

type Resource int

type Accessor struct {
	L *sync.Mutex
	R *Resource
}

/*
	생성할 때 Accessor{&resource, &sync.Mutex{}}와 같이 할당해주고 이 자원에 접근하는 메서드들에서
	Lock을 잘 활용하면 됩니다.

*/

func (acc *Accessor) USe() {
	// do something
	acc.L.Lock()

	// Use acc.R
	acc.L.Unlock()

	// Do something else
}

/*
	sync.RWMutex는 좀 더 복잡하다. 어떤 자원에 여러 프로세스가 동시에 읽어가는 상관없지만, 프로세스 하나라도
	쓰기를 한다면 다른 어떤 프로세스도 그 동안에 접근할 수 없는 경우에 이용된다. 물론 sync.Mutex를 이용하여
	읽어가는 것 역시 프로세스 하나만 허용하여도 문제없이 동작은 한다.

	그러나 성능이 많이 저하될 것이다. Go에서 기본으로 제공하는 맵이 RWMutex를 이용하기 적합한 성질을 가지고 있다.
	스레드 안전하지 않지만 읽기만 동시에 접근이 가능하다.

*/

type ConcurrentMap struct {
	M map[string]string
	L *sync.RWMutex
}

func (m ConcurrentMap) Get(key string) string {
	m.L.Lock()
	defer m.L.RUnlock()
	return m.M[key]
}

func (m ConcurrentMap) Set(key, value string) {
	m.L.Lock()
	m.M[key] = value
	m.L.Unlock()
}

/*
package main

import (
	"fmt"
	"sync"

	"github.com/golang_learn/chapter7"
)

func main() {

	m:=chapter7.ConcurrentMap{map[string]string{},&sync.RWMutex{}}
}


위의 코드는 간단한 스레드 안전 맵을 구현한 것이다. 다른 프로젝트에도 응용할 수 있을 것이다.const

RWMutex도 Mutex의 일종이다. 그렇기에 RLock과 RUnlock을 사용하지 않고 Lock과 UNlcok만 사용하면
Mutex와 동일하다. 이 특성은 Mutex를 넘겨야 하는 함수에 RWMutex를 넘겨서 활용할 수도 있음을 의미한다.
Mutex와 RWMutex 모두 sync.Locker 인터페이스를 구현하고 있기 때문에 이 함수가 Mutex 대신에 Locker를 받는
경우에 이것이 가능하다. 만일 RWMutex를 넘겨줄 때 RLocker로 동작하게 하고 싶다면 RLocker를 호출하면 된다.

*/
