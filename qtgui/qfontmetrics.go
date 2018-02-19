package qtgui

// /usr/include/qt/QtGui/qfontmetrics.h
// #include <qfontmetrics.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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

type QFontMetrics struct {
	*qtrt.CObject
}
type QFontMetrics_ITF interface {
	QFontMetrics_PTR() *QFontMetrics
}

func (ptr *QFontMetrics) QFontMetrics_PTR() *QFontMetrics { return ptr }

func (this *QFontMetrics) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontMetrics) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontMetricsFromPointer(cthis unsafe.Pointer) *QFontMetrics {
	return &QFontMetrics{&qtrt.CObject{cthis}}
}
func (*QFontMetrics) NewFromPointer(cthis unsafe.Pointer) *QFontMetrics {
	return NewQFontMetricsFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontmetrics.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontMetrics(const QFont &)
func NewQFontMetrics(arg0 QFont_ITF) *QFontMetrics {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetrics)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFontMetrics(const QFont &, QPaintDevice *)
func NewQFontMetrics_1(arg0 QFont_ITF, pd QPaintDevice_ITF /*777 QPaintDevice **/) *QFontMetrics {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pd != nil && pd.QPaintDevice_PTR() != nil {
		convArg1 = pd.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFontP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetrics)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFontMetrics()
func DeleteQFontMetrics(this *QFontMetrics) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QFontMetricsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qfontmetrics.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontMetrics & operator=(const QFontMetrics &)
func (this *QFontMetrics) Operator_equal(arg0 QFontMetrics_ITF) *QFontMetrics {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontMetrics_PTR() != nil {
		convArg0 = arg0.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QFontMetricsaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetrics)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:68
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QFontMetrics & operator=(QFontMetrics &&)
func (this *QFontMetrics) Operator_equal_1(other unsafe.Pointer /*333*/) *QFontMetrics {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QFontMetricsaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetrics)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFontMetrics &)
func (this *QFontMetrics) Swap(other QFontMetrics_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetrics_PTR() != nil {
		convArg0 = other.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QFontMetrics4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int ascent() const
func (this *QFontMetrics) Ascent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics6ascentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int capHeight() const
func (this *QFontMetrics) CapHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics9capHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int descent() const
func (this *QFontMetrics) Descent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics7descentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int height() const
func (this *QFontMetrics) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int leading() const
func (this *QFontMetrics) Leading() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics7leadingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineSpacing() const
func (this *QFontMetrics) LineSpacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics11lineSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int minLeftBearing() const
func (this *QFontMetrics) MinLeftBearing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics14minLeftBearingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int minRightBearing() const
func (this *QFontMetrics) MinRightBearing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics15minRightBearingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxWidth() const
func (this *QFontMetrics) MaxWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics8maxWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int xHeight() const
func (this *QFontMetrics) XHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics7xHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int averageCharWidth() const
func (this *QFontMetrics) AverageCharWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics16averageCharWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFont(QChar) const
func (this *QFontMetrics) InFont(arg0 qtcore.QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics6inFontE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFontUcs4(uint) const
func (this *QFontMetrics) InFontUcs4(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics10inFontUcs4Ej", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int leftBearing(QChar) const
func (this *QFontMetrics) LeftBearing(arg0 qtcore.QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics11leftBearingE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int rightBearing(QChar) const
func (this *QFontMetrics) RightBearing(arg0 qtcore.QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12rightBearingE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int width(const QString &, int) const
func (this *QFontMetrics) Width(arg0 string, len_ int) int {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int width(const QString &, int) const
func (this *QFontMetrics) Width__(arg0 string) int {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	len_ := -1
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:94
// index:1
// Public Visibility=Default Availability=Available
// [4] int width(const QString &, int, int) const
func (this *QFontMetrics) Width_1(arg0 string, len_ int, flags int) int {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:96
// index:2
// Public Visibility=Default Availability=Available
// [4] int width(QChar) const
func (this *QFontMetrics) Width_2(arg0 qtcore.QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics5widthE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int charWidth(const QString &, int) const
func (this *QFontMetrics) CharWidth(str string, pos int) int {
	var tmpArg0 = qtcore.NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics9charWidthERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pos)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:101
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(QChar) const
func (this *QFontMetrics) BoundingRect(arg0 qtcore.QChar_ITF /*123*/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:103
// index:1
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QString &) const
func (this *QFontMetrics) BoundingRect_1(text string) *qtcore.QRect /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:104
// index:2
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QRect &, int, const QString &, int, int *) const
func (this *QFontMetrics) BoundingRect_2(r qtcore.QRect_ITF, flags int, text string, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:104
// index:2
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QRect &, int, const QString &, int, int *) const
func (this *QFontMetrics) BoundingRect_2_(r qtcore.QRect_ITF, flags int, text string) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, int=Int, =Invalid,
	tabstops := 0
	// arg: 4, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:104
// index:2
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QRect &, int, const QString &, int, int *) const
func (this *QFontMetrics) BoundingRect_2_1(r qtcore.QRect_ITF, flags int, text string, tabstops int) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 4, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:105
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QRect boundingRect(int, int, int, int, int, const QString &, int, int *) const
func (this *QFontMetrics) BoundingRect_3(x int, y int, w int, h int, flags int, text string, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QRect /*123*/ {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:105
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QRect boundingRect(int, int, int, int, int, const QString &, int, int *) const
func (this *QFontMetrics) BoundingRect_3_(x int, y int, w int, h int, flags int, text string) *qtcore.QRect /*123*/ {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	// arg: 6, int=Int, =Invalid,
	tabstops := 0
	// arg: 7, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:105
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QRect boundingRect(int, int, int, int, int, const QString &, int, int *) const
func (this *QFontMetrics) BoundingRect_3_1(x int, y int, w int, h int, flags int, text string, tabstops int) *qtcore.QRect /*123*/ {
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	// arg: 7, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size(int, const QString &, int, int *) const
func (this *QFontMetrics) Size(flags int, str string, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QSize /*123*/ {
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size(int, const QString &, int, int *) const
func (this *QFontMetrics) Size__(flags int, str string) *qtcore.QSize /*123*/ {
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, int=Int, =Invalid,
	tabstops := 0
	// arg: 3, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size(int, const QString &, int, int *) const
func (this *QFontMetrics) Size__1(flags int, str string, tabstops int) *qtcore.QSize /*123*/ {
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:110
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect tightBoundingRect(const QString &) const
func (this *QFontMetrics) TightBoundingRect(text string) *qtcore.QRect /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics17tightBoundingRectERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elidedText(const QString &, Qt::TextElideMode, int, int) const
func (this *QFontMetrics) ElidedText(text string, mode int, width int, flags int) string {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics10elidedTextERK7QStringN2Qt13TextElideModeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, width, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontmetrics.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elidedText(const QString &, Qt::TextElideMode, int, int) const
func (this *QFontMetrics) ElidedText__(text string, mode int, width int) string {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, int=Int, =Invalid,
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics10elidedTextERK7QStringN2Qt13TextElideModeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, width, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontmetrics.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int underlinePos() const
func (this *QFontMetrics) UnderlinePos() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12underlinePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] int overlinePos() const
func (this *QFontMetrics) OverlinePos() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics11overlinePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] int strikeOutPos() const
func (this *QFontMetrics) StrikeOutPos() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics12strikeOutPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineWidth() const
func (this *QFontMetrics) LineWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetrics9lineWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QFontMetrics &) const
func (this *QFontMetrics) Operator_equal_equal(other QFontMetrics_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetrics_PTR() != nil {
		convArg0 = other.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetricseqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QFontMetrics &) const
func (this *QFontMetrics) Operator_not_equal(other QFontMetrics_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetrics_PTR() != nil {
		convArg0 = other.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QFontMetricsneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
