package internal

import (
	"math"
	"math/bits"
	"unsafe"
)

func f627(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int64
	_ = l18
	var l19 int64
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
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
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
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
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	var s8f32 float32
	_ = s8f32
	var s9f32 float32
	_ = s9f32
	s0i32 = ctx.g0
	s1i32 = 1376
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 232
		s0i32 = s0i32 + s1i32
		s0i32 = f152(ctx, s0i32)
		l10 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+84)])
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			s2i32 = l5
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
			s0i32 = f28(ctx, s0i32, s1i32, s2i32)
			l9 = s0i32
			s0i32 = l5
			s1i32 = 0
			ctx.Mem[int(s0i32+84)] = uint8(s1i32)
			s0i32 = l5
			s1i32 = l9
			ctx.Mem[int(s0i32+85)] = uint8(s1i32)
		}
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l20 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l21 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l22 = s0f32
		s0i32 = l4
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s1f32 = float32(math.Floor(float64(s1f32)))
		l23 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l23
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l23 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l23
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l23 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l23
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl2
		}
		s1i32 = -2147483648
	lbl2:
		l5 = s1i32
		s2i32 = -2147483645
		s3i32 = l5
		s4i32 = -2147483645
		if s3i32 > s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = -2
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint32(s1i32)
		s0i32 = l4
		s1f32 = l22
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l22 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l22
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l22 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l22
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l22 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l22
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl4
		}
		s1i32 = -2147483648
	lbl4:
		l5 = s1i32
		s2i32 = 2147483645
		s3i32 = l5
		s4i32 = 2147483645
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = 2
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+156)])) = uint32(s1i32)
		s0i32 = l4
		s1f32 = l21
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l21 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l21
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l21 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l21
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l21 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l21
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl6
		}
		s1i32 = -2147483648
	lbl6:
		l5 = s1i32
		s2i32 = 2147483645
		s3i32 = l5
		s4i32 = 2147483645
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = 2
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint32(s1i32)
		s0i32 = l4
		s1f32 = l20
		s1f32 = float32(math.Floor(float64(s1f32)))
		l20 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l20
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l20 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l20
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l20 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l20
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl8
		}
		s1i32 = -2147483648
	lbl8:
		l5 = s1i32
		s2i32 = -2147483645
		s3i32 = l5
		s4i32 = -2147483645
		if s3i32 > s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = -2
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l1
		s2i32 = l1
		s3i32 = 20
		s2i32 = s2i32 + s3i32
		s3i32 = l1
		s3i32 = int32(ctx.Mem[int(s3i32+40)])
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l4
		s3i32 = 144
		s2i32 = s2i32 + s3i32
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l1
		s1i32 = l4
		s2i32 = 144
		s1i32 = s1i32 + s2i32
		s0i32 = f153(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			s0i32 = 0
			goto lbl11
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+40)])
		if s1i32 != 0 {
			goto lbl13
		}
		s0i32 = l10
		s1i32 = l1
		s2i32 = l2
		f150(ctx, s0i32, s1i32, s2i32)
		s0i32 = l10
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1104)]))
		l2 = s0i32
		s0i32 = l10
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1100)]))
	lbl13:
		l8 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l9 = s0i32
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l6 = s0i32
		s0i32 = l4
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f32 = float32(s1i32)
		l20 = s1f32
		s2f32 = 1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)])) = s1f32
		s0i32 = l4
		s1i32 = l6
		s1f32 = float32(s1i32)
		l21 = s1f32
		s2f32 = 1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)])) = s1f32
		s0i32 = l4
		s1i32 = l9
		s1f32 = float32(s1i32)
		l22 = s1f32
		s2f32 = -1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = s1f32
		s0i32 = l4
		s1i32 = l5
		s1f32 = float32(s1i32)
		l23 = s1f32
		s2f32 = -1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = s1f32
		s0i32 = l4
		s1f32 = l20
		s2f32 = -1
		s1f32 = s1f32 + s2f32
		l20 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+228)])) = s1f32
		s0i32 = l4
		s1f32 = l22
		s2f32 = 1
		s1f32 = s1f32 + s2f32
		l22 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+220)])) = s1f32
		s0i32 = l4
		s1f32 = l23
		s2f32 = 1
		s1f32 = s1f32 + s2f32
		l23 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+216)])) = s1f32
		s0i32 = l4
		s1f32 = l21
		s2f32 = -1
		s1f32 = s1f32 + s2f32
		l21 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = s1f32
		s0f32 = l23
		s1f32 = l21
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
		s1i32 = 0
		s2f32 = l22
		s3f32 = l20
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = 1
		s2i32 = s2i32 ^ s3i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = uint64(s1i64)
			s0i32 = l4
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+216)])) = uint64(s1i64)
		}
		s0i32 = l4
		s1i32 = 200
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l4
		s1i32 = 216
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		s2i32 = l1
		s2i32 = int32(ctx.Mem[int(s2i32+42)])
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
	lbl11:
		l13 = s0i32
		s0i32 = l4
		s1i32 = 184
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l1
		l6 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		f275(ctx, s0i32, s1i32)
		s0i32 = l4
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l4
		s2i32 = 4
		s1i32 = s1i32 | s2i32
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 160
		s0i32 = s0i32 + s1i32
		l14 = s0i32
		s0i32 = l4
		s1i32 = 168
		s0i32 = s0i32 + s1i32
		l16 = s0i32
		s0i32 = l4
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 | s1i32
		l17 = s0i32
		s0i32 = 6
		l9 = s0i32
	lbl16:
		s0i32 = l9
		l1 = s0i32
		s0i32 = l6
		s1i32 = l4
		s2i32 = 144
		s1i32 = s1i32 + s2i32
		s0i32 = f276(ctx, s0i32, s1i32)
		l9 = s0i32
		s1i32 = 6
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l9
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl19
		case 1:
			goto lbl23
		case 2:
			goto lbl22
		case 3:
			goto lbl21
		case 4:
			goto lbl20
		case 5:
			goto lbl17
		default:
			goto lbl24
		}
	lbl24:
		s0i32 = l4
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)]))
		l18 = s0i64
		l19 = s0i64
		goto lbl16
	lbl23:
		s0i32 = l1
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l0 = s1i32
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l0
			s1i32 = int32(ctx.Mem[int(s1i32+0)])
		} else {
			s1i32 = 6
		}
		s2i32 = l4
		s3i32 = 144
		s2i32 = s2i32 + s3i32
		s3i32 = 3
		f262(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)]))
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+164)]))
		s0f32 = s0f32 + s1f32
		s1f32 = 0.5
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+156)]))
		s0f32 = s0f32 - s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s0f32 = float32(math.Ceil(float64(s0f32)))
		l20 = s0f32
		s1f32 = 2.1474835e+09
		s2f32 = l20
		s3f32 = 2.1474835e+09
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l20 = s0f32
		s1f32 = -2.1474835e+09
		s2f32 = l20
		s3f32 = -2.1474835e+09
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l20 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l20
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl26
		}
		s0i32 = -2147483648
	lbl26:
		l0 = s0i32
		s0i32 = l4
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = l8
		s2i32 = l13
		s3i32 = l12
		s4i32 = l2
		s5i32 = 33
		s6i32 = l4
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+144)]))
		s7i32 = l4
		s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+160)]))
		s6f32 = s6f32 + s7f32
		s7f32 = 0.5
		s6f32 = s6f32 * s7f32
		s7i32 = l4
		s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+152)]))
		s6f32 = s6f32 - s7f32
		s6f32 = float32(math.Abs(float64(s6f32)))
		s6f32 = float32(math.Ceil(float64(s6f32)))
		l20 = s6f32
		s7f32 = 2.1474835e+09
		s8f32 = l20
		s9f32 = 2.1474835e+09
		if s8f32 < s9f32 {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		if s8i32 != 0 {
			// s6f32 = s6f32
		} else {
			s6f32 = s7f32
		}
		l20 = s6f32
		s7f32 = -2.1474835e+09
		s8f32 = l20
		s9f32 = -2.1474835e+09
		if s8f32 > s9f32 {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		if s8i32 != 0 {
			// s6f32 = s6f32
		} else {
			s6f32 = s7f32
		}
		l20 = s6f32
		s6f32 = float32(math.Abs(float64(s6f32)))
		s7f32 = 2.1474836e+09
		if s6f32 < s7f32 {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			s6f32 = l20
			s6i32 = int32(math.Trunc(float64(s6f32)))
			goto lbl28
		}
		s6i32 = -2147483648
	lbl28:
		l7 = s6i32
		s7i32 = l0
		s8i32 = 1
		s7i32 = int32(uint32(s7i32) >> (uint32(s8i32) & 31))
		s6i32 = s6i32 + s7i32
		s7i32 = l7
		s8i32 = 1
		s7i32 = int32(uint32(s7i32) >> (uint32(s8i32) & 31))
		s8i32 = l0
		s7i32 = s7i32 + s8i32
		s8i32 = l7
		s9i32 = l0
		if uint32(s8i32) > uint32(s9i32) {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		if s8i32 != 0 {
			// s6i32 = s6i32
		} else {
			s6i32 = s7i32
		}
		s6i32 = int32(bits.LeadingZeros32(uint32(s6i32)))
		s5i32 = s5i32 - s6i32
		s6i32 = 1
		s5i32 = s5i32 >> (uint32(s6i32) & 31)
		l0 = s5i32
		s6i32 = 5
		s7i32 = l0
		s8i32 = 5
		if s7i32 < s8i32 {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i32 = s5i32
		} else {
			s5i32 = s6i32
		}
		s6i32 = l3
		f216(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		s0i32 = l14
		goto lbl18
	lbl22:
		s0i32 = l1
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l0 = s1i32
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l0
			s1i32 = int32(ctx.Mem[int(s1i32+0)])
		} else {
			s1i32 = 6
		}
		s2i32 = l4
		s3i32 = 144
		s2i32 = s2i32 + s3i32
		s3i32 = 3
		f262(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		s0i32 = l4
		s1i32 = 1360
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+160)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = 1352
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+152)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1344)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1368)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		s2i32 = l4
		s3i32 = 1344
		s2i32 = s2i32 + s3i32
		s2i32 = f233(ctx, s2i32)
		l7 = s2i32
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l0 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l11 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l11
			f13(ctx, s0i32)
		}
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 1
		s0i32 = s0i32 | s1i32
		l0 = s0i32
		s1i32 = 18
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l0
			s2i32 = 8
			s1i32 = f50(ctx, s1i32, s2i32)
			l0 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			goto lbl32
		}
		s0i32 = l4
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		l0 = s0i32
	lbl32:
		s0i32 = l4
		s1i32 = l4
		s2i32 = 1344
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s3i32 = l7
		s1i32 = f232(ctx, s1i32, s2i32, s3i32)
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
		s0i32 = 0
		l11 = s0i32
		s0i32 = l14
		s1i32 = l7
		s2i32 = 1
		if s1i32 < s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl18
		}
	lbl34:
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s0f32 = s0f32 + s1f32
		s1f32 = 0.5
		s0f32 = s0f32 * s1f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s0f32 = s0f32 - s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s0f32 = float32(math.Ceil(float64(s0f32)))
		l20 = s0f32
		s1f32 = 2.1474835e+09
		s2f32 = l20
		s3f32 = 2.1474835e+09
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l20 = s0f32
		s1f32 = -2.1474835e+09
		s2f32 = l20
		s3f32 = -2.1474835e+09
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l20 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l20
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl35
		}
		s0i32 = -2147483648
	lbl35:
		l7 = s0i32
		s0i32 = l0
		s1i32 = l8
		s2i32 = l13
		s3i32 = l12
		s4i32 = l2
		s5i32 = 33
		s6i32 = l0
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s7i32 = l0
		s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+16)]))
		s6f32 = s6f32 + s7f32
		s7f32 = 0.5
		s6f32 = s6f32 * s7f32
		s7i32 = l0
		s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+8)]))
		s6f32 = s6f32 - s7f32
		s6f32 = float32(math.Abs(float64(s6f32)))
		s6f32 = float32(math.Ceil(float64(s6f32)))
		l20 = s6f32
		s7f32 = 2.1474835e+09
		s8f32 = l20
		s9f32 = 2.1474835e+09
		if s8f32 < s9f32 {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		if s8i32 != 0 {
			// s6f32 = s6f32
		} else {
			s6f32 = s7f32
		}
		l20 = s6f32
		s7f32 = -2.1474835e+09
		s8f32 = l20
		s9f32 = -2.1474835e+09
		if s8f32 > s9f32 {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		if s8i32 != 0 {
			// s6f32 = s6f32
		} else {
			s6f32 = s7f32
		}
		l20 = s6f32
		s6f32 = float32(math.Abs(float64(s6f32)))
		s7f32 = 2.1474836e+09
		if s6f32 < s7f32 {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			s6f32 = l20
			s6i32 = int32(math.Trunc(float64(s6f32)))
			goto lbl37
		}
		s6i32 = -2147483648
	lbl37:
		l15 = s6i32
		s7i32 = l7
		s8i32 = 1
		s7i32 = int32(uint32(s7i32) >> (uint32(s8i32) & 31))
		s6i32 = s6i32 + s7i32
		s7i32 = l15
		s8i32 = 1
		s7i32 = int32(uint32(s7i32) >> (uint32(s8i32) & 31))
		s8i32 = l7
		s7i32 = s7i32 + s8i32
		s8i32 = l15
		s9i32 = l7
		if uint32(s8i32) > uint32(s9i32) {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		if s8i32 != 0 {
			// s6i32 = s6i32
		} else {
			s6i32 = s7i32
		}
		s6i32 = int32(bits.LeadingZeros32(uint32(s6i32)))
		s5i32 = s5i32 - s6i32
		s6i32 = 1
		s5i32 = s5i32 >> (uint32(s6i32) & 31)
		l7 = s5i32
		s6i32 = 5
		s7i32 = l7
		s8i32 = 5
		if s7i32 < s8i32 {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i32 = s5i32
		} else {
			s5i32 = s6i32
		}
		s6i32 = l3
		f216(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		s0i32 = l0
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l11
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl34
		}
		s0i32 = l14
		goto lbl18
	lbl21:
		s0i32 = l1
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l0 = s1i32
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l0
			s1i32 = int32(ctx.Mem[int(s1i32+0)])
		} else {
			s1i32 = 6
		}
		s2i32 = l4
		s3i32 = 144
		s2i32 = s2i32 + s3i32
		s3i32 = 4
		f262(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l4
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = l8
		s2i32 = l13
		s3i32 = l12
		s4i32 = l2
		s5i32 = l3
		f411(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = l16
		goto lbl18
	lbl20:
		s0i32 = l4
		s1i64 = l18
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = l19
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l0 = s1i32
			s2i32 = l6
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
			if uint32(s1i32) < uint32(s2i32) {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1i32 = l0
				s1i32 = int32(ctx.Mem[int(s1i32+0)])
			} else {
				s1i32 = 6
			}
			s2i32 = l4
			s3i32 = 144
			s2i32 = s2i32 + s3i32
			s3i32 = 2
			f262(ctx, s0i32, s1i32, s2i32, s3i32)
		}
		s0i32 = l4
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		s2i32 = l8
		s3i32 = l2
		s4i32 = l3
		if int(s4i32) < 0 || int(s4i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s4i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s4i32].numParams != 4 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl16
	lbl19:
		s0i32 = l1
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l0 = s1i32
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l0
			s1i32 = int32(ctx.Mem[int(s1i32+0)])
		} else {
			s1i32 = 6
		}
		s2i32 = l4
		s3i32 = 144
		s2i32 = s2i32 + s3i32
		s3i32 = 2
		f262(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l4
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		s2i32 = l8
		s3i32 = l2
		s4i32 = l3
		if int(s4i32) < 0 || int(s4i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s4i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s4i32].numParams != 4 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l17
	lbl18:
		l0 = s0i32
		s0i64 = l18
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)]))
		s2i64 = l18
		s3i32 = l9
		s4i32 = -1
		s3i32 = s3i32 + s4i32
		s4i32 = 4
		if uint32(s3i32) < uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		s2i32 = l1
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		l18 = s0i64
		s0i32 = l0
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i64
		goto lbl16
	lbl17:
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		s1i32 = l5
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l0
		f13(ctx, s0i32)
	lbl10:
		s0i32 = l10
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		s0i32 = f74(ctx, s0i32)
		s0i32 = l10
		f40(ctx, s0i32)
	}
	s0i32 = l4
	s1i32 = 1376
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
