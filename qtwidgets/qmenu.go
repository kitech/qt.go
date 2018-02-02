package qtwidgets

// /usr/include/qt/QtWidgets/qmenu.h
// #include <qmenu.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
// int columnCount()
func (this *QMenu) InheritColumnCount(f func() int) {
	ffiqt.SetAllInheritCallback(this, "columnCount", f)
}

// void changeEvent(class QEvent *)
func (this *QMenu) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QMenu) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QMenu) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QMenu) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QMenu) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QMenu) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void enterEvent(class QEvent *)
func (this *QMenu) InheritEnterEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "enterEvent", f)
}

// void leaveEvent(class QEvent *)
func (this *QMenu) InheritLeaveEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QMenu) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "hideEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QMenu) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void actionEvent(class QActionEvent *)
func (this *QMenu) InheritActionEvent(f func(arg0 *qtgui.QActionEvent /*777 QActionEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "actionEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QMenu) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "timerEvent", f)
}

// bool event(class QEvent *)
func (this *QMenu) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QMenu) InheritFocusNextPrevChild(f func(next bool) bool) {
	ffiqt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void initStyleOption(class QStyleOptionMenuItem *, const class QAction *)
func (this *QMenu) InheritInitStyleOption(f func(option *QStyleOptionMenuItem /*777 QStyleOptionMenuItem **/, action *QAction /*777 const QAction **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QMenu struct {
	*QWidget
}

func (this *QMenu) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QMenu) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQMenuFromPointer(cthis unsafe.Pointer) *QMenu {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMenu{bcthis0}
}
func (*QMenu) NewFromPointer(cthis unsafe.Pointer) *QMenu {
	return NewQMenuFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmenu.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QMenu) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenu(QWidget *)
func NewQMenu(parent *QWidget /*777 QWidget **/) *QMenu {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMenu(const QString &, QWidget *)
func NewQMenu_1(title *qtcore.QString, parent *QWidget /*777 QWidget **/) *QMenu {
	var convArg0 = title.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMenu()
func DeleteQMenu(this *QMenu) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmenu.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)
func (this *QMenu) AddAction(text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:80
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &)
func (this *QMenu) AddAction_1(icon *qtgui.QIcon, text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:81
// index:2
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &, const QObject *, const char *, const QKeySequence &)
func (this *QMenu) AddAction_2(text *qtcore.QString, receiver *qtcore.QObject /*777 const QObject **/, member string, shortcut *qtgui.QKeySequence) *QAction /*777 QAction **/ {
	var convArg0 = text.GetCthis()
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = shortcut.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:82
// index:3
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &, const QObject *, const char *, const QKeySequence &)
func (this *QMenu) AddAction_3(icon *qtgui.QIcon, text *qtcore.QString, receiver *qtcore.QObject /*777 const QObject **/, member string, shortcut *qtgui.QKeySequence) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = receiver.GetCthis()
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	var convArg4 = shortcut.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addMenu(QMenu *)
func (this *QMenu) AddMenu(menu *QMenu /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 = menu.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:157
// index:1
// Public Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QString &)
func (this *QMenu) AddMenu_1(title *qtcore.QString) *QMenu /*777 QMenu **/ {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:158
// index:2
// Public Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QIcon &, const QString &)
func (this *QMenu) AddMenu_2(icon *qtgui.QIcon, title *qtcore.QString) *QMenu /*777 QMenu **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addSeparator()
func (this *QMenu) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu12addSeparatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addSection(const QString &)
func (this *QMenu) AddSection(text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10addSectionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:163
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addSection(const QIcon &, const QString &)
func (this *QMenu) AddSection_1(icon *qtgui.QIcon, text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10addSectionERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertMenu(QAction *, QMenu *)
func (this *QMenu) InsertMenu(before *QAction /*777 QAction **/, menu *QMenu /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	var convArg1 = menu.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10insertMenuEP7QActionPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertSeparator(QAction *)
func (this *QMenu) InsertSeparator(before *QAction /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertSection(QAction *, const QString &)
func (this *QMenu) InsertSection(before *QAction /*777 QAction **/, text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:168
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * insertSection(QAction *, const QIcon &, const QString &)
func (this *QMenu) InsertSection_1(before *QAction /*777 QAction **/, icon *qtgui.QIcon, text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	var convArg1 = icon.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QMenu) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QMenu) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTearOffEnabled(_Bool)
func (this *QMenu) SetTearOffEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu17setTearOffEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:174
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTearOffEnabled()
func (this *QMenu) IsTearOffEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu16isTearOffEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTearOffMenuVisible()
func (this *QMenu) IsTearOffMenuVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu20isTearOffMenuVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showTearOffMenu()
func (this *QMenu) ShowTearOffMenu() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:178
// index:1
// Public Visibility=Default Availability=Available
// [-2] void showTearOffMenu(const QPoint &)
func (this *QMenu) ShowTearOffMenu_1(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hideTearOffMenu()
func (this *QMenu) HideTearOffMenu() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15hideTearOffMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultAction(QAction *)
func (this *QMenu) SetDefaultAction(arg0 *QAction /*777 QAction **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu16setDefaultActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * defaultAction()
func (this *QMenu) DefaultAction() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu13defaultActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveAction(QAction *)
func (this *QMenu) SetActiveAction(act *QAction /*777 QAction **/) {
	var convArg0 = act.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15setActiveActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * activeAction()
func (this *QMenu) ActiveAction() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu12activeActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void popup(const QPoint &, QAction *)
func (this *QMenu) Popup(pos *qtcore.QPoint, at *QAction /*777 QAction **/) {
	var convArg0 = pos.GetCthis()
	var convArg1 = at.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5popupERK6QPointP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * exec()
func (this *QMenu) Exec() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu4execEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:189
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * exec(const QPoint &, QAction *)
func (this *QMenu) Exec_1(pos *qtcore.QPoint, at *QAction /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 = pos.GetCthis()
	var convArg1 = at.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu4execERK6QPointP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:197
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QMenu) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:199
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect actionGeometry(QAction *)
func (this *QMenu) ActionGeometry(arg0 *QAction /*777 QAction **/) *qtcore.QRect /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu14actionGeometryEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * actionAt(const QPoint &)
func (this *QMenu) ActionAt(arg0 *qtcore.QPoint) *QAction /*777 QAction **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu8actionAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:202
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * menuAction()
func (this *QMenu) MenuAction() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu10menuActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:204
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title()
func (this *QMenu) Title() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu5titleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)
func (this *QMenu) SetTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:207
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QMenu) Icon() *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QMenu) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:210
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNoReplayFor(QWidget *)
func (this *QMenu) SetNoReplayFor(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu14setNoReplayForEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:219
// index:0
// Public Visibility=Default Availability=Available
// [1] bool separatorsCollapsible()
func (this *QMenu) SeparatorsCollapsible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu21separatorsCollapsibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSeparatorsCollapsible(_Bool)
func (this *QMenu) SetSeparatorsCollapsible(collapse bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu24setSeparatorsCollapsibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), collapse)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:222
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toolTipsVisible()
func (this *QMenu) ToolTipsVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu15toolTipsVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTipsVisible(_Bool)
func (this *QMenu) SetToolTipsVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu18setToolTipsVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToShow()
func (this *QMenu) AboutToShow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11aboutToShowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToHide()
func (this *QMenu) AboutToHide() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11aboutToHideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(QAction *)
func (this *QMenu) Triggered(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9triggeredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hovered(QAction *)
func (this *QMenu) Hovered(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7hoveredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:232
// index:0
// Protected Visibility=Default Availability=Available
// [4] int columnCount()
func (this *QMenu) ColumnCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmenu.h:234
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QMenu) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:235
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QMenu) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:236
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QMenu) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:237
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QMenu) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:238
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QMenu) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:240
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QMenu) WheelEvent(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:242
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void enterEvent(QEvent *)
func (this *QMenu) EnterEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10enterEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:243
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)
func (this *QMenu) LeaveEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10leaveEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QMenu) HideEvent(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:245
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QMenu) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:246
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)
func (this *QMenu) ActionEvent(arg0 *qtgui.QActionEvent /*777 QActionEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:247
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QMenu) TimerEvent(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:248
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QMenu) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:249
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QMenu) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:250
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionMenuItem *, const QAction *)
func (this *QMenu) InitStyleOption(option *QStyleOptionMenuItem /*777 QStyleOptionMenuItem **/, action *QAction /*777 const QAction **/) {
	var convArg0 = option.GetCthis()
	var convArg1 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
