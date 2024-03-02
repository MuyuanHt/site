package utils

import (
	"fmt"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"sync"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	var wg sync.WaitGroup
	s, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		t.Error(err)
		return
	}
	var check sync.Map
	t1 := time.Now()
	for i := 0; i < 200000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := s.NextVal()
			//t.Logf("%v", val)
			if _, ok := check.Load(val); ok {
				// id冲突检查
				t.Error(fmt.Errorf("error#unique: val:%v", val))
				return
			}
			check.Store(val, 0)
			if val == 0 {
				t.Error(fmt.Errorf("error"))
				return
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(t1)
	t.Logf("generate 20k ids elapsed: %v", elapsed)
}

func TestCreateSnowflakeId(t *testing.T) {
	s, err := InitSnowflake(0, 0)
	if err != nil {
		t.Error(err)
	}
	count := 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := GetNextSnowflakeId(s)
			t.Logf("第 %v 次：%v", i, id)
		}(i)
	}
	wg.Wait()
	t.Logf("count: %v", count)
}
