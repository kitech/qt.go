package qtwidgets

// /usr/include/qt/QtWidgets/qmenubar.h
// #include <qmenubar.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
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
type QMenuBar struct {
	*QWidget
}

func (this *QMenuBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QMenuBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQMenuBarFromPointer(cthis unsafe.Pointer) *QMenuBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMenuBar{bcthis0}
}
func (*QMenuBar) NewFromPointer(cthis unsafe.Pointer) *QMenuBar {
	return NewQMenuBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmenubar.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMenuBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:63
// index:0
// Public
// void QMenuBar(QWidget *)
func NewQMenuBar(parent *QWidget /*777 QWidget **/) *QMenuBar {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMenuBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmenubar.h:64
// index:0
// Public virtual
// void ~QMenuBar()
func DeleteQMenuBar(*QMenuBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:67
// index:0
// Public
// QAction * addAction(const QString &)
func (this *QMenuBar) AddAction(text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:68
// index:1
// Public
// QAction * addAction(const QString &, const QObject *, const char *)
func (this *QMenuBar) AddAction_1(text *qtcore.QString, receiver *qtcore.QObject /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var convArg0 = text.GetCthis()
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:70
// index:0
// Public
// QAction * addMenu(QMenu *)
func (this *QMenuBar) AddMenu(menu *QMenu /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 = menu.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuEP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:71
// index:1
// Public
// QMenu * addMenu(const QString &)
func (this *QMenuBar) AddMenu_1(title *qtcore.QString) *QMenu /*777 QMenu **/ {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:72
// index:2
// Public
// QMenu * addMenu(const QIcon &, const QString &)
func (this *QMenuBar) AddMenu_2(icon *qtgui.QIcon, title *qtcore.QString) *QMenu /*777 QMenu **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:75
// index:0
// Public
// QAction * addSeparator()
func (this *QMenuBar) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12addSeparatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:76
// index:0
// Public
// QAction * insertSeparator(QAction *)
func (this *QMenuBar) InsertSeparator(before *QAction /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:78
// index:0
// Public
// QAction * insertMenu(QAction *, QMenu *)
func (this *QMenuBar) InsertMenu(before *QAction /*777 QAction **/, menu *QMenu /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	var convArg1 = menu.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10insertMenuEP7QActionP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:80
// index:0
// Public
// void clear()
func (this *QMenuBar) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:82
// index:0
// Public
// QAction * activeAction()
func (this *QMenuBar) ActiveAction() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar12activeActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:83
// index:0
// Public
// void setActiveAction(QAction *)
func (this *QMenuBar) SetActiveAction(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15setActiveActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:85
// index:0
// Public
// void setDefaultUp(_Bool)
func (this *QMenuBar) SetDefaultUp(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12setDefaultUpEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:86
// index:0
// Public
// bool isDefaultUp()
func (this *QMenuBar) IsDefaultUp() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar11isDefaultUpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:88
// index:0
// Public virtual
// QSize sizeHint()
func (this *QMenuBar) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK8QMenuBar8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:89
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QMenuBar) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK8QMenuBar15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:90
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QMenuBar) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmenubar.h:92
// index:0
// Public
// QRect actionGeometry(QAction *)
func (this *QMenuBar) ActionGeometry(arg0 *QAction /*777 QAction **/) *qtcore.QRect /*123*/ {
	var convArg0 = arg0.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar14actionGeometryEP7QAction", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:93
// index:0
// Public
// QAction * actionAt(const QPoint &)
func (this *QMenuBar) ActionAt(arg0 *qtcore.QPoint) *QAction /*777 QAction **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar8actionAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:95
// index:0
// Public
// void setCornerWidget(QWidget *, Qt::Corner)
func (this *QMenuBar) SetCornerWidget(w *QWidget /*777 QWidget **/, corner int) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15setCornerWidgetEP7QWidgetN2Qt6CornerE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:96
// index:0
// Public
// QWidget * cornerWidget(Qt::Corner)
func (this *QMenuBar) CornerWidget(corner int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar12cornerWidgetEN2Qt6CornerE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenubar.h:102
// index:0
// Public
// bool isNativeMenuBar()
func (this *QMenuBar) IsNativeMenuBar() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar15isNativeMenuBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:103
// index:0
// Public
// void setNativeMenuBar(_Bool)
func (this *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar16setNativeMenuBarEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), nativeMenuBar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:106
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QMenuBar) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:109
// index:0
// Public
// void triggered(QAction *)
func (this *QMenuBar) Triggered(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9triggeredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:110
// index:0
// Public
// void hovered(QAction *)
func (this *QMenuBar) Hovered(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7hoveredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:113
// index:0
// Protected virtual
// void changeEvent(QEvent *)
func (this *QMenuBar) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:114
// index:0
// Protected virtual
// void keyPressEvent(QKeyEvent *)
func (this *QMenuBar) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:115
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QMenuBar) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:116
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QMenuBar) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:117
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QMenuBar) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:118
// index:0
// Protected virtual
// void leaveEvent(QEvent *)
func (this *QMenuBar) LeaveEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10leaveEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:119
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QMenuBar) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:120
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QMenuBar) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:121
// index:0
// Protected virtual
// void actionEvent(QActionEvent *)
func (this *QMenuBar) ActionEvent(arg0 *qtgui.QActionEvent /*777 QActionEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:122
// index:0
// Protected virtual
// void focusOutEvent(QFocusEvent *)
func (this *QMenuBar) FocusOutEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:123
// index:0
// Protected virtual
// void focusInEvent(QFocusEvent *)
func (this *QMenuBar) FocusInEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:124
// index:0
// Protected virtual
// void timerEvent(QTimerEvent *)
func (this *QMenuBar) TimerEvent(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:125
// index:0
// Protected virtual
// bool eventFilter(QObject *, QEvent *)
func (this *QMenuBar) EventFilter(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:126
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QMenuBar) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenubar.h:127
// index:0
// Protected
// void initStyleOption(QStyleOptionMenuItem *, const QAction *)
func (this *QMenuBar) InitStyleOption(option *QStyleOptionMenuItem /*777 QStyleOptionMenuItem **/, action *QAction /*777 const QAction **/) {
	var convArg0 = option.GetCthis()
	var convArg1 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
