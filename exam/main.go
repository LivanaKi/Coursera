package main

import "fmt"

func foobar(a *int, b int, c int)(int){

	*a = (*a) * b

	return ( *a - c)

}

func main(){

	var a, c int = 3, 4

       b := 1.5

	

	fmt.Println( " answer 1 is  " , foobar(&a, 3, 1) )

	fmt.Println( " answer 2 is  " , foobar(&a, 2, c) )

	fmt.Println( " answer 3 is  ", foobar(&a, int(b), c) )

	a = 2 

       b = 2.5

       fmt.Println(  " answer 4 is  ", foobar(&a,int(2 *b), c -1) )

	fmt.Println(  " answer 5 :  a is " , a , " , b is ", b)

	

}

