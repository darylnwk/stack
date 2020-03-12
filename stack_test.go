package stack_test

import (
	"testing"

	"github.com/darylnwk/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack_New(t *testing.T) {
	s := stack.New()

	assert.IsType(t, &stack.Stack{}, s)
}

func TestStack(t *testing.T) {
	stack := stack.New()

	stack.Push("hello", "world", 1)

	assert.Equal(t, 3, stack.Size())
	assert.Equal(t, 1, stack.Peek())
	assert.Equal(t, 1, stack.Pop())

	assert.Equal(t, 2, stack.Size())
	assert.Equal(t, "world", stack.Peek())
	assert.Equal(t, "world", stack.Pop())

	assert.Equal(t, 1, stack.Size())
	assert.Equal(t, "hello", stack.Peek())
	assert.Equal(t, "hello", stack.Pop())

	assert.Equal(t, 0, stack.Size())
	assert.Nil(t, stack.Peek())
	assert.Nil(t, stack.Pop())
}
