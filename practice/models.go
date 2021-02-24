package main

type Order struct {
	MaxPizzas         int
	TwoPersonsTeams   int
	ThreePersonsTeams int
	FourPersonsTeams  int
}

type Pizza struct {
	Index             int
	IngredientsAmount int
	Ingredients       []string
}
