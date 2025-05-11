package hash_test

import (
	"testing"

	"github.com/primefactor-io/commitment/pkg/hash"
)

func TestCommitVerify(t *testing.T) {
	t.Parallel()

	t.Run("Commit / Verify - Valid (single-data)", func(t *testing.T) {
		t.Parallel()

		data := []byte("Hello World")

		comm, _ := hash.Commit(data)
		isValid := hash.Verify(comm, data)

		if isValid != true {
			t.Error("Commitment verification failed")
		}
	})

	t.Run("Commit / Verify - Valid (multi-data)", func(t *testing.T) {
		t.Parallel()

		data1 := []byte("Hello")
		data2 := []byte("World")

		comm, _ := hash.Commit(data1, data2)
		isValid := hash.Verify(comm, data1, data2)

		if isValid != true {
			t.Error("Commitment verification failed")
		}
	})

	t.Run("Commit / Verify - Invalid (single-data)", func(t *testing.T) {
		t.Parallel()

		data1 := []byte("Hello World")
		data2 := []byte("Hello, World")

		comm, _ := hash.Commit(data1)
		isValid := hash.Verify(comm, data2)

		if isValid != false {
			t.Error("Commitment verification failed")
		}
	})

	t.Run("Commit / Verify - Invalid (multi-data)", func(t *testing.T) {
		t.Parallel()

		data1 := []byte("Hello")
		data2 := []byte("World")
		data3 := []byte(", World")

		comm, _ := hash.Commit(data1, data2)
		isValid := hash.Verify(comm, data1, data3)

		if isValid != false {
			t.Error("Commitment verification failed")
		}
	})
}
