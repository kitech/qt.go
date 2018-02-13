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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void initStyleOption(class QStyleOption *)
func (this *QGraphicsWidget) InheritInitStyleOption(f func(option *QStyleOption /*777 QStyleOption **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsWidget) InheritSizeHint(f func(which int, constraint *qtcore.QSizeF) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sizeHint", f)
}

// void updateGeometry()
func (this *QGraphicsWidget) InheritUpdateGeometry(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateGeometry", f)
}

// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
func (this *QGraphicsWidget) InheritItemChange(f func(change int, value *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "itemChange", f)
}

// QVariant propertyChange(const class QString &, const class QVariant &)
func (this *QGraphicsWidget) InheritPropertyChange(f func(propertyName string, value *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "propertyChange", f)
}

// bool sceneEvent(class QEvent *)
func (this *QGraphicsWidget) InheritSceneEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEvent", f)
}

// bool windowFrameEvent(class QEvent *)
func (this *QGraphicsWidget) InheritWindowFrameEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "windowFrameEvent", f)
}

// Qt::WindowFrameSection windowFrameSectionAt(const class QPointF &)
func (this *QGraphicsWidget) InheritWindowFrameSectionAt(f func(pos *qtcore.QPointF) int) {
	qtrt.SetAllInheritCallback(this, "windowFrameSectionAt", f)
}

// bool event(class QEvent *)
func (this *QGraphicsWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void changeEvent(class QEvent *)
func (this *QGraphicsWidget) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QGraphicsWidget) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsWidget) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QGraphicsWidget) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsWidget) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QGraphicsWidget) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void moveEvent(class QGraphicsSceneMoveEvent *)
func (this *QGraphicsWidget) InheritMoveEvent(f func(event *QGraphicsSceneMoveEvent /*777 QGraphicsSceneMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void polishEvent()
func (this *QGraphicsWidget) InheritPolishEvent(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "polishEvent", f)
}

// void resizeEvent(class QGraphicsSceneResizeEvent *)
func (this *QGraphicsWidget) InheritResizeEvent(f func(event *QGraphicsSceneResizeEvent /*777 QGraphicsSceneResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QGraphicsWidget) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) InheritHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) InheritHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// void grabMouseEvent(class QEvent *)
func (this *QGraphicsWidget) InheritGrabMouseEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "grabMouseEvent", f)
}

// void ungrabMouseEvent(class QEvent *)
func (this *QGraphicsWidget) InheritUngrabMouseEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "ungrabMouseEvent", f)
}

// void grabKeyboardEvent(class QEvent *)
func (this *QGraphicsWidget) InheritGrabKeyboardEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "grabKeyboardEvent", f)
}

// void ungrabKeyboardEvent(class QEvent *)
func (this *QGraphicsWidget) InheritUngrabKeyboardEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "ungrabKeyboardEvent", f)
}

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
// [8] const QMetaObject * metaObject()
func (this *QGraphicsWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsWidget(QGraphicsItem *, Qt::WindowFlags)
func NewQGraphicsWidget(parent *QGraphicsItem /*777 QGraphicsItem **/, wFlags int) *QGraphicsWidget {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsWidget()
func DeleteQGraphicsWidget(this *QGraphicsWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLayout * layout()
func (this *QGraphicsWidget) Layout() *QGraphicsLayout /*777 QGraphicsLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayout(QGraphicsLayout *)
func (this *QGraphicsWidget) SetLayout(layout *QGraphicsLayout /*777 QGraphicsLayout **/) {
	var convArg0 = layout.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()
func (this *QGraphicsWidget) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection()
func (this *QGraphicsWidget) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)
func (this *QGraphicsWidget) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetLayoutDirection()
func (this *QGraphicsWidget) UnsetLayoutDirection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget20unsetLayoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * style()
func (this *QGraphicsWidget) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)
func (this *QGraphicsWidget) SetStyle(style *QStyle /*777 QStyle **/) {
	var convArg0 = style.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget8setStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:96
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
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
func (this *QGraphicsWidget) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QPalette palette()
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
func (this *QGraphicsWidget) SetPalette(palette *qtgui.QPalette) {
	var convArg0 = palette.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoFillBackground()
func (this *QGraphicsWidget) AutoFillBackground() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget18autoFillBackgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoFillBackground(_Bool)
func (this *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget21setAutoFillBackgroundEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(const QSizeF &)
func (this *QGraphicsWidget) Resize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:106
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void resize(qreal, qreal)
func (this *QGraphicsWidget) Resize_1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:107
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size()
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
func (this *QGraphicsWidget) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:110
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setGeometry(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetGeometry_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF rect()
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
func (this *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18setContentsMarginsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFrameMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget21setWindowFrameMarginsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) GetWindowFrameMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetWindowFrameMargins()
func (this *QGraphicsWidget) UnsetWindowFrameMargins() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget23unsetWindowFrameMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:119
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF windowFrameGeometry()
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
// [32] QRectF windowFrameRect()
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
// [4] Qt::WindowFlags windowFlags()
func (this *QGraphicsWidget) WindowFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowType windowType()
func (this *QGraphicsWidget) WindowType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget10windowTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFlags(Qt::WindowFlags)
func (this *QGraphicsWidget) SetWindowFlags(wFlags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14setWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), wFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActiveWindow()
func (this *QGraphicsWidget) IsActiveWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget14isActiveWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowTitle(const QString &)
func (this *QGraphicsWidget) SetWindowTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14setWindowTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowTitle()
func (this *QGraphicsWidget) WindowTitle() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowTitleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::FocusPolicy focusPolicy()
func (this *QGraphicsWidget) FocusPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusPolicy(Qt::FocusPolicy)
func (this *QGraphicsWidget) SetFocusPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14setFocusPolicyEN2Qt11FocusPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTabOrder(QGraphicsWidget *, QGraphicsWidget *)
func (this *QGraphicsWidget) SetTabOrder(first *QGraphicsWidget /*777 QGraphicsWidget **/, second *QGraphicsWidget /*777 QGraphicsWidget **/) {
	var convArg0 = first.GetCthis()
	var convArg1 = second.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11setTabOrderEPS_S0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QGraphicsWidget_SetTabOrder(first *QGraphicsWidget /*777 QGraphicsWidget **/, second *QGraphicsWidget /*777 QGraphicsWidget **/) {
	var nilthis *QGraphicsWidget
	nilthis.SetTabOrder(first, second)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * focusWidget()
func (this *QGraphicsWidget) FocusWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int grabShortcut(const QKeySequence &, Qt::ShortcutContext)
func (this *QGraphicsWidget) GrabShortcut(sequence *qtgui.QKeySequence, context int) int {
	var convArg0 = sequence.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseShortcut(int)
func (this *QGraphicsWidget) ReleaseShortcut(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget15releaseShortcutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutEnabled(int, _Bool)
func (this *QGraphicsWidget) SetShortcutEnabled(id int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18setShortcutEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutAutoRepeat(int, _Bool)
func (this *QGraphicsWidget) SetShortcutAutoRepeat(id int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget21setShortcutAutoRepeatEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAction(QAction *)
func (this *QGraphicsWidget) AddAction(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9addActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertAction(QAction *, QAction *)
func (this *QGraphicsWidget) InsertAction(before *QAction /*777 QAction **/, action *QAction /*777 QAction **/) {
	var convArg0 = before.GetCthis()
	var convArg1 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12insertActionEP7QActionS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAction(QAction *)
func (this *QGraphicsWidget) RemoveAction(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12removeActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::WidgetAttribute, _Bool)
func (this *QGraphicsWidget) SetAttribute(attribute int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12setAttributeEN2Qt15WidgetAttributeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:159
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testAttribute(Qt::WidgetAttribute)
func (this *QGraphicsWidget) TestAttribute(attribute int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget13testAttributeEN2Qt15WidgetAttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attribute)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:164
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsWidget) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsWidget) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paintWindowFrame(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsWidget) PaintWindowFrame(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:168
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
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
// [8] QPainterPath shape()
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
func (this *QGraphicsWidget) GeometryChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget15geometryChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void layoutChanged()
func (this *QGraphicsWidget) LayoutChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget13layoutChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool close()
func (this *QGraphicsWidget) Close() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOption *)
func (this *QGraphicsWidget) InitStyleOption(option *QStyleOption /*777 QStyleOption **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsWidget) SizeHint(which int, constraint *qtcore.QSizeF) *qtcore.QSizeF /*123*/ {
	var convArg1 = constraint.GetCthis()
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
func (this *QGraphicsWidget) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const QVariant &)
func (this *QGraphicsWidget) ItemChange(change int, value *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var convArg1 = value.GetCthis()
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
func (this *QGraphicsWidget) PropertyChange(propertyName string, value *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(propertyName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = value.GetCthis()
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
func (this *QGraphicsWidget) SceneEvent(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10sceneEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:196
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool windowFrameEvent(QEvent *)
func (this *QGraphicsWidget) WindowFrameEvent(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget16windowFrameEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:197
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::WindowFrameSection windowFrameSectionAt(const QPointF &)
func (this *QGraphicsWidget) WindowFrameSectionAt(pos *qtcore.QPointF) int {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QGraphicsWidget) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QGraphicsWidget) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:203
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QGraphicsWidget) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QGraphicsWidget) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:207
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QGraphicsWidget) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QGraphicsWidget) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:209
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QGraphicsWidget) HideEvent(event *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:211
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QGraphicsSceneMoveEvent *)
func (this *QGraphicsWidget) MoveEvent(event *QGraphicsSceneMoveEvent /*777 QGraphicsSceneMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void polishEvent()
func (this *QGraphicsWidget) PolishEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11polishEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QGraphicsSceneResizeEvent *)
func (this *QGraphicsWidget) ResizeEvent(event *QGraphicsSceneResizeEvent /*777 QGraphicsSceneResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QGraphicsWidget) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) HoverMoveEvent(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) HoverLeaveEvent(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void grabMouseEvent(QEvent *)
func (this *QGraphicsWidget) GrabMouseEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget14grabMouseEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ungrabMouseEvent(QEvent *)
func (this *QGraphicsWidget) UngrabMouseEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void grabKeyboardEvent(QEvent *)
func (this *QGraphicsWidget) GrabKeyboardEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ungrabKeyboardEvent(QEvent *)
func (this *QGraphicsWidget) UngrabKeyboardEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QGraphicsWidget__ = int

const QGraphicsWidget__Type QGraphicsWidget__ = 11

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
