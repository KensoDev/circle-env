package circleenv

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func TestFileFinder(t *testing.T) { TestingT(t) }

type FileFinderSuite struct{}

var _ = Suite(&FileFinderSuite{})

func (s *FileFinderSuite) TestFindAndReplace(c *C) {
	dir, _ := os.Getwd()
	f := NewFileFinder(dir)
	f.FindAndReplaceTemplateFiles()
}
