package service

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

const (
	CountrySms = iota
	BandwidthSms
	ResponseTimeSms
	ProviderSms
)

func parseSMS(line string) (SMSData, bool) {
	sms := strings.Split(line, ";")

	switch {
	case len(sms) < 4:
		fallthrough
	case !isValidCountry(sms[CountrySms]):
		fallthrough
	case !isValidBandwidth(sms[BandwidthSms]):
		fallthrough
	case !isValidResponseTime(sms[ResponseTimeSms]):
		fallthrough
	case !isValidSMSProvider(sms[ProviderSms]):
		return SMSData{}, false
	}

	return SMSData{
		Country:      sms[CountrySms],
		Bandwidth:    sms[BandwidthSms],
		ResponseTime: sms[ResponseTimeSms],
		Provider:     sms[ProviderSms],
	}, true
}

func GetStatusSMS(csvPath string) ([]SMSData, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return []SMSData{}, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []SMSData{}, err
	}

	err = file.Close()
	if err != nil {
		return []SMSData{}, err
	}

	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)
	SMSList := make([]SMSData, 0)

	for scanner.Scan() {
		line := scanner.Text()
		sms, ok := parseSMS(line)

		if ok {
			SMSList = append(SMSList, sms)
		}
	}
	return SMSList, nil
}
