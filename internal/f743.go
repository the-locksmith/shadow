package internal

import (
	"math"
	"unsafe"
)

func f743(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s6i32 int32
	_ = s6i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0f32 = l5
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = 1
	s0f32 = f14(ctx, s0f32, s1f32)
	s1f32 = 255
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l5
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl0
	}
	s0i32 = 0
lbl0:
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l8 = s0i32
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = 1
	s0f32 = f14(ctx, s0f32, s1f32)
	s1f32 = 255
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l4
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l4
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl2
	}
	s0i32 = 0
lbl2:
	l7 = s0i32
	s0i32 = l8
	s0f32 = float32(uint32(s0i32))
	l4 = s0f32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l7
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s0f32 = float32(uint32(s0i32))
	l5 = s0f32
	s0f32 = l3
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = 1
	s0f32 = f14(ctx, s0f32, s1f32)
	s1f32 = 255
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l3 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l3
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl4
	}
	s0i32 = 0
lbl4:
	l7 = s0i32
	s0f32 = l4
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0f32 = l5
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l7
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l2
	s3f32 = 0
	s2f32 = f15(ctx, s2f32, s3f32)
	s3f32 = 1
	s2f32 = f14(ctx, s2f32, s3f32)
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	l2 = s2f32
	s3f32 = 4.2949673e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l2
	s4f32 = 0
	if s3f32 >= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		s2f32 = l2
		s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
		goto lbl6
	}
	s2i32 = 0
lbl6:
	s3i32 = l6
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s2i32 = s2i32 + s3i32
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s3f32 = l5
	s4f32 = l4
	s5f32 = l3
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, float32, float32, float32, float32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2f32, s3f32, s4f32, s5f32)
}
