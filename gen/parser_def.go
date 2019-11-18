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

type PNode interface {
	getPar() PNode
}

type ParData struct {
	par PNode
}

func (pd ParData) getPar() PNode { return pd.par }

type FuncData struct {
	name      string
	idx       int
	node      *FdefNode
	variables []string
}

type RootNode struct {
	prog      *ProgNode
	functions []FuncData
	funcMap   map[string]*FuncData
	ParData
}

type ProgNode struct {
	childs []PNode
	ParData
}
type FdefNode struct {
	name    string
	vars    *VarsNode
	content *BlockNode
	ParData
}
type BlockNode struct {
	stmts []*StmthNode
	ParData
}
type StmthNode struct {
	stmt *StmtNode
	flag StmthFlag
	ParData
}
type StmtNode struct {
	content *EqualNode
	flag    StmtFlag
	ParData
}
type EqualNode struct {
	childs []*CompNode
	ops    []EqualOpr
	ParData
}
type CompNode struct {
	childs []*ExprNode
	ops    []CompOpr
	ParData
}
type ExprNode struct {
	childs []*TermNode
	ops    []ExprOpr
	ParData
}
type TermNode struct {
	childs []*FactNode
	ops    []TermOpr
	ParData
}
type FactNode struct {
	lvals []*VarNode
	ops   []AssignOpr
	rval  *RvalNode
	ParData
}
type RvalNode struct {
	flag    RvalFlag
	content PNode
	ParData
}
type CallNode struct {
	childs []PNode
	ParData
}
type IfNode struct {
	childs []PNode
	ParData
}
type ForNode struct {
	childs []PNode
	ParData
}
type WhileNode struct {
	childs []PNode
	ParData
}
type VarsNode struct {
	args []*VarNode
	ParData
}
type VarNode struct {
	name  string
	isRef bool
	ParData
}
type NumNode struct {
	num int
	ParData
}
type StrNode struct {
	childs []PNode
	str    string
	ParData
}
type CharNode struct {
	childs []PNode
	char   uint8
	ParData
}

/* TODO: struct */
// type StructNode struct{ Node }
// type DefNode struct{ Node ResultType }

/* TODO: seq */
