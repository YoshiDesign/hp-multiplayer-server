package deque

import (
	"errors"
)

/**
 * Double-Ended Ring-Buffer Queue
 *
 * Keep in mind that we use a half-open index
 * [tail, head)
 * tail points to the logical index of the first value.
 * head points 1 beyond the back element.
 *
 * Fun Fact:
 * Go uses 2's complement integer behavior, so negative integers
 * will still wrap accordingly with our bitmask wrapping strat.
 */

type DeqRing[T any] struct {
	
	buffer 		[]T
	mask		int
	tail		int
	head  		int
	size		int
}

func (dq *DeqRing[T]) Initialize(capacity int) {
	// Bitmask Wrapping. capacity must be a power of 2
	dq.mask = capacity - 1
	dq.buffer = make([]T, capacity)
	dq.tail = 0
	dq.head = 0
	dq.size = 0
}

func (dq *DeqRing[T]) IsFull() bool {
	return dq.size == len(dq.buffer)
}

func (dq *DeqRing[T]) IsEmpty() bool {
	return dq.size == 0
}

func (dq *DeqRing[T]) PushBack(val T) bool {

	if dq.IsFull() {
		return false
	}

	idx := dq.head & dq.mask
	dq.buffer[idx] = val
	// Grow from head
	dq.head++
	dq.size++
	return true
}

func (dq *DeqRing[T]) PushFront(val T) bool {

	if dq.IsFull() {
		return false
	}

	// Grow the tail (decrements)
	dq.tail-- 
	idx := dq.tail & dq.mask 
	dq.buffer[idx] = val
	dq.size++
	return true
	
}

func (dq *DeqRing[T]) PopBack() (T, bool) {

	if dq.IsEmpty() {
		var zero T
		return zero, false
	}

	dq.head--
	out := dq.buffer[dq.head & dq.mask]
	dq.size--

	return out, true
}

func (dq *DeqRing[T]) PopFront() (T, bool) {

	if dq.IsEmpty() {
		var zero T
		return zero, false
	}

	out := dq.buffer[dq.tail & dq.mask]
	dq.tail++
	dq.size--

	return out, true
}

func (dq DeqRing[T]) Back() (T, error) {
	if dq.IsEmpty() {
		var zero T
		return zero, errors.New("dq is empty") 
	}

	return dq.buffer[(dq.head - 1) & dq.mask], nil
}


