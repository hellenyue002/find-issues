package fakes

import "net/http"

type HttpClient struct {
	GetCall struct {
		CallCount int
		Receives  struct {
			Url string
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
	}
}

func (h *HttpClient) Get(url string) (*http.Response, error) {
	h.GetCall.CallCount++
	h.GetCall.Receives.Url = url

	return h.GetCall.Returns.Response, h.GetCall.Returns.Error
}
