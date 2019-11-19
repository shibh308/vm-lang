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
	x * y + 12/(3 + 4) * 52155;
}
`
	s1 = `
	main(){
		x = 10;
		y = 30;
		x * y + 12/(3 + 4) * 52155;
	}
`
	tokens := lexString(strings.TrimSpace(s1))
	root := parseTokenSlice(tokens)
	root.generateOpCode()
}
