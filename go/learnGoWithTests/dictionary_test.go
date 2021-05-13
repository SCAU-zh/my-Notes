package learnGoWithTests

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	//查询存在的键值
	t.Run("search know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertSearch(t, got, want)
	})

	//查询不存在的键值
	t.Run("search unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		assertSearchUnknown(t, err, want)
	})
	t.Run("search unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		assertSearchUnknown(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	dictionary.Add(word, definition)
	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func assertSearch(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertSearchUnknown(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didnt get one")
	}
	if got.Error() != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
