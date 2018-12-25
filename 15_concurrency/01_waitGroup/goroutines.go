package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arr := []int{}
	chunk1 := 0
	chunk2 := 0
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Println("arr =", arr)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			chunk1 += arr[i]
		}

		fmt.Println("chunk1 =", chunk1)
	}()

	go func() {
		defer wg.Done()
		lengthArr := len(arr)
		for i := 5; i < lengthArr; i++ {
			chunk2 += arr[i]
		}
		fmt.Println("chunk2 =", chunk2)
	}()
	wg.Wait()

	fmt.Println("Total =", chunk1+chunk2)
	// wg.Wait()

	time.Sleep(2 * time.Second)
}

/*
	Như vậy có thể kết luận tất cả các hàm ở bên dưới wg.Wait() chỉ có thể thực thi khi wg done hết.
	Mỗi hàm trong wg.Done() trong goroutines sẽ thực thi khi goroutines chạy xong.
	Mỗi lần wg.Done() sẽ giảm wg đi 1 và cho tới khi wg về 0 wg.Wait() mới bắt đầu cho phép chức năng chạy xuống bên dưới.
*/
