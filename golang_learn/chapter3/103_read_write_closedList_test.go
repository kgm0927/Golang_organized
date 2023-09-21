package chapter3

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

/*


	3.4.5 그래프의 인접 리스트 읽고 쓰기

	그래프의 정점과 간성으로 이루어진 자료구조이다. 웹을 그래프
	구조로 볼 수 있고, 소셜 네트워크도 그래프 구조로 볼 수 있다.
	이번에 여러 종류의 인접 리스트 구조를 읽어보는 연습을 한다.

	이 형식은 자료가 나오기 전에 해당 자료의 길이를 미리 알려주고 시작하는 방식이다.




*/

func WriteTo_2(w io.Writer, adjList [][]int) error {
	size := len(adjList)

	if _, err := fmt.Fprintf(w, "%d", size); err != nil {
		return nil
	}

	for i := 0; i < size; i++ {
		lsize := len(adjList[i]) // 이차원 배열 개수

		if _, err := fmt.Fprintf(w, "\n%d", lsize); err != nil {
			return nil
		}

		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fprintf(w, " %d", adjList[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}

	return nil
}

func ReadFrom_2(r io.Reader, adjList *[][]int) error {
	var size int

	if _, err := fmt.Fscanf(r, "%d", &size); err != nil {
		return err
	}

	*adjList = make([][]int, size)

	for i := 0; i < size; i++ {
		var lsize int

		if _, err := fmt.Fscanf(r, "\n%d", &lsize); err != nil {
			return err
		}
		(*adjList)[i] = make([]int, lsize) // 역참조

		for j := 0; j < lsize; j++ {

			if _, err := fmt.Fscanf(r, " %d", &(*adjList)[i][j]); err != nil {
				return err
			}

		}

	}
	if _, err := fmt.Fscanf(r, "\n"); err != nil {
		return err
	}
	return nil

}

func TestWriteTo(t *testing.T) {
	adjList := [][]int{
		{3, 4},
		{0, 2},
		{3},
		{2, 4},
		{0},
	}

	w := bytes.NewBuffer(nil)
	if err := WriteTo_2(w, adjList); err != nil {
		t.Error(err)
	}
	expected := "5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n"
	if expected != w.String() {
		t.Logf("expected: %s\n", expected)
		t.Errorf("found: %s\n", w.String())
	}
}

func Example_readFrom2() {
	r := strings.NewReader("5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n")
	var adjList [][]int
	if err := ReadFrom_2(r, &adjList); err != nil {
		fmt.Println(err)
	}
	fmt.Println(adjList)
	// Output:
	// [[3 4] [0 2] [3] [2 4] [0]]

}
