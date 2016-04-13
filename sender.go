package circleenv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sender struct {
	Config *Configuration
}

func (sender *Sender) Send() {
	reader := NewFileReader(sender.Config.FileName)

	for i := 0; i < len(reader.EnvVars); i++ {
		j, err := json.Marshal(reader.EnvVars[i])
		req, err := http.NewRequest("POST", sender.url(), bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Updated: ", string(body))

		if err != nil {
			panic(err)
		}
	}
}

func (sender *Sender) url() string {
	return fmt.Sprintf("https://circleci.com/api/v1/project/%v/%v/envvar?circle-token=%v", sender.Config.UserName, sender.Config.ProjectName, sender.Config.GetToken())
}

func NewSender(configuration *Configuration) *Sender {
	return &Sender{Config: configuration}
}
