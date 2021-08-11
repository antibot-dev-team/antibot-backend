package analyzer

import (
	"strings"
)

const chromeDriverPrefix = "cdc_"

type ClientProperties struct {
	Languages       []string `json:"languages"`
	Plugins         []string `json:"plugins"`
	Window          []string `json:"custom_window"`
	UserAgent       string   `json:"ua"`
	HasWindowChrome bool     `json:"has_window_chrome"`
}

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

// AnalyzeProperties checks browser/client properties
// Returns false if properties are invalid, or true otherwise
func (a *Analyzer) AnalyzeProperties(properties ClientProperties) bool {
	// TODO: Add more checks here
	return a.analyzeLanguages(properties.Languages) &&
		a.analyzePlugins(properties.Plugins, properties.UserAgent) &&
		a.analyzeWindow(properties.Window) &&
		a.analyzeWindowChrome(properties.HasWindowChrome, properties.UserAgent)
}

// analyzeLanguages checks available languages
// If no languages available in browser - possibly bot
func (a *Analyzer) analyzeLanguages(languages []string) bool {
	return len(languages) != 0
}

// analyzePlugins checks installed plugins
// If no plugins available in browser - possibly bot
func (a *Analyzer) analyzePlugins(plugins []string, ua string) bool {

	// In Firefox-like browsers no plugins can be normal
	isGecko := strings.Contains(ua, "Gecko/")

	return isGecko || len(plugins) != 0
}

// analyzeWindow checks window environment
// If chromedriver prefix is present in window - possibly bot
// TODO: Add additional window checks - bot properties, bad client's properties, malicious objects, etc.
func (a *Analyzer) analyzeWindow(window []string) bool {
	for _, element := range window {
		if strings.Contains(element, chromeDriverPrefix) {
			return false
		}
	}

	return true
}

// analyzeWindowChrome checks if window.chrome is not present but client's browser is Chrome, Chromium or Opera
// While window.chrome is available in vanilla mode, it’s not available in headless mode.
func (a *Analyzer) analyzeWindowChrome(hasWindowChrome bool, ua string) bool {
	// TODO: check if window.chrome is present in Chrome on iOS. If so, check UA for "CriOS" substring.

	isChromeOrOpera := strings.Contains(ua, "Chrome") || strings.Contains(ua, "Chromium")
	isChromeOrOpera = isChromeOrOpera || strings.Contains(ua, "Opera") || strings.Contains(ua, "OPR")

	return !isChromeOrOpera || hasWindowChrome
}