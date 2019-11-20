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
	flagIf
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
	oprEq EqualOpr = iota
	oprNeq
)
const (
	oprGr CompOpr = iota
	oprLe
	oprGrEq
	oprLeEq
)
const (
	oprPlus ExprOpr = iota
	oprMinus
)
const (
	oprMul TermOpr = iota
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
	flagInc
	flagDec
	flagRev
)

const ignoreSize = 10

type PNode interface {
	getPar() PNode
	getChilds() []PNode
}

type ParData struct {
	par PNode
}

func (pd *ParData) getPar() PNode { return pd.par }

type FuncData struct {
	name      string
	idx       int
	line      int
	node      *FdefNode
	variables []string
	argCnt    int
	varCnt    int
	regSize   int
	varMap    map[string]int
}

type RootNode struct {
	prog      *ProgNode
	functions []FuncData
	funcMap   map[string]*FuncData
	code      []byteCode
	reg       []uint8
	ParData
}

func (p *RootNode) getChilds() []PNode { return []PNode{p.prog} }

type ProgNode struct {
	childs []PNode
	ParData
}

func (p *ProgNode) getChilds() []PNode { return p.childs }

type FdefNode struct {
	name    string
	vars    []PNode
	content *BlockNode
	ParData
}

func (p *FdefNode) getChilds() []PNode { return append(p.vars, p.content) }

type BlockNode struct {
	stmts []PNode
	ParData
}

func (p *BlockNode) getChilds() []PNode { return p.stmts }

type StmthNode struct {
	stmt PNode
	flag StmthFlag
	ParData
}

func (p *StmthNode) getChilds() []PNode { return []PNode{p.stmt} }

type StmtNode struct {
	content *FactNode
	flag    StmtFlag
	ParData
}

func (p *StmtNode) getChilds() []PNode { return []PNode{p.content} }

type EqualNode struct {
	childs []PNode
	ops    []EqualOpr
	ParData
}

func (p *EqualNode) getChilds() []PNode { return p.childs }

type CompNode struct {
	childs []PNode
	ops    []CompOpr
	ParData
}

func (p *CompNode) getChilds() []PNode { return p.childs }

type ExprNode struct {
	childs []PNode
	ops    []ExprOpr
	ParData
}

func (p *ExprNode) getChilds() []PNode { return p.childs }

type TermNode struct {
	childs []PNode
	ops    []TermOpr
	ParData
}

func (p *TermNode) getChilds() []PNode { return p.childs }

type FactNode struct {
	lvals []PNode
	ops   []AssignOpr
	rval  *EqualNode
	ParData
}

func (p *FactNode) getChilds() []PNode { return append(p.lvals, p.rval) }

type RvalNode struct {
	flag    RvalFlag
	content PNode
	ParData
}

func (p *RvalNode) getChilds() []PNode { return []PNode{p.content} }

type CallNode struct {
	name string
	args []PNode
	ParData
}

func (p *CallNode) getChilds() []PNode { return p.args }

type IfNode struct {
	comp    PNode
	content PNode
	ParData
}

func (p *IfNode) getChilds() []PNode { return []PNode{p.comp, p.content} }

type ForNode struct {
	childs []PNode
	ParData
}

func (p *ForNode) getChilds() []PNode { return p.childs }

type WhileNode struct {
	childs []PNode
	ParData
}

func (p *WhileNode) getChilds() []PNode { return p.childs }

type VarNode struct {
	name  string
	isRef bool
	ParData
}

func (p *VarNode) getChilds() []PNode { return []PNode{} }

type NumNode struct {
	num int
	ParData
}

func (p *NumNode) getChilds() []PNode { return []PNode{} }

type StrNode struct {
	childs []PNode
	str    string
	ParData
}

func (p *StrNode) getChilds() []PNode { return []PNode{} }

type CharNode struct {
	childs []PNode
	char   uint8
	ParData
}

func (p *CharNode) getChilds() []PNode { return []PNode{} }

/* TODO: struct */
// type StructNode struct{ Node }
// type DefNode struct{ Node ResultType }

/* TODO: seq */
