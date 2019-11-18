package main

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
	opLabel
	opDef
)

type byteCode struct {
	code opCode
	rand []int
}

func makeOpExtra(arg int, val int) byteCode { return byteCode{code: opExtra, rand: []int{arg, val}} }
func makeOpRead(dst int) byteCode           { return byteCode{code: opRead, rand: []int{dst}} }
func makeOpPrint(src int) byteCode          { return byteCode{code: opPrint, rand: []int{src}} }
func makeOpCopy(src int, dst int) byteCode  { return byteCode{code: opCopy, rand: []int{src, dst}} }
func makeOpAdd(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opAdd, rand: []int{src1, src2, dst}}
}
func makeOpSub(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opSub, rand: []int{src1, src2, dst}}
}
func makeOpMul(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opMul, rand: []int{src1, src2, dst}}
}
func makeOpDiv(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opDiv, rand: []int{src1, src2, dst}}
}
func makeOpMod(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opMod, rand: []int{src1, src2, dst}}
}
func makeOpEq(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opEq, rand: []int{src1, src2, dst}}
}
func makeOpNeq(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opNeq, rand: []int{src1, src2, dst}}
}
func makeOpGr(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opGr, rand: []int{src1, src2, dst}}
}
func makeOpLe(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opLe, rand: []int{src1, src2, dst}}
}
func makeOpGreq(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opGreq, rand: []int{src1, src2, dst}}
}
func makeOpLeeq(src1 int, src2 int, dst int) byteCode {
	return byteCode{code: opLeeq, rand: []int{src1, src2, dst}}
}
func makeOpJump(label int) byteCode        { return byteCode{code: opJump, rand: []int{label}} }
func makeOpIf(reg int, label int) byteCode { return byteCode{code: opIf, rand: []int{reg, label}} }
func makeReturn() byteCode                 { return byteCode{code: opReturn} }
func makeCall(reg1 int, reg2 int, def int) byteCode {
	return byteCode{code: opCall, rand: []int{reg1, reg2, def}}
}
func makeAssign(reg int, val int) byteCode { return byteCode{code: opAssign, rand: []int{reg, val}} }
func makeGet(dst int, mem int) byteCode    { return byteCode{code: opGet, rand: []int{dst, mem}} }
func makeSet(src int, mem int) byteCode    { return byteCode{code: opSet, rand: []int{src, mem}} }
func makeLabel() byteCode                  { return byteCode{code: opLabel} }
func makeDef() byteCode                    { return byteCode{code: opDef} }

func (root *RootNode) captureVariable() {
	for i, node := range root.prog.childs {
		switch fn := node.(type) {
		case *FdefNode:
			root.functions = append(root.functions, FuncData{idx: i, name: fn.name, node: fn})
		}
	}
	for _, funcData := range root.functions {
		// TODO: get assign nodes
	}
}

func (root *RootNode) genOpCode() []opCode {
	return true
}
