package internal

import (
	"unsafe"
)

func f1092(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s1i32 = l0
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	f12(ctx, s0i32)
lbl0:
	s0i32 = l1
	s1i32 = 116
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i64 = 51539607552
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l0
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		goto lbl1
	}
	s0i32 = l1
	s0i32 = f17(ctx, s0i32)
	l1 = s0i32
	s1i64 = 51539607552
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = 8
	s0i32 = s0i32 + s1i32
lbl1:
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
