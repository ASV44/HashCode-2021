package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputName := os.Args[1]
	simulation, streets, carsPaths, err := readFromFile("./input_data/" + inputName + ".txt")
	if err != nil {
		fmt.Println("Parsing input err ", err)
		return
	}

	//fmt.Println(simulation, streets)

	streetCarsCapacity := make(map[string]int)
	for _, carPath := range carsPaths {
		for _, street := range carPath.Path {
			streetCarsCapacity[street] += 1
		}
	}

	//fmt.Println(streetCarsCapacity)

	intersections := make(map[int][]string)
	for _, street := range streets {
		intersections[street.End] = append(intersections[street.End], street.Name)
	}

	//sortedIntersections := sortByStreets(intersections)
	//fmt.Println(sortedIntersections)

	var result []Intersection
	for key, _ := range intersections {
		if simulation.Time <= 0 {
			break
		}
		streetsIntersections := make(map[string]int)
		for _, street := range intersections[key] {
			if simulation.Time <= 0 {
				break
			}
			time := 0
			carAmount := streetCarsCapacity[street]
			if carAmount == 0 {
				continue
			}
			if carAmount > 7 {
				time = 2
			} else {
				time = 1
			}
			streetsIntersections[street] = time
			simulation.Time -= time
		}
		if len(streetsIntersections) > 0 {
			streetIntersection := Intersection{
				Index:   key,
				Streets: streetsIntersections,
			}

			result = append(result, streetIntersection)
		}
	}

	fmt.Println(len(result))
	for _, intersection := range result {
		fmt.Println(intersection.Index)
		fmt.Println(len(intersection.Streets))
		for name, time := range intersection.Streets {
			fmt.Println(name, " ", time)
		}
	}
}

func sortByStreets(intersections map[int][]string) PairList {
	pl := make(PairList, len(intersections))
	i := 0
	for k, v := range intersections {
		pl[i] = Pair{k, len(v)}
		i++
	}

	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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
