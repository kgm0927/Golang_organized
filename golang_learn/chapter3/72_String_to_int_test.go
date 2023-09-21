package chapter3

/*
예를 들어 "5"와 같이 문자열 형태로 된 숫자를 5와 같은 정수형으로 바꾸고 싶을 때, 단순히 int('5')를
이용하면 결과값이 5가 아니라 해당 문자의 유니코드의 코드 포인트 숫자로 변환되었으므로 다른 방법을
이용해야 한다.

한가지 방법은 strconv 패키지에 있는 함수를 이용하는 것이다. strconv.Atoi()*와 같이 같은 함수들이
문자열을 정수롤 바꾸저면 strconv.ParseInt()와 같은 함수들을 이용하면 64비트 정수, 혹은 10 진수가 아닌
수를 반환할 수 있다.

실수형은 strconv.ParseFloat()을 이용할 수 있다. 반대로 숫자를 문자열로 변환하는 경우는 strconv.Itoa() 및
strconv.FormatInt()를 이용하면 된다.


var i int
var k int64
var f float64
var s string
var err error
i,err=strconv.Atoi("350")
k,err=strconv.ParseInt("cc7fdd",16,32)
k,err=strconv.ParseInt("0xcc7fdd",0,32)
f,err=strconv.ParseFloat("3.14",64)
s=strconv.Itoa(340)
s=strconv.FormatInt(13402077,16)


다른 방법은 fmt 패키지에 있는 함수들을 이용하는 것이다. fmt.Sscanf()를 사용하여 문자열로부터 숫자 혹은
다른 형식을 읽을 수 있다. 반대로 숫자를 문자열로 바꿀 때에는 fmt.Sprint() 등을 이용하면 된다.


var s string;
s=fmt.Sprint(3.14)
s=fmt.Sprintf("%x",13402077)


*/
