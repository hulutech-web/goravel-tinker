// Code generated by 'yaegi extract text/scanner'. DO NOT EDIT.

//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/constant"
	"go/token"
	"reflect"
	"text/scanner"
)

func init() {
	Symbols["text/scanner/scanner"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Char":           reflect.ValueOf(constant.MakeFromLiteral("-5", token.INT, 0)),
		"Comment":        reflect.ValueOf(constant.MakeFromLiteral("-8", token.INT, 0)),
		"EOF":            reflect.ValueOf(constant.MakeFromLiteral("-1", token.INT, 0)),
		"Float":          reflect.ValueOf(constant.MakeFromLiteral("-4", token.INT, 0)),
		"GoTokens":       reflect.ValueOf(constant.MakeFromLiteral("1012", token.INT, 0)),
		"GoWhitespace":   reflect.ValueOf(constant.MakeFromLiteral("4294977024", token.INT, 0)),
		"Ident":          reflect.ValueOf(constant.MakeFromLiteral("-2", token.INT, 0)),
		"Int":            reflect.ValueOf(constant.MakeFromLiteral("-3", token.INT, 0)),
		"RawString":      reflect.ValueOf(constant.MakeFromLiteral("-7", token.INT, 0)),
		"ScanChars":      reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"ScanComments":   reflect.ValueOf(constant.MakeFromLiteral("256", token.INT, 0)),
		"ScanFloats":     reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"ScanIdents":     reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"ScanInts":       reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"ScanRawStrings": reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"ScanStrings":    reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"SkipComments":   reflect.ValueOf(constant.MakeFromLiteral("512", token.INT, 0)),
		"String":         reflect.ValueOf(constant.MakeFromLiteral("-6", token.INT, 0)),
		"TokenString":    reflect.ValueOf(scanner.TokenString),

		// type definitions
		"Position": reflect.ValueOf((*scanner.Position)(nil)),
		"Scanner":  reflect.ValueOf((*scanner.Scanner)(nil)),
	}
}