package circleenv

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestReplacer(t *testing.T) { TestingT(t) }

type ReplacerSuite struct{}

var _ = Suite(&ReplacerSuite{})

func (s *ReplacerSuite) TestReplaceContentAndWriteFile(c *C) {
	os.Setenv("TEST_DB_NAME", "test_database")
	os.Setenv("TEST_DB_USERNAME", "test_db_username")

	os.Remove("fixtures/run.sh")
	filename := "fixtures/run.sh.template"

	scanner := NewFileScanner(filename).Scan()
	r := NewReplacer(scanner)
	r.ReplaceContentAndWriteFile()
	c.Assert(FileExists("fixtures/run.sh"), Equals, true)

	bytes, _ := ioutil.ReadFile("fixtures/run.sh")
	content := string(bytes)

	c.Assert(strings.Contains(content, "test_database"), Equals, true)
	c.Assert(strings.Contains(content, "test_db_username"), Equals, true)
}
