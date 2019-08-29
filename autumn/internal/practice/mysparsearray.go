package practice

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

//稀疏数组 sparse array 压缩数组，将有效数据保存，重复的数据可以记录成一条
//案例 五子棋 存盘退出，续上盘 等功能实现
//记录 实际有效的数据
type NodeVal struct {
	Row int         `json:"row"`
	Col int         `json:"col"`
	Val interface{} `json:"val"`
}

//var sizeOfMyStruct = int(unsafe.Sizeof(NodeVal{}))

//func MyStructToBytes(s *NodeVal) []byte {
//	var x reflect.SliceHeader
//	x.Len = sizeOfMyStruct
//	x.Cap = sizeOfMyStruct
//	x.Data = uintptr(unsafe.Pointer(s))
//	return *(*[]byte)(unsafe.Pointer(&x))
//}

func Method01(dir string) {

	//原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//输出原始数组
	for _, v1 := range chessMap {

		for _, v2 := range v1 {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}

	//转换成稀疏数组
	//1)遍历原始数组，发现不为 0 ，保存一个node 结构体
	//2) 将其放到对应的切片
	var sparseArr []NodeVal

	var node1 = NodeVal{
		Row: 11,
		Col: 11,
		Val: 0,
	}
	sparseArr = append(sparseArr, node1)
	for i, v1 := range chessMap {

		for j, v2 := range v1 {
			if v2 != 0 {

				var node = NodeVal{
					Row: i,
					Col: j,
					Val: v2,
				}
				sparseArr = append(sparseArr, node)
			}
		}
	}

	//for k,v := range sparseArr {
	//	fmt.Printf("index=%d,val=%d \n",k,v)
	//}

	//将文件写到磁盘，再从磁盘读取出来，恢复成原始数组展示
	file, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("open file fail %d \t", err)
	}

	writer := bufio.NewWriter(file)
	for k, v := range sparseArr {
		fmt.Printf("index=%d,val=%d \n", k, v)
		con, _ := json.Marshal(v)
		fmt.Printf("%v \n", string(con))

		le, err2 := writer.WriteString(string(con) + "\n")

		if err2 != nil {
			fmt.Printf("write file fail %d \t", err2)
		}
		fmt.Printf("content:%v \n", le)
	}

	writer.Flush()

}

//将文件读取回来，恢复成原始的数组
func ReView(dir string) {

	dirfile, err := os.Open(dir)
	if err != nil {
		fmt.Printf("open file fail %d \t", err)
	}
	defer dirfile.Close()
	reader := bufio.NewReader(dirfile)

	var spaceArr []NodeVal
	for {
		line, err2 := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err2 == io.EOF {
			break
		}

		fmt.Printf("reader line = %v \n", line)
		val := NodeVal{}
		err3 := json.Unmarshal([]byte(line), &val)
		if err3 != nil {
			fmt.Printf("json transfer fail %d \t", err3)
		}
		spaceArr = append(spaceArr, val)
	}

	var arr [11][11]interface{}
	for k,v := range spaceArr {
		if k != 0{
			arr[v.Row][v.Col] = v.Val
		}
	}


	for _,j := range arr {
		for _,v := range j {
			if v==nil {
				v=0
			}
			fmt.Printf(" %v ", v)
		}

		fmt.Println()
	}

}
