package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Token interface{}

type TokenOpenBr struct{}
type TokenCloseBr struct{}
type TokenOpenWBr struct{}
type TokenCloseWBr struct{}
type TokenPlus struct{}
type TokenMinus struct{}
type TokenDPlus struct{}
type TokenDMinus struct{}
type TokenSlash struct{}
type TokenSemiColon struct{}
type TokenColon struct{}
type TokenPer struct{}
type TokenAst struct{}
type TokenAnd struct{}
type TokenPlusAssign struct{}
type TokenMinusAssign struct{}
type TokenMulAssign struct{}
type TokenDivAssign struct{}
type TokenModAssign struct{}
type TokenReturn struct{}
type TokenBreak struct{}
type TokenContinue struct{}
type TokenIf struct{}
type TokenWhile struct{}
type TokenFor struct{}
type TokenEqual struct{}
type TokenDEqual struct{}
type TokenNEqual struct{}
type TokenEx struct{}
type TokenComma struct{}
type TokenGr struct{}
type TokenGrEqual struct{}
type TokenLe struct{}
type TokenLeEqual struct{}
type TokenTrue struct{}
type TokenFalse struct{}
type TokenIgnore struct{}
type TokenVar struct {
	name string
}
type TokenNum struct {
	num int
}
type TokenStr struct {
	str string
}
type TokenChar struct {
	char uint8
}

var regVar = regexp.MustCompile(`[a-zA-Z]\w*`)
var regNum = regexp.MustCompile(`\d+`)
var regStr = regexp.MustCompile(`"\w+"`)
var regChar = regexp.MustCompile(`'\w'`)
var regSpace = regexp.MustCompile(`\s+`)

