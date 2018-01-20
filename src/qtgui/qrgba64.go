//  header block begin
// /usr/include/qt/QtGui/qrgba64.h
// #include <qrgba64.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.Cthis
}
func NewQRgba64FromPointer(cthis unsafe.Pointer) *QRgba64 {
	return &QRgba64{&qtrt.CObject{cthis}}
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
func (this *QRgba64) FromRgba64(c uint64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Ey", ffiqt.FFI_TYPE_POINTER, c)
	gopp.ErrPrint(err, rv)
	return rv
}
func QRgba64_FromRgba64(c uint64) {
	var nilthis *QRgba64
	nilthis.FromRgba64(c)
}

// /usr/include/qt/QtGui/qrgba64.h:77
// index:1
// Public static inline
// QRgba64 fromRgba64(quint16, quint16, quint16, quint16)
func (this *QRgba64) FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Etttt", ffiqt.FFI_TYPE_POINTER, red, green, blue, alpha)
	gopp.ErrPrint(err, rv)
	return rv
}
func QRgba64_FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) {
	var nilthis *QRgba64
	nilthis.FromRgba64_1(red, green, blue, alpha)
}

// /usr/include/qt/QtGui/qrgba64.h:84
// index:0
// Public static inline
// QRgba64 fromRgba(quint8, quint8, quint8, quint8)
func (this *QRgba64) FromRgba(red byte, green byte, blue byte, alpha byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648fromRgbaEhhhh", ffiqt.FFI_TYPE_POINTER, red, green, blue, alpha)
	gopp.ErrPrint(err, rv)
	return rv
}
func QRgba64_FromRgba(red byte, green byte, blue byte, alpha byte) {
	var nilthis *QRgba64
	nilthis.FromRgba(red, green, blue, alpha)
}

// /usr/include/qt/QtGui/qrgba64.h:92
// index:0
// Public static inline
// QRgba64 fromArgb32(uint)
func (this *QRgba64) FromArgb32(rgb uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromArgb32Ej", ffiqt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	return rv
}
func QRgba64_FromArgb32(rgb uint) {
	var nilthis *QRgba64
	nilthis.FromArgb32(rgb)
}

// /usr/include/qt/QtGui/qrgba64.h:97
// index:0
// Public inline
// bool isOpaque()
func (this *QRgba64) IsOpaque() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba648isOpaqueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:101
// index:0
// Public inline
// bool isTransparent()
func (this *QRgba64) IsTransparent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6413isTransparentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:106
// index:0
// Public inline
// quint16 red()
func (this *QRgba64) Red() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba643redEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:107
// index:0
// Public inline
// quint16 green()
func (this *QRgba64) Green() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645greenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:108
// index:0
// Public inline
// quint16 blue()
func (this *QRgba64) Blue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba644blueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:109
// index:0
// Public inline
// quint16 alpha()
func (this *QRgba64) Alpha() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645alphaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:110
// index:0
// Public inline
// void setRed(quint16)
func (this *QRgba64) SetRed(_red uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba646setRedEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &_red)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:111
// index:0
// Public inline
// void setGreen(quint16)
func (this *QRgba64) SetGreen(_green uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648setGreenEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &_green)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:112
// index:0
// Public inline
// void setBlue(quint16)
func (this *QRgba64) SetBlue(_blue uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba647setBlueEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &_blue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:113
// index:0
// Public inline
// void setAlpha(quint16)
func (this *QRgba64) SetAlpha(_alpha uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648setAlphaEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &_alpha)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:115
// index:0
// Public inline
// quint8 red8()
func (this *QRgba64) Red8() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba644red8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:116
// index:0
// Public inline
// quint8 green8()
func (this *QRgba64) Green8() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba646green8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:117
// index:0
// Public inline
// quint8 blue8()
func (this *QRgba64) Blue8() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645blue8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:118
// index:0
// Public inline
// quint8 alpha8()
func (this *QRgba64) Alpha8() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba646alpha8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:119
// index:0
// Public inline
// uint toArgb32()
func (this *QRgba64) ToArgb32() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba648toArgb32Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:123
// index:0
// Public inline
// ushort toRgb16()
func (this *QRgba64) ToRgb16() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba647toRgb16Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:128
// index:0
// Public inline
// QRgba64 premultiplied()
func (this *QRgba64) Premultiplied() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6413premultipliedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:137
// index:0
// Public inline
// QRgba64 unpremultiplied()
func (this *QRgba64) Unpremultiplied() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6415unpremultipliedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
