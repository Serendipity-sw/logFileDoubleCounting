package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
	"io/ioutil"
)

var (
	fileArray []string
	phoneCount map[string]int=make(map[string]int)
	urlMatch map[string]int=make(map[string]int)
	urlPhone map[string]name=map[string]name{}
	urlPathArray []string
	urlNameArray map[string]string=make(map[string]string)
)

type name  struct{
	list map[string]int
}

func main() {

	readFilde("./123123123.txt")
var asdfasdfsa []string
	for _,value:=range urlPathArray  {
		asdfasdfsa=strings.Split(value,"	")
		urlNameArray[asdfasdfsa[0]]=asdfasdfsa[1]
	}


	files, _ := ioutil.ReadDir("./files")
	for _,file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println(file.Name())
			readFile("./files/"+file.Name())
		}
	}





	var ok bool



	var phone []string
	//urlPhone map[string]name=map[string]name{}
	for _, value:=range fileArray  {
		phone=strings.Split(value,"	")
		phoneCount[phone[0]]=0
		_, ok = urlMatch[phone[2]]
		if !ok {
			urlMatch[phone[2]]=1
		}else{
			urlMatch[phone[2]]++
		}
		var list map[string]int =make(map[string]int)
		list[phone[0]]=0
		_,ok=urlPhone[phone[2]]
		if !ok {
			urlPhone[phone[2]]=name{list:list}
		}else{
			urlPhone[phone[2]].list[phone[0]]=0
		}
	}
	
	fmt.Println(len(fileArray))
	
	fmt.Println(len(phoneCount))

	fmt.Printf("%v \n",urlMatch)

	fmt.Println(len(urlMatch))

	phoneStr:=""
	phoneStr+=fmt.Sprintf("%s,%d,%d \r\n","0",len(fileArray),len(phoneCount))
	for index,value:=range urlPhone  {
		fmt.Printf("%s:%d ",index,len(value.list))
		phoneStr+=fmt.Sprintf("%s,%d,%d,%s \r\n",index,urlMatch[index],len(value.list),urlNameArray[index])
	}

	write,err:=os.OpenFile("./process.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("文件打开出错! err: %s \n",err.Error())
		return
	}
	write.WriteString(fmt.Sprintf("%d \n %d \n %v \n %d \n %s \n",len(fileArray),len(phoneCount),urlMatch,len(urlMatch),phoneStr))
	write.Close()
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

func readFilde(filePath string) {
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
			urlPathArray=append(urlPathArray,string(line))
			line=line[:0]
		}
	}
}

