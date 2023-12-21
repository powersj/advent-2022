package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyDoubleList(t *testing.T) {
	list := DoubleLinkedList{}

	require.Equal(t, list.Length(), 0)
	require.Equal(t, "", list.String())

	require.Error(t, list.Remove("testing"))
	require.False(t, list.Contains("testing"))

	_, err := list.Head()
	require.Error(t, err)
	_, err = list.Tail()
	require.Error(t, err)
}

func TestOneNodeDoubleList(t *testing.T) {
	expected := 42

	list := DoubleLinkedList{}
	require.NoError(t, list.Add(expected))
	require.Equal(t, list.Length(), 1)
	require.Equal(t, "42", list.String())

	require.Error(t, list.Remove("testing"))
	require.False(t, list.Contains("testing"))

	val, err := list.Head()
	require.NoError(t, err)
	require.Equal(t, expected, val)

	val, err = list.Tail()
	require.NoError(t, err)
	require.Equal(t, expected, val)

	require.True(t, list.Contains(expected))
	require.NoError(t, list.Remove(expected))
}

func TestRemoveHeadDoubleList(t *testing.T) {
	list := DoubleLinkedList{}
	require.NoError(t, list.Add(1))
	require.NoError(t, list.Add(2))

	require.NoError(t, list.Remove(1))
	require.Equal(t, list.Length(), 1)
	require.Equal(t, "2", list.String())

	require.NoError(t, list.Remove(2))
	require.Equal(t, list.Length(), 0)
	require.Equal(t, "", list.String())
}

func TestMultiNodeDoubleList(t *testing.T) {
	list := DoubleLinkedList{}
	require.NoError(t, list.Add(1))
	require.NoError(t, list.Add(2))
	require.NoError(t, list.Add(3))
	require.NoError(t, list.Add(4))
	require.Equal(t, 4, list.Length())

	val, err := list.Tail()
	require.NoError(t, err)
	require.Equal(t, 4, val)

	require.NoError(t, list.Remove(3))
	require.Equal(t, 3, list.Length())

	val, err = list.Tail()
	require.NoError(t, err)
	require.Equal(t, 4, val)

	require.NoError(t, list.Remove(4))
	require.Equal(t, 2, list.Length())

	val, err = list.Tail()
	require.NoError(t, err)
	require.Equal(t, 2, val)
}

func TestPrintDoubleList(t *testing.T) {
	list := DoubleLinkedList{}
	require.NoError(t, list.Add("a"))
	require.NoError(t, list.Add("b"))
	require.NoError(t, list.Add("c"))
	require.NoError(t, list.Add("d"))
	require.Equal(t, 4, list.Length())

	require.Equal(t, "a -> b -> c -> d", list.String())
}
