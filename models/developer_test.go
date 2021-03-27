package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TetstNewDeveloperWillReturnExpectedJSONPayload(t *testing.T) {
	// given
	name := "Rob"
	team := "DPE"
	skills := []string{"integration"}

	// when
	developer := NewDeveloper(&name, &team, &skills)
	bytes, _ := json.Marshal(developer)

	// then
	assert.Equal(t, `{}`, string(bytes))
}
