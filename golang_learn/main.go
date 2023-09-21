package main

import (
	"context"
	"fmt"

	"github.com/chapter7"
)

func main() {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	nums := chapter7.PlusOne_con(ctx, chapter7.PlusOne_con(ctx, chapter7.PlusOne_con(ctx, chapter7.PlusOne_con(ctx, chapter7.PlusOne_con(ctx, c)))))

	for num := range nums {
		fmt.Println(num)

		if num == 18 {
			cancel()
			break
		}
	}
}
