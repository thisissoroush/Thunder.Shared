package exception

import "fmt"

type ThunderException struct {
	Code    int
	Message string
}

func (e *ThunderException) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
