package integrationtests

import (
	"testing"
	"time"

	goselenium "github.com/itchio/go-selenium"
)

func Test_ElementWaitUntilElementPresent_CanSucceed(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error creating session", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/element-present.html")
	if err != nil {
		errorAndWrap(t, "Error navigating to URL", err)
	}

	by := goselenium.ByCSSSelector("#not-present-div")
	resp := driver.Wait(goselenium.UntilElementPresent(by), 20*time.Second, 0)
	if !resp {
		errorAndWrap(t, "Error waiting for element to be visible or it timed out", err)
	}

	printObjectResult(resp)
}

func Test_ElementWaitUntilElementPresent_NotFoundPriorToTimeoutFails(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error creating session", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/element-present.html")
	if err != nil {
		errorAndWrap(t, "Error navigating to URL", err)
	}

	by := goselenium.ByCSSSelector("#not-present-div")
	resp := driver.Wait(goselenium.UntilElementPresent(by), 1*time.Second, 0)
	if resp {
		errorAndWrap(t, "Error was not thrown when it should have been", err)
	}

	printObjectResult(resp)
}
