package internal

import (
	"math"
	"unsafe"
)

func f840(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
	var l24 float32
	_ = l24
	var l25 float32
	_ = l25
	var l26 float32
	_ = l26
	var l27 float32
	_ = l27
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
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l8
	s1f32 = l3
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	l4 = s1f32
	s2f32 = l4
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 - s2f32
	l4 = s1f32
	s2f32 = l4
	s1f32 = s1f32 * s2f32
	s2f32 = l4
	s3f32 = 0.3888889
	s2f32 = s2f32 * s3f32
	s3f32 = -0.33333334
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l8
	s1f32 = 1
	s2f32 = l4
	s1f32 = s1f32 - s2f32
	l5 = s1f32
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	s2f32 = l5
	s3f32 = 0.3888889
	s2f32 = s2f32 * s3f32
	s3f32 = -0.33333334
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	l20 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l8
	s1f32 = l4
	s2f32 = l4
	s3f32 = 1.5
	s4f32 = l4
	s5f32 = 1.1666666
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = 0.055555556
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l8
	s1f32 = l5
	s2f32 = l5
	s3f32 = 1.5
	s4f32 = l5
	s5f32 = 1.1666666
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = 0.055555556
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0f32 = 1.5
	s1f32 = l2
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	l4 = s1f32
	s2f32 = l4
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 - s2f32
	l4 = s1f32
	s2f32 = 1.1666666
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l19 = s0f32
	s0f32 = 1.5
	s1f32 = 1
	s2f32 = l4
	s1f32 = s1f32 - s2f32
	l5 = s1f32
	s2f32 = 1.1666666
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l21 = s0f32
	s0f32 = l2
	s1f32 = -1.5
	s0f32 = s0f32 + s1f32
	l26 = s0f32
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l23 = s0f32
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l22 = s0f32
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l2 = s1f32
	s0f32 = f14(ctx, s0f32, s1f32)
	l18 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l18
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l18
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl0
	}
	s0i32 = 0
lbl0:
	l13 = s0i32
	s0f32 = l4
	s1f32 = l19
	s0f32 = s0f32 * s1f32
	l19 = s0f32
	s0f32 = l5
	s1f32 = l21
	s0f32 = s0f32 * s1f32
	l21 = s0f32
	s0f32 = l22
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = l2
	s0f32 = f14(ctx, s0f32, s1f32)
	l18 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l18
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l18
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl2
	}
	s0i32 = 0
lbl2:
	l14 = s0i32
	s0f32 = l4
	s1f32 = 0.3888889
	s0f32 = s0f32 * s1f32
	l22 = s0f32
	s0f32 = l5
	s1f32 = 0.3888889
	s0f32 = s0f32 * s1f32
	l24 = s0f32
	s0f32 = l19
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l19 = s0f32
	s0f32 = l21
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l21 = s0f32
	s0f32 = l23
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = l2
	s0f32 = f14(ctx, s0f32, s1f32)
	l18 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l18
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l18
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl4
	}
	s0i32 = 0
lbl4:
	l15 = s0i32
	s0f32 = l22
	s1f32 = -0.33333334
	s0f32 = s0f32 + s1f32
	l18 = s0f32
	s0f32 = l4
	s1f32 = l4
	s0f32 = s0f32 * s1f32
	l23 = s0f32
	s0f32 = l24
	s1f32 = -0.33333334
	s0f32 = s0f32 + s1f32
	l22 = s0f32
	s0f32 = l5
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l24 = s0f32
	s0f32 = l4
	s1f32 = l19
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s0f32 = l5
	s1f32 = l21
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s0f32 = l26
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = l2
	s0f32 = f14(ctx, s0f32, s1f32)
	l2 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l2
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l2
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl6
	}
	s0i32 = 0
lbl6:
	l16 = s0i32
	s0f32 = l23
	s1f32 = l18
	s0f32 = s0f32 * s1f32
	l26 = s0f32
	s0f32 = l24
	s1f32 = l22
	s0f32 = s0f32 * s1f32
	l23 = s0f32
	s0f32 = l4
	s1f32 = 0.055555556
	s0f32 = s0f32 + s1f32
	l22 = s0f32
	s0f32 = l5
	s1f32 = 0.055555556
	s0f32 = s0f32 + s1f32
	l24 = s0f32
	s0f32 = l3
	s1f32 = -1.5
	s0f32 = s0f32 + s1f32
	l3 = s0f32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l27 = s0f32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0f32 = 0
	l4 = s0f32
	s0f32 = 0
	l5 = s0f32
	s0f32 = 0
	l18 = s0f32
lbl8:
	s0f32 = l24
	s1f32 = l20
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0f32 = l3
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = l27
	s0f32 = f14(ctx, s0f32, s1f32)
	l19 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l19
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l19
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl9
	}
	s0i32 = 0
lbl9:
	l7 = s0i32
	s0f32 = l26
	s1f32 = l20
	s0f32 = s0f32 * s1f32
	l19 = s0f32
	s1i32 = l6
	s2i32 = l7
	s3i32 = l17
	s2i32 = s2i32 * s3i32
	l7 = s2i32
	s3i32 = l13
	s2i32 = s2i32 + s3i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l10 = s1i32
	s2i32 = 24
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l22
	s2f32 = l20
	s1f32 = s1f32 * s2f32
	l21 = s1f32
	s2i32 = l6
	s3i32 = l7
	s4i32 = l14
	s3i32 = s3i32 + s4i32
	s4i32 = 2
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l11 = s2i32
	s3i32 = 24
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l2
	s3i32 = l6
	s4i32 = l7
	s5i32 = l15
	s4i32 = s4i32 + s5i32
	s5i32 = 2
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	s3i32 = s3i32 + s4i32
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l12 = s3i32
	s4i32 = 24
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s3f32 = float32(uint32(s3i32))
	s4f32 = 0.003921569
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = l5
	s4f32 = l23
	s5f32 = l20
	s4f32 = s4f32 * s5f32
	l20 = s4f32
	s5i32 = l6
	s6i32 = l7
	s7i32 = l16
	s6i32 = s6i32 + s7i32
	s7i32 = 2
	s6i32 = s6i32 << (uint32(s7i32) & 31)
	s5i32 = s5i32 + s6i32
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	l7 = s5i32
	s6i32 = 24
	s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
	s5f32 = float32(uint32(s5i32))
	s6f32 = 0.003921569
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l18
	s1f32 = l20
	s2i32 = l7
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l2
	s2i32 = l12
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l21
	s2i32 = l11
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l19
	s2i32 = l10
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l18 = s0f32
	s0f32 = l19
	s1i32 = l10
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l21
	s2i32 = l11
	s3i32 = 16
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l2
	s3i32 = l12
	s4i32 = 16
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 255
	s3i32 = s3i32 & s4i32
	s3f32 = float32(uint32(s3i32))
	s4f32 = 0.003921569
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = l4
	s4f32 = l20
	s5i32 = l7
	s6i32 = 16
	s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
	s6i32 = 255
	s5i32 = s5i32 & s6i32
	s5f32 = float32(uint32(s5i32))
	s6f32 = 0.003921569
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l25
	s1f32 = l20
	s2i32 = l7
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l2
	s2i32 = l12
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l21
	s2i32 = l11
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l19
	s2i32 = l10
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l25 = s0f32
	s0i32 = l9
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s1i32 = 4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l3
		s1f32 = 1
		s0f32 = s0f32 + s1f32
		l3 = s0f32
		s0i32 = l8
		s1i32 = l9
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l20 = s0f32
		goto lbl8
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l18
	s3f32 = l25
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
	s0i32 = l8
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
