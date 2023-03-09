package error

import "fmt"

type Error interface {
	GetFormatted() string
}

func ReportError(err Error) {
	fmt.Printf("%s\n", err.GetFormatted())
}
