//  header block begin
// /usr/include/qt/QtGui/qfontinfo.h
// #include <qfontinfo.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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

// /usr/include/qt/QtGui/qfontinfo.h:53
// index:0
// void QFontInfo(const class QFont &)
func NewQFontInfo(arg0 unsafe.Pointer) *QFontInfo {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFontInfoC2ERK5QFont", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontInfoFromPointer(cthis)
	return gothis
}
func NewQFontInfoFromPointer(cthis unsafe.Pointer) *QFontInfo {
	return &QFontInfo{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qfontinfo.h:55
// index:0
// void ~QFontInfo()
func DeleteQFontInfo(*QFontInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFontInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:59
// index:0
// inline
// void swap(class QFontInfo &)
func (this *QFontInfo) Swap(other unsafe.Pointer) {
	// 0: (, other QFontInfo &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QFontInfo4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:61
// index:0
// QString family()
func (this *QFontInfo) Family() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo6familyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:62
// index:0
// QString styleName()
func (this *QFontInfo) StyleName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9styleNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:63
// index:0
// int pixelSize()
func (this *QFontInfo) PixelSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9pixelSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:64
// index:0
// int pointSize()
func (this *QFontInfo) PointSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9pointSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:65
// index:0
// qreal pointSizeF()
func (this *QFontInfo) PointSizeF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo10pointSizeFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:66
// index:0
// bool italic()
func (this *QFontInfo) Italic() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo6italicEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:67
// index:0
// QFont::Style style()
func (this *QFontInfo) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo5styleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:68
// index:0
// int weight()
func (this *QFontInfo) Weight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo6weightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:69
// index:0
// inline
// bool bold()
func (this *QFontInfo) Bold() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo4boldEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:70
// index:0
// bool underline()
func (this *QFontInfo) Underline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9underlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:71
// index:0
// bool overline()
func (this *QFontInfo) Overline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo8overlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:72
// index:0
// bool strikeOut()
func (this *QFontInfo) StrikeOut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9strikeOutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:73
// index:0
// bool fixedPitch()
func (this *QFontInfo) FixedPitch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo10fixedPitchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:74
// index:0
// QFont::StyleHint styleHint()
func (this *QFontInfo) StyleHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo9styleHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:76
// index:0
// bool rawMode()
func (this *QFontInfo) RawMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo7rawModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:79
// index:0
// bool exactMatch()
func (this *QFontInfo) ExactMatch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QFontInfo10exactMatchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
