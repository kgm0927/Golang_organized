package chapter5

import (
	"container/heap"
	"fmt"
)

/*
		힙


	힙은 자료 중에 가장 작은 값을 O(log N)의 시간 복잡도로 꺼낼 수 있는 자료구조이다. 여기서
	가장 작은 값이라 하면 정렬의 순서상 가장 먼저 나오는 자료를 말한다.


*/

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

func ExampleCaseInsensitive_heap() {

	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})

	heap.Init(&apple)
	for apple.Len() > 0 {
		fmt.Println(heap.Pop(&apple))
	}
	// Output:
	// .

}

/*
	힙 정렬의 예이다. 단순히 정렬을 하는 것이라면 굳이 힙 정렬을 쓸 필요는 없을 수 있다.
	시간 복잡도로 보자면 힙 정렬은 O(n log n)이지만 일반적으로 빠른 정렬보다 느리고 메모리를
	여기저기 임의로 엑세스를 해야 하기 때문에 캐시를 효율적으로 사용할 수 없어서 더욱 느릴 수 밖에
	없다.

	그러나 무작위 빠른 정렬이나 합병 정렬은 모두 끝난 뒤에 정렬된 자료를 이용할 수 있지만 힙 정렬은
	좀 더 일찍 첫 자료를 받아볼 수 있다. 정렬이 다 끝나지 않은 상황에서도 지금까지 정렬된 자료를 이용할
	수 있는 것은 선택 정렬(selection sort)의 특징인데, 힙 정렬도 선택 정렬의 일종으로 볼 수 있다.

	힙 정렬이 효율적인 선택 정렬이기 때문에 이용가치가 충분히 있다. 느린 저장 장치에 자료를 저장하는
	경우라든지 자료를 정렬된 순서대로 받아보다가 어떤 조건을 만족하는 상황에서 더 이상 자료가 없어지는
	경우에 유용하게 이용할 수 있는 알고리즘이다.


	생산자와 소비자가 따로 있고 소비자는 현재 순서를 기다리고 있는 값 중에서 가장 작은(혹은 가장 큰)값부터
	소비해야 한다면 힙을 이용하는 것이 매우 효율적일 수 있다. 이것은 우선순위 큐라고 한다.


*/
