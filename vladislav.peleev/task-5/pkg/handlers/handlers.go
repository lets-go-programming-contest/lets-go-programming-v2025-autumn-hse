package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

func PrefixDecoratorFunc(ctx context.Context, inputChan chan string, outputChan chan string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg, ok := <-inputChan:
			if !ok {
				return nil
			}

			// проверка условия
			if strings.Contains(msg, "no decorator") { // ← по условию задачи
				return errors.New("can't be decorated")
			}

			// добавление префикса, если его нет
			if !strings.HasPrefix(msg, "decorated: ") {
				msg = "decorated: " + msg
			}

			// отправка результата
			select {
			case outputChan <- msg:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, inputChan chan string, outputChans []chan string) error {
	if len(outputChans) == 0 {
		return nil
	}

	counter := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg, ok := <-inputChan:
			if !ok {
				return nil
			}

			target := outputChans[counter%len(outputChans)]

			select {
			case target <- msg:
			case <-ctx.Done():
				return ctx.Err()
			}

			counter++
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputChans []chan string, outputChan chan string) error {
	if len(inputChans) == 0 {
		return nil
	}

	var wg sync.WaitGroup

	for _, ch := range inputChans {
		wg.Add(1)

		go func(in chan string) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case msg, ok := <-in:
					if !ok {
						return
					}

					// фильтрация
					if strings.Contains(msg, "no multiplexer") {
						continue
					}

					select {
					case outputChan <- msg:
					case <-ctx.Done():
						return
					}
				}
			}
		}(ch)
	}

	// ожидание завершения всех горутин
	go func() {
		wg.Wait()
	}()

	<-ctx.Done()
	return ctx.Err()
}
