package internal

import (
	"unsafe"
)

func f162(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
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
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f32 = 1
		s3i32 = l1
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s2f32 = s2f32 / s3f32
		l3 = s2f32
		s3f32 = 0
		s4f32 = l3
		s5f32 = 0
		s4f32 = s4f32 * s5f32
		s5f32 = 0
		if s4f32 == s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2f32 = s2f32
		} else {
			s2f32 = s3f32
		}
		l3 = s2f32
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l1
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2f32 = l3
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l3
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = f188(ctx, s1i32, s2f32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l1
		s1i32 = l2
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = f188(ctx, s1i32, s2f32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1i32 = l2
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1f32 = f188(ctx, s1i32, s2f32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l3 = s1f32
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+64)]))
		s1f32 = s1f32 * s2f32
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l4 = s2f32
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+76)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		l5 = s2f32
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+88)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l1
		s1f32 = l3
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+68)]))
		s1f32 = s1f32 * s2f32
		s2f32 = l4
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+80)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l5
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+92)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l3
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+72)]))
		s1f32 = s1f32 * s2f32
		s2f32 = l4
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+84)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l5
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+96)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s2i32 = 36
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = f188(ctx, s1i32, s2f32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l1
		s1i32 = l2
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = f188(ctx, s1i32, s2f32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1i32 = l2
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1f32 = f188(ctx, s1i32, s2f32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l3 = s1f32
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l1
		s1f32 = l3
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l3
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	}
}
