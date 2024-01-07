/*
Dining Philosophers Problem

This is a classic synchronization problem. Imagine you have five philosophers sitting around a table.
Each philosopher has a plate of spaghetti and a fork on each side.
Philosophers can do two things: think or eat.
However, they can’t eat without forks and each philosopher needs both forks to eat their spaghetti.

The challenge is to prevent the philosophers from starving, which happens if they can’t eat,
without getting into a deadlock, where each philosopher picks up one fork and waits forever for the other.

You can use a mutex to represent each fork.
A philosopher can try to lock each mutex (pick up each fork) when they want to eat.
After eating, they will unlock the mutexes (put down the forks).

This problem will help you understand how to use mutexes to handle concurrency and avoid race conditions.
Good luck! 😊
*/
package main

import (
	"fmt"
	"sync"
)

type State int

const (
	Thinking State = iota
	Eating
	Finished
)

type Fork struct {
	fnum string
	mu   sync.Mutex
}

func (f *Fork) pickup() string {
	f.mu.Lock()
	return f.fnum
}

func (f *Fork) putdown() string {
	f.mu.Unlock()
	return f.fnum
}

type Philosopher struct {
	name  string
	state State
	forks [2]*Fork
}

func (p *Philosopher) cycle(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		switch p.state {
		case Thinking:
			fmt.Printf("%s is thinking\n", p.name)
			for _, f := range p.forks {
				fmt.Printf("%s picked up fork %s\n", p.name, f.pickup())
			}
			p.state = Eating
		case Eating:
			fmt.Printf("%s is eating\n", p.name)
			for _, f := range p.forks {
				fmt.Printf("%s put down fork %s\n", p.name, f.putdown())
			}
			p.state = Finished
		case Finished:
			fmt.Printf("%s is full\n", p.name)
			return
		}
	}
}

func main() {
	forks := []*Fork{
		{fnum: "1"},
		{fnum: "2"},
		{fnum: "3"},
		{fnum: "4"},
		{fnum: "5"},
	}
	philosophers := []*Philosopher{
		{
			name: "Carolina",
			forks: [2]*Fork{
				forks[4],
				forks[0],
			},
		}, {
			name: "Guadalupe",
			forks: [2]*Fork{
				forks[0],
				forks[1]},
		}, {
			name: "Api",
			forks: [2]*Fork{
				forks[1],
				forks[2]},
		}, {
			name: "Gonza",
			forks: [2]*Fork{
				forks[2],
				forks[3],
			},
		}, {
			name: "Lean",
			forks: [2]*Fork{
				forks[3],
				forks[4],
			},
		},
	}

	var wg sync.WaitGroup
	for _, p := range philosophers {
		wg.Add(1)
		go p.cycle(&wg)
	}
	wg.Wait()
	fmt.Println("all philosophers are full")
}
