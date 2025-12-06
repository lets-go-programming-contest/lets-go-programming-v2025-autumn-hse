package handlers
import (
	"context"
)

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		_ = <-ctx.Done()
		return nil
	}
	
}