package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T){
	// var (
	// 	input =1
	// 	output =2
	// )
	// actual := AddOne(1)
	// if actual != output {
    //     t.Errorf("AddOne %d, input %d , actual = %d", input, output, actual)
    // }
	assert.Equal(t,AddOne(3),4, "AddOne")
	assert.NotEqual(t, 2,3)
	assert.Nil(t, nil, nil)
}
