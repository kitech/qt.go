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
	*QWidget
}

func (this *QDockWidget) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQDockWidgetFromPointer(cthis unsafe.Pointer) *QDockWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDockWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qdockwidget.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDockWidget) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:69
// index:0
// Public virtual
// void ~QDockWidget()
func DeleteQDockWidget(*QDockWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:71
// index:0
// Public
// QWidget * widget()
func (this *QDockWidget) Widget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:72
// index:0
// Public
// void setWidget(class QWidget *)
func (this *QDockWidget) SetWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:90
// index:0
// Public
// QDockWidget::DockWidgetFeatures features()
func (this *QDockWidget) Features() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget8featuresEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:92
// index:0
// Public
// void setFloating(_Bool)
func (this *QDockWidget) SetFloating(floating bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11setFloatingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &floating)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:93
// index:0
// Public inline
// bool isFloating()
func (this *QDockWidget) IsFloating() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10isFloatingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:96
// index:0
// Public
// Qt::DockWidgetAreas allowedAreas()
func (this *QDockWidget) AllowedAreas() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget12allowedAreasEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:98
// index:0
// Public
// void setTitleBarWidget(class QWidget *)
func (this *QDockWidget) SetTitleBarWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17setTitleBarWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:99
// index:0
// Public
// QWidget * titleBarWidget()
func (this *QDockWidget) TitleBarWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget14titleBarWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:101
// index:0
// Public inline
// bool isAreaAllowed(Qt::DockWidgetArea)
func (this *QDockWidget) IsAreaAllowed(area int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget13isAreaAllowedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &area)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:105
// index:0
// Public
// QAction * toggleViewAction()
func (this *QDockWidget) ToggleViewAction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget16toggleViewActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:110
// index:0
// Public
// void topLevelChanged(_Bool)
func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget15topLevelChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &topLevel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:112
// index:0
// Public
// void visibilityChanged(_Bool)
func (this *QDockWidget) VisibilityChanged(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17visibilityChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:113
// index:0
// Public
// void dockLocationChanged(Qt::DockWidgetArea)
func (this *QDockWidget) DockLocationChanged(area int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget19dockLocationChangedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:116
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QDockWidget) ChangeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:117
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QDockWidget) CloseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:118
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QDockWidget) PaintEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:119
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QDockWidget) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdockwidget.h:120
// index:0
// Protected
// void initStyleOption(class QStyleOptionDockWidget *)
func (this *QDockWidget) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
