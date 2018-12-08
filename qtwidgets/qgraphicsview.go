package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsview.h
// #include <qgraphicsview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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

// void setupViewport(QWidget *)
func (this *QGraphicsView) InheritSetupViewport(f func(widget *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setupViewport", f)
}

// bool event(QEvent *)
func (this *QGraphicsView) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool viewportEvent(QEvent *)
func (this *QGraphicsView) InheritViewportEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QGraphicsView) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QGraphicsView) InheritDragEnterEvent(f func(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QGraphicsView) InheritDragLeaveEvent(f func(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QGraphicsView) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QGraphicsView) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QGraphicsView) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// bool focusNextPrevChild(bool)
func (this *QGraphicsView) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QGraphicsView) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QGraphicsView) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsView) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QGraphicsView) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QGraphicsView) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QGraphicsView) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QGraphicsView) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QGraphicsView) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QGraphicsView) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QGraphicsView) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QGraphicsView) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void showEvent(QShowEvent *)
func (this *QGraphicsView) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsView) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void drawBackground(QPainter *, const QRectF &)
func (this *QGraphicsView) InheritDrawBackground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawBackground", f)
}

// void drawForeground(QPainter *, const QRectF &)
func (this *QGraphicsView) InheritDrawForeground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawForeground", f)
}

// void drawItems(QPainter *, int, QGraphicsItem **, const QStyleOptionGraphicsItem *)
func (this *QGraphicsView) InheritDrawItems(f func(painter *qtgui.QPainter /*777 QPainter **/, numItems int, items []interface{}, options []interface{}) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawItems", f)
}

/*

 */
type QGraphicsView struct {
	*QAbstractScrollArea
}
type QGraphicsView_ITF interface {
	QAbstractScrollArea_ITF
	QGraphicsView_PTR() *QGraphicsView
}

func (ptr *QGraphicsView) QGraphicsView_PTR() *QGraphicsView { return ptr }

