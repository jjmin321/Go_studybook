package main

import "fmt"

func main() {
	MAX := 5
	var queue [5]int
	front, rear := 0, 0

	enqueue := func(input int) {
		if (rear+1)%MAX == front {
			fmt.Println("enqueue : Queue overflow")
			return
		}
		rear = (rear + 1) % MAX
		queue[rear] = input
		fmt.Println("enqueue : ", queue)
	}

	dequeue := func() {
		if front == rear {
			fmt.Println("dequeue : Queue underflow")
			return
		}
		front = (front + 1) % MAX
		fmt.Println("dequeue : ", queue[front])
	}

	print_queue := func() {
		fmt.Print("print_queue : ")
		for i := (front + 1) % MAX; i != (rear+1)%MAX; i = (i + 1) % MAX {
			fmt.Print(queue[i], "\t")
		}
	}

	enqueue(7)
	enqueue(4)
	enqueue(13)
	enqueue(11)
	enqueue(15)
	dequeue()
	dequeue()
	enqueue(16)
	print_queue()

}
