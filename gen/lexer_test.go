package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	s1 := `
main(){
	x = 10;
	y = 30;
	str_sample = "test";
	char_sample = 'c';
	x * y + 12/(3 + 4) * 52155;
}
`
	tokens := lexString(strings.TrimSpace(s1))
	for i, token := range tokens {
		fmt.Printf("%03d: %s\n", i+1, reflect.TypeOf(token))
	}
}
