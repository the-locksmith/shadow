package internal

import (
	"unsafe"
)

func f1556(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
	var l8 int64
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 0
		ctx.Mem[int(s0i32+32)] = uint8(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0i32 = s0i32 - s1i32
		s0i64 = int64(uint32(s0i32))
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1i32 = s1i32 - s2i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = 32
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s0i64 = s0i64 | s1i64
		l7 = s0i64
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l5 = s0i32
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		}
		s0i32 = l1
		s1i32 = 36
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l4
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		l8 = s0i64
		s0i32 = l4
		s1i64 = l7
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = l8
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l4
		f463(ctx, s0i32, s1i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l1
		f12(ctx, s0i32)
	lbl3:
		s0i32 = l6
		s1i32 = l3
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3i32 = l2
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s0i32 = f457(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 2
			ctx.Mem[int(s0i32+53)] = uint8(s1i32)
		}
		s0i32 = l0
		s1i32 = l3
		f253(ctx, s0i32, s1i32)
		s0i32 = l3
		s0i32 = f41(ctx, s0i32)
		goto lbl0
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+72)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l2 = s0i32
		goto lbl5
	}
	s0i32 = 0
	l2 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l3
	s1i32 = 1
	ctx.Mem[int(s0i32+53)] = uint8(s1i32)
lbl5:
	s0i32 = l0
	s1i32 = l1
	s2i32 = 36
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	f389(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = l4
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
