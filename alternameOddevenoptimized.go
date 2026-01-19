package main

import(
	"fmt"
)


func p_odd(chans chan int,n int){
	defer wg.Done()
	for i:=0;i<n/2;i++{
		 val,ok:=<-chans
		 if !ok{
			
		 }else{
			fmt.Println(val)
			if val==n{
				close(chans)
				return
			}
			
			chans <-val+1
		 }


	}
}

func p_even(chans chan int,n int){
	defer wg.Done()
	for i:=0;i<=n/2;i++{
		 val,ok:=<-chans
		 if !ok{
			
		 }else{
			fmt.Println(val)
			if val==n{
				close(chans)
				return
			}
			
			chans <-val+1
		 }


	}
}
func optimized(){
	n:=10
	chans:=make(chan int)

	wg.Add(2)
	go p_odd(chans,n)
	go p_even(chans,n)
	chans <-1
	wg.Wait()

}