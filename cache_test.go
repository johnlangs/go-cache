package cache

import "testing"

func TestSet(t *testing.T) {
	c := CreateCache()

	err := c.Set("A", 1)
	if err != nil {
		t.Errorf("Failed to Set cache value")
	}
}

func TestGet(t *testing.T) {
	c := CreateCache()

	err := c.Set("A", 1)
	if err != nil {
		t.Errorf("Failed to Set cache value")
	}

	var result interface{}
	result, ok := c.Get("A")
	if !ok {
		t.Errorf("Failed to Get cache value: Not Ok on Get")
	}

	resultInt, ok := result.(int)
	if !ok {
		t.Errorf("Failed to Get cache value: Not Ok on type conversion")
	}
	if resultInt != 1 {
		t.Errorf("Failed to Get cache value: Item not equal to original inserted value")
	}
}
