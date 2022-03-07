package dummydb

import "testing"

func TestBasicQuery(t *testing.T) {
	t.Run("Set Get", func(t *testing.T) {
		db := NewDB("test-db")

		type obj struct {
			name    string
			version string
		}

		o1 := obj{
			name:    "test",
			version: "1.0",
		}

		err := db.Set("key1", o1)
		if err != nil {
			t.Errorf("should not be error: %v", err)
		}

		fetch, err := db.Get("key1")
		if err != nil {
			t.Errorf("error %v", err)
		}

		if v, ok := fetch.(obj); ok {
			t.Logf("fetched: %v", v)
		}

	})
}
