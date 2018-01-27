package qtwidgets

// /usr/include/qt/QtWidgets/qstylepainter.h
// #include <qstylepainter.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
type QStylePainter struct {
	*qtgui.QPainter
}

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
// Public inline
// void QStylePainter()
func NewQStylePainter() *QStylePainter {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStylePainterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstylepainter.h:55
// index:1
// Public inline
// void QStylePainter(QWidget *)
func NewQStylePainter_1(w *QWidget /*777 QWidget **/) *QStylePainter {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainterC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStylePainterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstylepainter.h:56
// index:2
// Public inline
// void QStylePainter(QPaintDevice *, QWidget *)
func NewQStylePainter_2(pd *qtgui.QPaintDevice /*777 QPaintDevice **/, w *QWidget /*777 QWidget **/) *QStylePainter {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = pd.GetCthis()
	var convArg1 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainterC2EP12QPaintDeviceP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStylePainterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstylepainter.h:57
// index:0
// Public inline
// bool begin(QWidget *)
func (this *QStylePainter) Begin(w *QWidget /*777 QWidget **/) bool {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter5beginEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstylepainter.h:58
// index:1
// Public inline
// bool begin(QPaintDevice *, QWidget *)
func (this *QStylePainter) Begin_1(pd *qtgui.QPaintDevice /*777 QPaintDevice **/, w *QWidget /*777 QWidget **/) bool {
	var convArg0 = pd.GetCthis()
	var convArg1 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstylepainter.h:64
// index:0
// Public inline
// void drawPrimitive(QStyle::PrimitiveElement, const QStyleOption &)
func (this *QStylePainter) DrawPrimitive(pe int, opt *QStyleOption) {
	var convArg1 = opt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter13drawPrimitiveEN6QStyle16PrimitiveElementERK12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pe, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:65
// index:0
// Public inline
// void drawControl(QStyle::ControlElement, const QStyleOption &)
func (this *QStylePainter) DrawControl(ce int, opt *QStyleOption) {
	var convArg1 = opt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter11drawControlEN6QStyle14ControlElementERK12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ce, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:66
// index:0
// Public inline
// void drawComplexControl(QStyle::ComplexControl, const QStyleOptionComplex &)
func (this *QStylePainter) DrawComplexControl(cc int, opt *QStyleOptionComplex) {
	var convArg1 = opt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter18drawComplexControlEN6QStyle14ComplexControlERK19QStyleOptionComplex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:67
// index:0
// Public inline
// void drawItemText(const QRect &, int, const QPalette &, bool, const QString &, QPalette::ColorRole)
func (this *QStylePainter) DrawItemText(r *qtcore.QRect, flags int, pal *qtgui.QPalette, enabled bool, text *qtcore.QString, textRole int) {
	var convArg0 = r.GetCthis()
	var convArg2 = pal.GetCthis()
	var convArg4 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter12drawItemTextERK5QRectiRK8QPalettebRK7QStringNS3_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, enabled, convArg4, textRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:69
// index:0
// Public inline
// void drawItemPixmap(const QRect &, int, const QPixmap &)
func (this *QStylePainter) DrawItemPixmap(r *qtcore.QRect, flags int, pixmap *qtgui.QPixmap) {
	var convArg0 = r.GetCthis()
	var convArg2 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:70
// index:0
// Public inline
// QStyle * style()
func (this *QStylePainter) Style() *QStyle /*777 QStyle **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStylePainter5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
