package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("Creating a task", func(t *testing.T) {
		err := moppleCreate("description")

		assertError(t, err, "task sucessfully created")
	})

	t.Run("Deleting a task", func(t *testing.T) {
		err := moppleDelete(1)

		assertError(t, err, "task sucessfully deleted")
	})

	t.Run("List tasks", func(t *testing.T) {
		err := moppleList()

		assertError(t, err, "list of tasks retrieved")
	})

	t.Run("try to remove the task from empty list", func(t *testing.T) {
		err := moppleDelete(2)

		assertError(t, err, "nothing to delete")
	})
}

func assertError(t testing.TB, err error, want string) {
	t.Helper()

	if err.Error() != want {
		t.Errorf("got : %q , want %q", err, want)
	}
}
