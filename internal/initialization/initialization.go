package initialization

import (
	"cube/internal/output"
	"time"
)

func Initialization() {

	output.Output("Wait, the die is being rolled...\n")
	time.Sleep(time.Second * 2)

}
