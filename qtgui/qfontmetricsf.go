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
// extern C begin: 38
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

type QFontMetricsF struct {
	*qtrt.CObject
}
type QFontMetricsF_ITF interface {
	QFontMetricsF_PTR() *QFontMetricsF
}

func (ptr *QFontMetricsF) QFontMetricsF_PTR() *QFontMetricsF { return ptr }

func (this *QFontMetricsF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontMetricsF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontMetricsFFromPointer(cthis unsafe.Pointer) *QFontMetricsF {
	return &QFontMetricsF{&qtrt.CObject{cthis}}
}
func (*QFontMetricsF) NewFromPointer(cthis unsafe.Pointer) *QFontMetricsF {
	return NewQFontMetricsFFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontmetrics.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontMetricsF(const QFont &)
func NewQFontMetricsF(arg0 QFont_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetricsF)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:135
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFontMetricsF(const QFont &, QPaintDevice *)
func NewQFontMetricsF_1(arg0 QFont_ITF, pd QPaintDevice_ITF /*777 QPaintDevice **/) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pd != nil && pd.QPaintDevice_PTR() != nil {
		convArg1 = pd.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFontP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetricsF)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:136
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QFontMetricsF(const QFontMetrics &)
func NewQFontMetricsF_2(arg0 QFontMetrics_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontMetrics_PTR() != nil {
		convArg0 = arg0.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK12QFontMetrics", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetricsF)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFontMetricsF()
func DeleteQFontMetricsF(this *QFontMetricsF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qfontmetrics.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontMetricsF & operator=(const QFontMetricsF &)
func (this *QFontMetricsF) Operator_equal(arg0 QFontMetricsF_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontMetricsF_PTR() != nil {
		convArg0 = arg0.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetricsF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:141
// index:1
// Public Visibility=Default Availability=Available
// [8] QFontMetricsF & operator=(const QFontMetrics &)
func (this *QFontMetricsF) Operator_equal_1(arg0 QFontMetrics_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontMetrics_PTR() != nil {
		convArg0 = arg0.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFaSERK12QFontMetrics", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetricsF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:143
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QFontMetricsF & operator=(QFontMetricsF &&)
func (this *QFontMetricsF) Operator_equal_2(other unsafe.Pointer /*333*/) *QFontMetricsF {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetricsF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:147
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFontMetricsF &)
func (this *QFontMetricsF) Swap(other QFontMetricsF_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetricsF_PTR() != nil {
		convArg0 = other.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsF4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent() const
func (this *QFontMetricsF) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal capHeight() const
func (this *QFontMetricsF) CapHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF9capHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent() const
func (this *QFontMetricsF) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal height() const
func (this *QFontMetricsF) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leading() const
func (this *QFontMetricsF) Leading() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF7leadingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lineSpacing() const
func (this *QFontMetricsF) LineSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF11lineSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minLeftBearing() const
func (this *QFontMetricsF) MinLeftBearing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF14minLeftBearingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minRightBearing() const
func (this *QFontMetricsF) MinRightBearing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF15minRightBearingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maxWidth() const
func (this *QFontMetricsF) MaxWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF8maxWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal xHeight() const
func (this *QFontMetricsF) XHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF7xHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal averageCharWidth() const
func (this *QFontMetricsF) AverageCharWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF16averageCharWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFont(QChar) const
func (this *QFontMetricsF) InFont(arg0 qtcore.QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF6inFontE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFontUcs4(uint) const
func (this *QFontMetricsF) InFontUcs4(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF10inFontUcs4Ej", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leftBearing(QChar) const
func (this *QFontMetricsF) LeftBearing(arg0 qtcore.QChar_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF11leftBearingE5QChar", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rightBearing(QChar) const
func (this *QFontMetricsF) RightBearing(arg0 qtcore.QChar_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12rightBearingE5QChar", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width(const QString &) const
func (this *QFontMetricsF) Width(string string) float64 {
	var tmpArg0 = qtcore.NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthERK7QString", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:169
// index:1
// Public Visibility=Default Availability=Available
// [8] qreal width(QChar) const
func (this *QFontMetricsF) Width_1(arg0 qtcore.QChar_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthE5QChar", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:171
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QString &) const
func (this *QFontMetricsF) BoundingRect(string string) *qtcore.QRectF /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:172
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(QChar) const
func (this *QFontMetricsF) BoundingRect_1(arg0 qtcore.QChar_ITF /*123*/) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:173
// index:2
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &, int, int *) const
func (this *QFontMetricsF) BoundingRect_2(r qtcore.QRectF_ITF, flags int, string string, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(string)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:173
// index:2
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &, int, int *) const
func (this *QFontMetricsF) BoundingRect_2_(r qtcore.QRectF_ITF, flags int, string string) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(string)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, int=Int, =Invalid,
	tabstops := int(0)
	// arg: 4, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:173
// index:2
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &, int, int *) const
func (this *QFontMetricsF) BoundingRect_2_1(r qtcore.QRectF_ITF, flags int, string string, tabstops int) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(string)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 4, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:174
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(int, const QString &, int, int *) const
func (this *QFontMetricsF) Size(flags int, str string, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QSizeF /*123*/ {
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:174
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(int, const QString &, int, int *) const
func (this *QFontMetricsF) Size__(flags int, str string) *qtcore.QSizeF /*123*/ {
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, int=Int, =Invalid,
	tabstops := int(0)
	// arg: 3, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:174
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(int, const QString &, int, int *) const
func (this *QFontMetricsF) Size__1(flags int, str string, tabstops int) *qtcore.QSizeF /*123*/ {
	var tmpArg1 = qtcore.NewQString_5(str)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, int *=Pointer, =Invalid,
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:176
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF tightBoundingRect(const QString &) const
func (this *QFontMetricsF) TightBoundingRect(text string) *qtcore.QRectF /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF17tightBoundingRectERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elidedText(const QString &, Qt::TextElideMode, qreal, int) const
func (this *QFontMetricsF) ElidedText(text string, mode int, width float64, flags int) string {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF10elidedTextERK7QStringN2Qt13TextElideModeEdi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, width, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontmetrics.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elidedText(const QString &, Qt::TextElideMode, qreal, int) const
func (this *QFontMetricsF) ElidedText__(text string, mode int, width float64) string {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, int=Int, =Invalid,
	flags := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF10elidedTextERK7QStringN2Qt13TextElideModeEdi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, width, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontmetrics.h:180
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal underlinePos() const
func (this *QFontMetricsF) UnderlinePos() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12underlinePosEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal overlinePos() const
func (this *QFontMetricsF) OverlinePos() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF11overlinePosEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal strikeOutPos() const
func (this *QFontMetricsF) StrikeOutPos() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12strikeOutPosEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lineWidth() const
func (this *QFontMetricsF) LineWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF9lineWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:185
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QFontMetricsF &) const
func (this *QFontMetricsF) Operator_equal_equal(other QFontMetricsF_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetricsF_PTR() != nil {
		convArg0 = other.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsFeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QFontMetricsF &) const
func (this *QFontMetricsF) Operator_not_equal(other QFontMetricsF_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetricsF_PTR() != nil {
		convArg0 = other.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsFneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
