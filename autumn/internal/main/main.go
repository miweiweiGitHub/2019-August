package main

import "autumn/internal/practice"

func main() {

	head := &practice.CityNode{}
	node1 := &practice.CityNode{
		No:1,
		Name:"北京",
	}
	node2 := &practice.CityNode{
		No:2,
		Name:"上海",
	}
	node3 := &practice.CityNode{
		No:3,
		Name:"天津",
	}

	practice.AddCityNode(head,node1)
	practice.AddCityNode(head,node2)
	practice.AddCityNode(head,node3)
	practice.ShowCityList(head)

}
