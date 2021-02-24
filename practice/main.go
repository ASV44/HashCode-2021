package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// a
	// b
	// c
	// d
	// e
	inputName := os.Args[1]
	order, pizzas, err := readFromFile("./input_data/" + inputName)
	if err != nil {
		fmt.Println("Parsing input err ", err)
		return
	}

	//fmt.Println(order)
	//for _, pizza := range pizzas {
	//	fmt.Println(pizza)
	//}

	sort.Sort(ByIngredientsAmount(pizzas))

	var delivery [][]Pizza
	for i := 0; i < order.TwoPersonsTeams; i++ {
		deliveryPizzas, remainingPizzas, err := getDeliveryPizza(pizzas, 2)
		if err != nil {
			break
		}
		pizzas = remainingPizzas
		delivery = append(delivery, deliveryPizzas)
		//for _, deliveryPizza := range deliveryPizzas {
		//	fmt.Print(deliveryPizza.Index, " ")
		//}
		//fmt.Println()
	}

	for i := 0; i < order.ThreePersonsTeams; i++ {
		deliveryPizzas, remainingPizzas, err := getDeliveryPizza(pizzas, 3)
		if err != nil {
			break
		}
		pizzas = remainingPizzas
		delivery = append(delivery, deliveryPizzas)
		//fmt.Print("3", " ")
		//for _, deliveryPizza := range deliveryPizzas {
		//	fmt.Print(deliveryPizza.Index, " ")
		//}
		//fmt.Println()
	}

	for i := 0; i < order.FourPersonsTeams; i++ {
		deliveryPizzas, remainingPizzas, err := getDeliveryPizza(pizzas, 4)
		if err != nil {
			break
		}
		pizzas = remainingPizzas
		delivery = append(delivery, deliveryPizzas)
		//fmt.Print("4", " ")
		//for _, deliveryPizza := range deliveryPizzas {
		//	fmt.Print(deliveryPizza.Index, " ")
		//}
		//fmt.Println()
	}

	fmt.Println(len(delivery))
	for _, d := range delivery {
		fmt.Print(len(d), " ")
		for _, pizzaDelivery := range d {
			fmt.Print(pizzaDelivery.Index, " ")
		}
		fmt.Println()
	}
}

func readFromFile(fileName string) (Order, []Pizza, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	orderData := strings.Fields(scanner.Text())
	order, err := parseOrderData(orderData)
	if err != nil {
		return order, []Pizza{}, err
	}

	var pizzas []Pizza
	pizzaIndex := 0
	for scanner.Scan() {
		pizza, err := parsePizzaInput(pizzaIndex, scanner.Text())
		if err != nil {
			return Order{}, nil, err
		}

		pizzas = append(pizzas, pizza)
		pizzaIndex += 1
	}

	return order, pizzas, nil
}

func parseOrderData(orderData []string) (Order, error) {
	order := Order{}
	for index, data := range orderData {
		dataValue, err := strconv.Atoi(data)
		if err != nil {
			return Order{}, err
		}
		switch index {
		case 0:
			order.MaxPizzas = dataValue
		case 1:
			order.TwoPersonsTeams = dataValue
		case 2:
			order.ThreePersonsTeams = dataValue
		case 3:
			order.FourPersonsTeams = dataValue
		}
	}

	return order, nil
}

func parsePizzaInput(pizzaIndex int, pizzaData string) (Pizza, error) {
	data := strings.Fields(pizzaData)
	ingredientsAmountValue, ingredients := data[0], data[1:]
	ingredientsAmount, err := strconv.Atoi(ingredientsAmountValue)
	if err != nil {
		return Pizza{}, err
	}

	return Pizza{Index: pizzaIndex, IngredientsAmount: ingredientsAmount, Ingredients: ingredients}, nil
}

type ByIngredientsAmount []Pizza

func (pizza ByIngredientsAmount) Len() int      { return len(pizza) }
func (pizza ByIngredientsAmount) Swap(i, j int) { pizza[i], pizza[j] = pizza[j], pizza[i] }
func (pizza ByIngredientsAmount) Less(i, j int) bool {
	return pizza[i].IngredientsAmount < pizza[j].IngredientsAmount
}

func getDeliveryPizza(pizzas []Pizza, teamSize int) ([]Pizza, []Pizza, error) {
	if len(pizzas) < teamSize {
		return nil, nil, errors.New("less pizzas than team members")
	}

	//var deliveryPizzas []Pizza
	//best := 0
	//worst := 0
	//for i := 0; i < teamSize; i++ {
	//	deliveryPizzas = append(deliveryPizzas, pizzas[i])
	//	best += pizzas[i].IngredientsAmount
	//	if worst < pizzas[i].IngredientsAmount {
	//		worst = pizzas[i].IngredientsAmount
	//	}
	//}
	//
	//for i := teamSize; i < len(pizzas); i++ {
	//
	//}

	deliveryPizzas := pizzas[0:teamSize]
	remainingPizzas := pizzas[teamSize:]

	return deliveryPizzas, remainingPizzas, nil
}
