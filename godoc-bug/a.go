package bug

import (
	"go/ast"
	. "go/ast"
)

func X(Node) Node {
	return nil
}

func Y(Node) ast.Node {
	return nil
}
