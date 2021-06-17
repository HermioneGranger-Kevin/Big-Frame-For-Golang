package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	_ "fmt"
)

//构建模拟时间数据点结构体
type Data struct {
	T1 [100]Datapoint
	T2 [100]Datapoint
}

//构建一个时间数据点结构体，这个需要用json格式
type Datapoint struct {
	B string `json:"begintime"`
	E string `json:"endtime"`
}

//主函数
func main() {
	http.HandleFunc("/api" , handlerData)
	//错误处理
	err := http.ListenAndServe(":8080" , nil)
	//调用函数，向http发起请求，url: localhost:8080/api
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//取浏览器器上输入：http://localhost:8080/api，可以看的见效果
	// fmt.Println(traverseCalculatePoint())//输出测试代码，用于输出遍历的100个数据点，查看效果 
}

//定义 timeCalculate 函数，生成一个数据点：调用一次，返回一个数据点
func calculateTime(begintime string) *Datapoint {

	datapoint := new(Datapoint)

	//生成随机的begintime
	bt, _:=time.Parse("15:04",begintime)//获取指定日期的时间戳
	//func (t Time) Add(d Duration) Time：时间加时间间隔，示例：计算1小时后及1小时前时间
	bt = bt.Add(time.Duration(rand.Intn(60)) * time.Minute)
	bt = bt.Add(time.Duration(rand.Intn(12)) * time.Hour)
	//生成的开始时间格式化为字符串
	datapoint.B = bt.Format("15:04")

	//计算相应的endtime
	bt = bt.Add(time.Hour)//先加上一小时再说
	bt = bt.Add(time.Duration(rand.Intn(61)) * time.Minute)//加上 [0,60] 分钟，0~60分钟之间有61个数字
	datapoint.E = bt.Format("15:04")//生成的结束时间格式化为字符串
	return datapoint//返回数据点
}

//定义 traverseCalculatePoint 函数，遍历100个时间数据点
func traverseCalculatePoint() Data {
	var data = Data{}
	//使用 range 遍历 T1，T2 数组并计算
	for k,_ := range data.T1{
		data.T1[k] = *calculateTime("08:00")
		data.T2[k] = *calculateTime("20:00")
	}
	return data
}

//定义 handlerData 函数，完成HTTP协议下数据格式JSON化
func handlerData(writer http.ResponseWriter, request *http.Request) {
	var data = Data{}
	//在get方法下向http发起请求
	if request.Method == "GET" {
		//遍历好的100个数据点保存在data中
		data = traverseCalculatePoint()
		//以 json 格式导出 data 数据，注意细节：,_ := json
		data_json,_ := json.Marshal(data)
		//设置 content-type ，让客户端明晰数据格式为json，注意首字母大写：H,S,W
		writer.Header().Set("Content-type","application/json")
		writer.Write(data_json)
	}
}
