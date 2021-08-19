package analyzer

import (
	"strings"

	"github.com/mileusna/useragent"
)

const chromeDriverPrefix = "cdc_"

type ClientProperties struct {
	Languages       []string `json:"languages"`
	Plugins         []string `json:"plugins"`
	Window          []string `json:"custom_window"`
	UserAgent       string   `json:"ua"`
	HasWindowChrome bool     `json:"has_window_chrome"`
	Webdriver       bool     `json:"webdriver"`
	ConsistentPerms bool     `json:"consistent_permissions"`
	EvalLength      int      `json:"eval_length"`
	ProductSub      string   `json:"product_sub"`
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
		a.analyzeWindowChrome(properties.HasWindowChrome, properties.UserAgent) &&
		a.analyzeWebdriver(properties.Webdriver) &&
		a.analyzePermissions(properties.ConsistentPerms) &&
		a.analyzeEvalLength(properties.UserAgent, properties.EvalLength) &&
		a.analyzeProductSub(properties.UserAgent, properties.ProductSub)
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

	ua = strings.ToLower(ua)

	isChromeOrOpera := strings.Contains(ua, "chrome") || strings.Contains(ua, "chromium")
	isChromeOrOpera = isChromeOrOpera || strings.Contains(ua, "opera") || strings.Contains(ua, "opr")

	return !isChromeOrOpera || hasWindowChrome
}

// analyzePermissions checks if permissions are working as intended.
// If permissions query leads to contradictory results - possibly bot.
func (a *Analyzer) analyzePermissions(consistentPerms bool) bool {
	return consistentPerms
}

// analyzeEvalLength checks if browser specified in UserAgent is consistent with value of eval.toString().length
// If it is inconsistent - possibly dishonest client.
func (a *Analyzer) analyzeEvalLength(UserAgent string, evalLength int) bool {
	// Browser -> evalLength
	BrowserToLength := map[string]int{
		"Firefox":           37,
		"Safari":            37,
		"Chrome":            33,
		"Opera":             33,
		"Internet Explorer": 39,
	}

	client := ua.Parse(UserAgent)
	usualLength, ok := BrowserToLength[client.Name]

	// If evalLength for given browser is unknown - consider client honest
	if !ok {
		return true
	}

	if usualLength != evalLength {
		return false
	}

	return true
}

// analyzeProductSub checks if browser specified in UserAgent is consistent with navigator.productSub
// If it is inconsistent - possibly dishonest client.
func (a *Analyzer) analyzeProductSub(UserAgent, productSub string) bool {
	// Safari, Chrome and Opera always have this productSub
	chromeBuildNumber := "20030107"

	client := ua.Parse(UserAgent)
	browser := client.Name

	if browser == "Opera" || browser == "Chrome" || browser == "Safari" {
		if productSub != chromeBuildNumber {
			return false
		}
	}

	return true
}