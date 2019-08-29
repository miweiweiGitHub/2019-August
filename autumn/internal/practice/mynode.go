package practice

import "github.com/bilibili/kratos/pkg/log"

type CityNode struct {
	No       int
	Name     string
	Next     *CityNode
}

func  AddCityNode(head *CityNode, element *CityNode) {

	temp := head
	for{
		if temp.Next==nil{
			temp.Next = element
			break
		}
		temp = temp.Next
	}

}
//顺序添加，根据node的节点
func  InsertCityNode(head *CityNode, element *CityNode) {

	temp := head
	if temp == nil{

	}


	for{
		if temp.Next==nil{
			temp.Next = element
			break
		}
		temp = temp.Next
	}

}



func  ShowCityList(head *CityNode){

	temp := head
	for{
		log.Info("no=%d,name=%d",temp.No,temp.Name)
		if temp.Next == nil{
			break
		}
		temp = temp.Next
	}

}

