package circleenv

import (
	"io/ioutil"
	"os"
	"strings"
)

type Replacer struct {
	Scanner *FileScanner
}

func NewReplacer(scanner *FileScanner) *Replacer {
	return &Replacer{
		Scanner: scanner,
	}
}

func (r *Replacer) ReplaceContentAndWriteFile() {

	newfilename := strings.Replace(r.Scanner.FileName, ".template", "", -1)

	for _, match := range r.Scanner.Matches {
		envValue := os.Getenv(match.MatchEnvVariable)

		if envValue != "" {
			r.Scanner.FileContent = strings.Replace(r.Scanner.FileContent,
				match.MatchString, envValue, -1)
		}
	}

	err := ioutil.WriteFile(newfilename, []byte(r.Scanner.FileContent), 0644)
	if err != nil {
		panic(err)
	}
}
