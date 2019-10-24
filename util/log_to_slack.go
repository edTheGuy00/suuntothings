package util

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gobuffalo/envy"
)

// LogToSlack sends logs to slack channel
func LogToSlack(content string) {
	_, err := http.PostForm("https://slack.com/api/files.upload", url.Values{
		"token":           {envy.Get("SLACK_TOKEN", "")},
		"channels":        {envy.Get("SLACK_LOG_CHANNEL", "")},
		"content":         {content},
		"filename":        {"suunto.json"},
		"filetype":        {"json"},
		"initial_comment": {"new file"},
	})

	if err != nil {
		log.Fatalln(err)
	}
}
