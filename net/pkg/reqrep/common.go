package reqrep

import (
	"fmt"
	"os"
)

type Serializable interface {
	ToBytes() []byte
	FromBytes(data []byte)
}

func die(format string, v ...any) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}
