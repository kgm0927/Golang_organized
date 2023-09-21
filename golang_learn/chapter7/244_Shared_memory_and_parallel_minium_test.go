package chapter7

import (
	"fmt"
	"sync"
)

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}

	return min
}

// 아래는 간단한 예제 및 테스트이다.

func ExampleMin() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}))
	// Output:
	// .
}

/*
	이제 '병렬 버전'을 만들어보겠다. 병렬 버전은 별다른 어려움이 없다. 사람 4명이서 가장 작은 수를 찾는다고 생각해보면,
	전체를 넷으로 나누어서 각자 가장 작은 수를 찾은 다음, 그중에 가장 작은 수를 한 번 더 찾으면 된다. 그래서 아주 쉬운
	병렬(emgarrassingly parallel) 문제이다.



*/

func ParallelMin(a []int, n int) int {
	if len(a) < n { // n은 몇개의 고루틴을 쓸지 정하는 변수이다.
		return Min(a)
	}

	mins := make([]int, n)

	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(i int) {

			defer wg.Done()

			begin, end := i*size, (i+1)*size

			if end > len(a) {
				end = len(a)
			}
			mins[i] = Min(a[begin:end])
		}(i)
	}

	wg.Wait()

	return Min(mins)
}

/*
	위의 예제에서 고루틴 내에 mins[i]라는 결과를 넣어준다. 고루틴들이 메모리를 공유하는 모델이기 때문에 같은 배열에 접근이 가능한 것이다.
	i번째 고루틴은 i번째에 값을 넣게 된다.

*/

func ExampleParallelMin() {
	fmt.Println(ParallelMin([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}, 4))
	// Output:
	// .
}
