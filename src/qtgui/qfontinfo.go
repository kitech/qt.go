//  header block begin
// /usr/include/qt/QtGui/qfontinfo.h
// #include <qfontinfo.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QFontInfo struct {
	*qtrt.CObject
}

func (this *QFontInfo) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQFontInfoFromPointer(cthis unsafe.Pointer) *QFontInfo {
	return &QFontInfo{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qfontinfo.h:53
// index:0
// Public
// void QFontInfo(const class QFont &)
func NewQFontInfo(arg0 *QFont) *QFontInfo {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFontInfoC2ERK5QFont", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontinfo.h:55
// index:0
// Public
// void ~QFontInfo()
func DeleteQFontInfo(*QFontInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFontInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:59
// index:0
// Public inline
// void swap(class QFontInfo &)
func (this *QFontInfo) Swap(other *QFontInfo) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFontInfo4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:61
// index:0
// Public
// QString family()
func (this *QFontInfo) Family() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo6familyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:62
// index:0
// Public
// QString styleName()
func (this *QFontInfo) StyleName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9styleNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:63
// index:0
// Public
// int pixelSize()
func (this *QFontInfo) PixelSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9pixelSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:64
// index:0
// Public
// int pointSize()
func (this *QFontInfo) PointSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9pointSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:65
// index:0
// Public
// qreal pointSizeF()
func (this *QFontInfo) PointSizeF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo10pointSizeFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:66
// index:0
// Public
// bool italic()
func (this *QFontInfo) Italic() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo6italicEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:67
// index:0
// Public
// QFont::Style style()
func (this *QFontInfo) Style() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:68
// index:0
// Public
// int weight()
func (this *QFontInfo) Weight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo6weightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:69
// index:0
// Public inline
// bool bold()
func (this *QFontInfo) Bold() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo4boldEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:70
// index:0
// Public
// bool underline()
func (this *QFontInfo) Underline() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9underlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:71
// index:0
// Public
// bool overline()
func (this *QFontInfo) Overline() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo8overlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:72
// index:0
// Public
// bool strikeOut()
func (this *QFontInfo) StrikeOut() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9strikeOutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:73
// index:0
// Public
// bool fixedPitch()
func (this *QFontInfo) FixedPitch() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo10fixedPitchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:74
// index:0
// Public
// QFont::StyleHint styleHint()
func (this *QFontInfo) StyleHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9styleHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:76
// index:0
// Public
// bool rawMode()
func (this *QFontInfo) RawMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo7rawModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontinfo.h:79
// index:0
// Public
// bool exactMatch()
func (this *QFontInfo) ExactMatch() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo10exactMatchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
