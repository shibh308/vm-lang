package main

type ResultType int
type StmtFlag int
type EqualFlag int
type CompFlag int
type ExprFlag int
type TermFlag int
type AssignFlag int
type RvalFlag int

const (
	typeInt ResultType = iota
	typeBool
	typeSeq
	typeVoid
)
const (
	flagEqual StmtFlag = iota
	flagReturn
	flagBreak
	flagContinue
)
const (
	flagSingleEqual EqualFlag = iota
	flagEq
	flagNeq
)
const (
	flagSingleComp CompFlag = iota
	flagGr
	flagLe
	flagGrEq
	flagLeEq
)
const (
	flagSingleExpr ExprFlag = iota
	flagPlus
	flagMinus
)
const (
	flagSingleTerm TermFlag = iota
	flagMul
	flagDiv
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

type PNode interface {
	parseTokens(*[]Token, int) (int, bool)
}

type RootNode struct{ prog *ProgNode }
type ProgNode struct{ childs []PNode }
type FdefNode struct {
	name    string
	vars    *VarsNode
	content *BlockNode
}
type BlockNode struct{ childs []PNode }
type StmthNode struct{ childs []PNode }
type StmtNode struct{ childs []PNode }
type EqualNode struct {
	childs []PNode
	result ResultType
}
type CompNode struct {
	childs []PNode
	result ResultType
}
type ExprNode struct {
	childs []PNode
	result ResultType
}
type TermNode struct {
	childs []PNode
	result ResultType
}
type FactNode struct {
	childs []PNode
	result ResultType
}
type RvalNode struct {
	childs []PNode
	result ResultType
}
type CallNode struct {
	childs []PNode
	result ResultType
}
type IfNode struct {
	childs []PNode
	result ResultType
}
type ForNode struct{ childs []PNode }
type WhileNode struct{ childs []PNode }
type VarsNode struct {
	args []*VarNode
}
type RetNode struct {
	childs []PNode
	result ResultType
}
type VarNode struct {
	name   string
	isRef  bool
	result ResultType
}
type NumNode struct {
	childs []PNode
	num    int
	result ResultType
}
type StrNode struct {
	childs []PNode
	str    string
	result ResultType
}
type CharNode struct {
	childs []PNode
	char   uint8
	result ResultType
}
type AssignNode struct {
	childs []PNode
	result ResultType
}

/* TODO struct */
// type StructNode struct{ Node }
// type DefNode struct{ Node ResultType }

/* TODO seq */
