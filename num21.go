package main

import "fmt"

// класс собака
type Dog struct{}

// реакция собаки
func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав!")
}

// класс кошка
type Cat struct{}

// реакция кошки, она немного посложнее и если ее не позвать - она молчит
func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Мяу-мяу...")
	}
}

// целевой интерфейс - Target
type AnimalAdapter interface {
	Reaction()
}

// адаптер для собаки
type DogAdapter struct {
	*Dog
}

// реакция собаки
func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}

// конструктор адаптера для собаки
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// адаптер для кошки
type CatAdapter struct {
	*Cat
}

// реакция кошки
func (adapter *CatAdapter) Reaction() {
	// адаптер автоматически зовет кошку isCall = true
	adapter.MeowMeow(true)
}

// конструктор адаптера для кота
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

// класс - жена
type Wife struct {
}

// реакция жены - адаптер не нужен, нужный метод итак есть
func (w *Wife) Reaction() {
	fmt.Println("Здравствуй, Дорогой")
}

// основной метод для демонстрации
func main() {
	myFamily := [3]AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{}), &Wife{}}
	//
	fmt.Println("Открываете дверь и заходите домой\n")
	for _, member := range myFamily {
		member.Reaction()
	}
}
