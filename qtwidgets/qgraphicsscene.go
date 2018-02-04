package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsscene.h
// #include <qgraphicsscene.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
// bool event(class QEvent *)
func (this *QGraphicsScene) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QGraphicsScene) InheritEventFilter(f func(watched *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsScene) InheritContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/)) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDropEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsScene) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsScene) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void helpEvent(class QGraphicsSceneHelpEvent *)
func (this *QGraphicsScene) InheritHelpEvent(f func(event *QGraphicsSceneHelpEvent /*777 QGraphicsSceneHelpEvent **/)) {
	qtrt.SetAllInheritCallback(this, "helpEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsScene) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsScene) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMousePressEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void wheelEvent(class QGraphicsSceneWheelEvent *)
func (this *QGraphicsScene) InheritWheelEvent(f func(event *QGraphicsSceneWheelEvent /*777 QGraphicsSceneWheelEvent **/)) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsScene) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/)) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void drawBackground(class QPainter *, const class QRectF &)
func (this *QGraphicsScene) InheritDrawBackground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF)) {
	qtrt.SetAllInheritCallback(this, "drawBackground", f)
}

// void drawForeground(class QPainter *, const class QRectF &)
func (this *QGraphicsScene) InheritDrawForeground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF)) {
	qtrt.SetAllInheritCallback(this, "drawForeground", f)
}

// void drawItems(class QPainter *, int, class QGraphicsItem **, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsScene) InheritDrawItems(f func(painter *qtgui.QPainter /*777 QPainter **/, numItems int, items []interface{}, options []interface{}, widget *QWidget /*777 QWidget **/)) {
	qtrt.SetAllInheritCallback(this, "drawItems", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QGraphicsScene) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

type QGraphicsScene struct {
	*qtcore.QObject
}

func (this *QGraphicsScene) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGraphicsScene) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGraphicsSceneFromPointer(cthis unsafe.Pointer) *QGraphicsScene {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsScene{bcthis0}
}
func (*QGraphicsScene) NewFromPointer(cthis unsafe.Pointer) *QGraphicsScene {
	return NewQGraphicsSceneFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsScene) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(QObject *)
