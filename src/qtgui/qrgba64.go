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

// /usr/include/qt/QtGui/qrgba64.h:69
// index:0
// inline
// void QRgba64()
func NewQRgba64() *QRgba64 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba64C2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRgba64FromPointer(cthis)
	return gothis
}
func NewQRgba64FromPointer(cthis unsafe.Pointer) *QRgba64 {
	return &QRgba64{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qrgba64.h:72
// index:0
// static inline
// QRgba64 fromRgba64(quint64)
func (this *QRgba64) FromRgba64(c uint64) {
	// 0: (c quint64), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Ey", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRgba64_FromRgba64(c uint64) {
	// 0: (c quint64), (c)
	var nilthis *QRgba64
	nilthis.FromRgba64(c)
}

// /usr/include/qt/QtGui/qrgba64.h:77
// index:1
// static inline
// QRgba64 fromRgba64(quint16, quint16, quint16, quint16)
func (this *QRgba64) FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) {
	// 1: (red quint16, green quint16, blue quint16, alpha quint16), (red, green, blue, alpha)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Etttt", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRgba64_FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) {
	// 1: (red quint16, green quint16, blue quint16, alpha quint16), (red, green, blue, alpha)
	var nilthis *QRgba64
	nilthis.FromRgba64_1(red, green, blue, alpha)
}

// /usr/include/qt/QtGui/qrgba64.h:84
// index:0
// static inline
// QRgba64 fromRgba(quint8, quint8, quint8, quint8)
func (this *QRgba64) FromRgba(red byte, green byte, blue byte, alpha byte) {
	// 0: (red quint8, green quint8, blue quint8, alpha quint8), (red, green, blue, alpha)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648fromRgbaEhhhh", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRgba64_FromRgba(red byte, green byte, blue byte, alpha byte) {
	// 0: (red quint8, green quint8, blue quint8, alpha quint8), (red, green, blue, alpha)
	var nilthis *QRgba64
	nilthis.FromRgba(red, green, blue, alpha)
}

// /usr/include/qt/QtGui/qrgba64.h:92
// index:0
// static inline
// QRgba64 fromArgb32(uint)
func (this *QRgba64) FromArgb32(rgb uint) {
	// 0: (rgb uint), (rgb)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba6410fromArgb32Ej", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRgba64_FromArgb32(rgb uint) {
	// 0: (rgb uint), (rgb)
	var nilthis *QRgba64
	nilthis.FromArgb32(rgb)
}

// /usr/include/qt/QtGui/qrgba64.h:97
// index:0
// inline
// bool isOpaque()
func (this *QRgba64) IsOpaque() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba648isOpaqueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:101
// index:0
// inline
// bool isTransparent()
func (this *QRgba64) IsTransparent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6413isTransparentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:106
// index:0
// inline
// quint16 red()
func (this *QRgba64) Red() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba643redEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:107
// index:0
// inline
// quint16 green()
func (this *QRgba64) Green() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645greenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:108
// index:0
// inline
// quint16 blue()
func (this *QRgba64) Blue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba644blueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:109
// index:0
// inline
// quint16 alpha()
func (this *QRgba64) Alpha() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645alphaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:110
// index:0
// inline
// void setRed(quint16)
func (this *QRgba64) SetRed(_red uint16) {
	// 0: (, _red quint16), (&_red)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba646setRedEt", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &_red)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:111
// index:0
// inline
// void setGreen(quint16)
func (this *QRgba64) SetGreen(_green uint16) {
	// 0: (, _green quint16), (&_green)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648setGreenEt", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &_green)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:112
// index:0
// inline
// void setBlue(quint16)
func (this *QRgba64) SetBlue(_blue uint16) {
	// 0: (, _blue quint16), (&_blue)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba647setBlueEt", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &_blue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:113
// index:0
// inline
// void setAlpha(quint16)
func (this *QRgba64) SetAlpha(_alpha uint16) {
	// 0: (, _alpha quint16), (&_alpha)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRgba648setAlphaEt", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &_alpha)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:115
// index:0
// inline
// quint8 red8()
func (this *QRgba64) Red8() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba644red8Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:116
// index:0
// inline
// quint8 green8()
func (this *QRgba64) Green8() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba646green8Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:117
// index:0
// inline
// quint8 blue8()
func (this *QRgba64) Blue8() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba645blue8Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:118
// index:0
// inline
// quint8 alpha8()
func (this *QRgba64) Alpha8() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba646alpha8Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:119
// index:0
// inline
// uint toArgb32()
func (this *QRgba64) ToArgb32() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba648toArgb32Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:123
// index:0
// inline
// ushort toRgb16()
func (this *QRgba64) ToRgb16() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba647toRgb16Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:128
// index:0
// inline
// QRgba64 premultiplied()
func (this *QRgba64) Premultiplied() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6413premultipliedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:137
// index:0
// inline
// QRgba64 unpremultiplied()
func (this *QRgba64) Unpremultiplied() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRgba6415unpremultipliedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
