package circleenv

type EnvVar struct {
	Key   string `json:"name"`
	Value string `json:"value"`
}

type Configuration struct {
	FileName    string
	ProjectName string
	UserName    string
	Token       string
}

type Match struct {
	MatchString      string
	MatchEnvVariable string
}

func (c *Configuration) GetToken() string {
	return c.Token
}
