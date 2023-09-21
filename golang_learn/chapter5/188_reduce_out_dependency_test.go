package chapter5

import "os"

/*
테스트 기술을 갈고 닦기 위해서는 인터페이스가 필요하다. 많은 경우에 테스트를 할 때,
많은 경우에 테스트를 할 때, 외부 리소스를 접근하는 것을 막고 싶은 경우가 있다.
*/

/*func Save(f *os.File) {

}*/

/*
	이제 이 함수를 테스트하려면 파일을 넘겨줘야 한다. 구조체로 구현되어 있는 파일을
	넘겨받는 대신에 io.Writer와 같은 것을 받으면 실제 구현에서는 파일을 넘기면 되고,
	테스트에서는 파일을 넘기지 않고 bytes.Buffer와 같은 것을 넘겨서 테스트 할 수 있다.


*/

/*func Save(w io.Writer) {

}*/

/*
	가능하면 만들어져 있는 인터페이스를 받아서 동작하게 코드를 작성시키면 유연하게 테스트를 할 수 있다.
	닫는 연산이 필요하면 io.WriterCloser를 io.Writer 대신에 이용할 수 있다. 순차적으로 쓰는 경우가
	아니라면 io.WriteSeeker 및 io.WriterAt을 이용하면 된다.


	더 나아가서 파일 시스템을 이용하는 함수들이 있는 파일의 이름을 바꾸고 파일의 목록을 살펴보는 일을 하는 코드가
	있을 때는 어떻게 테스트를 하는 것이 좋을까? 파일시스템에 접근하는 경우에는 파일시스템 인터페이스를 만들어서 이용하는
	것이 유연성을 높이는데 도움이 된다.


	작성한 프로그램이 파일 이름 변경과 삭제를 해야 하는 경우라면 다음과 같이 인터페이스를 만든다.
*/

type FileSystem interface {
	Rename(oldpath, newpath string) error
	Remove(name string) error
}

type OSFileSystem struct{}

// 실제 구현은 다음과 같이 한다.
func (fs OSFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs OSFileSystem) Remove(name string) error {
	return os.Remove(name)
}

// 이렇게 하고 이제 구현 부분에는 이 인터페이스를 사용합니다.

func ManageFiles(fs FileSystem) {

}

/*
	이렇게 하면 OSFileSystem을 이용하여 호출하면 실제 파일 시스템을 이용하고, 테스트 용도로 가짜 파일
	시스템을 만들어서 이용할 수 있다.

	이렇듯 인터페이스를 잘 활용하면 외부 의존성을 줄이는데 많은 도움이 된다. Go언어의 표준 라이브러리에도
	이와 같은 테크닉을 이용하는 부분이 있다. 아래 주소를 통해 열람이 가능하다.

	golang.org/src/net/http/fs.go

	우리 예제에서는 빈 구조체를 이용했지만, 책이 쓰여지는 현재 표준 라이브러리 코드의 예제는 string형을
	Dir형으로 이름 붙여 정의하고 있다. 그리고 Open 메서드 하나만 이용하기 때문에 해당 메서드를 넣어서
	인터페이스를 정의하고 있다.

	golang.org/src/net/http/fs_test.go


	테스트하는 코드를 보면 testFileSystem이라는 구조체에 open이라는 함수 필드가 있고 이 필드를
	설정하는 것에 따라서 가짜 파일 시스템을 만들 수 있게 되어 있다. Open 메서드가 불릴 때마다
	실제로 파일을 여는 대신에 어떤 일을 할지 정해서 테스트할 수 있는 것이다.

*/
