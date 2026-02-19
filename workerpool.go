package main

import (
	"fmt"
	//"sync"
	"time"
	"context"
)

//var wg sync.WaitGroup

func helper(id int,jobs <-chan int,ctx context.Context){

	defer wg.Done()
	for{
		select{
		
	case <-ctx.Done():
		fmt.Println("cancelliing gorotine",id)
		return
	case val,ok:=<-jobs:
		if !ok{
			return
		}
		time.Sleep(1*time.Second)
		fmt.Printf("goroutine %d processing value %d\n",id,val)
	}
	}
	// for val:=range jobs{
	// 	fmt.Printf("goroutine %d processes value %d \n",id,val)
	// }
}

func workerpool(){
	jobs :=make(chan int,5)
//	result:=make(chan int ,5)
	ctx,cancel:=context.WithCancel(context.Background())
	for i:=1;i<=3;i++{
		wg.Add(1)
		go helper(i,jobs,ctx)
	}

	for j:=1;j<15;j++{
		jobs<-j
	}
	
	go func(){
		close(jobs)
		time.Sleep(500*time.Millisecond)
		
		cancel()
	}()

	wg.Wait()
	// close(result)

	// for k:=range result{
	// 	fmt.Println(k)
	// }
}