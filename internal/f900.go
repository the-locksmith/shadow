package internal

import (
	"unsafe"
)

func f900(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 int32
	_ = l6
	var l7 float32
	_ = l7
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
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 4
	s1i32 = s1i32 & s2i32
	s2i32 = l6
	s3i32 = 4
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 16
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l6
	s3i32 = 2
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 1
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l6
	s2i32 = s2i32 ^ s3i32
	l6 = s2i32
	s3i32 = 5
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 32
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l6
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 8
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l6
	s3i32 = 1
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 2
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.015625
	s1f32 = s1f32 * s2f32
	s2f32 = -0.4921875
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	l7 = s0f32
	s1f32 = l4
	s0f32 = s0f32 + s1f32
	s1f32 = l5
	s0f32 = f14(ctx, s0f32, s1f32)
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	l4 = s0f32
	s0f32 = l7
	s1f32 = l3
	s0f32 = s0f32 + s1f32
	s1f32 = l5
	s0f32 = f14(ctx, s0f32, s1f32)
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	l3 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l7
	s3f32 = l2
	s2f32 = s2f32 + s3f32
	s3f32 = l5
	s2f32 = f14(ctx, s2f32, s3f32)
	s3f32 = 0
	s2f32 = f15(ctx, s2f32, s3f32)
	s3f32 = l3
	s4f32 = l4
	s5f32 = l5
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
