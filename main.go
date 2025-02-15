package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

var (
	residentsEvents = map[string][]func(r *Resident){
		"Наконец-то, найден спутник в жизни!!!": {
			func(r *Resident) { r.Married = true },
		},
		"Устроился на новую работу": {
			func(r *Resident) {},
		},
		"Поругался с супругой/ом.": {
			func(r *Resident) {},
		},
		"Развод, больше я не в браке.": {
			func(r *Resident) { r.Married = false },
		},
		"Смерть": {
			func(r *Resident) { r.Alive = false },
		},
	}

	animalsEvents = map[string][]func(a *Animal){
		"Покусал прохожего": {
			func(a *Animal) {},
		},
		"Сломал лапу": {
			func(a *Animal) {},
		},
		"Смерть": {
			func(a *Animal) { a.Alive = false },
		},
	}
)

type Village []VillageElement

func (v *Village) AddElement(e VillageElement) {
	*v = append(*v, e)
}

func (v *Village) DeleteElement(e VillageElement) {
	if slices.Contains(*v, e) {
		index := slices.Index(*v, e)
		*v = append((*v)[:index], (*v)[index+1:]...)
	}
}

func (v *Village) UpdateAll() {
	for _, elem := range *v {
		elem.Update()
	}
}

func (v *Village) ShowAllInfo() string {
	result := ""
	for _, elem := range *v {
		result += elem.FlushInfo()
	}
	return result + "\n"
}

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func (r Resident) MarriageStatus() string {
	return map[bool]string{true: "в браке", false: "холост"}[r.Married]
}

func (r *Resident) Update() {
	if r.Alive {
		r.AddAge()
		r.RandomEvent(1)
	}
}

func (r *Resident) AddAge() {
	r.Age++
}

func (r *Resident) FlushInfo() string {
	result := fmt.Sprintf("Житель %s (возраст: %d), статус: %s.\n", r.Name, r.Age, r.MarriageStatus())
	if len(r.Events) == 0 {
		return result + "События: Нет\n\n"
	}
	result += fmt.Sprintf("События: \n%s\n\n", strings.Join(r.Events, "\n"))
	r.Events = r.Events[:0]
	return result
}

func (r *Resident) RandomEvent(chance int) {
	EventList := []string{}
	for k := range residentsEvents {
		EventList = append(EventList, k)
	}
	randomChoice := rand.Intn(chance)
	if randomChoice == 0 {
		event := EventList[rand.Intn(len(EventList))]
		r.Events = append(r.Events, event)
		for _, Func := range residentsEvents[event] {
			Func(r)
		}
	}
}

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

func (a *Animal) Update() {
	if a.Alive {
		a.AddAge()
		a.RandomEvent(2)
	}
}

func (a *Animal) AddAge() {
	a.Age++
}

func (a *Animal) FlushInfo() string {
	result := fmt.Sprintf("Животное %s (%s, возраст: %d).\n", a.Name, a.Type, a.Age)
	if len(a.Events) == 0 {
		return result + "События: Нет\n\n"
	}
	result += fmt.Sprintf("События: \n%s\n\n", strings.Join(a.Events, "\n"))
	a.Events = a.Events[:0]
	return result
}

func (a *Animal) RandomEvent(chance int) {
	EventList := []string{}
	for k := range animalsEvents {
		EventList = append(EventList, k)
	}
	randomChoice := rand.Intn(chance)
	if randomChoice == 0 {
		event := EventList[rand.Intn(len(EventList))]
		a.Events = append(a.Events, event)
		for _, Func := range animalsEvents[event] {
			Func(a)
		}
	}
}

func main() {
	village := Village{}

	// Создаем жителей деревни
	resident1 := &Resident{Name: "Алиса", Age: 30, Married: false, Alive: true, Events: []string{}}
	resident2 := &Resident{Name: "Борис", Age: 40, Married: true, Alive: true, Events: []string{}}

	// Создаем животных
	animal1 := &Animal{Name: "Бобик", Age: 5, Type: "собака", Alive: true, Events: []string{}}
	animal2 := &Animal{Name: "Мурка", Age: 3, Type: "кошка", Alive: true, Events: []string{}}

	// Добавляем элементы в деревню
	village.AddElement(resident1)
	village.AddElement(resident2)
	village.AddElement(animal1)
	village.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		village.UpdateAll()
		fmt.Println(village.ShowAllInfo())
	}
}
