package internal

import (
	"unsafe"
)

func f148(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	s1i32 = l0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l6 = s0i32
	s1i32 = l3
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l6
		s2i32 = l2
		s3i32 = l1
		s4i32 = -1
		s3i32 = s3i32 ^ s4i32
		s4i32 = l3
		s3i32 = s3i32 + s4i32
		s4i32 = l4
		f414(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		return
	}
	s0i32 = l1
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l6
		s2i32 = l2
		s3i32 = 256
		s4i32 = l1
		s3i32 = s3i32 - s4i32
		s4i32 = l4
		f414(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
	}
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l8 = s0i32
	s1i32 = l6
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l1 = s0i32
	s1i32 = l2
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l1
		s2i32 = l6
		s3i32 = l7
		s4i32 = l0
		s5i32 = -1
		s4i32 = s4i32 ^ s5i32
		s5i32 = l2
		s4i32 = s4i32 + s5i32
		s5i32 = 255
		s4i32 = s4i32 & s5i32
		s5i32 = l4
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		goto lbl3
	}
	s0i32 = l0
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l1
		s2i32 = l6
		s3i32 = l7
		s4i32 = 0
		s5i32 = l0
		s4i32 = s4i32 - s5i32
		s5i32 = 255
		s4i32 = s4i32 & s5i32
		s5i32 = l4
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	}
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l9 = s0i32
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l4
	s1i32 = l1
	s2i32 = l6
	s3i32 = l10
	s4i32 = l7
	s5i32 = l4
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl6:
	s0i32 = l2
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l4
	s1i32 = l9
	s2i32 = l6
	s3i32 = l7
	s4i32 = l1
	s5i32 = l4
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl3:
	s0i32 = l3
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l8
	s2i32 = l2
	s3i32 = l1
	s4i32 = l4
	f414(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl0:
}
