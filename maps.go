package main

import "fmt"

type Vertexs struct {
	Lat, long float64
}

var m map[string]Vertexs

func main_maps() {
	m = make(map[string]Vertexs)
	m["Bell Labs"] = Vertexs{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
