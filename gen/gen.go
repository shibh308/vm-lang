package main

import (
	"encoding/binary"
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
	fmt.Printf("%3d", b.code)
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
		funcData.argCnt = len(variables)
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
		}
		if root.reg[i] == 0 {
			match++
		} else {
			match = 0
		}
	}
	for j := 0; j < size; j++ {
		root.reg[i-size+j] = 1
	}
	return i - size
}

func (root *RootNode) useReg() int {
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

func (root *RootNode) unUseReg(i int) {
	if i != -1 && root.reg[i] != 2 {
		root.reg[i] = 0
	}
}

func (root *RootNode) genByteCode(p PNode, funcData *FuncData) int {
	switch node := p.(type) {
	case *FdefNode:
		funcData.line = len(root.code)
		ret := root.genByteCode(node.content, funcData)
		if ret == -1 {
			_, _ = fmt.Fprintln(os.Stderr, "missing a return statement in function block")
			os.Exit(1)
		}
		root.makeOpCopy(ret, 0)
		root.unUseReg(ret)
		root.makeOpReturn()
		return -1
	case *CallNode:
		s := node.name
		var argCnt int
		var funcIdx int
		builtIn := false
		if s == "read" {
			builtIn = true
			argCnt = 0
		} else if s == "print" {
			builtIn = true
			argCnt = 1
		} else {
			callFunc, exist := root.funcMap[s]
			if !exist {
				_, _ = fmt.Fprintf(os.Stderr, `calling an undeclared function "%s"`, s)
				os.Exit(1)
			}
			argCnt = callFunc.argCnt
			funcIdx = callFunc.idx
		}
		st := root.useMultiRegs(argCnt, funcData)
		for i, argNode := range node.args {
			reg := root.genByteCode(argNode, funcData)
			root.makeOpCopy(reg, st+i)
			root.unUseReg(reg)
		}
		reg := root.useReg()
		if builtIn {
			switch s {
			case "read":
				root.makeOpRead(reg)
			case "print":
				root.makeOpPrint(st)
				root.unUseReg(reg)
				reg = st
			}
		} else {
			root.makeOpCall(st, reg, funcIdx)
		}
		for i := 0; i < len(node.args); i++ {
			root.unUseReg(st + i)
		}
		return reg
	case *IfNode:
		comp := root.genByteCode(node.comp, funcData)
		idx := len(root.code)
		root.makeOpIf(comp, 0)
		root.unUseReg(comp)
		root.genByteCode(node.content, funcData)
		root.code[idx].rand[1] = len(root.code) - 1
		return -1
	case *StmtNode:
		switch node.flag {
		case flagReturn:
			ret := root.genByteCode(node.content, funcData)
			root.makeOpCopy(ret, 0)
			root.unUseReg(ret)
			root.makeOpReturn()
			return -1
		default:
			return root.genByteCode(node.content, funcData)
		}
	case *EqualNode:
		src1 := root.genByteCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genByteCode(node.childs[i+1], funcData)
			dst := root.useReg()
			switch node.ops[i] {
			case oprEq:
				root.makeOpEq(src1, src2, dst)
			case oprNeq:
				root.makeOpNeq(src1, src2, dst)
			}
			root.unUseReg(src1)
			root.unUseReg(src2)
			src1 = dst
		}
		return src1
	case *CompNode:
		src1 := root.genByteCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genByteCode(node.childs[i+1], funcData)
			dst := root.useReg()
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
			root.unUseReg(src1)
			root.unUseReg(src2)
			src1 = dst
		}
		return src1
	case *ExprNode:
		src1 := root.genByteCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genByteCode(node.childs[i+1], funcData)
			dst := root.useReg()
			switch node.ops[i] {
			case oprPlus:
				root.makeOpAdd(src1, src2, dst)
			case oprMinus:
				root.makeOpSub(src1, src2, dst)
			}
			root.unUseReg(src1)
			root.unUseReg(src2)
			src1 = dst
		}
		return src1
	case *TermNode:
		src1 := root.genByteCode(node.childs[0], funcData)
		for i := 0; i < len(node.ops); i++ {
			src2 := root.genByteCode(node.childs[i+1], funcData)
			dst := root.useReg()
			switch node.ops[i] {
			case oprMul:
				root.makeOpMul(src1, src2, dst)
			case oprDiv:
				root.makeOpDiv(src1, src2, dst)
			}
			root.unUseReg(src1)
			root.unUseReg(src2)
			src1 = dst
		}
		return src1
	case *FactNode:
		rvalReg := root.genByteCode(node.rval, funcData)
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
			root.unUseReg(rvalReg)
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
			reg := root.useReg()
			root.makeOpAssign(reg, num)
			return reg
		case flagBracket:
			reg := root.genByteCode(node.content, funcData)
			return reg
		case flagCall:
			reg := root.genByteCode(node.content, funcData)
			return reg
		default:
			_, _ = fmt.Fprintf(os.Stderr, `unknown rvalFlag type: %d\n`, node.flag)
			return -1
		}
	default:
		ret := -1
		for _, child := range p.getChilds() {
			if ret != -1 {
				root.unUseReg(ret)
			}
			ret = root.genByteCode(child, funcData)
		}
		return ret
	}
}

