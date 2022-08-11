package main

func main(){
	x := 0
	for x < 5 {
		println(x)
		x++
	}

	for i:=0;i<10;i++ {
		println(i)
	}

	v := []int{100,200,300}
	for i,n := range v {
		println(i , " " , n)
	}
}
