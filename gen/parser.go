package main

import (
	"fmt"
	"os"
	"reflect"
)

func (p RootNode) parseTokens(tokens []Token) error {
	tokensLen := len(tokens)
	for i := 0; i < 10; i++ {
		tokens = append(tokens, TokenIgnore{})
	}
	pr := new(ProgNode)
	if i, err := pr.parseTokens(&tokens, 0); err || (i != tokensLen) {
		if err {
			msg := "parse Error at Token %d, %s\n"
			return fmt.Errorf(msg, i+1, reflect.TypeOf(tokens[i]))
		}
	}
	return nil
}

func (p ProgNode) parseTokens(tokens *[]Token, i int) (int, bool) {
	var fdefNode = new(FdefNode)
	if _i, err := fdefNode.parseTokens(tokens, i); err {
		return _i, true
	} else {
		i = _i
		var node PNode = fdefNode
		p.childs = append(p.childs, node)
	}

	for i < len(*tokens) {
		fdefNode = new(FdefNode)
		if _i, err := fdefNode.parseTokens(tokens, i); err {
			return _i, false
		} else {
			i = _i
			var node PNode = fdefNode
			p.childs = append(p.childs, node)
		}
	}
	return i, false
}

func (p FdefNode) parseTokens(tokens *[]Token, i int) (int, bool) {
	switch v := (*tokens)[i].(type) {
	case TokenVar:
		p.name = v.name
		i++
	default:
		return i, true
	}

	var varsNode = new(VarsNode)
	if _i, err := varsNode.parseTokens(tokens, i); err {
		fmt.Println(_i)
		return _i, true
	} else {
		i = _i
		p.vars = varsNode
	}

	var blockNode = new(BlockNode)
	if _i, err := blockNode.parseTokens(tokens, i); err {
		return _i, true
	} else {
		i = _i
		p.content = blockNode
	}

	return i, false
}

func (p VarsNode) parseTokens(tokens *[]Token, i int) (int, bool) {
	switch (*tokens)[i].(type) {
	case TokenOpenBr:
		i++
	default:
		return i, true
	}

	firstArg := true
	for {
		if !firstArg {
			flag := false
			switch (*tokens)[i].(type) {
			case TokenComma:
				i++
				flag = true
			}
			if !flag {
				break
			}
		} else {
			firstArg = false
		}

		var varNode = new(VarNode)
		if _i, err := varNode.parseTokens(tokens, i); err {
			break
		} else {
			i = _i
			p.args = append(p.args, varNode)
		}
	}

	switch (*tokens)[i].(type) {
	case TokenCloseBr:
		i++
	default:
		return i, true
	}
	return i, false
}

func (p VarNode) parseTokens(tokens *[]Token, i int) (int, bool) {
	varNode := new(VarNode)
	switch v := (*tokens)[i].(type) {
	case TokenVar:
		i++
		varNode.name = v.name
	default:
		return i, true
	}
	switch (*tokens)[i].(type) {
	case TokenAnd:
		i++
		varNode.isRef = true
	}
	return i, false
}

func (p BlockNode) parseTokens(tokens *[]Token, i int) (int, bool) {
	return i, true
}

func parseTokenSlice(tokens []Token) {
	root := RootNode{}
	if err := root.parseTokens(tokens); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
