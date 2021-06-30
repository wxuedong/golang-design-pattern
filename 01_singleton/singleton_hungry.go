/**
  @Author:wxuedong
  @Date:2021/6/30
  @Note:
**/
package singleton

type example2 struct {
	name string
}

var instance2 *example2

func init() {
	instance2 = new(example2)
	instance2.name = "Initialize singleton"
}

// The hungry mode will create a singleton object when the package is loaded.
//  When the object is not used in the program, a part of the space is wasted.
//  Compared with the lazy mode, it is safer but slows down the program startup speed.
func GetInstance2() *example2 {
	return instance2
}
