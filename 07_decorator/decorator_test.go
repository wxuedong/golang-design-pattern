/**
  @Author:wxuedong
  @Date:2021/7/29
  @Note:
**/
package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}