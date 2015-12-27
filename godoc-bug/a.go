package bug

import (
	"go/ast"
	. "go/ast"
)

func X() Node {
	return nil
}

func Y() ast.Node {
	return nil
}
