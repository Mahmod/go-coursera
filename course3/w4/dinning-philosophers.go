package main

import (
	"fmt"
	"sync"
	"time"
)

const numberOfPhilosophers = 5
const eatLimit = 3

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	number              int
	leftChopstick, rightChopstick *Chopstick
	permissionChannel   chan bool
}

func (p Philosopher) eat() {
	for i := 0; i < eatLimit; i++ {
		<-p.permissionChannel // Requesting permission to eat
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("round:%d ,starting to eat %d\n",i, p.number)
		time.Sleep(time.Second) // Eating
		fmt.Printf("round:%d ,finishing eating %d\n",i, p.number)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		p.permissionChannel <- true // Informing completion
	}
}

func host(permissionChannel chan bool) {
	for i := 0; i < eatLimit*numberOfPhilosophers; i++ {
		permissionChannel <- true
		<-permissionChannel
	}
	close(permissionChannel)
}

func main() {
	chopsticks := make([]*Chopstick, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, numberOfPhilosophers)
	permissionChannel := make(chan bool, 2) // Buffer of 2 to allow 2 philosophers to eat concurrently

	for i := 0; i < numberOfPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			number:            i + 1,
			leftChopstick:     chopsticks[i],
			rightChopstick:    chopsticks[(i+1)%numberOfPhilosophers],
			permissionChannel: permissionChannel,
		}
	}

	go host(permissionChannel)

	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go func(philosopher *Philosopher) {
			defer wg.Done()
			philosopher.eat()
		}(philosopher)
	}
	wg.Wait()
}
