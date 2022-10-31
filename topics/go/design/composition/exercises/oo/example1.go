package main

type Animal struct {
	Name string
	mean bool
}

func (a Animal) Walk()  {

}

type Dog struct {
	Animal
	BarkStrength int
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

func (d Dog) Walk() {
}

func (dog *Dog) MakeNoise() {
	barkStrength := dog.BarkStrength

	if dog.mean {
		barkStrength *= 5
	}

	dog.DoNoise(barkStrength, "bark")
}

func (c *Cat) MakeNoise() {
	meowStrength := c.MeowStrength
	if c.Basics.mean {
		meowStrength *= 5
	}
}

type AnimalSounder interface {
	MakeNoise()
}

func (animal *Animal) DoNoise(strength int, sound string) {
	if animal.mean {

	}
}
func main() {
	d := Dog{
		Animal{},
		2,
	}

	_ = d
}
