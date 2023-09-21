package chapter3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 다음은 단순히 문자열 슬라이스를 출력하는 함수이다.
//

func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {

		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}

	}
	return nil
}

func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ExampleWriteTo() {
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	// Output:
	// bill@mail.com
	// tom@mail.com
	// jane@mail.com

}

func ExampleReadFrom() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}
