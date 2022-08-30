# hashset
github.com/yudeguang/hashset, https://github.com/yudeguang/noGcMap, https://github.com/yudeguang/noGcStaticMap 为同一系列的无GC类型 

本包除了提供若干基础类型的hashset 还提供无GC类型的StringThreadSafeNoGC


```go
package main

import (
	"github.com/yudeguang/hashset"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	s := hashset.NewStringThreadSafeNoGC()
	s.Add("1")
	log.Println(s.Contains("1"))
	s.Remove("1")
	log.Println(s.Contains("1"))
}


```
