// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// ---- lexer ----

// This lexer is similar to the one described in Chapter 13.
type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// ---- parser ----

// Parse parses the input string as an arithmetic expression.
//
//   expr = num                         a literal number, e.g., 3.14159
//        | id                          a variable name, e.g., x
//        | id '(' expr ',' ... ')'     a function call
//        | '-' expr                    a unary operator (+-)
//        | expr '+' expr               a binary operator (+-*/)
//
func Parse(input string) (_ Expr, _ []Var, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() // initial lookahead
	e, vars := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, removeDuplicate(vars), nil
}

func parseExpr(lex *lexer) (Expr, []Var) { return parseBinary(lex, nil, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, vars []Var, prec1 int) (Expr, []Var) {
	lhs, vars := parseUnary(lex, vars)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next() // consume operator
			var rhs Expr
			rhs, vars = parseBinary(lex, vars, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs, vars
}

// unary = '+' expr | primary
func parseUnary(lex *lexer, vars []Var) (Expr, []Var) {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next() // consume '+' or '-'
		e, vars := parseUnary(lex, vars)
		return unary{op, e}, vars
	}
	return parsePrimary(lex, vars)
}

// primary = id
//         | id '(' expr ',' ... ',' expr ')'
//         | num
//         | '(' expr ')'
func parsePrimary(lex *lexer, vars []Var) (Expr, []Var) {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next() // consume Ident
		if lex.token != '(' {
			v := Var(id)
			return v, append(vars, v)
		}
		lex.next() // consume '('
		var args []Expr
		if lex.token != ')' {
			for {
				e, v := parseExpr(lex)
				args = append(args, e)
				vars = append(vars, v...)
				if lex.token != ',' {
					break
				}
				lex.next() // consume ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %q, want ')'", lex.token)
				panic(lexPanic(msg))
			}
		}
		lex.next() // consume ')'
		return call{id, args}, vars

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // consume number
		return literal(f), vars

	case '(':
		lex.next() // consume ')'
		e, v := parseExpr(lex)
		vars = append(vars, v...)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e, vars
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}

func removeDuplicate(s []Var) []Var {
	setMap := make(map[Var]bool)
	var set []Var

	for _, v := range s {
		setMap[v] = true
	}

	for k, _ := range setMap {
		set = append(set, k)
	}

	return set
}
