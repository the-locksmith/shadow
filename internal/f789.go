package internal

import (
	"math"
	"unsafe"
)

func f789(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	var s7f32 float32
	_ = s7f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l12 = s0f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l13 = s1f32
	s2f32 = l4
	s2i32 = int32(math.Float32bits(s2f32))
	l8 = s2i32
	s3i32 = 2147483647
	s2i32 = s2i32 & s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	l4 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	l6 = s1i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 1.1920929e-07
	s1f32 = s1f32 * s2f32
	s2f32 = -124.22552
	s1f32 = s1f32 + s2f32
	s2i32 = l6
	s3i32 = 8388607
	s2i32 = s2i32 & s3i32
	s3i32 = 1056964608
	s2i32 = s2i32 | s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l10 = s2f32
	s3f32 = 1.4980303
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = 1.72588
	s3f32 = l10
	s4f32 = 0.35208872
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l10 = s0f32
	s1f32 = 121.274055
	s0f32 = s0f32 + s1f32
	s1f32 = l10
	s2f32 = l10
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 - s2f32
	l10 = s1f32
	s2f32 = 1.4901291
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = 27.728024
	s2f32 = 4.8425255
	s3f32 = l10
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 8.388608e+06
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l10
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l10
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl0
	}
	s0i32 = 0
lbl0:
	s0f32 = math.Float32frombits(uint32(s0i32))
	s1f32 = l4
	s2f32 = l4
	s3f32 = 1
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = l4
	s2f32 = l4
	s3f32 = 0
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l15 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l10 = s0f32
	s1f32 = l14
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l14 = s2f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 1.442695
	s0f32 = s0f32 * s1f32
	l11 = s0f32
	s1f32 = 121.274055
	s0f32 = s0f32 + s1f32
	s1f32 = l11
	s2f32 = l11
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	s2f32 = 1.4901291
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = 27.728024
	s2f32 = 4.8425255
	s3f32 = l11
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 8.388608e+06
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l11 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l11
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l11
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl2
	}
	s0i32 = 0
lbl2:
	l6 = s0i32
	s0f32 = l15
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l11 = s1f32
	s2i32 = l6
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = 1
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = l8
	s2i32 = -2147483648
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 | s1i32
	l8 = s0i32
	s0f32 = l12
	s1f32 = l13
	s2f32 = l3
	s2i32 = int32(math.Float32bits(s2f32))
	l7 = s2i32
	s3i32 = 2147483647
	s2i32 = s2i32 & s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l15 = s2f32
	s1f32 = s1f32 * s2f32
	l3 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	l6 = s1i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 1.1920929e-07
	s1f32 = s1f32 * s2f32
	s2f32 = -124.22552
	s1f32 = s1f32 + s2f32
	s2i32 = l6
	s3i32 = 8388607
	s2i32 = s2i32 & s3i32
	s3i32 = 1056964608
	s2i32 = s2i32 | s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l4 = s2f32
	s3f32 = 1.4980303
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = 1.72588
	s3f32 = l4
	s4f32 = 0.35208872
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s1f32 = 121.274055
	s0f32 = s0f32 + s1f32
	s1f32 = l4
	s2f32 = l4
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 - s2f32
	l4 = s1f32
	s2f32 = 1.4901291
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = 27.728024
	s2f32 = 4.8425255
	s3f32 = l4
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 8.388608e+06
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
		goto lbl4
	}
	s0i32 = 0
