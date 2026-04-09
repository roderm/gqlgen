package graphql

import (
	"context"

	"github.com/vektah/gqlparser/v2/ast"
)

// ExecutableSchemaState stores generated executable schema dependencies.
// Generated code defines its local executableSchema type from this one.
type ExecutableSchemaState[R any, D any, C any] struct {
	SchemaData     *ast.Schema
	Resolvers      R
	Directives     D
	ComplexityRoot C

	Sources             []*ast.Source
	AfterInputUnmarshal []func(ctx context.Context, obj interface{}, value interface{}) error
}

func (e *ExecutableSchemaState[R, D, C]) AfterUnmarshalInput(ctx context.Context, obj interface{}, value interface{}) error {
	for _, fn := range e.AfterInputUnmarshal {
		if err := fn(ctx, obj, value); err != nil {
			return err
		}
	}
	return nil
}
