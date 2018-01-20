//  header block begin
// /usr/include/qt/QtWidgets/qstyle.h
// #include <qstyle.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyle struct {
	*qtcore.QObject
}

func (this *QStyle) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qstyle.h:66
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStyle) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:70
// index:0
// void QStyle(class QStylePrivate &)
func NewQStyle(dd unsafe.Pointer) *QStyle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyleC1ER13QStylePrivate", ffiqt.FFI_TYPE_VOID, cthis, dd)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleFromPointer(cthis)
	return gothis
}
func NewQStyleFromPointer(cthis unsafe.Pointer) *QStyle {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStyle{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyle.h:73
// index:1
// void QStyle()
func NewQStyle_1() *QStyle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyleC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyle.h:74
// index:0
// virtual
// void ~QStyle()
func DeleteQStyle(*QStyle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:76
// index:0
// virtual
// void polish(class QWidget *)
func (this *QStyle) Polish(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:79
// index:1
// virtual
// void polish(class QApplication *)
func (this *QStyle) Polish_1(application unsafe.Pointer) {
	// 1: (, application QApplication *), (application)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.GetCthis(), application)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:82
// index:2
// virtual
// void polish(class QPalette &)
func (this *QStyle) Polish_2(palette unsafe.Pointer) {
	// 2: (, palette QPalette &), (palette)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishER8QPalette", ffiqt.FFI_TYPE_VOID, this.GetCthis(), palette)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:77
// index:0
// virtual
// void unpolish(class QWidget *)
func (this *QStyle) Unpolish(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:80
// index:1
// virtual
// void unpolish(class QApplication *)
func (this *QStyle) Unpolish_1(application unsafe.Pointer) {
	// 1: (, application QApplication *), (application)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.GetCthis(), application)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:84
// index:0
// virtual
// QRect itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
func (this *QStyle) ItemTextRect(fm unsafe.Pointer, r unsafe.Pointer, flags int, enabled bool, text unsafe.Pointer) {
	// 0: (, fm const QFontMetrics &, r const QRect &, flags int, enabled bool, text const QString &), (fm, r, &flags, &enabled, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fm, r, &flags, &enabled, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:88
// index:0
// virtual
// QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
func (this *QStyle) ItemPixmapRect(r unsafe.Pointer, flags int, pixmap unsafe.Pointer) {
	// 0: (, r const QRect &, flags int, pixmap const QPixmap &), (r, &flags, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &flags, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:90
// index:0
// virtual
// void drawItemText(class QPainter *, const class QRect &, int, const class QPalette &, _Bool, const class QString &, class QPalette::ColorRole)
func (this *QStyle) DrawItemText(painter unsafe.Pointer, rect unsafe.Pointer, flags int, pal unsafe.Pointer, enabled bool, text unsafe.Pointer, textRole int) {
	// 0: (, painter QPainter *, rect const QRect &, flags int, pal const QPalette &, enabled bool, text const QString &, textRole QPalette::ColorRole), (painter, rect, &flags, pal, &enabled, text, &textRole)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, &flags, pal, &enabled, text, &textRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:94
// index:0
// virtual
// void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
func (this *QStyle) DrawItemPixmap(painter unsafe.Pointer, rect unsafe.Pointer, alignment int, pixmap unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRect &, alignment int, pixmap const QPixmap &), (painter, rect, &alignment, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, &alignment, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:97
// index:0
// virtual
// QPalette standardPalette()
func (this *QStyle) StandardPalette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle15standardPaletteEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:204
// index:0
// pure virtual
// void drawPrimitive(enum QStyle::PrimitiveElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QStyle) DrawPrimitive(pe int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, pe QStyle::PrimitiveElement, opt const QStyleOption *, p QPainter *, w const QWidget *), (&pe, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle13drawPrimitiveENS_16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pe, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:275
// index:0
// pure virtual
// void drawControl(enum QStyle::ControlElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QStyle) DrawControl(element int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, element QStyle::ControlElement, opt const QStyleOption *, p QPainter *, w const QWidget *), (&element, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle11drawControlENS_14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &element, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:364
// index:0
// pure virtual
// QRect subElementRect(enum QStyle::SubElement, const class QStyleOption *, const class QWidget *)
func (this *QStyle) SubElementRect(subElement int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, subElement QStyle::SubElement, option const QStyleOption *, widget const QWidget *), (&subElement, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14subElementRectENS_10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &subElement, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:443
// index:0
// pure virtual
// void drawComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, class QPainter *, const class QWidget *)
func (this *QStyle) DrawComplexControl(cc int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, cc QStyle::ComplexControl, opt const QStyleOptionComplex *, p QPainter *, widget const QWidget *), (&cc, opt, p, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle18drawComplexControlENS_14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cc, opt, p, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:445
// index:0
// pure virtual
// QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, const class QPoint &, const class QWidget *)
func (this *QStyle) HitTestComplexControl(cc int, opt unsafe.Pointer, pt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, cc QStyle::ComplexControl, opt const QStyleOptionComplex *, pt const QPoint &, widget const QWidget *), (&cc, opt, pt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle21hitTestComplexControlENS_14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cc, opt, pt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:447
// index:0
// pure virtual
// QRect subControlRect(enum QStyle::ComplexControl, const class QStyleOptionComplex *, enum QStyle::SubControl, const class QWidget *)
func (this *QStyle) SubControlRect(cc int, opt unsafe.Pointer, sc int, widget unsafe.Pointer) {
	// 0: (, cc QStyle::ComplexControl, opt const QStyleOptionComplex *, sc QStyle::SubControl, widget const QWidget *), (&cc, opt, &sc, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14subControlRectENS_14ComplexControlEPK19QStyleOptionComplexNS_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cc, opt, &sc, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:582
// index:0
// pure virtual
// int pixelMetric(enum QStyle::PixelMetric, const class QStyleOption *, const class QWidget *)
func (this *QStyle) PixelMetric(metric int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, metric QStyle::PixelMetric, option const QStyleOption *, widget const QWidget *), (&metric, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle11pixelMetricENS_11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &metric, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:614
// index:0
// pure virtual
// QSize sizeFromContents(enum QStyle::ContentsType, const class QStyleOption *, const class QSize &, const class QWidget *)
func (this *QStyle) SizeFromContents(ct int, opt unsafe.Pointer, contentsSize unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, ct QStyle::ContentsType, opt const QStyleOption *, contentsSize const QSize &, w const QWidget *), (&ct, opt, contentsSize, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle16sizeFromContentsENS_12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ct, opt, contentsSize, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:748
// index:0
// pure virtual
// int styleHint(enum QStyle::StyleHint, const class QStyleOption *, const class QWidget *, class QStyleHintReturn *)
func (this *QStyle) StyleHint(stylehint int, opt unsafe.Pointer, widget unsafe.Pointer, returnData unsafe.Pointer) {
	// 0: (, stylehint QStyle::StyleHint, opt const QStyleOption *, widget const QWidget *, returnData QStyleHintReturn *), (&stylehint, opt, widget, returnData)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle9styleHintENS_9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &stylehint, opt, widget, returnData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:828
// index:0
// pure virtual
// QPixmap standardPixmap(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QStyle) StandardPixmap(standardPixmap int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, standardPixmap QStyle::StandardPixmap, opt const QStyleOption *, widget const QWidget *), (&standardPixmap, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14standardPixmapENS_14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &standardPixmap, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:831
// index:0
// pure virtual
// QIcon standardIcon(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QStyle) StandardIcon(standardIcon int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, standardIcon QStyle::StandardPixmap, option const QStyleOption *, widget const QWidget *), (&standardIcon, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12standardIconENS_14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &standardIcon, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:834
// index:0
// pure virtual
// QPixmap generatedIconPixmap(class QIcon::Mode, const class QPixmap &, const class QStyleOption *)
func (this *QStyle) GeneratedIconPixmap(iconMode int, pixmap unsafe.Pointer, opt unsafe.Pointer) {
	// 0: (, iconMode QIcon::Mode, pixmap const QPixmap &, opt const QStyleOption *), (&iconMode, pixmap, opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &iconMode, pixmap, opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:837
// index:0
// static
// QRect visualRect(Qt::LayoutDirection, const class QRect &, const class QRect &)
func (this *QStyle) VisualRect(direction int, boundingRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	// 0: (direction Qt::LayoutDirection, boundingRect const QRect &, logicalRect const QRect &), (direction, boundingRect, logicalRect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle10visualRectEN2Qt15LayoutDirectionERK5QRectS4_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyle_VisualRect(direction int, boundingRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	// 0: (direction Qt::LayoutDirection, boundingRect const QRect &, logicalRect const QRect &), (direction, boundingRect, logicalRect)
	var nilthis *QStyle
	nilthis.VisualRect(direction, boundingRect, logicalRect)
}

// /usr/include/qt/QtWidgets/qstyle.h:839
// index:0
// static
// QPoint visualPos(Qt::LayoutDirection, const class QRect &, const class QPoint &)
func (this *QStyle) VisualPos(direction int, boundingRect unsafe.Pointer, logicalPos unsafe.Pointer) {
	// 0: (direction Qt::LayoutDirection, boundingRect const QRect &, logicalPos const QPoint &), (direction, boundingRect, logicalPos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle9visualPosEN2Qt15LayoutDirectionERK5QRectRK6QPoint", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyle_VisualPos(direction int, boundingRect unsafe.Pointer, logicalPos unsafe.Pointer) {
	// 0: (direction Qt::LayoutDirection, boundingRect const QRect &, logicalPos const QPoint &), (direction, boundingRect, logicalPos)
	var nilthis *QStyle
	nilthis.VisualPos(direction, boundingRect, logicalPos)
}

// /usr/include/qt/QtWidgets/qstyle.h:841
// index:0
// static
// int sliderPositionFromValue(int, int, int, int, _Bool)
func (this *QStyle) SliderPositionFromValue(min int, max int, val int, space int, upsideDown bool) {
	// 0: (min int, max int, val int, space int, upsideDown bool), (min, max, val, space, upsideDown)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle23sliderPositionFromValueEiiiib", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyle_SliderPositionFromValue(min int, max int, val int, space int, upsideDown bool) {
	// 0: (min int, max int, val int, space int, upsideDown bool), (min, max, val, space, upsideDown)
	var nilthis *QStyle
	nilthis.SliderPositionFromValue(min, max, val, space, upsideDown)
}

// /usr/include/qt/QtWidgets/qstyle.h:843
// index:0
// static
// int sliderValueFromPosition(int, int, int, int, _Bool)
func (this *QStyle) SliderValueFromPosition(min int, max int, pos int, space int, upsideDown bool) {
	// 0: (min int, max int, pos int, space int, upsideDown bool), (min, max, pos, space, upsideDown)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle23sliderValueFromPositionEiiiib", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyle_SliderValueFromPosition(min int, max int, pos int, space int, upsideDown bool) {
	// 0: (min int, max int, pos int, space int, upsideDown bool), (min, max, pos, space, upsideDown)
	var nilthis *QStyle
	nilthis.SliderValueFromPosition(min, max, pos, space, upsideDown)
}

// /usr/include/qt/QtWidgets/qstyle.h:845
// index:0
// static
// Qt::Alignment visualAlignment(Qt::LayoutDirection, Qt::Alignment)
func (this *QStyle) VisualAlignment(direction int, alignment int) {
	// 0: (direction Qt::LayoutDirection, QFlags<Qt::AlignmentFlag> alignment), (direction, alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle15visualAlignmentEN2Qt15LayoutDirectionE6QFlagsINS0_13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyle_VisualAlignment(direction int, alignment int) {
	// 0: (direction Qt::LayoutDirection, QFlags<Qt::AlignmentFlag> alignment), (direction, alignment)
	var nilthis *QStyle
	nilthis.VisualAlignment(direction, alignment)
}

// /usr/include/qt/QtWidgets/qstyle.h:846
// index:0
// static
// QRect alignedRect(Qt::LayoutDirection, Qt::Alignment, const class QSize &, const class QRect &)
func (this *QStyle) AlignedRect(direction int, alignment int, size unsafe.Pointer, rectangle unsafe.Pointer) {
	// 0: (direction Qt::LayoutDirection, QFlags<Qt::AlignmentFlag> alignment, size const QSize &, rectangle const QRect &), (direction, alignment, size, rectangle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle11alignedRectEN2Qt15LayoutDirectionE6QFlagsINS0_13AlignmentFlagEERK5QSizeRK5QRect", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyle_AlignedRect(direction int, alignment int, size unsafe.Pointer, rectangle unsafe.Pointer) {
	// 0: (direction Qt::LayoutDirection, QFlags<Qt::AlignmentFlag> alignment, size const QSize &, rectangle const QRect &), (direction, alignment, size, rectangle)
	var nilthis *QStyle
	nilthis.AlignedRect(direction, alignment, size, rectangle)
}

// /usr/include/qt/QtWidgets/qstyle.h:849
// index:0
// pure virtual
// int layoutSpacing(class QSizePolicy::ControlType, class QSizePolicy::ControlType, Qt::Orientation, const class QStyleOption *, const class QWidget *)
func (this *QStyle) LayoutSpacing(control1 int, control2 int, orientation int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, control1 QSizePolicy::ControlType, control2 QSizePolicy::ControlType, orientation Qt::Orientation, option const QStyleOption *, widget const QWidget *), (&control1, &control2, &orientation, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &control1, &control2, &orientation, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:852
// index:0
// int combinedLayoutSpacing(class QSizePolicy::ControlTypes, class QSizePolicy::ControlTypes, Qt::Orientation, class QStyleOption *, class QWidget *)
func (this *QStyle) CombinedLayoutSpacing(controls1 int, controls2 int, orientation int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QFlags<QSizePolicy::ControlType> controls1, QFlags<QSizePolicy::ControlType> controls2, orientation Qt::Orientation, option QStyleOption *, widget QWidget *), (&controls1, &controls2, &orientation, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle21combinedLayoutSpacingE6QFlagsIN11QSizePolicy11ControlTypeEES3_N2Qt11OrientationEP12QStyleOptionP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &controls1, &controls2, &orientation, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:856
// index:0
// const QStyle * proxy()
func (this *QStyle) Proxy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle5proxyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
