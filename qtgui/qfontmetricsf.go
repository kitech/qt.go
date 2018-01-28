package qtgui

// /usr/include/qt/QtGui/qfontmetrics.h
// #include <qfontmetrics.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QFontMetricsF struct {
	*qtrt.CObject
}

func (this *QFontMetricsF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontMetricsF) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQFontMetricsFFromPointer(cthis unsafe.Pointer) *QFontMetricsF {
	return &QFontMetricsF{&qtrt.CObject{cthis}}
}
func (*QFontMetricsF) NewFromPointer(cthis unsafe.Pointer) *QFontMetricsF {
	return NewQFontMetricsFFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontmetrics.h:134
// index:0
// Public
// void QFontMetricsF(const QFont &)
func NewQFontMetricsF(arg0 *QFont) *QFontMetricsF {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFont", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:135
// index:1
// Public
// void QFontMetricsF(const QFont &, QPaintDevice *)
func NewQFontMetricsF_1(arg0 *QFont, pd *QPaintDevice /*777 QPaintDevice **/) *QFontMetricsF {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	var convArg1 = pd.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:136
// index:2
// Public
// void QFontMetricsF(const QFontMetrics &)
func NewQFontMetricsF_2(arg0 *QFontMetrics) *QFontMetricsF {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK12QFontMetrics", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:138
// index:0
// Public
// void ~QFontMetricsF()
func DeleteQFontMetricsF(*QFontMetricsF) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:147
// index:0
// Public inline
// void swap(QFontMetricsF &)
func (this *QFontMetricsF) Swap(other *QFontMetricsF) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsF4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:149
// index:0
// Public
// qreal ascent()
func (this *QFontMetricsF) Ascent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF6ascentEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:150
// index:0
// Public
// qreal capHeight()
func (this *QFontMetricsF) CapHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF9capHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:151
// index:0
// Public
// qreal descent()
func (this *QFontMetricsF) Descent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF7descentEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:152
// index:0
// Public
// qreal height()
func (this *QFontMetricsF) Height() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF6heightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:153
// index:0
// Public
// qreal leading()
func (this *QFontMetricsF) Leading() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF7leadingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:154
// index:0
// Public
// qreal lineSpacing()
func (this *QFontMetricsF) LineSpacing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF11lineSpacingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:155
// index:0
// Public
// qreal minLeftBearing()
func (this *QFontMetricsF) MinLeftBearing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF14minLeftBearingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:156
// index:0
// Public
// qreal minRightBearing()
func (this *QFontMetricsF) MinRightBearing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF15minRightBearingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:157
// index:0
// Public
// qreal maxWidth()
func (this *QFontMetricsF) MaxWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF8maxWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:159
// index:0
// Public
// qreal xHeight()
func (this *QFontMetricsF) XHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF7xHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:160
// index:0
// Public
// qreal averageCharWidth()
func (this *QFontMetricsF) AverageCharWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF16averageCharWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:162
// index:0
// Public
// bool inFont(QChar)
func (this *QFontMetricsF) InFont(arg0 *qtcore.QChar /*123*/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF6inFontE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:163
// index:0
// Public
// bool inFontUcs4(uint)
func (this *QFontMetricsF) InFontUcs4(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF10inFontUcs4Ej", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:165
// index:0
// Public
// qreal leftBearing(QChar)
func (this *QFontMetricsF) LeftBearing(arg0 *qtcore.QChar /*123*/) float64 {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF11leftBearingE5QChar", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:166
// index:0
// Public
// qreal rightBearing(QChar)
func (this *QFontMetricsF) RightBearing(arg0 *qtcore.QChar /*123*/) float64 {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12rightBearingE5QChar", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:167
// index:0
// Public
// qreal width(const QString &)
func (this *QFontMetricsF) Width(string *qtcore.QString) float64 {
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthERK7QString", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:169
// index:1
// Public
// qreal width(QChar)
func (this *QFontMetricsF) Width_1(arg0 *qtcore.QChar /*123*/) float64 {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthE5QChar", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:171
// index:0
// Public
// QRectF boundingRect(const QString &)
func (this *QFontMetricsF) BoundingRect(string *qtcore.QString) *qtcore.QRectF /*123*/ {
	var convArg0 = string.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:172
// index:1
// Public
// QRectF boundingRect(QChar)
func (this *QFontMetricsF) BoundingRect_1(arg0 *qtcore.QChar /*123*/) *qtcore.QRectF /*123*/ {
	var convArg0 = arg0.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectE5QChar", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:173
// index:2
// Public
// QRectF boundingRect(const QRectF &, int, const QString &, int, int *)
func (this *QFontMetricsF) BoundingRect_2(r *qtcore.QRectF, flags int, string *qtcore.QString, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QRectF /*123*/ {
	var convArg0 = r.GetCthis()
	var convArg2 = string.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, flags, convArg2, tabstops, &tabarray)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:174
// index:0
// Public
// QSizeF size(int, const QString &, int, int *)
func (this *QFontMetricsF) Size(flags int, str *qtcore.QString, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QSizeF /*123*/ {
	var convArg1 = str.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), flags, convArg1, tabstops, &tabarray)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:176
// index:0
// Public
// QRectF tightBoundingRect(const QString &)
func (this *QFontMetricsF) TightBoundingRect(text *qtcore.QString) *qtcore.QRectF /*123*/ {
	var convArg0 = text.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF17tightBoundingRectERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:178
// index:0
// Public
// QString elidedText(const QString &, Qt::TextElideMode, qreal, int)
func (this *QFontMetricsF) ElidedText(text *qtcore.QString, mode int, width float64, flags int) *qtcore.QString /*123*/ {
	var convArg0 = text.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF10elidedTextERK7QStringN2Qt13TextElideModeEdi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode, width, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:180
// index:0
// Public
// qreal underlinePos()
func (this *QFontMetricsF) UnderlinePos() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12underlinePosEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:181
// index:0
// Public
// qreal overlinePos()
func (this *QFontMetricsF) OverlinePos() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF11overlinePosEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:182
// index:0
// Public
// qreal strikeOutPos()
func (this *QFontMetricsF) StrikeOutPos() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12strikeOutPosEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:183
// index:0
// Public
// qreal lineWidth()
func (this *QFontMetricsF) LineWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF9lineWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

//  body block end
