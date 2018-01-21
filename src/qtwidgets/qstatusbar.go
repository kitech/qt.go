package qtwidgets

// /usr/include/qt/QtWidgets/qstatusbar.h
// #include <qstatusbar.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
func NewQStatusBarFromPointer(cthis unsafe.Pointer) *QStatusBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QStatusBar{bcthis0}
}

// /usr/include/qt/QtWidgets/qstatusbar.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStatusBar) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStatusBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public
// void QStatusBar(class QWidget *)
func NewQStatusBar(parent *QWidget /*444 QWidget **/) *QStatusBar {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStatusBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:60
// index:0
// Public virtual
// void ~QStatusBar()
func DeleteQStatusBar(*QStatusBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:62
// index:0
// Public
// void addWidget(class QWidget *, int)
func (this *QStatusBar) AddWidget(widget *QWidget /*444 QWidget **/, stretch int) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar9addWidgetEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:63
// index:0
// Public
// int insertWidget(int, class QWidget *, int)
func (this *QStatusBar) InsertWidget(index int, widget *QWidget /*444 QWidget **/, stretch int) int {
	var convArg1 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar12insertWidgetEiP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1, &stretch)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:64
// index:0
// Public
// void addPermanentWidget(class QWidget *, int)
func (this *QStatusBar) AddPermanentWidget(widget *QWidget /*444 QWidget **/, stretch int) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:65
// index:0
// Public
// int insertPermanentWidget(int, class QWidget *, int)
func (this *QStatusBar) InsertPermanentWidget(index int, widget *QWidget /*444 QWidget **/, stretch int) int {
	var convArg1 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1, &stretch)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:66
// index:0
// Public
// void removeWidget(class QWidget *)
func (this *QStatusBar) RemoveWidget(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar12removeWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:68
// index:0
// Public
// void setSizeGripEnabled(_Bool)
func (this *QStatusBar) SetSizeGripEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar18setSizeGripEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:69
// index:0
// Public
// bool isSizeGripEnabled()
func (this *QStatusBar) IsSizeGripEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStatusBar17isSizeGripEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstatusbar.h:71
// index:0
// Public
// QString currentMessage()
func (this *QStatusBar) CurrentMessage() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStatusBar14currentMessageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstatusbar.h:74
// index:0
// Public
// void showMessage(const class QString &, int)
func (this *QStatusBar) ShowMessage(text *qtcore.QString, timeout int) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar11showMessageERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &timeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:75
// index:0
// Public
// void clearMessage()
func (this *QStatusBar) ClearMessage() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar12clearMessageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:79
// index:0
// Public
// void messageChanged(const class QString &)
func (this *QStatusBar) MessageChanged(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar14messageChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:82
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QStatusBar) ShowEvent(arg0 *qtgui.QShowEvent /*444 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:83
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QStatusBar) PaintEvent(arg0 *qtgui.QPaintEvent /*444 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:84
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QStatusBar) ResizeEvent(arg0 *qtgui.QResizeEvent /*444 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:87
// index:0
// Protected
// void reformat()
func (this *QStatusBar) Reformat() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar8reformatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:88
// index:0
// Protected
// void hideOrShow()
func (this *QStatusBar) HideOrShow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar10hideOrShowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:89
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QStatusBar) Event(arg0 *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
