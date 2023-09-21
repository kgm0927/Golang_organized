package chapter3

import "fmt"

/*
	3.2.4 슬라이스 용량

	슬라이스는 연속된 메모리 공간을 활용하는 것이라고 용량에 제한이 있을 수 밖에 없다.

	make([]int,5)와 같이 다섯 개의 빈 공간을 미리 할당하거나 []int{0,0,0,0,0}과 같이 다섯
	개의 정수로 초기화한 경우는 길이뿐만 아니라 용량도 5로 맞춰지게 된다. 공간의 낭비가 없다.
	그런데 여기서 슬라이스 덧붙이기를 이용하여 하나를 더 덧붙이면 용량이 부족하므로 슬라이스 전체를
	복사하게 된다.
*/

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2))
	fmt.Println()
	sliced3 := nums[:4]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()
	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)

	// Output:
	// .
}

/*
	얼마나 덧붙일 공간이 있느냐에 따라서 용량이 결정되므로, 뒤에 2개를 잘라낸 경우에는
	길이는 2만큼 줄어들지만, 기둥 뒤에 공간이 있듯이 여전히 2만큼 공간이 뒤에 더 있으므로
	용량은 여전히 5가 된다. 반면 앞에 2개를 잘라낸 경우에는 길이가 2만큼 줄어들고, 뒤에 공간이
	없으므로 용량도 3으로 줄어든다.


	잘라냈더라도 뒤에 공간이 있으면 공간을 살릴 수도 있다. sliced3은 nums에서 잘라낸 것이 아니라 이미
	잘라진 sliced1의 뒷 공간을 살려낸 것이다.

	마지막으로 nums[2]에 100을 넣어보았다. 사실 nums에서 잘려진 슬라이스들은 모두 동일한 메모리를 보고
	있으므로 맨 마지막에서 보는 것과 같이 nums[2]만 수정해도 다른 슬라이스 모두 수정이 일어난다.


	슬라이스의 길이를 지정해서 새로 생성할 수 있듯이 용량 역시 미리 지정해서 생성할 수 있다.

	nums:=make([]int,3,5)

	길이는 3이지만 용량은 5인 슬라이스를 만든 것이다. 사실상 다음과 동일하다.const

	nums:=make([]int,5)
	num=nums[:3]

	길이는 3이지만 용량은 5인 슬라이스를 만든 것이다. 이는 다음과 같다.

	nums:=make([]int,5)
	nums=nums[:3]


	빈 슬라이스를 생성하고 싶지만, 미리 공간을 예약하고 싶을 때가 있는데, 즉, N개까지 길이가 늘어나더라도
	복사가 일어나지 않게 하고 싶은 경우다. 추가될 개수를 미리 예측할 수 있으면 성능에 도움이 된다.

	nums:=make([]int,0,15)

	위의 경우는 15개까지 미리 예약을 해 두었으므로 nums=append(nums,x)한 경우에도 15개를 넘어가지 않으면
	복사가 일어나지 않는다.
*/
