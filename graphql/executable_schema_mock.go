// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package graphql

import (
	"context"
	"errors"
	"sync"

	"github.com/vektah/gqlparser/v2/ast"
)

// Ensure, that ExecutableSchemaMock does implement ExecutableSchema.
// If this is not the case, regenerate this file with moq.
var _ ExecutableSchema = &ExecutableSchemaMock{}

// ExecutableSchemaMock is a mock implementation of ExecutableSchema.
//
// 	func TestSomethingThatUsesExecutableSchema(t *testing.T) {
//
// 		// make and configure a mocked ExecutableSchema
// 		mockedExecutableSchema := &ExecutableSchemaMock{
// 			ComplexityFunc: func(typeName string, fieldName string, childComplexity int, args map[string]interface{}) (int, bool) {
// 				panic("mock out the Complexity method")
// 			},
// 			ExecFunc: func(ctx context.Context) ResponseHandler {
// 				panic("mock out the Exec method")
// 			},
// 			SchemaFunc: func() *ast.Schema {
// 				panic("mock out the Schema method")
// 			},
// 		}
//
// 		// use mockedExecutableSchema in code that requires ExecutableSchema
// 		// and then make assertions.
//
// 	}
type ExecutableSchemaMock struct {
	// ComplexityFunc mocks the Complexity method.
	ComplexityFunc func(typeName string, fieldName string, childComplexity int, args map[string]interface{}) (int, bool)

	// ExecFunc mocks the Exec method.
	ExecFunc func(ctx context.Context) ResponseHandler

	// SchemaFunc mocks the Schema method.
	SchemaFunc func() *ast.Schema

	// calls tracks calls to the methods.
	calls struct {
		// Complexity holds details about calls to the Complexity method.
		Complexity []struct {
			// TypeName is the typeName argument value.
			TypeName string
			// FieldName is the fieldName argument value.
			FieldName string
			// ChildComplexity is the childComplexity argument value.
			ChildComplexity int
			// Args is the args argument value.
			Args map[string]interface{}
		}
		// Exec holds details about calls to the Exec method.
		Exec []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Schema holds details about calls to the Schema method.
		Schema []struct {
		}
	}
	lockComplexity sync.RWMutex
	lockExec       sync.RWMutex
	lockSchema     sync.RWMutex
}

// Complexity calls ComplexityFunc.
func (mock *ExecutableSchemaMock) Complexity(typeName string, fieldName string, childComplexity int, args map[string]interface{}) (int, bool) {
	if mock.ComplexityFunc == nil {
		panic("ExecutableSchemaMock.ComplexityFunc: method is nil but ExecutableSchema.Complexity was just called")
	}
	callInfo := struct {
		TypeName        string
		FieldName       string
		ChildComplexity int
		Args            map[string]interface{}
	}{
		TypeName:        typeName,
		FieldName:       fieldName,
		ChildComplexity: childComplexity,
		Args:            args,
	}
	mock.lockComplexity.Lock()
	mock.calls.Complexity = append(mock.calls.Complexity, callInfo)
	mock.lockComplexity.Unlock()
	return mock.ComplexityFunc(typeName, fieldName, childComplexity, args)
}

// ComplexityCalls gets all the calls that were made to Complexity.
// Check the length with:
//     len(mockedExecutableSchema.ComplexityCalls())
func (mock *ExecutableSchemaMock) ComplexityCalls() []struct {
	TypeName        string
	FieldName       string
	ChildComplexity int
	Args            map[string]interface{}
} {
	var calls []struct {
		TypeName        string
		FieldName       string
		ChildComplexity int
		Args            map[string]interface{}
	}
	mock.lockComplexity.RLock()
	calls = mock.calls.Complexity
	mock.lockComplexity.RUnlock()
	return calls
}

// Exec calls ExecFunc.
func (mock *ExecutableSchemaMock) Exec(ctx context.Context) ResponseHandler {
	if mock.ExecFunc == nil {
		panic("ExecutableSchemaMock.ExecFunc: method is nil but ExecutableSchema.Exec was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockExec.Lock()
	mock.calls.Exec = append(mock.calls.Exec, callInfo)
	mock.lockExec.Unlock()
	return mock.ExecFunc(ctx)
}

// ExecCalls gets all the calls that were made to Exec.
// Check the length with:
//     len(mockedExecutableSchema.ExecCalls())
func (mock *ExecutableSchemaMock) ExecCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockExec.RLock()
	calls = mock.calls.Exec
	mock.lockExec.RUnlock()
	return calls
}

// Schema calls SchemaFunc.
func (mock *ExecutableSchemaMock) Schema() *ast.Schema {
	if mock.SchemaFunc == nil {
		panic("ExecutableSchemaMock.SchemaFunc: method is nil but ExecutableSchema.Schema was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSchema.Lock()
	mock.calls.Schema = append(mock.calls.Schema, callInfo)
	mock.lockSchema.Unlock()
	return mock.SchemaFunc()
}

func (mock *ExecutableSchemaMock) Extend(ExecutableSchema) error {
	return errors.New("not implemented")
}
func (mock *ExecutableSchemaMock) SchemaSources() []*ast.Source {
	return []*ast.Source{}
}
// SchemaCalls gets all the calls that were made to Schema.
// Check the length with:
//     len(mockedExecutableSchema.SchemaCalls())
func (mock *ExecutableSchemaMock) SchemaCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSchema.RLock()
	calls = mock.calls.Schema
	mock.lockSchema.RUnlock()
	return calls
}
func (mock *ExecutableSchemaMock) GetRootResolvers() SchemaRootResolvers {
	return SchemaRootResolvers{}
}
func (mock *ExecutableSchemaMock) GetFieldResolvers() []map[string]func(context.Context, CollectedField, interface{}) (func(*FieldSet, int) bool, error) {
	return nil
}
