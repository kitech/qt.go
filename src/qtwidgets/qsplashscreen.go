//  header block begin
// /usr/include/qt/QtWidgets/qsplashscreen.h
// #include <qsplashscreen.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QSplashScreen struct {
	*QWidget
}

func (this *QSplashScreen) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSplashScreen) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSplashScreen10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:57
// index:0
// void QSplashScreen(const class QPixmap &, Qt::WindowFlags)
func NewQSplashScreen(pixmap unsafe.Pointer, f int) *QSplashScreen {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreenC2ERK7QPixmap6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, pixmap, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(cthis)
	return gothis
}
func NewQSplashScreenFromPointer(cthis unsafe.Pointer) *QSplashScreen {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplashScreen{bcthis0}
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:58
// index:1
// void QSplashScreen(class QWidget *, const class QPixmap &, Qt::WindowFlags)
func NewQSplashScreen_1(parent unsafe.Pointer, pixmap unsafe.Pointer, f int) *QSplashScreen {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreenC2EP7QWidgetRK7QPixmap6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, pixmap, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:59
// index:0
// virtual
// void ~QSplashScreen()
func DeleteQSplashScreen(*QSplashScreen) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreenD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:61
// index:0
// void setPixmap(const class QPixmap &)
func (this *QSplashScreen) SetPixmap(pixmap unsafe.Pointer) {
	// 0: (, pixmap const QPixmap &), (pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:62
// index:0
// const QPixmap pixmap()
func (this *QSplashScreen) Pixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSplashScreen6pixmapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:63
// index:0
// void finish(class QWidget *)
func (this *QSplashScreen) Finish(w unsafe.Pointer) {
	// 0: (, w QWidget *), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen6finishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:64
// index:0
// void repaint()
func (this *QSplashScreen) Repaint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen7repaintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:65
// index:0
// QString message()
func (this *QSplashScreen) Message() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSplashScreen7messageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// void showMessage(const class QString &, int, const class QColor &)
func (this *QSplashScreen) ShowMessage(message unsafe.Pointer, alignment int, color unsafe.Pointer) {
	// 0: (, message const QString &, alignment int, color const QColor &), (message, &alignment, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), message, &alignment, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:70
// index:0
// void clearMessage()
func (this *QSplashScreen) ClearMessage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen12clearMessageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:73
// index:0
// void messageChanged(const class QString &)
func (this *QSplashScreen) MessageChanged(message unsafe.Pointer) {
	// 0: (, message const QString &), (message)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen14messageChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), message)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:76
// index:0
// virtual
// bool event(class QEvent *)
func (this *QSplashScreen) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:77
// index:0
// virtual
// void drawContents(class QPainter *)
func (this *QSplashScreen) DrawContents(painter unsafe.Pointer) {
	// 0: (, painter QPainter *), (painter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen12drawContentsEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:78
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QSplashScreen) MousePressEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
