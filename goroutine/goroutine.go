package goroutine

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func thread() {
	runtime.GOMAXPROCS(2)
	go say("world")
	say("hello")
}

func Test() {
	//thread()

	//TestChannel()

	//testBufferChannel()

	testChannelSelect()
}
