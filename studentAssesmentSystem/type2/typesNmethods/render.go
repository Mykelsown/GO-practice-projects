package type2

import (
	"fmt"
	"os"
	"text/template"
)

type DisplayData struct {
	Receipent string
	Course string
	Score float64
	Grade string
}

func render(data any) {
	message := `
══════════════════════════════════════
   🎓  CERTIFICATE OF COMPLETION
══════════════════════════════════════
  Recipient : {{ .Receipent}}
  Course    : {{ .Course}}
  Score     : {{ .Score}}%
  Grade     : {{ .Grade}}

    Congratulations on completing the course!
══════════════════════════════════════
`
	tmpl, err := template.New("student-assesment-system").Parse(message) 
	logError(err)
	err = tmpl.Execute(os.Stdout, data)
	logError(err)
}

func logError(err error) {
  if err != nil {
	fmt.Errorf(err.Error())
	os.Exit(1)
  }
} 