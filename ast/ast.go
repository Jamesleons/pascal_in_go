package ast

import (
	"fmt"
	"pascal_in_go/token"
)

//BinNode represents the binary expr
type BinNode struct {
	Left  Expr
	Right Expr
	Tok   token.Token
}

//NumNode holds the number of token
type NumNode struct {
	Tok   token.Token `json:"tok"`
	Value string      `json:"value"`
}

//VarNode represents the Variable node
type VarNode struct {
	Tok     token.Token
	Literal string
}

//Expr interface represent the expr,  expr is the unit for program
type Expr interface {
	ToStr() string
}

//Unary Node
type Unary struct {
	Op   string
	Expr Expr
}

//ToStr for BinNode
func (binNode BinNode) ToStr() string {
	left := fmt.Sprint(binNode.Left)
	right := fmt.Sprint(binNode.Right)
	op := fmt.Sprint(binNode.Tok)
	return fmt.Sprint(left, right, op)
}

//ToStr for NumNode
func (numNode NumNode) ToStr() string {
	return fmt.Sprint(numNode.Tok)
}

//ToStr for UnaryNode
func (unary Unary) ToStr() string {
	return fmt.Sprint(unary)
}

//ToStr for VarNode
func (va VarNode) ToStr() string {
	return va.Literal
}

type AssignStatement struct {
	Left  VarNode
	Op    token.Token
	Right Expr
}

func (asg AssignStatement) ToStr() string {
	return fmt.Sprint(asg)
}

type Compound struct {
	Children []Expr
}

func (compound Compound) ToStr() string {
	return fmt.Sprint(compound)
}

type Statement struct {
	Statement Expr
}

func (st Statement) ToStr() string {
	return fmt.Sprint(st)
}

type NoOp struct {
}

func (noop NoOp) ToStr() string {
	return "noop"
}

type Program struct {
	Block Block
	Name  string
}

func (prog Program) ToStr() string {
	return fmt.Sprint(prog)
}

type Block struct {
	Decl     Decl
	Compound Compound
}

func (blk Block) ToStr() string {
	return fmt.Sprint(blk)
}

type Decl struct {
	VarDeclList   []VarDecl
	ProceDeclList []Procedure
}

func (decl Decl) ToStr() string {
	return fmt.Sprint(decl)
}

type VarDecl struct {
	Node VarNode
	Type token.Type
}

func (vardecl VarDecl) ToStr() string {
	return fmt.Sprint(vardecl)
}

type Procedure struct {
	Name  string
	Block Block
}

func (procedure Procedure) ToStr() string {
	return fmt.Sprint(procedure)
}
