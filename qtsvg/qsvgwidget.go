package qtsvg

// /usr/include/qt/QtSvg/qsvgwidget.h
// #include <qsvgwidget.h>
// #include <QtSvg>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"

//  ext block end

//  body block begin

// void paintEvent(QPaintEvent *)
func (this *QSvgWidget) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

/*

 */
type QSvgWidget struct {
	*qtwidgets.QWidget
}
type QSvgWidget_ITF interface {
	qtwidgets.QWidget_ITF
	QSvgWidget_PTR() *QSvgWidget
}

func (ptr *QSvgWidget) QSvgWidget_PTR() *QSvgWidget { return ptr }

func (this *QSvgWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QSvgWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = qtwidgets.NewQWidgetFromPointer(cthis)
}
func NewQSvgWidgetFromPointer(cthis unsafe.Pointer) *QSvgWidget {
	bcthis0 := qtwidgets.NewQWidgetFromPointer(cthis)
	return &QSvgWidget{bcthis0}
}
func (*QSvgWidget) NewFromPointer(cthis unsafe.Pointer) *QSvgWidget {
	return NewQSvgWidgetFromPointer(cthis)
}

// /usr/include/qt/QtSvg/qsvgwidget.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSvgWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSvgWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtSvg/qsvgwidget.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSvgWidget(QWidget *)

/*
Constructs a new SVG display widget with the given parent.
*/
func (*QSvgWidget) NewForInherit(parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QSvgWidget {
	return NewQSvgWidget(parent)
}
func NewQSvgWidget(parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QSvgWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgWidget")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgwidget.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSvgWidget(QWidget *)

/*
Constructs a new SVG display widget with the given parent.
*/
func (*QSvgWidget) NewForInheritp() *QSvgWidget {
	return NewQSvgWidgetp()
}
func NewQSvgWidgetp() *QSvgWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgWidget")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgwidget.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSvgWidget(const QString &, QWidget *)

/*
Constructs a new SVG display widget with the given parent.
*/
func (*QSvgWidget) NewForInherit1(file string, parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QSvgWidget {
	return NewQSvgWidget1(file, parent)
}
func NewQSvgWidget1(file string, parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QSvgWidget {
	var tmpArg0 = qtcore.NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidgetC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgWidget")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgwidget.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSvgWidget(const QString &, QWidget *)

/*
Constructs a new SVG display widget with the given parent.
*/
func (*QSvgWidget) NewForInherit1p(file string) *QSvgWidget {
	return NewQSvgWidget1p(file)
}
func NewQSvgWidget1p(file string) *QSvgWidget {
	var tmpArg0 = qtcore.NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidgetC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSvgWidget")
	return gothis
}

// /usr/include/qt/QtSvg/qsvgwidget.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSvgWidget()

/*

 */
func DeleteQSvgWidget(this *QSvgWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtSvg/qsvgwidget.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QSvgRenderer * renderer() const

/*
Returns the renderer used to display the contents of the widget.
*/
func (this *QSvgWidget) Renderer() *QSvgRenderer /*777 QSvgRenderer **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSvgWidget8rendererEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSvgRendererFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtSvg/qsvgwidget.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QSvgWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSvgWidget8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtSvg/qsvgwidget.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(const QString &)

/*
Loads the contents of the specified SVG file and updates the widget.
*/
func (this *QSvgWidget) Load(file string) {
	var tmpArg0 = qtcore.NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidget4loadERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgwidget.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(const QByteArray &)

/*
Loads the contents of the specified SVG file and updates the widget.
*/
func (this *QSvgWidget) Load1(contents qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if contents != nil && contents.QByteArray_PTR() != nil {
		convArg0 = contents.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidget4loadERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvgwidget.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QSvgWidget) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSvgWidget10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtwidgets.KeepMe()
	}
}

//  keep block end
