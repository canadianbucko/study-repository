package main

import (
	"context"
	"fmt"
	"time"
)

func Foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo завершилась!!  :(  ")
			return
		default:
			fmt.Println("ха-ха, я Foo и я спамлю в вывод")
		}
		time.Sleep(300 * time.Millisecond)
	}
}
func Boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("кто-то прибил мой контекст middleCtx :(( я завершаюсь..")
			return
		default:
			fmt.Println("ха-ха, я Boo !!! я спамлю в вывод")
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func Moo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("кто-то прибил мой контекст parentCtx :(( я завершаюсь..")
			return
		default:
			fmt.Println("ха-ха, я Moo хаха")
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)

	go Foo(childCtx)
	time.Sleep(1 * time.Second)
	childCancel()

	go Boo(middleCtx)
	time.Sleep(1 * time.Second)
	middleCancel()

	go Moo(parentCtx)
	time.Sleep(1 * time.Second)
	parentCancel()

	time.Sleep(1 * time.Second)
	fmt.Println("main завершился!")

}
