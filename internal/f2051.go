package internal

import (
	"unsafe"
)

func f2051(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
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
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l4 = s0i32
		s1i32 = 16
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l4
		s0i32 = s0i32 | s1i32
		s1i32 = 132184095
		s0i32 = s0i32 & s1i32
		s1i32 = l1
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		l4 = s1i32
		s2i32 = 16
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l4
		s1i32 = s1i32 | s2i32
		s2i32 = 132184095
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		l4 = s1i32
		s2i32 = 16
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l4
		s1i32 = s1i32 | s2i32
		s2i32 = 132184095
		s1i32 = s1i32 & s2i32
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
		l5 = s1i32
		s2i32 = 16
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l5
		s1i32 = s1i32 | s2i32
		s2i32 = 132184095
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)])))
		l5 = s2i32
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l5
		s2i32 = s2i32 | s3i32
		s3i32 = 132184095
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)])))
		l5 = s2i32
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l5
		s2i32 = s2i32 | s3i32
		s3i32 = 132184095
		s2i32 = s2i32 & s3i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s2i32 = l4
		s3i32 = l2
		s3i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+2)])))
		l4 = s3i32
		s4i32 = 16
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = l4
		s3i32 = s3i32 | s4i32
		s4i32 = 132184095
		s3i32 = s3i32 & s4i32
		s4i32 = l1
		s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+2)])))
		l4 = s4i32
		s5i32 = 16
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = l4
		s4i32 = s4i32 | s5i32
		s5i32 = 132184095
		s4i32 = s4i32 & s5i32
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+2)])))
		l4 = s4i32
		s5i32 = 16
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = l4
		s4i32 = s4i32 | s5i32
		s5i32 = 132184095
		s4i32 = s4i32 & s5i32
		s5i32 = 1
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s3i32 = s3i32 + s4i32
		s4i32 = 1
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s2i32 = 20
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 2016
		s1i32 = s1i32 & s2i32
		s2i32 = l4
		s3i32 = 4
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 63519
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l6
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l5
		l4 = s0i32
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
