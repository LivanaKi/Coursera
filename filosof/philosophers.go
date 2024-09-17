package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func (p Philo) lockChops() {
	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Println("starting to eat", p.id+1)
}

func (p Philo) unlockChops() {
	p.leftCS.Unlock()
	p.rightCS.Unlock()

	fmt.Println("finishing eating", p.id+1)
}

func philo(id int, p *Philo) {
	host := make(chan *Philo, 2)
	defer wg.Done()
	for i := 0; i < 3; i++ {
		host <- p
		p.lockChops()
		time.Sleep(time.Millisecond * time.Duration(100+id*50)) 
		p.unlockChops()
		<- host
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make(map[int]*Philo)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}
	for i, _ := range philos {
		wg.Add(1)
		go philo(i, philos[i])
	}
	wg.Wait()

}
