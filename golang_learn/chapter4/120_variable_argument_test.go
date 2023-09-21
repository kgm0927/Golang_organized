package chapter4

import "io"

/*
	4.1.5 가변인자

	append 함수와 같이, 넘겨받을 수 있는 인자의 개수가 정해져 있지 않는 함수를 만드려면 어떻게 해야 할까?
	단, 이 인자들의 자료형은 동일하다. Go언어는 값을 넘기기 때문에 그 값들을 슬라이스에 담아서 넘기면
	되지 않을까? 다음과 같이 변수 w에 들어 있는 값과 변수 x,y,z에 담겨 있는 정수 값을 넘길 수 있다.


*/

func f(w io.Writer, nums []int) {

}

/*
	동일한 방법이지만 눈으로 보기에 다르게 쓸 수 있다. 이미 3장에서 보았던 것이다.
	이번에는 위에서 쓴 WriteTo를 다음과 같이 바꿔써 본다.

*/

func WriteTo(w io.Writer, lines ...string) (n int64, err error) {

	return
}

/*
	이와 같이 lines를 가변인자로 변경하여도 lines는 슬라이스가 된다. 호출할 때 원래는 슬라이스를
	넘겨주어야 했지만 이제는 다음과 같이 나열하는 것으로 호출할 수 있다. 이 방식을 이용한 함수들은
	append, fmt에 많이 있다.

	WriteTo(w,"hello","world","Go language")
*/

/*
	그러면 이미 슬라이스로 갖고 있는 자료를 가변인자를 두고 있는 함수로 넘기려면 어떻게 해야
	할까? 그냥 슬라이스 하나를 넘기면 그 슬라이스 하나를 담고 있는 슬라이스로 만들어서 넘겨줄
	것이기 때문에 우리가 원하는 대로 되지 않는다.

	그래서 호출 시 점 세개를 붙이면 슬라이스를 넘길 수 있다.

	lines:=[]strings {"hello","world","Go language"}
	WriteTo(w,lines...)

*/
