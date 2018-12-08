package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicswidget.h
// #include <qgraphicswidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void initStyleOption(QStyleOption *)
func (this *QGraphicsWidget) InheritInitStyleOption(f func(option *QStyleOption /*777 QStyleOption **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// QSizeF sizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsWidget) InheritSizeHint(f func(which int, constraint *qtcore.QSizeF) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sizeHint", f)
}

// void updateGeometry()
func (this *QGraphicsWidget) InheritUpdateGeometry(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateGeometry", f)
}

// QVariant itemChange(QGraphicsItem::GraphicsItemChange, const QVariant &)
func (this *QGraphicsWidget) InheritItemChange(f func(change int, value *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "itemChange", f)
}

// QVariant propertyChange(const QString &, const QVariant &)
func (this *QGraphicsWidget) InheritPropertyChange(f func(propertyName string, value *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "propertyChange", f)
}

// bool sceneEvent(QEvent *)
func (this *QGraphicsWidget) InheritSceneEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEvent", f)
}

// bool windowFrameEvent(QEvent *)
func (this *QGraphicsWidget) InheritWindowFrameEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "windowFrameEvent", f)
}

// Qt::WindowFrameSection windowFrameSectionAt(const QPointF &)
func (this *QGraphicsWidget) InheritWindowFrameSectionAt(f func(pos *qtcore.QPointF) int) {
	qtrt.SetAllInheritCallback(this, "windowFrameSectionAt", f)
}

// bool event(QEvent *)
func (this *QGraphicsWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void changeEvent(QEvent *)
func (this *QGraphicsWidget) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(QCloseEvent *)
func (this *QGraphicsWidget) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QGraphicsWidget) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// bool focusNextPrevChild(bool)
func (this *QGraphicsWidget) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QGraphicsWidget) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QGraphicsWidget) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void moveEvent(QGraphicsSceneMoveEvent *)
func (this *QGraphicsWidget) InheritMoveEvent(f func(event *QGraphicsSceneMoveEvent /*777 QGraphicsSceneMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void polishEvent()
func (this *QGraphicsWidget) InheritPolishEvent(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "polishEvent", f)
}

// void resizeEvent(QGraphicsSceneResizeEvent *)
func (this *QGraphicsWidget) InheritResizeEvent(f func(event *QGraphicsSceneResizeEvent /*777 QGraphicsSceneResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QGraphicsWidget) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hoverMoveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) InheritHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) InheritHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// void grabMouseEvent(QEvent *)
func (this *QGraphicsWidget) InheritGrabMouseEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "grabMouseEvent", f)
}

// void ungrabMouseEvent(QEvent *)
func (this *QGraphicsWidget) InheritUngrabMouseEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "ungrabMouseEvent", f)
}

// void grabKeyboardEvent(QEvent *)
func (this *QGraphicsWidget) InheritGrabKeyboardEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "grabKeyboardEvent", f)
}

// void ungrabKeyboardEvent(QEvent *)
func (this *QGraphicsWidget) InheritUngrabKeyboardEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "ungrabKeyboardEvent", f)
}

/*

 */
type QGraphicsWidget struct {
	*QGraphicsLayoutItem
}
type QGraphicsWidget_ITF interface {
	QGraphicsLayoutItem_ITF
	QGraphicsWidget_PTR() *QGraphicsWidget
}

func (ptr *QGraphicsWidget) QGraphicsWidget_PTR() *QGraphicsWidget { return ptr }

func (this *QGraphicsWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayoutItem.GetCthis()
	}
}
func (this *QGraphicsWidget) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayoutItem = NewQGraphicsLayoutItemFromPointer(cthis)
}
func NewQGraphicsWidgetFromPointer(cthis unsafe.Pointer) *QGraphicsWidget {
	bcthis0 := NewQGraphicsLayoutItemFromPointer(cthis)
	return &QGraphicsWidget{bcthis0}
}
func (*QGraphicsWidget) NewFromPointer(cthis unsafe.Pointer) *QGraphicsWidget {
	return NewQGraphicsWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsWidget(QGraphicsItem *, Qt::WindowFlags)

/*
Constructs a QGraphicsWidget instance. The optional parent argument is passed to QGraphicsItem's constructor. The optional wFlags argument specifies the widget's window flags (e.g., whether the widget should be a window, a tool, a popup, etc).
*/
func (*QGraphicsWidget) NewForInherit(parent QGraphicsItem_ITF /*777 QGraphicsItem **/, wFlags int) *QGraphicsWidget {
	return NewQGraphicsWidget(parent, wFlags)
}
func NewQGraphicsWidget(parent QGraphicsItem_ITF /*777 QGraphicsItem **/, wFlags int) *QGraphicsWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsWidget(QGraphicsItem *, Qt::WindowFlags)

/*
Constructs a QGraphicsWidget instance. The optional parent argument is passed to QGraphicsItem's constructor. The optional wFlags argument specifies the widget's window flags (e.g., whether the widget should be a window, a tool, a popup, etc).
*/
func (*QGraphicsWidget) NewForInheritp() *QGraphicsWidget {
	return NewQGraphicsWidgetp()
}
func NewQGraphicsWidgetp() *QGraphicsWidget {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	wFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsWidget(QGraphicsItem *, Qt::WindowFlags)

/*
Constructs a QGraphicsWidget instance. The optional parent argument is passed to QGraphicsItem's constructor. The optional wFlags argument specifies the widget's window flags (e.g., whether the widget should be a window, a tool, a popup, etc).
*/
func (*QGraphicsWidget) NewForInheritp1(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsWidget {
	return NewQGraphicsWidgetp1(parent)
}
func NewQGraphicsWidgetp1(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	wFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsWidget()

/*

 */
func DeleteQGraphicsWidget(this *QGraphicsWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLayout * layout() const

/*
Returns this widget's layout, or 0 if no layout is currently managing this widget.

Note: Getter function for property layout.

See also setLayout().
*/
func (this *QGraphicsWidget) Layout() *QGraphicsLayout /*777 QGraphicsLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayout(QGraphicsLayout *)

/*
Sets the layout for this widget to layout. Any existing layout manager is deleted before the new layout is assigned. If layout is 0, the widget is left without a layout. Existing subwidgets' geometries will remain unaffected.

All widgets that are currently managed by layout or all of its sublayouts, are automatically reparented to this item. The layout is then invalidated, and the child widget geometries are adjusted according to this item's geometry() and contentsMargins(). Children who are not explicitly managed by layout remain unaffected by the layout after it has been assigned to this widget.

QGraphicsWidget takes ownership of layout.

Note: Setter function for property layout.

See also layout(), QGraphicsLinearLayout::addItem(), and QGraphicsLayout::invalidate().
*/
func (this *QGraphicsWidget) SetLayout(layout QGraphicsLayout_ITF /*777 QGraphicsLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QGraphicsLayout_PTR() != nil {
		convArg0 = layout.QGraphicsLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()

/*
Adjusts the size of the widget to its effective preferred size hint.

This function is called implicitly when the item is shown for the first time.

See also effectiveSizeHint() and Qt::MinimumSize.
*/
func (this *QGraphicsWidget) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection() const

/*

 */
func (this *QGraphicsWidget) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)

/*

 */
func (this *QGraphicsWidget) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetLayoutDirection()

/*

 */
func (this *QGraphicsWidget) UnsetLayoutDirection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget20unsetLayoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * style() const

/*
Returns a pointer to the widget's style. If this widget does not have any explicitly assigned style, the scene's style is returned instead. In turn, if the scene does not have any assigned style, this function returns QApplication::style().

See also setStyle().
*/
func (this *QGraphicsWidget) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)

/*
Sets the widget's style to style. QGraphicsWidget does not take ownership of style.

If no style is assigned, or style is 0, the widget will use QGraphicsScene::style() (if this has been set). Otherwise the widget will use QApplication::style().

This function sets the Qt::WA_SetStyle attribute if style is not 0; otherwise it clears the attribute.

See also style().
*/
func (this *QGraphicsWidget) SetStyle(style QStyle_ITF /*777 QStyle **/) {
	var convArg0 unsafe.Pointer
	if style != nil && style.QStyle_PTR() != nil {
		convArg0 = style.QStyle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget8setStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:96
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
func (this *QGraphicsWidget) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QGraphicsWidget) SetFont(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QPalette palette() const

/*

 */
func (this *QGraphicsWidget) Palette() *qtgui.QPalette /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget7paletteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &)

/*

 */
func (this *QGraphicsWidget) SetPalette(palette qtgui.QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if palette != nil && palette.QPalette_PTR() != nil {
		convArg0 = palette.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoFillBackground() const

/*

 */
func (this *QGraphicsWidget) AutoFillBackground() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget18autoFillBackgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoFillBackground(bool)

/*

 */
func (this *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget21setAutoFillBackgroundEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(const QSizeF &)

/*

 */
func (this *QGraphicsWidget) Resize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:106
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void resize(qreal, qreal)

/*

 */
func (this *QGraphicsWidget) Resize1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:107
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size() const

/*

 */
func (this *QGraphicsWidget) Size() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:109
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)

/*

 */
func (this *QGraphicsWidget) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:110
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setGeometry(qreal, qreal, qreal, qreal)

/*

 */
func (this *QGraphicsWidget) SetGeometry1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF rect() const

/*
Returns the item's local rect as a QRectF. This function is equivalent to QRectF(QPointF(), size()).

See also setGeometry() and resize().
*/
func (this *QGraphicsWidget) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(qreal, qreal, qreal, qreal)

/*
Sets the widget's contents margins to left, top, right and bottom.

Contents margins are used by the assigned layout to define the placement of subwidgets and layouts. Margins are particularly useful for widgets that constrain subwidgets to only a section of its own geometry. For example, a group box with a layout will place subwidgets inside its frame, but below the title.

Changing a widget's contents margins will always trigger an update(), and any assigned layout will be activated automatically. The widget will then receive a ContentsRectChange event.

See also getContentsMargins() and setGeometry().
*/
func (this *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18setContentsMarginsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *) const

/*
Reimplemented from QGraphicsLayoutItem::getContentsMargins().

Gets the widget's contents margins. The margins are stored in left, top, right and bottom, as pointers to qreals. Each argument can be omitted by passing 0.

See also setContentsMargins().
*/
func (this *QGraphicsWidget) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFrameMargins(qreal, qreal, qreal, qreal)

/*
Sets the widget's window frame margins to left, top, right and bottom. The default frame margins are provided by the style, and they depend on the current window flags.

If you would like to draw your own window decoration, you can set your own frame margins to override the default margins.

See also unsetWindowFrameMargins(), getWindowFrameMargins(), and windowFrameRect().
*/
func (this *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget21setWindowFrameMarginsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *) const

/*
Gets the widget's window frame margins. The margins are stored in left, top, right and bottom as pointers to qreals. Each argument can be omitted by passing 0.

See also setWindowFrameMargins() and windowFrameRect().
*/
func (this *QGraphicsWidget) GetWindowFrameMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetWindowFrameMargins()

/*
Resets the window frame margins to the default value, provided by the style.

See also setWindowFrameMargins(), getWindowFrameMargins(), and windowFrameRect().
*/
func (this *QGraphicsWidget) UnsetWindowFrameMargins() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget23unsetWindowFrameMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:119
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF windowFrameGeometry() const

/*
Returns the widget's geometry in parent coordinates including any window frame.

See also windowFrameRect(), getWindowFrameMargins(), and setWindowFrameMargins().
*/
func (this *QGraphicsWidget) WindowFrameGeometry() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget19windowFrameGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:120
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF windowFrameRect() const

/*
Returns the widget's local rect including any window frame.

See also windowFrameGeometry(), getWindowFrameMargins(), and setWindowFrameMargins().
*/
func (this *QGraphicsWidget) WindowFrameRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget15windowFrameRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:123
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowFlags windowFlags() const

/*

 */
func (this *QGraphicsWidget) WindowFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowType windowType() const

/*
Returns the widgets window type.

See also windowFlags(), isWindow(), and isPanel().
*/
func (this *QGraphicsWidget) WindowType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget10windowTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFlags(Qt::WindowFlags)

/*

 */
func (this *QGraphicsWidget) SetWindowFlags(wFlags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14setWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), wFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActiveWindow() const

/*
Returns true if this widget's window is in the active window, or if the widget does not have a window but is in an active scene (i.e., a scene that currently has focus).

The active window is the window that either contains a child widget that currently has input focus, or that itself has input focus.

See also QGraphicsScene::activeWindow(), QGraphicsScene::setActiveWindow(), and isActive().
*/
func (this *QGraphicsWidget) IsActiveWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget14isActiveWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowTitle(const QString &)

/*

 */
func (this *QGraphicsWidget) SetWindowTitle(title string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14setWindowTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowTitle() const

/*

 */
func (this *QGraphicsWidget) WindowTitle() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowTitleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::FocusPolicy focusPolicy() const

/*

 */
func (this *QGraphicsWidget) FocusPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusPolicy(Qt::FocusPolicy)

/*

 */
func (this *QGraphicsWidget) SetFocusPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14setFocusPolicyEN2Qt11FocusPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTabOrder(QGraphicsWidget *, QGraphicsWidget *)

/*
Moves the second widget around the ring of focus widgets so that keyboard focus moves from the first widget to the second widget when the Tab key is pressed.

Note that since the tab order of the second widget is changed, you should order a chain like this:


  setTabOrder(a, b); // a to b
  setTabOrder(b, c); // a to b to c
  setTabOrder(c, d); // a to b to c to d



not like this:


  // WRONG
  setTabOrder(c, d); // c to d
  setTabOrder(a, b); // a to b AND c to d
  setTabOrder(b, c); // a to b to c, but not c to d



If first is 0, this indicates that second should be the first widget to receive input focus should the scene gain Tab focus (i.e., the user hits Tab so that focus passes into the scene). If second is 0, this indicates that first should be the first widget to gain focus if the scene gained BackTab focus.

By default, tab order is defined implicitly using widget creation order.

See also focusPolicy and Keyboard Focus in Widgets.
*/
func (this *QGraphicsWidget) SetTabOrder(first QGraphicsWidget_ITF /*777 QGraphicsWidget **/, second QGraphicsWidget_ITF /*777 QGraphicsWidget **/) {
	var convArg0 unsafe.Pointer
	if first != nil && first.QGraphicsWidget_PTR() != nil {
		convArg0 = first.QGraphicsWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if second != nil && second.QGraphicsWidget_PTR() != nil {
		convArg1 = second.QGraphicsWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11setTabOrderEPS_S0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QGraphicsWidget_SetTabOrder(first QGraphicsWidget_ITF /*777 QGraphicsWidget **/, second QGraphicsWidget_ITF /*777 QGraphicsWidget **/) {
	var nilthis *QGraphicsWidget
	nilthis.SetTabOrder(first, second)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * focusWidget() const

/*
If this widget, a child or descendant of this widget currently has input focus, this function will return a pointer to that widget. If no descendant widget has input focus, 0 is returned.

See also QGraphicsItem::focusItem() and QWidget::focusWidget().
*/
func (this *QGraphicsWidget) FocusWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int grabShortcut(const QKeySequence &, Qt::ShortcutContext)

/*
Adds a shortcut to Qt's shortcut system that watches for the given key sequence in the given context. If the context is Qt::ApplicationShortcut, the shortcut applies to the application as a whole. Otherwise, it is either local to this widget, Qt::WidgetShortcut, or to the window itself, Qt::WindowShortcut. For widgets that are not part of a window (i.e., top-level widgets and their children), Qt::WindowShortcut shortcuts apply to the scene.

If the same key sequence has been grabbed by several widgets, when the key sequence occurs a QEvent::Shortcut event is sent to all the widgets to which it applies in a non-deterministic order, but with the ``ambiguous'' flag set to true.

Warning: You should not normally need to use this function; instead create QActions with the shortcut key sequences you require (if you also want equivalent menu options and toolbar buttons), or create QShortcuts if you just need key sequences. Both QAction and QShortcut handle all the event filtering for you, and provide signals which are triggered when the user triggers the key sequence, so are much easier to use than this low-level function.

This function was introduced in  Qt 4.5.

See also releaseShortcut(), setShortcutEnabled(), and QWidget::grabShortcut().
*/
func (this *QGraphicsWidget) GrabShortcut(sequence qtgui.QKeySequence_ITF, context int) int {
	var convArg0 unsafe.Pointer
	if sequence != nil && sequence.QKeySequence_PTR() != nil {
		convArg0 = sequence.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int grabShortcut(const QKeySequence &, Qt::ShortcutContext)

/*
Adds a shortcut to Qt's shortcut system that watches for the given key sequence in the given context. If the context is Qt::ApplicationShortcut, the shortcut applies to the application as a whole. Otherwise, it is either local to this widget, Qt::WidgetShortcut, or to the window itself, Qt::WindowShortcut. For widgets that are not part of a window (i.e., top-level widgets and their children), Qt::WindowShortcut shortcuts apply to the scene.

If the same key sequence has been grabbed by several widgets, when the key sequence occurs a QEvent::Shortcut event is sent to all the widgets to which it applies in a non-deterministic order, but with the ``ambiguous'' flag set to true.

Warning: You should not normally need to use this function; instead create QActions with the shortcut key sequences you require (if you also want equivalent menu options and toolbar buttons), or create QShortcuts if you just need key sequences. Both QAction and QShortcut handle all the event filtering for you, and provide signals which are triggered when the user triggers the key sequence, so are much easier to use than this low-level function.

This function was introduced in  Qt 4.5.

See also releaseShortcut(), setShortcutEnabled(), and QWidget::grabShortcut().
*/
func (this *QGraphicsWidget) GrabShortcutp(sequence qtgui.QKeySequence_ITF) int {
	var convArg0 unsafe.Pointer
	if sequence != nil && sequence.QKeySequence_PTR() != nil {
		convArg0 = sequence.QKeySequence_PTR().GetCthis()
	}
	// arg: 1, Qt::ShortcutContext=Elaborated, Qt::ShortcutContext=Enum, , Invalid
	context := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseShortcut(int)

/*
Removes the shortcut with the given id from Qt's shortcut system. The widget will no longer receive QEvent::Shortcut events for the shortcut's key sequence (unless it has other shortcuts with the same key sequence).

Warning: You should not normally need to use this function since Qt's shortcut system removes shortcuts automatically when their parent widget is destroyed. It is best to use QAction or QShortcut to handle shortcuts, since they are easier to use than this low-level function. Note also that this is an expensive operation.

This function was introduced in  Qt 4.5.

See also grabShortcut(), setShortcutEnabled(), and QWidget::releaseShortcut().
*/
func (this *QGraphicsWidget) ReleaseShortcut(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget15releaseShortcutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutEnabled(int, bool)

/*
If enabled is true, the shortcut with the given id is enabled; otherwise the shortcut is disabled.

Warning: You should not normally need to use this function since Qt's shortcut system enables/disables shortcuts automatically as widgets become hidden/visible and gain or lose focus. It is best to use QAction or QShortcut to handle shortcuts, since they are easier to use than this low-level function.

This function was introduced in  Qt 4.5.

See also grabShortcut(), releaseShortcut(), and QWidget::setShortcutEnabled().
*/
func (this *QGraphicsWidget) SetShortcutEnabled(id int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18setShortcutEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutEnabled(int, bool)

/*
If enabled is true, the shortcut with the given id is enabled; otherwise the shortcut is disabled.

Warning: You should not normally need to use this function since Qt's shortcut system enables/disables shortcuts automatically as widgets become hidden/visible and gain or lose focus. It is best to use QAction or QShortcut to handle shortcuts, since they are easier to use than this low-level function.

This function was introduced in  Qt 4.5.

See also grabShortcut(), releaseShortcut(), and QWidget::setShortcutEnabled().
*/
func (this *QGraphicsWidget) SetShortcutEnabledp(id int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18setShortcutEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutAutoRepeat(int, bool)

/*
If enabled is true, auto repeat of the shortcut with the given id is enabled; otherwise it is disabled.

This function was introduced in  Qt 4.5.

See also grabShortcut(), releaseShortcut(), and QWidget::setShortcutAutoRepeat().
*/
func (this *QGraphicsWidget) SetShortcutAutoRepeat(id int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget21setShortcutAutoRepeatEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutAutoRepeat(int, bool)

/*
If enabled is true, auto repeat of the shortcut with the given id is enabled; otherwise it is disabled.

This function was introduced in  Qt 4.5.

See also grabShortcut(), releaseShortcut(), and QWidget::setShortcutAutoRepeat().
*/
func (this *QGraphicsWidget) SetShortcutAutoRepeatp(id int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget21setShortcutAutoRepeatEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAction(QAction *)

/*
Appends the action action to this widget's list of actions.

All QGraphicsWidgets have a list of QActions, however they can be represented graphically in many different ways. The default use of the QAction list (as returned by actions()) is to create a context QMenu.

A QGraphicsWidget should only have one of each action and adding an action it already has will not cause the same action to be in the widget twice.

This function was introduced in  Qt 4.5.

See also removeAction(), insertAction(), actions(), and QWidget::addAction().
*/
func (this *QGraphicsWidget) AddAction(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9addActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertAction(QAction *, QAction *)

/*
Inserts the action action to this widget's list of actions, before the action before. It appends the action if before is 0 or before is not a valid action for this widget.

A QGraphicsWidget should only have one of each action.

This function was introduced in  Qt 4.5.

See also removeAction(), addAction(), QMenu, actions(), and QWidget::insertActions().
*/
func (this *QGraphicsWidget) InsertAction(before QAction_ITF /*777 QAction **/, action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg1 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12insertActionEP7QActionS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAction(QAction *)

/*
Removes the action action from this widget's list of actions.

This function was introduced in  Qt 4.5.

See also insertAction(), actions(), insertAction(), and QWidget::removeAction().
*/
func (this *QGraphicsWidget) RemoveAction(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12removeActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::WidgetAttribute, bool)

/*
If on is true, this function enables attribute; otherwise attribute is disabled.

See the class documentation for QGraphicsWidget for a complete list of which attributes are supported, and what they are for.

See also testAttribute() and QWidget::setAttribute().
*/
func (this *QGraphicsWidget) SetAttribute(attribute int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12setAttributeEN2Qt15WidgetAttributeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::WidgetAttribute, bool)

/*
If on is true, this function enables attribute; otherwise attribute is disabled.

See the class documentation for QGraphicsWidget for a complete list of which attributes are supported, and what they are for.

See also testAttribute() and QWidget::setAttribute().
*/
func (this *QGraphicsWidget) SetAttributep(attribute int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12setAttributeEN2Qt15WidgetAttributeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:159
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testAttribute(Qt::WidgetAttribute) const

/*
Returns true if attribute is enabled for this widget; otherwise, returns false.

See also setAttribute().
*/
func (this *QGraphicsWidget) TestAttribute(attribute int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget13testAttributeEN2Qt15WidgetAttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:164
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const

/*
Reimplemented from QGraphicsItem::type().
*/
func (this *QGraphicsWidget) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
Reimplemented from QGraphicsItem::paint().
*/
func (this *QGraphicsWidget) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
Reimplemented from QGraphicsItem::paint().
*/
func (this *QGraphicsWidget) Paintp(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paintWindowFrame(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
This virtual function is called by QGraphicsScene to draw the window frame for windows using painter, option, and widget, in local coordinates. The base implementation uses the current style to render the frame and title bar.

You can reimplement this function in a subclass of QGraphicsWidget to provide custom rendering of the widget's window frame.

See also QGraphicsItem::paint().
*/
func (this *QGraphicsWidget) PaintWindowFrame(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paintWindowFrame(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
This virtual function is called by QGraphicsScene to draw the window frame for windows using painter, option, and widget, in local coordinates. The base implementation uses the current style to render the frame and title bar.

You can reimplement this function in a subclass of QGraphicsWidget to provide custom rendering of the widget's window frame.

See also QGraphicsItem::paint().
*/
func (this *QGraphicsWidget) PaintWindowFramep(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:168
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect() const

/*
Reimplemented from QGraphicsItem::boundingRect().
*/
func (this *QGraphicsWidget) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:169
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape() const

/*
Reimplemented from QGraphicsItem::shape().
*/
func (this *QGraphicsWidget) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void geometryChanged()

/*
This signal gets emitted whenever the geometry is changed in setGeometry().

Note: Notifier signal for property geometry. Notifier signal for property size.
*/
func (this *QGraphicsWidget) GeometryChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget15geometryChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void layoutChanged()

/*

 */
func (this *QGraphicsWidget) LayoutChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget13layoutChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool close()

/*
Call this function to close the widget.

Returns true if the widget was closed; otherwise returns false. This slot will first send a QCloseEvent to the widget, which may or may not accept the event. If the event was ignored, nothing happens. If the event was accepted, it will hide() the widget.

If the widget has the Qt::WA_DeleteOnClose attribute set it will be deleted.
*/
func (this *QGraphicsWidget) Close() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOption *) const

/*
Populates a style option object for this widget based on its current state, and stores the output in option. The default implementation populates option with the following properties.


 Style Option PropertyValue
state & QStyle::State_EnabledCorresponds to QGraphicsItem::isEnabled().
state & QStyle::State_HasFocusCorresponds to QGraphicsItem::hasFocus().
state & QStyle::State_MouseOverCorresponds to QGraphicsItem::isUnderMouse().
directionCorresponds to QGraphicsWidget::layoutDirection().
rectCorresponds to QGraphicsWidget::rect().toRect().
paletteCorresponds to QGraphicsWidget::palette().
fontMetricsCorresponds to QFontMetrics(QGraphicsWidget::font()).


Subclasses of QGraphicsWidget should call the base implementation, and then test the type of option using qstyleoption_cast<>() or test QStyleOption::Type before storing widget-specific options.

For example:


  void MyGroupBoxWidget::initStyleOption(QStyleOption *option) const
  {
      QGraphicsWidget::initStyleOption(option);
      if (QStyleOptionGroupBox *box = qstyleoption_cast<QStyleOptionGroupBox *>(option)) {
          // Add group box specific state.
          box->flat = isFlat();
          ...
      }
  }



See also QStyleOption::initFrom().
*/
func (this *QGraphicsWidget) InitStyleOption(option QStyleOption_ITF /*777 QStyleOption **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg0 = option.QStyleOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsWidget) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsWidget) SizeHintp(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometry()

/*
Reimplemented from QGraphicsLayoutItem::updateGeometry().

If this widget is currently managed by a layout, this function notifies the layout that the widget's size hints have changed and the layout may need to resize and reposition the widget accordingly.

Call this function if the widget's sizeHint() has changed.

See also QGraphicsLayout::invalidate().
*/
func (this *QGraphicsWidget) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant itemChange(QGraphicsItem::GraphicsItemChange, const QVariant &)

/*
Reimplemented from QGraphicsItem::itemChange().

QGraphicsWidget uses the base implementation of this function to catch and deliver events related to state changes in the item. Because of this, it is very important that subclasses call the base implementation.

change specifies the type of change, and value is the new value.

For example, QGraphicsWidget uses ItemVisibleChange to deliver Show and Hide events, ItemPositionHasChanged to deliver Move events, and ItemParentChange both to deliver ParentChange events, and for managing the focus chain.

QGraphicsWidget enables the ItemSendsGeometryChanges flag by default in order to track position changes.

See also QGraphicsItem::itemChange().
*/
func (this *QGraphicsWidget) ItemChange(change int, value qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant propertyChange(const QString &, const QVariant &)

/*

 */
func (this *QGraphicsWidget) PropertyChange(propertyName string, value qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString5(propertyName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14propertyChangeERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:195
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool sceneEvent(QEvent *)

/*
Reimplemented from QGraphicsItem::sceneEvent().

QGraphicsWidget's implementation of sceneEvent() simply passes event to QGraphicsWidget::event(). You can handle all events for your widget in event() or in any of the convenience functions; you should not have to reimplement this function in a subclass of QGraphicsWidget.

See also QGraphicsItem::sceneEvent().
*/
func (this *QGraphicsWidget) SceneEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10sceneEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:196
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool windowFrameEvent(QEvent *)

/*
This event handler, for event, receives events for the window frame if this widget is a window. Its base implementation provides support for default window frame interaction such as moving, resizing, etc.

You can reimplement this handler in a subclass of QGraphicsWidget to provide your own custom window frame interaction support.

Returns true if event has been recognized and processed; otherwise, returns false.

See also event().
*/
func (this *QGraphicsWidget) WindowFrameEvent(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget16windowFrameEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:197
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::WindowFrameSection windowFrameSectionAt(const QPointF &) const

/*
Returns the window frame section at position pos, or Qt::NoSection if there is no window frame section at this position.

This function is used in QGraphicsWidget's base implementation for window frame interaction.

You can reimplement this function if you want to customize how a window can be interactively moved or resized. For instance, if you only want to allow a window to be resized by the bottom right corner, you can reimplement this function to return Qt::NoSection for all sections except Qt::BottomRightSection.

This function was introduced in  Qt 4.4.

See also windowFrameEvent(), paintWindowFrame(), and windowFrameGeometry().
*/
func (this *QGraphicsWidget) WindowFrameSectionAt(pos qtcore.QPointF_ITF) int {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().

Handles the event. QGraphicsWidget handles the following events:


 EventUsage
PolishDelivered to the widget some time after it has been shown.
GraphicsSceneMoveDelivered to the widget after its local position has changed.
GraphicsSceneResizeDelivered to the widget after its size has changed.
ShowDelivered to the widget before it has been shown.
HideDelivered to the widget after it has been hidden.
PaletteChangeDelivered to the widget after its palette has changed.
FontChangeDelivered to the widget after its font has changed.
EnabledChangeDelivered to the widget after its enabled state has changed.
StyleChangeDelivered to the widget after its style has changed.
LayoutDirectionChangeDelivered to the widget after its layout direction has changed.
ContentsRectChangeDelivered to the widget after its contents margins/ contents rect has changed.
*/
func (this *QGraphicsWidget) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
This event handler can be reimplemented to handle state changes.

The state being changed in this event can be retrieved through event.

Change events include: QEvent::ActivationChange, QEvent::EnabledChange, QEvent::FontChange, QEvent::StyleChange, QEvent::PaletteChange, QEvent::ParentChange, QEvent::LayoutDirectionChange, and QEvent::ContentsRectChange.
*/
func (this *QGraphicsWidget) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:203
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)

/*
This event handler, for event, can be reimplemented in a subclass to receive widget close events. The default implementation accepts the event.

See also close() and QCloseEvent.
*/
func (this *QGraphicsWidget) CloseEvent(event qtgui.QCloseEvent_ITF /*777 QCloseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QCloseEvent_PTR() != nil {
		convArg0 = event.QCloseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QGraphicsItem::focusInEvent().
*/
func (this *QGraphicsWidget) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:207
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Finds a new widget to give the keyboard focus to, as appropriate for Tab and Shift+Tab, and returns true if it can find a new widget; returns false otherwise. If next is true, this function searches forward; if next is false, it searches backward.

Sometimes, you will want to reimplement this function to provide special focus handling for your widget and its subwidgets. For example, a web browser might reimplement it to move its current active link forward or backward, and call the base implementation only when it reaches the last or first link on the page.

Child widgets call focusNextPrevChild() on their parent widgets, but only the window that contains the child widgets decides where to redirect focus. By reimplementing this function for an object, you gain control of focus traversal for all child widgets.

See also focusPolicy().
*/
func (this *QGraphicsWidget) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QGraphicsItem::focusOutEvent().
*/
func (this *QGraphicsWidget) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:209
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
This event handler, for Hide events, is delivered after the widget has been hidden, for example, setVisible(false) has been called for the widget or one of its ancestors when the widget was previously shown.

You can reimplement this event handler to detect when your widget is hidden. Calling QEvent::accept() or QEvent::ignore() on event has no effect.

See also showEvent(), QWidget::hideEvent(), and ItemVisibleChange.
*/
func (this *QGraphicsWidget) HideEvent(event qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHideEvent_PTR() != nil {
		convArg0 = event.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:211
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QGraphicsSceneMoveEvent *)

/*
This event handler, for GraphicsSceneMove events, is delivered after the widget has moved (e.g., its local position has changed).

This event is only delivered when the item is moved locally. Calling setTransform() or moving any of the item's ancestors does not affect the item's local position.

You can reimplement this event handler to detect when your widget has moved. Calling QEvent::accept() or QEvent::ignore() on event has no effect.

See also ItemPositionChange and ItemPositionHasChanged.
*/
func (this *QGraphicsWidget) MoveEvent(event QGraphicsSceneMoveEvent_ITF /*777 QGraphicsSceneMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMoveEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void polishEvent()

/*
This event is delivered to the item by the scene at some point after it has been constructed, but before it is shown or otherwise accessed through the scene. You can use this event handler to do last-minute initializations of the widget which require the item to be fully constructed.

The base implementation does nothing.
*/
func (this *QGraphicsWidget) PolishEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11polishEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QGraphicsSceneResizeEvent *)

/*
This event handler, for GraphicsSceneResize events, is delivered after the widget has been resized (i.e., its local size has changed). event contains both the old and the new size.

This event is only delivered when the widget is resized locally; calling setTransform() on the widget or any of its ancestors or view, does not affect the widget's local size.

You can reimplement this event handler to detect when your widget has been resized. Calling QEvent::accept() or QEvent::ignore() on event has no effect.

See also geometry() and setGeometry().
*/
func (this *QGraphicsWidget) ResizeEvent(event QGraphicsSceneResizeEvent_ITF /*777 QGraphicsSceneResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneResizeEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
This event handler, for Show events, is delivered before the widget has been shown, for example, setVisible(true) has been called for the widget or one of its ancestors when the widget was previously hidden.

You can reimplement this event handler to detect when your widget is shown. Calling QEvent::accept() or QEvent::ignore() on event has no effect.

See also hideEvent(), QWidget::showEvent(), and ItemVisibleChange.
*/
func (this *QGraphicsWidget) ShowEvent(event qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QShowEvent_PTR() != nil {
		convArg0 = event.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QGraphicsSceneHoverEvent *)

/*
Reimplemented from QGraphicsItem::hoverMoveEvent().
*/
func (this *QGraphicsWidget) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QGraphicsSceneHoverEvent *)

/*
Reimplemented from QGraphicsItem::hoverLeaveEvent().
*/
func (this *QGraphicsWidget) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void grabMouseEvent(QEvent *)

/*
This event handler, for event, can be reimplemented in a subclass to receive notifications for QEvent::GrabMouse events.

See also grabMouse() and grabKeyboard().
*/
func (this *QGraphicsWidget) GrabMouseEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14grabMouseEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ungrabMouseEvent(QEvent *)

/*
This event handler, for event, can be reimplemented in a subclass to receive notifications for QEvent::UngrabMouse events.

See also ungrabMouse() and ungrabKeyboard().
*/
func (this *QGraphicsWidget) UngrabMouseEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void grabKeyboardEvent(QEvent *)

/*
This event handler, for event, can be reimplemented in a subclass to receive notifications for QEvent::GrabKeyboard events.

See also grabKeyboard() and grabMouse().
*/
func (this *QGraphicsWidget) GrabKeyboardEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ungrabKeyboardEvent(QEvent *)

/*
This event handler, for event, can be reimplemented in a subclass to receive notifications for QEvent::UngrabKeyboard events.

See also ungrabKeyboard() and ungrabMouse().
*/
func (this *QGraphicsWidget) UngrabKeyboardEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QGraphicsWidget__ = int

//
const QGraphicsWidget__Type QGraphicsWidget__ = 11

func (this *QGraphicsWidget) ItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsWidget_ItemName(val int) string {
	var nilthis *QGraphicsWidget
	return nilthis.ItemName(val)
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
		log.Println(123)
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
