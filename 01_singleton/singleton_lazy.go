/**
  @Author:wxuedong
  @Date:2021/6/29
  @Note:
**/
package singleton

type singleton struct {
	name string
}

// private
var instance *singleton

//The lazy mode has thread safety issues.
//If multiple threads call this method at the same time, it will detect that the instance is nil, and multiple objects will be created, so the hungry mode appears
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	instance.name = "Not thread safe"
	return instance
}