func (this *QGraphicsView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QGraphicsView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQGraphicsViewFromPointer(cthis unsafe.Pointer) *QGraphicsView {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QGraphicsView{bcthis0}
}
func (*QGraphicsView) NewFromPointer(cthis unsafe.Pointer) *QGraphicsView {
	return NewQGraphicsViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsView(QWidget *)

/*
Constructs a QGraphicsView. parent is passed to QWidget's constructor.
*/
func (*QGraphicsView) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QGraphicsView {
	return NewQGraphicsView(parent)
}
func NewQGraphicsView(parent QWidget_ITF /*777 QWidget **/) *QGraphicsView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsView")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsView(QWidget *)

/*
Constructs a QGraphicsView. parent is passed to QWidget's constructor.
*/
func (*QGraphicsView) NewForInheritp() *QGraphicsView {
	return NewQGraphicsViewp()
}
func NewQGraphicsViewp() *QGraphicsView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsView")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:118
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsView(QGraphicsScene *, QWidget *)

/*
Constructs a QGraphicsView. parent is passed to QWidget's constructor.
*/
func (*QGraphicsView) NewForInherit1(scene QGraphicsScene_ITF /*777 QGraphicsScene **/, parent QWidget_ITF /*777 QWidget **/) *QGraphicsView {
	return NewQGraphicsView1(scene, parent)
}
func NewQGraphicsView1(scene QGraphicsScene_ITF /*777 QGraphicsScene **/, parent QWidget_ITF /*777 QWidget **/) *QGraphicsView {
	var convArg0 unsafe.Pointer
	if scene != nil && scene.QGraphicsScene_PTR() != nil {
		convArg0 = scene.QGraphicsScene_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP14QGraphicsSceneP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsView")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:118
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsView(QGraphicsScene *, QWidget *)

/*
Constructs a QGraphicsView. parent is passed to QWidget's constructor.
*/
func (*QGraphicsView) NewForInherit1p(scene QGraphicsScene_ITF /*777 QGraphicsScene **/) *QGraphicsView {
	return NewQGraphicsView1p(scene)
}
func NewQGraphicsView1p(scene QGraphicsScene_ITF /*777 QGraphicsScene **/) *QGraphicsView {
	var convArg0 unsafe.Pointer
	if scene != nil && scene.QGraphicsScene_PTR() != nil {
		convArg0 = scene.QGraphicsScene_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP14QGraphicsSceneP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsView")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsView()

/*

 */
func DeleteQGraphicsView(this *QGraphicsView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:121
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QGraphicsView) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [4] QPainter::RenderHints renderHints() const

/*

 */
func (this *QGraphicsView) RenderHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView11renderHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHint(QPainter::RenderHint, bool)

/*
If enabled is true, the render hint hint is enabled; otherwise it is disabled.

See also renderHints.
*/
func (this *QGraphicsView) SetRenderHint(hint int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13setRenderHintEN8QPainter10RenderHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHint(QPainter::RenderHint, bool)

/*
If enabled is true, the render hint hint is enabled; otherwise it is disabled.

See also renderHints.
*/
func (this *QGraphicsView) SetRenderHintp(hint int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13setRenderHintEN8QPainter10RenderHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHints(QPainter::RenderHints)

/*

 */
func (this *QGraphicsView) SetRenderHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14setRenderHintsE6QFlagsIN8QPainter10RenderHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QGraphicsView) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*

 */
func (this *QGraphicsView) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::ViewportAnchor transformationAnchor() const

/*

 */
func (this *QGraphicsView) TransformationAnchor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView20transformationAnchorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformationAnchor(QGraphicsView::ViewportAnchor)

/*

 */
func (this *QGraphicsView) SetTransformationAnchor(anchor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView23setTransformationAnchorENS_14ViewportAnchorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), anchor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::ViewportAnchor resizeAnchor() const

/*

 */
func (this *QGraphicsView) ResizeAnchor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12resizeAnchorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeAnchor(QGraphicsView::ViewportAnchor)

/*

 */
func (this *QGraphicsView) SetResizeAnchor(anchor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15setResizeAnchorENS_14ViewportAnchorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), anchor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::ViewportUpdateMode viewportUpdateMode() const

/*

 */
func (this *QGraphicsView) ViewportUpdateMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView18viewportUpdateModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewportUpdateMode(QGraphicsView::ViewportUpdateMode)

/*

 */
func (this *QGraphicsView) SetViewportUpdateMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView21setViewportUpdateModeENS_18ViewportUpdateModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::OptimizationFlags optimizationFlags() const

/*

 */
func (this *QGraphicsView) OptimizationFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView17optimizationFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizationFlag(QGraphicsView::OptimizationFlag, bool)

/*
Enables flag if enabled is true; otherwise disables flag.

See also optimizationFlags.
*/
func (this *QGraphicsView) SetOptimizationFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView19setOptimizationFlagENS_16OptimizationFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizationFlag(QGraphicsView::OptimizationFlag, bool)

/*
Enables flag if enabled is true; otherwise disables flag.

See also optimizationFlags.
*/
func (this *QGraphicsView) SetOptimizationFlagp(flag int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView19setOptimizationFlagENS_16OptimizationFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizationFlags(QGraphicsView::OptimizationFlags)

/*

 */
func (this *QGraphicsView) SetOptimizationFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView20setOptimizationFlagsE6QFlagsINS_16OptimizationFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::DragMode dragMode() const

/*

 */
func (this *QGraphicsView) DragMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView8dragModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragMode(QGraphicsView::DragMode)

/*

 */
func (this *QGraphicsView) SetDragMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView11setDragModeENS_8DragModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemSelectionMode rubberBandSelectionMode() const

/*

 */
func (this *QGraphicsView) RubberBandSelectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView23rubberBandSelectionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRubberBandSelectionMode(Qt::ItemSelectionMode)

/*

 */
func (this *QGraphicsView) SetRubberBandSelectionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView26setRubberBandSelectionModeEN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:149
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rubberBandRect() const

/*
This functions returns the current rubber band area (in viewport coordinates) if the user is currently doing an itemselection with rubber band. When the user is not using the rubber band this functions returns (a null) QRectF().

Notice that part of this QRect can be outise the visual viewport. It can e.g contain negative values.

This function was introduced in  Qt 5.1.

See also rubberBandSelectionMode and rubberBandChanged().
*/
func (this *QGraphicsView) RubberBandRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView14rubberBandRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:152
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::CacheMode cacheMode() const

/*

 */
func (this *QGraphicsView) CacheMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView9cacheModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheMode(QGraphicsView::CacheMode)

/*

 */
func (this *QGraphicsView) SetCacheMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setCacheModeE6QFlagsINS_13CacheModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetCachedContent()

/*
Resets any cached content. Calling this function will clear QGraphicsView's cache. If the current cache mode is CacheNone, this function does nothing.

This function is called automatically for you when the backgroundBrush or QGraphicsScene::backgroundBrush properties change; you only need to call this function if you have reimplemented QGraphicsScene::drawBackground() or QGraphicsView::drawBackground() to draw a custom background, and need to trigger a full redraw.

See also cacheMode().
*/
func (this *QGraphicsView) ResetCachedContent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView18resetCachedContentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:156
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInteractive() const

/*

 */
func (this *QGraphicsView) IsInteractive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView13isInteractiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInteractive(bool)

/*

 */
func (this *QGraphicsView) SetInteractive(allowed bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14setInteractiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), allowed)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsScene * scene() const

/*
Returns a pointer to the scene that is currently visualized in the view. If no scene is currently visualized, 0 is returned.

See also setScene().
*/
func (this *QGraphicsView) Scene() *QGraphicsScene /*777 QGraphicsScene **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5sceneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScene(QGraphicsScene *)

/*
Sets the current scene to scene. If scene is already being viewed, this function does nothing.

When a scene is set on a view, the QGraphicsScene::changed() signal is automatically connected to this view's updateScene() slot, and the view's scroll bars are adjusted to fit the size of the scene.

The view does not take ownership of scene.

See also scene().
*/
func (this *QGraphicsView) SetScene(scene QGraphicsScene_ITF /*777 QGraphicsScene **/) {
	var convArg0 unsafe.Pointer
	if scene != nil && scene.QGraphicsScene_PTR() != nil {
		convArg0 = scene.QGraphicsScene_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView8setSceneEP14QGraphicsScene", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:162
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF sceneRect() const

/*

 */
func (this *QGraphicsView) SceneRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView9sceneRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSceneRect(const QRectF &)

/*

 */
func (this *QGraphicsView) SetSceneRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:164
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setSceneRect(qreal, qreal, qreal, qreal)

/*

 */
func (this *QGraphicsView) SetSceneRect1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:166
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix matrix() const

/*
Returns the current transformation matrix for the view. If no current transformation is set, the identity matrix is returned.

See also setMatrix(), transform(), rotate(), scale(), shear(), and translate().
*/
func (this *QGraphicsView) Matrix() *qtgui.QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &, bool)

/*
Sets the view's current transformation matrix to matrix.

If combine is true, then matrix is combined with the current matrix; otherwise, matrix replaces the current matrix. combine is false by default.

The transformation matrix tranforms the scene into view coordinates. Using the default transformation, provided by the identity matrix, one pixel in the view represents one unit in the scene (e.g., a 10x10 rectangular item is drawn using 10x10 pixels in the view). If a 2x2 scaling matrix is applied, the scene will be drawn in 1:2 (e.g., a 10x10 rectangular item is then drawn using 20x20 pixels in the view).

Example:


  QGraphicsScene scene;
  scene.addText("GraphicsView rotated clockwise");

  QGraphicsView view(&scene);
  view.rotate(90); // the text is rendered with a 90 degree clockwise rotation
  view.show();



To simplify interation with items using a transformed view, QGraphicsView provides mapTo... and mapFrom... functions that can translate between scene and view coordinates. For example, you can call mapToScene() to map a view coordinate to a floating point scene coordinate, or mapFromScene() to map from floating point scene coordinates to view coordinates.

See also matrix(), setTransform(), rotate(), scale(), shear(), and translate().
*/
func (this *QGraphicsView) SetMatrix(matrix qtgui.QMatrix_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &, bool)

/*
Sets the view's current transformation matrix to matrix.

If combine is true, then matrix is combined with the current matrix; otherwise, matrix replaces the current matrix. combine is false by default.

The transformation matrix tranforms the scene into view coordinates. Using the default transformation, provided by the identity matrix, one pixel in the view represents one unit in the scene (e.g., a 10x10 rectangular item is drawn using 10x10 pixels in the view). If a 2x2 scaling matrix is applied, the scene will be drawn in 1:2 (e.g., a 10x10 rectangular item is then drawn using 20x20 pixels in the view).

Example:


  QGraphicsScene scene;
  scene.addText("GraphicsView rotated clockwise");

  QGraphicsView view(&scene);
  view.rotate(90); // the text is rendered with a 90 degree clockwise rotation
  view.show();



To simplify interation with items using a transformed view, QGraphicsView provides mapTo... and mapFrom... functions that can translate between scene and view coordinates. For example, you can call mapToScene() to map a view coordinate to a floating point scene coordinate, or mapFromScene() to map from floating point scene coordinates to view coordinates.

See also matrix(), setTransform(), rotate(), scale(), shear(), and translate().
*/
func (this *QGraphicsView) SetMatrixp(matrix qtgui.QMatrix_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid, , Invalid
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetMatrix()

/*
Resets the view transformation matrix to the identity matrix.

See also resetTransform().
*/
func (this *QGraphicsView) ResetMatrix() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView11resetMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:169
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transform() const

/*
Returns the current transformation matrix for the view. If no current transformation is set, the identity matrix is returned.

See also setTransform(), rotate(), scale(), shear(), and translate().
*/
func (this *QGraphicsView) Transform() *qtgui.QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView9transformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:170
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform viewportTransform() const

/*
Returns a matrix that maps scene coordinates to viewport coordinates.

See also mapToScene() and mapFromScene().
*/
func (this *QGraphicsView) ViewportTransform() *qtgui.QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView17viewportTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:171
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTransformed() const

/*
Returns true if the view is transformed (i.e., a non-identity transform has been assigned, or the scrollbars are adjusted).

This function was introduced in  Qt 4.6.

See also setTransform(), horizontalScrollBar(), and verticalScrollBar().
*/
func (this *QGraphicsView) IsTransformed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView13isTransformedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, bool)

/*
Sets the view's current transformation matrix to matrix.

If combine is true, then matrix is combined with the current matrix; otherwise, matrix replaces the current matrix. combine is false by default.

The transformation matrix tranforms the scene into view coordinates. Using the default transformation, provided by the identity matrix, one pixel in the view represents one unit in the scene (e.g., a 10x10 rectangular item is drawn using 10x10 pixels in the view). If a 2x2 scaling matrix is applied, the scene will be drawn in 1:2 (e.g., a 10x10 rectangular item is then drawn using 20x20 pixels in the view).

Example:


  QGraphicsScene scene;
  scene.addText("GraphicsView rotated clockwise");

  QGraphicsView view(&scene);
  view.rotate(90); // the text is rendered with a 90 degree clockwise rotation
  view.show();



To simplify interation with items using a transformed view, QGraphicsView provides mapTo... and mapFrom... functions that can translate between scene and view coordinates. For example, you can call mapToScene() to map a view coordiate to a floating point scene coordinate, or mapFromScene() to map from floating point scene coordinates to view coordinates.

See also transform(), rotate(), scale(), shear(), and translate().
*/
func (this *QGraphicsView) SetTransform(matrix qtgui.QTransform_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, bool)

/*
Sets the view's current transformation matrix to matrix.

If combine is true, then matrix is combined with the current matrix; otherwise, matrix replaces the current matrix. combine is false by default.

The transformation matrix tranforms the scene into view coordinates. Using the default transformation, provided by the identity matrix, one pixel in the view represents one unit in the scene (e.g., a 10x10 rectangular item is drawn using 10x10 pixels in the view). If a 2x2 scaling matrix is applied, the scene will be drawn in 1:2 (e.g., a 10x10 rectangular item is then drawn using 20x20 pixels in the view).

Example:


  QGraphicsScene scene;
  scene.addText("GraphicsView rotated clockwise");

  QGraphicsView view(&scene);
  view.rotate(90); // the text is rendered with a 90 degree clockwise rotation
  view.show();



To simplify interation with items using a transformed view, QGraphicsView provides mapTo... and mapFrom... functions that can translate between scene and view coordinates. For example, you can call mapToScene() to map a view coordiate to a floating point scene coordinate, or mapFromScene() to map from floating point scene coordinates to view coordinates.

See also transform(), rotate(), scale(), shear(), and translate().
*/
func (this *QGraphicsView) SetTransformp(matrix qtgui.QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid, , Invalid
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetTransform()

/*
Resets the view transformation to the identity matrix.

See also transform() and setTransform().
*/
func (this *QGraphicsView) ResetTransform() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14resetTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotate(qreal)

/*
Rotates the current view transformation angle degrees clockwise.

See also setTransform(), transform(), scale(), shear(), and translate().
*/
func (this *QGraphicsView) Rotate(angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6rotateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scale(qreal, qreal)

/*
Scales the current view transformation by (sx, sy).

See also setTransform(), transform(), rotate(), shear(), and translate().
*/
func (this *QGraphicsView) Scale(sx float64, sy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView5scaleEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void shear(qreal, qreal)

/*
Shears the current view transformation by (sh, sv).

See also setTransform(), transform(), rotate(), scale(), and translate().
*/
func (this *QGraphicsView) Shear(sh float64, sv float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView5shearEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)

/*
Translates the current view transformation by (dx, dy).

See also setTransform(), transform(), rotate(), and shear().
*/
func (this *QGraphicsView) Translate(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void centerOn(const QPointF &)

/*
Scrolls the contents of the viewport to ensure that the scene coordinate pos, is centered in the view.

Because pos is a floating point coordinate, and the scroll bars operate on integer coordinates, the centering is only an approximation.

Note: If the item is close to or outside the border, it will be visible in the view, but not centered.

See also ensureVisible().
*/
func (this *QGraphicsView) CenterOn(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:180
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void centerOn(qreal, qreal)

/*
Scrolls the contents of the viewport to ensure that the scene coordinate pos, is centered in the view.

Because pos is a floating point coordinate, and the scroll bars operate on integer coordinates, the centering is only an approximation.

Note: If the item is close to or outside the border, it will be visible in the view, but not centered.

See also ensureVisible().
*/
func (this *QGraphicsView) CenterOn1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:181
// index:2
// Public Visibility=Default Availability=Available
// [-2] void centerOn(const QGraphicsItem *)

/*
Scrolls the contents of the viewport to ensure that the scene coordinate pos, is centered in the view.

Because pos is a floating point coordinate, and the scroll bars operate on integer coordinates, the centering is only an approximation.

Note: If the item is close to or outside the border, it will be visible in the view, but not centered.

See also ensureVisible().
*/
func (this *QGraphicsView) CenterOn2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisible(rect qtcore.QRectF_ITF, xmargin int, ymargin int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisiblep(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisiblep1(rect qtcore.QRectF_ITF, xmargin int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisible1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisible1p(x float64, y float64, w float64, h float64) {
	// arg: 4, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 5, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisible1p1(x float64, y float64, w float64, h float64, xmargin int) {
	// arg: 5, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QGraphicsItem *, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisible2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, xmargin int, ymargin int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QGraphicsItem *, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisible2p(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QGraphicsItem *, int, int)

/*
Scrolls the contents of the viewport so that the scene rectangle rect is visible, with margins specified in pixels by xmargin and ymargin. If the specified rect cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

See also centerOn().
*/
func (this *QGraphicsView) EnsureVisible2p1(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, xmargin int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fitInView(const QRectF &, Qt::AspectRatioMode)

/*
Scales the view matrix and scrolls the scroll bars to ensure that the scene rectangle rect fits inside the viewport. rect must be inside the scene rect; otherwise, fitInView() cannot guarantee that the whole rect is visible.

This function keeps the view's rotation, translation, or shear. The view is scaled according to aspectRatioMode. rect will be centered in the view if it does not fit tightly.

It's common to call fitInView() from inside a reimplementation of resizeEvent(), to ensure that the whole scene, or parts of the scene, scales automatically to fit the new size of the viewport as the view is resized. Note though, that calling fitInView() from inside resizeEvent() can lead to unwanted resize recursion, if the new transformation toggles the automatic state of the scrollbars. You can toggle the scrollbar policies to always on or always off to prevent this (see horizontalScrollBarPolicy() and verticalScrollBarPolicy()).

If rect is empty, or if the viewport is too small, this function will do nothing.

See also setTransform(), ensureVisible(), and centerOn().
*/
func (this *QGraphicsView) FitInView(rect qtcore.QRectF_ITF, aspectRadioMode int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewERK6QRectFN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fitInView(const QRectF &, Qt::AspectRatioMode)

/*
Scales the view matrix and scrolls the scroll bars to ensure that the scene rectangle rect fits inside the viewport. rect must be inside the scene rect; otherwise, fitInView() cannot guarantee that the whole rect is visible.

This function keeps the view's rotation, translation, or shear. The view is scaled according to aspectRatioMode. rect will be centered in the view if it does not fit tightly.

It's common to call fitInView() from inside a reimplementation of resizeEvent(), to ensure that the whole scene, or parts of the scene, scales automatically to fit the new size of the viewport as the view is resized. Note though, that calling fitInView() from inside resizeEvent() can lead to unwanted resize recursion, if the new transformation toggles the automatic state of the scrollbars. You can toggle the scrollbar policies to always on or always off to prevent this (see horizontalScrollBarPolicy() and verticalScrollBarPolicy()).

If rect is empty, or if the viewport is too small, this function will do nothing.

See also setTransform(), ensureVisible(), and centerOn().
*/
func (this *QGraphicsView) FitInViewp(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRadioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewERK6QRectFN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:186
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void fitInView(qreal, qreal, qreal, qreal, Qt::AspectRatioMode)

/*
Scales the view matrix and scrolls the scroll bars to ensure that the scene rectangle rect fits inside the viewport. rect must be inside the scene rect; otherwise, fitInView() cannot guarantee that the whole rect is visible.

This function keeps the view's rotation, translation, or shear. The view is scaled according to aspectRatioMode. rect will be centered in the view if it does not fit tightly.

It's common to call fitInView() from inside a reimplementation of resizeEvent(), to ensure that the whole scene, or parts of the scene, scales automatically to fit the new size of the viewport as the view is resized. Note though, that calling fitInView() from inside resizeEvent() can lead to unwanted resize recursion, if the new transformation toggles the automatic state of the scrollbars. You can toggle the scrollbar policies to always on or always off to prevent this (see horizontalScrollBarPolicy() and verticalScrollBarPolicy()).

If rect is empty, or if the viewport is too small, this function will do nothing.

See also setTransform(), ensureVisible(), and centerOn().
*/
func (this *QGraphicsView) FitInView1(x float64, y float64, w float64, h float64, aspectRadioMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEddddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:186
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void fitInView(qreal, qreal, qreal, qreal, Qt::AspectRatioMode)

/*
Scales the view matrix and scrolls the scroll bars to ensure that the scene rectangle rect fits inside the viewport. rect must be inside the scene rect; otherwise, fitInView() cannot guarantee that the whole rect is visible.

This function keeps the view's rotation, translation, or shear. The view is scaled according to aspectRatioMode. rect will be centered in the view if it does not fit tightly.

It's common to call fitInView() from inside a reimplementation of resizeEvent(), to ensure that the whole scene, or parts of the scene, scales automatically to fit the new size of the viewport as the view is resized. Note though, that calling fitInView() from inside resizeEvent() can lead to unwanted resize recursion, if the new transformation toggles the automatic state of the scrollbars. You can toggle the scrollbar policies to always on or always off to prevent this (see horizontalScrollBarPolicy() and verticalScrollBarPolicy()).

If rect is empty, or if the viewport is too small, this function will do nothing.

See also setTransform(), ensureVisible(), and centerOn().
*/
func (this *QGraphicsView) FitInView1p(x float64, y float64, w float64, h float64) {
	// arg: 4, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRadioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEddddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:188
// index:2
// Public Visibility=Default Availability=Available
// [-2] void fitInView(const QGraphicsItem *, Qt::AspectRatioMode)

/*
Scales the view matrix and scrolls the scroll bars to ensure that the scene rectangle rect fits inside the viewport. rect must be inside the scene rect; otherwise, fitInView() cannot guarantee that the whole rect is visible.

This function keeps the view's rotation, translation, or shear. The view is scaled according to aspectRatioMode. rect will be centered in the view if it does not fit tightly.

It's common to call fitInView() from inside a reimplementation of resizeEvent(), to ensure that the whole scene, or parts of the scene, scales automatically to fit the new size of the viewport as the view is resized. Note though, that calling fitInView() from inside resizeEvent() can lead to unwanted resize recursion, if the new transformation toggles the automatic state of the scrollbars. You can toggle the scrollbar policies to always on or always off to prevent this (see horizontalScrollBarPolicy() and verticalScrollBarPolicy()).

If rect is empty, or if the viewport is too small, this function will do nothing.

See also setTransform(), ensureVisible(), and centerOn().
*/
func (this *QGraphicsView) FitInView2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, aspectRadioMode int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEPK13QGraphicsItemN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:188
// index:2
// Public Visibility=Default Availability=Available
// [-2] void fitInView(const QGraphicsItem *, Qt::AspectRatioMode)

/*
Scales the view matrix and scrolls the scroll bars to ensure that the scene rectangle rect fits inside the viewport. rect must be inside the scene rect; otherwise, fitInView() cannot guarantee that the whole rect is visible.

This function keeps the view's rotation, translation, or shear. The view is scaled according to aspectRatioMode. rect will be centered in the view if it does not fit tightly.

It's common to call fitInView() from inside a reimplementation of resizeEvent(), to ensure that the whole scene, or parts of the scene, scales automatically to fit the new size of the viewport as the view is resized. Note though, that calling fitInView() from inside resizeEvent() can lead to unwanted resize recursion, if the new transformation toggles the automatic state of the scrollbars. You can toggle the scrollbar policies to always on or always off to prevent this (see horizontalScrollBarPolicy() and verticalScrollBarPolicy()).

If rect is empty, or if the viewport is too small, this function will do nothing.

See also setTransform(), ensureVisible(), and centerOn().
*/
func (this *QGraphicsView) FitInView2p(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRadioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEPK13QGraphicsItemN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)

/*
Renders the source rect, which is in view coordinates, from the scene into target, which is in paint device coordinates, using painter. This function is useful for capturing the contents of the view onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing to QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...

  QGraphicsView view(&scene);
  view.show();
  ...

  QPrinter printer(QPrinter::HighResolution);
  printer.setPageSize(QPrinter::A4);
  QPainter painter(&printer);

  // print, fitting the viewport contents into a full page
  view.render(&painter);

  // print the upper half of the viewport into the lower.
  // half of the page.
  QRect viewport = view.viewport()->rect();
  view.render(&painter,
              QRectF(0, printer.height() / 2,
                     printer.width(), printer.height() / 2),
              viewport.adjusted(0, 0, 0, -viewport.height() / 2));



If source is a null rect, this function will use viewport()->rect() to determine what to draw. If target is a null rect, the full dimensions of painter's paint device (e.g., for a QPrinter, the page size) will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsScene::render().
*/
func (this *QGraphicsView) Render(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF, source qtcore.QRect_ITF, aspectRatioMode int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRectF_PTR() != nil {
		convArg1 = target.QRectF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if source != nil && source.QRect_PTR() != nil {
		convArg2 = source.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)

/*
Renders the source rect, which is in view coordinates, from the scene into target, which is in paint device coordinates, using painter. This function is useful for capturing the contents of the view onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing to QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...

  QGraphicsView view(&scene);
  view.show();
  ...

  QPrinter printer(QPrinter::HighResolution);
  printer.setPageSize(QPrinter::A4);
  QPainter painter(&printer);

  // print, fitting the viewport contents into a full page
  view.render(&painter);

  // print the upper half of the viewport into the lower.
  // half of the page.
  QRect viewport = view.viewport()->rect();
  view.render(&painter,
              QRectF(0, printer.height() / 2,
                     printer.width(), printer.height() / 2),
              viewport.adjusted(0, 0, 0, -viewport.height() / 2));



If source is a null rect, this function will use viewport()->rect() to determine what to draw. If target is a null rect, the full dimensions of painter's paint device (e.g., for a QPrinter, the page size) will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsScene::render().
*/
func (this *QGraphicsView) Renderp(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 1, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QRect &=LValueReference, QRect=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)

/*
Renders the source rect, which is in view coordinates, from the scene into target, which is in paint device coordinates, using painter. This function is useful for capturing the contents of the view onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing to QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...

  QGraphicsView view(&scene);
  view.show();
  ...

  QPrinter printer(QPrinter::HighResolution);
  printer.setPageSize(QPrinter::A4);
  QPainter painter(&printer);

  // print, fitting the viewport contents into a full page
  view.render(&painter);

  // print the upper half of the viewport into the lower.
  // half of the page.
  QRect viewport = view.viewport()->rect();
  view.render(&painter,
              QRectF(0, printer.height() / 2,
                     printer.width(), printer.height() / 2),
              viewport.adjusted(0, 0, 0, -viewport.height() / 2));



If source is a null rect, this function will use viewport()->rect() to determine what to draw. If target is a null rect, the full dimensions of painter's paint device (e.g., for a QPrinter, the page size) will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsScene::render().
*/
func (this *QGraphicsView) Renderp1(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRectF_PTR() != nil {
		convArg1 = target.QRectF_PTR().GetCthis()
	}
	// arg: 2, const QRect &=LValueReference, QRect=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)

/*
Renders the source rect, which is in view coordinates, from the scene into target, which is in paint device coordinates, using painter. This function is useful for capturing the contents of the view onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing to QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...

  QGraphicsView view(&scene);
  view.show();
  ...

  QPrinter printer(QPrinter::HighResolution);
  printer.setPageSize(QPrinter::A4);
  QPainter painter(&printer);

  // print, fitting the viewport contents into a full page
  view.render(&painter);

  // print the upper half of the viewport into the lower.
  // half of the page.
  QRect viewport = view.viewport()->rect();
  view.render(&painter,
              QRectF(0, printer.height() / 2,
                     printer.width(), printer.height() / 2),
              viewport.adjusted(0, 0, 0, -viewport.height() / 2));



If source is a null rect, this function will use viewport()->rect() to determine what to draw. If target is a null rect, the full dimensions of painter's paint device (e.g., for a QPrinter, the page size) will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsScene::render().
*/
func (this *QGraphicsView) Renderp2(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF, source qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRectF_PTR() != nil {
		convArg1 = target.QRectF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if source != nil && source.QRect_PTR() != nil {
		convArg2 = source.QRect_PTR().GetCthis()
	}
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:194
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items() const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items() *QGraphicsItemList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:195
// index:1
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(const QPoint &) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items1(pos qtcore.QPoint_ITF) *QGraphicsItemList /*lll*/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:196
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(int, int) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items2(x int, y int) *QGraphicsItemList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:197
// index:3
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(const QRect &, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items3(rect qtcore.QRect_ITF, mode int) *QGraphicsItemList /*lll*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK5QRectN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:197
// index:3
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(const QRect &, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items3p(rect qtcore.QRect_ITF) *QGraphicsItemList /*lll*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	// arg: 1, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK5QRectN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:198
// index:4
// Public inline Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(int, int, int, int, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items4(x int, y int, w int, h int, mode int) *QGraphicsItemList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEiiiiN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:198
// index:4
// Public inline Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(int, int, int, int, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items4p(x int, y int, w int, h int) *QGraphicsItemList /*lll*/ {
	// arg: 4, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEiiiiN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:199
// index:5
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(const QPolygon &, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items5(polygon qtgui.QPolygon_ITF, mode int) *QGraphicsItemList /*lll*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygon_PTR() != nil {
		convArg0 = polygon.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK8QPolygonN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:199
// index:5
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(const QPolygon &, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items5p(polygon qtgui.QPolygon_ITF) *QGraphicsItemList /*lll*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygon_PTR() != nil {
		convArg0 = polygon.QPolygon_PTR().GetCthis()
	}
	// arg: 1, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK8QPolygonN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:200
// index:6
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(const QPainterPath &, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items6(path qtgui.QPainterPath_ITF, mode int) *QGraphicsItemList /*lll*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK12QPainterPathN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:200
// index:6
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> items(const QPainterPath &, Qt::ItemSelectionMode) const

/*
Returns a list of all the items in the associated scene, in descending stacking order (i.e., the first item in the returned list is the uppermost item).

See also QGraphicsScene::items() and Sorting.
*/
func (this *QGraphicsView) Items6p(path qtgui.QPainterPath_ITF) *QGraphicsItemList /*lll*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 1, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK12QPainterPathN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * itemAt(const QPoint &) const

/*
Returns the item at position pos, which is in viewport coordinates. If there are several items at this position, this function returns the topmost item.

Example:


  void CustomView::mousePressEvent(QMouseEvent *event)
  {
      if (QGraphicsItem *item = itemAt(event->pos())) {
          qDebug() << "You clicked on item" << item;
      } else {
          qDebug("You didn't click on an item.");
      }
  }



See also items() and Sorting.
*/
func (this *QGraphicsView) ItemAt(pos qtcore.QPoint_ITF) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsItem * itemAt(int, int) const

/*
Returns the item at position pos, which is in viewport coordinates. If there are several items at this position, this function returns the topmost item.

Example:


  void CustomView::mousePressEvent(QMouseEvent *event)
  {
      if (QGraphicsItem *item = itemAt(event->pos())) {
          qDebug() << "You clicked on item" << item;
      } else {
          qDebug("You didn't click on an item.");
      }
  }



See also items() and Sorting.
*/
func (this *QGraphicsView) ItemAt1(x int, y int) *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:204
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToScene(const QPoint &) const

/*
Returns the viewport coordinate point mapped to scene coordinates.

Note: It can be useful to map the whole rectangle covered by the pixel at point instead of the point itself. To do this, you can call mapToScene(QRect(point, QSize(2, 2))).

See also mapFromScene().
*/
func (this *QGraphicsView) MapToScene(point qtcore.QPoint_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:205
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToScene(const QRect &) const

/*
Returns the viewport coordinate point mapped to scene coordinates.

Note: It can be useful to map the whole rectangle covered by the pixel at point instead of the point itself. To do this, you can call mapToScene(QRect(point, QSize(2, 2))).

See also mapFromScene().
*/
func (this *QGraphicsView) MapToScene1(rect qtcore.QRect_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:206
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToScene(const QPolygon &) const

/*
Returns the viewport coordinate point mapped to scene coordinates.

Note: It can be useful to map the whole rectangle covered by the pixel at point instead of the point itself. To do this, you can call mapToScene(QRect(point, QSize(2, 2))).

See also mapFromScene().
*/
func (this *QGraphicsView) MapToScene2(polygon qtgui.QPolygon_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygon_PTR() != nil {
		convArg0 = polygon.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:207
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapToScene(const QPainterPath &) const

/*
Returns the viewport coordinate point mapped to scene coordinates.

Note: It can be useful to map the whole rectangle covered by the pixel at point instead of the point itself. To do this, you can call mapToScene(QRect(point, QSize(2, 2))).

See also mapFromScene().
*/
func (this *QGraphicsView) MapToScene3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:212
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QPointF mapToScene(int, int) const

/*
Returns the viewport coordinate point mapped to scene coordinates.

Note: It can be useful to map the whole rectangle covered by the pixel at point instead of the point itself. To do this, you can call mapToScene(QRect(point, QSize(2, 2))).

See also mapFromScene().
*/
func (this *QGraphicsView) MapToScene4(x int, y int) *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:213
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF mapToScene(int, int, int, int) const

/*
Returns the viewport coordinate point mapped to scene coordinates.

Note: It can be useful to map the whole rectangle covered by the pixel at point instead of the point itself. To do this, you can call mapToScene(QRect(point, QSize(2, 2))).

See also mapFromScene().
*/
func (this *QGraphicsView) MapToScene5(x int, y int, w int, h int) *qtgui.QPolygonF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:208
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFromScene(const QPointF &) const

/*
Returns the scene coordinate point to viewport coordinates.

See also mapToScene().
*/
func (this *QGraphicsView) MapFromScene(point qtcore.QPointF_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:209
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygon mapFromScene(const QRectF &) const

/*
Returns the scene coordinate point to viewport coordinates.

See also mapToScene().
*/
func (this *QGraphicsView) MapFromScene1(rect qtcore.QRectF_ITF) *qtgui.QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:210
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygon mapFromScene(const QPolygonF &) const

/*
Returns the scene coordinate point to viewport coordinates.

See also mapToScene().
*/
func (this *QGraphicsView) MapFromScene2(polygon qtgui.QPolygonF_ITF) *qtgui.QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:211
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapFromScene(const QPainterPath &) const

/*
Returns the scene coordinate point to viewport coordinates.

See also mapToScene().
*/
func (this *QGraphicsView) MapFromScene3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:214
// index:4
// Public inline Visibility=Default Availability=Available
// [8] QPoint mapFromScene(qreal, qreal) const

/*
Returns the scene coordinate point to viewport coordinates.

See also mapToScene().
*/
func (this *QGraphicsView) MapFromScene4(x float64, y float64) *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:215
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygon mapFromScene(qreal, qreal, qreal, qreal) const

/*
Returns the scene coordinate point to viewport coordinates.

See also mapToScene().
*/
func (this *QGraphicsView) MapFromScene5(x float64, y float64, w float64, h float64) *qtgui.QPolygon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:217
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QGraphicsView) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush backgroundBrush() const

/*

 */
func (this *QGraphicsView) BackgroundBrush() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView15backgroundBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundBrush(const QBrush &)

/*

 */
func (this *QGraphicsView) SetBackgroundBrush(brush qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView18setBackgroundBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush foregroundBrush() const

/*

 */
func (this *QGraphicsView) ForegroundBrush() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView15foregroundBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setForegroundBrush(const QBrush &)

/*

 */
func (this *QGraphicsView) SetForegroundBrush(brush qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView18setForegroundBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidateScene(const QRectF &, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of layers inside rect. rect is in scene coordinates. Any cached content for layers inside rect is unconditionally invalidated and redrawn.

You can call this function to notify QGraphicsView of changes to the background or the foreground of the scene. It is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled background caching.

Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but QGraphicsScene::BackgroundLayer is passed.

See also QGraphicsScene::invalidate() and update().
*/
func (this *QGraphicsView) InvalidateScene(rect qtcore.QRectF_ITF, layers int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15invalidateSceneERK6QRectF6QFlagsIN14QGraphicsScene10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidateScene(const QRectF &, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of layers inside rect. rect is in scene coordinates. Any cached content for layers inside rect is unconditionally invalidated and redrawn.

You can call this function to notify QGraphicsView of changes to the background or the foreground of the scene. It is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled background caching.

Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but QGraphicsScene::BackgroundLayer is passed.

See also QGraphicsScene::invalidate() and update().
*/
func (this *QGraphicsView) InvalidateScenep() {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, QGraphicsScene::SceneLayers=Elaborated, QGraphicsScene::SceneLayers=Typedef, QFlags<QGraphicsScene::SceneLayer>, Unexposed
	layers := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15invalidateSceneERK6QRectF6QFlagsIN14QGraphicsScene10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidateScene(const QRectF &, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of layers inside rect. rect is in scene coordinates. Any cached content for layers inside rect is unconditionally invalidated and redrawn.

You can call this function to notify QGraphicsView of changes to the background or the foreground of the scene. It is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled background caching.

Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but QGraphicsScene::BackgroundLayer is passed.

See also QGraphicsScene::invalidate() and update().
*/
func (this *QGraphicsView) InvalidateScenep1(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, QGraphicsScene::SceneLayers=Elaborated, QGraphicsScene::SceneLayers=Typedef, QFlags<QGraphicsScene::SceneLayer>, Unexposed
	layers := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15invalidateSceneERK6QRectF6QFlagsIN14QGraphicsScene10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateSceneRect(const QRectF &)

/*
Notifies QGraphicsView that the scene's scene rect has changed. rect is the new scene rect. If the view already has an explicitly set scene rect, this function does nothing.

See also sceneRect and QGraphicsScene::sceneRectChanged().
*/
func (this *QGraphicsView) UpdateSceneRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15updateSceneRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:232
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rubberBandChanged(QRect, QPointF, QPointF)

/*
This signal is emitted when the rubber band rect is changed. The viewport Rect is specified by rubberBandRect. The drag start position and drag end position are provided in scene points with fromScenePoint and toScenePoint.

When rubberband selection ends this signal will be emitted with null vales.

This function was introduced in  Qt 5.1.

See also rubberBandRect().
*/
func (this *QGraphicsView) RubberBandChanged(viewportRect qtcore.QRect_ITF /*123*/, fromScenePoint qtcore.QPointF_ITF /*123*/, toScenePoint qtcore.QPointF_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if viewportRect != nil && viewportRect.QRect_PTR() != nil {
		convArg0 = viewportRect.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if fromScenePoint != nil && fromScenePoint.QPointF_PTR() != nil {
		convArg1 = fromScenePoint.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if toScenePoint != nil && toScenePoint.QPointF_PTR() != nil {
		convArg2 = toScenePoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView17rubberBandChangedE5QRect7QPointFS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:236
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setupViewport(QWidget *)

/*
Reimplemented from QAbstractScrollArea::setupViewport().

This slot is called by QAbstractScrollArea after setViewport() has been called. Reimplement this function in a subclass of QGraphicsView to initialize the new viewport widget before it is used.

See also setViewport().
*/
func (this *QGraphicsView) SetupViewport(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13setupViewportEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:240
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QGraphicsView) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:241
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)

/*
Reimplemented from QAbstractScrollArea::viewportEvent().
*/
func (this *QGraphicsView) ViewportEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().
*/
func (this *QGraphicsView) ContextMenuEvent(event qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QContextMenuEvent_PTR() != nil {
		convArg0 = event.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:247
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QWidget::dragEnterEvent().
*/
func (this *QGraphicsView) DragEnterEvent(event qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragEnterEvent_PTR() != nil {
		convArg0 = event.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:248
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QWidget::dragLeaveEvent().
*/
func (this *QGraphicsView) DragLeaveEvent(event qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragLeaveEvent_PTR() != nil {
		convArg0 = event.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:249
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QWidget::dragMoveEvent().
*/
func (this *QGraphicsView) DragMoveEvent(event qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragMoveEvent_PTR() != nil {
		convArg0 = event.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:250
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QGraphicsView) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:252
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QGraphicsView) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:253
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QWidget::focusNextPrevChild().
*/
func (this *QGraphicsView) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:254
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QGraphicsView) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:255
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QGraphicsView) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:256
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyReleaseEvent().
*/
func (this *QGraphicsView) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:257
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseDoubleClickEvent().
*/
func (this *QGraphicsView) MouseDoubleClickEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:258
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QGraphicsView) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:259
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QGraphicsView) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:260
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QGraphicsView) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:262
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QGraphicsView) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:264
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QGraphicsView) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:265
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QGraphicsView) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:266
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*
Reimplemented from QAbstractScrollArea::scrollContentsBy().
*/
func (this *QGraphicsView) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:267
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QGraphicsView) ShowEvent(event qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QShowEvent_PTR() != nil {
		convArg0 = event.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:268
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
Reimplemented from QWidget::inputMethodEvent().
*/
func (this *QGraphicsView) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QInputMethodEvent_PTR() != nil {
		convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:270
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawBackground(QPainter *, const QRectF &)

/*
Draws the background of the scene using painter, before any items and the foreground are drawn. Reimplement this function to provide a custom background for this view.

If all you want is to define a color, texture or gradient for the background, you can call setBackgroundBrush() instead.

All painting is done in scene coordinates. rect is the exposed rectangle.

The default implementation fills rect using the view's backgroundBrush. If no such brush is defined (the default), the scene's drawBackground() function is called instead.

See also drawForeground() and QGraphicsScene::drawBackground().
*/
func (this *QGraphicsView) DrawBackground(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:271
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawForeground(QPainter *, const QRectF &)

/*
Draws the foreground of the scene using painter, after the background and all items are drawn. Reimplement this function to provide a custom foreground for this view.

If all you want is to define a color, texture or gradient for the foreground, you can call setForegroundBrush() instead.

All painting is done in scene coordinates. rect is the exposed rectangle.

The default implementation fills rect using the view's foregroundBrush. If no such brush is defined (the default), the scene's drawForeground() function is called instead.

See also drawBackground() and QGraphicsScene::drawForeground().
*/
func (this *QGraphicsView) DrawForeground(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:272
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawItems(QPainter *, int, QGraphicsItem **, const QStyleOptionGraphicsItem *)

/*

 */
func (this *QGraphicsView) DrawItems(painter qtgui.QPainter_ITF /*777 QPainter **/, numItems int, items []interface{}, options []interface{}) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numItems, items, options)
	qtrt.ErrPrint(err, rv)
}

/*
This enums describe the possible anchors that QGraphicsView can use when the user resizes the view or when the view is transformed.



See also resizeAnchor and transformationAnchor.

*/
type QGraphicsView__ViewportAnchor = int

// No anchor, i.e. the view leaves the scene's position unchanged.
const QGraphicsView__NoAnchor QGraphicsView__ViewportAnchor = 0

// The scene point at the center of the view is used as the anchor.
const QGraphicsView__AnchorViewCenter QGraphicsView__ViewportAnchor = 1

// The point under the mouse is used as the anchor.
const QGraphicsView__AnchorUnderMouse QGraphicsView__ViewportAnchor = 2

func (this *QGraphicsView) ViewportAnchorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsView_ViewportAnchorItemName(val int) string {
	var nilthis *QGraphicsView
	return nilthis.ViewportAnchorItemName(val)
}

/*


 */
type QGraphicsView__CacheModeFlag = int

//
const QGraphicsView__CacheNone QGraphicsView__CacheModeFlag = 0

//
const QGraphicsView__CacheBackground QGraphicsView__CacheModeFlag = 1

func (this *QGraphicsView) CacheModeFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsView_CacheModeFlagItemName(val int) string {
	var nilthis *QGraphicsView
	return nilthis.CacheModeFlagItemName(val)
}

/*
This enum describes the default action for the view when pressing and dragging the mouse over the viewport.



See also dragMode and QGraphicsScene::setSelectionArea().

*/
type QGraphicsView__DragMode = int

// Nothing happens; the mouse event is ignored.
const QGraphicsView__NoDrag QGraphicsView__DragMode = 0

// The cursor changes into a pointing hand, and dragging the mouse around will scroll the scrolbars. This mode works both in interactive and non-interactive mode.
const QGraphicsView__ScrollHandDrag QGraphicsView__DragMode = 1

// A rubber band will appear. Dragging the mouse will set the rubber band geometry, and all items covered by the rubber band are selected. This mode is disabled for non-interactive views.
const QGraphicsView__RubberBandDrag QGraphicsView__DragMode = 2

func (this *QGraphicsView) DragModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsView_DragModeItemName(val int) string {
	var nilthis *QGraphicsView
	return nilthis.DragModeItemName(val)
}

/*
This enum describes how QGraphicsView updates its viewport when the scene contents change or are exposed.



This enum was introduced or modified in  Qt 4.3.

See also viewportUpdateMode.

*/
type QGraphicsView__ViewportUpdateMode = int

// When any visible part of the scene changes or is reexposed, QGraphicsView will update the entire viewport. This approach is fastest when QGraphicsView spends more time figuring out what to draw than it would spend drawing (e.g., when very many small items are repeatedly updated). This is the preferred update mode for viewports that do not support partial updates, such as QGLWidget, and for viewports that need to disable scroll optimization.
const QGraphicsView__FullViewportUpdate QGraphicsView__ViewportUpdateMode = 0

// QGraphicsView will determine the minimal viewport region that requires a redraw, minimizing the time spent drawing by avoiding a redraw of areas that have not changed. This is QGraphicsView's default mode. Although this approach provides the best performance in general, if there are many small visible changes on the scene, QGraphicsView might end up spending more time finding the minimal approach than it will spend drawing.
const QGraphicsView__MinimalViewportUpdate QGraphicsView__ViewportUpdateMode = 1

// QGraphicsView will attempt to find an optimal update mode by analyzing the areas that require a redraw.
const QGraphicsView__SmartViewportUpdate QGraphicsView__ViewportUpdateMode = 2

// QGraphicsView will never update its viewport when the scene changes; the user is expected to control all updates. This mode disables all (potentially slow) item visibility testing in QGraphicsView, and is suitable for scenes that either require a fixed frame rate, or where the viewport is otherwise updated externally.
const QGraphicsView__NoViewportUpdate QGraphicsView__ViewportUpdateMode = 3

// The bounding rectangle of all changes in the viewport will be redrawn. This mode has the advantage that QGraphicsView searches only one region for changes, minimizing time spent determining what needs redrawing. The disadvantage is that areas that have not changed also need to be redrawn.
const QGraphicsView__BoundingRectViewportUpdate QGraphicsView__ViewportUpdateMode = 4

func (this *QGraphicsView) ViewportUpdateModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsView_ViewportUpdateModeItemName(val int) string {
	var nilthis *QGraphicsView
	return nilthis.ViewportUpdateModeItemName(val)
}

/*


 */
type QGraphicsView__OptimizationFlag = int

//
const QGraphicsView__DontClipPainter QGraphicsView__OptimizationFlag = 1

//
const QGraphicsView__DontSavePainterState QGraphicsView__OptimizationFlag = 2

//
const QGraphicsView__DontAdjustForAntialiasing QGraphicsView__OptimizationFlag = 4

//
const QGraphicsView__IndirectPainting QGraphicsView__OptimizationFlag = 8

func (this *QGraphicsView) OptimizationFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsView_OptimizationFlagItemName(val int) string {
	var nilthis *QGraphicsView
	return nilthis.OptimizationFlagItemName(val)
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