func (root *RootNode) generateByteCode() {
	err := root.captureVariable()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	for i := 0; i < len(root.functions); i++ {
		funcData := &root.functions[i]
		for key, _ := range funcData.varMap {
			funcData.varMap[key]++
		}
		root.reg = make([]uint8, funcData.varCnt+1)
		for i := 0; i < funcData.varCnt+1; i++ {
			root.reg[i] = 2
		}
		root.genByteCode(funcData.node, funcData)
		funcData.code = root.code
		root.code = []byteCode{}
		funcData.regSize = len(root.reg)
	}
}

func (root *RootNode) printByteCode() {
	for i, funcData := range root.functions {
		fmt.Printf(`%4d:  "%s"  %d %d %d`+"\n", i, funcData.name, len(funcData.code), funcData.regSize, funcData.argCnt)
		for i, byteCode := range funcData.code {
			fmt.Printf("%4d:  ", i)
			byteCode.print()
		}
	}
}

func write(f *os.File, val uint32) {
	err := binary.Write(f, binary.LittleEndian, val)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "binary.Write fialed: %s\n", err)
		os.Exit(1)
	}
}

func (root *RootNode) writeByteCode(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "os.Create failed: %s\n", err)
		os.Exit(1)
	}
	write(f, uint32(len(root.functions)))
	for _, fn := range root.functions {
		write(f, uint32(fn.line))
		write(f, uint32(len(fn.code)))
		write(f, uint32(fn.regSize)|(uint32(fn.argCnt)<<16))
		for _, byteCode := range fn.code {
			/* TODO: opExtra */
			val := 0
			op := byteCode.code
			switch byteCode.code {
			// reg:0
			case opJump, opReturn:
				if len(byteCode.rand) == 1 {
					val += byteCode.rand[0]
				}
				val <<= 6
				val += int(op)
			// reg:1
			case opRead, opPrint, opIf, opAssign, opGet, opSet:
				if len(byteCode.rand) == 2 {
					val += byteCode.rand[1]
				}
				val <<= 9
				val += byteCode.rand[0]
				val <<= 6
				val += int(op)
			// reg:2
			case opCopy, opCall:
				if len(byteCode.rand) == 3 {
					val += byteCode.rand[2]
				}
				val <<= 9
				val += byteCode.rand[1]
				val <<= 9
				val += byteCode.rand[0]
				val <<= 6
				val += int(op)
			// reg:3
			default:
				if len(byteCode.rand) == 4 {
					val += byteCode.rand[3]
				}
				val <<= 9
				val += byteCode.rand[2]
				val <<= 9
				val += byteCode.rand[1]
				val <<= 9
				val += byteCode.rand[0]
				val <<= 6
				val += int(op)
			}
			write(f, uint32(val))
		}
	}
}
