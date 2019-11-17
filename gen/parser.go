package main

import (
	"fmt"
	"os"
	"reflect"
)

func makeSyntaxTree(tokens []Token) (*RootNode, error) {
	p := new(RootNode)
	tokensLen := len(tokens)
	for i := 0; i < ignoreSize; i++ {
		tokens = append(tokens, TokenIgnore{})
	}
	node, _i := makeProgNode(&tokens, 0)
	if node == nil || (_i != tokensLen) {
		msg := "parse error at Token %d, %s\n"
		return nil, fmt.Errorf(msg, _i+1, reflect.TypeOf(tokens[_i]))
	}
	p.prog = node
	return p, nil
}

func makeProgNode(tokens *[]Token, i int) (*ProgNode, int) {
	p := new(ProgNode)
	if node, _i := makeFdefNode(tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}

	for i < len(*tokens)-ignoreSize {
		if node, _i := makeFdefNode(tokens, i); node == nil {
			return p, i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
	return p, i
}

func makeFdefNode(tokens *[]Token, i int) (*FdefNode, int) {
	p := new(FdefNode)
	switch v := (*tokens)[i].(type) {
	case TokenVar:
		p.name = v.name
		i++
	default:
		return nil, i
	}

	if node, _i := makeVarsNode(tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.vars = node
	}

	if node, _i := makeBlockNode(tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.content = node
	}

	return p, i
}

func makeVarsNode(tokens *[]Token, i int) (*VarsNode, int) {
	p := new(VarsNode)
	switch (*tokens)[i].(type) {
	case TokenOpenBr:
		i++
	default:
		return nil, i
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

		if node, _i := makeVarNode(tokens, i); node == nil {
			break
		} else {
			i = _i
			p.args = append(p.args, node)
		}
	}

	switch (*tokens)[i].(type) {
	case TokenCloseBr:
		i++
	default:
		return nil, i
	}
	return p, i
}

func makeVarNode(tokens *[]Token, i int) (*VarNode, int) {
	p := new(VarNode)
	switch v := (*tokens)[i].(type) {
	case TokenVar:
		i++
		p.name = v.name
	default:
		return nil, i
	}
	switch (*tokens)[i].(type) {
	case TokenAnd:
		i++
		p.isRef = true
	}
	return p, i
}

func makeBlockNode(tokens *[]Token, i int) (*BlockNode, int) {
	p := new(BlockNode)
	switch (*tokens)[i].(type) {
	case TokenOpenWBr:
		i++
		for {
			if node, _i := makeStmthNode(tokens, i); node == nil {
				break
			} else {
				i = _i
				p.stmts = append(p.stmts, node)
			}
		}
		if len(p.stmts) == 0 {
			return nil, i
		}
		switch (*tokens)[i].(type) {
		case TokenCloseWBr:
			i++
		default:
			return nil, i
		}
	default:
		if node, _i := makeStmthNode(tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.stmts = []*StmthNode{node}
		}
	}
	return p, i
}

func makeStmthNode(tokens *[]Token, i int) (*StmthNode, int) {
	p := new(StmthNode)
	if node, _i := makeStmtNode(tokens, i); node != nil {
		i = _i
		p.stmt = node
		p.flag = flagSingleStmth
		switch (*tokens)[i].(type) {
		case TokenSemiColon:
			i++
			return p, i
		default:
			return nil, i
		}
	}
	/* TODO: for, while */
	return nil, i
}

func makeStmtNode(tokens *[]Token, i int) (*StmtNode, int) {
	p := new(StmtNode)
	p.flag = flagSingleStmt
	switch (*tokens)[i].(type) {
	case TokenReturn:
		i++
		p.flag = flagReturn
	}
	if node, _i := makeEqualNode(tokens, i); node != nil {
		i = _i
		p.content = node
		return p, i
	}
	/* TODO: break, continue */
	return nil, i
}

func makeEqualNode(tokens *[]Token, i int) (*EqualNode, int) {
	p := new(EqualNode)
	if node, _i := makeCompNode(tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenDEqual:
			i++
			p.ops = append(p.ops, opEq)
		case TokenNEqual:
			i++
			p.ops = append(p.ops, opNeq)
		default:
			return p, i
		}
		if node, _i := makeCompNode(tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}

func makeCompNode(tokens *[]Token, i int) (*CompNode, int) {
	p := new(CompNode)
	if node, _i := makeExprNode(tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenLe:
			i++
			p.ops = append(p.ops, opLe)
		case TokenLeEqual:
			i++
			p.ops = append(p.ops, opLeEq)
		case TokenGr:
			i++
			p.ops = append(p.ops, opGr)
		case TokenGrEqual:
			i++
			p.ops = append(p.ops, opGrEq)
		default:
			return p, i
		}
		if node, _i := makeExprNode(tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}

func makeExprNode(tokens *[]Token, i int) (*ExprNode, int) {
	p := new(ExprNode)
	if node, _i := makeTermNode(tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenPlus:
			i++
			p.ops = append(p.ops, opPlus)
		case TokenMinus:
			i++
			p.ops = append(p.ops, opMinus)
		default:
			return p, i
		}
		if node, _i := makeTermNode(tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}

func makeTermNode(tokens *[]Token, i int) (*TermNode, int) {
	p := new(TermNode)
	if node, _i := makeFactNode(tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenAst:
			i++
			p.ops = append(p.ops, opMul)
		case TokenSlash:
			i++
			p.ops = append(p.ops, opDiv)
		case TokenPer:
			i++
			p.ops = append(p.ops, opMod)
		default:
			return p, i
		}
		if node, _i := makeFactNode(tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}
func makeFactNode(tokens *[]Token, i int) (*FactNode, int) {
	p := new(FactNode)
	for {
		if node, _i := makeVarNode(tokens, i); node == nil {
			break
		} else {
			flag := true
			switch (*tokens)[_i].(type) {
			case TokenEqual:
				i = _i + 1
				p.ops = append(p.ops, opAssign)
			case TokenPlusAssign:
				i = _i + 1
				p.ops = append(p.ops, opPlusAssign)
			case TokenMinusAssign:
				i = _i + 1
				p.ops = append(p.ops, opMinusAssign)
			case TokenMulAssign:
				i = _i + 1
				p.ops = append(p.ops, opMulAssign)
			case TokenDivAssign:
				i = _i + 1
				p.ops = append(p.ops, opDivAssign)
			default:
				flag = false
			}
			if !flag {
				break
			}
			p.childs = append(p.childs, node)
		}
	}
	if node, _i := makeRvalNode(tokens, i); node == nil {
		return nil, 0
	} else {
		i = _i
		p.rval = node
	}
	return p, i
}

func makeRvalNode(tokens *[]Token, i int) (*RvalNode, int) {
	/* TODO: call, str, char, if, inc, dec, not, true, false */
	p := new(RvalNode)
	if node, _i := makeVarNode(tokens, i); node != nil {
		i = _i
		p.flag = flagVar
		p.content = node
	} else {
		switch v := (*tokens)[i].(type) {
		case TokenNum:
			i++
			node := new(NumNode)
			node.num = v.num
			p.flag = flagNum
			p.content = node
		case TokenOpenBr:
			i++
			if node, _i := makeEqualNode(tokens, i); node == nil {
				return nil, _i
			} else {
				p.flag = flagBracket
				p.content = node
				i = _i
				switch (*tokens)[i].(type) {
				case TokenCloseBr:
					i++
					return p, i
				default:
					return nil, i
				}
			}
		default:
			return nil, i
		}
	}
	return p, i
}

func parseTokenSlice(tokens []Token) {
	root, err := makeSyntaxTree(tokens)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(root)
}
