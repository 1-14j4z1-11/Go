package main

import (
	"eval"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			getExpression(w, req)
		} else if req.Method == "POST" {
			postExpression(w, req)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:50000", nil))
}

type args struct {
	Expr  string
	Value string
}

var expressionTemplate = `<html>
<body>
<form name="expr" action="/" method="post">
<p>
Expression：<input type="text" name="expression" value="{{.Expr}}" size="40">
<input type="submit" value="Calculate">
</p>
<p>
Value = {{.Value}}
</p>
</form>
</body>
</html>
`

var exprForm = template.Must(template.New("table").Parse(expressionTemplate))

func getExpression(w http.ResponseWriter, req *http.Request) {
	exprForm.Execute(w, args{"", ""})
}

func postExpression(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Form parse error : %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Parse error"))
		return
	}

	exprStr := req.Form.Get("expression")
	expr, vars, err := eval.Parse(exprStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Expression parse error : %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Parse error"))
		return
	}
	if len(vars) != 0 {
		fmt.Fprintf(os.Stderr, "contains vars\n")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Expression containing vars is unsupported"))
		return
	}
	if err = expr.Check(nil); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid expression"))
		return
	}

	value := expr.Eval(nil)
	exprForm.Execute(w, args{exprStr, fmt.Sprintf("%f", value)})
}
