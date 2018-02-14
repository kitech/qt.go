package qtgui

// /usr/include/qt/QtGui/qfontinfo.h
// #include <qfontinfo.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QFontInfo struct {
	*qtrt.CObject
}
type QFontInfo_ITF interface {
	QFontInfo_PTR() *QFontInfo
}

func (ptr *QFontInfo) QFontInfo_PTR() *QFontInfo { return ptr }

func (this *QFontInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontInfoFromPointer(cthis unsafe.Pointer) *QFontInfo {
	return &QFontInfo{&qtrt.CObject{cthis}}
}
func (*QFontInfo) NewFromPointer(cthis unsafe.Pointer) *QFontInfo {
	return NewQFontInfoFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontinfo.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontInfo(const QFont &)
func NewQFontInfo(arg0 QFont_ITF) *QFontInfo {
	var convArg0 = arg0.QFont_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFontInfoC2ERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontInfo)
	return gothis
}

// /usr/include/qt/QtGui/qfontinfo.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFontInfo()
func DeleteQFontInfo(this *QFontInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFontInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qfontinfo.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFontInfo &)
func (this *QFontInfo) Swap(other QFontInfo_ITF) {
	var convArg0 = other.QFontInfo_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFontInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QString family()
func (this *QFontInfo) Family() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo6familyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontinfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleName()
func (this *QFontInfo) StyleName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9styleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontinfo.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int pixelSize()
func (this *QFontInfo) PixelSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9pixelSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int pointSize()
func (this *QFontInfo) PointSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9pointSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal pointSizeF()
func (this *QFontInfo) PointSizeF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo10pointSizeFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool italic()
func (this *QFontInfo) Italic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo6italicEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Style style()
func (this *QFontInfo) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight()
func (this *QFontInfo) Weight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo6weightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontinfo.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool bold()
func (this *QFontInfo) Bold() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo4boldEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underline()
func (this *QFontInfo) Underline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9underlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overline()
func (this *QFontInfo) Overline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo8overlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strikeOut()
func (this *QFontInfo) StrikeOut() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9strikeOutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fixedPitch()
func (this *QFontInfo) FixedPitch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo10fixedPitchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::StyleHint styleHint()
func (this *QFontInfo) StyleHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo9styleHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfontinfo.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rawMode()
func (this *QFontInfo) RawMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo7rawModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontinfo.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exactMatch()
func (this *QFontInfo) ExactMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFontInfo10exactMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

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
