package main

import (
	"strings"
	"testing"
)

func TestGen(t *testing.T) {

	s1 := `
func(a, b, c){
	d = a + b;
	a + d + c;
}
main(){
	x = 10;
	y = 30;
	func(x, y, x + 10) * 3;
}
`
	testCode("s1", s1)

	s2 := `
main(){
	x = 10;
	y = 0;
	if(x == 10){
		y = 1;
	}
	y;
}
`
	testCode("s2", s2)
	s3 := `
fib(n){
	if(n == 1)
		return 1;
	if(n == 0)
		return 0;
	fib(n - 1) + fib(n - 2);
}
main(){
	fib(32);
}
`
	testCode("s3", s3)
}

func testCode(name string, code string) {
	tokens := lexString(strings.TrimSpace(code))
	root := parseTokenSlice(tokens)
	root.generateByteCode()
	root.printByteCode()
	filename := "./out_" + name + ".scbc"
	root.writeByteCode(filename)
}
