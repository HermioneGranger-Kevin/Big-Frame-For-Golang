# 第一个小项目的介绍

+ 小项目名称：模拟生成随机时间
+ 小项目源码位置：<https://github.com/HermioneGranger-Kevin/Big-Frame-For-Golang/blob/start/create_random_numbers_of_time.go>

### 第一、小项目的具体要求

+ 数据格式要求如下所示：
+ T1 开始时间取值范围为 [8:00 ～ 20:00) 之间的随机数
+ T2 开始时间取值范围为 [20:00 ～ 8:00) 之间的随机数
+ T1 和 T2 的结束时间为对应的开始时间 + 1小时 + 60分钟内的随机值（包括 60 分钟）
+ T1 和 T2 各 100 个数据点
+ 数据通过 HTTP 协议以 JSON 格式返回，接口地址定义为 GET /api
+ 只允许使用 Go 内置的包

### 第二、小知识存放
1. JSON数据解析方法：JSON数据：其实JSON数据就是一段字符串而已，只不过有不同意义的分隔符将其分割开来而已，我们看上面的符号，里面有[] ,{}等符号，其中：
	1. []中括号代表的是一个数组；
	1. {}大括号代表的是一个对象
	1. 双引号“”表示的是属性值
	1. 冒号：代表的是前后之间的关系，冒号前面是属性的名称，后面是属性的值，这个值可以是基本数据类型，也可以是引用数据类型。
	1. 一些介绍：https://www.cnblogs.com/abella/p/11125685.html
	1. 功能函数：func Marshal（v interface {}）（[] byte，error）：元帅返回v的JSON编码。
1. GenerateCalculatepoint：
	1. rand.Int：获取随机数，不加随机种子，每次遍历获取都是重复的一些随机数据
	1. rand.Seed(time.Now().UnixNano())：设置随机数种子，加上这行代码，可以保证每次随机（估计是少了：过程）都是随机的
	1. go中for的使用：https://studygolang.com/articles/4573
1.timeCalculate：
	1. time.Parse（解析）
	1. time.Duration（持续时间）
	1. time.Format：格式化为字符串
	1. t, _ := time.Parse("15:04", beginTime)：获取指定日期的时间戳：
		1. （时间戳是使用数字签名技术产生的数据，签名的对象包括了原始文件信息、签名参数、签名时间等信息。
		1.	时间戳系统用来产生和管理时间戳，对签名对象进行数字签名产生时间戳，以证明原始文件在签名时间之前已经存在。）
1. go语言的nil：https://blog.csdn.net/zf766045962/article/details/81005037


### 第三、参考资料：
+ Go学习指南：http://golang.org
+ Golang标准库API文档：http://studygolang.com/pkgdoc
+ 某博客：https://www.cnblogs.com/aresxin/p/go-yu-yantime-bao-de-shi-yong.html
