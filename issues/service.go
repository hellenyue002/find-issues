package issues

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type httpClient interface {
	Get(url string) (*http.Response, error)
}

type Issues struct {
	Number int
	Title  string
	Url    string
	Labels []Label
}

type Label struct {
	Name string
}

type Service struct {
	repo   string
	client httpClient
}

func NewService(repo string, client httpClient) Service {
	return Service{
		repo:   repo,
		client: client,
	}
}

func (i Service) Get() ([]Issues, error) {
	u := &url.URL{
		Scheme: "https",
		// TODO: Accept a full path and parse the org/repo from it?
		// go run main.go github.com/ghc-tdd/spike
		// api.github.com/repos/ghc-tdd/spike/issues
		Path: fmt.Sprintf("api.github.com/repos/%s/issues", i.repo),
	}

	res, err := i.client.Get(u.String())
	if err != nil {
		return []Issues{}, err
	}

	if res.StatusCode != 200 {
		return []Issues{}, fmt.Errorf("invalid path %s", u.String())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Issues{}, err
	}

	var issues []Issues
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return []Issues{}, err
	}

	return issues, nil
}
