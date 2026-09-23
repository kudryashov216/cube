package initialization

import (
	"bufio"
	"time"
)

func Initialization(writer *bufio.Writer) {

	writer.Write([]byte("Wait, the die is being rolled...\n"))
	writer.Flush()

	time.Sleep(time.Second * 2)

}
