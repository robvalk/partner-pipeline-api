package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonMarshals(t *testing.T) {
	developer := Developer{
		Name: "Dev Dangerfield",
	}

	_, err := json.Marshal(developer)

	if err != nil {
		t.Error("JSON marchaling through an error", err)
	}
}

func TestJsonUnmarshals(t *testing.T) {
	someJson := "{\"name\": \"Dev Dangerfield\"}"

	developer := Developer{}

	err := json.Unmarshal([]byte(someJson), &developer)

	if err != nil {
		t.Error("JSON marchaling through an error", err)
	}

	assert.Equal(t, "Dev Dangerfield", developer.Name)
}
