package pokecache

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	timeDuration := 5 * time.Second
	fmt.Println("Test Case 1 - Create Empty Cache")
	expected := Cache{
		Mapper: make(map[string]CacheEntry),
		Mux:    &sync.Mutex{},
	}
	actual := NewCache(timeDuration)
	if(len(expected.Mapper) != len(actual.Mapper)) {
		t.Errorf("Expected " + strconv.Itoa(len(expected.Mapper)) + " but found " + strconv.Itoa(len(actual.Mapper)))
	}
}

func TestAdd(t *testing.T) {
	timeDuration := 2 * time.Second
	fmt.Println("Test Case 1 - Add new cache into Cache")
	expectedLength := 1
	actual := NewCache(timeDuration)
	actual.Add("key", []byte("TestString"))
	if(expectedLength != len(actual.Mapper)) {
		t.Errorf("Expected 1 but found " + strconv.Itoa(len(actual.Mapper)))
	}
}

func TestGet(t *testing.T) {
	timeDuration := 2 * time.Second
	fmt.Println("Test Case 1 - Get existing key")
	cache := NewCache(timeDuration)
	cache.Add("key", []byte("TestString"))
	_, ok := cache.Get("key")
	if(!ok) {
		t.Errorf("Expected 'key' but found ''")
	}
	fmt.Println("Test Case 2 - Get non-existant key")
	cache = NewCache(timeDuration)
	_, ok = cache.Get("key")
	if(ok) {
		t.Errorf("Expected '' but found 'key'")
	}
}

func TestReaploop(t *testing.T) {
	timeDuration := 2 * time.Second
	cache := NewCache(timeDuration)
	fmt.Println("Test Case 1 - Key Removal")
	cache.Add("key", []byte("TestString"))
	time.Sleep(timeDuration + time.Second)
	_, ok := cache.Get("key")
	if(ok) {
		t.Errorf("Expected '' but found 'key'")
	}
	fmt.Println("Test Case 2 - Non-key Removal")
	cache = NewCache(timeDuration)
	cache.Add("key", []byte("TestString"))
	_, ok = cache.Get("key")
	if(!ok) {
		t.Errorf("Expected 'key' but found ''")
	}
}