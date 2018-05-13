package deptwo // import "github.com/shenderov/dep-two"

import (
	"fmt"
	"github.com/shenderov/dep-one"
)

func Hello() {
	fmt.Println("This is dep two")
	depone.Hello()
}
