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

// void setupViewport(class QWidget *)
func (this *QGraphicsView) InheritSetupViewport(f func(widget *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setupViewport", f)
}

// bool event(class QEvent *)
func (this *QGraphicsView) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool viewportEvent(class QEvent *)
func (this *QGraphicsView) InheritViewportEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QGraphicsView) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(class QDragEnterEvent *)
func (this *QGraphicsView) InheritDragEnterEvent(f func(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QGraphicsView) InheritDragLeaveEvent(f func(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QGraphicsView) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QGraphicsView) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsView) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QGraphicsView) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsView) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsView) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsView) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QGraphicsView) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QGraphicsView) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QGraphicsView) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QGraphicsView) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QGraphicsView) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QGraphicsView) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QGraphicsView) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QGraphicsView) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void showEvent(class QShowEvent *)
func (this *QGraphicsView) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsView) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void drawBackground(class QPainter *, const class QRectF &)
func (this *QGraphicsView) InheritDrawBackground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawBackground", f)
}

// void drawForeground(class QPainter *, const class QRectF &)
func (this *QGraphicsView) InheritDrawForeground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawForeground", f)
}

// void drawItems(class QPainter *, int, class QGraphicsItem **, const class QStyleOptionGraphicsItem *)
func (this *QGraphicsView) InheritDrawItems(f func(painter *qtgui.QPainter /*777 QPainter **/, numItems int, items []interface{}, options []interface{}) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawItems", f)
}

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
func (this *QGraphicsView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsView(QWidget *)
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
func NewQGraphicsView__() *QGraphicsView {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
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
func NewQGraphicsView_1(scene QGraphicsScene_ITF /*777 QGraphicsScene **/, parent QWidget_ITF /*777 QWidget **/) *QGraphicsView {
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
func NewQGraphicsView_1_(scene QGraphicsScene_ITF /*777 QGraphicsScene **/) *QGraphicsView {
	var convArg0 unsafe.Pointer
	if scene != nil && scene.QGraphicsScene_PTR() != nil {
		convArg0 = scene.QGraphicsScene_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record,
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
func (this *QGraphicsView) RenderHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView11renderHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHint(QPainter::RenderHint, _Bool)
func (this *QGraphicsView) SetRenderHint(hint int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13setRenderHintEN8QPainter10RenderHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHint(QPainter::RenderHint, _Bool)
func (this *QGraphicsView) SetRenderHint__(hint int) {
	// arg: 1, bool=Bool, =Invalid,
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13setRenderHintEN8QPainter10RenderHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderHints(QPainter::RenderHints)
func (this *QGraphicsView) SetRenderHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14setRenderHintsE6QFlagsIN8QPainter10RenderHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const
func (this *QGraphicsView) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)
func (this *QGraphicsView) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::ViewportAnchor transformationAnchor() const
func (this *QGraphicsView) TransformationAnchor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView20transformationAnchorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformationAnchor(enum QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetTransformationAnchor(anchor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView23setTransformationAnchorENS_14ViewportAnchorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), anchor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::ViewportAnchor resizeAnchor() const
func (this *QGraphicsView) ResizeAnchor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView12resizeAnchorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeAnchor(enum QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetResizeAnchor(anchor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15setResizeAnchorENS_14ViewportAnchorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), anchor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::ViewportUpdateMode viewportUpdateMode() const
func (this *QGraphicsView) ViewportUpdateMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView18viewportUpdateModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewportUpdateMode(enum QGraphicsView::ViewportUpdateMode)
func (this *QGraphicsView) SetViewportUpdateMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView21setViewportUpdateModeENS_18ViewportUpdateModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::OptimizationFlags optimizationFlags() const
func (this *QGraphicsView) OptimizationFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView17optimizationFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizationFlag(enum QGraphicsView::OptimizationFlag, _Bool)
func (this *QGraphicsView) SetOptimizationFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView19setOptimizationFlagENS_16OptimizationFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizationFlag(enum QGraphicsView::OptimizationFlag, _Bool)
func (this *QGraphicsView) SetOptimizationFlag__(flag int) {
	// arg: 1, bool=Bool, =Invalid,
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView19setOptimizationFlagENS_16OptimizationFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizationFlags(QGraphicsView::OptimizationFlags)
func (this *QGraphicsView) SetOptimizationFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView20setOptimizationFlagsE6QFlagsINS_16OptimizationFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsView::DragMode dragMode() const
func (this *QGraphicsView) DragMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView8dragModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragMode(enum QGraphicsView::DragMode)
func (this *QGraphicsView) SetDragMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView11setDragModeENS_8DragModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemSelectionMode rubberBandSelectionMode() const
func (this *QGraphicsView) RubberBandSelectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView23rubberBandSelectionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRubberBandSelectionMode(Qt::ItemSelectionMode)
func (this *QGraphicsView) SetRubberBandSelectionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView26setRubberBandSelectionModeEN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:149
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rubberBandRect() const
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
func (this *QGraphicsView) CacheMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView9cacheModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheMode(QGraphicsView::CacheMode)
func (this *QGraphicsView) SetCacheMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setCacheModeE6QFlagsINS_13CacheModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetCachedContent()
func (this *QGraphicsView) ResetCachedContent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView18resetCachedContentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:156
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInteractive() const
func (this *QGraphicsView) IsInteractive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView13isInteractiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInteractive(_Bool)
func (this *QGraphicsView) SetInteractive(allowed bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14setInteractiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), allowed)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsScene * scene() const
func (this *QGraphicsView) Scene() *QGraphicsScene /*777 QGraphicsScene **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView5sceneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScene(QGraphicsScene *)
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
func (this *QGraphicsView) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:166
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix matrix() const
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
// [-2] void setMatrix(const QMatrix &, _Bool)
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
// [-2] void setMatrix(const QMatrix &, _Bool)
func (this *QGraphicsView) SetMatrix__(matrix qtgui.QMatrix_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetMatrix()
func (this *QGraphicsView) ResetMatrix() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView11resetMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:169
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transform() const
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
func (this *QGraphicsView) IsTransformed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView13isTransformedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, _Bool)
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
// [-2] void setTransform(const QTransform &, _Bool)
func (this *QGraphicsView) SetTransform__(matrix qtgui.QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetTransform()
func (this *QGraphicsView) ResetTransform() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14resetTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotate(qreal)
func (this *QGraphicsView) Rotate(angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6rotateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scale(qreal, qreal)
func (this *QGraphicsView) Scale(sx float64, sy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView5scaleEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void shear(qreal, qreal)
func (this *QGraphicsView) Shear(sh float64, sv float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView5shearEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)
func (this *QGraphicsView) Translate(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void centerOn(const QPointF &)
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
func (this *QGraphicsView) CenterOn_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:181
// index:2
// Public Visibility=Default Availability=Available
// [-2] void centerOn(const QGraphicsItem *)
func (this *QGraphicsView) CenterOn_2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) {
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
func (this *QGraphicsView) EnsureVisible__(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	xmargin := int(50)
	// arg: 2, int=Int, =Invalid,
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)
func (this *QGraphicsView) EnsureVisible__1(rect qtcore.QRectF_ITF, xmargin int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsView) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsView) EnsureVisible_1_(x float64, y float64, w float64, h float64) {
	// arg: 4, int=Int, =Invalid,
	xmargin := int(50)
	// arg: 5, int=Int, =Invalid,
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsView) EnsureVisible_1_1(x float64, y float64, w float64, h float64, xmargin int) {
	// arg: 5, int=Int, =Invalid,
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QGraphicsItem *, int, int)
func (this *QGraphicsView) EnsureVisible_2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, xmargin int, ymargin int) {
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
func (this *QGraphicsView) EnsureVisible_2_(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	xmargin := int(50)
	// arg: 2, int=Int, =Invalid,
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QGraphicsItem *, int, int)
func (this *QGraphicsView) EnsureVisible_2_1(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, xmargin int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fitInView(const QRectF &, Qt::AspectRatioMode)
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
func (this *QGraphicsView) FitInView__(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectRadioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewERK6QRectFN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:186
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void fitInView(qreal, qreal, qreal, qreal, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_1(x float64, y float64, w float64, h float64, aspectRadioMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEddddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:186
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void fitInView(qreal, qreal, qreal, qreal, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_1_(x float64, y float64, w float64, h float64) {
	// arg: 4, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectRadioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEddddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:188
// index:2
// Public Visibility=Default Availability=Available
// [-2] void fitInView(const QGraphicsItem *, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, aspectRadioMode int) {
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
func (this *QGraphicsView) FitInView_2_(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectRadioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEPK13QGraphicsItemN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)
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
func (this *QGraphicsView) Render__(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 1, const QRectF &=LValueReference, QRectF=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const QRect &=LValueReference, QRect=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)
func (this *QGraphicsView) Render__1(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRectF_PTR() != nil {
		convArg1 = target.QRectF_PTR().GetCthis()
	}
	// arg: 2, const QRect &=LValueReference, QRect=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)
func (this *QGraphicsView) Render__2(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF, source qtcore.QRect_ITF) {
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
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * itemAt(const QPoint &) const
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
func (this *QGraphicsView) ItemAt_1(x int, y int) *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:204
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToScene(const QPoint &) const
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
func (this *QGraphicsView) MapToScene_1(rect qtcore.QRect_ITF) *qtgui.QPolygonF /*123*/ {
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
func (this *QGraphicsView) MapToScene_2(polygon qtgui.QPolygon_ITF) *qtgui.QPolygonF /*123*/ {
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
func (this *QGraphicsView) MapToScene_3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
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
func (this *QGraphicsView) MapToScene_4(x int, y int) *qtcore.QPointF /*123*/ {
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
func (this *QGraphicsView) MapToScene_5(x int, y int, w int, h int) *qtgui.QPolygonF /*123*/ {
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
func (this *QGraphicsView) MapFromScene_1(rect qtcore.QRectF_ITF) *qtgui.QPolygon /*123*/ {
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
func (this *QGraphicsView) MapFromScene_2(polygon qtgui.QPolygonF_ITF) *qtgui.QPolygon /*123*/ {
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
func (this *QGraphicsView) MapFromScene_3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
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
func (this *QGraphicsView) MapFromScene_4(x float64, y float64) *qtcore.QPoint /*123*/ {
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
func (this *QGraphicsView) MapFromScene_5(x float64, y float64, w float64, h float64) *qtgui.QPolygon /*123*/ {
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
func (this *QGraphicsView) InvalidateScene__() {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record,
	var convArg0 unsafe.Pointer
	// arg: 1, QGraphicsScene::SceneLayers=Elaborated, QGraphicsScene::SceneLayers=Typedef, QFlags<QGraphicsScene::SceneLayer>
	layers := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15invalidateSceneERK6QRectF6QFlagsIN14QGraphicsScene10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidateScene(const QRectF &, QGraphicsScene::SceneLayers)
func (this *QGraphicsView) InvalidateScene__1(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, QGraphicsScene::SceneLayers=Elaborated, QGraphicsScene::SceneLayers=Typedef, QFlags<QGraphicsScene::SceneLayer>
	layers := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView15invalidateSceneERK6QRectF6QFlagsIN14QGraphicsScene10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateSceneRect(const QRectF &)
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
// [1] bool focusNextPrevChild(_Bool)
func (this *QGraphicsView) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:254
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
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
func (this *QGraphicsView) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:267
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
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
func (this *QGraphicsView) DrawItems(painter qtgui.QPainter_ITF /*777 QPainter **/, numItems int, items []interface{}, options []interface{}) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numItems, items, options)
	qtrt.ErrPrint(err, rv)
}

type QGraphicsView__ViewportAnchor = int

const QGraphicsView__NoAnchor QGraphicsView__ViewportAnchor = 0
const QGraphicsView__AnchorViewCenter QGraphicsView__ViewportAnchor = 1
const QGraphicsView__AnchorUnderMouse QGraphicsView__ViewportAnchor = 2

type QGraphicsView__CacheModeFlag = int

const QGraphicsView__CacheNone QGraphicsView__CacheModeFlag = 0
const QGraphicsView__CacheBackground QGraphicsView__CacheModeFlag = 1

type QGraphicsView__DragMode = int

const QGraphicsView__NoDrag QGraphicsView__DragMode = 0
const QGraphicsView__ScrollHandDrag QGraphicsView__DragMode = 1
const QGraphicsView__RubberBandDrag QGraphicsView__DragMode = 2

type QGraphicsView__ViewportUpdateMode = int

const QGraphicsView__FullViewportUpdate QGraphicsView__ViewportUpdateMode = 0
const QGraphicsView__MinimalViewportUpdate QGraphicsView__ViewportUpdateMode = 1
const QGraphicsView__SmartViewportUpdate QGraphicsView__ViewportUpdateMode = 2
const QGraphicsView__NoViewportUpdate QGraphicsView__ViewportUpdateMode = 3
const QGraphicsView__BoundingRectViewportUpdate QGraphicsView__ViewportUpdateMode = 4

type QGraphicsView__OptimizationFlag = int

const QGraphicsView__DontClipPainter QGraphicsView__OptimizationFlag = 1
const QGraphicsView__DontSavePainterState QGraphicsView__OptimizationFlag = 2
const QGraphicsView__DontAdjustForAntialiasing QGraphicsView__OptimizationFlag = 4
const QGraphicsView__IndirectPainting QGraphicsView__OptimizationFlag = 8

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
