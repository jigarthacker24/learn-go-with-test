package maps

import (
	"testing"
)

func TestDelete(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		word := "test"
		def := "this is test def"
		dict := Dictionary{word: def}

		dict.Delete(word)

		_, err := dict.Search(word)

		assertError(t, err, ErrNotFound)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "This is just a test"

		dict := Dictionary{word: def}
		newDef := "This is new def for test"

		err := dict.Update(word, newDef)

		assertError(t, err, nil)
		assertDictionaryDefinition(t, dict, word, newDef)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "This is just a test"

		dict := Dictionary{word: def}
		newWord := "test2"
		newDef := "This is test2 def"

		err := dict.Update(newWord, newDef)
		assertError(t, err, ErrWordNotExists)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def := "This is just a test"
		err := dict.Add(word, def)
		assertError(t, err, nil)
		assertDictionaryDefinition(t, dict, word, def)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "This is just a test"
		dict := Dictionary{word: def}
		err := dict.Add(word, "This is new def for test")
		assertError(t, err, ErrWordExists)
		assertDictionaryDefinition(t, dict, word, def)
	})

}
func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "This is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		assertString(t, got, "This is just a test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("Want error. No Error received")
		}

		assertError(t, err, ErrNotFound)
	})

}

func assertString(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:%q, want:%q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got:%q, want:%q", got, want)
	}
}

func assertDictionaryDefinition(t testing.TB, dict Dictionary, word, def string) {
	t.Helper()
	got, err := dict.Search(word)

	if err != nil {
		t.Fatalf("Error received. Error: %q", err)
	}
	if got != def {
		t.Errorf("got: %q, want: %q, given:%q", got, def, "test")
	}

}
