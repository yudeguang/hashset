package hashset

import (
	"log"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
func TestSetThreadSafeNoGC(t *testing.T) {
	//定义并发数及测试量
	chanNum := 10
	var max = 1000000
	var maxDeleted = 500000
	var setOffical = NewThreadSafe()
	var setThreadNoGC = NewThreadSafeNoGC()
	log.Println("开始测试并发写")
	//Add
	var wgSet sync.WaitGroup
	wgSet.Add(chanNum)
	setThreadNoGC.Add("")
	setOffical.Add("")
	for i := 0; i < chanNum; i++ {
		go func(i int, wg *sync.WaitGroup) {
			for i := 0; i < max; i++ {
				setThreadNoGC.Add(strconv.Itoa(i))
				setOffical.Add(strconv.Itoa(i))
			}
			wgSet.Done()
		}(i, &wgSet)
	}
	wgSet.Wait()
	log.Println("开始测试并发读取")
	//get
	var wgGet sync.WaitGroup
	wgGet.Add(chanNum)
	for i := 0; i < chanNum; i++ {
		go func(i int, wg *sync.WaitGroup) {
			for i := 0; i < max+100; i++ {
				existOffical := setOffical.Contains(strconv.Itoa(i))
				existAny := setThreadNoGC.Contains(strconv.Itoa(i))
				if existOffical != existAny {
					t.Fatalf("unexpected value obtained; got %v want %v", existAny, existOffical)
				}
			}
			wgGet.Done()
		}(i, &wgGet)
	}
	wgGet.Wait()

	//get key empty
	exist := setThreadNoGC.Contains("")
	if !exist {
		t.Fatalf("unexpected value obtained; got %v want %v", exist, true)
	}

	log.Println("开始测试并发删除")
	//delete
	var wgDelete sync.WaitGroup
	wgDelete.Add(chanNum)
	for i := 0; i < chanNum; i++ {
		go func(i int, wg *sync.WaitGroup) {
			for i := -100; i < maxDeleted; i++ {
				setThreadNoGC.Remove(strconv.Itoa(i))
				setOffical.Remove(strconv.Itoa(i))
			}
			wgDelete.Done()
		}(i, &wgDelete)
	}
	wgDelete.Wait()
	log.Println("开始测试并发删除之后与官方结果对比")
	//get 删除之后再用get测试
	var wgGetAfterDeleted sync.WaitGroup
	wgGetAfterDeleted.Add(chanNum)
	for i := 0; i < chanNum; i++ {
		go func(i int, wg *sync.WaitGroup) {
			for i := 0; i < max+100; i++ {
				existOffical := setOffical.Contains(strconv.Itoa(i))
				existAny := setThreadNoGC.Contains(strconv.Itoa(i))
				if existOffical != existAny {
					t.Fatalf("unexpected value obtained; got %v want %v", existAny, existOffical)
				}

			}
			wgGetAfterDeleted.Done()
		}(i, &wgGetAfterDeleted)
	}
	wgGetAfterDeleted.Wait()
	log.Println("清除官方map前GC耗时")
	gcRunAndPrint()
	setOffical = nil
	log.Println("清除官方map后GC耗时")
	gcRunAndPrint()
	log.Println("开始测试同时并发读写删,因为读写锁耗时会比较长")
	//边写，边删
	var wgSetGetDelete sync.WaitGroup
	wgSetGetDelete.Add(chanNum)
	for i := 0; i < chanNum; i++ {
		go func(i int, wg *sync.WaitGroup) {
			for i := 0; i < max+100; i++ {
				setThreadNoGC.Add(strconv.Itoa(i))
				setThreadNoGC.Contains(strconv.Itoa(i - 1))
				setThreadNoGC.Remove(strconv.Itoa(i - 1))
			}
			wgSetGetDelete.Done()
		}(i, &wgSetGetDelete)
	}
	wgSetGetDelete.Wait()
}

//一般是第二次执行才是真实耗时
func gcRunAndPrint() {
	runtime.GC()
	debug.FreeOSMemory()
	begin_time := time.Now()
	runtime.GC()
	debug.FreeOSMemory()
	log.Println("GC耗时:", "耗时:", time.Now().Sub(begin_time).String())
}
