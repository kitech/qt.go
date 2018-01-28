package qtwidgets

// /usr/include/qt/QtWidgets/qmdisubwindow.h
// #include <qmdisubwindow.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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
type QMdiSubWindow struct {
	*QWidget
}

func (this *QMdiSubWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QMdiSubWindow) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQMdiSubWindowFromPointer(cthis unsafe.Pointer) *QMdiSubWindow {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMdiSubWindow{bcthis0}
}
func (*QMdiSubWindow) NewFromPointer(cthis unsafe.Pointer) *QMdiSubWindow {
	return NewQMdiSubWindowFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMdiSubWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:69
// index:0
// Public
// void QMdiSubWindow(QWidget *, Qt::WindowFlags)
func NewQMdiSubWindow(parent *QWidget /*777 QWidget **/, flags int) *QMdiSubWindow {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMdiSubWindowFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:70
// index:0
// Public virtual
// void ~QMdiSubWindow()
func DeleteQMdiSubWindow(*QMdiSubWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:72
// index:0
// Public virtual
// QSize sizeHint()
func (this *QMdiSubWindow) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK13QMdiSubWindow8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:73
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QMdiSubWindow) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK13QMdiSubWindow15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:75
// index:0
// Public
// void setWidget(QWidget *)
func (this *QMdiSubWindow) SetWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:76
// index:0
// Public
// QWidget * widget()
func (this *QMdiSubWindow) Widget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:78
// index:0
// Public
// QWidget * maximizedButtonsWidget()
func (this *QMdiSubWindow) MaximizedButtonsWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:79
// index:0
// Public
// QWidget * maximizedSystemMenuIconWidget()
func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:81
// index:0
// Public
// bool isShaded()
func (this *QMdiSubWindow) IsShaded() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow8isShadedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:83
// index:0
// Public
// void setOption(enum QMdiSubWindow::SubWindowOption, _Bool)
func (this *QMdiSubWindow) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9setOptionENS_15SubWindowOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:84
// index:0
// Public
// bool testOption(enum QMdiSubWindow::SubWindowOption)
func (this *QMdiSubWindow) TestOption(arg0 int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10testOptionENS_15SubWindowOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:86
// index:0
// Public
// void setKeyboardSingleStep(int)
func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow21setKeyboardSingleStepEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:87
// index:0
// Public
// int keyboardSingleStep()
func (this *QMdiSubWindow) KeyboardSingleStep() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow18keyboardSingleStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:89
// index:0
// Public
// void setKeyboardPageStep(int)
func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow19setKeyboardPageStepEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:90
// index:0
// Public
// int keyboardPageStep()
func (this *QMdiSubWindow) KeyboardPageStep() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow16keyboardPageStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:93
// index:0
// Public
// void setSystemMenu(QMenu *)
func (this *QMdiSubWindow) SetSystemMenu(systemMenu *QMenu /*777 QMenu **/) {
	var convArg0 = systemMenu.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow13setSystemMenuEP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:94
// index:0
// Public
// QMenu * systemMenu()
func (this *QMdiSubWindow) SystemMenu() *QMenu /*777 QMenu **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10systemMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:97
// index:0
// Public
// QMdiArea * mdiArea()
func (this *QMdiSubWindow) MdiArea() *QMdiArea /*777 QMdiArea **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow7mdiAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMdiAreaFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:101
// index:0
// Public
// void aboutToActivate()
func (this *QMdiSubWindow) AboutToActivate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow15aboutToActivateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:105
// index:0
// Public
// void showSystemMenu()
func (this *QMdiSubWindow) ShowSystemMenu() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow14showSystemMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:107
// index:0
// Public
// void showShaded()
func (this *QMdiSubWindow) ShowShaded() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10showShadedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:110
// index:0
// Protected virtual
// bool eventFilter(QObject *, QEvent *)
func (this *QMdiSubWindow) EventFilter(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = object.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:111
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QMdiSubWindow) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:112
// index:0
// Protected virtual
// void showEvent(QShowEvent *)
func (this *QMdiSubWindow) ShowEvent(showEvent *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = showEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:113
// index:0
// Protected virtual
// void hideEvent(QHideEvent *)
func (this *QMdiSubWindow) HideEvent(hideEvent *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = hideEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:114
// index:0
// Protected virtual
// void changeEvent(QEvent *)
func (this *QMdiSubWindow) ChangeEvent(changeEvent *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = changeEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:115
// index:0
// Protected virtual
// void closeEvent(QCloseEvent *)
func (this *QMdiSubWindow) CloseEvent(closeEvent *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = closeEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:116
// index:0
// Protected virtual
// void leaveEvent(QEvent *)
func (this *QMdiSubWindow) LeaveEvent(leaveEvent *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = leaveEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10leaveEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:117
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QMdiSubWindow) ResizeEvent(resizeEvent *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = resizeEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:118
// index:0
// Protected virtual
// void timerEvent(QTimerEvent *)
func (this *QMdiSubWindow) TimerEvent(timerEvent *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = timerEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:119
// index:0
// Protected virtual
// void moveEvent(QMoveEvent *)
func (this *QMdiSubWindow) MoveEvent(moveEvent *qtgui.QMoveEvent /*777 QMoveEvent **/) {
	var convArg0 = moveEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9moveEventEP10QMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:120
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QMdiSubWindow) PaintEvent(paintEvent *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = paintEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:121
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QMdiSubWindow) MousePressEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:122
// index:0
// Protected virtual
// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QMdiSubWindow) MouseDoubleClickEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:123
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QMdiSubWindow) MouseReleaseEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:124
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QMdiSubWindow) MouseMoveEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:125
// index:0
// Protected virtual
// void keyPressEvent(QKeyEvent *)
func (this *QMdiSubWindow) KeyPressEvent(keyEvent *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = keyEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:127
// index:0
// Protected virtual
// void contextMenuEvent(QContextMenuEvent *)
func (this *QMdiSubWindow) ContextMenuEvent(contextMenuEvent *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = contextMenuEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:129
// index:0
// Protected virtual
// void focusInEvent(QFocusEvent *)
func (this *QMdiSubWindow) FocusInEvent(focusInEvent *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = focusInEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:130
// index:0
// Protected virtual
// void focusOutEvent(QFocusEvent *)
func (this *QMdiSubWindow) FocusOutEvent(focusOutEvent *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = focusOutEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:131
// index:0
// Protected virtual
// void childEvent(QChildEvent *)
func (this *QMdiSubWindow) ChildEvent(childEvent *qtcore.QChildEvent /*777 QChildEvent **/) {
	var convArg0 = childEvent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QMdiSubWindow__SubWindowOption = int

const QMdiSubWindow__AllowOutsideAreaHorizontally QMdiSubWindow__SubWindowOption = 1
const QMdiSubWindow__AllowOutsideAreaVertically QMdiSubWindow__SubWindowOption = 2
const QMdiSubWindow__RubberBandResize QMdiSubWindow__SubWindowOption = 4
const QMdiSubWindow__RubberBandMove QMdiSubWindow__SubWindowOption = 8

//  body block end
