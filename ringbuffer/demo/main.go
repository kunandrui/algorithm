package main

import "fmt"

type T interface{}

type Ring struct {
	buf  []T
	size int
	w    int
	r    int
}

func NewRing(size int) *Ring {
	r := &Ring{
		buf:  make([]T, size),
		size: size,
		w:    0,
		r:    0,
	}
	return r
}

func (r *Ring) IsEmpty() bool {
	return r.w == r.r
}

func (r *Ring) Read() (T, error) {
	if r.IsEmpty() {
		return nil, fmt.Errorf("ring is empty")
	}
	v := r.buf[r.r]
	r.r++
	if r.r == r.size {
		r.r = 0
	}
	return v, nil
}

func (r *Ring) Write(v T) {
	r.buf[r.w] = v
	r.w++
	if r.w == r.size {
		r.w = 0
	}
	if r.w == r.r {
		r.grow()
	}
}

func (r *Ring) grow() {
	var size int
	if r.size < 1024 {
		size = r.size * 2
	} else {
		size = r.size + r.size/4
	}
	buf := make([]T, size)

	copy(buf[0:], r.buf[r.r:])
	copy(buf[r.size-r.r:], r.buf[0:r.r])
	r.r = 0
	r.w = r.size
	r.size = size
}