func NewQGraphicsScene(parent *qtcore.QObject /*777 QObject **/) *QGraphicsScene {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:125
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(const QRectF &, QObject *)
func NewQGraphicsScene_1(sceneRect *qtcore.QRectF, parent *qtcore.QObject /*777 QObject **/) *QGraphicsScene {
	var convArg0 = sceneRect.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2ERK6QRectFP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:126
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(qreal, qreal, qreal, qreal, QObject *)
func NewQGraphicsScene_2(x float64, y float64, width float64, height float64, parent *qtcore.QObject /*777 QObject **/) *QGraphicsScene {
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EddddP7QObject", qtrt.FFI_TYPE_POINTER, x, y, width, height, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsScene()
func DeleteQGraphicsScene(this *QGraphicsScene) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:129
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF sceneRect()
func (this *QGraphicsScene) SceneRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene9sceneRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal width()
func (this *QGraphicsScene) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:131
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height()
func (this *QGraphicsScene) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSceneRect(const QRectF &)
func (this *QGraphicsScene) SetSceneRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:133
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRectF &, Qt::AspectRatioMode)
func (this *QGraphicsScene) Render(painter *qtgui.QPainter /*777 QPainter **/, target *qtcore.QRectF, source *qtcore.QRectF, aspectRatioMode int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = target.GetCthis()
	var convArg2 = source.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsScene::ItemIndexMethod itemIndexMethod()
func (this *QGraphicsScene) ItemIndexMethod() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene15itemIndexMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemIndexMethod(enum QGraphicsScene::ItemIndexMethod)
func (this *QGraphicsScene) SetItemIndexMethod(method int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18setItemIndexMethodENS_15ItemIndexMethodE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), method)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortCacheEnabled()
func (this *QGraphicsScene) IsSortCacheEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene18isSortCacheEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortCacheEnabled(_Bool)
func (this *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene19setSortCacheEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] int bspTreeDepth()
func (this *QGraphicsScene) BspTreeDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene12bspTreeDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBspTreeDepth(int)
func (this *QGraphicsScene) SetBspTreeDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15setBspTreeDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF itemsBoundingRect()
func (this *QGraphicsScene) ItemsBoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene17itemsBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * itemAt(const QPointF &, const QTransform &)
func (this *QGraphicsScene) ItemAt(pos *qtcore.QPointF, deviceTransform *qtgui.QTransform) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg0 = pos.GetCthis()
	var convArg1 = deviceTransform.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:179
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsItem * itemAt(qreal, qreal, const QTransform &)
func (this *QGraphicsScene) ItemAt_1(x float64, y float64, deviceTransform *qtgui.QTransform) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg2 = deviceTransform.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtEddRK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath selectionArea()
func (this *QGraphicsScene) SelectionArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectionAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, const QTransform &)
func (this *QGraphicsScene) SetSelectionArea(path *qtgui.QPainterPath, deviceTransform *qtgui.QTransform) {
	var convArg0 = path.GetCthis()
	var convArg1 = deviceTransform.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionMode, const QTransform &)
func (this *QGraphicsScene) SetSelectionArea_1(path *qtgui.QPainterPath, mode int, deviceTransform *qtgui.QTransform) {
	var convArg0 = path.GetCthis()
	var convArg2 = deviceTransform.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:186
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const QTransform &)
func (this *QGraphicsScene) SetSelectionArea_2(path *qtgui.QPainterPath, selectionOperation int, mode int, deviceTransform *qtgui.QTransform) {
	var convArg0 = path.GetCthis()
	var convArg3 = deviceTransform.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, selectionOperation, mode, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroyItemGroup(QGraphicsItemGroup *)
func (this *QGraphicsScene) DestroyItemGroup(group *QGraphicsItemGroup /*777 QGraphicsItemGroup **/) {
	var convArg0 = group.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsItem *)
func (this *QGraphicsScene) AddItem(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addItemEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(const QRectF &, const QPen &, const QBrush &)
func (this *QGraphicsScene) AddEllipse(rect *qtcore.QRectF, pen *qtgui.QPen, brush *qtgui.QBrush) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	var convArg0 = rect.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)
func (this *QGraphicsScene) AddEllipse_1(x float64, y float64, w float64, h float64, pen *qtgui.QPen, brush *qtgui.QBrush) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	var convArg4 = pen.GetCthis()
	var convArg5 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:194
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLineItem * addLine(const QLineF &, const QPen &)
func (this *QGraphicsScene) AddLine(line *qtcore.QLineF, pen *qtgui.QPen) *QGraphicsLineItem /*777 QGraphicsLineItem **/ {
	var convArg0 = line.GetCthis()
	var convArg1 = pen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLineItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const QPen &)
func (this *QGraphicsScene) AddLine_1(x1 float64, y1 float64, x2 float64, y2 float64, pen *qtgui.QPen) *QGraphicsLineItem /*777 QGraphicsLineItem **/ {
	var convArg4 = pen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineEddddRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLineItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:195
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPathItem * addPath(const QPainterPath &, const QPen &, const QBrush &)
func (this *QGraphicsScene) AddPath(path *qtgui.QPainterPath, pen *qtgui.QPen, brush *qtgui.QBrush) *QGraphicsPathItem /*777 QGraphicsPathItem **/ {
	var convArg0 = path.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsPathItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPixmapItem * addPixmap(const QPixmap &)
func (this *QGraphicsScene) AddPixmap(pixmap *qtgui.QPixmap) *QGraphicsPixmapItem /*777 QGraphicsPixmapItem **/ {
	var convArg0 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9addPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsPixmapItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:197
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPolygonItem * addPolygon(const QPolygonF &, const QPen &, const QBrush &)
func (this *QGraphicsScene) AddPolygon(polygon *qtgui.QPolygonF, pen *qtgui.QPen, brush *qtgui.QBrush) *QGraphicsPolygonItem /*777 QGraphicsPolygonItem **/ {
	var convArg0 = polygon.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(const QRectF &, const QPen &, const QBrush &)
func (this *QGraphicsScene) AddRect(rect *qtcore.QRectF, pen *qtgui.QPen, brush *qtgui.QBrush) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	var convArg0 = rect.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:206
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)
func (this *QGraphicsScene) AddRect_1(x float64, y float64, w float64, h float64, pen *qtgui.QPen, brush *qtgui.QBrush) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	var convArg4 = pen.GetCthis()
	var convArg5 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsTextItem * addText(const QString &, const QFont &)
func (this *QGraphicsScene) AddText(text *qtcore.QString, font *qtgui.QFont) *QGraphicsTextItem /*777 QGraphicsTextItem **/ {
	var convArg0 = text.GetCthis()
	var convArg1 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addTextERK7QStringRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsSimpleTextItem * addSimpleText(const QString &, const QFont &)
func (this *QGraphicsScene) AddSimpleText(text *qtcore.QString, font *qtgui.QFont) *QGraphicsSimpleTextItem /*777 QGraphicsSimpleTextItem **/ {
	var convArg0 = text.GetCthis()
	var convArg1 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsSimpleTextItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsProxyWidget * addWidget(QWidget *, Qt::WindowFlags)
func (this *QGraphicsScene) AddWidget(widget *QWidget /*777 QWidget **/, wFlags int) *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9addWidgetEP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, wFlags)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QGraphicsItem *)
func (this *QGraphicsScene) RemoveItem(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10removeItemEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:210
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * focusItem()
func (this *QGraphicsScene) FocusItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene9focusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusItem(QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) SetFocusItem(item *QGraphicsItem /*777 QGraphicsItem **/, focusReason int) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12setFocusItemEP13QGraphicsItemN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:212
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus()
func (this *QGraphicsScene) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)
func (this *QGraphicsScene) SetFocus(focusReason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFocus()
func (this *QGraphicsScene) ClearFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10clearFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStickyFocus(_Bool)
func (this *QGraphicsScene) SetStickyFocus(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14setStickyFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:217
// index:0
// Public Visibility=Default Availability=Available
// [1] bool stickyFocus()
func (this *QGraphicsScene) StickyFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene11stickyFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * mouseGrabberItem()
func (this *QGraphicsScene) MouseGrabberItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene16mouseGrabberItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush backgroundBrush()
func (this *QGraphicsScene) BackgroundBrush() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene15backgroundBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundBrush(const QBrush &)
func (this *QGraphicsScene) SetBackgroundBrush(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18setBackgroundBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush foregroundBrush()
func (this *QGraphicsScene) ForegroundBrush() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene15foregroundBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setForegroundBrush(const QBrush &)
func (this *QGraphicsScene) SetForegroundBrush(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18setForegroundBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:227
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsScene) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void update(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) Update(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6updateEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:257
// index:1
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)
func (this *QGraphicsScene) Update_1(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void invalidate(qreal, qreal, qreal, qreal, QGraphicsScene::SceneLayers)
func (this *QGraphicsScene) Invalidate(x float64, y float64, w float64, h float64, layers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateEdddd6QFlagsINS_10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, layers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:258
// index:1
// Public Visibility=Default Availability=Available
// [-2] void invalidate(const QRectF &, QGraphicsScene::SceneLayers)
func (this *QGraphicsScene) Invalidate_1(rect *qtcore.QRectF, layers int) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateERK6QRectF6QFlagsINS_10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * style()
func (this *QGraphicsScene) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)
func (this *QGraphicsScene) SetStyle(style *QStyle /*777 QStyle **/) {
	var convArg0 = style.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene8setStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:239
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QGraphicsScene) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QGraphicsScene) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:242
// index:0
// Public Visibility=Default Availability=Available
// [16] QPalette palette()
func (this *QGraphicsScene) Palette() *qtgui.QPalette /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene7paletteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &)
func (this *QGraphicsScene) SetPalette(palette *qtgui.QPalette) {
	var convArg0 = palette.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:245
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QGraphicsScene) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:246
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * activePanel()
func (this *QGraphicsScene) ActivePanel() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene11activePanelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActivePanel(QGraphicsItem *)
func (this *QGraphicsScene) SetActivePanel(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:248
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * activeWindow()
func (this *QGraphicsScene) ActiveWindow() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene12activeWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveWindow(QGraphicsWidget *)
func (this *QGraphicsScene) SetActiveWindow(widget *QGraphicsWidget /*777 QGraphicsWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:251
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sendEvent(QGraphicsItem *, QEvent *)
func (this *QGraphicsScene) SendEvent(item *QGraphicsItem /*777 QGraphicsItem **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = item.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:253
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minimumRenderSize()
func (this *QGraphicsScene) MinimumRenderSize() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene17minimumRenderSizeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:254
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumRenderSize(qreal)
func (this *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene20setMinimumRenderSizeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void advance()
func (this *QGraphicsScene) Advance() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7advanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSelection()
func (this *QGraphicsScene) ClearSelection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14clearSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QGraphicsScene) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:264
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QGraphicsScene) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:265
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QGraphicsScene) EventFilter(watched *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = watched.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:266
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsScene) ContextMenuEvent(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:267
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragEnterEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:268
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragMoveEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:269
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragLeaveEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:270
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DropEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:271
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QGraphicsScene) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:272
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QGraphicsScene) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:273
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void helpEvent(QGraphicsSceneHelpEvent *)
func (this *QGraphicsScene) HelpEvent(event *QGraphicsSceneHelpEvent /*777 QGraphicsSceneHelpEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:274
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QGraphicsScene) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:275
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsScene) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:276
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MousePressEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:277
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseMoveEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:278
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseReleaseEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:279
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseDoubleClickEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:280
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QGraphicsSceneWheelEvent *)
func (this *QGraphicsScene) WheelEvent(event *QGraphicsSceneWheelEvent /*777 QGraphicsSceneWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:281
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsScene) InputMethodEvent(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:283
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawBackground(QPainter *, const QRectF &)
func (this *QGraphicsScene) DrawBackground(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:284
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawForeground(QPainter *, const QRectF &)
func (this *QGraphicsScene) DrawForeground(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:285
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawItems(QPainter *, int, QGraphicsItem **, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsScene) DrawItems(painter *qtgui.QPainter /*777 QPainter **/, numItems int, items []interface{}, options []interface{}, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg4 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numItems, items, options, convArg4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:295
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QGraphicsScene) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneRectChanged(const QRectF &)
func (this *QGraphicsScene) SceneRectChanged(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16sceneRectChangedERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()
func (this *QGraphicsScene) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusItemChanged(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) FocusItemChanged(newFocus *QGraphicsItem /*777 QGraphicsItem **/, oldFocus *QGraphicsItem /*777 QGraphicsItem **/, reason int) {
	var convArg0 = newFocus.GetCthis()
	var convArg1 = oldFocus.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16focusItemChangedEP13QGraphicsItemS1_N2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, reason)
	gopp.ErrPrint(err, rv)
}

type QGraphicsScene__ItemIndexMethod = int

const QGraphicsScene__BspTreeIndex QGraphicsScene__ItemIndexMethod = 0
const QGraphicsScene__NoIndex QGraphicsScene__ItemIndexMethod = -1

type QGraphicsScene__SceneLayer = int

const QGraphicsScene__ItemLayer QGraphicsScene__SceneLayer = 1
const QGraphicsScene__BackgroundLayer QGraphicsScene__SceneLayer = 2
const QGraphicsScene__ForegroundLayer QGraphicsScene__SceneLayer = 4
const QGraphicsScene__AllLayers QGraphicsScene__SceneLayer = 65535

//  body block end
