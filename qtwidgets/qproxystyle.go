package qtwidgets

// /usr/include/qt/QtWidgets/qproxystyle.h
// #include <qproxystyle.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QProxyStyle) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QProxyStyle struct {
	*QCommonStyle
}
type QProxyStyle_ITF interface {
	QCommonStyle_ITF
	QProxyStyle_PTR() *QProxyStyle
}

func (ptr *QProxyStyle) QProxyStyle_PTR() *QProxyStyle { return ptr }

func (this *QProxyStyle) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QCommonStyle.GetCthis()
	}
}
func (this *QProxyStyle) SetCthis(cthis unsafe.Pointer) {
	this.QCommonStyle = NewQCommonStyleFromPointer(cthis)
}
func NewQProxyStyleFromPointer(cthis unsafe.Pointer) *QProxyStyle {
	bcthis0 := NewQCommonStyleFromPointer(cthis)
	return &QProxyStyle{bcthis0}
}
func (*QProxyStyle) NewFromPointer(cthis unsafe.Pointer) *QProxyStyle {
	return NewQProxyStyleFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QProxyStyle) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qproxystyle.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProxyStyle(QStyle *)
func NewQProxyStyle(style *QStyle /*777 QStyle **/) *QProxyStyle {
	var convArg0 = style.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyleC2EP6QStyle", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProxyStyleFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qproxystyle.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QProxyStyle(const QString &)
func NewQProxyStyle_1(key string) *QProxyStyle {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyleC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProxyStyleFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qproxystyle.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QProxyStyle()
func DeleteQProxyStyle(this *QProxyStyle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * baseStyle()
func (this *QProxyStyle) BaseStyle() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle9baseStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qproxystyle.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseStyle(QStyle *)
func (this *QProxyStyle) SetBaseStyle(style *QStyle /*777 QStyle **/) {
	var convArg0 = style.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle12setBaseStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPrimitive(enum QStyle::PrimitiveElement, const QStyleOption *, QPainter *, const QWidget *)
func (this *QProxyStyle) DrawPrimitive(element int, option *QStyleOption /*777 const QStyleOption **/, painter *qtgui.QPainter /*777 QPainter **/, widget *QWidget /*777 const QWidget **/) {
	var convArg1 = option.GetCthis()
	var convArg2 = painter.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawControl(enum QStyle::ControlElement, const QStyleOption *, QPainter *, const QWidget *)
func (this *QProxyStyle) DrawControl(element int, option *QStyleOption /*777 const QStyleOption **/, painter *qtgui.QPainter /*777 QPainter **/, widget *QWidget /*777 const QWidget **/) {
	var convArg1 = option.GetCthis()
	var convArg2 = painter.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, QPainter *, const QWidget *)
func (this *QProxyStyle) DrawComplexControl(control int, option *QStyleOptionComplex /*777 const QStyleOptionComplex **/, painter *qtgui.QPainter /*777 QPainter **/, widget *QWidget /*777 const QWidget **/) {
	var convArg1 = option.GetCthis()
	var convArg2 = painter.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawItemText(QPainter *, const QRect &, int, const QPalette &, _Bool, const QString &, QPalette::ColorRole)
func (this *QProxyStyle) DrawItemText(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, flags int, pal *qtgui.QPalette, enabled bool, text string, textRole int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg3 = pal.GetCthis()
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags, convArg3, enabled, convArg5, textRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawItemPixmap(QPainter *, const QRect &, int, const QPixmap &)
func (this *QProxyStyle) DrawItemPixmap(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, alignment int, pixmap *qtgui.QPixmap) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg3 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeFromContents(enum QStyle::ContentsType, const QStyleOption *, const QSize &, const QWidget *)
func (this *QProxyStyle) SizeFromContents(type_ int, option *QStyleOption /*777 const QStyleOption **/, size *qtcore.QSize, widget *QWidget /*777 const QWidget **/) *qtcore.QSize /*123*/ {
	var convArg1 = option.GetCthis()
	var convArg2 = size.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subElementRect(enum QStyle::SubElement, const QStyleOption *, const QWidget *)
func (this *QProxyStyle) SubElementRect(element int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subControlRect(enum QStyle::ComplexControl, const QStyleOptionComplex *, enum QStyle::SubControl, const QWidget *)
func (this *QProxyStyle) SubControlRect(cc int, opt *QStyleOptionComplex /*777 const QStyleOptionComplex **/, sc int, widget *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, sc, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect itemTextRect(const QFontMetrics &, const QRect &, int, _Bool, const QString &)
func (this *QProxyStyle) ItemTextRect(fm *qtgui.QFontMetrics, r *qtcore.QRect, flags int, enabled bool, text string) *qtcore.QRect /*123*/ {
	var convArg0 = fm.GetCthis()
	var convArg1 = r.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(text)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags, enabled, convArg4)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect itemPixmapRect(const QRect &, int, const QPixmap &)
func (this *QProxyStyle) ItemPixmapRect(r *qtcore.QRect, flags int, pixmap *qtgui.QPixmap) *qtcore.QRect /*123*/ {
	var convArg0 = r.GetCthis()
	var convArg2 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, const QPoint &, const QWidget *)
func (this *QProxyStyle) HitTestComplexControl(control int, option *QStyleOptionComplex /*777 const QStyleOptionComplex **/, pos *qtcore.QPoint, widget *QWidget /*777 const QWidget **/) int {
	var convArg1 = option.GetCthis()
	var convArg2 = pos.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(enum QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *)
func (this *QProxyStyle) StyleHint(hint int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/, returnData *QStyleHintReturn /*777 QStyleHintReturn **/) int {
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	var convArg3 = returnData.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(enum QStyle::PixelMetric, const QStyleOption *, const QWidget *)
func (this *QProxyStyle) PixelMetric(metric int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) int {
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *)
func (this *QProxyStyle) LayoutSpacing(control1 int, control2 int, orientation int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) int {
	var convArg3 = option.GetCthis()
	var convArg4 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *)
func (this *QProxyStyle) StandardIcon(standardIcon int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtgui.QIcon /*123*/ {
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap standardPixmap(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *)
func (this *QProxyStyle) StandardPixmap(standardPixmap int, opt *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtgui.QPixmap /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardPixmap, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap generatedIconPixmap(QIcon::Mode, const QPixmap &, const QStyleOption *)
func (this *QProxyStyle) GeneratedIconPixmap(iconMode int, pixmap *qtgui.QPixmap, opt *QStyleOption /*777 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	var convArg1 = pixmap.GetCthis()
	var convArg2 = opt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), iconMode, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QPalette standardPalette()
func (this *QProxyStyle) StandardPalette() *qtgui.QPalette /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle15standardPaletteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QWidget *)
func (this *QProxyStyle) Polish(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:90
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QPalette &)
func (this *QProxyStyle) Polish_1(pal *qtgui.QPalette) {
	var convArg0 = pal.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle6polishER8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:91
// index:2
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QApplication *)
func (this *QProxyStyle) Polish_2(app *QApplication /*777 QApplication **/) {
	var convArg0 = app.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP12QApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QWidget *)
func (this *QProxyStyle) Unpolish(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:94
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QApplication *)
func (this *QProxyStyle) Unpolish_1(app *QApplication /*777 QApplication **/) {
	var convArg0 = app.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP12QApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QProxyStyle) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
