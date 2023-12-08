package scraper

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	TARGET_SEPERATOR = "|"
)

type Target struct {
	Method string
	Url    string
}

func ConvertToTarget(inc string) Target {
	return Target{
		Method: strings.Split(inc, TARGET_SEPERATOR)[0],
		Url:    strings.Join(strings.Split(inc, TARGET_SEPERATOR)[1:], TARGET_SEPERATOR),
	}
}

func Scrape(
	output io.Writer,
	targets []Target,
) error {
	client := &http.Client{}

	for _, target := range targets {
		fmt.Fprintf(
			output,
			"Request : %s|%s\n",
			target.Method,
			target.Url,
		)

		req, _ := http.NewRequest(target.Method, target.Url, nil)
		resp, err := client.Do(req)

		if err != nil {
			return fmt.Errorf("error http request on %s/%s - %s", target.Method, target.Url, err)
		}

		fmt.Println(resp.Status)
	}

	return nil
}
