package internal

import (
	"math"
	"unsafe"
)

func f1697(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0f32
	s0i32 = 1
	l3 = s0i32
	s0i32 = 1
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l6 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f32 = s1f32 - s2f32
	l7 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = 1
	s1f32 = l5
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l7
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l8
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
lbl0:
	l4 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = l6
	s0f32 = s0f32 - s1f32
	l6 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l6
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l5
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	l3 = s0i32
lbl1:
	s0i32 = 0
	s1i32 = l3
	s2i32 = l4
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	s1i32 = l3
	s2i32 = l4
	s1i32 = s1i32 | s2i32
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = 2
	s1i32 = l0
	s1i32 = f618(ctx, s1i32)
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	s1i32 = l0
	s1f32 = f692(ctx, s1i32)
	l5 = s1f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	s1f32 = l5
	s2f32 = 1
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2f32 = l5
	f339(ctx, s0i32, s1i32, s2f32)
	s0i32 = l1
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = 3
lbl2:
	l0 = s0i32
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
