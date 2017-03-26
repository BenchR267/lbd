package ast

import "github.com/BenchR267/lbd/lexer/token"

// NodeType is the type of a given node
type NodeType int

const (
	Program NodeType = iota
)

// Node represents one node inside the AST
type Node struct {
	token token.Token
}
