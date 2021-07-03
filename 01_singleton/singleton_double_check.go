/**
  @Author:wxuedong
  @Date:2021/6/30
  @Note:
**/
package singleton

import "sync"

type example3 struct {
	name string
}

var instance3 *example3
var once sync.Once

func GetInstance3() *example3 {
	//The Do factory_method of sync.Once can realize that the callback is only run once during the running of the program
	//Atom
	once.Do(func() {
		instance3 = new(example3)
		instance3.name = "Assign singleton for the first time"
	})
	return instance3
}
