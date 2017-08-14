2017年6月3日  开始学习golang语言
以下为学习过程整理的相关笔记

go特点：
    静态编译
    垃圾回收
    简洁的符号和语法
    平坦的类型系统
    基于CSP的并发模型
    高效简单的工具链
    丰富的标准库

go语言的应用：
    Docker，火热的容器化技术
    Kubernets，Goole Borg的开源实现
    Etcd，类似Zookeeper的高可用key-value存储

运行和编译
    go run hellO.go
    go build hello.go
    ./hello
    file hello #可以使用file命令查看该文件隶属于的系统

    linux下编译：
    编译成linux版本：GOOS=linux go build hello.go   生成hello
    编译成windows版本：GOOS=windows go build hello.go   生成hello.exe
    编译成mac版本：GOOS=darwin go build hello.go   生成hello.mac

    windows下编译
    set GOOS=windows

缩进问题
    gofmt -w hello.go   自动缩进
    goimports -w hello.go  自动引入包

一些常用命令
    go build
    go test
    go get
    godoc -http=:9000

go中的指针

go中的作用域

go的堆和栈

go的数据类型
    整数
        int,int32,int64,uint,uint32,uint64
        int
        int8    [-128 , 127]
        int16   [-32768 , 32767]
        int32   [-2**16, 2**16 -1]
        int64   [-2**32, 2**32 -1]
        uint
        uint8   [0, 2**8 - 1]
        uint16  [0, 2**16 - 1]
        uint32  [0, 2**32]
        uint64  [0, 2**64]
    字符串
        字符串本身不可修改
        通过跟[]byte来相互转换来实现修改
    布尔
    浮点
    常量 const 
    
go的条件语句
    if
    switch
    for
    
复合数组类型
    数组, 切片
    map
    结构体
    序列化和反序列化
    
map
    hash方式的
    无序的
    o(1)的访问时间   //时间复杂度
    
结构体
    type 名字 struct {
        Id  int
        Name string
    
}


序列化和反序列化
    序列化    把内存的数据存到磁盘          json.Marshal
    反序列化  把磁盘的数据读取后加载到内存  json.Unmarshal

过程式编程
    函数
    错误异常处理
    go调用外部程序
    文件读取的多种方式
    
匿名函数

读取文件的几种方式
    file.Read
    ioutil.ReadFile
    bufio.Scanner
    bufio.Reader
    io.Copy
    
type是用来定义类的
    package main
    type Path []Point
    func (path Path) Distance() float64{

}

为什么需要接口


go的总揽
    携程和channel
    并发爬虫
    多线程下载器
    携程
    - 类似线程的处理
    - 并发
    
    启动携程
    - go关键字加函数
    - 
    
    