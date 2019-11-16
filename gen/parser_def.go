package main

type ResultType int
type StmtFlag int
type StmthFlag int
type EqualOp int
type CompOp int
type ExprOp int
type TermOp int
type AssignFlag int
type RvalFlag int

const (
	typeInt ResultType = iota
	typeBool
	typeSeq
	typeVoid
)
const (
	flagSingleStmth StmthFlag = iota
	flagFor
	flagWhile
)
const (
	flagSingleStmt StmtFlag = iota
	flagReturn
	flagBreak
	flagContinue
)
const (
	opSingleEqual EqualOp = iota
	opEq
	opNeq
)
const (
	opSingleComp CompOp = iota
	opGr
	opLe
	opGrEq
	opLeEq
)
const (
	opSingleExpr ExprOp = iota
	opPlus
	opMinus
)
const (
	opSingleTerm TermOp = iota
	opMul
	opDiv
)
const (
	flagAssign AssignFlag = iota
	flagPlusAssign
	flagMinusAssign
	flagMulAssign
	flagDivAssign
)
const (
	flagCall RvalFlag = iota
	flagVarh
	flagNum
	flagStr
	flagChar
	flagBool
	flagBracket
	flagIf
	flagInc
	flagDec
	flagRev
)

const ignoreSize = 10

type PNode interface{}

type RootNode struct{ prog *ProgNode }
type ProgNode struct{ childs []PNode }
type FdefNode struct {
	name    string
	vars    *VarsNode
	content *BlockNode
}
type BlockNode struct{ stmts []*StmthNode }
type StmthNode struct {
	stmt *StmtNode
	flag StmthFlag
}
type StmtNode struct {
	content *EqualNode
	flag    StmtFlag
}
type EqualNode struct {
	childs []PNode
	ops    []EqualOp
}
type CompNode struct {
	childs []PNode
	ops    []CompOp
}
type ExprNode struct {
	childs []PNode
	ops    []ExprOp
}
type TermNode struct {
	childs []PNode
	ops    []TermOp
}
type FactNode struct {
	childs []PNode
}
type RvalNode struct {
	childs []PNode
}
type CallNode struct {
	childs []PNode
}
type IfNode struct {
	childs []PNode
}
type ForNode struct{ childs []PNode }
type WhileNode struct{ childs []PNode }
type VarsNode struct {
	args []*VarNode
}
type VarNode struct {
	name  string
	isRef bool
}
type NumNode struct {
	childs []PNode
	num    int
}
type StrNode struct {
	childs []PNode
	str    string
}
type CharNode struct {
	childs []PNode
	char   uint8
}
type AssignNode struct {
	childs []PNode
}

/* TODO: struct */
// type StructNode struct{ Node }
// type DefNode struct{ Node ResultType }

/* TODO: seq */
