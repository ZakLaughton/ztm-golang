package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Expected queue length %v, got %v", i, len(q.items))
		}
		if !q.Append(i) {
			t.Errorf("Expected Append to return true, got false")
		}
	}
	if q.Append(4) {
		t.Errorf("should not be able to add to a full queue")
	}

}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able to get item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order: %v, want %v", item, i)
		}
	}
	// Queue is empty
	item, ok := q.Next()
	if ok {
		t.Errorf("should not be any more items in queue, got: %v", item)
	}
}
