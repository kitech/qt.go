package qtgui

// /usr/include/qt/QtGui/qglyphrun.h
// #include <qglyphrun.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QGlyphRun struct {
	*qtrt.CObject
}
type QGlyphRun_ITF interface {
	QGlyphRun_PTR() *QGlyphRun
}

func (ptr *QGlyphRun) QGlyphRun_PTR() *QGlyphRun { return ptr }

func (this *QGlyphRun) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGlyphRun) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGlyphRunFromPointer(cthis unsafe.Pointer) *QGlyphRun {
	return &QGlyphRun{&qtrt.CObject{cthis}}
}
func (*QGlyphRun) NewFromPointer(cthis unsafe.Pointer) *QGlyphRun {
	return NewQGlyphRunFromPointer(cthis)
}

// /usr/include/qt/QtGui/qglyphrun.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGlyphRun()
func NewQGlyphRun() *QGlyphRun {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRunC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGlyphRunFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGlyphRun)
	return gothis
}

// /usr/include/qt/QtGui/qglyphrun.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QGlyphRun()
func DeleteQGlyphRun(this *QGlyphRun) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRunD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qglyphrun.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QGlyphRun &)
func (this *QGlyphRun) Swap(other *QGlyphRun) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QRawFont rawFont()
func (this *QGlyphRun) RawFont() *QRawFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun7rawFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRawFont)
	return rv2
}

// /usr/include/qt/QtGui/qglyphrun.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawFont(const QRawFont &)
func (this *QGlyphRun) SetRawFont(rawFont *QRawFont) {
	var convArg0 = rawFont.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun10setRawFontERK8QRawFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawData(const quint32 *, const QPointF *, int)
func (this *QGlyphRun) SetRawData(glyphIndexArray unsafe.Pointer /*666*/, glyphPositionArray *qtcore.QPointF /*777 const QPointF **/, size int) {
	var convArg1 = glyphPositionArray.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun10setRawDataEPKjPK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &glyphIndexArray, convArg1, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QGlyphRun) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverline(_Bool)
func (this *QGlyphRun) SetOverline(overline bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun11setOverlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overline)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overline()
func (this *QGlyphRun) Overline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun8overlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnderline(_Bool)
func (this *QGlyphRun) SetUnderline(underline bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun12setUnderlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), underline)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underline()
func (this *QGlyphRun) Underline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun9underlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStrikeOut(_Bool)
func (this *QGlyphRun) SetStrikeOut(strikeOut bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun12setStrikeOutEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), strikeOut)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strikeOut()
func (this *QGlyphRun) StrikeOut() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun9strikeOutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRightToLeft(_Bool)
func (this *QGlyphRun) SetRightToLeft(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun14setRightToLeftEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRightToLeft()
func (this *QGlyphRun) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qglyphrun.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(enum QGlyphRun::GlyphRunFlag, _Bool)
func (this *QGlyphRun) SetFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun7setFlagENS_12GlyphRunFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QGlyphRun::GlyphRunFlags)
func (this *QGlyphRun) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun8setFlagsE6QFlagsINS_12GlyphRunFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] QGlyphRun::GlyphRunFlags flags()
func (this *QGlyphRun) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundingRect(const QRectF &)
func (this *QGlyphRun) SetBoundingRect(boundingRect *qtcore.QRectF) {
	var convArg0 = boundingRect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGlyphRun15setBoundingRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:113
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGlyphRun) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qglyphrun.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QGlyphRun) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGlyphRun7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QGlyphRun__GlyphRunFlag = int

const QGlyphRun__Overline QGlyphRun__GlyphRunFlag = 1
const QGlyphRun__Underline QGlyphRun__GlyphRunFlag = 2
const QGlyphRun__StrikeOut QGlyphRun__GlyphRunFlag = 4
const QGlyphRun__RightToLeft QGlyphRun__GlyphRunFlag = 8
const QGlyphRun__SplitLigature QGlyphRun__GlyphRunFlag = 16

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
