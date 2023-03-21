package main

import (
	"fmt"
	"log"
)

type Person struct {
	name string
	age  uint8
}

func (p *Person) incAge() {
	p.age++
	if p.age > 30 {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%s has a new life, he (she) is %d years old", p.name, p.age)
				if p.age > 35 {
					panic(fmt.Sprintf("%s super old!!!", p.name))
				}
			}
		}()
		panic(fmt.Sprintf("%s too old!!!", p.name))
	}
}

func main() {

	alex := &Person{
		name: "Alex",
		age:  27,
	}

	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s, welcome to club: 'А вот в наши времена'", alex.name)
		}
	}()
	for i := 0; i < 20; i++ {
		alex.incAge()
	}
	fmt.Printf("Person data: %+v\n", alex)
}
