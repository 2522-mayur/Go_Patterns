package main

import(
	"fmt"
)


func PEven(n int,evenChan chan bool,oddChan chan bool){
defer wg.Done()
	for i:=2;i<=n;i=i+2{
		<-oddChan
		fmt.Println(i)
		if i<n{
			evenChan<-true
		}else{
			close(evenChan)
		}
	}
}

func pOdd(n int,evenChan chan bool,oddChan chan bool){
defer wg.Done()
	for i:=1;i<=n;i=i+2{
		<-evenChan
		fmt.Println(i)
		if i<n{
			oddChan<-true
		}else{
			close(oddChan)
		}
	}
}

func alternateEvenOdd(){
	oddChan:=make(chan bool)
	evenChan:=make(chan bool)

	wg.Add(2)
	go PEven(10,evenChan,oddChan)
	go pOdd(10,evenChan,oddChan)
evenChan<-true
	wg.Wait()
}