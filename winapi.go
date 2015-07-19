// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winapi

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go winapi.go

type MEMORYSTATUSEX struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}

//sys	GlobalMemoryStatusEx(buf *MEMORYSTATUSEX) (err error) = kernel32.GlobalMemoryStatusEx
