package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func processMessage(m message) error {
	fmt.Printf("Size: %v, Date: %v, Snippet: %q\n", m.size, m.date, m.snippet)
	log.Println("Extract validation url")
	url, err := extractValidationUrl(strings.NewReader(m.body))
	if err != nil {
		return fmt.Errorf("Url not found in message %v: %v", m.gmailID, m)
	}

	log.Println("Follow redirect url")
	content, err := followUrl(url)
	if err != nil {
		return err
	}
	log.Println("Redirect url OK")

	log.Println("Extract redirect url")
	redirectUrl, err := extractRedirectUrl(strings.NewReader(string(content)))
	if err != nil {
		return fmt.Errorf("Redirect url not found in response from validation link: %s", string(content))
	}

	content2, err := followUrl(redirectUrl)
	if err != nil {
		return err
	}

	players, err := parseAttendeePage(strings.NewReader(string(content2)))
	if err != nil {
		return fmt.Errorf("error parsing attendee page: %v", err)
	}

	if len(players) > 0 {
		fmt.Println(players)
		for _, player := range players {
			matched, _ := regexp.MatchString(`Tupangiu`, player)
			if matched {
				fmt.Println("You have been selected !!!!")
			}
		}
	}
	return nil
}
