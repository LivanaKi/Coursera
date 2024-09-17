package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func setup(){
	fmt.Println("Init")
}

func dostuff(){
	once.Do(setup)
	fmt.Println("hello")
	wg.Done()
}


func main(){
	wg.Add(2)
	go dostuff()
	go dostuff()
	wg.Wait()
}
/*
var i int = 0
var mut sync.Mutex
var wg sync.WaitGroup

func inc() {
	i = i + 1
	wg.Done()
}


func inc(){
	mut.Lock()
	i = i + 1
	defer wg.Done()
	mut.Unlock()
}
func main(){
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}*/