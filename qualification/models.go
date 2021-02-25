package main

type Simulation struct {
	Time          int
	Intersections int
	Streets       int
	Cars          int
	Score         int
}

type Street struct {
	Name     string
	Start    int
	End      int
	Duration int
}

type CarPath struct {
	Streets int
	Path    []string
}
