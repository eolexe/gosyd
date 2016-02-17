package pretty

import (
	"fmt"
)

func Println(a ...interface{}) (n int, errno error) {
	return fmt.Println("I AM FROM VENDOR >:]")
}
