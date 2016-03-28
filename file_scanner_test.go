package circleenv

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestFileScanner(t *testing.T) { TestingT(t) }

type FileScannerSuite struct{}

var _ = Suite(&FileScannerSuite{})

func (s *FileScannerSuite) TestScanFileContent(c *C) {
	filename := "fixtures/run.sh.template"
	scanner := NewFileScanner(filename)
	scanner.Scan()
	c.Assert(len(scanner.Matches), Equals, 5)
	c.Assert(scanner.Matches[0].MatchString, Equals, "<TEST_CODE>")
	c.Assert(scanner.Matches[0].MatchEnvVariable, Equals, "TEST_CODE")
	c.Assert(scanner.Matches[1].MatchString, Equals, "<TEST_DB_NAME>")
	c.Assert(scanner.Matches[1].MatchEnvVariable, Equals, "TEST_DB_NAME")
}
