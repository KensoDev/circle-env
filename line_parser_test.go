package circleenv

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) { TestingT(t) }

type ParserSuite struct{}

var _ = Suite(&ParserSuite{})

func (s *ParserSuite) TestInitialParsing(c *C) {
	content, _ := ioutil.ReadFile("fixtures/single_line.env")
	envvar, err := ParseLine(content)
	c.Assert(err, Equals, nil)
	c.Assert(envvar.Key, Equals, "HOME")
	c.Assert(envvar.Value, Equals, "/home")
}
