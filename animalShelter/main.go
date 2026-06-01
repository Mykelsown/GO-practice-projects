package main

import "fmt"

type Animal interface {
	Sound() string
	Fee() float64
}

type Dog struct {
	Name        string
	Breed       string
	AdoptionFee float64
}

type Cat struct {
	Name        string
	Indoor      bool
	AdoptionFee float64
}

// ================= Interface Methods +++++++++++++++
func (d Dog) Sound() string {
	return "woof"
}

func (c Cat) Sound() string {
	return "meow"
}

func (d Dog) Fee() float64 {
	return d.AdoptionFee
}

func (c Cat) Fee() float64 {
	return c.AdoptionFee
}

// +++++++++++++++++ STRINGER ===========================
func (d Dog) String() string {
	return fmt.Sprintf("DOG: %s (%s) - $%.2f", d.Name, d.Breed, d.AdoptionFee)
}

func (c Cat) String() string {
	location := "Outdoor"
	if c.Indoor{
		location = "Indoor"
	}
	return fmt.Sprintf("CAT: %s [%v] - $%.2f", c.Name, location, c.AdoptionFee)
}

func Cheapest[T Animal](animals []T) T {
	lowest := animals[0].Fee()
	cheapestAnimal := animals[0]
	for _, animal := range animals[1:] {
		if animal.Fee() <= lowest {
			lowest = animal.Fee()
			cheapestAnimal = animal
		}
	}
	return cheapestAnimal
}

func main() {
	//===================== DOGS
	dog1 := Dog{
		Name:        "Bounce",
		Breed:       "Pit Bull",
		AdoptionFee: 197.54,
	}
	dog2 := Dog{
		Name:        "Brave",
		Breed:       "caucasian",
		AdoptionFee: 281.91,
	}
	dogs := []Dog{dog1, dog2}

	//+++++++++++++++++++++ CATS
	cat1 := Cat{
		Name:        "Whiskers",
		Indoor:        false,
		AdoptionFee: 64.34,
	}
	cat2 := Cat{
		Name:        "Scream",
		Indoor:        true,
		AdoptionFee: 105.68,
	}
	cats := []Cat{cat2, cat1}

	// mixed := []Animal{dog1, cat1}
	fmt.Println("Cheapest dog: ", Cheapest(dogs))
	fmt.Println("Cheapest cat: ", Cheapest(cats))
	// fmt.Println(Cheapest(mixed))

}
