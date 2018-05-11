package glog

import (
	"fmt"
	"log"
	"os"
	"time"
)

//初始化日志
func  Init(fileName string) {
	fileName = fmt.Sprintf(fileName+"-%s.log", time.Now().Format("20060102"))

	f,err:=os.Create(fileName)

	//outfile, err := os.OpenFile(n.Name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开
	if err != nil {
		fmt.Println(*f, "open failed")
		os.Exit(1)
	}

	log.SetOutput(f)  //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	defer func() {
		fmt.Println("log init finish! filename:"+fileName)
		//outfile.Close()

	}()

}
