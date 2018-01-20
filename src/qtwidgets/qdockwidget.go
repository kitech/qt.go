//  header block begin
// /usr/include/qt/QtWidgets/qdockwidget.h
// #include <qdockwidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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

// /usr/include/qt/QtWidgets/qdockwidget.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDockWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:66
// index:0
// void QDockWidget(const class QString &, class QWidget *, Qt::WindowFlags)
func NewQDockWidget(title unsafe.Pointer, parent unsafe.Pointer, flags int) *QDockWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidgetC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, title, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(cthis)
	return gothis
}
func NewQDockWidgetFromPointer(cthis unsafe.Pointer) *QDockWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDockWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qdockwidget.h:68
// index:1
// void QDockWidget(class QWidget *, Qt::WindowFlags)
func NewQDockWidget_1(parent unsafe.Pointer, flags int) *QDockWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidgetC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(cthis)
	return gothis
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget6widgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:72
// index:0
// void setWidget(class QWidget *)
func (this *QDockWidget) SetWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:89
// index:0
// void setFeatures(QDockWidget::DockWidgetFeatures)
func (this *QDockWidget) SetFeatures(features int) {
	// 0: (, QFlags<QDockWidget::DockWidgetFeature> features), (features)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11setFeaturesE6QFlagsINS_17DockWidgetFeatureEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), features)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:90
// index:0
// QDockWidget::DockWidgetFeatures features()
func (this *QDockWidget) Features() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget8featuresEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:92
// index:0
// void setFloating(_Bool)
func (this *QDockWidget) SetFloating(floating bool) {
	// 0: (, floating bool), (&floating)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11setFloatingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &floating)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:93
// index:0
// inline
// bool isFloating()
func (this *QDockWidget) IsFloating() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget10isFloatingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:95
// index:0
// void setAllowedAreas(Qt::DockWidgetAreas)
func (this *QDockWidget) SetAllowedAreas(areas int) {
	// 0: (, QFlags<Qt::DockWidgetArea> areas), (&areas)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget15setAllowedAreasE6QFlagsIN2Qt14DockWidgetAreaEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &areas)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:96
// index:0
// Qt::DockWidgetAreas allowedAreas()
func (this *QDockWidget) AllowedAreas() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget12allowedAreasEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:98
// index:0
// void setTitleBarWidget(class QWidget *)
func (this *QDockWidget) SetTitleBarWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17setTitleBarWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:99
// index:0
// QWidget * titleBarWidget()
func (this *QDockWidget) TitleBarWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget14titleBarWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:101
// index:0
// inline
// bool isAreaAllowed(Qt::DockWidgetArea)
func (this *QDockWidget) IsAreaAllowed(area int) {
	// 0: (, area Qt::DockWidgetArea), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget13isAreaAllowedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:105
// index:0
// QAction * toggleViewAction()
func (this *QDockWidget) ToggleViewAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget16toggleViewActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:109
// index:0
// void featuresChanged(class QDockWidget::DockWidgetFeatures)
func (this *QDockWidget) FeaturesChanged(features int) {
	// 0: (, QFlags<QDockWidget::DockWidgetFeature> features), (&features)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget15featuresChangedE6QFlagsINS_17DockWidgetFeatureEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &features)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:110
// index:0
// void topLevelChanged(_Bool)
func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	// 0: (, topLevel bool), (&topLevel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget15topLevelChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &topLevel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:111
// index:0
// void allowedAreasChanged(Qt::DockWidgetAreas)
func (this *QDockWidget) AllowedAreasChanged(allowedAreas int) {
	// 0: (, QFlags<Qt::DockWidgetArea> allowedAreas), (&allowedAreas)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget19allowedAreasChangedE6QFlagsIN2Qt14DockWidgetAreaEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &allowedAreas)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:112
// index:0
// void visibilityChanged(_Bool)
func (this *QDockWidget) VisibilityChanged(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget17visibilityChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:113
// index:0
// void dockLocationChanged(Qt::DockWidgetArea)
func (this *QDockWidget) DockLocationChanged(area int) {
	// 0: (, area Qt::DockWidgetArea), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget19dockLocationChangedEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:116
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QDockWidget) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:117
// index:0
// virtual
// void closeEvent(class QCloseEvent *)
func (this *QDockWidget) CloseEvent(event unsafe.Pointer) {
	// 0: (, event QCloseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:118
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QDockWidget) PaintEvent(event unsafe.Pointer) {
	// 0: (, event QPaintEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:119
// index:0
// virtual
// bool event(class QEvent *)
func (this *QDockWidget) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDockWidget5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:120
// index:0
// void initStyleOption(class QStyleOptionDockWidget *)
func (this *QDockWidget) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionDockWidget *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
