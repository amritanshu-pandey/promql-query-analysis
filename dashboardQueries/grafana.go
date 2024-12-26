package dashboardQueries

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GrafanaClient struct {
	url                 string
	serviceAccountToken string
	insecure            bool
	httpClient          *http.Client
}

func CreateGrafanaClient(url string, serviceAccountToken string, insecure bool, skipVerify bool) (*GrafanaClient, error) {
	gc := &GrafanaClient{
		url:                 url,
		serviceAccountToken: serviceAccountToken,
		insecure:            insecure,
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	if skipVerify {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	gc.httpClient = &http.Client{Transport: customTransport}

	err := gc.checkConnection()
	if err != nil {
		return nil, err
	}
	return gc, nil
}

func (g *GrafanaClient) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.serviceAccountToken))
	resp, err := g.httpClient.Do(req)
	return resp, err
}

func (g *GrafanaClient) checkConnection() error {
	resp, err := g.get(fmt.Sprintf("%s/api/user", g.url))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("grafana returned HTTP status %s", resp.Status)
	}
	fmt.Printf("grafana returned HTTP status %s\n", resp.Status)
	return nil
}

func (g *GrafanaClient) Dashboard(id string) (any, error) {
	resp, err := g.get(fmt.Sprintf("%s/api/dashboards/uid/%s", g.url, id))
	if err != nil {
		return nil, err
	}
	dashboardString, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var dashboard any
	err = json.Unmarshal(dashboardString, &dashboard)
	if err != nil {
		fmt.Println("Unable to unmarshal dashboard response")
		return "", err
	}
	return dashboard, nil
}
