package chapter4

/*
	객체지향을 저원하는 언어들은 메서드들을  public 혹은 private으로 지정하여 접근을 조절할 수 있다.

	Go언어 답게 예약어가 따로 없이 모든 것을 관통하는 하나의 법칙이 있다. 바로 식별자 이름의 첫 글자가
	대문자인지 소문지인지 구분하는 방식이다.

	메서드의 이름이 대문자로 시작하면 해당 메서드는 다른 모듈에서 보이기 때문에 호출이 가능하다. 만일
	메서드의 이름이 소문자로 시작되면 다른 모듈에서는 보이지 않는다.

	마찬가지로 법칙이 모듈의 전역에 정의된 자료형, 변수, 상수, 함수 모두에 적용이 된다. 메서드가 아닌
	함수 역시 대문자로 시작되는 것만 모듈 밖에서 보인다. 명명된 자료형도 마찬가지다.


	한 모듈은 여러 파일로 구성되어 있다. 소문자로 시작하는 이름이라도 같은 모듈을 구성하는 다른 파일에서도
	보인다. 테스트 파일 역시 같은 모듈에 속한다면 소문자로 시작하는 함수나 메서드를 테스트하는데 아무런 문제가
	없다. 클래스 단위로 접근이 제한되기 때문에 공개되지 않은 메서드를 테스트할때 friend 등을 쓰는 c++과 다른 점이다.


*/
