package analyzer

import "strings"

const chromeDriverPrefix = "cdc_"

type ClientProperties struct {
	Languages []string `json:"languages"`
	Plugins   []string `json:"plugins"`
	Window    []string `json:"custom_window"`
	Webdriver bool     `json:"webdriver"`
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
		a.analyzePlugins(properties.Plugins) &&
		a.analyzeWindow(properties.Window) &&
		a.analyzeWebdriver(properties.Webdriver)
}

// analyzeWebdriver checks navigator.webdriver property value
// If navigator.webdriver == True - possibly bot
func (a *Analyzer) analyzeWebdriver(webdriver bool) bool {
	return !webdriver
}

// analyzeLanguages checks available languages
// If no languages available in browser - possibly bot
func (a *Analyzer) analyzeLanguages(languages []string) bool {
	return len(languages) != 0
}

// analyzePlugins checks installed plugins
// If no plugins available in browser - possibly bot
// TODO: Check additionally for firefox browser: in firefox no plugins can be normal
func (a *Analyzer) analyzePlugins(plugins []string) bool {
	return len(plugins) != 0
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
