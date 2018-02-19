package qtwidgets

// /usr/include/qt/QtWidgets/qsplashscreen.h
// #include <qsplashscreen.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
func (this *QSplashScreen) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void drawContents(class QPainter *)
func (this *QSplashScreen) InheritDrawContents(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawContents", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QSplashScreen) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

type QSplashScreen struct {
	*QWidget
}
type QSplashScreen_ITF interface {
	QWidget_ITF
	QSplashScreen_PTR() *QSplashScreen
}

func (ptr *QSplashScreen) QSplashScreen_PTR() *QSplashScreen { return ptr }

func (this *QSplashScreen) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QSplashScreen) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQSplashScreenFromPointer(cthis unsafe.Pointer) *QSplashScreen {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplashScreen{bcthis0}
}
func (*QSplashScreen) NewFromPointer(cthis unsafe.Pointer) *QSplashScreen {
	return NewQSplashScreenFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QSplashScreen) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSplashScreen10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(const QPixmap &, Qt::WindowFlags)
func NewQSplashScreen(pixmap qtgui.QPixmap_ITF, f int) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2ERK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(const QPixmap &, Qt::WindowFlags)
func NewQSplashScreen__() *QSplashScreen {
	// arg: 0, const QPixmap &=LValueReference, QPixmap=Record,
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2ERK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(const QPixmap &, Qt::WindowFlags)
func NewQSplashScreen__1(pixmap qtgui.QPixmap_ITF) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2ERK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(QWidget *, const QPixmap &, Qt::WindowFlags)
func NewQSplashScreen_1(parent QWidget_ITF /*777 QWidget **/, pixmap qtgui.QPixmap_ITF, f int) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2EP7QWidgetRK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(QWidget *, const QPixmap &, Qt::WindowFlags)
func NewQSplashScreen_1_(parent QWidget_ITF /*777 QWidget **/) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QPixmap &=LValueReference, QPixmap=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2EP7QWidgetRK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(QWidget *, const QPixmap &, Qt::WindowFlags)
func NewQSplashScreen_1_1(parent QWidget_ITF /*777 QWidget **/, pixmap qtgui.QPixmap_ITF) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2EP7QWidgetRK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSplashScreen()
func DeleteQSplashScreen(this *QSplashScreen) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &)
func (this *QSplashScreen) SetPixmap(pixmap qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen9setPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:62
// index:0
// Public Visibility=Default Availability=Available
// [32] const QPixmap pixmap() const
func (this *QSplashScreen) Pixmap() *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSplashScreen6pixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finish(QWidget *)
func (this *QSplashScreen) Finish(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen6finishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void repaint()
func (this *QSplashScreen) Repaint() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen7repaintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString message() const
func (this *QSplashScreen) Message() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSplashScreen7messageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int, const QColor &)
func (this *QSplashScreen) ShowMessage(message string, alignment int, color qtgui.QColor_ITF) {
	var tmpArg0 = qtcore.NewQString_5(message)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg2 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int, const QColor &)
func (this *QSplashScreen) ShowMessage__(message string) {
	var tmpArg0 = qtcore.NewQString_5(message)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	alignment := 0 /*Qt::AlignLeft*/
	// arg: 2, const QColor &=LValueReference, QColor=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int, const QColor &)
func (this *QSplashScreen) ShowMessage__1(message string, alignment int) {
	var tmpArg0 = qtcore.NewQString_5(message)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, const QColor &=LValueReference, QColor=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMessage()
func (this *QSplashScreen) ClearMessage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen12clearMessageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void messageChanged(const QString &)
func (this *QSplashScreen) MessageChanged(message string) {
	var tmpArg0 = qtcore.NewQString_5(message)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen14messageChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSplashScreen) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawContents(QPainter *)
func (this *QSplashScreen) DrawContents(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen12drawContentsEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:78
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QSplashScreen) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
}

//  keep block end
