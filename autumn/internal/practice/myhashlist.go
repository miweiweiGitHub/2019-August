package practice

import (
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
)

//雇员
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//链表
//这里的 链表不带表头，直接存放节点
type EmpLink struct {
	Head *Emp
}
//添加成员方法
func (this *EmpLink) Insert(emp *Emp) {
	//将成员加进来，并且成员编号从小到大
	cur := this.Head
	if cur==nil{
		this.Head = emp
		return
	}

	var temp *Emp = nil
	for {
		//循环找位置
		if cur != nil{
			if cur.Id >= emp.Id{
				//找到要添加的位置了
				break
			}
			//temp 位置在cur 的前面一个
			temp = cur
			cur = cur.Next
		}else {
			//说明当前添加的 id 最大
			break
		}

	}

	if temp ==nil{
		temp = emp
		this.Head = temp
	}else{
		temp.Next = emp
	}

	emp.Next = cur

}

func (this *EmpLink) ShowLink(i int) {
	if this.Head ==nil{
		log.Info("链表[%v] 为空",i)
	}

	emp := this.Head
	for  {
		if emp!=nil {
			fmt.Printf("链表[%v] id=%v,name=%v ->",i,emp.Id,emp.Name)
			emp = emp.Next
		}else {
			break
		}
	}
	fmt.Println()
}

func (this *EmpLink) FindById(i int) *Emp{

	cur := this.Head
	if cur ==nil{
		log.Error("当前的链表为空")
		return nil
	}

	for{
		if cur ==nil {
			log.Error("当前链表未找到对应的数据")
			return nil
		}
		if cur.Id == i{
			fmt.Println("找到对应 成员：",cur)
			return cur
		}
		cur = cur.Next
	}

}

func (this *EmpLink) DeleteById(i int) bool {
	cur := this.Head
	if cur ==nil{
		log.Error("当前的链表为空")
		return false
	}

	var temp *Emp = nil
	for{
		if cur == nil{
			log.Error("当前的链表未找到删除数据")
			return false
		}

		if cur.Id == i{
			fmt.Println("删除目标 成员：",cur)
			//第一个就匹配到
			if temp == nil {
				cur = cur.Next
				this.Head = cur
				return true
			}

			//其他情况，指针后移一位
			temp.Next = cur.Next
			return true

		}
		temp = cur
		cur = cur.Next
	}



}

//链表数组
type LinkTable struct {
	LinkArr [7]EmpLink
}

func (this *LinkTable) InsertEmp(e *Emp) {
	linkNo := this.HashFun(e.Id)
	this.LinkArr[linkNo].Insert(e)

}

func (this *LinkTable) HashFun(id int) int {
	return id % 7
}

func (this *LinkTable)ShowAll()  {
	for i:=0;i<len(this.LinkArr) ;i++  {
		this.LinkArr[i].ShowLink(i)

	}
}

func (this *LinkTable) FindById(i int)  *Emp{
	no := this.HashFun(i)
	return  this.LinkArr[no].FindById(i)
}

func (this *LinkTable) DeleteById(i int) bool {

	no := this.HashFun(i)
	return  this.LinkArr[no].DeleteById(i)
}

func testMain()  {

	key := ""
	id := 0
	name := ""
	var table LinkTable
lop1:
	for {
		fmt.Println("===================test====================")
		fmt.Println("input   添加成员")
		fmt.Println("show    展示成员")
		fmt.Println("find    查找成员")
		fmt.Println("delete  查找成员")
		fmt.Println("exit    退出")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入成员id")
			fmt.Scanln(&id)
			fmt.Println("请输入成员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			table.InsertEmp(emp)
		case "show":
			table.ShowAll()
		case "find":
			fmt.Println("请输入查找id")
			fmt.Scanln(&id)
			table.FindById(id)
		case "delete":
			fmt.Println("请输入删除id")
			fmt.Scanln(&id)
			table.DeleteById(id)
		case "exit":
			break lop1
		default:
			log.Error("请输入错误")
		}

	}
}
