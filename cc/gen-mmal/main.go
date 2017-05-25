// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"go/token"
	"log"
	"os"
	"sort"
	"unsafe"

	"github.com/cznic/cc"
	"github.com/cznic/xc"
	"github.com/davecgh/go-spew/spew"
)

type TargetArch string

const (
	Arch32    TargetArch = "i386"
	Arch48    TargetArch = "x86_48"
	Arch64    TargetArch = "x86_64"
	ArchArm32 TargetArch = "arm"
	ArchArm64 TargetArch = "aarch64"
)

var builtinBase = `
#define __builtin_va_list void *
#define __asm(x)
#define __inline inline
#define __inline__ inline
#define __signed
#define __signed__
#define __const const
#define __extension__
#define __attribute__(x)
#define __restrict
#define __volatile__

#define __builtin_inff() (0)
#define __builtin_infl() (0)
#define __builtin_inf() (0)
#define __builtin_fabsf(x) (0)
#define __builtin_fabsl(x) (0)
#define __builtin_fabs(x) (0)

#define __INTRINSIC_PROLOG(name)
`

var basePredefines = `
#define __STDC_HOSTED__ 1
#define __STDC_VERSION__ 199901L
#define __STDC__ 1
#define __GNUC__ 4
#define __POSIX_C_DEPRECATED(ver)

#define __FLT_MIN__ 0
#define __DBL_MIN__ 0
#define __LDBL_MIN__ 0

void __GO__(char*, ...);
`

var archPredefines = map[TargetArch]string{
	Arch32: `#define __i386__ 1`,
	Arch48: `#define __x86_64__ 1`,
	Arch64: `#define __x86_64__ 1`,
	ArchArm32: `#define __ARM_EABI__ 1
#define __arm__ 1`,
	ArchArm64: `#define __ARM_EABI__ 1
#define __aarch64__ 1`,
	// TODO: https://sourceforge.net/p/predef/wiki/Architectures/
}

