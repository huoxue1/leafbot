package utils

import (
	"github.com/mxschmitt/playwright-go"
	log "github.com/sirupsen/logrus"
)

func GetPWScreen(url string, device string) ([]byte, error) {

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

	var userAgent string
	if device == "android" {
		userAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Mobile Safari/537.36 Edg/92.0.902.67"
	} else {
		userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.67"
	}

	page, err := browser.NewPage(playwright.BrowserNewContextOptions{UserAgent: playwright.String(userAgent)})

	defer func() {
		page.Close()
		pw.Stop()
	}()
	if err != nil {
		log.Errorf("could not create page: %v", err)
		return nil, err
	}
	if _, err = page.Goto(url); err != nil {
		log.Errorf("could not goto: %v", err)
		return nil, err
	}
	data, err := page.Screenshot(playwright.PageScreenshotOptions{
		FullPage: playwright.Bool(true),
	})
	return data, err

}
