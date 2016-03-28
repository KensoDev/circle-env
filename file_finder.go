package circleenv

import (
	"os"
	"path/filepath"
	"strings"
)

type FileFinder struct {
	Path string
}

func NewFileFinder(path string) *FileFinder {
	return &FileFinder{Path: path}
}

func (finder *FileFinder) FindAndReplaceTemplateFiles() {
	filepath.Walk(finder.Path, finder.computeFile)
}

func (finder *FileFinder) computeFile(path string, f os.FileInfo, err error) (e error) {
	if strings.HasSuffix(f.Name(), ".template") {
		base := filepath.Base(path)
		dir := filepath.Dir(path)
		fullpath := filepath.Join(dir, base)
		scanner := NewFileScanner(fullpath).Scan()
		replacer := NewReplacer(scanner)
		replacer.ReplaceContentAndWriteFile()
	}
	return
}
