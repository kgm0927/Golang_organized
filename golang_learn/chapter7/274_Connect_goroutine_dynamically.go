package chapter7

import "context"

func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

type Intpipe_2 func(context.Context, <-chan int) <-chan int

func FilterMultiPle(n int) IntPipe {

}
