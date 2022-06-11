package services_test

import (
	"testing"

	"github.com/sing3demons/gotest/services"
)

func TestCheckGrade(t *testing.T) {
	grade := services.CheckGrade(81)
	expected := "A"

	if grade != expected {
		t.Errorf("got %v expected %v", grade, expected)
	}
}
