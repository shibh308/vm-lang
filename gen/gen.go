package main

import (
	"fmt"
	"os"
)

const (
	opExtra opCode = iota

	opRead
	opPrint

	opCopy

	opAdd
	opSub
	opMul
	opDiv
	opMod

	opEq
	opNeq
	opGr
	opLe
	opGreq
	opLeeq

	opJump
	opIf
	opCall

	opDef
	opReturn
	opAssign

	opGet
	opSet
)

func (b byteCode) print() {
	var s string
	switch b.code {
	case opExtra:
		s = "extra"
	case opRead:
		s = "read"
	case opPrint:
		s = "print"
	case opCopy:
		s = "copy"
	case opAdd:
		s = "add"
	case opSub:
		s = "sub"
	case opMul:
		s = "mul"
	case opDiv:
		s = "div"
	case opEq:
		s = "eq"
	case opNeq:
		s = "neq"
	case opGr:
		s = "gr"
	case opLe:
		s = "le"
	case opGreq:
		s = "greq"
	case opLeeq:
		s = "leeq"
	case opJump:
		s = "jump"
	case opIf:
		s = "if"
	case opCall:
		s = "call"
	case opDef:
		s = "def"
	case opReturn:
		s = "return"
	case opAssign:
		s = "assign"
	case opGet:
		s = "get"
	case opSet:
		s = "set"
	default:
		s = "???"
	}
	fmt.Printf("%6s ", s)
	for _, arg := range b.rand {
		fmt.Printf(" %d", arg)
	}
	fmt.Printf("\n")
}

type byteCode struct {
	code opCode
	rand []int
}

