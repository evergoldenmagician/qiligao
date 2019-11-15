package main

import"fmt"

func main(){
	r:=2


	defer func(){
		r++
	fmt.Println("日月共辉",r)
	}()
r++
	fmt.Println("究极日月",r)


}
