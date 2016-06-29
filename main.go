package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

var (
	fileArray []string
	phoneCount map[string]int=make(map[string]int)
	urlMatch map[string]int=make(map[string]int)
)

func main() {

	var ok bool

	readFile("./jizhihui.log.2016-06-28-00")

	var phone []string

	for _, value:=range fileArray  {
		phone=strings.Split(value,"	")
		phoneCount[phone[0]]=0
		_, ok = urlMatch[phone[2]]
		if !ok {
			urlMatch[phone[2]]=1
			continue
		}
		urlMatch[phone[2]]++
	}
	fmt.Println(len(phoneCount))

	fmt.Printf("%v",urlMatch)
}

/**
逐行读取文件
创建人:邵炜
创建时间:2016年6月29日15:41:51
输入参数:文件路劲
 */
func readFile(filePath string) {
	var(
		readAll =false
		readByte []byte
		line []byte
	)
	read,err:=os.Open(filePath)
	if err != nil {
		fmt.Printf("文件打开失败! err: %s \n",err.Error())
		return
	}
	defer read.Close()
	buf:=bufio.NewReader(read)
	for err!=io.EOF {
		if err != nil {
			fmt.Printf("文件读取错误! err: %s \n",err.Error())
		}
		if readAll {
			readByte,readAll,err=buf.ReadLine()
			line=append(line,readByte...)
		}else{
			readByte,readAll,err=buf.ReadLine()
			line=append(line,readByte...)
			if len(strings.TrimSpace(string(line)))==0 {
				continue
			}
			fileArray=append(fileArray,string(line))
			line=line[:0]
		}
	}
}