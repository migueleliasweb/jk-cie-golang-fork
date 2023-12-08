package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

const (
	TARGET_SEPERATOR = "|"
)

type Target struct {
	Method string
	Url    string
}

func convertToTarget(inc string) Target {
	return Target{
		Method: strings.Split(inc, TARGET_SEPERATOR)[0],
		Url:    strings.Join(strings.Split(inc, TARGET_SEPERATOR)[1:], TARGET_SEPERATOR),
	}
}

func Scrape(cmd *cobra.Command, args []string) error {

	inc, err := cmd.Flags().GetStringSlice("target")

	if err != nil {
		return fmt.Errorf("error loading targets, %s", err)
	}

	targets := []Target{}

	for _, v := range inc {
		targets = append(targets, convertToTarget(v))
	}

	client := &http.Client{}

	for _, target := range targets {
		fmt.Printf("Request : %s/%s\n", target.Method, target.Url)

		req, err := http.NewRequest(target.Method, target.Url, nil)
		resp, err := client.Do(req)

		if err != nil {
			return fmt.Errorf("error http request on %s/%s - %s", target.Method, target.Url, err)
		}

		fmt.Println(resp.Status)
	}

	return nil
}
