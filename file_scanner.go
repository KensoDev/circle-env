package circleenv

import (
	"io/ioutil"
	"os"
	"regexp"
)

type FileScanner struct {
	FileName    string
	FileContent string
	Matches     []Match
}

func NewFileScanner(filename string) *FileScanner {
	return &FileScanner{
		FileName: filename,
	}
}

func (r *FileScanner) Scan() *FileScanner {
	r.FileContent = r.ReadFile()
	re := regexp.MustCompile("<(.*?)>")
	matches := re.FindAllStringSubmatch(r.FileContent, -1)
	for _, match := range matches {
		m := Match{MatchString: match[0], MatchEnvVariable: match[1]}
		r.Matches = append(r.Matches, m)
	}
	return r
}

func (r *FileScanner) ReadFile() string {
	bytes, err := ioutil.ReadFile(r.FileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
