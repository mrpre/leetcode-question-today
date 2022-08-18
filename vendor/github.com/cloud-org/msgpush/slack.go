package msgpush

import (
	"fmt"
	"net/http"	
	"encoding/json"
	"io/ioutil"
	//"github.com/imroc/req"
	"bytes"
)

type Content struct {
	Content string `json:"content"`
}

type Ats struct {
	IsAtAll bool `json:"isAtAll"`
}

type slackTextContent struct {
	At Ats `json:"at"`
	Msgtype string `json:"msgtype"`
	Text Content `json:"text"`
}

type Slack struct {
	ReqUrl string
	Pre    string
}

func NewSlack(token string) *Slack {
	return &Slack{ReqUrl: token}
}

func (s *Slack) Send(content string) error {
	return s.sendText(content)
}

/*
func (s *Slack) SendMrDown(title, content string) error {
	msg := fmt.Sprintf(`{
	"blocks": [
		{
			"type": "header",
			"text": {
				"type": "plain_text",
				"text": "%s",
				"emoji": true
			}
		},
		{
			"type": "section",
			"text": {
				"type": "plain_text",
				"text": "%s",
				"emoji": true
			}
		}
	]
}`, title, content)
	_, err := req.Post(s.ReqUrl, req.BodyJSON(msg))
	return err
}

func (s *Slack) SendMrDown(title, content string) error {
        msg := fmt.Sprintf(`{
            {
                    "msgtype": "markdown",
                    "markdown": {
                            "title": "%s",
                            "text": "%s",
                    },
                    "at": {
                    	"isAtAll": true,
                    },
            },
}`, title, content)
        _, err := req.Post(s.ReqUrl, req.BodyJSON(msg))
        return err
}
*/

func (s *Slack) sendText(content string) error {
	customreq := &slackTextContent{
		Text: Content{
			Content: content,
		},
		Msgtype: "text",
		At: Ats{IsAtAll: true},
	}
	reqb, _:=json.Marshal(customreq)
	creq, err := http.NewRequest("POST", s.ReqUrl, bytes.NewBuffer(reqb))
	creq.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(creq)
	if err !=nil {
		return err
	}
	defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return err
}

func (s *Slack) String() string {
	return "slack"
}
