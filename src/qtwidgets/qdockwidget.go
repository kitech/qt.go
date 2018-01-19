//  header block begin
// /usr/include/qt/QtWidgets/qdockwidget.h
// #include <qdockwidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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
type QDockWidget struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qdockwidget.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDockWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:69
// index:0
// virtual
// void ~QDockWidget()
func DeleteQDockWidget(*QDockWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:71
// index:0
// QWidget * widget()
func (this *QDockWidget) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget6widgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:72
// index:0
// void setWidget(class QWidget *)
func (this *QDockWidget) SetWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:90
// index:0
// QDockWidget::DockWidgetFeatures features()
func (this *QDockWidget) Features() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget8featuresEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:92
// index:0
// void setFloating(_Bool)
func (this *QDockWidget) SetFloating(floating bool) {
	// 0: (, bool floating), (&floating)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11setFloatingEb", ffiqt.FFI_TYPE_VOID, this.cthis, &floating)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:93
// index:0
// inline
// bool isFloating()
func (this *QDockWidget) IsFloating() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10isFloatingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:96
// index:0
// Qt::DockWidgetAreas allowedAreas()
func (this *QDockWidget) AllowedAreas() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget12allowedAreasEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:98
// index:0
// void setTitleBarWidget(class QWidget *)
func (this *QDockWidget) SetTitleBarWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17setTitleBarWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:99
// index:0
// QWidget * titleBarWidget()
func (this *QDockWidget) TitleBarWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget14titleBarWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:101
// index:0
// inline
// bool isAreaAllowed(Qt::DockWidgetArea)
func (this *QDockWidget) IsAreaAllowed(area int) {
	// 0: (, Qt::DockWidgetArea area), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget13isAreaAllowedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.cthis, &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:105
// index:0
// QAction * toggleViewAction()
func (this *QDockWidget) ToggleViewAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget16toggleViewActionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:110
// index:0
// void topLevelChanged(_Bool)
func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	// 0: (, bool topLevel), (&topLevel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget15topLevelChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &topLevel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:112
// index:0
// void visibilityChanged(_Bool)
func (this *QDockWidget) VisibilityChanged(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17visibilityChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:113
// index:0
// void dockLocationChanged(Qt::DockWidgetArea)
func (this *QDockWidget) DockLocationChanged(area int) {
	// 0: (, Qt::DockWidgetArea area), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget19dockLocationChangedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.cthis, &area)
	gopp.ErrPrint(err, rv)
}

//  body block end
