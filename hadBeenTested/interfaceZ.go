package main

type tester interface {
	test()
	string() string
}

type data struct{}

func (*data) test()         {}
func (data) string() string { return "a" }

func main() {
	var d data
	var t tester = &d
	t.test()
	println(t.string())
}
