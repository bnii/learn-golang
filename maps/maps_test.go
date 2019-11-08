package maps

import (
	"errors"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("unknown  word", func(t *testing.T) {
		got, err := dictionary.Search("mas")
		wantErr := errors.New("not found")
		assert.Equal(t, wantErr, err)
		assert.Zero(t, got)
	})

}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		myKey := "test"
		myVal := "testVal"
		errAdd := dictionary.Add(myKey, myVal)
		gotVal, err := dictionary.Search(myKey)

		assert.Nil(t, errAdd)
		assert.Nil(t, err)
		assert.Equal(t, myVal, gotVal)
	})
	t.Run("existing word", func(t *testing.T) {
		myKey := "test"
		dictionary := Dictionary{myKey: "val1"}
		errAdd := dictionary.Add(myKey, "val2")

		assert.Equal(t, ErrAlreadyExists, errAdd)
	})
}

func TestUpdate(t *testing.T) {
	const key = "test"
	t.Run("update existing", func(t *testing.T) {
		dictionary := Dictionary{key: "origval"}
		errUpdate := dictionary.Update(key, "newval")
		assert.Nil(t, errUpdate)
		got, errSearch := dictionary.Search(key)
		assert.Equal(t, "newval", got)
		assert.Nil(t, errSearch)
	})
	t.Run("update non existing", func(t *testing.T) {
		dictionary := Dictionary{}
		errUpdate := dictionary.Update(key, "newval")

		assert.Equal(t, ErrUpdateNonExisting, errUpdate)

	})
}
func TestDelete(t *testing.T) {
	t.Run("delete existing", func(t *testing.T) {
		dictionary := Dictionary{"test": "someval"}
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("delete non-existing", func(t *testing.T) {
		dictionary := Dictionary{}
		errDelete := dictionary.Delete("test")
		assert.Equal(t, ErrDeleteNonExisting, errDelete)

	})
}
