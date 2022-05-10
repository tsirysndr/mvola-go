package mvola

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	defaultTestTimeout = time.Second * 1
	accessToken        = "eyJ4NXQiOiJPRE5tWkRFMll6UTRNVEkxTVRZME1tSmhaR00yTUdWa1lUZGhOall5TWpnM01XTmpNalJqWWpnMll6bGpNRGRsWWpZd05ERmhZVGd6WkRoa1lUVm1OZyIsImtpZCI6Ik9ETm1aREUyWXpRNE1USTFNVFkwTW1KaFpHTTJNR1ZrWVRkaE5qWXlNamczTVdOak1qUmpZamcyWXpsak1EZGxZall3TkRGaFlUZ3paRGhrWVRWbU5nX1JTMjU2IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJ0c2lyeS5zbmRyQGdtYWlsLmNvbUBjYXJib24uc3VwZXIiLCJhdXQiOiJBUFBMSUNBVElPTiIsImF1ZCI6IjlCV3pjcmdGRW1FdHJXT1EwRW5FUU1UOEtEZ2EiLCJuYmYiOjE2NTIyMDYzODgsImF6cCI6IjlCV3pjcmdGRW1FdHJXT1EwRW5FUU1UOEtEZ2EiLCJzY29wZSI6IkVYVF9JTlRfTVZPTEFfU0NPUEUiLCJpc3MiOiJodHRwczpcL1wvYXBpbS5wcmVwLnRlbG1hLm1nOjk0NDNcL29hdXRoMlwvdG9rZW4iLCJleHAiOjE2NTIyMDk5ODgsImlhdCI6MTY1MjIwNjM4OCwianRpIjoiZWI2NWQ2NzgtYmNlYy00Yjg4LWI1MzAtNGM2YmY3ODA4MWU1In0.nx-kk21G1OTofEf0q9Iya_ESTtTJ9Po7AgcV_vV8ROKBAOrAH6HE2ckkK4S9S1WjM0zBQl7e0qtOIYgaN9Oge4BL0ORuvcY5eqammuGAtTku95GzYlDKo0PwQGC_pdbFltKLwnbvA66a8SkXSwiL3OgoC6NRYqmMYF_qQnwF5ZcGVlFnwp-yzLf2ojzgGlHY2o8gDyaEqwThxARmLK5vjcQdv5GoO0h72vrWIfmKZydq1MLQDzLYhSyRZ-q5luoJeTU7tlToIOaBtI-ppTliSTt4TdZi5-5qnLw45wDKGUy5IELs2LSc348JDPq9w7q1Fs2HCH26j7ugw39HZbJ_2Q"
)

func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	client := &http.Client{
		Transport: transport,
	}
	return client, mux, server
}

type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	assert.Equal(t, expectedMethod, req.Method)
}
