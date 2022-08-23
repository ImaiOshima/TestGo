package main

func test(p **int) {
	x := 100
	//使用了
	*p = &x
}
func main() {
	var p *int
	//因为p是指针 所以&p就是双指针 代表的是p指针的地址
	test(&p)
	print(*p)
}
