package stresstest

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockRequester struct {
	statusCodes []int
}

func (m *mockRequester) MakeRequest(url, method string, headers []string, data []byte, timeout time.Duration) int {
	code := m.statusCodes[0]
	if len(m.statusCodes) > 1 {
		m.statusCodes = m.statusCodes[1:]
	}
	return code
}

func mockDefaultRequester(t *testing.T, statusCodes ...int) {
	t.Cleanup(func() {
		DefaultRequester = &defaultRequester{}
	})
	DefaultRequester = &mockRequester{statusCodes: statusCodes}
}

func TestStressTest_Run(t *testing.T) {
	// Assemble
	mockDefaultRequester(t, 200)
	st := &StressTest{
		Requests:    10,
		Concurrency: 4,
		Url:         "http://example.com",
	}

	// Act
	rpt, err := st.Run()

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 10, rpt.TotalRequests)
	assert.Equal(t, 10, rpt.SuccessfulRequests)
	assert.Equal(t, map[int]int{}, rpt.FailedRequests)
}

func TestStressTest_Run_Failure(t *testing.T) {
	// Assemble
	mockDefaultRequester(t, 400)
	st := &StressTest{
		Requests:    10,
		Concurrency: 4,
		Url:         "http://example.com",
	}

	// Act
	rpt, err := st.Run()

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 10, rpt.TotalRequests)
	assert.Equal(t, 0, rpt.SuccessfulRequests)
	assert.Equal(t, map[int]int{400: 10}, rpt.FailedRequests)
}

func TestStressTest_Run_Distribution(t *testing.T) {
	// Assemble
	mockDefaultRequester(t, 200)
	st := &StressTest{
		Requests:    293,
		Concurrency: 21,
		Url:         "http://example.com",
	}

	// Act
	rpt, err := st.Run()

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 293, rpt.TotalRequests)
	assert.Equal(t, 293, rpt.SuccessfulRequests)
	assert.Equal(t, map[int]int{}, rpt.FailedRequests)
}

func TestStressTest_Run_SuccessAndFailuresAndNotPerformed(t *testing.T) {
	// Assemble
	mockDefaultRequester(t, 200, 400, 0)
	st := &StressTest{
		Requests:    10,
		Concurrency: 5,
		Url:         "http://example.com",
	}

	// Act
	rpt, err := st.Run()

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 2, rpt.TotalRequests)
	assert.Equal(t, 1, rpt.SuccessfulRequests)
	assert.Equal(t, map[int]int{0: 8, 400: 1}, rpt.FailedRequests)
}

func TestStressTest_Run_SuccessAndFailures(t *testing.T) {
	// Assemble
	mockDefaultRequester(t, 200, 400, 500)
	st := &StressTest{
		Requests:    3,
		Concurrency: 3,
		Url:         "http://example.com",
	}

	// Act
	rpt, err := st.Run()

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 3, rpt.TotalRequests)
	assert.Equal(t, 1, rpt.SuccessfulRequests)
	assert.Equal(t, map[int]int{400: 1, 500: 1}, rpt.FailedRequests)
}
