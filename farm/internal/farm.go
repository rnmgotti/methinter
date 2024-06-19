package internal

import "fmt"

func (d Dog) MakeNoise() {
	fmt.Println("гавгав!")
}

func (c Cat) MakeNoise() {
	fmt.Println("мурмяу")
}

func (c Cow) MakeNoise() {
	fmt.Println("Корова здесь")
}

func (c Cow) Milk() {
	fmt.Println("Количество литров молока у коровы:", c.milk+1)
}

func (g Goat) MakeNoise() {
	fmt.Println("Коза здесь")
}

func (g Goat) Milk() {
	fmt.Println("Количество литров молока у козы:", g.milk+1)
}

func (c Chicken) MakeNoise() {
	fmt.Println("Курица здесь")
}

func (c Chicken) Egg() {
	fmt.Println("количество яиц у курицы:", c.egg+1)
}

func (o Ostrich) MakeNoise() {
	fmt.Println("Страус здесь")
}

func (o Ostrich) Egg() {
	fmt.Println("количество яиц у страуса:", o.egg+1)
}

func (q Quail) MakeNoise() {
	fmt.Println("Перепелка здесь")
}

func (q Quail) Egg() {
	fmt.Println("количество яиц у перепеоки:", q.egg+1)
}

func RollCallAnimalMilk(animals []AnimalMilk) {
	fmt.Println("Перекличка AnimalMilk:")
	for _, animal := range animals {
		animal.MakeNoise()
		animal.Milk()
	}
}

func RollCallAnimalEgg(animals []AnimalEgg) {
	fmt.Println("Перекличка AnimalEgg:")
	for _, animal := range animals {
		animal.MakeNoise()
		animal.Egg()
	}
}

func Farm() {
	cow := Cow{}
	goat := Goat{}
	chicken := Chicken{}
	ostrich := Ostrich{}
	quail := Quail{}

	MilkAnimal := []AnimalMilk{cow, goat}
	RollCallAnimalMilk(MilkAnimal)

	EggAnimal := []AnimalEgg{chicken, ostrich, quail}
	RollCallAnimalEgg(EggAnimal)
}
