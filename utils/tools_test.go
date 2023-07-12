package utils

import "testing"

func TestGenerateRequestID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := GenerateRequestID()
		t.Log(id)
	}
}

func TestGenerateRequestTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		time := GenerateRequestTime()
		t.Log(time)
	}
}
