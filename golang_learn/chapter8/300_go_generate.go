package chapter8

import (
	"flag"
	"strings"
	"text/template"
)

/*
	8.2.4 go generate

	매크로를 활용해야 하는 경우라면 지금까지 나열한 방법으로 해결되지 않는다.
	C 언어 등에서 제공하는 매크로는 전처리기를 통하여 소스 코드를 확장하여 컴파일한다.
	비슷한 도구가 go 도구에 있다.

	go generate를 이용하면 임의의 명령을 수행하여 프로그램 코드를 생성할 수 있다.
	//go:generate 뒤에 명령을 붙여서 코드에 넣은 뒤에 go generate를 수행하면 된다.


	enum값을 const를 이용하여 지정하였는데, 이것을 문자열로 바꾸고 싶으면 반복적인
	코드를 작성해야 한다. C 언어에서는 매크로를 이용하여 쉽게 해결할 수 있는데, Go에서는
	stringer를 활용해야 한다. 아래 명령으로 stringer를 설치한다.





	4장에서 만든 MutliSet이 있을 것이다. 문자열로 된 자료를 중복해서 집어넣을 수 있는 컨테이너이다.
	문자열뿐만 아니라 다양한 자료형에 대해 동작하게 만들고 싶다. 그러나 템플릿이나 매크로를 지원하지
	않기 때문에 굳이 오히려 go generate를 이용하는 것이 편리하다.


	6장에서 배운 탬플릿을 이용하여 소스 코드를 생성하는 코드를 작성했다. html/template과 달리 text/template을 이용한다.


	명령줄 플래그로 package_name, multiset_typename, element_typename, output을 받는다. 이중에서는 앞에 3개는
	코드를 생성할 때는 쓰이는 문자열이고 output은 출력할 파일 이름이다.


*/

var (
	packageName = flag.String(
		"package_name",
		"main",
		"package name",
	)
	multisetTypename = flag.String(
		"multiset_typename",
		"MultiSet",
		"container type",
	)

	elementTypename = flag.String(
		"element_typename",
		"string",
		"element type",
	)
	output = flag.String(
		"output",
		"",
		"output filename",
	)
)

var tmpl = template.Must(template.New("multiset").Parse(`package {{.PackageName}}

import "fmt"

type {{.MultisetTypename}} map[{{.ElementTypename}}]int

func New{{.MultisetTypename}}() {{.MultisetTypename}}{
	return {{.MultisetTypename}}{}
}

func (m {{.MultisetTypename}}) Insert(val {{.MultisetTypename}}){
	m[val]++
}


func (m {{.MultisetTypename}})Erase(val {{.MultisetTypename}}){
	if _,exists:=,[val];!exists{
		return
	}
	m[val]--

	if m[val]<=0{
		delete(m,val)
	}
}

func (m {{.MultisetTypename}}) Count(val {{.MultisetTypename}}) int{
	return m[val]
}

func (m {{.MultisetTypename}}) String()string{
	vals:=""
	for val,count:=range m{
		for i:=0; i<count; i++{
			vals+=fmt.Sprint(val)+ " "
		}
	}
	return "{ "+ vals+ " }"
}

`))

// outputFilename returns a filename either output string if not empty
// or lowercased multisetTypename.go.

func outputFilename(output, multisetTypename string) string {
	if output != "" {
		return output
	}
	return strings.ToLower(multisetTypename + ".go")
}
