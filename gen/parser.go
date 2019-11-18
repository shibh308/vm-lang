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
	node, _i := makeProgNode(p, &tokens, 0)
	if node == nil || (_i != tokensLen) {
		msg := "parse error at Token %d, %s\n"
		return nil, fmt.Errorf(msg, _i+1, reflect.TypeOf(tokens[_i]))
	}
	p.prog = node
	return p, nil
}

func makeProgNode(parent PNode, tokens *[]Token, i int) (*ProgNode, int) {
	p := new(ProgNode)
	p.par = parent
	if node, _i := makeFdefNode(p, tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}

	for i < len(*tokens)-ignoreSize {
		if node, _i := makeFdefNode(p, tokens, i); node == nil {
			return p, i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
	return p, i
}

func makeFdefNode(parent PNode, tokens *[]Token, i int) (*FdefNode, int) {
	p := new(FdefNode)
	p.par = parent
	switch v := (*tokens)[i].(type) {
	case TokenVar:
		p.name = v.name
		i++
	default:
		return nil, i
	}

	if node, _i := makeVarsNode(p, tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.vars = node
	}

	if node, _i := makeBlockNode(p, tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.content = node
	}

	return p, i
}

func makeVarsNode(parent PNode, tokens *[]Token, i int) (*VarsNode, int) {
	p := new(VarsNode)
	p.par = parent
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

		if node, _i := makeVarNode(p, tokens, i); node == nil {
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

func makeVarNode(parent PNode, tokens *[]Token, i int) (*VarNode, int) {
	p := new(VarNode)
	p.par = parent
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

func makeBlockNode(parent PNode, tokens *[]Token, i int) (*BlockNode, int) {
	p := new(BlockNode)
	p.par = parent
	switch (*tokens)[i].(type) {
	case TokenOpenWBr:
		i++
		for {
			if node, _i := makeStmthNode(p, tokens, i); node == nil {
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
		if node, _i := makeStmthNode(p, tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.stmts = []PNode{node}
		}
	}
	return p, i
}

func makeStmthNode(parent PNode, tokens *[]Token, i int) (*StmthNode, int) {
	p := new(StmthNode)
	p.par = parent
	if node, _i := makeStmtNode(p, tokens, i); node != nil {
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

func makeStmtNode(parent PNode, tokens *[]Token, i int) (*StmtNode, int) {
	p := new(StmtNode)
	p.par = parent
	p.flag = flagSingleStmt
	switch (*tokens)[i].(type) {
	case TokenReturn:
		i++
		p.flag = flagReturn
	}
	if node, _i := makeEqualNode(p, tokens, i); node != nil {
		i = _i
		p.content = node
		return p, i
	}
	/* TODO: break, continue */
	return nil, i
}

func makeEqualNode(parent PNode, tokens *[]Token, i int) (*EqualNode, int) {
	p := new(EqualNode)
	p.par = parent
	if node, _i := makeCompNode(p, tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenDEqual:
			i++
			p.ops = append(p.ops, oprEq)
		case TokenNEqual:
			i++
			p.ops = append(p.ops, oprNeq)
		default:
			return p, i
		}
		if node, _i := makeCompNode(p, tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}

func makeCompNode(parent PNode, tokens *[]Token, i int) (*CompNode, int) {
	p := new(CompNode)
	p.par = parent
	if node, _i := makeExprNode(p, tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenLe:
			i++
			p.ops = append(p.ops, oprLe)
		case TokenLeEqual:
			i++
			p.ops = append(p.ops, oprLeEq)
		case TokenGr:
			i++
			p.ops = append(p.ops, oprGr)
		case TokenGrEqual:
			i++
			p.ops = append(p.ops, oprGrEq)
		default:
			return p, i
		}
		if node, _i := makeExprNode(p, tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}

func makeExprNode(parent PNode, tokens *[]Token, i int) (*ExprNode, int) {
	p := new(ExprNode)
	p.par = parent
	if node, _i := makeTermNode(p, tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenPlus:
			i++
			p.ops = append(p.ops, oprPlus)
		case TokenMinus:
			i++
			p.ops = append(p.ops, oprMinus)
		default:
			return p, i
		}
		if node, _i := makeTermNode(p, tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}

func makeTermNode(parent PNode, tokens *[]Token, i int) (*TermNode, int) {
	p := new(TermNode)
	p.par = parent
	if node, _i := makeFactNode(p, tokens, i); node == nil {
		return nil, _i
	} else {
		i = _i
		p.childs = append(p.childs, node)
	}
	for {
		switch (*tokens)[i].(type) {
		case TokenAst:
			i++
			p.ops = append(p.ops, oprMul)
		case TokenSlash:
			i++
			p.ops = append(p.ops, oprDiv)
		case TokenPer:
			i++
			p.ops = append(p.ops, oprMod)
		default:
			return p, i
		}
		if node, _i := makeFactNode(p, tokens, i); node == nil {
			return nil, _i
		} else {
			i = _i
			p.childs = append(p.childs, node)
		}
	}
}

func makeFactNode(parent PNode, tokens *[]Token, i int) (*FactNode, int) {
	p := new(FactNode)
	p.par = parent
	for {
		if node, _i := makeVarNode(p, tokens, i); node == nil {
			break
		} else {
			flag := true
			switch (*tokens)[_i].(type) {
			case TokenEqual:
				i = _i + 1
				p.ops = append(p.ops, oprAssign)
			case TokenPlusAssign:
				i = _i + 1
				p.ops = append(p.ops, oprPlusAssign)
			case TokenMinusAssign:
				i = _i + 1
				p.ops = append(p.ops, oprMinusAssign)
			case TokenMulAssign:
				i = _i + 1
				p.ops = append(p.ops, oprMulAssign)
			case TokenDivAssign:
				i = _i + 1
				p.ops = append(p.ops, oprDivAssign)
			case TokenModAssign:
				i = _i + 1
				p.ops = append(p.ops, oprModAssign)
			default:
				flag = false
			}
			if !flag {
				break
			}
			p.lvals = append(p.lvals, node)
		}
	}
	if node, _i := makeRvalNode(p, tokens, i); node == nil {
		return nil, 0
	} else {
		i = _i
		p.rval = node
	}
	return p, i
}

func makeRvalNode(parent PNode, tokens *[]Token, i int) (*RvalNode, int) {
	p := new(RvalNode)
	p.par = parent
	/* TODO: call, str, char, if, inc, dec, not, true, false */
	if node, _i := makeVarNode(p, tokens, i); node != nil {
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
			if node, _i := makeEqualNode(p, tokens, i); node == nil {
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

func parseTokenSlice(tokens []Token) *RootNode {
	root, err := makeSyntaxTree(tokens)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return root
}
