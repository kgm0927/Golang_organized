package chapter3

/*
		3.4.1 io.Reader와 io.Writer


입출력은 io.Reader와 io.Writer 인터페이스와 파생된 다른 인터페이스를 이용한다. 각각은 간단히 말해서
바이트들을 읽고 쓸 수 있는 인터페이스이다.


fmt 패키지에는 F로 시작하는 함수들이 io.Reader와 io.Writer를 인자를 받는다.


fmt.Fprintln(os.Stdout,s)는 fmt.Println(s)와 동일하고, fmt.Fprintf(os.Stdout,format, ...)은 fmt.Println(format, ...)과
동일하다.


따라서 기본 입출력시 역시 파일을 읽고 쓰는 것과 거의 동일한 방법으로 일고 쓸 수 있다. 함수를 작성할 때, io.Reader 혹은 io.Writer 등을
받아서 처리하게 하면 입출력, 파일, 네트워크 등 모두 적용할 수 있으며, 테스트 등을 할 때 좋다. 따라서 파일 입출력을 하는 상황으로 가정하고
설명을 주지만 실제로 코드 작성은 io.Reader와 io.Writer기준으로 한다.



	3.4.2 파일 읽기


	f,err:=os.Open(filename)
	if err!=nil{
		return err		// 혹은 다른 에러 처리
	}

	defer f.Close()

	var num int
	if _, err:=fmt.Fscanf(f,"%d \n",&num); err==nil{
			// 읽은 num 값 사용
		}


	파일은 다음과 같이 읽을 수 있다.


	os.Open()은 반환값이 둘이다. 하나는 파일 오브젝트이고 다른 하나는 에러이다. 이에 에러값이 nil이 되면
	성공적으로 파일을 연 것이 된다. 그렇게 f를 사용한다. 파일을 열지 못한 경우에는 앞으로 진행이 되지 않은
	경우가 많으므로 대부분은 해당 에러를 반환할 수 있고, 반환이 일어나다가 에러를 처리할 수 있는 곳에서 다른
	방법을 이용하거나 에러 로그를 남기거나 아니면 프로그램을 중단시킬 수 있다.

	여튼 파일에 제대로 열리지 않은 경우에는 if문 다음 줄로 계속 프로그램이 실행되게 해서는 안된다.


	'defer'는 해당 함수를 벗어날 때 호출할 함수를 등록하는 역할을 한다. 함수나 반복문을 빠져나가는 곳이 한 군데가
	아닌 경우가 많은데 저렇게 파일을 열거나 리소스를 획득하는 것에 성공한 경우는 defer를 이용하면 깔끔하다.
	수행은 역순으로 되기 때문에 순서가 꼬이는 경우가 적다.


	여기서 fmt.Fscanf()를 이용하여 파일을 읽어보았다. %d는 10진수 정수 형태로 읽는다. _부분에서는 몇 개를 읽었는지 의미가
	없으니 밑줄로 무시한다.
*/
