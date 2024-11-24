package genome

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/zygote-sh/zygote/internal/template"
)

func readGenomeFile() (*Templates, error) {

	// get file from viper
	// file := viper.GetString("genome-file")
	// TODO: Move to Filesystem package
	file := os.ExpandEnv("${ZDOTDIR}/zygote.genome")
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	var s Templates
	for fileScanner.Scan() {
		s = append(s, Template(fileScanner.Text()))
	}
	readFile.Close()
	return &s, nil
}
