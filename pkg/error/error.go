package error

import (
	"fmt"
	"runtime"
)

// Error ...
func Error(err interface{}) error {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("ERROR: [%s:%d] %v \n", fn, line, err)

		switch err.(type) {
		case string:
			return fmt.Errorf(err.(string))
		case error:
			return err.(error)
		default:
			return fmt.Errorf("%v", err)
		}
	}
	return nil
}
