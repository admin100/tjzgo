package tjzgo

import (
	"fmt"
	"os"
)

func Fatal(v ...interface{}) {
	fmt.Sprint(v...)
	os.Exit(1)
}
