//  header block begin
// /usr/include/qt/QtWidgets/qgraphicswidget.h
// #include <qgraphicswidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QGraphicsWidget struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:66
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:84
// index:0
// virtual
// void ~QGraphicsWidget()
func DeleteQGraphicsWidget(*QGraphicsWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:85
// index:0
// QGraphicsLayout * layout()
func (this *QGraphicsWidget) Layout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget6layoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:86
// index:0
// void setLayout(class QGraphicsLayout *)
func (this *QGraphicsWidget) SetLayout(layout unsafe.Pointer) {
	// 0: (, QGraphicsLayout * layout), (layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout", ffiqt.FFI_TYPE_VOID, this.cthis, layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:87
// index:0
// void adjustSize()
func (this *QGraphicsWidget) AdjustSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget10adjustSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:89
// index:0
// Qt::LayoutDirection layoutDirection()
func (this *QGraphicsWidget) LayoutDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget15layoutDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:90
// index:0
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QGraphicsWidget) SetLayoutDirection(direction int) {
	// 0: (, Qt::LayoutDirection direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:91
// index:0
// void unsetLayoutDirection()
func (this *QGraphicsWidget) UnsetLayoutDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget20unsetLayoutDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:93
// index:0
// QStyle * style()
func (this *QGraphicsWidget) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget5styleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:94
// index:0
// void setStyle(class QStyle *)
func (this *QGraphicsWidget) SetStyle(style unsafe.Pointer) {
	// 0: (, QStyle * style), (style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget8setStyleEP6QStyle", ffiqt.FFI_TYPE_VOID, this.cthis, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:96
// index:0
// QFont font()
func (this *QGraphicsWidget) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:97
// index:0
// void setFont(const class QFont &)
func (this *QGraphicsWidget) SetFont(font unsafe.Pointer) {
	// 0: (, const QFont & font), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:99
// index:0
// QPalette palette()
func (this *QGraphicsWidget) Palette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget7paletteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:100
// index:0
// void setPalette(const class QPalette &)
func (this *QGraphicsWidget) SetPalette(palette unsafe.Pointer) {
	// 0: (, const QPalette & palette), (palette)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget10setPaletteERK8QPalette", ffiqt.FFI_TYPE_VOID, this.cthis, palette)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:102
// index:0
// bool autoFillBackground()
func (this *QGraphicsWidget) AutoFillBackground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget18autoFillBackgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:103
// index:0
// void setAutoFillBackground(_Bool)
func (this *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget21setAutoFillBackgroundEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:105
// index:0
// void resize(const class QSizeF &)
func (this *QGraphicsWidget) Resize(size unsafe.Pointer) {
	// 0: (, const QSizeF & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:106
// index:1
// inline
// void resize(qreal, qreal)
func (this *QGraphicsWidget) Resize_1(w float64, h float64) {
	// 1: (, qreal w, qreal h), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:107
// index:0
// QSizeF size()
func (this *QGraphicsWidget) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:109
// index:0
// virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsWidget) SetGeometry(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:110
// index:1
// inline
// void setGeometry(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetGeometry_1(x float64, y float64, w float64, h float64) {
	// 1: (, qreal x, qreal y, qreal w, qreal h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:111
// index:0
// inline
// QRectF rect()
func (this *QGraphicsWidget) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4rectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:113
// index:0
// void setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	// 0: (, qreal left, qreal top, qreal right, qreal bottom), (&left, &top, &right, &bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget18setContentsMarginsEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:114
// index:0
// virtual
// void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) GetContentsMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	// 0: (, qreal * left, qreal * top, qreal * right, qreal * bottom), (left, top, right, bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:116
// index:0
// void setWindowFrameMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	// 0: (, qreal left, qreal top, qreal right, qreal bottom), (&left, &top, &right, &bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget21setWindowFrameMarginsEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:117
// index:0
// void getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) GetWindowFrameMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	// 0: (, qreal * left, qreal * top, qreal * right, qreal * bottom), (left, top, right, bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:118
// index:0
// void unsetWindowFrameMargins()
func (this *QGraphicsWidget) UnsetWindowFrameMargins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget23unsetWindowFrameMarginsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:119
// index:0
// QRectF windowFrameGeometry()
func (this *QGraphicsWidget) WindowFrameGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget19windowFrameGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:120
// index:0
// QRectF windowFrameRect()
func (this *QGraphicsWidget) WindowFrameRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget15windowFrameRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:123
// index:0
// Qt::WindowFlags windowFlags()
func (this *QGraphicsWidget) WindowFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:124
// index:0
// Qt::WindowType windowType()
func (this *QGraphicsWidget) WindowType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget10windowTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:126
// index:0
// bool isActiveWindow()
func (this *QGraphicsWidget) IsActiveWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget14isActiveWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:127
// index:0
// void setWindowTitle(const class QString &)
func (this *QGraphicsWidget) SetWindowTitle(title unsafe.Pointer) {
	// 0: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14setWindowTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:128
// index:0
// QString windowTitle()
func (this *QGraphicsWidget) WindowTitle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowTitleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:131
// index:0
// Qt::FocusPolicy focusPolicy()
func (this *QGraphicsWidget) FocusPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusPolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:132
// index:0
// void setFocusPolicy(Qt::FocusPolicy)
func (this *QGraphicsWidget) SetFocusPolicy(policy int) {
	// 0: (, Qt::FocusPolicy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14setFocusPolicyEN2Qt11FocusPolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:133
// index:0
// static
// void setTabOrder(class QGraphicsWidget *, class QGraphicsWidget *)
func (this *QGraphicsWidget) SetTabOrder(first unsafe.Pointer, second unsafe.Pointer) {
	// 0: (QGraphicsWidget * first, QGraphicsWidget * second), (first, second)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11setTabOrderEPS_S0_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGraphicsWidget_SetTabOrder(first unsafe.Pointer, second unsafe.Pointer) {
	// 0: (QGraphicsWidget * first, QGraphicsWidget * second), (first, second)
	var nilthis *QGraphicsWidget
	nilthis.SetTabOrder(first, second)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:134
// index:0
// QGraphicsWidget * focusWidget()
func (this *QGraphicsWidget) FocusWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:137
// index:0
// int grabShortcut(const class QKeySequence &, Qt::ShortcutContext)
func (this *QGraphicsWidget) GrabShortcut(sequence unsafe.Pointer, context int) {
	// 0: (, const QKeySequence & sequence, Qt::ShortcutContext context), (sequence, &context)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_VOID, this.cthis, sequence, &context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:138
// index:0
// void releaseShortcut(int)
func (this *QGraphicsWidget) ReleaseShortcut(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget15releaseShortcutEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:139
// index:0
// void setShortcutEnabled(int, _Bool)
func (this *QGraphicsWidget) SetShortcutEnabled(id int, enabled bool) {
	// 0: (, int id, bool enabled), (&id, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget18setShortcutEnabledEib", ffiqt.FFI_TYPE_VOID, this.cthis, &id, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:140
// index:0
// void setShortcutAutoRepeat(int, _Bool)
func (this *QGraphicsWidget) SetShortcutAutoRepeat(id int, enabled bool) {
	// 0: (, int id, bool enabled), (&id, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget21setShortcutAutoRepeatEib", ffiqt.FFI_TYPE_VOID, this.cthis, &id, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:145
// index:0
// void addAction(class QAction *)
func (this *QGraphicsWidget) AddAction(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget9addActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:153
// index:0
// void insertAction(class QAction *, class QAction *)
func (this *QGraphicsWidget) InsertAction(before unsafe.Pointer, action unsafe.Pointer) {
	// 0: (, QAction * before, QAction * action), (before, action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12insertActionEP7QActionS1_", ffiqt.FFI_TYPE_VOID, this.cthis, before, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:154
// index:0
// void removeAction(class QAction *)
func (this *QGraphicsWidget) RemoveAction(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12removeActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:155
// index:0
// QList<QAction *> actions()
func (this *QGraphicsWidget) Actions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget7actionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:158
// index:0
// void setAttribute(Qt::WidgetAttribute, _Bool)
func (this *QGraphicsWidget) SetAttribute(attribute int, on bool) {
	// 0: (, Qt::WidgetAttribute attribute, bool on), (&attribute, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12setAttributeEN2Qt15WidgetAttributeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &attribute, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:159
// index:0
// bool testAttribute(Qt::WidgetAttribute)
func (this *QGraphicsWidget) TestAttribute(attribute int) {
	// 0: (, Qt::WidgetAttribute attribute), (&attribute)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget13testAttributeEN2Qt15WidgetAttributeE", ffiqt.FFI_TYPE_VOID, this.cthis, &attribute)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:164
// index:0
// virtual
// int type()
func (this *QGraphicsWidget) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:166
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:167
// index:0
// virtual
// void paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) PaintWindowFrame(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:168
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsWidget) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:169
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsWidget) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:178
// index:0
// void geometryChanged()
func (this *QGraphicsWidget) GeometryChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget15geometryChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:179
// index:0
// void layoutChanged()
func (this *QGraphicsWidget) LayoutChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget13layoutChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:182
// index:0
// bool close()
func (this *QGraphicsWidget) Close() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget5closeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
