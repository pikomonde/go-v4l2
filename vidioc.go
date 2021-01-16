package v4l2

import (
	"github.com/reiver/go-v4l2/framesize"
	"github.com/reiver/go-v4l2/format"

	"unsafe"
)

// ioctl codes for V4L2 (Video4Linux version 2) video devices.
const (
	// A Golang conversion of the following C code:
	//
	// #define  VIDIOC_QUERYCAP  _IOR('V',  0, struct v4l2_capability)
	CONST_VIDIOC_QUERYCAP = (CONST_IOC_READ                      << CONST_IOC_DIRSHIFT)  |
	                        (uintptr('V')                        << CONST_IOC_TYPESHIFT) |
	                        (0                                   << CONST_IOC_NRSHIFT)   |
	                        (unsafe.Sizeof(internalCapability{}) << CONST_IOC_SIZESHIFT)

	// A Golang conversion of the following C code:
	//
	// #define  VIDIOC_RESERVED   _IO('V',  1)
	CONST_VIDIOC_RESERVED = (CONST_IOC_NONE                      << CONST_IOC_DIRSHIFT)  |
	                        (uintptr('V')                        << CONST_IOC_TYPESHIFT) |
	                        (1                                   << CONST_IOC_NRSHIFT)   |
	                        (0                                   << CONST_IOC_SIZESHIFT)

	// A Golang conversion of the following C code:
	//
	// #define  VIDIOC_ENUM_FMT _IOWR('V',  2, struct v4l2_fmtdesc)
	CONST_VIDIOC_ENUM_FMT = ((CONST_IOC_READ | CONST_IOC_WRITE)     << CONST_IOC_DIRSHIFT)  |
	                        (uintptr('V')                           << CONST_IOC_TYPESHIFT) |
	                        (2                                      << CONST_IOC_NRSHIFT)   |
	                        (unsafe.Sizeof(internalFormatFamily{})  << CONST_IOC_SIZESHIFT)

	// A Golang conversion of the following C code:
	//
	// #define VIDIOC_S_FMT     _IOWR('V',  5, struct v4l2_format)
	CONST_VIDIOC_S_FMT = ((CONST_IOC_READ | CONST_IOC_WRITE)    << CONST_IOC_DIRSHIFT)  |
	                        (uintptr('V')                       << CONST_IOC_TYPESHIFT) |
	                        (5                                  << CONST_IOC_NRSHIFT)   |
	                        (unsafe.Sizeof(v4l2_format.Type{})  << CONST_IOC_SIZESHIFT)

	// A Golang conversion of the following C code:
	//
	// #define VIDIOC_ENUM_FRAMESIZES  _IOWR('V', 74, struct v4l2_frmsizeenum)
	CONST_VIDIOC_ENUM_FRAMESIZES = ((CONST_IOC_READ | CONST_IOC_WRITE)    << CONST_IOC_DIRSHIFT)  |
	                               (uintptr('V')                          << CONST_IOC_TYPESHIFT) |
	                               (74                                    << CONST_IOC_NRSHIFT)   |
	                               (unsafe.Sizeof(v4l2_framesize.Type{})  << CONST_IOC_SIZESHIFT)

)

