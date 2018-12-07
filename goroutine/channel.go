package goroutine

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for i, v := range a {
		sum += v
		if i == len(a)-1 {
			fmt.Printf("(%d) = %d\n", v, sum)
		} else {
			fmt.Printf("(%d) + ", v)
		}
	}

	c <- sum //send sum to channel c
}

func TestChannel() {
	a := []int{1, 3, 5, 7, -8}
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		fmt.Printf("x: %2d, now: %d\n", x, time.Now().Nanosecond())
	}
	close(c)
}

func testBufferChannel() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("i: %2d, now: %d\n", i, time.Now().Nanosecond())
		fmt.Println(i)
	}
}

func fibonacciSelect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit...")
			return
		}
	}
}

func testChannelSelect() {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 0
	}()
	fibonacciSelect(c, q)
}

//如果存在多个channel的时候,我们该如何操作呢,Go里面提供了一个关键字select,通过select可以监听channel上的数据流动。
// select默认是阻塞的,只有当监听的channel中有发送或接收可以进行时才会运行,
// 当多个channel都准备好的时候,select是随机的选择一个执行的
