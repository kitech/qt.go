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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QProxyStyle) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QProxyStyle) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qproxystyle.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProxyStyle(QStyle *)

/*
Constructs a QProxyStyle object for overriding behavior in the specified style, or in the default native style if style is not specified.

Ownership of style is transferred to QProxyStyle.
*/
func NewQProxyStyle(style QStyle_ITF /*777 QStyle **/) *QProxyStyle {
	var convArg0 unsafe.Pointer
	if style != nil && style.QStyle_PTR() != nil {
		convArg0 = style.QStyle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyleC2EP6QStyle", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProxyStyleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProxyStyle")
	return gothis
}

// /usr/include/qt/QtWidgets/qproxystyle.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProxyStyle(QStyle *)

/*
Constructs a QProxyStyle object for overriding behavior in the specified style, or in the default native style if style is not specified.

Ownership of style is transferred to QProxyStyle.
*/
func NewQProxyStyle__() *QProxyStyle {
	// arg: 0, QStyle *=Pointer, QStyle=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyleC2EP6QStyle", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProxyStyleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProxyStyle")
	return gothis
}

// /usr/include/qt/QtWidgets/qproxystyle.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QProxyStyle(const QString &)

/*
Constructs a QProxyStyle object for overriding behavior in the specified style, or in the default native style if style is not specified.

Ownership of style is transferred to QProxyStyle.
*/
func NewQProxyStyle_1(key string) *QProxyStyle {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyleC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProxyStyleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProxyStyle")
	return gothis
}

// /usr/include/qt/QtWidgets/qproxystyle.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QProxyStyle()

/*

 */