lbl4:
	s0f32 = math.Float32frombits(uint32(s0i32))
	s1f32 = l3
	s2f32 = l3
	s3f32 = 1
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = l3
	s2f32 = l3
	s3f32 = 0
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = l11
	s2f32 = l10
	s3f32 = l15
	s4f32 = l14
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = 1.442695
	s2f32 = s2f32 * s3f32
	l4 = s2f32
	s3f32 = 121.274055
	s2f32 = s2f32 + s3f32
	s3f32 = l4
	s4f32 = l4
	s4f32 = float32(math.Floor(float64(s4f32)))
	s3f32 = s3f32 - s4f32
	l4 = s3f32
	s4f32 = 1.4901291
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	s3f32 = 27.728024
	s4f32 = 4.8425255
	s5f32 = l4
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 / s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = 8.388608e+06
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	l4 = s2f32
	s3f32 = 4.2949673e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l4
	s4f32 = 0
	if s3f32 >= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		s2f32 = l4
		s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
		goto lbl6
	}
	s2i32 = 0
lbl6:
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s3f32 = 1
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = l7
	s2i32 = -2147483648
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 | s1i32
	l7 = s0i32
	s0f32 = l12
	s1f32 = l13
	s2f32 = l2
	s2i32 = int32(math.Float32bits(s2f32))
	l9 = s2i32
	s3i32 = 2147483647
	s2i32 = s2i32 & s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l4 = s2f32
	s1f32 = s1f32 * s2f32
	l2 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	l6 = s1i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 1.1920929e-07
	s1f32 = s1f32 * s2f32
	s2f32 = -124.22552
	s1f32 = s1f32 + s2f32
	s2i32 = l6
	s3i32 = 8388607
	s2i32 = s2i32 & s3i32
	s3i32 = 1056964608
	s2i32 = s2i32 | s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l3 = s2f32
	s3f32 = 1.4980303
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = 1.72588
	s3f32 = l3
	s4f32 = 0.35208872
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s1f32 = 121.274055
	s0f32 = s0f32 + s1f32
	s1f32 = l3
	s2f32 = l3
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 - s2f32
	l3 = s1f32
	s2f32 = 1.4901291
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = 27.728024
	s2f32 = 4.8425255
	s3f32 = l3
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 8.388608e+06
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
		goto lbl8
	}
	s0i32 = 0
lbl8:
	l6 = s0i32
	s0i32 = l8
	s0f32 = math.Float32frombits(uint32(s0i32))
	l12 = s0f32
	s0i32 = l7
	s0f32 = math.Float32frombits(uint32(s0i32))
	l13 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = l2
	s4f32 = l2
	s5f32 = 1
	if s4f32 != s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s3f32 = l2
	s4f32 = l2
	s5f32 = 0
	if s4f32 != s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s3f32 = l11
	s4f32 = l10
	s5f32 = l4
	s6f32 = l14
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = 1.442695
	s4f32 = s4f32 * s5f32
	l3 = s4f32
	s5f32 = 121.274055
	s4f32 = s4f32 + s5f32
	s5f32 = l3
	s6f32 = l3
	s6f32 = float32(math.Floor(float64(s6f32)))
	s5f32 = s5f32 - s6f32
	l3 = s5f32
	s6f32 = 1.4901291
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 - s5f32
	s5f32 = 27.728024
	s6f32 = 4.8425255
	s7f32 = l3
	s6f32 = s6f32 - s7f32
	s5f32 = s5f32 / s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = 8.388608e+06
	s4f32 = s4f32 * s5f32
	s5f32 = 0.5
	s4f32 = s4f32 + s5f32
	l3 = s4f32
	s5f32 = 4.2949673e+09
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s5f32 = l3
	s6f32 = 0
	if s5f32 >= s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	s4i32 = s4i32 & s5i32
	if s4i32 != 0 {
		s4f32 = l3
		s4i32 = int32(uint32(math.Trunc(float64(s4f32))))
		goto lbl10
	}
	s4i32 = 0
lbl10:
	s4f32 = math.Float32frombits(uint32(s4i32))
	s3f32 = s3f32 + s4f32
	s4f32 = l2
	s5f32 = 1
	if s4f32 <= s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s2i32 = int32(math.Float32bits(s2f32))
	s3i32 = l9
	s4i32 = -2147483648
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 | s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = l13
	s4f32 = l12
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
