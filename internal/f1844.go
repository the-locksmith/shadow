package internal

import (
	"math"
	"unsafe"
)

func f1844(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int64
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+536)]))
	l3 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s1i32 = 28
	s0i32 = s0i32 | s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+540)]))
	s2i32 = l3
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 532
		s0i32 = s0i32 + s1i32
		s1i32 = 28
		s2i32 = 4
		f18(ctx, s0i32, s1i32, s2i32)
		s0i32 = 0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+536)]))
		l3 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l2 = s0i32
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s2i32 = 28
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+536)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = 1
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+564)]))
	s3i32 = 6
	s2i32 = s2i32 + s3i32
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s1f32 = float32(s1i32)
	l13 = s1f32
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l14
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl1
	}
	s0i32 = -2147483648
lbl1:
	l2 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l14
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl3
	}
	s0i32 = -2147483648
lbl3:
	l5 = s0i32
	s1i32 = l2
	s2i32 = l5
	s3i32 = l2
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l4 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l6 = s0i32
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = 6
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l9 = s0i32
	s0i32 = l2
	s1i32 = l5
	s2i32 = l4
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l8 = s0i32
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s1i32 = 6
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l2 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l14
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl5
	}
	s0i32 = -2147483648
lbl5:
	l5 = s0i32
	s0i32 = l2
	s1i32 = l9
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l11 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl7
	}
	s0i32 = -2147483648
lbl7:
	l1 = s0i32
	s0i32 = l11
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l6
	s1i32 = l8
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s0i32 = l5
	s1i32 = l1
	s2i32 = l4
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l1
	s2i32 = l5
	s3i32 = l4
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l11 = s1i32
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s1i32 = 32768
	s0i32 = s0i32 + s1i32
	s1i32 = 65535
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 16
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l6
		s0i32 = i32DivS(s0i32, s1i32)
		goto lbl10
	}
	s0i32 = l7
	s1i32 = l1
	s1i64 = int64(s1i32)
	s2i64 = 16
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	s2i32 = l6
	s2i64 = int64(s2i32)
	s1i64 = i64DivS(s1i64, s2i64)
	l12 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = -2147483647
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 2147483647
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 24
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i64 = l12
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i64 = l12
	s3i64 = -2147483647
	if s2i64 < s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
lbl10:
	l1 = s0i32
	s0i32 = l3
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint16(s1i32)
	s0i32 = l3
	s1i32 = l9
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = -1
	s2i32 = 1
	s3i32 = l4
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l6 = s1i32
	ctx.Mem[int(s0i32+27)] = uint8(s1i32)
	s0i32 = l3
	s1i32 = l11
	s2i32 = l1
	s2i64 = int64(s2i32)
	s3i32 = l10
	s4i32 = -64
	s3i32 = s3i32 & s4i32
	s4i32 = 32
	s3i32 = s3i32 | s4i32
	s4i32 = l8
	s3i32 = s3i32 - s4i32
	s3i64 = int64(s3i32)
	s2i64 = s2i64 * s3i64
	s3i64 = 16
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	s2i32 = int32(s2i64)
	s1i32 = s1i32 + s2i32
	s2i32 = 10
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l1
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l10 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l4
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l4 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+27)])
	s1i32 = l6
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l9
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			goto lbl9
		}
		s0i32 = l2
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		s0i32 = l1
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		goto lbl9
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l8 = s0i32
	s0i32 = l2
	s1i32 = l4
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l8
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l9
		s1i32 = l8
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l9
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			goto lbl9
		}
		s0i32 = l1
		s1i32 = l6
		ctx.Mem[int(s0i32+27)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l8
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		goto lbl9
	}
	s0i32 = l5
	s1i32 = l8
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l2
	s1i32 = l4
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		goto lbl9
	}
	s0i32 = l1
	s1i32 = l6
	ctx.Mem[int(s0i32+27)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l4
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	goto lbl9
lbl15:
	s0i32 = l0
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	goto lbl9
lbl12:
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s0i32 = f217(ctx, s0i32)
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl9:
	s0i32 = l7
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
