/**
  @Author:wxuedong
  @Date:2021/7/15
  @Note:
**/
package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg")

	assert.Nil(t, err)
}