package qtwidgets

// /usr/include/qt/QtWidgets/qstylepainter.h
// #include <qstylepainter.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

/*

 */
type QStylePainter struct {
	*qtgui.QPainter
}
type QStylePainter_ITF interface {
	qtgui.QPainter_ITF
	QStylePainter_PTR() *QStylePainter
}

func (ptr *QStylePainter) QStylePainter_PTR() *QStylePainter { return ptr }

func (this *QStylePainter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPainter.GetCthis()
	}
}
func (this *QStylePainter) SetCthis(cthis unsafe.Pointer) {
	this.QPainter = qtgui.NewQPainterFromPointer(cthis)
}
func NewQStylePainterFromPointer(cthis unsafe.Pointer) *QStylePainter {
	bcthis0 := qtgui.NewQPainterFromPointer(cthis)
	return &QStylePainter{bcthis0}
}
func (*QStylePainter) NewFromPointer(cthis unsafe.Pointer) *QStylePainter {
	return NewQStylePainterFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:54
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStylePainter()

/*
Constructs a QStylePainter.
*/
func (*QStylePainter) NewForInherit() *QStylePainter {
	return NewQStylePainter()
}
func NewQStylePainter() *QStylePainter {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStylePainterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStylePainter)
	return gothis
}

// /usr/include/qt/QtWidgets/qstylepainter.h:55
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QStylePainter(QWidget *)

/*
Constructs a QStylePainter.
*/
func (*QStylePainter) NewForInherit1(w QWidget_ITF /*777 QWidget **/) *QStylePainter {
	return NewQStylePainter1(w)
}
func NewQStylePainter1(w QWidget_ITF /*777 QWidget **/) *QStylePainter {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainterC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStylePainterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStylePainter)
	return gothis
}

// /usr/include/qt/QtWidgets/qstylepainter.h:56
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QStylePainter(QPaintDevice *, QWidget *)

/*
Constructs a QStylePainter.
*/
func (*QStylePainter) NewForInherit2(pd qtgui.QPaintDevice_ITF /*777 QPaintDevice **/, w QWidget_ITF /*777 QWidget **/) *QStylePainter {
	return NewQStylePainter2(pd, w)
}
func NewQStylePainter2(pd qtgui.QPaintDevice_ITF /*777 QPaintDevice **/, w QWidget_ITF /*777 QWidget **/) *QStylePainter {
	var convArg0 unsafe.Pointer
	if pd != nil && pd.QPaintDevice_PTR() != nil {
		convArg0 = pd.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg1 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainterC2EP12QPaintDeviceP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStylePainterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStylePainter)
	return gothis
}

// /usr/include/qt/QtWidgets/qstylepainter.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool begin(QWidget *)

/*
Begin painting operations on the specified widget. Returns true if the painter is ready to use; otherwise returns false.

This is automatically called by the constructor that takes a QWidget.
*/
func (this *QStylePainter) Begin(w QWidget_ITF /*777 QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter5beginEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstylepainter.h:58
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool begin(QPaintDevice *, QWidget *)

/*
Begin painting operations on the specified widget. Returns true if the painter is ready to use; otherwise returns false.

This is automatically called by the constructor that takes a QWidget.
*/
func (this *QStylePainter) Begin1(pd qtgui.QPaintDevice_ITF /*777 QPaintDevice **/, w QWidget_ITF /*777 QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if pd != nil && pd.QPaintDevice_PTR() != nil {
		convArg0 = pd.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg1 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstylepainter.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawPrimitive(QStyle::PrimitiveElement, const QStyleOption &)

/*
Use the widget's style to draw a primitive element pe specified by QStyleOption option.

See also QStyle::drawPrimitive().
*/
func (this *QStylePainter) DrawPrimitive(pe int, opt QStyleOption_ITF) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter13drawPrimitiveEN6QStyle16PrimitiveElementERK12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pe, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawControl(QStyle::ControlElement, const QStyleOption &)

/*
Use the widget's style to draw a control element ce specified by QStyleOption option.

See also QStyle::drawControl().
*/
func (this *QStylePainter) DrawControl(ce int, opt QStyleOption_ITF) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter11drawControlEN6QStyle14ControlElementERK12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ce, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawComplexControl(QStyle::ComplexControl, const QStyleOptionComplex &)

/*
Use the widget's style to draw a complex control cc specified by the QStyleOptionComplex option.

See also QStyle::drawComplexControl().
*/
func (this *QStylePainter) DrawComplexControl(cc int, opt QStyleOptionComplex_ITF) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter18drawComplexControlEN6QStyle14ComplexControlERK19QStyleOptionComplex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawItemText(const QRect &, int, const QPalette &, bool, const QString &, QPalette::ColorRole)

/*
Draws the text in rectangle rect and palette pal. The text is aligned and wrapped according to flags.

The pen color is specified with textRole. The enabled bool indicates whether or not the item is enabled; when reimplementing this bool should influence how the item is drawn.

See also QStyle::drawItemText() and Qt::Alignment.
*/
func (this *QStylePainter) DrawItemText(r qtcore.QRect_ITF, flags int, pal qtgui.QPalette_ITF, enabled bool, text string, textRole int) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg2 = pal.QPalette_PTR().GetCthis()
	}
	var tmpArg4 = qtcore.NewQString5(text)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter12drawItemTextERK5QRectiRK8QPalettebRK7QStringNS3_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, enabled, convArg4, textRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawItemText(const QRect &, int, const QPalette &, bool, const QString &, QPalette::ColorRole)

/*
Draws the text in rectangle rect and palette pal. The text is aligned and wrapped according to flags.

The pen color is specified with textRole. The enabled bool indicates whether or not the item is enabled; when reimplementing this bool should influence how the item is drawn.

See also QStyle::drawItemText() and Qt::Alignment.
*/
func (this *QStylePainter) DrawItemTextp(r qtcore.QRect_ITF, flags int, pal qtgui.QPalette_ITF, enabled bool, text string) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg2 = pal.QPalette_PTR().GetCthis()
	}
	var tmpArg4 = qtcore.NewQString5(text)
	var convArg4 = tmpArg4.GetCthis()
	// arg: 5, QPalette::ColorRole=Elaborated, QPalette::ColorRole=Enum, , Invalid
	textRole := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter12drawItemTextERK5QRectiRK8QPalettebRK7QStringNS3_9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, enabled, convArg4, textRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void drawItemPixmap(const QRect &, int, const QPixmap &)

/*
Draws the pixmap in rectangle rect. The pixmap is aligned according to flags.

See also QStyle::drawItemPixmap() and Qt::Alignment.
*/
func (this *QStylePainter) DrawItemPixmap(r qtcore.QRect_ITF, flags int, pixmap qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg2 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStyle * style() const

/*
Return the current style used by the QStylePainter.
*/
func (this *QStylePainter) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStylePainter5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQStylePainter(this *QStylePainter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStylePainterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11331() {
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
