package request

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testSuccess struct {
	UserID int `json:"userId"`
}
type testFalse struct {
	UserID string `json:"userId"`
}

func TestDoGetHTTPRequest(t *testing.T) {
	var (
		url     = `https://jsonplaceholder.typicode.com/posts/1`
		result  = testSuccess{}
		result1 = testFalse{}
		err     error
	)
	Convey("TestDoGetHTTPRequest testing process...", t, func() {
		Convey("Check for success testResult", func() {
			err = DoGetHTTPRequest(url, true, &result)
			So(err, ShouldBeNil)
			So(result.UserID, ShouldHaveSameTypeAs, 1)
		})
		Convey("Check for Invalid URL", func() {
			err = DoGetHTTPRequest("url", false, &result1)
			So(err, ShouldNotBeNil)
		})
		Convey("Check for Invalid request", func() {
			err = DoGetHTTPRequest(url+"DINGDONG", true, &result)
			So(err, ShouldNotBeNil)
		})
		Convey("Check for JSON Unmarshalling", func() {
			err = DoGetHTTPRequest(url, false, &result1)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestDoHTTPRequest(t *testing.T) {
	var (
		url     = `https://jsonplaceholder.typicode.com/posts/1`
		result  = testSuccess{}
		result1 = testFalse{}
		err     error
	)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	Convey("TestDoGetHTTPRequest testing process...", t, func() {
		Convey("Check for success testResult", func() {
			err = DoHTTPRequest("get", url, headers, strings.NewReader(""), true, &result)
			So(err, ShouldBeNil)
			So(result.UserID, ShouldHaveSameTypeAs, 1)
		})
		Convey("Check for Invalid URL", func() {
			err = DoHTTPRequest("get", "url", nil, strings.NewReader(""), false, &result)
			So(err, ShouldNotBeNil)
		})
		Convey("Check for JSON Unmarshalling", func() {
			err = DoHTTPRequest("get", url, headers, strings.NewReader(""), false, &result1)
			So(err, ShouldNotBeNil)
		})
	})
}
