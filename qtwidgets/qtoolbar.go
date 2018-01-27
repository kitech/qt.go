package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbar.h
// #include <qtoolbar.h>
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
type QToolBar struct {
	*QWidget
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QToolBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:79
// index:0
// Public
// void QToolBar(const QString &, QWidget *)
func NewQToolBar(title *qtcore.QString, parent *QWidget /*777 QWidget **/) *QToolBar {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = title.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBarC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:80
// index:1
// Public
// void QToolBar(QWidget *)
func NewQToolBar_1(parent *QWidget /*777 QWidget **/) *QToolBar {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:81
// index:0
// Public virtual
// void ~QToolBar()
func DeleteQToolBar(*QToolBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:83
// index:0
// Public
// void setMovable(bool)
func (this *QToolBar) SetMovable(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar10setMovableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:84
// index:0
// Public
// bool isMovable()
func (this *QToolBar) IsMovable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar9isMovableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:87
// index:0
// Public
// Qt::ToolBarAreas allowedAreas()
func (this *QToolBar) AllowedAreas() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar12allowedAreasEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:89
// index:0
// Public inline
// bool isAreaAllowed(Qt::ToolBarArea)
func (this *QToolBar) IsAreaAllowed(area int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar13isAreaAllowedEN2Qt11ToolBarAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:92
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QToolBar) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:93
// index:0
// Public
// Qt::Orientation orientation()
func (this *QToolBar) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:95
// index:0
// Public
// void clear()
func (this *QToolBar) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:98
// index:0
// Public
// QAction * addAction(const QString &)
func (this *QToolBar) AddAction(text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:99
// index:1
// Public
// QAction * addAction(const QIcon &, const QString &)
func (this *QToolBar) AddAction_1(icon *qtgui.QIcon, text *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:100
// index:2
// Public
// QAction * addAction(const QString &, const QObject *, const char *)
func (this *QToolBar) AddAction_2(text *qtcore.QString, receiver *qtcore.QObject /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var convArg0 = text.GetCthis()
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QStringPK7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:101
// index:3
// Public
// QAction * addAction(const QIcon &, const QString &, const QObject *, const char *)
func (this *QToolBar) AddAction_3(icon *qtgui.QIcon, text *qtcore.QString, receiver *qtcore.QObject /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = receiver.GetCthis()
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:155
// index:0
// Public
// QAction * addSeparator()
func (this *QToolBar) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar12addSeparatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:156
// index:0
// Public
// QAction * insertSeparator(QAction *)
func (this *QToolBar) InsertSeparator(before *QAction /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:158
// index:0
// Public
// QAction * addWidget(QWidget *)
func (this *QToolBar) AddWidget(widget *QWidget /*777 QWidget **/) *QAction /*777 QAction **/ {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:159
// index:0
// Public
// QAction * insertWidget(QAction *, QWidget *)
func (this *QToolBar) InsertWidget(before *QAction /*777 QAction **/, widget *QWidget /*777 QWidget **/) *QAction /*777 QAction **/ {
	var convArg0 = before.GetCthis()
	var convArg1 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar12insertWidgetEP7QActionP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:161
// index:0
// Public
// QRect actionGeometry(QAction *)
func (this *QToolBar) ActionGeometry(action *QAction /*777 QAction **/) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar14actionGeometryEP7QAction", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:162
// index:0
// Public
// QAction * actionAt(const QPoint &)
func (this *QToolBar) ActionAt(p *qtcore.QPoint) *QAction /*777 QAction **/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar8actionAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:163
// index:1
// Public inline
// QAction * actionAt(int, int)
func (this *QToolBar) ActionAt_1(x int, y int) *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar8actionAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:165
// index:0
// Public
// QAction * toggleViewAction()
func (this *QToolBar) ToggleViewAction() *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar16toggleViewActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:167
// index:0
// Public
// QSize iconSize()
func (this *QToolBar) IconSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar8iconSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:168
// index:0
// Public
// Qt::ToolButtonStyle toolButtonStyle()
func (this *QToolBar) ToolButtonStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar15toolButtonStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:170
// index:0
// Public
// QWidget * widgetForAction(QAction *)
func (this *QToolBar) WidgetForAction(action *QAction /*777 QAction **/) *QWidget /*777 QWidget **/ {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar15widgetForActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:172
// index:0
// Public
// bool isFloatable()
func (this *QToolBar) IsFloatable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar11isFloatableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:173
// index:0
// Public
// void setFloatable(bool)
func (this *QToolBar) SetFloatable(floatable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar12setFloatableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), floatable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:174
// index:0
// Public
// bool isFloating()
func (this *QToolBar) IsFloating() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar10isFloatingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:177
// index:0
// Public
// void setIconSize(const QSize &)
func (this *QToolBar) SetIconSize(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:178
// index:0
// Public
// void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QToolBar) SetToolButtonStyle(toolButtonStyle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar18setToolButtonStyleEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:181
// index:0
// Public
// void actionTriggered(QAction *)
func (this *QToolBar) ActionTriggered(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15actionTriggeredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:182
// index:0
// Public
// void movableChanged(bool)
func (this *QToolBar) MovableChanged(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar14movableChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:184
// index:0
// Public
// void orientationChanged(Qt::Orientation)
func (this *QToolBar) OrientationChanged(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar18orientationChangedEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:185
// index:0
// Public
// void iconSizeChanged(const QSize &)
func (this *QToolBar) IconSizeChanged(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15iconSizeChangedERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:186
// index:0
// Public
// void toolButtonStyleChanged(Qt::ToolButtonStyle)
func (this *QToolBar) ToolButtonStyleChanged(toolButtonStyle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:187
// index:0
// Public
// void topLevelChanged(bool)
func (this *QToolBar) TopLevelChanged(topLevel bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15topLevelChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), topLevel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:188
// index:0
// Public
// void visibilityChanged(bool)
func (this *QToolBar) VisibilityChanged(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar17visibilityChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:191
// index:0
// Protected virtual
// void actionEvent(QActionEvent *)
func (this *QToolBar) ActionEvent(event *qtgui.QActionEvent /*777 QActionEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:192
// index:0
// Protected virtual
// void changeEvent(QEvent *)
func (this *QToolBar) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:193
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QToolBar) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:194
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QToolBar) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:195
// index:0
// Protected
// void initStyleOption(QStyleOptionToolBar *)
func (this *QToolBar) InitStyleOption(option *QStyleOptionToolBar /*777 QStyleOptionToolBar **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
