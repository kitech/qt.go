package qtgui

// /usr/include/qt/QtGui/qrgba64.h
// #include <qrgba64.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QRgba64 struct {
	*qtrt.CObject
}

func (this *QRgba64) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRgba64) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQRgba64FromPointer(cthis unsafe.Pointer) *QRgba64 {
	return &QRgba64{&qtrt.CObject{cthis}}
}
func (*QRgba64) NewFromPointer(cthis unsafe.Pointer) *QRgba64 {
	return NewQRgba64FromPointer(cthis)
}

// /usr/include/qt/QtGui/qrgba64.h:69
// index:0
// Public inline
// void QRgba64()
func NewQRgba64() *QRgba64 {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba64C2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRgba64FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qrgba64.h:72
// index:0
// Public static inline
// QRgba64 fromRgba64(quint64)
func (this *QRgba64) FromRgba64(c uint64) *QRgba64 /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Ey", ffiqt.FFI_TYPE_POINTER, c)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QRgba64_FromRgba64(c uint64) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromRgba64(c)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:77
// index:1
// Public static inline
// QRgba64 fromRgba64(quint16, quint16, quint16, quint16)
func (this *QRgba64) FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) *QRgba64 /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Etttt", ffiqt.FFI_TYPE_POINTER, red, green, blue, alpha)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QRgba64_FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromRgba64_1(red, green, blue, alpha)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:84
// index:0
// Public static inline
// QRgba64 fromRgba(quint8, quint8, quint8, quint8)
func (this *QRgba64) FromRgba(red byte, green byte, blue byte, alpha byte) *QRgba64 /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648fromRgbaEhhhh", ffiqt.FFI_TYPE_POINTER, red, green, blue, alpha)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QRgba64_FromRgba(red byte, green byte, blue byte, alpha byte) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromRgba(red, green, blue, alpha)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:92
// index:0
// Public static inline
// QRgba64 fromArgb32(uint)
func (this *QRgba64) FromArgb32(rgb uint) *QRgba64 /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromArgb32Ej", ffiqt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QRgba64_FromArgb32(rgb uint) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromArgb32(rgb)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:97
// index:0
// Public inline
// bool isOpaque()
func (this *QRgba64) IsOpaque() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba648isOpaqueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrgba64.h:101
// index:0
// Public inline
// bool isTransparent()
func (this *QRgba64) IsTransparent() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6413isTransparentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrgba64.h:106
// index:0
// Public inline
// quint16 red()
func (this *QRgba64) Red() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba643redEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:107
// index:0
// Public inline
// quint16 green()
func (this *QRgba64) Green() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645greenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:108
// index:0
// Public inline
// quint16 blue()
func (this *QRgba64) Blue() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba644blueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:109
// index:0
// Public inline
// quint16 alpha()
func (this *QRgba64) Alpha() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645alphaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:110
// index:0
// Public inline
// void setRed(quint16)
func (this *QRgba64) SetRed(_red uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba646setRedEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), _red)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:111
// index:0
// Public inline
// void setGreen(quint16)
func (this *QRgba64) SetGreen(_green uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648setGreenEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), _green)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:112
// index:0
// Public inline
// void setBlue(quint16)
func (this *QRgba64) SetBlue(_blue uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba647setBlueEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), _blue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:113
// index:0
// Public inline
// void setAlpha(quint16)
func (this *QRgba64) SetAlpha(_alpha uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648setAlphaEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), _alpha)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:115
// index:0
// Public inline
// quint8 red8()
func (this *QRgba64) Red8() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba644red8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:116
// index:0
// Public inline
// quint8 green8()
func (this *QRgba64) Green8() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba646green8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:117
// index:0
// Public inline
// quint8 blue8()
func (this *QRgba64) Blue8() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645blue8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:118
// index:0
// Public inline
// quint8 alpha8()
func (this *QRgba64) Alpha8() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba646alpha8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:119
// index:0
// Public inline
// uint toArgb32()
func (this *QRgba64) ToArgb32() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba648toArgb32Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:123
// index:0
// Public inline
// ushort toRgb16()
func (this *QRgba64) ToRgb16() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba647toRgb16Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:128
// index:0
// Public inline
// QRgba64 premultiplied()
func (this *QRgba64) Premultiplied() *QRgba64 /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6413premultipliedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qrgba64.h:137
// index:0
// Public inline
// QRgba64 unpremultiplied()
func (this *QRgba64) Unpremultiplied() *QRgba64 /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6415unpremultipliedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QRgba64__Shifts = int

const QRgba64__RedShift QRgba64__Shifts = 0
const QRgba64__GreenShift QRgba64__Shifts = 16
const QRgba64__BlueShift QRgba64__Shifts = 32
const QRgba64__AlphaShift QRgba64__Shifts = 48

//  body block end
