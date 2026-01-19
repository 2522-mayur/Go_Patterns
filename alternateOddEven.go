package main 

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printodd(evenchan <- chan int,oddchan chan<- int,n int){
	defer wg.Done()

for i:=1;i<=n;i=i+2{
	<-evenchan
	fmt.Println(i)
	if i<n {oddchan<-1}

}

}
func printeven(evenchan chan<- int,oddchan <-chan int,n int){
defer wg.Done()
for i:=2;i<=n;i=i+2{
	<-oddchan
	fmt.Println(i)
	if i<n {evenchan <-1}

}

}

func alternate(){

	n:=5
	evenchan:=make(chan int)
	oddchan:=make(chan int)

	wg.Add(2)

	go printeven(evenchan,oddchan,n)
	go printodd(evenchan,oddchan,n)

	evenchan <-1
	wg.Wait()
	close(evenchan)
	close(oddchan)



}