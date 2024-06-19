package internal

type AnimalEgg interface {
	MakeNoise()
	Egg()
}
type AnimalMilk interface {
	MakeNoise()
	Milk()
}

type Cat struct{}
type Dog struct{}
type Cow struct{ milk int }
type Goat struct{ milk int }
type Chicken struct{ egg int }
type Ostrich struct{ egg int }
type Quail struct{ egg int }
