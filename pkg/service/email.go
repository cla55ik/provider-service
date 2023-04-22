package service

import (
	"bufio"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

const (
	CountryEmail = iota
	ProviderEmail
	DeliveryTimeEmail
)

func parseEmail(line string) (EmailData, bool) {
	email := strings.Split(line, ";")

	switch {
	case len(email) != 3:
		fallthrough
	case !isValidCountry(email[CountryEmail]):
		fallthrough
	case !isValidEmailProvider(email[ProviderEmail]):
		fallthrough
	case !isValidDeliveryTime(email[DeliveryTimeEmail]):
		return EmailData{}, false
	}

	deliveryTime, _ := strconv.Atoi(email[DeliveryTimeEmail])

	return EmailData{
		Country:      email[CountryEmail],
		Provider:     email[ProviderEmail],
		DeliveryTime: deliveryTime,
	}, true
}

func GetStatusEmail(csvPath string) ([]EmailData, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return []EmailData{}, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []EmailData{}, err
	}

	err = file.Close()
	if err != nil {
		return []EmailData{}, err
	}

	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)
	EmailList := make([]EmailData, 0)

	for scanner.Scan() {
		line := scanner.Text()
		email, ok := parseEmail(line)

		if ok {
			EmailList = append(EmailList, email)
		}
	}

	return EmailList, nil
}

func GetSlowFastEmailProvider(data []EmailData, code string) (slow []EmailData, fast []EmailData) {

	emailsByCountry := make([]EmailData, 0)
	for _, email := range data {
		if email.Country == code {
			emailsByCountry = append(emailsByCountry, email)
		}
	}

	sort.SliceStable(emailsByCountry, func(i, j int) bool {
		return emailsByCountry[i].DeliveryTime < emailsByCountry[j].DeliveryTime
	})

	if len(emailsByCountry) < 3 {
		return emailsByCountry, emailsByCountry
	}

	return emailsByCountry[len(emailsByCountry)-3:], emailsByCountry[:3]
}