func DeleteQProxyStyle(this *QProxyStyle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * baseStyle() const

/*
Returns the proxy base style object. If no base style is set on the proxy style, QProxyStyle will create an instance of the application style instead.

See also setBaseStyle() and QStyle.
*/
func (this *QProxyStyle) BaseStyle() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle9baseStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qproxystyle.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseStyle(QStyle *)

/*
Sets the base style that should be proxied.

Ownership of style is transferred to QProxyStyle.

If style is zero, a desktop-dependant style will be assigned automatically.

See also baseStyle().
*/
func (this *QProxyStyle) SetBaseStyle(style QStyle_ITF /*777 QStyle **/) {
	var convArg0 unsafe.Pointer
	if style != nil && style.QStyle_PTR() != nil {
		convArg0 = style.QStyle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle12setBaseStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPrimitive(enum QStyle::PrimitiveElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawPrimitive().
*/
func (this *QProxyStyle) DrawPrimitive(element int, option QStyleOption_ITF /*777 const QStyleOption **/, painter qtgui.QPainter_ITF /*777 QPainter **/, widget QWidget_ITF /*777 const QWidget **/) {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg2 = painter.QPainter_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg3 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPrimitive(enum QStyle::PrimitiveElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawPrimitive().
*/
func (this *QProxyStyle) DrawPrimitive__(element int, option QStyleOption_ITF /*777 const QStyleOption **/, painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg2 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawControl(enum QStyle::ControlElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawControl().
*/
func (this *QProxyStyle) DrawControl(element int, option QStyleOption_ITF /*777 const QStyleOption **/, painter qtgui.QPainter_ITF /*777 QPainter **/, widget QWidget_ITF /*777 const QWidget **/) {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg2 = painter.QPainter_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg3 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawControl(enum QStyle::ControlElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawControl().
*/
func (this *QProxyStyle) DrawControl__(element int, option QStyleOption_ITF /*777 const QStyleOption **/, painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg2 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawComplexControl().
*/
func (this *QProxyStyle) DrawComplexControl(control int, option QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, painter qtgui.QPainter_ITF /*777 QPainter **/, widget QWidget_ITF /*777 const QWidget **/) {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionComplex_PTR() != nil {
		convArg1 = option.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg2 = painter.QPainter_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg3 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawComplexControl().
*/
func (this *QProxyStyle) DrawComplexControl__(control int, option QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionComplex_PTR() != nil {
		convArg1 = option.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg2 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawItemText(QPainter *, const QRect &, int, const QPalette &, _Bool, const QString &, QPalette::ColorRole) const

/*
Reimplemented from QStyle::drawItemText().
*/
func (this *QProxyStyle) DrawItemText(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, flags int, pal qtgui.QPalette_ITF, enabled bool, text string, textRole int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg3 = pal.QPalette_PTR().GetCthis()
	}
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags, convArg3, enabled, convArg5, textRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawItemText(QPainter *, const QRect &, int, const QPalette &, _Bool, const QString &, QPalette::ColorRole) const

/*
Reimplemented from QStyle::drawItemText().
*/
func (this *QProxyStyle) DrawItemText__(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, flags int, pal qtgui.QPalette_ITF, enabled bool, text string) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg3 = pal.QPalette_PTR().GetCthis()
	}
	var tmpArg5 = qtcore.NewQString_5(text)
	var convArg5 = tmpArg5.GetCthis()
	// arg: 6, QPalette::ColorRole=Elaborated, QPalette::ColorRole=Enum,
	textRole := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags, convArg3, enabled, convArg5, textRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawItemPixmap(QPainter *, const QRect &, int, const QPixmap &) const

/*
Reimplemented from QStyle::drawItemPixmap().
*/
func (this *QProxyStyle) DrawItemPixmap(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, alignment int, pixmap qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg3 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeFromContents(enum QStyle::ContentsType, const QStyleOption *, const QSize &, const QWidget *) const

/*
Reimplemented from QStyle::sizeFromContents().
*/
func (this *QProxyStyle) SizeFromContents(type_ int, option QStyleOption_ITF /*777 const QStyleOption **/, size qtcore.QSize_ITF, widget QWidget_ITF /*777 const QWidget **/) *qtcore.QSize /*123*/ {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg2 = size.QSize_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg3 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subElementRect(enum QStyle::SubElement, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::subElementRect().
*/
func (this *QProxyStyle) SubElementRect(element int, option QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subControlRect(enum QStyle::ComplexControl, const QStyleOptionComplex *, enum QStyle::SubControl, const QWidget *) const

/*
Reimplemented from QStyle::subControlRect().
*/
func (this *QProxyStyle) SubControlRect(cc int, opt QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, sc int, widget QWidget_ITF /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg3 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, sc, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect itemTextRect(const QFontMetrics &, const QRect &, int, _Bool, const QString &) const

/*
Reimplemented from QStyle::itemTextRect().
*/
func (this *QProxyStyle) ItemTextRect(fm qtgui.QFontMetrics_ITF, r qtcore.QRect_ITF, flags int, enabled bool, text string) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if fm != nil && fm.QFontMetrics_PTR() != nil {
		convArg0 = fm.QFontMetrics_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg1 = r.QRect_PTR().GetCthis()
	}
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
// [16] QRect itemPixmapRect(const QRect &, int, const QPixmap &) const

/*
Reimplemented from QStyle::itemPixmapRect().
*/
func (this *QProxyStyle) ItemPixmapRect(r qtcore.QRect_ITF, flags int, pixmap qtgui.QPixmap_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg2 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, const QPoint &, const QWidget *) const

/*
Reimplemented from QStyle::hitTestComplexControl().
*/
func (this *QProxyStyle) HitTestComplexControl(control int, option QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, pos qtcore.QPoint_ITF, widget QWidget_ITF /*777 const QWidget **/) int {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionComplex_PTR() != nil {
		convArg1 = option.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg2 = pos.QPoint_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg3 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, const QPoint &, const QWidget *) const

/*
Reimplemented from QStyle::hitTestComplexControl().
*/
func (this *QProxyStyle) HitTestComplexControl__(control int, option QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, pos qtcore.QPoint_ITF) int {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionComplex_PTR() != nil {
		convArg1 = option.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg2 = pos.QPoint_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(enum QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QProxyStyle) StyleHint(hint int, option QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/, returnData QStyleHintReturn_ITF /*777 QStyleHintReturn **/) int {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if returnData != nil && returnData.QStyleHintReturn_PTR() != nil {
		convArg3 = returnData.QStyleHintReturn_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(enum QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QProxyStyle) StyleHint__(hint int) int {
	// arg: 1, const QStyleOption *=Pointer, QStyleOption=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, QStyleHintReturn *=Pointer, QStyleHintReturn=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(enum QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QProxyStyle) StyleHint__1(hint int, option QStyleOption_ITF /*777 const QStyleOption **/) int {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, QStyleHintReturn *=Pointer, QStyleHintReturn=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(enum QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QProxyStyle) StyleHint__2(hint int, option QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) int {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 3, QStyleHintReturn *=Pointer, QStyleHintReturn=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(enum QStyle::PixelMetric, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::pixelMetric().
*/
func (this *QProxyStyle) PixelMetric(metric int, option QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) int {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(enum QStyle::PixelMetric, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::pixelMetric().
*/
func (this *QProxyStyle) PixelMetric__(metric int) int {
	// arg: 1, const QStyleOption *=Pointer, QStyleOption=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(enum QStyle::PixelMetric, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::pixelMetric().
*/
func (this *QProxyStyle) PixelMetric__1(metric int, option QStyleOption_ITF /*777 const QStyleOption **/) int {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::layoutSpacing().

This slot is called by layoutSpacing() to determine the spacing that should be used between control1 and control2 in a layout. orientation specifies whether the controls are laid out side by side or stacked vertically. The option parameter can be used to pass extra information about the parent widget. The widget parameter is optional and can also be used if option is 0.

The default implementation returns -1.

See also combinedLayoutSpacing().
*/
func (this *QProxyStyle) LayoutSpacing(control1 int, control2 int, orientation int, option QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) int {
	var convArg3 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg3 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg4 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::layoutSpacing().

This slot is called by layoutSpacing() to determine the spacing that should be used between control1 and control2 in a layout. orientation specifies whether the controls are laid out side by side or stacked vertically. The option parameter can be used to pass extra information about the parent widget. The widget parameter is optional and can also be used if option is 0.

The default implementation returns -1.

See also combinedLayoutSpacing().
*/
func (this *QProxyStyle) LayoutSpacing__(control1 int, control2 int, orientation int) int {
	// arg: 3, const QStyleOption *=Pointer, QStyleOption=Record,
	var convArg3 unsafe.Pointer
	// arg: 4, const QWidget *=Pointer, QWidget=Record,
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::layoutSpacing().

This slot is called by layoutSpacing() to determine the spacing that should be used between control1 and control2 in a layout. orientation specifies whether the controls are laid out side by side or stacked vertically. The option parameter can be used to pass extra information about the parent widget. The widget parameter is optional and can also be used if option is 0.

The default implementation returns -1.

See also combinedLayoutSpacing().
*/
func (this *QProxyStyle) LayoutSpacing__1(control1 int, control2 int, orientation int, option QStyleOption_ITF /*777 const QStyleOption **/) int {
	var convArg3 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg3 = option.QStyleOption_PTR().GetCthis()
	}
	// arg: 4, const QWidget *=Pointer, QWidget=Record,
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qproxystyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardIcon().

Returns an icon for the given standardIcon.

Reimplement this slot to provide your own icons in a QStyle subclass. The option argument can be used to pass extra information required to find the appropriate icon. The widget argument is optional and can also be used to help find the icon.
*/
func (this *QProxyStyle) StandardIcon(standardIcon int, option QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) *qtgui.QIcon /*123*/ {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardIcon().

Returns an icon for the given standardIcon.

Reimplement this slot to provide your own icons in a QStyle subclass. The option argument can be used to pass extra information required to find the appropriate icon. The widget argument is optional and can also be used to help find the icon.
*/
func (this *QProxyStyle) StandardIcon__(standardIcon int) *qtgui.QIcon /*123*/ {
	// arg: 1, const QStyleOption *=Pointer, QStyleOption=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardIcon().

Returns an icon for the given standardIcon.

Reimplement this slot to provide your own icons in a QStyle subclass. The option argument can be used to pass extra information required to find the appropriate icon. The widget argument is optional and can also be used to help find the icon.
*/
func (this *QProxyStyle) StandardIcon__1(standardIcon int, option QStyleOption_ITF /*777 const QStyleOption **/) *qtgui.QIcon /*123*/ {
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg1 = option.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap standardPixmap(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardPixmap().
*/
func (this *QProxyStyle) StandardPixmap(standardPixmap int, opt QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardPixmap, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap standardPixmap(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardPixmap().
*/
func (this *QProxyStyle) StandardPixmap__(standardPixmap int, opt QStyleOption_ITF /*777 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardPixmap, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap generatedIconPixmap(QIcon::Mode, const QPixmap &, const QStyleOption *) const

/*
Reimplemented from QStyle::generatedIconPixmap().
*/
func (this *QProxyStyle) GeneratedIconPixmap(iconMode int, pixmap qtgui.QPixmap_ITF, opt QStyleOption_ITF /*777 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg2 = opt.QStyleOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QProxyStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), iconMode, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qproxystyle.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QPalette standardPalette() const

/*
Reimplemented from QStyle::standardPalette().
*/
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

/*
Reimplemented from QStyle::polish().
*/
func (this *QProxyStyle) Polish(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:90
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QPalette &)

/*
Reimplemented from QStyle::polish().
*/
func (this *QProxyStyle) Polish_1(pal qtgui.QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg0 = pal.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle6polishER8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:91
// index:2
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QApplication *)

/*
Reimplemented from QStyle::polish().
*/
func (this *QProxyStyle) Polish_2(app QApplication_ITF /*777 QApplication **/) {
	var convArg0 unsafe.Pointer
	if app != nil && app.QApplication_PTR() != nil {
		convArg0 = app.QApplication_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP12QApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QWidget *)

/*
Reimplemented from QStyle::unpolish().
*/
func (this *QProxyStyle) Unpolish(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:94
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QApplication *)

/*
Reimplemented from QStyle::unpolish().
*/
func (this *QProxyStyle) Unpolish_1(app QApplication_ITF /*777 QApplication **/) {
	var convArg0 unsafe.Pointer
	if app != nil && app.QApplication_PTR() != nil {
		convArg0 = app.QApplication_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP12QApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QProxyStyle) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
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
		log.Println(123)
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
