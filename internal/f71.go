package internal

import (
	"unsafe"
)

func f71(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	l1 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl1
	case 1:
		goto lbl0
	default:
		goto lbl2
	}
lbl2:
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+4)])
	l1 = s1i32
	s2i32 = 1
	s3i32 = l1
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 16
	s0i32 = f17(ctx, s0i32)
	l1 = s0i32
	s1i32 = 0
	s2i32 = 0
	s0i32 = f1(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = 2
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	goto lbl0
lbl1:
lbl3:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl0:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l0 = s0i32
lbl4:
	s0i32 = l0
	s0i32 = f7(ctx, s0i32)
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 29100
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 27
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	}
}
