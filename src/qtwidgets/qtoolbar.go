//  header block begin
// /usr/include/qt/QtWidgets/qtoolbar.h
// #include <qtoolbar.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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
type QToolBar struct {
	*QWidget
}

func (this *QToolBar) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qtoolbar.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QToolBar) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:79
// index:0
// void QToolBar(const class QString &, class QWidget *)
func NewQToolBar(title unsafe.Pointer, parent unsafe.Pointer) *QToolBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBarC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, title, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(cthis)
	return gothis
}
func NewQToolBarFromPointer(cthis unsafe.Pointer) *QToolBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QToolBar{bcthis0}
}

// /usr/include/qt/QtWidgets/qtoolbar.h:80
// index:1
// void QToolBar(class QWidget *)
func NewQToolBar_1(parent unsafe.Pointer) *QToolBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:81
// index:0
// virtual
// void ~QToolBar()
func DeleteQToolBar(*QToolBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:83
// index:0
// void setMovable(_Bool)
func (this *QToolBar) SetMovable(movable bool) {
	// 0: (, movable bool), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar10setMovableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:84
// index:0
// bool isMovable()
func (this *QToolBar) IsMovable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar9isMovableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:86
// index:0
// void setAllowedAreas(Qt::ToolBarAreas)
func (this *QToolBar) SetAllowedAreas(areas int) {
	// 0: (, QFlags<Qt::ToolBarArea> areas), (&areas)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15setAllowedAreasE6QFlagsIN2Qt11ToolBarAreaEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &areas)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:87
// index:0
// Qt::ToolBarAreas allowedAreas()
func (this *QToolBar) AllowedAreas() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar12allowedAreasEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:89
// index:0
// inline
// bool isAreaAllowed(Qt::ToolBarArea)
func (this *QToolBar) IsAreaAllowed(area int) {
	// 0: (, area Qt::ToolBarArea), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar13isAreaAllowedEN2Qt11ToolBarAreaE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:92
// index:0
// void setOrientation(Qt::Orientation)
func (this *QToolBar) SetOrientation(orientation int) {
	// 0: (, orientation Qt::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:93
// index:0
// Qt::Orientation orientation()
func (this *QToolBar) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:95
// index:0
// void clear()
func (this *QToolBar) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:98
// index:0
// QAction * addAction(const class QString &)
func (this *QToolBar) AddAction(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:99
// index:1
// QAction * addAction(const class QIcon &, const class QString &)
func (this *QToolBar) AddAction_1(icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, icon const QIcon &, text const QString &), (icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:100
// index:2
// QAction * addAction(const class QString &, const class QObject *, const char *)
func (this *QToolBar) AddAction_2(text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 2: (, text const QString &, receiver const QObject *, member const char *), (text, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QStringPK7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:101
// index:3
// QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *)
func (this *QToolBar) AddAction_3(icon unsafe.Pointer, text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 3: (, icon const QIcon &, text const QString &, receiver const QObject *, member const char *), (icon, text, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, text, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:155
// index:0
// QAction * addSeparator()
func (this *QToolBar) AddSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar12addSeparatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:156
// index:0
// QAction * insertSeparator(class QAction *)
func (this *QToolBar) InsertSeparator(before unsafe.Pointer) {
	// 0: (, before QAction *), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:158
// index:0
// QAction * addWidget(class QWidget *)
func (this *QToolBar) AddWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar9addWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:159
// index:0
// QAction * insertWidget(class QAction *, class QWidget *)
func (this *QToolBar) InsertWidget(before unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, before QAction *, widget QWidget *), (before, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar12insertWidgetEP7QActionP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:161
// index:0
// QRect actionGeometry(class QAction *)
func (this *QToolBar) ActionGeometry(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar14actionGeometryEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:162
// index:0
// QAction * actionAt(const class QPoint &)
func (this *QToolBar) ActionAt(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar8actionAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:163
// index:1
// inline
// QAction * actionAt(int, int)
func (this *QToolBar) ActionAt_1(x int, y int) {
	// 1: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar8actionAtEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:165
// index:0
// QAction * toggleViewAction()
func (this *QToolBar) ToggleViewAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar16toggleViewActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:167
// index:0
// QSize iconSize()
func (this *QToolBar) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:168
// index:0
// Qt::ToolButtonStyle toolButtonStyle()
func (this *QToolBar) ToolButtonStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar15toolButtonStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:170
// index:0
// QWidget * widgetForAction(class QAction *)
func (this *QToolBar) WidgetForAction(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar15widgetForActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:172
// index:0
// bool isFloatable()
func (this *QToolBar) IsFloatable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar11isFloatableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:173
// index:0
// void setFloatable(_Bool)
func (this *QToolBar) SetFloatable(floatable bool) {
	// 0: (, floatable bool), (&floatable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar12setFloatableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &floatable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:174
// index:0
// bool isFloating()
func (this *QToolBar) IsFloating() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar10isFloatingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:177
// index:0
// void setIconSize(const class QSize &)
func (this *QToolBar) SetIconSize(iconSize unsafe.Pointer) {
	// 0: (, iconSize const QSize &), (iconSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), iconSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:178
// index:0
// void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QToolBar) SetToolButtonStyle(toolButtonStyle int) {
	// 0: (, toolButtonStyle Qt::ToolButtonStyle), (&toolButtonStyle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar18setToolButtonStyleEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:181
// index:0
// void actionTriggered(class QAction *)
func (this *QToolBar) ActionTriggered(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15actionTriggeredEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:182
// index:0
// void movableChanged(_Bool)
func (this *QToolBar) MovableChanged(movable bool) {
	// 0: (, movable bool), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar14movableChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:183
// index:0
// void allowedAreasChanged(Qt::ToolBarAreas)
func (this *QToolBar) AllowedAreasChanged(allowedAreas int) {
	// 0: (, QFlags<Qt::ToolBarArea> allowedAreas), (&allowedAreas)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar19allowedAreasChangedE6QFlagsIN2Qt11ToolBarAreaEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &allowedAreas)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:184
// index:0
// void orientationChanged(Qt::Orientation)
func (this *QToolBar) OrientationChanged(orientation int) {
	// 0: (, orientation Qt::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar18orientationChangedEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:185
// index:0
// void iconSizeChanged(const class QSize &)
func (this *QToolBar) IconSizeChanged(iconSize unsafe.Pointer) {
	// 0: (, iconSize const QSize &), (iconSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15iconSizeChangedERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), iconSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:186
// index:0
// void toolButtonStyleChanged(Qt::ToolButtonStyle)
func (this *QToolBar) ToolButtonStyleChanged(toolButtonStyle int) {
	// 0: (, toolButtonStyle Qt::ToolButtonStyle), (&toolButtonStyle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:187
// index:0
// void topLevelChanged(_Bool)
func (this *QToolBar) TopLevelChanged(topLevel bool) {
	// 0: (, topLevel bool), (&topLevel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar15topLevelChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &topLevel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:188
// index:0
// void visibilityChanged(_Bool)
func (this *QToolBar) VisibilityChanged(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar17visibilityChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:191
// index:0
// virtual
// void actionEvent(class QActionEvent *)
func (this *QToolBar) ActionEvent(event unsafe.Pointer) {
	// 0: (, event QActionEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:192
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QToolBar) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:193
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QToolBar) PaintEvent(event unsafe.Pointer) {
	// 0: (, event QPaintEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:194
// index:0
// virtual
// bool event(class QEvent *)
func (this *QToolBar) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBar5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:195
// index:0
// void initStyleOption(class QStyleOptionToolBar *)
func (this *QToolBar) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionToolBar *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
