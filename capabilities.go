package goselenium

import (
	"encoding/json"
)

// Browser defines a supported selenium enabled browser.
type Browser interface {
	BrowserName() string
}

// Browser represents a browser to run within Selenium.
type browser struct {
	browserName string
}

// ChromeOptions defines chromedriver-specific options
type ChromeOptions interface {
	Binary() string
	SetBinary(binary string)
	Args() []string
	SetArgs(args []string)

	AsMap() map[string]interface{}
}

type chromeOptions struct {
	binary string
	args   []string
}

// Binary returns the executable chromedriver will launch when a new session is created.
func (co *chromeOptions) Binary() string {
	return co.binary
}

// SetBinary defines the executable chromedriver will launch when a new session is created.
func (co *chromeOptions) SetBinary(binary string) {
	co.binary = binary
}

// Args returns the command-line argument chromedriver will pass to the binary.
func (co *chromeOptions) Args() []string {
	return co.args
}

// SetArgs defines the command-line argument chromedriver will pass to the binary.
func (co *chromeOptions) SetArgs(args []string) {
	co.args = args
}

// AsMap returns a version of chrome options suitable for sending
// over to chromedriver over the wire.
func (co *chromeOptions) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"binary":      co.Binary(),
		"args":        co.Args(),
		"windowTypes": []string{"webview"},
	}
}

// BrowserName returns the browser name assigned to the current browser object.
func (b browser) BrowserName() string {
	return b.browserName
}

// FirefoxBrowser returns a Firefox browser object.
func FirefoxBrowser() Browser {
	return browser{"firefox"}
}

// ChromeBrowser returns a Chrome browser object.
func ChromeBrowser() Browser {
	return browser{"chrome"}
}

// AndroidBrowser returns an Android browser object.
func AndroidBrowser() Browser {
	return browser{"android"}
}

// HTMLUnitBrowser returns a HTMLUnit browser object.
func HTMLUnitBrowser() Browser {
	return browser{"htmlunit"}
}

// InternetExplorerBrowser returns an IE browser object.
func InternetExplorerBrowser() Browser {
	return browser{"internetexplorer"}
}

// IPhoneBrowser returns an IPhone browser object.
func IPhoneBrowser() Browser {
	return browser{"iphone"}
}

// IPadBrowser returns an IPad browser object.
func IPadBrowser() Browser {
	return browser{"ipad"}
}

// OperaBrowser returns an Opera browser object.
func OperaBrowser() Browser {
	return browser{"opera"}
}

// SafariBrowser returns a Safari browser object.
func SafariBrowser() Browser {
	return browser{"safari"}
}

// Capabilities represents the capabilities defined in the W3C specification.
// The main capability is the browser, which can be set by calling one of the
// \wBrowser\(\) methods.
type Capabilities struct {
	browser       Browser
	chromeOptions ChromeOptions
}

// Browser yields the browser capability assigned to the current Capabilities
// object..
func (c *Capabilities) Browser() Browser {
	if c.browser != nil {
		return c.browser
	}

	return browser{}
}

// SetBrowser sets the browser capability to be one of the allowed browsers.
func (c *Capabilities) SetBrowser(b Browser) {
	c.browser = b
}

// NewChromeOptions returns an empty set of ChromeOptions
func NewChromeOptions() ChromeOptions {
	return &chromeOptions{
		binary: "",
		args:   []string{},
	}
}

// SetChromeOptions sets the chrome options for the current Capabilities object.
func (c *Capabilities) SetChromeOptions(co ChromeOptions) {
	c.chromeOptions = co
}

func (c *Capabilities) toJSON() (string, error) {
	capabilities := map[string]map[string]interface{}{
		"desiredCapabilities": {
			"browserName": c.browser.BrowserName(),
		},
	}

	if c.chromeOptions != nil {
		capabilities["desiredCapabilities"]["chromeOptions"] = c.chromeOptions.AsMap()
	}

	capabilitiesJSON, err := json.Marshal(capabilities)
	if err != nil {
		return "", err
	}

	return string(capabilitiesJSON), nil
}
