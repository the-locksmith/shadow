package internal

import (
	"unsafe"
)

func f288(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
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
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s0i32 = l2
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1i32
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	l8 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l17 = s0f32
	s1f32 = 0.05
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 12
		s0i32 = s0i32 | s1i32
		s1i32 = l2
		s2i32 = l6
		s3i32 = 12
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l14 = s1f32
		s2f32 = 0.05
		if s1f32 <= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s1i32 = l6
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l13 = s0f32
		s0i32 = l2
		s1i32 = l9
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0f32
		s0i32 = l5
		s1f32 = l17
		s2f32 = l14
		s3f32 = l17
		s2f32 = s2f32 - s3f32
		s3f32 = l17
		s4f32 = -0.05
		s3f32 = s3f32 + s4f32
		l10 = s3f32
		s4f32 = l10
		s5f32 = l14
		s6f32 = -0.05
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 / s4f32
		l10 = s3f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l5
		s1f32 = l11
		s2f32 = l13
		s3f32 = l11
		s2f32 = s2f32 - s3f32
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l5
		s1f32 = l12
		s2f32 = l16
		s3f32 = l12
		s2f32 = s2f32 - s3f32
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l5
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l2
	s2i32 = l6
	s3i32 = 12
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l14 = s1f32
	s2f32 = 0.05
	if s1f32 > s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l6
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s0i32 = l2
	s1i32 = l9
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l11 = s0f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s0i32 = l5
	s1f32 = l17
	s2f32 = l14
	s3f32 = l17
	s2f32 = s2f32 - s3f32
	s3f32 = l17
	s4f32 = -0.05
	s3f32 = s3f32 + s4f32
	l10 = s3f32
	s4f32 = l10
	s5f32 = l14
	s6f32 = -0.05
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 / s4f32
	l10 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l5
	s1f32 = l11
	s2f32 = l13
	s3f32 = l11
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l5
	s1f32 = l12
	s2f32 = l16
	s3f32 = l12
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l5
	s1i32 = 12
	s0i32 = s0i32 | s1i32
lbl0:
	l3 = s0i32
	s0i32 = l2
	s1i32 = l6
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l7 = s0i32
	s0f32 = l14
	s1f32 = 0.05
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l7
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l15 = s0f32
		s1f32 = 0.05
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l2
		s1i32 = l7
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l13 = s0f32
		s0i32 = l2
		s1i32 = l6
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0f32
		s0i32 = l3
		s1f32 = l14
		s2f32 = l15
		s3f32 = l14
		s2f32 = s2f32 - s3f32
		s3f32 = l14
		s4f32 = -0.05
		s3f32 = s3f32 + s4f32
		l10 = s3f32
		s4f32 = l10
		s5f32 = l15
		s6f32 = -0.05
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 / s4f32
		l10 = s3f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l3
		s1f32 = l11
		s2f32 = l13
		s3f32 = l11
		s2f32 = s2f32 - s3f32
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l3
		s1f32 = l12
		s2f32 = l16
		s3f32 = l12
		s2f32 = s2f32 - s3f32
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		goto lbl2
	}
	s0i32 = l3
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l7
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l15 = s0f32
	s1f32 = 0.05
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		goto lbl2
	}
	s0i32 = l2
	s1i32 = l7
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s0i32 = l2
	s1i32 = l6
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l11 = s0f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s0i32 = l3
	s1f32 = l14
	s2f32 = l15
	s3f32 = l14
	s2f32 = s2f32 - s3f32
	s3f32 = l14
	s4f32 = -0.05
	s3f32 = s3f32 + s4f32
	l10 = s3f32
	s4f32 = l10
	s5f32 = l15
	s6f32 = -0.05
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 / s4f32
	l10 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l3
	s1f32 = l11
	s2f32 = l13
	s3f32 = l11
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l3
	s1f32 = l12
	s2f32 = l16
	s3f32 = l12
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l3
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl2:
	s0i32 = l2
	s1i32 = l7
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0f32 = l15
	s1f32 = 0.05
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l17
		s1f32 = 0.05
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l2
		s1i32 = l9
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l13 = s0f32
		s0i32 = l2
		s1i32 = l7
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0f32
		s0i32 = l3
		s1f32 = l15
		s2f32 = l17
		s3f32 = l15
		s2f32 = s2f32 - s3f32
		s3f32 = l15
		s4f32 = -0.05
		s3f32 = s3f32 + s4f32
		l10 = s3f32
		s4f32 = l10
		s5f32 = l17
		s6f32 = -0.05
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 / s4f32
		l10 = s3f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l3
		s1f32 = l11
		s2f32 = l13
		s3f32 = l11
		s2f32 = s2f32 - s3f32
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l3
		s1f32 = l12
		s2f32 = l16
		s3f32 = l12
		s2f32 = s2f32 - s3f32
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		goto lbl5
	}
	s0i32 = l3
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l8
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l15 = s0f32
	s1f32 = 0.05
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		goto lbl5
	}
	s0i32 = l2
	s1i32 = l9
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s0i32 = l2
	s1i32 = l7
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l11 = s0f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s0i32 = l3
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l10 = s1f32
	s2f32 = l15
	s3f32 = l10
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s4f32 = -0.05
	s3f32 = s3f32 + s4f32
	l10 = s3f32
	s4f32 = l10
	s5f32 = l15
	s6f32 = -0.05
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 / s4f32
	l10 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l3
	s1f32 = l11
	s2f32 = l13
	s3f32 = l11
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l3
	s1f32 = l12
	s2f32 = l16
	s3f32 = l12
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l3
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl5:
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s1i32 = 12
	s0i32 = i32DivS(s0i32, s1i32)
	l8 = s0i32
	s0i32 = 0
	l3 = s0i32
	s0i32 = l1
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl9:
		s0i32 = l5
		s1i32 = l3
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l13 = s0f32
		s0i32 = l0
		s1i32 = l3
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1f32 = 1
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1f32 = s1f32 / s2f32
		l10 = s1f32
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l13
		s2f32 = l10
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l8
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
	}
	s0i32 = l8
	return s0i32
}
