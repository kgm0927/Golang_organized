package chapter3

/*3.1.1 유니코드 처리

유니코드는 전 세계의 모든 문자를 일관되게 표현하기 위한 산업 표준이다.

Go 언어의 소스 코드는 UTF-8로 되어 있다. 따로 코드 상에 표시된 문자열 역시 UTF-8로 인코딩되어 있다.
이 문자열이 자연스럽게 UTF-8로 인코딩된 문자열이 된다. 그런데 UTF-8은 유니코드 포인트를 나타내기 위한 바이트 수가
가변적이다. 0부터 127까지는 1 바이트로 표현되지만, 6바이트를 늘어날 수 있다.const


rune형태: 고대 유럽 어족에서 쓰이던 문자로서 여기서는 int32 즉 32비트 정수의 별칭(alias)이며, 유니코드 포인트 하나를 담을 수 있다.

그리고 함수 len은 바이트 단위로 출력하기 때문에 9를 입력한다.*/
var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numdEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numdEnds != 0
		}
	}
	return result
}

/*
HasConsonantSuffix 함수를 한글이 아닌 문자가 있는 경우는 무시하고 지나친다. 어쨌든 s에 있는
유니코드를 하나씩 r로 꺼내고, 이것이 "가"에서 "힣"까지의 한글이면 result를 업데이트한다.

r-start 하면 몇 번째 한글인지가 나오고 (즉, "가"의 경우에는 0, "각"은 1, ...), 이것을 28로 나누어서
나머지가 0이면 받침이 없는 것이다.
*/
