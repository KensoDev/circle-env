package circleenv

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestConfiguration(t *testing.T) { TestingT(t) }

type ConfigurationSuite struct{}

var _ = Suite(&ConfigurationSuite{})

func (s *ConfigurationSuite) TestConfiguratrionToken(c *C) {
	configuration := &Configuration{
		Token: "some-token",
	}
	c.Assert(configuration.GetToken(), Equals, "some-token")
}
