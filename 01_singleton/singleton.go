/**
  @Author:wxuedong
  @Date:2021/6/29
  @Note:
**/
package _1_singleton

type singleton struct {
}

// private
var instance *singleton

// public
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}     // not thread safe
	}
	return instance
}
