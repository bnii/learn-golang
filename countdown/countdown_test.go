package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockSleeper struct {
	mock.Mock
}

func (m *MockSleeper) Sleep() {
	m.Called()
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	mockSleeper := new(MockSleeper)
	mockSleeper.On("Sleep").Return()
	Countdown(buffer, mockSleeper)

	assert.Equal(t, `3
2
1
Go!`, buffer.String())

	mockSleeper.AssertNumberOfCalls(t, "Sleep", 3)
}
