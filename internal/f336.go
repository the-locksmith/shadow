package internal

import (
	"unsafe"
)

func f336(ctx *Context, l0 int32, l1 float32, l2 int32, l3 int32) {
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
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
	var s6f32 float32
	_ = s6f32
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l8 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l9 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0f32
		s0i32 = l2
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l4 = s1f32
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
		l7 = s3f32
		s2f32 = s2f32 * s3f32
		l6 = s2f32
		s3f32 = l4
		s2f32 = s2f32 - s3f32
		l10 = s2f32
		s3f32 = l10
		s2f32 = s2f32 + s3f32
		s3f32 = l4
		s4i32 = l0
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
		s5f32 = l6
		s6f32 = l6
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 + s4f32
		s4f32 = l1
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s3f32 = l1
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l7
		s3f32 = -1
		s2f32 = s2f32 + s3f32
		l4 = s2f32
		s3f32 = l4
		s2f32 = s2f32 + s3f32
		l4 = s2f32
		s3f32 = 0
		s4f32 = l4
		s3f32 = s3f32 - s4f32
		s4f32 = l1
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s3f32 = l1
		s2f32 = s2f32 * s3f32
		s3f32 = 1
		s2f32 = s2f32 + s3f32
		l6 = s2f32
		s1f32 = s1f32 / s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l2
		s1f32 = l5
		s2f32 = l9
		s3f32 = l7
		s2f32 = s2f32 * s3f32
		l4 = s2f32
		s3f32 = l5
		s2f32 = s2f32 - s3f32
		l7 = s2f32
		s3f32 = l7
		s2f32 = s2f32 + s3f32
		s3f32 = l5
		s4f32 = l8
		s5f32 = l4
		s6f32 = l4
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 + s4f32
		s4f32 = l1
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s3f32 = l1
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l6
		s1f32 = s1f32 / s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	}
	s0i32 = l3
	if s0i32 != 0 {
		s0f32 = l1
		s1f32 = 0
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l5 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l7 = s0f32
		goto lbl4
	lbl5:
		s0f32 = l1
		s1f32 = 1
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
			l5 = s0f32
			s0i32 = l0
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l4 = s0f32
			s0i32 = l0
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l7 = s0f32
			s0i32 = l0
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l6 = s0f32
			goto lbl3
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l5 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l4 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l6 = s0f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l7 = s1f32
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0f32 = l4
		s1f32 = l5
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0f32
	lbl4:
		s0f32 = l5
		s1f32 = l6
		s0f32 = s0f32 - s1f32
		l5 = s0f32
		s0f32 = l7
		s1f32 = l4
		s0f32 = s0f32 - s1f32
		goto lbl2
	lbl3:
		s0f32 = l4
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l9 = s1f32
		s0f32 = s0f32 - s1f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		l4 = s1f32
		s0f32 = s0f32 * s1f32
		l8 = s0f32
		s1f32 = l4
		s2f32 = l5
		s3f32 = l9
		s2f32 = s2f32 - s3f32
		l5 = s2f32
		s1f32 = s1f32 * s2f32
		s2f32 = l5
		s1f32 = s1f32 - s2f32
		s2f32 = l1
		s1f32 = s1f32 * s2f32
		s2f32 = l5
		s3f32 = l8
		s2f32 = s2f32 - s3f32
		s3f32 = l8
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l1
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l5 = s0f32
		s0f32 = l6
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l8 = s1f32
		s0f32 = s0f32 - s1f32
		s1f32 = l4
		s0f32 = s0f32 * s1f32
		l6 = s0f32
		s1f32 = l7
		s2f32 = l8
		s1f32 = s1f32 - s2f32
		l7 = s1f32
		s2f32 = l4
		s1f32 = s1f32 * s2f32
		s2f32 = l7
		s1f32 = s1f32 - s2f32
		s2f32 = l1
		s1f32 = s1f32 * s2f32
		s2f32 = l7
		s3f32 = l6
		s2f32 = s2f32 - s3f32
		s3f32 = l6
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l1
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
	lbl2:
		l1 = s0f32
		s0i32 = l3
		s1f32 = l5
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l3
		s1f32 = l1
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	}
}
