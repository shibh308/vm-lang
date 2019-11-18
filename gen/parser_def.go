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

const ignoreSize = 10

type PNode interface{}

type FuncData struct {
	name      string
	idx       int
	node      *FdefNode
	variables []string
}

type RootNode struct {
	prog      *ProgNode
	functions []FuncData
}

type ProgNode struct {
	childs []PNode
	par    PNode
}
type FdefNode struct {
	name    string
	vars    *VarsNode
	content *BlockNode
	par     PNode
}
type BlockNode struct {
	stmts []*StmthNode
	par   PNode
}
type StmthNode struct {
	stmt *StmtNode
	flag StmthFlag
	par  PNode
}
type StmtNode struct {
	content *EqualNode
	flag    StmtFlag
	par     PNode
}
type EqualNode struct {
	childs []*CompNode
	ops    []EqualOpr
	par    PNode
}
type CompNode struct {
	childs []*ExprNode
	ops    []CompOpr
	par    PNode
}
type ExprNode struct {
	childs []*TermNode
	ops    []ExprOpr
	par    PNode
}
type TermNode struct {
	childs []*FactNode
	ops    []TermOpr
	par    PNode
}
type FactNode struct {
	childs []*VarNode
	ops    []AssignOpr
	rval   *RvalNode
	par    PNode
}
type RvalNode struct {
	flag    RvalFlag
	content PNode
	par     PNode
}
type CallNode struct {
	childs []PNode
	par    PNode
}
type IfNode struct {
	childs []PNode
	par    PNode
}
type ForNode struct {
	childs []PNode
	par    PNode
}
type WhileNode struct {
	childs []PNode
	par    PNode
}
type VarsNode struct {
	args []*VarNode
	par  PNode
}
type VarNode struct {
	name  string
	isRef bool
	par   PNode
}
type NumNode struct {
	num int
	par PNode
}
type StrNode struct {
	childs []PNode
	str    string
	par    PNode
}
type CharNode struct {
	childs []PNode
	char   uint8
	par    PNode
}

/* TODO: struct */
// type StructNode struct{ Node }
// type DefNode struct{ Node ResultType }

/* TODO: seq */
