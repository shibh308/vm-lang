package main

type ResultType int
type StmtFlag int
type StmthFlag int
type EqualOpr int
type CompOpr int
type ExprOpr int
type TermOpr int
type AssignOpr int
type RvalFlag int

type opCode int

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
	oprSingleEqual EqualOpr = iota
	oprEq
	oprNeq
)
const (
	oprSingleComp CompOpr = iota
	oprGr
	oprLe
	oprGrEq
	oprLeEq
)
const (
	oprSingleExpr ExprOpr = iota
	oprPlus
	oprMinus
)
const (
	oprSingleTerm TermOpr = iota
	oprMul
	oprDiv
	oprMod
)
const (
	oprAssign AssignOpr = iota
	oprPlusAssign
	oprMinusAssign
	oprMulAssign
	oprDivAssign
	oprModAssign
)
const (
	flagCall RvalFlag = iota
	flagVar
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

const ()

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
	childs []*CompNode
	ops    []EqualOpr
}
type CompNode struct {
	childs []*ExprNode
	ops    []CompOpr
}
type ExprNode struct {
	childs []*TermNode
	ops    []ExprOpr
}
type TermNode struct {
	childs []*FactNode
	ops    []TermOpr
}
type FactNode struct {
	childs []*VarNode
	ops    []AssignOpr
	rval   *RvalNode
}
type RvalNode struct {
	flag    RvalFlag
	content PNode
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
	num int
}
type StrNode struct {
	childs []PNode
	str    string
}
type CharNode struct {
	childs []PNode
	char   uint8
}

/* TODO: struct */
// type StructNode struct{ Node }
// type DefNode struct{ Node ResultType }

/* TODO: seq */
