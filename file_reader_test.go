package circleenv

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestFileReader(t *testing.T) { TestingT(t) }

type FileReaderSuite struct{}

var _ = Suite(&FileReaderSuite{})

func (s *FileReaderSuite) TestMultipleLinesRead(c *C) {
	filename := "fixtures/.env.test"
	fileReader := NewFileReader(filename)
	c.Assert(len(fileReader.Lines), Equals, 3)
}

func (s *FileReaderSuite) TestSingleLineRead(c *C) {
	filename := "fixtures/single_line.env"
	fileReader := NewFileReader(filename)
	c.Assert(len(fileReader.Lines), Equals, 1)
}

func (s *FileReaderSuite) TestEntries(c *C) {
	filename := "fixtures/.env.test"
	fileReader := NewFileReader(filename)
	c.Assert(len(fileReader.Lines), Equals, 3)
	c.Assert(len(fileReader.EnvVars), Equals, 3)
}
