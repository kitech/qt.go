//  header block begin
// /usr/include/qt/QtWidgets/qmenubar.h
// #include <qmenubar.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 69
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
type QMenuBar struct {
	*QWidget
}

func (this *QMenuBar) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qmenubar.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMenuBar) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:63
// index:0
// void QMenuBar(class QWidget *)
func NewQMenuBar(parent unsafe.Pointer) *QMenuBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMenuBarFromPointer(cthis)
	return gothis
}
func NewQMenuBarFromPointer(cthis unsafe.Pointer) *QMenuBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMenuBar{bcthis0}
}

// /usr/include/qt/QtWidgets/qmenubar.h:64
// index:0
// virtual
// void ~QMenuBar()
func DeleteQMenuBar(*QMenuBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:67
// index:0
// QAction * addAction(const class QString &)
func (this *QMenuBar) AddAction(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:68
// index:1
// QAction * addAction(const class QString &, const class QObject *, const char *)
func (this *QMenuBar) AddAction_1(text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 1: (, text const QString &, receiver const QObject *, member const char *), (text, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:70
// index:0
// QAction * addMenu(class QMenu *)
func (this *QMenuBar) AddMenu(menu unsafe.Pointer) {
	// 0: (, menu QMenu *), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:71
// index:1
// QMenu * addMenu(const class QString &)
func (this *QMenuBar) AddMenu_1(title unsafe.Pointer) {
	// 1: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:72
// index:2
// QMenu * addMenu(const class QIcon &, const class QString &)
func (this *QMenuBar) AddMenu_2(icon unsafe.Pointer, title unsafe.Pointer) {
	// 2: (, icon const QIcon &, title const QString &), (icon, title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:75
// index:0
// QAction * addSeparator()
func (this *QMenuBar) AddSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12addSeparatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:76
// index:0
// QAction * insertSeparator(class QAction *)
func (this *QMenuBar) InsertSeparator(before unsafe.Pointer) {
	// 0: (, before QAction *), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:78
// index:0
// QAction * insertMenu(class QAction *, class QMenu *)
func (this *QMenuBar) InsertMenu(before unsafe.Pointer, menu unsafe.Pointer) {
	// 0: (, before QAction *, menu QMenu *), (before, menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10insertMenuEP7QActionP5QMenu", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:80
// index:0
// void clear()
func (this *QMenuBar) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:82
// index:0
// QAction * activeAction()
func (this *QMenuBar) ActiveAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar12activeActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:83
// index:0
// void setActiveAction(class QAction *)
func (this *QMenuBar) SetActiveAction(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15setActiveActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:85
// index:0
// void setDefaultUp(_Bool)
func (this *QMenuBar) SetDefaultUp(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12setDefaultUpEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:86
// index:0
// bool isDefaultUp()
func (this *QMenuBar) IsDefaultUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar11isDefaultUpEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:88
// index:0
// virtual
// QSize sizeHint()
func (this *QMenuBar) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:89
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QMenuBar) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:90
// index:0
// virtual
// int heightForWidth(int)
func (this *QMenuBar) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:92
// index:0
// QRect actionGeometry(class QAction *)
func (this *QMenuBar) ActionGeometry(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar14actionGeometryEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:93
// index:0
// QAction * actionAt(const class QPoint &)
func (this *QMenuBar) ActionAt(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar8actionAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:95
// index:0
// void setCornerWidget(class QWidget *, Qt::Corner)
func (this *QMenuBar) SetCornerWidget(w unsafe.Pointer, corner int) {
	// 0: (, w QWidget *, corner Qt::Corner), (w, &corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15setCornerWidgetEP7QWidgetN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w, &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:96
// index:0
// QWidget * cornerWidget(Qt::Corner)
func (this *QMenuBar) CornerWidget(corner int) {
	// 0: (, corner Qt::Corner), (&corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar12cornerWidgetEN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:102
// index:0
// bool isNativeMenuBar()
func (this *QMenuBar) IsNativeMenuBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar15isNativeMenuBarEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:103
// index:0
// void setNativeMenuBar(_Bool)
func (this *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	// 0: (, nativeMenuBar bool), (&nativeMenuBar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar16setNativeMenuBarEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nativeMenuBar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:104
// index:0
// QPlatformMenuBar * platformMenuBar()
func (this *QMenuBar) PlatformMenuBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15platformMenuBarEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:106
// index:0
// virtual
// void setVisible(_Bool)
func (this *QMenuBar) SetVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:109
// index:0
// void triggered(class QAction *)
func (this *QMenuBar) Triggered(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9triggeredEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:110
// index:0
// void hovered(class QAction *)
func (this *QMenuBar) Hovered(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7hoveredEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:113
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QMenuBar) ChangeEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:114
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QMenuBar) KeyPressEvent(arg0 unsafe.Pointer) {
	// 0: (, QKeyEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:115
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QMenuBar) MouseReleaseEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:116
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QMenuBar) MousePressEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:117
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QMenuBar) MouseMoveEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:118
// index:0
// virtual
// void leaveEvent(class QEvent *)
func (this *QMenuBar) LeaveEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10leaveEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:119
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QMenuBar) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:120
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QMenuBar) ResizeEvent(arg0 unsafe.Pointer) {
	// 0: (, QResizeEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:121
// index:0
// virtual
// void actionEvent(class QActionEvent *)
func (this *QMenuBar) ActionEvent(arg0 unsafe.Pointer) {
	// 0: (, QActionEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:122
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QMenuBar) FocusOutEvent(arg0 unsafe.Pointer) {
	// 0: (, QFocusEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:123
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QMenuBar) FocusInEvent(arg0 unsafe.Pointer) {
	// 0: (, QFocusEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:124
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QMenuBar) TimerEvent(arg0 unsafe.Pointer) {
	// 0: (, QTimerEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:125
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QMenuBar) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QObject * arg0, QEvent * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:126
// index:0
// virtual
// bool event(class QEvent *)
func (this *QMenuBar) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:127
// index:0
// void initStyleOption(class QStyleOptionMenuItem *, const class QAction *)
func (this *QMenuBar) InitStyleOption(option unsafe.Pointer, action unsafe.Pointer) {
	// 0: (, option QStyleOptionMenuItem *, action const QAction *), (option, action)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, action)
	gopp.ErrPrint(err, rv)
}

//  body block end
