package stack_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/pinvest/tdd-with-go/stack"
	"testing"
)

func TestNewSet(t *testing.T) {
	t.Run("The stack is empty on construction", func(t *testing.T) {
		s := stack.New()
		assert.True(t, s.IsEmpty())
	})

	t.Run("The stack has size 0 on construction", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, 0, s.Size())
	})
}

func TestPushElements(t *testing.T) {
	t.Run("After n pushes to an empty stack (n > 0), the stack is not empty and its size equals n", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		assert.False(t, s.IsEmpty())
		assert.Equal(t, 2, s.Size())
	})
}

func TestPushAndThenPopElement(t *testing.T) {
	t.Run("If one pushes x then pops, the value popped is x, and the size is one less than old size", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 3, s.Size())
		value, _ := s.Pop()
		expectedValue := 3
		assert.Equal(t, &expectedValue, value)
		assert.Equal(t, 2, s.Size())
	})
}

func TestPushAndThenPeekElement(t *testing.T) {
	t.Run("If one pushes x then pops, the value popped is x, and the size is one less than old size", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 3, s.Size())
		value, _ := s.Peek()
		expectedValue := 3
		assert.Equal(t, &expectedValue, value)
		assert.Equal(t, 3, s.Size())
	})
}

func TestError(t *testing.T) {
	t.Run("Popping from an empty stack return an error, ErrNoSuchElement", func(t *testing.T) {
		s := stack.New()
		value, err := s.Pop()
		assert.Nil(t, value)
		assert.Equal(t, stack.ErrNoSuchElement, err)
	})

	t.Run("Peeking into an empty stack return an error, ErrNoSuchElement", func(t *testing.T) {
		s := stack.New()
		value, err := s.Peek()
		assert.Nil(t, value)
		assert.Equal(t, stack.ErrNoSuchElement, err)
	})
}
