package internal

import (
	"unsafe"
)

func f736(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	lbl2:
		s0i32 = l3
		s1i32 = 8
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l1
			s1i32 = 32
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l3
			s1i32 = 8
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l0 = s0i32
			s0i32 = l3
			s1i32 = -8
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l0
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl2
			}
			goto lbl0
		}
		s0i32 = l3
		s1i32 = 4
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = -4
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l1
			s1i32 = 16
			s0i32 = s0i32 + s1i32
			l1 = s0i32
		}
		s0i32 = l3
		s1i32 = 2
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = -2
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l1
			s1i32 = 8
			s0i32 = s0i32 + s1i32
			l1 = s0i32
		}
		s0i32 = l3
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	s0i32 = l3
	s1i32 = l1
	s2i32 = l2
	s3i32 = l4
	s4i32 = 468
	f98(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl0:
}
