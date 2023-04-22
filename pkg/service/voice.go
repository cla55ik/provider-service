package service

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country        string  `json:"country"`
	Load           string  `json:"bandwidth"`
	ResponseTime   string  `json:"response_time"`
	Provider       string  `json:"provider"`
	Stability      float32 `json:"connection_stability"`
	TTFB           int     `json:"ttfb"`
	Purity         int     `json:"voice_purity"`
	MedianDuration int     `json:"median_of_call_time"`
}

// Indexes of voice data
const (
	CountryVoice = iota
	LoadVoice
	ResponseTimeVoice
	ProviderVoice
	StabilityVoice
	TtfbVoice
	PurityVoice
	MedianDurationVoice
)

func parseVoiceData(line string) (VoiceCallData, bool) {
	voice := strings.Split(line, ";")

	switch {
	case len(voice) != 8:
		fallthrough
	case !isValidCountry(voice[CountryVoice]):
		fallthrough
	case !isValidLoad(voice[LoadVoice]):
		fallthrough
	case !isValidResponseTime(voice[ResponseTimeVoice]):
		fallthrough
	case !isValidVoiceProvider(voice[ProviderVoice]):
		fallthrough
	case !isValidStability(voice[StabilityVoice]):
		fallthrough
	case !isValidPurity(voice[PurityVoice]):
		fallthrough
	case !isValidTTFB(voice[TtfbVoice]):
		fallthrough
	case !isMedianDuration(voice[MedianDurationVoice]):
		return VoiceCallData{}, false
	}

	load := voice[LoadVoice]
	responseTime := voice[ResponseTimeVoice]
	stability64, _ := strconv.ParseFloat(voice[StabilityVoice], 32)
	ttfb, _ := strconv.Atoi(voice[ResponseTimeVoice])
	purity, _ := strconv.Atoi(voice[PurityVoice])
	medianDuration, _ := strconv.Atoi(voice[MedianDurationVoice])

	return VoiceCallData{
		Country:        voice[CountryVoice],
		Load:           load,
		ResponseTime:   responseTime,
		Provider:       voice[ProviderVoice],
		Stability:      float32(stability64),
		TTFB:           ttfb,
		Purity:         purity,
		MedianDuration: medianDuration,
	}, true
}

func GetStatusVoice(csvPath string) ([]VoiceCallData, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return []VoiceCallData{}, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []VoiceCallData{}, err
	}

	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)
	VoiceList := make([]VoiceCallData, 0)

	for scanner.Scan() {
		line := scanner.Text()
		voice, ok := parseVoiceData(line)

		if ok {
			VoiceList = append(VoiceList, voice)
		}
	}

	return VoiceList, nil
}
