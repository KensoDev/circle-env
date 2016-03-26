package circleenv

import (
	"bufio"
	"log"
	"os"
)

type FileReader struct {
	Lines   []string
	EnvVars []*EnvVar
}

func NewFileReader(filename string) (reader FileReader) {
	reader = FileReader{
		Lines: []string{},
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		reader.Lines = append(reader.Lines, line)
		envvar, err := ParseLine([]byte(line))

		if err == nil {
			reader.EnvVars = append(reader.EnvVars, envvar)
		}

	}

	return
}
