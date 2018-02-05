package mysql

import (
	"testing"

	todo "github.com/olin-fse/integration-demo"
)

func newTestItem() todo.Item {
	i := todo.Item{
		Title:       randomString(16),
		Description: randomString(256),
	}
	return i
}

func CreateItem(t *testing.T) {
	s := newTestStore()
	defer s.Close()

	i := newTestItem()

	// Insert an item
	createdItem, err := s.CreateItem(i)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the item's id was set
	if createdItem.ID == 0 {
		t.Fatal("id not set")
	}

}
