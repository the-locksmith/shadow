package internal

import (
	"unsafe"
)

func f1906(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	l3 = s0i32
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l2 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
}
