package qtwidgets

// /usr/include/qt/QtWidgets/qdockwidget.h
// #include <qdockwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 37
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QDockWidget struct {
	*QWidget
}

func (this *QDockWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QDockWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQDockWidgetFromPointer(cthis unsafe.Pointer) *QDockWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDockWidget{bcthis0}
}
func (*QDockWidget) NewFromPointer(cthis unsafe.Pointer) *QDockWidget {
	return NewQDockWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDockWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdockwidget.h:66
// index:0
// Public
// void QDockWidget(const class QString &, class QWidget *, Qt::WindowFlags)
func NewQDockWidget(title *qtcore.QString, parent *QWidget /*777 QWidget **/, flags int) *QDockWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = title.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidgetC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdockwidget.h:68
// index:1
// Public
// void QDockWidget(class QWidget *, Qt::WindowFlags)
func NewQDockWidget_1(parent *QWidget /*777 QWidget **/, flags int) *QDockWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidgetC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(cthis)
	return gothis
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
func (this *QDockWidget) Widget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdockwidget.h:72
// index:0
// Public
// void setWidget(class QWidget *)
func (this *QDockWidget) SetWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:90
// index:0
// Public
// QDockWidget::DockWidgetFeatures features()
func (this *QDockWidget) Features() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget8featuresEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:92
// index:0
// Public
// void setFloating(_Bool)
func (this *QDockWidget) SetFloating(floating bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11setFloatingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), floating)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:93
// index:0
// Public inline
// bool isFloating()
func (this *QDockWidget) IsFloating() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10isFloatingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdockwidget.h:96
// index:0
// Public
// Qt::DockWidgetAreas allowedAreas()
func (this *QDockWidget) AllowedAreas() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget12allowedAreasEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:98
// index:0
// Public
// void setTitleBarWidget(class QWidget *)
func (this *QDockWidget) SetTitleBarWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17setTitleBarWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:99
// index:0
// Public
// QWidget * titleBarWidget()
func (this *QDockWidget) TitleBarWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget14titleBarWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdockwidget.h:101
// index:0
// Public inline
// bool isAreaAllowed(Qt::DockWidgetArea)
func (this *QDockWidget) IsAreaAllowed(area int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget13isAreaAllowedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdockwidget.h:105
// index:0
// Public
// QAction * toggleViewAction()
func (this *QDockWidget) ToggleViewAction() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget16toggleViewActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdockwidget.h:110
// index:0
// Public
// void topLevelChanged(_Bool)
func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget15topLevelChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), topLevel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:112
// index:0
// Public
// void visibilityChanged(_Bool)
func (this *QDockWidget) VisibilityChanged(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17visibilityChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:113
// index:0
// Public
// void dockLocationChanged(Qt::DockWidgetArea)
func (this *QDockWidget) DockLocationChanged(area int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget19dockLocationChangedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:116
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QDockWidget) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:117
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QDockWidget) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:118
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QDockWidget) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:119
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QDockWidget) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdockwidget.h:120
// index:0
// Protected
// void initStyleOption(class QStyleOptionDockWidget *)
func (this *QDockWidget) InitStyleOption(option *QStyleOptionDockWidget /*777 QStyleOptionDockWidget **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QDockWidget__DockWidgetFeature = int

const QDockWidget__DockWidgetClosable QDockWidget__DockWidgetFeature = 1
const QDockWidget__DockWidgetMovable QDockWidget__DockWidgetFeature = 2
const QDockWidget__DockWidgetFloatable QDockWidget__DockWidgetFeature = 4
const QDockWidget__DockWidgetVerticalTitleBar QDockWidget__DockWidgetFeature = 8
const QDockWidget__DockWidgetFeatureMask QDockWidget__DockWidgetFeature = 15
const QDockWidget__AllDockWidgetFeatures QDockWidget__DockWidgetFeature = 7
const QDockWidget__NoDockWidgetFeatures QDockWidget__DockWidgetFeature = 0
const QDockWidget__Reserved QDockWidget__DockWidgetFeature = 255

//  body block end
