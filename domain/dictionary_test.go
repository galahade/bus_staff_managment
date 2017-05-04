package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDictionary_QueryUnique(t *testing.T) {
	setGORMShowSQL()
	dic := Dictionary{
		Name: "后勤",
		Type: 1,
	}
	err := (&dic).QueryUnique()
	assert.Nil(t, err)

}
