package main

import (
	"syscall"
	"unsafe"
)

//go:wasmimport wasi_snapshot_preview1 sock_recv
//go:noescape
func sock_recv(fd int32, ri_data unsafe.Pointer, ri_data_len int32, ri_flags int32, ro_data_len int32, ro_flags unsafe.Pointer) syscall.Errno

//go:wasmimport wasi_snapshot_preview1 sock_send
//go:noescape
func sock_send(fd int32, si_data unsafe.Pointer, si_data_len int32, si_flags int32, ret_data_len int32) syscall.Errno
