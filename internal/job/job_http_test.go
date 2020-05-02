package job

import (
	"errors"
	apiPb "github.com/squzy/squzy_generated/generated/proto/v1"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type httpToolsMock struct {
}

func (h httpToolsMock) SendRequestTimeoutStatusCode(req *http.Request, timeout time.Duration, expectedCode int, ) (int, []byte, error) {
	return 0, nil, nil
}

func (h httpToolsMock) SendRequestTimeout(req *http.Request, timeout time.Duration) (int, []byte, error) {
	panic("implement me")
}

func (h httpToolsMock) CreateRequest(method string, url string, headers *map[string]string, log string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	return req
}

type httpToolsMockError struct {
}

func (h httpToolsMockError) SendRequestTimeoutStatusCode(req *http.Request, timeout time.Duration, expectedCode int, ) (int, []byte, error) {
	return 0, nil, errors.New("safsaf")
}

func (h httpToolsMockError) SendRequestTimeout(req *http.Request, timeout time.Duration) (int, []byte, error) {
	panic("implement me")
}

func (h httpToolsMockError) GetWithRedirectsWithStatusCode(url string, expectedCode int) (int, []byte, error) {
	panic("implement me")
}

func (h httpToolsMockError) GetWithRedirects(url string) (int, []byte, error) {
	panic("implement me")
}

func (h httpToolsMockError) CreateRequest(method string, url string, headers *map[string]string, log string) *http.Request {
	rq, _ := http.NewRequest(method, url, nil)
	return rq
}

func (h httpToolsMockError) SendRequest(req *http.Request) (int, []byte, error) {
	panic("implement me")
}

func (h httpToolsMockError) SendRequestWithStatusCode(req *http.Request, expectedCode int) (int, []byte, error) {
	return 0, nil, errors.New("safsaf")
}

func (h httpToolsMock) SendRequest(req *http.Request) (int, []byte, error) {
	panic("implement me")
}

func (h httpToolsMock) SendRequestWithStatusCode(req *http.Request, expectedCode int) (int, []byte, error) {
	return 0, nil, nil
}

func TestNewHttpJob(t *testing.T) {
	t.Run("Should: implement interface", func(t *testing.T) {
		s := NewHttpJob(http.MethodGet, "", map[string]string{}, 0, http.StatusOK, &httpToolsMock{})
		assert.Implements(t, (*Job)(nil), s)
	})

}

func TestJobHTTP_Do(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := NewHttpJob(http.MethodGet, "", map[string]string{}, 0, http.StatusOK, &httpToolsMock{})
		assert.Equal(t, apiPb.SchedulerResponseCode_OK, s.Do("").GetLogData().Code)
	})
	t.Run("Should: return error because long request", func(t *testing.T) {
		s := NewHttpJob(http.MethodGet, "", map[string]string{}, 0, http.StatusOK, &httpToolsMock{})
		assert.Equal(t, apiPb.SchedulerResponseCode_OK, s.Do("").GetLogData().Code)
	})
	t.Run("Should: not return error with headers", func(t *testing.T) {
		s := NewHttpJob(http.MethodGet, "", map[string]string{
			"test": "asf",
		}, http.StatusOK, 0, &httpToolsMock{})
		assert.Equal(t, apiPb.SchedulerResponseCode_OK, s.Do("").GetLogData().Code)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := NewHttpJob(http.MethodGet, "", map[string]string{}, 0, http.StatusOK, &httpToolsMockError{})
		assert.Equal(t, apiPb.SchedulerResponseCode_Error, s.Do("").GetLogData().Code)
	})
}
