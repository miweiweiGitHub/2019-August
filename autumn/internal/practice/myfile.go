package practice

import (
	"bufio"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	"io"
	"io/ioutil"
	"os"
)

func Base(dir string) {
	file, err := os.Open(dir)
	if err != nil{
		log.Error("open file fail err:{%d}",err)
	}
	log.Info("file:{%d}",file)

	defer file.Close() //需要及时关闭file句柄，否则会有内存泄漏

	//带缓冲区的file 读取
	reader := bufio.NewReader(file)
	for{
		str, err2 := reader.ReadString('\n')

		if err2 == io.EOF{
			log.Info("reader file end")
			break
		}

		fmt.Print(str)
	}
}

func OnceReader(dir string)  {
	//一次性读取完
	content, err := ioutil.ReadFile(dir)

	if err!=nil{
		log.Error("reader file fail err:{%d}",err)
	}
	log.Info("content:{%v}",string(content))
}

func WriteBuf(dir string)  {
	file, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE, 0666)
	if err!=nil{
		log.Error("open file fail err:{%d}",err)
	}
	str := "hello world!"
	writer := bufio.NewWriter(file)

	for i:=0;i<5 ;i++  {
		writer.WriteString(str+"\n")
	}
	//将缓存数据真正写入到文件中
	writer.Flush()
}

//图片，音频，视频 文件拷贝
func OtherResource(souDir string, detDir string) (int64 ,error) {

	srcfile, err1 := os.Open(souDir)

	if err1!=nil{
		fmt.Printf("open file fail err=%v",err1)
	}
	defer srcfile.Close()
	reader := bufio.NewReader(srcfile)

	dstfile, err2 := os.OpenFile(detDir, os.O_CREATE|os.O_WRONLY, 0666)

	if err2!=nil{
		fmt.Printf("open file fail err=%v",err2)
	}
	defer dstfile.Close()
	writer := bufio.NewWriter(dstfile)

	return io.Copy(writer,reader)

}



