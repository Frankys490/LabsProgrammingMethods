package labs

import (
	"bufio"
	"os"
)

func scanFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var str string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str += scanner.Text()
		str += ";"
	}
	return str
}
