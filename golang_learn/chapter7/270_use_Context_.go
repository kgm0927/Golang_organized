package chapter7

import "context"

//
func PlusOne_con(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for num := range in {
			select {
			case out <- num + 1:
			case <-ctx.Done():
				return

			}
		}
	}()
	return out
}
