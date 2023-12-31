package chapter4

/*
	4.4.2 path/filepath 패키지

	path/filepath는 파일 이름 경로를 다루는 패키지이다.

	이 함수는 지정된 디렉터리(폴더) 경로 아래에 있는 파일들에 대하여 어떤 일을 할 수 있는
	함수이다. 디렉터리 안에 디렉터리가 있다면 그것도 추적해 들어간다. 음악 파일들을 폴더별로
	모아두었다면 이 함수를 통해 음악들을 모두 재생목록에 넣는 등의 일을 알 수 있다.


	어떤 일을 하게 만들어야 할 지를 미리 정해둔다면 비슷한 패턴의 함수들이 많이 만들어져야 할 것이다.
	특정 디렉터리 밑의 모든 파일의 소유권을 바꾸는 함수, 특정 디렉터리 밑의 모든 파일을 재생하는 함수,
	특정 디렉터리 밑의 모든 파일의 이름을 그저 출력하는 함수, 슬라이스에 넣는 함수 여러 가지가 있을 수 있다.
	이러한 함수들 모두 '고계 함수'로 이루어져 있다.

*/
