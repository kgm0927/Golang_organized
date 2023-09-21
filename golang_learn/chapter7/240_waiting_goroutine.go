package chapter7

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

/*
		7.1.2 고루틴 기다리기

	여러 이미지 파일을 동시에 다운로드 받은 뒤에 zip으로 압축하는 프로그램을 만들어본다.

	아래 코드는 main 함수 내에 들어간다.

	func main() {
	var urls = []string{
		"http://image.com/img01.jpg",
		"http://image.com/img02.jpg",
		"http://image.com/img03.jpg",
	}

	for _, url := range urls {
		go func(url string) {
			if _, err := chapter7.Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)

		filenames, err := filepath.Glob("*.jpg")
		if err != nil {
			log.Fatal(err)
		}

		err = chapter7.WriteZip("images.zip", filenames)

		if err != nil {
			log.Fatal(err)
		}
	}

}


*/

// 이제 download 함수를 만들어야 한다.

func Download(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	filename, err := urlToFilename(url)
	fmt.Printf("filename:%s\n", filename)
	if err != nil {
		return "", err
	}

	f, err := os.Create(filename)

	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return filename, err
}

/*
	filepath.Base 함수는 경로에서 가장 마지막 부분을 반환한다. 따라서 저 위의 URL들에
	대하여 img01.jpg, img02.jpg, img03.jpg을 반환한다.

*/

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	return filepath.Base(url.Path), nil
}

func WriteZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)

	if err != nil {
		return err
	}

	zw := zip.NewWriter(outf)

	for _, filename := range filenames {

		w, err := zw.Create(filename)

		if err != nil {
			return err
		}

		f, err := os.Open(filename)

		if err != nil {
			return err
		}

		defer f.Close()
		_, err = io.Copy(w, f)

		if err != nil {
			return err
		}

	}

	return zw.Close()
}

/*
	이 프로그램을 실행하면 어떻게 될까? 원하는 대로 압축이 되나? 파일이 다운로드되기 전에
	압축하려고 들 수 있기 때문에 압축이 되지 않는다. 저 위의 코드에 따르면 각각의 파일을
	다운로드하는 것과 압축하는 것 중 어느 것을 수행하여도 좋은 동시성이 있는 작업들이다.

	그러나 실제로는 그렇지 않는다. 파일 다운로드가 완료되지 않으면 압축도 할 수 없다. 즉,
	압축은 파일들이 모두 다운로드될 때까지 기다린 후 수행해야 된다. 이럴 때 이용할 수 있는
	것이 sync.WaitGroup이다.

	사용법은 그렇게 어렵지 않다. main 함수에서 for 반복문 부분을 이렇게 고치면 된다.

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := chapter7.Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()


*/

/*
	wg에는 기본값이 0으로 맞춰져 있는 카운터가 들어 있다. Wait() 함수는 이 카운터가 0이 될 때까지
	기다린다. Add() 함수는 호출될 때마다 숫자를 더한다.
	wg.Done()은 사실상 wg.Add(-1)과 같다고 보면 되지만,
	다 되었다는 것을 표현해두는 것이 더 알기 쉽기 때문에 wg.Done()을 쓰는 것이 좋다.

	그렇기 때문에 처음에 WaitGroup을 만들자마자 URL 개수만큼 카운터를 증가시키고 각각의 고루틴에서는 작업이
	완료될 때마다 카운터를 감소시킨다. 모든 고루틴이 끝나면 카운터가 0이 되는데, 이 상태가 되기 전까지는
	wg.Wait() 부분에서 멈춰 서 있다. 우리가 원하는 동작이다.


	미리 고루틴이 몇 개 생길지 알기 때문에 이렇게 작성하지만, '고루틴이 몇 개 생길지 알기 어렵거나 따로 수를
	세어야 하는 경우에는' 고루틴을 띄워보내기 전에 wg.Add(1)을 수행하여 하나씩 카운터를 증가시킬 수 있다.


	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if _, err := chapter7.Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()


	마찬가지로 동작한다. 주의할 점은 wg.Add(1)을 '고루틴 내부에 포함시키면 안된다는 점'이다. 고루틴 내부의 wg.Add(1)이 수행되기
	전에 메인 고루틴이 wg.Wait()을 통과할 가능성이 있기 때문이다. 이런 상태를 레이스 컨디션(race condition)이라고 한다.


	나중에는 멀티코어를 이용하여 코드를 작성할 수 있다.

*/
