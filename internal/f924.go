package internal

import (
	"unsafe"
)

func f924(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = l2
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l8 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l11 = s1i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l9 = s2i32
	s3i32 = 16383
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l8
	s1i32 = l11
	s2i32 = l9
	s3i32 = 18
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = 256
	s1i32 = l9
	s2i32 = 14
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 15
	s1i32 = s1i32 & s2i32
	l9 = s1i32
	s2i32 = 4
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l11 = s1i32
	s0i32 = s0i32 - s1i32
	l15 = s0i32
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])))
	l5 = s0i32
	s1i32 = 255
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl2:
		s0i32 = l3
		s1i32 = l12
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l4 = s2i32
		s3i32 = 16
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 65532
		s2i32 = s2i32 & s3i32
		l10 = s2i32
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l6 = s1i32
		s2i32 = 16711935
		s1i32 = s1i32 & s2i32
		s2i32 = l11
		s3i32 = l4
		s4i32 = 14
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 15
		s3i32 = s3i32 & s4i32
		l7 = s3i32
		s4i32 = l9
		s3i32 = s3i32 * s4i32
		l0 = s3i32
		s2i32 = s2i32 - s3i32
		l13 = s2i32
		s1i32 = s1i32 * s2i32
		s2i32 = l7
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l7 = s2i32
		s3i32 = l0
		s2i32 = s2i32 - s3i32
		l14 = s2i32
		s3i32 = l8
		s4i32 = l4
		s5i32 = 16383
		s4i32 = s4i32 & s5i32
		s5i32 = 2
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		l4 = s4i32
		s3i32 = s3i32 + s4i32
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		l16 = s3i32
		s4i32 = 16711935
		s3i32 = s3i32 & s4i32
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = l12
		s2i32 = s2i32 + s3i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l4 = s2i32
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s3i32 = l0
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l15
		s3i32 = l7
		s2i32 = s2i32 - s3i32
		s3i32 = l0
		s2i32 = s2i32 + s3i32
		l7 = s2i32
		s3i32 = l8
		s4i32 = l10
		s3i32 = s3i32 + s4i32
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		l10 = s3i32
		s4i32 = 16711935
		s3i32 = s3i32 & s4i32
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 16711935
		s1i32 = s1i32 & s2i32
		s2i32 = l5
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 16711935
		s1i32 = s1i32 & s2i32
		s2i32 = l6
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s3i32 = l13
		s2i32 = s2i32 * s3i32
		s3i32 = l16
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 16711935
		s3i32 = s3i32 & s4i32
		s4i32 = l14
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = l7
		s4i32 = l10
		s5i32 = 8
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = 16711935
		s4i32 = s4i32 & s5i32
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 16711935
		s3i32 = s3i32 & s4i32
		s4i32 = l0
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s3i32 = l5
		s2i32 = s2i32 * s3i32
		s3i32 = -16711936
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l3
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l2
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l2
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl2
		}
		goto lbl0
		panic("unreachable executed")
		panic("unreachable executed")
	}
lbl3:
	s0i32 = l3
	s1i32 = l12
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l5 = s2i32
	s3i32 = 16
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 65532
	s2i32 = s2i32 & s3i32
	l4 = s2i32
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l10 = s1i32
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l11
	s3i32 = l5
	s4i32 = 14
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 15
	s3i32 = s3i32 & s4i32
	l6 = s3i32
	s4i32 = l9
	s3i32 = s3i32 * s4i32
	l0 = s3i32
	s2i32 = s2i32 - s3i32
	l7 = s2i32
	s1i32 = s1i32 * s2i32
	s2i32 = l6
	s3i32 = 4
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	l6 = s2i32
	s3i32 = l0
	s2i32 = s2i32 - s3i32
	l13 = s2i32
	s3i32 = l8
	s4i32 = l5
	s5i32 = 16383
	s4i32 = s4i32 & s5i32
	s5i32 = 2
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	l5 = s4i32
	s3i32 = s3i32 + s4i32
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l14 = s3i32
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = l12
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l5 = s2i32
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l0
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l15
	s3i32 = l6
	s2i32 = s2i32 - s3i32
	s3i32 = l0
	s2i32 = s2i32 + s3i32
	l6 = s2i32
	s3i32 = l4
	s4i32 = l8
	s3i32 = s3i32 + s4i32
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l4 = s3i32
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l10
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l7
	s2i32 = s2i32 * s3i32
	s3i32 = l14
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s4i32 = l13
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 + s3i32
	s3i32 = l6
	s4i32 = l4
	s5i32 = 8
	s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
	s5i32 = 16711935
	s4i32 = s4i32 & s5i32
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s4i32 = l0
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 + s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l2
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl3
	}
lbl0:
}
