package main

type Utils interface {
	Greet()
	Calc(a int, b int) int
}

type Random interface {
	GenerateRandomInt() int
}

type MyString string

func (ms MyString) Error() string {
	return string(ms)
}

type Person struct{}

func (ms MyString) Greet() {}

func (ms MyString) Calc(c int, d int) int {
	return c
}

func (ms MyString) GenerateRandomInt() int {
	return 0
}

func main() {

	var ms MyString = "Johan"

	var u Utils = ms

	var r Random = ms

	var name MyString = r.(MyString)

	_, _, _ = u, r, name

	var e error = MyString("ada error")

	_ = e
}
