package validator

import "fmt"

func actualMessage(v interface{}, name string) string {
	return fmt.Sprintf("%s (%v)", name, v)
}
