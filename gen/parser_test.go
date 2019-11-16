package main

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	s1 := `
main(){
	x = 10;
	y = 30;
	x * y + 12/(3 + 4) * 52155;
}
`
	tokens := lexString(strings.TrimSpace(s1))
	parseTokenSlice(tokens)
}