func main() {
	decls, err := Declarations(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(decls[1])
	for _, d := range decls {
		if d.Name == "mmal_buffer_header_acquire" {
			spew.Dump(d)
		}
	}
}

func model() *cc.Model {
	p := int(unsafe.Sizeof(uintptr(0)))
	i := int(unsafe.Sizeof(int(0)))
	return &cc.Model{
		Items: map[cc.Kind]cc.ModelItem{
			cc.Ptr:               {Size: p, Align: p, StructAlign: p},
			cc.UintPtr:           {Size: p, Align: p, StructAlign: p},
			cc.Void:              {Size: 0, Align: 1, StructAlign: 1},
			cc.Char:              {Size: 1, Align: 1, StructAlign: 1},
			cc.SChar:             {Size: 1, Align: 1, StructAlign: 1},
			cc.UChar:             {Size: 1, Align: 1, StructAlign: 1},
			cc.Short:             {Size: 2, Align: 2, StructAlign: 2},
			cc.UShort:            {Size: 2, Align: 2, StructAlign: 2},
			cc.Int:               {Size: 4, Align: 4, StructAlign: 4},
			cc.UInt:              {Size: 4, Align: 4, StructAlign: 4},
			cc.Long:              {Size: i, Align: i, StructAlign: i},
			cc.ULong:             {Size: i, Align: i, StructAlign: i},
			cc.LongLong:          {Size: 8, Align: 8, StructAlign: 8},
			cc.ULongLong:         {Size: 8, Align: 8, StructAlign: 8},
			cc.Float:             {Size: 4, Align: 4, StructAlign: 4},
			cc.Double:            {Size: 8, Align: 8, StructAlign: 8},
			cc.LongDouble:        {Size: 8, Align: 8, StructAlign: 8},
			cc.Bool:              {Size: 1, Align: 1, StructAlign: 1},
			cc.FloatComplex:      {Size: 8, Align: 8, StructAlign: 8},
			cc.DoubleComplex:     {Size: 16, Align: 16, StructAlign: 16},
			cc.LongDoubleComplex: {Size: 16, Align: 16, StructAlign: 16},
		},
	}
}

// Declaration is a description of a C function declaration.
type Declaration struct {
	Pos         token.Pos
	Name        string
	Return      cc.Type
	CParameters []cc.Parameter
	Variadic    bool
}

// Position returns the token position of the declaration.
func (d *Declaration) Position() token.Position { return xc.FileSet.Position(d.Pos) }

// Parameter is a C function parameter.
type Parameter struct{ Parameter cc.Parameter }

// Name returns the name of the parameter.
func (p *Parameter) Name() string { return string(xc.Dict.S(p.Parameter.Name)) }

// Type returns the C type of the parameter.
func (p *Parameter) Type() cc.Type { return p.Parameter.Type }

// Kind returns the C kind of the parameter.
func (p *Parameter) Kind() cc.Kind { return p.Parameter.Type.Kind() }

// Elem returns the pointer type of a pointer parameter or the element type of an
// array parameter.
func (p *Parameter) Elem() cc.Type { return p.Parameter.Type.Element() }

// Parameters returns the declaration's CParameters converted to a []Parameter.
func (d *Declaration) Parameters() []Parameter {
	p := make([]Parameter, len(d.CParameters))
	for i, c := range d.CParameters {
		p[i] = Parameter{c}
	}
	return p
}

// Declarations returns the C function declarations in the givel set of file paths.
func Declarations(paths []string) ([]Declaration, error) {
	_, _, sysIncludePaths, err := cc.HostConfig()
	archBits := Arch64
	predefined := builtinBase + basePredefines
	if archDefs, ok := archPredefines[archBits]; ok {
		predefined += fmt.Sprintf("\n%s", archDefs)
	}
	// model := models[cfg.archBits]
	t, err := cc.Parse(predefined, os.Args[1:], model(),
		cc.SysIncludePaths(sysIncludePaths),
		cc.EnableAlignOf(),
		cc.EnableAlternateKeywords(),
		cc.EnableAnonymousStructFields(),
		cc.EnableAsm(),
		cc.EnableBuiltinClassifyType(),
		cc.EnableBuiltinConstantP(),
		cc.EnableComputedGotos(),
		cc.EnableDefineOmitCommaBeforeDDD(),
		cc.EnableDlrInIdentifiers(),
		cc.EnableEmptyDeclarations(),
		cc.EnableEmptyDefine(),
		cc.EnableEmptyStructs(),
		cc.EnableImaginarySuffix(),
		cc.EnableImplicitFuncDef(),
		cc.EnableImplicitIntType(),
		cc.EnableIncludeNext(),
		cc.EnableLegacyDesignators(),
		cc.EnableNonConstStaticInitExpressions(),
		cc.EnableNoreturn(),
		cc.EnableOmitConditionalOperand(),
		cc.EnableOmitFuncArgTypes(),
		cc.EnableOmitFuncRetType(),
		cc.EnableParenthesizedCompoundStatemen(),
		cc.EnableStaticAssert(),
		cc.EnableTypeOf(),
		cc.EnableUndefExtraTokens(),
		cc.EnableUnsignedEnums(),
		cc.EnableWideBitFieldTypes(),
		cc.EnableWideEnumValues(),
	)
	if err != nil {
		return nil, fmt.Errorf("gen-mmal: failed to parse: %v", err)
	}

	var decls []Declaration
	for ; t != nil; t = t.TranslationUnit {
		if t.ExternalDeclaration.Case != 1 /* Declaration */ {
			log.Println(t.ExternalDeclaration.FunctionDefinition)
			continue
		}

		d := t.ExternalDeclaration.Declaration
		if d.Case != 0 {
			// Other case is 1: StaticAssertDeclaration.
			continue
		}

		init := d.InitDeclaratorListOpt
		if init == nil {
			continue
		}
		idl := init.InitDeclaratorList
		if idl.InitDeclaratorList != nil {
			// We do not want comma-separated lists.
			continue
		}
		id := idl.InitDeclarator
		if id.Case != 0 {
			// We do not want assignments.
			continue
		}

		declarator := id.Declarator
		if declarator.Type.Kind() != cc.Function {
			// We want only functions.
			continue
		}
		params, variadic := declarator.Type.Parameters()
		name, _ := declarator.Identifier()
		decls = append(decls, Declaration{
			Pos:         declarator.Pos(),
			Name:        string(xc.Dict.S(name)),
			Return:      declarator.Type.Result(),
			CParameters: params,
			Variadic:    variadic,
		})
	}

	sort.Sort(byPosition(decls))

	return decls, nil
}

type byPosition []Declaration

func (d byPosition) Len() int { return len(d) }
func (d byPosition) Less(i, j int) bool {
	iPos := d[i].Position()
	jPos := d[j].Position()
	if iPos.Filename == jPos.Filename {
		return iPos.Line < jPos.Line
	}
	return iPos.Filename < jPos.Filename
}
func (d byPosition) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
