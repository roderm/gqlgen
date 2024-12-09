package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.58

import (
	"context"

	"github.com/99designs/gqlgen/_examples/enum/model"
)

// IntTyped is the resolver for the intTyped field.
func (r *queryResolver) IntTyped(ctx context.Context, arg model.IntTyped) (model.IntTyped, error) {
	return arg, nil
}

// IntUntyped is the resolver for the intUntyped field.
func (r *queryResolver) IntUntyped(ctx context.Context, arg int) (int, error) {
	return arg, nil
}

// IntTypedN is the resolver for the intTypedN field.
func (r *queryResolver) IntTypedN(ctx context.Context, arg *model.IntTyped) (*model.IntTyped, error) {
	return arg, nil
}

// IntUntypedN is the resolver for the intUntypedN field.
func (r *queryResolver) IntUntypedN(ctx context.Context, arg *int) (*int, error) {
	return arg, nil
}

// StringTyped is the resolver for the stringTyped field.
func (r *queryResolver) StringTyped(ctx context.Context, arg model.StringTyped) (model.StringTyped, error) {
	return arg, nil
}

// StringUntyped is the resolver for the stringUntyped field.
func (r *queryResolver) StringUntyped(ctx context.Context, arg string) (string, error) {
	return arg, nil
}

// StringTypedN is the resolver for the stringTypedN field.
func (r *queryResolver) StringTypedN(ctx context.Context, arg *model.StringTyped) (*model.StringTyped, error) {
	return arg, nil
}

// StringUntypedN is the resolver for the stringUntypedN field.
func (r *queryResolver) StringUntypedN(ctx context.Context, arg *string) (*string, error) {
	return arg, nil
}

// BoolTyped is the resolver for the boolTyped field.
func (r *queryResolver) BoolTyped(ctx context.Context, arg model.BoolTyped) (model.BoolTyped, error) {
	return arg, nil
}

// BoolUntyped is the resolver for the boolUntyped field.
func (r *queryResolver) BoolUntyped(ctx context.Context, arg bool) (bool, error) {
	return arg, nil
}

// BoolTypedN is the resolver for the boolTypedN field.
func (r *queryResolver) BoolTypedN(ctx context.Context, arg *model.BoolTyped) (*model.BoolTyped, error) {
	return arg, nil
}

// BoolUntypedN is the resolver for the boolUntypedN field.
func (r *queryResolver) BoolUntypedN(ctx context.Context, arg *bool) (*bool, error) {
	return arg, nil
}

// VarTyped is the resolver for the varTyped field.
func (r *queryResolver) VarTyped(ctx context.Context, arg model.VarTyped) (model.VarTyped, error) {
	return arg, nil
}

// VarUntyped is the resolver for the varUntyped field.
func (r *queryResolver) VarUntyped(ctx context.Context, arg bool) (bool, error) {
	return arg, nil
}

// InPackage is the resolver for the inPackage field.
func (r *queryResolver) InPackage(ctx context.Context, arg InPackage) (InPackage, error) {
	return arg, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
