package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 4, 5, 7, 12}
	var s []int = primes[1:4] //定义切片, 注意切片中不需要指定元素个数,但数组需要指定
	fmt.Println(s)
	fmt.Println(&s[0])
	fmt.Println(&primes[1])

	var s1 []int
	s1 = s                       //两个切片相等
	fmt.Println(&s1[0] == &s[0]) //切片的地址是相等的

	nams := [4]string{"a", "b", "c", "d"}
	fmt.Println(nams)
	a := nams[0:2]
	b := nams[1:3]
	fmt.Println(a, b)

	b[0] = "XXXXX" //注意这里是将b的第0个元素修改为XXXX
	fmt.Println(a, b)
	fmt.Println(nams) //把原数组的值做改变
	c := a[1:2]
	c[0] = "YYY"
	fmt.Println(c)

	var p [2]*string
	p[0] = &nams[0]
	p[1] = &nams[1]
	*p[0] = "AAA"
	fmt.Println(p)
	fmt.Println(nams)

	ss := []int{2, 3, 5, 6, 8, 9}
	ss = ss[1:4]
	fmt.Println(ss) //输出[3, 5, 6]

	ss = ss[:2]     //对第一次切片继续切片
	fmt.Println(ss) //输出[5, 6]

	ss = ss[1:]     //对第二次切片继续切片
	fmt.Println(ss) //输出[5]

	aa := []int{2, 3, 5, 6, 8, 9}
	aa = aa[:0]
	fmt.Println(aa, len(aa), cap(aa)) //cap打印切片的容量,潜在的容量

	aa = aa[:4]
	fmt.Println(aa, len(aa), cap(aa))

	aa = aa[2:]
	fmt.Println(aa, len(aa), cap(aa))

	//make创建切片
	ff := make([]int, 5) //长度为5
	fmt.Println(ff)
	ff[0] = 1
	fmt.Println(ff)

	bb := make([]int, 0, 5) //长度为0,容量为5
	fmt.Println("bb:", bb, len(bb), cap(bb))
	bb1 := bb[:2]
	fmt.Println("bb1:", bb1, len(bb1), cap(bb1))
	bb2 := bb1[2:5]
	fmt.Println("bb2:", bb2, len(bb2), cap(bb2))

	var dd []int //定义一个空切片 dd
	fmt.Println(dd)

	dd = append(dd, 0)
	fmt.Println(dd)
	dd = append(dd, 1)
	fmt.Println(dd)
	dd = append(dd, 2, 3, 4, 5)
	fmt.Println(dd)
	ee := append(dd, 6)
	fmt.Println(ee)
	ee[0] = 111
	fmt.Println(ee, dd)
	//两个切片相加
	ee = append(ee, dd...) //ee与dd两个切片相加,注意后面必须的三个点
	fmt.Println(ee)

	//反正切片
	abc := []int{2, 3, 5, 7, 11, 13, 16}
	fmt.Println(abc)
	for i := 0; i < len(abc)/2; i++ {
		abc[i], abc[len(abc)-i-1] = abc[len(abc)-i-1], abc[i]
	}
	fmt.Println(abc)

}

//切片可以理解为对数组的一个引用