func (root *RootNode) makeOpExtra(arg int, val int) {
	root.code = append(root.code, byteCode{code: opExtra, rand: []int{arg, val}})
}
func (root *RootNode) makeOpRead(dst int) {
	root.code = append(root.code, byteCode{code: opRead, rand: []int{dst}})
}
func (root *RootNode) makeOpPrint(src int) {
	root.code = append(root.code, byteCode{code: opPrint, rand: []int{src}})
}
func (root *RootNode) makeOpCopy(src int, dst int) {
	root.code = append(root.code, byteCode{code: opCopy, rand: []int{src, dst}})
}
func (root *RootNode) makeOpAdd(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opAdd, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpSub(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opSub, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpMul(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opMul, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpDiv(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opDiv, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpMod(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opMod, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpEq(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opEq, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpNeq(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opNeq, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpGr(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opGr, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpLe(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opLe, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpGreq(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opGreq, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpLeeq(src1 int, src2 int, dst int) {
	root.code = append(root.code, byteCode{code: opLeeq, rand: []int{src1, src2, dst}})
}
func (root *RootNode) makeOpJump(label int) {
	root.code = append(root.code, byteCode{code: opJump, rand: []int{label}})
}
func (root *RootNode) makeOpIf(reg int, label int) {
	root.code = append(root.code, byteCode{code: opIf, rand: []int{reg, label}})
}
func (root *RootNode) makeOpDef(def int) {
	root.code = append(root.code, byteCode{code: opDef, rand: []int{def}})
}
func (root *RootNode) makeOpReturn() { root.code = append(root.code, byteCode{code: opReturn}) }
func (root *RootNode) makeOpCall(src int, dst int, def int) {
	root.code = append(root.code, byteCode{code: opCall, rand: []int{src, dst, def}})
}
func (root *RootNode) makeOpAssign(reg int, val int) {
	root.code = append(root.code, byteCode{code: opAssign, rand: []int{reg, val}})
}
func (root *RootNode) makeOpGet(dst int, mem int) {
	root.code = append(root.code, byteCode{code: opGet, rand: []int{dst, mem}})
}
func (root *RootNode) makeOpSet(src int, mem int) {
	root.code = append(root.code, byteCode{code: opSet, rand: []int{src, mem}})
}

func getVariables(p PNode) []string {
	var names []string
	switch node := p.(type) {
	case *FactNode:
		for _, lval := range node.lvals {
			switch lv := lval.(type) {
			case *VarNode:
				names = append(names, lv.name)
			}
		}
		names = append(names, getVariables(node.rval)...)
	default:
		for _, child := range node.getChilds() {
			names = append(names, getVariables(child)...)
		}
	}
	return names
}

func (root *RootNode) captureVariable() error {
	root.funcMap = map[string]*FuncData{}
	for _, node := range root.prog.childs {
		switch fn := node.(type) {
		case *FdefNode:
			_, exist := root.funcMap[fn.name]
			if exist {
				return fmt.Errorf(`function redeclared "%s"\n`, fn.name)
			}
			root.functions = append(root.functions, FuncData{name: fn.name, node: fn, variables: []string{}, varMap: map[string]int{}})
		}
	}
	fl := false
	for i := 0; i < len(root.functions); i++ {
		if root.functions[i].name == "main" {
			if i != 0 {
				root.functions[i], root.functions[0] = root.functions[0], root.functions[i]
			}
			fl = true
			break
		}
	}
	if !fl {
		return fmt.Errorf(`main function is not defined\n`)
	}
	for i := 0; i < len(root.functions); i++ {
		funcData := &root.functions[i]
		root.funcMap[funcData.name] = &root.functions[i]
		funcData.idx = i
		var variables []string
		for _, node := range funcData.node.vars {
			switch arg := node.(type) {
			case *VarNode:
				variables = append(variables, arg.name)
			}
		}
		variables = append(variables, getVariables(funcData.node.content)...)
		funcData.varCnt = len(variables)
		funcData.variables = variables
		funcData.varMap = map[string]int{}
		for i, variable := range variables {
			funcData.varMap[variable] = i
		}
	}
	return nil
}

func (f *FuncData) getReg(name string) int {
	return f.varMap[name]
}

func (root *RootNode) useMultiRegs(size int, funcData *FuncData) int {
	var i int
	match := 0
	for i = 0; match < size; i++ {
		if i >= len(root.reg) {
			root.reg = append(root.reg, 0)
			break
		}
		if root.reg[i] == 0 {
			match++
		} else if root.reg[i] == 1 {
			match = 0
		}
	}
	return i - size
}

func (root *RootNode) useReg(funcData *FuncData) int {
	var i int
	for i = 0; ; i++ {
		if i >= len(root.reg) {
			root.reg = append(root.reg, 1)
			break
		}
		if root.reg[i] == 0 {
			root.reg[i] = 1
			break
		}
	}
	return i
}

func (root *RootNode) unUseReg(i int, funcData *FuncData) {
	if root.reg[i] != 2 {
		root.reg[i] = 0
	}
}

func (root *RootNode) genOpCode(p PNode, funcData *FuncData) int {
	switch node := p.(type) {
	case *FdefNode:
		root.makeOpDef(funcData.idx)
		root.makeOpAssign(1, funcData.varCnt)
		ret := root.genOpCode(node.content, funcData)
		root.makeOpCopy(funcData.varCnt+2, ret)
		root.makeOpReturn()
		return -1
	case *CallNode:
		s := node.name
		callFunc := root.funcMap[s]
		st := root.useMultiRegs(funcData.varCnt, funcData)
		for i, argNode := range node.args {
			reg := root.genOpCode(argNode, funcData)
			root.makeOpCopy(reg, st+i)
		}
		reg := root.useReg(funcData)
		root.makeOpCall(st, reg, callFunc.idx)
		for i := 0; i < len(node.args); i++ {
			root.unUseReg(st+i, funcData)
		}
		return reg
	case *EqualNode:
		src1 := root.genOpCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genOpCode(node.childs[i+1], funcData)
			dst := root.useReg(funcData)
			switch node.ops[i] {
			case oprEq:
				root.makeOpEq(src1, src2, dst)
			case oprNeq:
				root.makeOpNeq(src1, src2, dst)
			}
			root.unUseReg(src1, funcData)
			root.unUseReg(src2, funcData)
			src1 = dst
		}
		return src1
	case *CompNode:
		src1 := root.genOpCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genOpCode(node.childs[i+1], funcData)
			dst := root.useReg(funcData)
			switch node.ops[i] {
			case oprGr:
				root.makeOpGr(src1, src2, dst)
			case oprLe:
				root.makeOpLe(src1, src2, dst)
			case oprGrEq:
				root.makeOpGreq(src1, src2, dst)
			case oprLeEq:
				root.makeOpLeeq(src1, src2, dst)
			}
			root.unUseReg(src1, funcData)
			root.unUseReg(src2, funcData)
			src1 = dst
		}
		return src1
	case *ExprNode:
		src1 := root.genOpCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genOpCode(node.childs[i+1], funcData)
			dst := root.useReg(funcData)
			switch node.ops[i] {
			case oprPlus:
				root.makeOpAdd(src1, src2, dst)
			case oprMinus:
				root.makeOpSub(src1, src2, dst)
			}
			root.unUseReg(src1, funcData)
			root.unUseReg(src2, funcData)
			src1 = dst
		}
		return src1
	case *TermNode:
		src1 := root.genOpCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genOpCode(node.childs[i+1], funcData)
			dst := root.useReg(funcData)
			switch node.ops[i] {
			case oprMul:
				root.makeOpMul(src1, src2, dst)
			case oprDiv:
				root.makeOpDiv(src1, src2, dst)
			}
			root.unUseReg(src1, funcData)
			root.unUseReg(src2, funcData)
			src1 = dst
		}
		return src1
	case *FactNode:
		rvalReg := root.genOpCode(node.rval, funcData)
		for i := len(node.ops) - 1; i >= 0; i-- {
			var lvalReg int
			switch lv := node.lvals[i].(type) {
			case *VarNode:
				lvalReg = funcData.getReg(lv.name)
			}
			switch node.ops[i] {
			case oprAssign:
				root.makeOpCopy(rvalReg, lvalReg)
			case oprPlusAssign:
				root.makeOpAdd(lvalReg, rvalReg, lvalReg)
			case oprMinusAssign:
				root.makeOpSub(lvalReg, rvalReg, lvalReg)
			case oprMulAssign:
				root.makeOpMul(lvalReg, rvalReg, lvalReg)
			case oprDivAssign:
				root.makeOpDiv(lvalReg, rvalReg, lvalReg)
			case oprModAssign:
				root.makeOpMod(lvalReg, rvalReg, lvalReg)
			}
			rvalReg = lvalReg
		}
		return rvalReg
	case *RvalNode:
		switch node.flag {
		case flagVar:
			var name string
			switch variable := node.content.(type) {
			case *VarNode:
				name = variable.name
			}
			return funcData.getReg(name)
		case flagNum:
			var num int
			switch numNode := node.content.(type) {
			case *NumNode:
				num = numNode.num
			}
			reg := root.useReg(funcData)
			root.makeOpAssign(reg, num)
			return reg
		case flagBracket:
			reg := root.genOpCode(node.content, funcData)
			return reg
		case flagCall:
			reg := root.genOpCode(node.content, funcData)
			return reg
		default:
			_, _ = fmt.Fprintf(os.Stderr, `unknown rvalFlag type: %d\n`, node.flag)
			return -1
		}
	default:
		var ret int
		for _, child := range p.getChilds() {
			ret = root.genOpCode(child, funcData)
		}
		return ret
	}
}

func (root *RootNode) generateOpCode() {
	err := root.captureVariable()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	for _, funcData := range root.functions {
		root.reg = make([]uint8, funcData.varCnt+3)
		for i := 0; i < funcData.varCnt+3; i++ {
			root.reg[i] = 2
		}
		root.genOpCode(funcData.node, &funcData)
	}
	for _, byteCode := range root.code {
		byteCode.print()
	}
}
