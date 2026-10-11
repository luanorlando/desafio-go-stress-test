package loadtest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoadTestTestSuite struct {
	suite.Suite
}

// TestRun_Success200 valida se todas as requisições bem-sucedidas são contabilizadas corretamente.
func (s *LoadTestTestSuite) TestRun_Success200() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer server.Close()

	cfg := Config{
		URL:         server.URL,
		Requests:    10,
		Concurrency: 3,
	}

	metrics := Run(cfg)

	assert.Equal(s.T(), 10, metrics.TotalRequests)
	assert.Equal(s.T(), int64(10), metrics.Success200)
	assert.Equal(s.T(), 10, metrics.StatusDist[http.StatusOK])
	assert.True(s.T(), metrics.TotalDuration > 0)
}

// TestRun_MixedStatusCodes valida se a distribuição de diferentes status HTTP (ex: 200 e 500) funciona.
func (s *LoadTestTestSuite) TestRun_MixedStatusCodes() {
	count := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		if count%2 == 0 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	cfg := Config{
		URL:         server.URL,
		Requests:    6,
		Concurrency: 2,
	}

	metrics := Run(cfg)

	assert.Equal(s.T(), 6, metrics.TotalRequests)
	assert.Equal(s.T(), int64(3), metrics.Success200)
	assert.Equal(s.T(), 3, metrics.StatusDist[http.StatusOK])
	assert.Equal(s.T(), 3, metrics.StatusDist[http.StatusInternalServerError])
}

func (s *LoadTestTestSuite) TestRun_ConnectionError() {
	cfg := Config{
		URL:         "http://localhost:1",
		Requests:    3,
		Concurrency: 1,
	}

	metrics := Run(cfg)

	assert.Equal(s.T(), 3, metrics.TotalRequests)
	assert.Equal(s.T(), int64(0), metrics.Success200)
	assert.Equal(s.T(), 3, metrics.StatusDist[0])
}

func TestLoadTestTestSuite(t *testing.T) {
	suite.Run(t, new(LoadTestTestSuite))
}
