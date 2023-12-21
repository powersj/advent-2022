package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyStack(t *testing.T) {
	s := Stack{}
	require.Equal(t, uint64(0), s.Size())
	require.Nil(t, s.Pop())
	require.Nil(t, s.Pop())
}

func TestOneNodeStack(t *testing.T) {
	s := Stack{}
	s.Push(1)
	require.Equal(t, uint64(1), s.Size())
	require.Equal(t, 1, s.Pop())
	require.Zero(t, s.Size())
}

func TestMultiNodeStack(t *testing.T) {
	s := Stack{}
	for i := 1; i <= 5; i++ {
		s.Push(i)
		require.Equal(t, uint64(i), s.Size())
	}

	for i := 0; i < 5; i++ {
		s.Pop()
	}

	require.Equal(t, uint64(0), s.Size())
}
