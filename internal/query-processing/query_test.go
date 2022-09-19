package queryprocessing

import (
	"testing"
)

// test function
func TestUpdateQuery(t *testing.T) {
	limitFlag := "project=\"DataV2\""
	defaultInput := "node_cpu_seconds_total"
	expectedString := "node_cpu_seconds_total{project=\"DataV2\"}"
	actualString := UpdateQuery(defaultInput, limitFlag)

	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}

func TestUpdateQueryCase_2(t *testing.T) {
	limitFlag := "project=\"DataV2\""
	defaultInput := "node_cpu_seconds_total{}"
	expectedString := "node_cpu_seconds_total{project=\"DataV2\"}"
	actualString := UpdateQuery(defaultInput, limitFlag)

	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}

func TestUpdateQueryCase_3(t *testing.T) {
	limitFlag := "project=\"DataV2\""
	defaultInput := "node_cpu_seconds_total{service=\"Vector\"}"
	expectedString := "node_cpu_seconds_total{project=\"DataV2\",service=\"Vector\"}"
	actualString := UpdateQuery(defaultInput, limitFlag)

	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
