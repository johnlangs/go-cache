package cache

import (
	"slices"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	c := CreateCache(0, 0, false, -1)

	ok := c.Set("A", 1)
	if ok != true {
		t.Errorf("Failed to Set cache value")
	}
}

func TestGet(t *testing.T) {
	c := CreateCache(0, 0, false, -1)

	ok := c.Set("A", 1)
	if ok != true {
		t.Errorf("Failed to Set cache value")
	}

	var result interface{}
	result, ok = c.Get("A")
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

func TestDelete(t *testing.T) {
	c := CreateCache(0, 0, false, -1)

	ok := c.Set("A", 1)
	if ok != true {
		t.Errorf("Failed to Set cache value")
	}

	var result interface{}
	result, ok = c.Get("A")
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

	c.Delete("A")
	result, ok = c.Get("A")
	if result != nil {
		t.Errorf("Failed to delete value: Not nil")
	}
	if ok {
		t.Errorf("Failed to delete value: Access was Ok")
	}
}

func TestStructStorage(t *testing.T) {
	type thisStruct struct {
		num  int
		str  string
		nums []int
	}

	someObj := thisStruct{10, "Hello!", []int{1, 2, 3}}
	c := CreateCache(0, 0, false, -1)

	ok := c.Set("struct", someObj)
	if ok != true {
		t.Errorf("Failed to Set cache value")
	}

	var result interface{}
	result, ok = c.Get("struct")
	if !ok {
		t.Errorf("Failed to Get cache value: Not Ok on Get")
	}

	resultStruct, ok := result.(thisStruct)
	if !ok {
		t.Errorf("Failed to Get cache value: Not Ok on type conversion")
	}
	if (resultStruct.num != 10) || (resultStruct.str != "Hello!") || (!slices.Equal(resultStruct.nums, []int{1, 2, 3})) {
		t.Errorf("Failed to Get cache value: Item not equal to original inserted value")
	}

	c.Delete("struct")
	result, ok = c.Get("struct")
	if result != nil {
		t.Errorf("Failed to delete value: Not nil")
	}
	if ok {
		t.Errorf("Failed to delete value: Access was Ok")
	}
}

func TestStructPointerStorage(t *testing.T) {
	type thisStruct struct {
		num  int
		str  string
		nums []int
	}

	someObj := &thisStruct{10, "Hello!", []int{1, 2, 3}}
	c := CreateCache(0, 0, false, -1)

	ok := c.Set("struct", someObj)
	if ok != true {
		t.Errorf("Failed to Set cache value")
	}

	var result interface{}
	result, ok = c.Get("struct")
	if !ok {
		t.Errorf("Failed to Get cache value: Not Ok on Get")
	}

	resultStruct, ok := result.(*thisStruct)
	if !ok {
		t.Errorf("Failed to Get cache value: Not Ok on type conversion")
	}
	if (resultStruct.num != 10) || (resultStruct.str != "Hello!") || (!slices.Equal(resultStruct.nums, []int{1, 2, 3})) {
		t.Errorf("Failed to Get cache value: Item not equal to original inserted value")
	}

	c.Delete("struct")
	result, ok = c.Get("struct")
	if result != nil {
		t.Errorf("Failed to delete value: Not nil")
	}
	if ok {
		t.Errorf("Failed to delete value: Access was Ok")
	}
}

func TestMaxKeys(t *testing.T) {
	c := CreateCache(5, 1, false, 5)

	ok := c.Set("A", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("B", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("C", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("D", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("E", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}

	ok = c.Set("F", 1)
	if ok {
		t.Errorf("Set value past key amount max. Current: %d, Max: %d", c.currentKeys, c.maxKeys)
	}
}

func TestBelowMaxKeys(t *testing.T) {
	c := CreateCache(5, 1, false, 5)

	ok := c.Set("A", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("B", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("C", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("D", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}
	ok = c.Set("E", 1)
	if !ok {
		t.Errorf("Failed to set value")
	}

	c.Delete("A")
	c.Delete("B")

	ok = c.Set("F", 1)
	if !ok {
		t.Errorf("Set value when below max amount of keys. Current: %d, Max: %d", c.currentKeys, c.maxKeys)
	}
}

func TestLifetimeWatcher(t *testing.T) {
	c := CreateCache(5, 1, true, -1)
	c.Set("A", 1)

	time.Sleep(1 * time.Second)
	val, ok := c.Get("A")

	if !ok || val == nil {
		t.Errorf("Failed to get value before lifetime excceded")
	}

	time.Sleep(5 * time.Second)
	val, ok = c.Get("A")

	if ok || val != nil {
		t.Errorf("Failed to delete value after lifetime excceded")
	}
}
