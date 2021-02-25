package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputName := os.Args[1]
	simulation, streets, carsPaths, err := readFromFile("./" + inputName + ".txt")
	if err != nil {
		fmt.Println("Parsing input err ", err)
		return
	}

	fmt.Println(simulation, streets, carsPaths)
}

func readFromFile(fileName string) (Simulation, []Street, []CarPath, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	simulation := Simulation{}
	simulationData := strings.Fields(scanner.Text())
	for dataIndex, data := range simulationData {
		switch dataIndex {
		case 0:
			simulation.Time, err = strconv.Atoi(data)
		case 1:
			simulation.Intersections, err = strconv.Atoi(data)
		case 2:
			simulation.Streets, err = strconv.Atoi(data)
		case 3:
			simulation.Cars, err = strconv.Atoi(data)
		case 4:
			simulation.Score, err = strconv.Atoi(data)
		}
		if err != nil {
			return simulation, nil, nil, err
		}
	}

	var streets []Street
	for i := 0; i < simulation.Streets; i++ {
		scanner.Scan()
		streetData := strings.Fields(scanner.Text())
		start, err := strconv.Atoi(streetData[0])
		end, err := strconv.Atoi(streetData[1])
		duration, err := strconv.Atoi(streetData[3])
		street := Street{
			Name:     streetData[2],
			Start:    start,
			End:      end,
			Duration: duration,
		}
		if err != nil {
			return simulation, nil, nil, err
		}

		streets = append(streets, street)
	}

	var carsPaths []CarPath
	for i := 0; i < simulation.Cars; i++ {
		scanner.Scan()
		carData := strings.Fields(scanner.Text())
		streetsNumber, err := strconv.Atoi(carData[0])
		if err != nil {
			return simulation, nil, nil, err
		}

		carPath := CarPath{Streets: streetsNumber}
		carData = carData[1:]
		var carPathStreets []string
		for _, streetName := range carData {
			carPathStreets = append(carPathStreets, streetName)
		}

		carPath.Path = carPathStreets

		carsPaths = append(carsPaths, carPath)
	}

	return simulation, streets, carsPaths, nil
}
