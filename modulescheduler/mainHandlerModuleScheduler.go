package moduleinternaltimer

import (
	"fmt"
	"sync"
)

//ModuleScheduler планировщик задач
type ModuleScheduler struct {
}

var once sync.Once
var mst ModuleScheduler

//MainHandlerModuleScheduler планировщик задач, реализует выполнение задач по заплонировааному времени
func MainHandlerModuleScheduler() *ModuleScheduler {
	fmt.Println("func 'MainHandlerInternalTimer', START...")

	once.Do(func() {
		mst = ModuleScheduler{}
	})

	return &mst
}
