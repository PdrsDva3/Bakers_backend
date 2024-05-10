package customerr

import (
	"fmt"
)

var errorTypes = [9]string{
	"transaction error: %v",               // 0
	"json error: %v",                      // 1
	"query error: %v",                     // 2
	"rollback error: %v",                  // 3
	"query error: %v, rollback error: %v", // 4
	"commit error: %v",                    // 5
	"scan error: %v",                      // 6
	"row error : %v",                      // 7
	"row count doesnt equals 1: %v",       // 8
}

func ErrorMessage(mode int, a ...any) error {
	return fmt.Errorf(errorTypes[mode], a...)
}
