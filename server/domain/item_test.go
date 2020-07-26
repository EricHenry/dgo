package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestItem(t *testing.T) {
	id, _ := uuid.NewRandom()
	item := NewItem(id, "Chex Mix", "This is Chex Mix", 10.00)
	ID := item.GetID()

	if ID != id {
		t.Errorf("ID is not correct. Wanted %s, Got %s", "abc123", ID)
	}
}