func matchString(s *string, i *int) (Token, error) {
	switch {
	case strings.HasPrefix((*s)[*i:], `(`):
		*i += 1
		return TokenOpenBr{}, nil
	case strings.HasPrefix((*s)[*i:], `)`):
		*i += 1
		return TokenCloseBr{}, nil
	case strings.HasPrefix((*s)[*i:], `{`):
		*i += 1
		return TokenOpenWBr{}, nil
	case strings.HasPrefix((*s)[*i:], `}`):
		*i += 1
		return TokenCloseWBr{}, nil
	case strings.HasPrefix((*s)[*i:], `++`):
		*i += 2
		return TokenDPlus{}, nil
	case strings.HasPrefix((*s)[*i:], `--`):
		*i += 2
		return TokenDMinus{}, nil
	case strings.HasPrefix((*s)[*i:], `+=`):
		*i += 2
		return TokenPlusAssign{}, nil
	case strings.HasPrefix((*s)[*i:], `-=`):
		*i += 2
		return TokenMinusAssign{}, nil
	case strings.HasPrefix((*s)[*i:], `*=`):
		*i += 2
		return TokenMulAssign{}, nil
	case strings.HasPrefix((*s)[*i:], `/=`):
		*i += 2
		return TokenDivAssign{}, nil
	case strings.HasPrefix((*s)[*i:], `%=`):
		*i += 2
		return TokenModAssign{}, nil
	case strings.HasPrefix((*s)[*i:], `-`):
		*i += 1
		return TokenMinus{}, nil
	case strings.HasPrefix((*s)[*i:], `+`):
		*i += 1
		return TokenPlus{}, nil
	case strings.HasPrefix((*s)[*i:], `/`):
		*i += 1
		return TokenSlash{}, nil
	case strings.HasPrefix((*s)[*i:], `;`):
		*i += 1
		return TokenSemiColon{}, nil
	case strings.HasPrefix((*s)[*i:], `:`):
		*i += 1
		return TokenColon{}, nil
	case strings.HasPrefix((*s)[*i:], `%`):
		*i += 1
		return TokenPer{}, nil
	case strings.HasPrefix((*s)[*i:], `*`):
		*i += 1
		return TokenAst{}, nil
	case strings.HasPrefix((*s)[*i:], `&`):
		*i += 1
		return TokenAnd{}, nil
	case strings.HasPrefix((*s)[*i:], `true`):
		*i += 4
		return TokenTrue{}, nil
	case strings.HasPrefix((*s)[*i:], `false`):
		*i += 5
		return TokenFalse{}, nil
	case strings.HasPrefix((*s)[*i:], `return`):
		*i += 6
		return TokenReturn{}, nil
	case strings.HasPrefix((*s)[*i:], `break`):
		*i += 5
		return TokenBreak{}, nil
	case strings.HasPrefix((*s)[*i:], `continue`):
		*i += 8
		return TokenContinue{}, nil
	case strings.HasPrefix((*s)[*i:], `false`):
		*i += 5
		return TokenFalse{}, nil
	case strings.HasPrefix((*s)[*i:], `for`):
		*i += 3
		return TokenFor{}, nil
	case strings.HasPrefix((*s)[*i:], `if`):
		*i += 2
		return TokenIf{}, nil
	case strings.HasPrefix((*s)[*i:], `while`):
		*i += 5
		return TokenWhile{}, nil
	case strings.HasPrefix((*s)[*i:], `==`):
		*i += 2
		return TokenDEqual{}, nil
	case strings.HasPrefix((*s)[*i:], `!=`):
		*i += 2
		return TokenNEqual{}, nil
	case strings.HasPrefix((*s)[*i:], `=`):
		*i += 1
		return TokenEqual{}, nil
	case strings.HasPrefix((*s)[*i:], `>=`):
		*i += 2
		return TokenGrEqual{}, nil
	case strings.HasPrefix((*s)[*i:], `>`):
		*i += 1
		return TokenGr{}, nil
	case strings.HasPrefix((*s)[*i:], `<=`):
		*i += 2
		return TokenLeEqual{}, nil
	case strings.HasPrefix((*s)[*i:], `<`):
		*i += 1
		return TokenLe{}, nil
	case strings.HasPrefix((*s)[*i:], `!`):
		*i += 1
		return TokenEx{}, nil
	case strings.HasPrefix((*s)[*i:], `,`):
		*i += 1
		return TokenComma{}, nil
	}
	if sl := regVar.FindStringIndex((*s)[*i:]); sl != nil && sl[0] == 0 {
		name := (*s)[*i+sl[0] : *i+sl[1]]
		*i += sl[1]
		return TokenVar{name}, nil
	}
	if sl := regNum.FindStringIndex((*s)[*i:]); sl != nil && sl[0] == 0 {
		numStr := (*s)[*i+sl[0] : *i+sl[1]]
		*i += sl[1]
		num, _ := strconv.Atoi(numStr)
		return TokenNum{num}, nil
	}
	if sl := regStr.FindStringIndex((*s)[*i:]); sl != nil && sl[0] == 0 {
		*i += sl[1]
		str := (*s)[*i+sl[0]+1 : *i+sl[1]-1]
		return TokenStr{str}, nil
	}
	if sl := regChar.FindStringIndex((*s)[*i:]); sl != nil && sl[0] == 0 {
		*i += sl[1]
		ch := (*s)[*i+sl[0]+1]
		return TokenChar{ch}, nil
	}
	if sl := regSpace.FindStringIndex((*s)[*i:]); sl != nil && sl[0] == 0 {
		*i += sl[1]
		return TokenIgnore{}, nil
	}

	msg := `parse error at %d, line %d, "%s\n"`
	lineNum := strings.Count((*s)[:*i], "\n") + 1
	l := strings.LastIndex((*s)[:*i], "\n") + 1
	r := *i + strings.Index((*s)[*i:], "\n")
	return TokenIgnore{}, fmt.Errorf(msg, *i, lineNum, strings.TrimSpace((*s)[l:r]))
}

func lexString(s string) []Token {
	i := 0
	var tokens []Token
	for i < len(s) {
		token, err := matchString(&s, &i)
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		if reflect.TypeOf(token) != reflect.TypeOf(TokenIgnore{}) {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
