package tests

import (
	"fmt"
	"testing"
)

type geodeDB struct{}

//Call geode server and GET results depends on key
func (g geodeDB) Get(key string) ([]string, bool) {
	return []string{"A", "B", "C", "D", "E"}, true
}

//Call geode server and PUSH results
func (g geodeDB) Push(key string, values []string) {}

type kvdbClient interface {
	Get(key string) ([]string, bool)
	Push(key string, values []string)
}

type inMemoryDB struct {
	m map[string][]string
}

func (im *inMemoryDB) Get(key string) ([]string, bool) {
	if im.m == nil {
		im.m = make(map[string][]string)
	}
	result, ok := im.m[key]
	return result, ok
}
func (im *inMemoryDB) Push(key string, values []string) {
	if im.m == nil {
		im.m = make(map[string][]string)
	}
	im.m[key] = values
}

func SaveAndGet(cli kvdbClient, key string, values []string) bool {
	cli.Push(key, values)
	//Doing something
	results, ok := cli.Get(key)
	fmt.Print(results)
	return ok
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

func Test_SaveAndGet(t *testing.T) {
	//when
	ok := SaveAndGet(&inMemoryDB{}, "key1", []string{"A", "B", "C"})
	//then
	if !ok {
		t.Errorf("SaveAndGet operation failed")
	}
}
