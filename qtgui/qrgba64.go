package qtgui

// /usr/include/qt/QtGui/qrgba64.h
// #include <qrgba64.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QRgba64 struct {
	*qtrt.CObject
}
type QRgba64_ITF interface {
	QRgba64_PTR() *QRgba64
}

func (ptr *QRgba64) QRgba64_PTR() *QRgba64 { return ptr }

func (this *QRgba64) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRgba64) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRgba64FromPointer(cthis unsafe.Pointer) *QRgba64 {
	return &QRgba64{&qtrt.CObject{cthis}}
}
func (*QRgba64) NewFromPointer(cthis unsafe.Pointer) *QRgba64 {
	return NewQRgba64FromPointer(cthis)
}

// /usr/include/qt/QtGui/qrgba64.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRgba64()

/*
Default constructs an instance of QRgba64.
*/
func NewQRgba64() *QRgba64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba64C2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRgba64)
	return gothis
}

// /usr/include/qt/QtGui/qrgba64.h:72
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRgba64 fromRgba64(quint64)

/*
Returns c as a QRgba64 struct.

See also fromArgb32().
*/
func (this *QRgba64) FromRgba64(c uint64) *QRgba64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Ey", qtrt.FFI_TYPE_POINTER, c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}
func QRgba64_FromRgba64(c uint64) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromRgba64(c)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:77
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QRgba64 fromRgba64(quint16, quint16, quint16, quint16)

/*
Returns c as a QRgba64 struct.

See also fromArgb32().
*/
func (this *QRgba64) FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) *QRgba64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba6410fromRgba64Etttt", qtrt.FFI_TYPE_POINTER, red, green, blue, alpha)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}
func QRgba64_FromRgba64_1(red uint16, green uint16, blue uint16, alpha uint16) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromRgba64_1(red, green, blue, alpha)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:84
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRgba64 fromRgba(quint8, quint8, quint8, quint8)

/*
Constructs a QRgba64 value from the four 8-bit color channels red, green, blue and alpha.

See also fromArgb32().
*/
func (this *QRgba64) FromRgba(red byte, green byte, blue byte, alpha byte) *QRgba64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba648fromRgbaEhhhh", qtrt.FFI_TYPE_POINTER, red, green, blue, alpha)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}
func QRgba64_FromRgba(red byte, green byte, blue byte, alpha byte) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromRgba(red, green, blue, alpha)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:92
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRgba64 fromArgb32(uint)

/*
Constructs a QRgba64 value from the 32bit ARGB value rgb.

See also fromRgba().
*/
func (this *QRgba64) FromArgb32(rgb uint) *QRgba64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba6410fromArgb32Ej", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}
func QRgba64_FromArgb32(rgb uint) *QRgba64 /*123*/ {
	var nilthis *QRgba64
	rv := nilthis.FromArgb32(rgb)
	return rv
}

// /usr/include/qt/QtGui/qrgba64.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isOpaque() const

/*
Returns whether the color is fully opaque.

See also isTransparent() and alpha().
*/
func (this *QRgba64) IsOpaque() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba648isOpaqueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrgba64.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTransparent() const

/*
Returns whether the color is transparent.

See also isOpaque() and alpha().
*/
func (this *QRgba64) IsTransparent() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba6413isTransparentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrgba64.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [2] quint16 red() const

/*
Returns the 16-bit red color component.

See also setRed().
*/
func (this *QRgba64) Red() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba643redEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [2] quint16 green() const

/*
Returns the 16-bit green color component.

See also setGreen().
*/
func (this *QRgba64) Green() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba645greenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [2] quint16 blue() const

/*
Returns the 16-bit blue color component.

See also setBlue().
*/
func (this *QRgba64) Blue() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba644blueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [2] quint16 alpha() const

/*
Returns the 16-bit alpha channel.

See also setAlpha().
*/
func (this *QRgba64) Alpha() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba645alphaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRed(quint16)

/*
Sets the red color component of this color to red.

See also red().
*/
func (this *QRgba64) SetRed(_red uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba646setRedEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), _red)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setGreen(quint16)

/*
Sets the green color component of this color to green.

See also green().
*/
func (this *QRgba64) SetGreen(_green uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba648setGreenEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), _green)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBlue(quint16)

/*
Sets the blue color component of this color to blue.

See also blue().
*/
func (this *QRgba64) SetBlue(_blue uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba647setBlueEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), _blue)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAlpha(quint16)

/*
Sets the alpha of this color to alpha.

See also alpha().
*/
func (this *QRgba64) SetAlpha(_alpha uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba648setAlphaEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), _alpha)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgba64.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] quint8 red8() const

/*
Returns the red color component as an 8-bit.
*/
func (this *QRgba64) Red8() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba644red8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [1] quint8 green8() const

/*
Returns the green color component as an 8-bit.
*/
func (this *QRgba64) Green8() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba646green8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [1] quint8 blue8() const

/*
Returns the blue color component as an 8-bit.
*/
func (this *QRgba64) Blue8() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba645blue8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [1] quint8 alpha8() const

/*
Returns the alpha channel as an 8-bit.
*/
func (this *QRgba64) Alpha8() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba646alpha8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [4] uint toArgb32() const

/*
Returns the color as a 32-bit ARGB value.

See also fromArgb32().
*/
func (this *QRgba64) ToArgb32() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba648toArgb32Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [2] ushort toRgb16() const

/*
Returns the color as a 16-bit RGB value.

See also toArgb32().
*/
func (this *QRgba64) ToRgb16() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba647toRgb16Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtGui/qrgba64.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRgba64 premultiplied() const

/*
Returns the color with the alpha premultiplied.

See also unpremultiplied().
*/
func (this *QRgba64) Premultiplied() *QRgba64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba6413premultipliedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}

// /usr/include/qt/QtGui/qrgba64.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRgba64 unpremultiplied() const

/*
Returns the color with the alpha unpremultiplied.

See also premultiplied().
*/
func (this *QRgba64) Unpremultiplied() *QRgba64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRgba6415unpremultipliedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}

// /usr/include/qt/QtGui/qrgba64.h:151
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRgba64 operator=(quint64)

/*

 */
func (this *QRgba64) Operator_equal(_rgba uint64) *QRgba64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba64aSEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), _rgba)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}

func DeleteQRgba64(this *QRgba64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRgba64D2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QRgba64__Shifts = int

//
const QRgba64__RedShift QRgba64__Shifts = 0

//
const QRgba64__GreenShift QRgba64__Shifts = 16

//
const QRgba64__BlueShift QRgba64__Shifts = 32

//
const QRgba64__AlphaShift QRgba64__Shifts = 48

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
