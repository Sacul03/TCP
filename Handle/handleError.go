package Handle

import (
	"fmt"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
