package practice

import (
	"errors"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
)

//数组来实现queue
//非环形

//这个队列 一个数组有大小，size
// font 输出指针 初始化为 -1
//real 输入指针 初始化 -1

//AddQueue 加入队列
//GetQueue 取出队列
//ShowQueue 展示队列

type ArrQueue struct {
	Array   [3]int
	Maxsize int
	Font    int
	Rear    int
}

func (this *ArrQueue) AddQueue(val int) (err error) {

	//判断当前的队列是否已经装满
	if this.Rear == this.Maxsize-1 {
		return errors.New("arr queue full")
	}
	this.Rear++
	this.Array[this.Rear] = val
	return

}

func (this *ArrQueue) GetQueue() (val int, err error) {

	//判断当前的队列是为空
	if this.Rear == this.Font {
		return -1, errors.New("arr queue empty")
	}
	this.Font++
	return this.Array[this.Font], err

}

func (this *ArrQueue) ShowQueue() {

	for i := this.Font + 1; i <= this.Rear; i++ {
		log.Info("array[%d]=%d", i, this.Array[i])
	}
}

func (this *ArrQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	this.Array[this.Rear] = val
	this.Rear = (this.Rear+1)%this.Maxsize
	return
}

func (this *ArrQueue) Pop() (val int, err error) {

	if this.IsEmpty(){
		return -1, errors.New("queue is empty")
	}
	val = this.Array[this.Font]
	this.Font = (this.Font+1)%this.Maxsize
	return
}

func (this *ArrQueue) Show(){

	size := this.Size()
	if size == 0{
		log.Info("queue ie empty")
	}
	temp := this.Font
	for i := 0; i<size;i++  {
		log.Info("queue array[%d]=%d",temp,this.Array[temp])
		temp = (temp+1)%this.Maxsize
	}

}

func (this *ArrQueue) IsFull() bool {
	return (this.Rear+1)%this.Maxsize == this.Font
}

func (this *ArrQueue) IsEmpty() bool {
	return this.Font == this.Rear
}

func (this *ArrQueue) Size() int {
	return (this.Rear + this.Maxsize - this.Font) % this.Maxsize
}

func test()  {
	queue := ArrQueue{
		Maxsize:3,
		Rear:0,
		Font:0,
	}

	var key string
	var val int
loop1:
	for   {
		fmt.Println(" add  表示添加数据到队列")
		fmt.Println(" get  表示取出队列数据")
		fmt.Println(" show 表示打印队列")
		fmt.Println(" exit 表示退出")
		fmt.Scanln(&key)
		switch key {
		case"add":
			fmt.Println("请输入你要加入队列的数据：")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err!=nil{
				log.Error("add queue err:%d",err)
			}

		case"get":

			val,err := queue.Pop()
			if err!=nil{
				log.Error("get queue err:%d",err)
			}
			log.Info("get queue val:%d",val)

		case"show":
			queue.Show()
		case"exit":
			break loop1

		default:
			log.Error("输入错误！！！")
		}

	}

}