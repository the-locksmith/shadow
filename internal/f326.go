package internal

import (
	"unsafe"
)

func f326(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s0i32 = f654(ctx, s0i32, s1i32, s2i32)
	goto lbl0
lbl1:
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l0
		f245(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+40)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		ctx.Mem[int(s0i32+41)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l4
		s1i32 = f92(ctx, s1i32)
		ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	}
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	f359(ctx, s0i32, s1i32, s2i32, s3i32)
lbl0:
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+40)])
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		s3i32 = -1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		ctx.Mem[int(s1i32+41)] = uint8(s2i32)
		goto lbl5
	}
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
	l1 = s2i32
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	ctx.Mem[int(s1i32+41)] = uint8(s2i32)
	s1i32 = l1
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1i32 = f92(ctx, s1i32)
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l1
		s1i32 = f81(ctx, s1i32, s2i32)
		s1i32 = l1
		f143(ctx, s1i32)
		s1i32 = l0
		s2i32 = 1
		ctx.Mem[int(s1i32+40)] = uint8(s2i32)
		goto lbl5
	}
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+40)])
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
lbl5:
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	goto lbl3
lbl4:
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	s1i32 = f92(ctx, s1i32)
lbl3:
	ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
}
