package qtwidgets

// /usr/include/qt/QtWidgets/qstatusbar.h
// #include <qstatusbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void showEvent(class QShowEvent *)
func (this *QStatusBar) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QStatusBar) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QStatusBar) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void reformat()
func (this *QStatusBar) InheritReformat(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "reformat", f)
}

// void hideOrShow()
func (this *QStatusBar) InheritHideOrShow(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideOrShow", f)
}

// bool event(class QEvent *)
func (this *QStatusBar) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QStatusBar struct {
	*QWidget
}

func (this *QStatusBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QStatusBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQStatusBarFromPointer(cthis unsafe.Pointer) *QStatusBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QStatusBar{bcthis0}
}
func (*QStatusBar) NewFromPointer(cthis unsafe.Pointer) *QStatusBar {
	return NewQStatusBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStatusBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStatusBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusBar(QWidget *)
func NewQStatusBar(parent *QWidget /*777 QWidget **/) *QStatusBar {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStatusBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStatusBar()
func DeleteQStatusBar(this *QStatusBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int)
func (this *QStatusBar) AddWidget(widget *QWidget /*777 QWidget **/, stretch int) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar9addWidgetEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *, int)
func (this *QStatusBar) InsertWidget(index int, widget *QWidget /*777 QWidget **/, stretch int) int {
	var convArg1 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar12insertWidgetEiP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPermanentWidget(QWidget *, int)
func (this *QStatusBar) AddPermanentWidget(widget *QWidget /*777 QWidget **/, stretch int) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertPermanentWidget(int, QWidget *, int)
func (this *QStatusBar) InsertPermanentWidget(index int, widget *QWidget /*777 QWidget **/, stretch int) int {
	var convArg1 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)
func (this *QStatusBar) RemoveWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar12removeWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeGripEnabled(_Bool)
func (this *QStatusBar) SetSizeGripEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar18setSizeGripEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSizeGripEnabled()
func (this *QStatusBar) IsSizeGripEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStatusBar17isSizeGripEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstatusbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString currentMessage()
func (this *QStatusBar) CurrentMessage() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStatusBar14currentMessageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qstatusbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int)
func (this *QStatusBar) ShowMessage(text string, timeout int) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar11showMessageERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMessage()
func (this *QStatusBar) ClearMessage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar12clearMessageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void messageChanged(const QString &)
func (this *QStatusBar) MessageChanged(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar14messageChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:82
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QStatusBar) ShowEvent(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QStatusBar) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QStatusBar) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:87
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void reformat()
func (this *QStatusBar) Reformat() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar8reformatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:88
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void hideOrShow()
func (this *QStatusBar) HideOrShow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar10hideOrShowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QStatusBar) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStatusBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
