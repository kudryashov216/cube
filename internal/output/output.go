package output

import (
	"bufio"
	"os"
)

func Output(line string) {

	var writer *bufio.Writer

	writer = bufio.NewWriter(os.Stdout)

	writer.Write([]byte(line))
	writer.Flush()

}
