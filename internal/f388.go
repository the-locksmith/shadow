package internal

import (
	"unsafe"
)

func f388(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l3
	s1i32 = 120
	s2i32 = 4
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s0i32 = l3
	s1i32 = l4
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1209
	s2i32 = l4
	s3i32 = l5
	s2i32 = s2i32 - s3i32
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s1i32 = l0
	s2i32 = l1
	s3i32 = l2
	s0i32 = f1511(ctx, s0i32, s1i32, s2i32, s3i32)
	s1i32 = 0
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	return s0i32
}
