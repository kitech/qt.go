package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbar.h
// #include <qtoolbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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

// void actionEvent(class QActionEvent *)
func (this *QToolBar) InheritActionEvent(f func(event *qtgui.QActionEvent /*777 QActionEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "actionEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QToolBar) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QToolBar) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// bool event(class QEvent *)
func (this *QToolBar) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void initStyleOption(class QStyleOptionToolBar *)
func (this *QToolBar) InheritInitStyleOption(f func(option *QStyleOptionToolBar /*777 QStyleOptionToolBar **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QToolBar struct {
	*QWidget
}
type QToolBar_ITF interface {
	QWidget_ITF
	QToolBar_PTR() *QToolBar
}

func (ptr *QToolBar) QToolBar_PTR() *QToolBar { return ptr }

func (this *QToolBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QToolBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQToolBarFromPointer(cthis unsafe.Pointer) *QToolBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QToolBar{bcthis0}
}
func (*QToolBar) NewFromPointer(cthis unsafe.Pointer) *QToolBar {
	return NewQToolBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QToolBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(const QString &, QWidget *)
func NewQToolBar(title string, parent *QWidget /*777 QWidget **/) *QToolBar {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:80
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(QWidget *)
func NewQToolBar_1(parent *QWidget /*777 QWidget **/) *QToolBar {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QToolBar()
func DeleteQToolBar(this *QToolBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovable(_Bool)
func (this *QToolBar) SetMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar10setMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMovable()
func (this *QToolBar) IsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar9isMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAllowedAreas(Qt::ToolBarAreas)
func (this *QToolBar) SetAllowedAreas(areas int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15setAllowedAreasE6QFlagsIN2Qt11ToolBarAreaEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), areas)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolBarAreas allowedAreas()
func (this *QToolBar) AllowedAreas() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar12allowedAreasEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAreaAllowed(Qt::ToolBarArea)
func (this *QToolBar) IsAreaAllowed(area int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar13isAreaAllowedEN2Qt11ToolBarAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QToolBar) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QToolBar) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QToolBar) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)
func (this *QToolBar) AddAction(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:99
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &)
func (this *QToolBar) AddAction_1(icon *qtgui.QIcon, text string) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:100
// index:2
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &, const QObject *, const char *)
func (this *QToolBar) AddAction_2(text string, receiver *qtcore.QObject /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QStringPK7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:101
// index:3
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &, const QObject *, const char *)
func (this *QToolBar) AddAction_3(icon *qtgui.QIcon, text string, receiver *qtcore.QObject /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = receiver.GetCthis()
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addSeparator()
func (this *QToolBar) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar12addSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertSeparator(QAction *)
func (this *QToolBar) InsertSeparator(before *QAction /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15insertSeparatorEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addWidget(QWidget *)
func (this *QToolBar) AddWidget(widget *QWidget /*777 QWidget **/) *QAction /*777 QAction **/ {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertWidget(QAction *, QWidget *)
func (this *QToolBar) InsertWidget(before *QAction /*777 QAction **/, widget *QWidget /*777 QWidget **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	var convArg1 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar12insertWidgetEP7QActionP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:161
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect actionGeometry(QAction *)
func (this *QToolBar) ActionGeometry(action *QAction /*777 QAction **/) *qtcore.QRect /*123*/ {
	var convArg0 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar14actionGeometryEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * actionAt(const QPoint &)
func (this *QToolBar) ActionAt(p *qtcore.QPoint) *QAction /*777 QAction **/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar8actionAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:163
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QAction * actionAt(int, int)
func (this *QToolBar) ActionAt_1(x int, y int) *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar8actionAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * toggleViewAction()
func (this *QToolBar) ToggleViewAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar16toggleViewActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
func (this *QToolBar) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:168
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolButtonStyle toolButtonStyle()
func (this *QToolBar) ToolButtonStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar15toolButtonStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widgetForAction(QAction *)
func (this *QToolBar) WidgetForAction(action *QAction /*777 QAction **/) *QWidget /*777 QWidget **/ {
	var convArg0 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar15widgetForActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:172
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFloatable()
func (this *QToolBar) IsFloatable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar11isFloatableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFloatable(_Bool)
func (this *QToolBar) SetFloatable(floatable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar12setFloatableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), floatable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:174
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFloating()
func (this *QToolBar) IsFloating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar10isFloatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)
func (this *QToolBar) SetIconSize(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QToolBar) SetToolButtonStyle(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar18setToolButtonStyleEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actionTriggered(QAction *)
func (this *QToolBar) ActionTriggered(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15actionTriggeredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void movableChanged(_Bool)
func (this *QToolBar) MovableChanged(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar14movableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void allowedAreasChanged(Qt::ToolBarAreas)
func (this *QToolBar) AllowedAreasChanged(allowedAreas int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar19allowedAreasChangedE6QFlagsIN2Qt11ToolBarAreaEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), allowedAreas)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void orientationChanged(Qt::Orientation)
func (this *QToolBar) OrientationChanged(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar18orientationChangedEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconSizeChanged(const QSize &)
func (this *QToolBar) IconSizeChanged(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15iconSizeChangedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toolButtonStyleChanged(Qt::ToolButtonStyle)
func (this *QToolBar) ToolButtonStyleChanged(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void topLevelChanged(_Bool)
func (this *QToolBar) TopLevelChanged(topLevel bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15topLevelChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), topLevel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibilityChanged(_Bool)
func (this *QToolBar) VisibilityChanged(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar17visibilityChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)
func (this *QToolBar) ActionEvent(event *qtgui.QActionEvent /*777 QActionEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar11actionEventEP12QActionEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QToolBar) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:193
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QToolBar) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QToolBar) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:195
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionToolBar *)
func (this *QToolBar) InitStyleOption(option *QStyleOptionToolBar /*777 QStyleOptionToolBar **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
