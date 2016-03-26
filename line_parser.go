package circleenv

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	assignRegex = regexp.MustCompile(`^([^=]+)=(.*)$`)
)

type LineParseError struct {
	Message string
}

func (e *LineParseError) Error() string {
	return e.Message
}

func ParseLine(content []byte) (*EnvVar, error) {
	s := string(content[:])
	s = strings.TrimSpace(s)

	groups := assignRegex.FindStringSubmatch(s)

	if groups != nil {
		envvar := &EnvVar{
			Key:   groups[1],
			Value: groups[2],
		}
		return envvar, nil
	}
	err := &LineParseError{
		fmt.Sprintf("Could not parse the line: %v", s),
	}
	return nil, err
}
