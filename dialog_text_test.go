package slack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTextInput(t *testing.T) {
	name := "internalName"
	label := "Human Readable"
	value := "Pre filled text"
	optional := true
	textInput := NewTextInput(name, label, value, optional)
	assert.Equal(t, InputTypeText, textInput.Type)
	assert.Equal(t, name, textInput.Name)
	assert.Equal(t, label, textInput.Label)
	assert.Equal(t, value, textInput.Value)
	assert.Equal(t, optional, textInput.Optional)
}

func TestNewTextAreaInput(t *testing.T) {
	name := "internalName"
	label := "Human Readable"
	value := "Pre filled text"
	textInput := NewTextAreaInput(name, label, value)
	assert.Equal(t, InputTypeTextArea, textInput.Type)
	assert.Equal(t, name, textInput.Name)
	assert.Equal(t, label, textInput.Label)
	assert.Equal(t, value, textInput.Value)
}
