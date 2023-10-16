package chapter7

import (
	"context"
	"fmt"
)

func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

type Intpipe_2 func(context.Context, <-chan int) <-chan int

func FilterMultiPle(n int) Intpipe_2 {
	return func(ctx context.Context, in <-chan int) <-chan int {

		out := make(chan int)

		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}

				select {
				case out <- x:
				case <-ctx.Done():
					return
				}

			}
		}()

		return out

	}
}

/*
	FilterMultiple은 n의 배수를 걸러내는 파이프라인을 반환한다. 이 패턴도 눈에 익혀두시면 좋다.
	클로져를 이용하여 함수를 반환한 이유는 파이프라인 함수형을 맞춰서 다른 파이프라인에 연결해서
	안전하게 쓸 수 있게 하기 위함이다.


*/

func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)

		for {

			select {
			case i := <-c:
				c = FilterMultiPle(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

/*
	Primes는 무한 소수 생성기이다. 여기서 어려운 게 <-c와 같이 받는 부분에서 막혀 있을 때
	ctx가 취소가 될 수 있고 여기서 받은 뒤에 그 값을 out<-i와 같이 보낼 때 막혀 있다가
	ctx 취소가 될 수 있기 때문에 위의 형태처럼 select를 다중으로 만들어주어야 한다.
	이렇게 작성하여야 취소되었을 때 막혀서 계속 살아있는 좀비 고루틴이 없어지게 된다.


*/

func PrintPrimes(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for prime := range Primes(ctx) {
		if prime > max {
			break
		}
		fmt.Println(prime, " ")
	}
	fmt.Println()
}

/*
위는 Primes를 이용하는 코드 예제이다. max 숫자가 나올 때까지 Primes에서 소수를 순서대로 꺼내어 이용하다가 범위가
넘어버리면 반복문을 빠져 가간다. defer cancel()이 있기 때문에 함수가 종료되었을 때, Primes 함수로 넘어간 ctx가
취소되고, 생성되었던 고루틴들이 모두 소멸될 수 있게 된다.
*/
