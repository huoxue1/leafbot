package utils

import (
	"github.com/mxschmitt/playwright-go"
	log "github.com/sirupsen/logrus"
)

func GetPWScreen(url string) ([]byte, error) {

	pw, err := playwright.Run()
	if err != nil {
		log.Errorf("could not start playwright: %v", err)
		return nil, err
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Errorf("could not launch browser: %v", err)
		return nil, err
	}

	page, err := browser.NewPage()

	defer func() {
		page.Close()
		pw.Stop()
	}()
	if err != nil {
		log.Errorf("could not create page: %v", err)
		return nil, err
	}
	if _, err = page.Goto(url); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	data, err := page.Screenshot(playwright.PageScreenshotOptions{
		FullPage: playwright.Bool(true),
	})
	return data, err

}
