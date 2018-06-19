package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsscene.h
// #include <qgraphicsscene.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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

// bool event(QEvent *)
func (this *QGraphicsScene) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QGraphicsScene) InheritEventFilter(f func(watched *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void contextMenuEvent(QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsScene) InheritContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) InheritDropEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QGraphicsScene) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QGraphicsScene) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void helpEvent(QGraphicsSceneHelpEvent *)
func (this *QGraphicsScene) InheritHelpEvent(f func(event *QGraphicsSceneHelpEvent /*777 QGraphicsSceneHelpEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "helpEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QGraphicsScene) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsScene) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMousePressEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) InheritMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void wheelEvent(QGraphicsSceneWheelEvent *)
func (this *QGraphicsScene) InheritWheelEvent(f func(event *QGraphicsSceneWheelEvent /*777 QGraphicsSceneWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsScene) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void drawBackground(QPainter *, const QRectF &)
func (this *QGraphicsScene) InheritDrawBackground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawBackground", f)
}

// void drawForeground(QPainter *, const QRectF &)
func (this *QGraphicsScene) InheritDrawForeground(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawForeground", f)
}

// void drawItems(QPainter *, int, QGraphicsItem **, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsScene) InheritDrawItems(f func(painter *qtgui.QPainter /*777 QPainter **/, numItems int, items []interface{}, options []interface{}, widget *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawItems", f)
}

// bool focusNextPrevChild(bool)
func (this *QGraphicsScene) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

/*

 */
type QGraphicsScene struct {
	*qtcore.QObject
}
type QGraphicsScene_ITF interface {
	qtcore.QObject_ITF
	QGraphicsScene_PTR() *QGraphicsScene
}

func (ptr *QGraphicsScene) QGraphicsScene_PTR() *QGraphicsScene { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsScene) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(QObject *)

/*
Constructs a QGraphicsScene object. The parent parameter is passed to QObject's constructor.
*/
func NewQGraphicsScene(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsScene {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsScene")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(QObject *)

/*
Constructs a QGraphicsScene object. The parent parameter is passed to QObject's constructor.
*/
func NewQGraphicsScene__() *QGraphicsScene {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsScene")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:125
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(const QRectF &, QObject *)

/*
Constructs a QGraphicsScene object. The parent parameter is passed to QObject's constructor.
*/
func NewQGraphicsScene_1(sceneRect qtcore.QRectF_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsScene {
	var convArg0 unsafe.Pointer
	if sceneRect != nil && sceneRect.QRectF_PTR() != nil {
		convArg0 = sceneRect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2ERK6QRectFP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsScene")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:125
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(const QRectF &, QObject *)

/*
Constructs a QGraphicsScene object. The parent parameter is passed to QObject's constructor.
*/
func NewQGraphicsScene_1_(sceneRect qtcore.QRectF_ITF) *QGraphicsScene {
	var convArg0 unsafe.Pointer
	if sceneRect != nil && sceneRect.QRectF_PTR() != nil {
		convArg0 = sceneRect.QRectF_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2ERK6QRectFP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsScene")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:126
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(qreal, qreal, qreal, qreal, QObject *)

/*
Constructs a QGraphicsScene object. The parent parameter is passed to QObject's constructor.
*/
func NewQGraphicsScene_2(x float64, y float64, width float64, height float64, parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsScene {
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg4 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EddddP7QObject", qtrt.FFI_TYPE_POINTER, x, y, width, height, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsScene")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:126
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScene(qreal, qreal, qreal, qreal, QObject *)

/*
Constructs a QGraphicsScene object. The parent parameter is passed to QObject's constructor.
*/
func NewQGraphicsScene_2_(x float64, y float64, width float64, height float64) *QGraphicsScene {
	// arg: 4, QObject *=Pointer, QObject=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EddddP7QObject", qtrt.FFI_TYPE_POINTER, x, y, width, height, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsScene")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsScene()

/*

 */
func DeleteQGraphicsScene(this *QGraphicsScene) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsSceneD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:129
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF sceneRect() const

/*

 */
func (this *QGraphicsScene) SceneRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene9sceneRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal width() const

/*
This convenience function is equivalent to calling sceneRect().width().

See also height().
*/
func (this *QGraphicsScene) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:131
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height() const

/*
This convenience function is equivalent to calling sceneRect().height().

See also width().
*/
func (this *QGraphicsScene) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSceneRect(const QRectF &)

/*

 */
func (this *QGraphicsScene) SetSceneRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:133
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setSceneRect(qreal, qreal, qreal, qreal)

/*

 */
func (this *QGraphicsScene) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRectF &, Qt::AspectRatioMode)

/*
Renders the source rect from scene into target, using painter. This function is useful for capturing the contents of the scene onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing with QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...
  QPrinter printer(QPrinter::HighResolution);
  printer.setPaperSize(QPrinter::A4);

  QPainter painter(&printer);
  scene.render(&painter);



If source is a null rect, this function will use sceneRect() to determine what to render. If target is a null rect, the dimensions of painter's paint device will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsView::render().
*/
func (this *QGraphicsScene) Render(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF, source qtcore.QRectF_ITF, aspectRatioMode int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRectF_PTR() != nil {
		convArg1 = target.QRectF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if source != nil && source.QRectF_PTR() != nil {
		convArg2 = source.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRectF &, Qt::AspectRatioMode)

/*
Renders the source rect from scene into target, using painter. This function is useful for capturing the contents of the scene onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing with QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...
  QPrinter printer(QPrinter::HighResolution);
  printer.setPaperSize(QPrinter::A4);

  QPainter painter(&printer);
  scene.render(&painter);



If source is a null rect, this function will use sceneRect() to determine what to render. If target is a null rect, the dimensions of painter's paint device will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsView::render().
*/
func (this *QGraphicsScene) Render__(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 1, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRectF &, Qt::AspectRatioMode)

/*
Renders the source rect from scene into target, using painter. This function is useful for capturing the contents of the scene onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing with QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...
  QPrinter printer(QPrinter::HighResolution);
  printer.setPaperSize(QPrinter::A4);

  QPainter painter(&printer);
  scene.render(&painter);



If source is a null rect, this function will use sceneRect() to determine what to render. If target is a null rect, the dimensions of painter's paint device will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsView::render().
*/
func (this *QGraphicsScene) Render__1(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRectF_PTR() != nil {
		convArg1 = target.QRectF_PTR().GetCthis()
	}
	// arg: 2, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QRectF &, const QRectF &, Qt::AspectRatioMode)

/*
Renders the source rect from scene into target, using painter. This function is useful for capturing the contents of the scene onto a paint device, such as a QImage (e.g., to take a screenshot), or for printing with QPrinter. For example:


  QGraphicsScene scene;
  scene.addItem(...
  ...
  QPrinter printer(QPrinter::HighResolution);
  printer.setPaperSize(QPrinter::A4);

  QPainter painter(&printer);
  scene.render(&painter);



If source is a null rect, this function will use sceneRect() to determine what to render. If target is a null rect, the dimensions of painter's paint device will be used.

The source rect contents will be transformed according to aspectRatioMode to fit into the target rect. By default, the aspect ratio is kept, and source is scaled to fit in target.

See also QGraphicsView::render().
*/
func (this *QGraphicsScene) Render__2(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRectF_ITF, source qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRectF_PTR() != nil {
		convArg1 = target.QRectF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if source != nil && source.QRectF_PTR() != nil {
		convArg2 = source.QRectF_PTR().GetCthis()
	}
	// arg: 3, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectRatioMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsScene::ItemIndexMethod itemIndexMethod() const

/*

 */
func (this *QGraphicsScene) ItemIndexMethod() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene15itemIndexMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemIndexMethod(QGraphicsScene::ItemIndexMethod)

/*

 */
func (this *QGraphicsScene) SetItemIndexMethod(method int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18setItemIndexMethodENS_15ItemIndexMethodE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), method)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortCacheEnabled() const

/*

 */
func (this *QGraphicsScene) IsSortCacheEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene18isSortCacheEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortCacheEnabled(bool)

/*

 */
func (this *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene19setSortCacheEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] int bspTreeDepth() const

/*

 */
func (this *QGraphicsScene) BspTreeDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene12bspTreeDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBspTreeDepth(int)

/*

 */
func (this *QGraphicsScene) SetBspTreeDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15setBspTreeDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF itemsBoundingRect() const

/*
Calculates and returns the bounding rect of all items on the scene. This function works by iterating over all items, and because of this, it can be slow for large scenes.

See also sceneRect().
*/
func (this *QGraphicsScene) ItemsBoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene17itemsBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * itemAt(const QPointF &, const QTransform &) const

/*
Returns the topmost visible item at the specified position, or 0 if there are no items at this position.

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

Note: See items() for a definition of which items are considered visible by this function.

This function was introduced in  Qt 4.6.

See also items(), collidingItems(), and Sorting.
*/
func (this *QGraphicsScene) ItemAt(pos qtcore.QPointF_ITF, deviceTransform qtgui.QTransform_ITF) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if deviceTransform != nil && deviceTransform.QTransform_PTR() != nil {
		convArg1 = deviceTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:179
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsItem * itemAt(qreal, qreal, const QTransform &) const

/*
Returns the topmost visible item at the specified position, or 0 if there are no items at this position.

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

Note: See items() for a definition of which items are considered visible by this function.

This function was introduced in  Qt 4.6.

See also items(), collidingItems(), and Sorting.
*/
func (this *QGraphicsScene) ItemAt_1(x float64, y float64, deviceTransform qtgui.QTransform_ITF) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg2 unsafe.Pointer
	if deviceTransform != nil && deviceTransform.QTransform_PTR() != nil {
		convArg2 = deviceTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtEddRK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath selectionArea() const

/*
Returns the selection area that was previously set with setSelectionArea(), or an empty QPainterPath if no selection area has been set.

See also setSelectionArea().
*/
func (this *QGraphicsScene) SelectionArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectionAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, const QTransform &)

/*
Sets the selection area to path. All items within this area are immediately selected, and all items outside are unselected. You can get the list of all selected items by calling selectedItems().

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

For an item to be selected, it must be marked as selectable (QGraphicsItem::ItemIsSelectable).

This function was introduced in  Qt 4.6.

See also clearSelection() and selectionArea().
*/
func (this *QGraphicsScene) SetSelectionArea(path qtgui.QPainterPath_ITF, deviceTransform qtgui.QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if deviceTransform != nil && deviceTransform.QTransform_PTR() != nil {
		convArg1 = deviceTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionMode, const QTransform &)

/*
Sets the selection area to path. All items within this area are immediately selected, and all items outside are unselected. You can get the list of all selected items by calling selectedItems().

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

For an item to be selected, it must be marked as selectable (QGraphicsItem::ItemIsSelectable).

This function was introduced in  Qt 4.6.

See also clearSelection() and selectionArea().
*/
func (this *QGraphicsScene) SetSelectionArea_1(path qtgui.QPainterPath_ITF, mode int, deviceTransform qtgui.QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if deviceTransform != nil && deviceTransform.QTransform_PTR() != nil {
		convArg2 = deviceTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionMode, const QTransform &)

/*
Sets the selection area to path. All items within this area are immediately selected, and all items outside are unselected. You can get the list of all selected items by calling selectedItems().

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

For an item to be selected, it must be marked as selectable (QGraphicsItem::ItemIsSelectable).

This function was introduced in  Qt 4.6.

See also clearSelection() and selectionArea().
*/
func (this *QGraphicsScene) SetSelectionArea_1_(path qtgui.QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 1, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	// arg: 2, const QTransform &=LValueReference, QTransform=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:185
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionMode, const QTransform &)

/*
Sets the selection area to path. All items within this area are immediately selected, and all items outside are unselected. You can get the list of all selected items by calling selectedItems().

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

For an item to be selected, it must be marked as selectable (QGraphicsItem::ItemIsSelectable).

This function was introduced in  Qt 4.6.

See also clearSelection() and selectionArea().
*/
func (this *QGraphicsScene) SetSelectionArea_1_1(path qtgui.QPainterPath_ITF, mode int) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 2, const QTransform &=LValueReference, QTransform=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:186
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const QTransform &)

/*
Sets the selection area to path. All items within this area are immediately selected, and all items outside are unselected. You can get the list of all selected items by calling selectedItems().

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

For an item to be selected, it must be marked as selectable (QGraphicsItem::ItemIsSelectable).

This function was introduced in  Qt 4.6.

See also clearSelection() and selectionArea().
*/
func (this *QGraphicsScene) SetSelectionArea_2(path qtgui.QPainterPath_ITF, selectionOperation int, mode int, deviceTransform qtgui.QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if deviceTransform != nil && deviceTransform.QTransform_PTR() != nil {
		convArg3 = deviceTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, selectionOperation, mode, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:186
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const QTransform &)

/*
Sets the selection area to path. All items within this area are immediately selected, and all items outside are unselected. You can get the list of all selected items by calling selectedItems().

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

For an item to be selected, it must be marked as selectable (QGraphicsItem::ItemIsSelectable).

This function was introduced in  Qt 4.6.

See also clearSelection() and selectionArea().
*/
func (this *QGraphicsScene) SetSelectionArea_2_(path qtgui.QPainterPath_ITF, selectionOperation int) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 2, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	// arg: 3, const QTransform &=LValueReference, QTransform=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, selectionOperation, mode, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:186
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setSelectionArea(const QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const QTransform &)

/*
Sets the selection area to path. All items within this area are immediately selected, and all items outside are unselected. You can get the list of all selected items by calling selectedItems().

deviceTransform is the transformation that applies to the view, and needs to be provided if the scene contains items that ignore transformations.

For an item to be selected, it must be marked as selectable (QGraphicsItem::ItemIsSelectable).

This function was introduced in  Qt 4.6.

See also clearSelection() and selectionArea().
*/
func (this *QGraphicsScene) SetSelectionArea_2_1(path qtgui.QPainterPath_ITF, selectionOperation int, mode int) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 3, const QTransform &=LValueReference, QTransform=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, selectionOperation, mode, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroyItemGroup(QGraphicsItemGroup *)

/*
Reparents all items in group to group's parent item, then removes group from the scene, and finally deletes it. The items' positions and transformations are mapped from the group to the group's parent.

See also createItemGroup() and QGraphicsItemGroup::removeFromGroup().
*/
func (this *QGraphicsScene) DestroyItemGroup(group QGraphicsItemGroup_ITF /*777 QGraphicsItemGroup **/) {
	var convArg0 unsafe.Pointer
	if group != nil && group.QGraphicsItemGroup_PTR() != nil {
		convArg0 = group.QGraphicsItemGroup_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsItem *)

/*
Adds or moves the item and all its childen to this scene. This scene takes ownership of the item.

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

If the item is already in a different scene, it will first be removed from its old scene, and then added to this scene as a top-level.

QGraphicsScene will send ItemSceneChange notifications to item while it is added to the scene. If item does not currently belong to a scene, only one notification is sent. If it does belong to scene already (i.e., it is moved to this scene), QGraphicsScene will send an addition notification as the item is removed from its previous scene.

If the item is a panel, the scene is active, and there is no active panel in the scene, then the item will be activated.

See also removeItem(), addEllipse(), addLine(), addPath(), addPixmap(), addRect(), addText(), addWidget(), and Sorting.
*/
func (this *QGraphicsScene) AddItem(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addItemEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(const QRectF &, const QPen &, const QBrush &)

/*
Creates and adds an ellipse item to the scene, and returns the item pointer. The geometry of the ellipse is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addLine(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddEllipse(rect qtcore.QRectF_ITF, pen qtgui.QPen_ITF, brush qtgui.QBrush_ITF) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg2 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(const QRectF &, const QPen &, const QBrush &)

/*
Creates and adds an ellipse item to the scene, and returns the item pointer. The geometry of the ellipse is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addLine(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddEllipse__(rect qtcore.QRectF_ITF) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(const QRectF &, const QPen &, const QBrush &)

/*
Creates and adds an ellipse item to the scene, and returns the item pointer. The geometry of the ellipse is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addLine(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddEllipse__1(rect qtcore.QRectF_ITF, pen qtgui.QPen_ITF) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)

/*
Creates and adds an ellipse item to the scene, and returns the item pointer. The geometry of the ellipse is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addLine(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddEllipse_1(x float64, y float64, w float64, h float64, pen qtgui.QPen_ITF, brush qtgui.QBrush_ITF) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	var convArg4 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg4 = pen.QPen_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg5 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)

/*
Creates and adds an ellipse item to the scene, and returns the item pointer. The geometry of the ellipse is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addLine(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddEllipse_1_(x float64, y float64, w float64, h float64) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	// arg: 4, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg5 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)

/*
Creates and adds an ellipse item to the scene, and returns the item pointer. The geometry of the ellipse is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addLine(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddEllipse_1_1(x float64, y float64, w float64, h float64, pen qtgui.QPen_ITF) *QGraphicsEllipseItem /*777 QGraphicsEllipseItem **/ {
	var convArg4 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg4 = pen.QPen_PTR().GetCthis()
	}
	// arg: 5, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg5 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:194
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLineItem * addLine(const QLineF &, const QPen &)

/*
Creates and adds a line item to the scene, and returns the item pointer. The geometry of the line is defined by line, and its pen is initialized to pen.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddLine(line qtcore.QLineF_ITF, pen qtgui.QPen_ITF) *QGraphicsLineItem /*777 QGraphicsLineItem **/ {
	var convArg0 unsafe.Pointer
	if line != nil && line.QLineF_PTR() != nil {
		convArg0 = line.QLineF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLineItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:194
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLineItem * addLine(const QLineF &, const QPen &)

/*
Creates and adds a line item to the scene, and returns the item pointer. The geometry of the line is defined by line, and its pen is initialized to pen.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddLine__(line qtcore.QLineF_ITF) *QGraphicsLineItem /*777 QGraphicsLineItem **/ {
	var convArg0 unsafe.Pointer
	if line != nil && line.QLineF_PTR() != nil {
		convArg0 = line.QLineF_PTR().GetCthis()
	}
	// arg: 1, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLineItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const QPen &)

/*
Creates and adds a line item to the scene, and returns the item pointer. The geometry of the line is defined by line, and its pen is initialized to pen.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddLine_1(x1 float64, y1 float64, x2 float64, y2 float64, pen qtgui.QPen_ITF) *QGraphicsLineItem /*777 QGraphicsLineItem **/ {
	var convArg4 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg4 = pen.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineEddddRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2, convArg4)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLineItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const QPen &)

/*
Creates and adds a line item to the scene, and returns the item pointer. The geometry of the line is defined by line, and its pen is initialized to pen.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addPath(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddLine_1_(x1 float64, y1 float64, x2 float64, y2 float64) *QGraphicsLineItem /*777 QGraphicsLineItem **/ {
	// arg: 4, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineEddddRK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2, convArg4)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLineItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:195
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPathItem * addPath(const QPainterPath &, const QPen &, const QBrush &)

/*
Creates and adds a path item to the scene, and returns the item pointer. The geometry of the path is defined by path, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddPath(path qtgui.QPainterPath_ITF, pen qtgui.QPen_ITF, brush qtgui.QBrush_ITF) *QGraphicsPathItem /*777 QGraphicsPathItem **/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg2 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsPathItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:195
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPathItem * addPath(const QPainterPath &, const QPen &, const QBrush &)

/*
Creates and adds a path item to the scene, and returns the item pointer. The geometry of the path is defined by path, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddPath__(path qtgui.QPainterPath_ITF) *QGraphicsPathItem /*777 QGraphicsPathItem **/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 1, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsPathItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:195
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPathItem * addPath(const QPainterPath &, const QPen &, const QBrush &)

/*
Creates and adds a path item to the scene, and returns the item pointer. The geometry of the path is defined by path, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddPath__1(path qtgui.QPainterPath_ITF, pen qtgui.QPen_ITF) *QGraphicsPathItem /*777 QGraphicsPathItem **/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsPathItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPixmapItem * addPixmap(const QPixmap &)

/*
Creates and adds a pixmap item to the scene, and returns the item pointer. The pixmap is defined by pixmap.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPath(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddPixmap(pixmap qtgui.QPixmap_ITF) *QGraphicsPixmapItem /*777 QGraphicsPixmapItem **/ {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9addPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsPixmapItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:197
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPolygonItem * addPolygon(const QPolygonF &, const QPen &, const QBrush &)

/*
Creates and adds a polygon item to the scene, and returns the item pointer. The polygon is defined by polygon, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPath(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddPolygon(polygon qtgui.QPolygonF_ITF, pen qtgui.QPen_ITF, brush qtgui.QBrush_ITF) *QGraphicsPolygonItem /*777 QGraphicsPolygonItem **/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg2 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:197
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPolygonItem * addPolygon(const QPolygonF &, const QPen &, const QBrush &)

/*
Creates and adds a polygon item to the scene, and returns the item pointer. The polygon is defined by polygon, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPath(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddPolygon__(polygon qtgui.QPolygonF_ITF) *QGraphicsPolygonItem /*777 QGraphicsPolygonItem **/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	// arg: 1, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:197
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsPolygonItem * addPolygon(const QPolygonF &, const QPen &, const QBrush &)

/*
Creates and adds a polygon item to the scene, and returns the item pointer. The polygon is defined by polygon, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPath(), addRect(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddPolygon__1(polygon qtgui.QPolygonF_ITF, pen qtgui.QPen_ITF) *QGraphicsPolygonItem /*777 QGraphicsPolygonItem **/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(const QRectF &, const QPen &, const QBrush &)

/*
Creates and adds a rectangle item to the scene, and returns the item pointer. The geometry of the rectangle is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0). For example, if a QRect(50, 50, 100, 100) is added, its top-left corner will be at (50, 50) relative to the origin in the items coordinate system.

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddRect(rect qtcore.QRectF_ITF, pen qtgui.QPen_ITF, brush qtgui.QBrush_ITF) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg2 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(const QRectF &, const QPen &, const QBrush &)

/*
Creates and adds a rectangle item to the scene, and returns the item pointer. The geometry of the rectangle is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0). For example, if a QRect(50, 50, 100, 100) is added, its top-left corner will be at (50, 50) relative to the origin in the items coordinate system.

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddRect__(rect qtcore.QRectF_ITF) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(const QRectF &, const QPen &, const QBrush &)

/*
Creates and adds a rectangle item to the scene, and returns the item pointer. The geometry of the rectangle is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0). For example, if a QRect(50, 50, 100, 100) is added, its top-left corner will be at (50, 50) relative to the origin in the items coordinate system.

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddRect__1(rect qtcore.QRectF_ITF, pen qtgui.QPen_ITF) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg1 = pen.QPen_PTR().GetCthis()
	}
	// arg: 2, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:206
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)

/*
Creates and adds a rectangle item to the scene, and returns the item pointer. The geometry of the rectangle is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0). For example, if a QRect(50, 50, 100, 100) is added, its top-left corner will be at (50, 50) relative to the origin in the items coordinate system.

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddRect_1(x float64, y float64, w float64, h float64, pen qtgui.QPen_ITF, brush qtgui.QBrush_ITF) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	var convArg4 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg4 = pen.QPen_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg5 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:206
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)

/*
Creates and adds a rectangle item to the scene, and returns the item pointer. The geometry of the rectangle is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0). For example, if a QRect(50, 50, 100, 100) is added, its top-left corner will be at (50, 50) relative to the origin in the items coordinate system.

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddRect_1_(x float64, y float64, w float64, h float64) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	// arg: 4, const QPen &=LValueReference, QPen=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg5 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:206
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const QPen &, const QBrush &)

/*
Creates and adds a rectangle item to the scene, and returns the item pointer. The geometry of the rectangle is defined by rect, and its pen and brush are initialized to pen and brush.

Note that the item's geometry is provided in item coordinates, and its position is initialized to (0, 0). For example, if a QRect(50, 50, 100, 100) is added, its top-left corner will be at (50, 50) relative to the origin in the items coordinate system.

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addText(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddRect_1_1(x float64, y float64, w float64, h float64, pen qtgui.QPen_ITF) *QGraphicsRectItem /*777 QGraphicsRectItem **/ {
	var convArg4 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg4 = pen.QPen_PTR().GetCthis()
	}
	// arg: 5, const QBrush &=LValueReference, QBrush=Record, , Invalid
	var convArg5 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, convArg4, convArg5)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsTextItem * addText(const QString &, const QFont &)

/*
Creates and adds a text item to the scene, and returns the item pointer. The text string is initialized to text, and its font is initialized to font.

The item's position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addRect(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddText(text string, font qtgui.QFont_ITF) *QGraphicsTextItem /*777 QGraphicsTextItem **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg1 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addTextERK7QStringRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsTextItem * addText(const QString &, const QFont &)

/*
Creates and adds a text item to the scene, and returns the item pointer. The text string is initialized to text, and its font is initialized to font.

The item's position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addRect(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddText__(text string) *QGraphicsTextItem /*777 QGraphicsTextItem **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QFont &=LValueReference, QFont=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7addTextERK7QStringRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsSimpleTextItem * addSimpleText(const QString &, const QFont &)

/*
Creates and adds a QGraphicsSimpleTextItem to the scene, and returns the item pointer. The text string is initialized to text, and its font is initialized to font.

The item's position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addRect(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddSimpleText(text string, font qtgui.QFont_ITF) *QGraphicsSimpleTextItem /*777 QGraphicsSimpleTextItem **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg1 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsSimpleTextItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsSimpleTextItem * addSimpleText(const QString &, const QFont &)

/*
Creates and adds a QGraphicsSimpleTextItem to the scene, and returns the item pointer. The text string is initialized to text, and its font is initialized to font.

The item's position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addRect(), addItem(), and addWidget().
*/
func (this *QGraphicsScene) AddSimpleText__(text string) *QGraphicsSimpleTextItem /*777 QGraphicsSimpleTextItem **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QFont &=LValueReference, QFont=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsSimpleTextItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsProxyWidget * addWidget(QWidget *, Qt::WindowFlags)

/*
Creates a new QGraphicsProxyWidget for widget, adds it to the scene, and returns a pointer to the proxy. wFlags set the default window flags for the embedding proxy widget.

The item's position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

Note that widgets with the Qt::WA_PaintOnScreen widget attribute set and widgets that wrap an external application or controller are not supported. Examples are QGLWidget and QAxWidget.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addRect(), addText(), addSimpleText(), and addItem().
*/
func (this *QGraphicsScene) AddWidget(widget QWidget_ITF /*777 QWidget **/, wFlags int) *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9addWidgetEP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsProxyWidget * addWidget(QWidget *, Qt::WindowFlags)

/*
Creates a new QGraphicsProxyWidget for widget, adds it to the scene, and returns a pointer to the proxy. wFlags set the default window flags for the embedding proxy widget.

The item's position is initialized to (0, 0).

If the item is visible (i.e., QGraphicsItem::isVisible() returns true), QGraphicsScene will emit changed() once control goes back to the event loop.

Note that widgets with the Qt::WA_PaintOnScreen widget attribute set and widgets that wrap an external application or controller are not supported. Examples are QGLWidget and QAxWidget.

See also addEllipse(), addLine(), addPixmap(), addPixmap(), addRect(), addText(), addSimpleText(), and addItem().
*/
func (this *QGraphicsScene) AddWidget__(widget QWidget_ITF /*777 QWidget **/) *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	wFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9addWidgetEP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QGraphicsItem *)

/*
Removes the item item and all its children from the scene. The ownership of item is passed on to the caller (i.e., QGraphicsScene will no longer delete item when destroyed).

See also addItem().
*/
func (this *QGraphicsScene) RemoveItem(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10removeItemEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:210
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * focusItem() const

/*
When the scene is active, this functions returns the scene's current focus item, or 0 if no item currently has focus. When the scene is inactive, this functions returns the item that will gain input focus when the scene becomes active.

The focus item receives keyboard input when the scene receives a key event.

See also setFocusItem(), QGraphicsItem::hasFocus(), and isActive().
*/
func (this *QGraphicsScene) FocusItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene9focusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusItem(QGraphicsItem *, Qt::FocusReason)

/*
Sets the scene's focus item to item, with the focus reason focusReason, after removing focus from any previous item that may have had focus.

If item is 0, or if it either does not accept focus (i.e., it does not have the QGraphicsItem::ItemIsFocusable flag enabled), or is not visible or not enabled, this function only removes focus from any previous focusitem.

If item is not 0, and the scene does not currently have focus (i.e., hasFocus() returns false), this function will call setFocus() automatically.

See also focusItem(), hasFocus(), and setFocus().
*/
func (this *QGraphicsScene) SetFocusItem(item QGraphicsItem_ITF /*777 QGraphicsItem **/, focusReason int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12setFocusItemEP13QGraphicsItemN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, focusReason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusItem(QGraphicsItem *, Qt::FocusReason)

/*
Sets the scene's focus item to item, with the focus reason focusReason, after removing focus from any previous item that may have had focus.

If item is 0, or if it either does not accept focus (i.e., it does not have the QGraphicsItem::ItemIsFocusable flag enabled), or is not visible or not enabled, this function only removes focus from any previous focusitem.

If item is not 0, and the scene does not currently have focus (i.e., hasFocus() returns false), this function will call setFocus() automatically.

See also focusItem(), hasFocus(), and setFocus().
*/
func (this *QGraphicsScene) SetFocusItem__(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, Qt::FocusReason=Elaborated, Qt::FocusReason=Enum, , Invalid
	focusReason := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12setFocusItemEP13QGraphicsItemN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, focusReason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:212
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus() const

/*
Returns true if the scene has focus; otherwise returns false. If the scene has focus, it will will forward key events from QKeyEvent to any item that has focus.

See also setFocus() and setFocusItem().
*/
func (this *QGraphicsScene) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)

/*
Sets focus on the scene by sending a QFocusEvent to the scene, passing focusReason as the reason. If the scene regains focus after having previously lost it while an item had focus, the last focus item will receive focus with focusReason as the reason.

If the scene already has focus, this function does nothing.

See also hasFocus(), clearFocus(), and setFocusItem().
*/
func (this *QGraphicsScene) SetFocus(focusReason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focusReason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)

/*
Sets focus on the scene by sending a QFocusEvent to the scene, passing focusReason as the reason. If the scene regains focus after having previously lost it while an item had focus, the last focus item will receive focus with focusReason as the reason.

If the scene already has focus, this function does nothing.

See also hasFocus(), clearFocus(), and setFocusItem().
*/
func (this *QGraphicsScene) SetFocus__() {
	// arg: 0, Qt::FocusReason=Elaborated, Qt::FocusReason=Enum, , Invalid
	focusReason := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focusReason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFocus()

/*
Clears focus from the scene. If any item has focus when this function is called, it will lose focus, and regain focus again once the scene regains focus.

A scene that does not have focus ignores key events.

See also hasFocus(), setFocus(), and setFocusItem().
*/
func (this *QGraphicsScene) ClearFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10clearFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStickyFocus(bool)

/*

 */
func (this *QGraphicsScene) SetStickyFocus(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14setStickyFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:217
// index:0
// Public Visibility=Default Availability=Available
// [1] bool stickyFocus() const

/*

 */
func (this *QGraphicsScene) StickyFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene11stickyFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * mouseGrabberItem() const

/*
Returns the current mouse grabber item, or 0 if no item is currently grabbing the mouse. The mouse grabber item is the item that receives all mouse events sent to the scene.

An item becomes a mouse grabber when it receives and accepts a mouse press event, and it stays the mouse grabber until either of the following events occur:


If the item receives a mouse release event when there are no other buttons pressed, it loses the mouse grab.
If the item becomes invisible (i.e., someone calls item->setVisible(false)), or if it becomes disabled (i.e., someone calls item->setEnabled(false)), it loses the mouse grab.
If the item is removed from the scene, it loses the mouse grab.


If the item loses its mouse grab, the scene will ignore all mouse events until a new item grabs the mouse (i.e., until a new item receives a mouse press event).
*/
func (this *QGraphicsScene) MouseGrabberItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene16mouseGrabberItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush backgroundBrush() const

/*

 */
func (this *QGraphicsScene) BackgroundBrush() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene15backgroundBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundBrush(const QBrush &)

/*

 */
func (this *QGraphicsScene) SetBackgroundBrush(brush qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18setBackgroundBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush foregroundBrush() const

/*

 */
func (this *QGraphicsScene) ForegroundBrush() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene15foregroundBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setForegroundBrush(const QBrush &)

/*

 */
func (this *QGraphicsScene) SetForegroundBrush(brush qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18setForegroundBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:227
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
This method is used by input methods to query a set of properties of the scene to be able to support complex input method operations as support for surrounding text and reconversions.

The query parameter specifies which property is queried.

See also QWidget::inputMethodQuery().
*/
func (this *QGraphicsScene) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void update(qreal, qreal, qreal, qreal)

/*
Schedules a redraw of the area rect on the scene.

See also sceneRect() and changed().
*/
func (this *QGraphicsScene) Update(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6updateEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:257
// index:1
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)

/*
Schedules a redraw of the area rect on the scene.

See also sceneRect() and changed().
*/
func (this *QGraphicsScene) Update_1(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:257
// index:1
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)

/*
Schedules a redraw of the area rect on the scene.

See also sceneRect() and changed().
*/
func (this *QGraphicsScene) Update_1_() {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void invalidate(qreal, qreal, qreal, qreal, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of the layers in rect on the scene. Any cached content in layers is unconditionally invalidated and redrawn.

You can use this function overload to notify QGraphicsScene of changes to the background or the foreground of the scene. This function is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled CacheBackground.

Example:


  QRectF TileScene::rectForTile(int x, int y) const
  {
      // Return the rectangle for the tile at position (x, y).
      return QRectF(x * tileWidth, y * tileHeight, tileWidth, tileHeight);
  }

  void TileScene::setTile(int x, int y, const QPixmap &pixmap)
  {
      // Sets or replaces the tile at position (x, y) with pixmap.
      if (x >= 0 && x < numTilesH && y >= 0 && y < numTilesV) {
          tiles[y][x] = pixmap;
          invalidate(rectForTile(x, y), BackgroundLayer);
      }
  }

  void TileScene::drawBackground(QPainter *painter, const QRectF &exposed)
  {
      // Draws all tiles that intersect the exposed area.
      for (int y = 0; y < numTilesV; ++y) {
          for (int x = 0; x < numTilesH; ++x) {
              QRectF rect = rectForTile(x, y);
              if (exposed.intersects(rect))
                  painter->drawPixmap(rect.topLeft(), tiles[y][x]);
          }
      }
  }



Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but BackgroundLayer is passed.

See also QGraphicsView::resetCachedContent().
*/
func (this *QGraphicsScene) Invalidate(x float64, y float64, w float64, h float64, layers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateEdddd6QFlagsINS_10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void invalidate(qreal, qreal, qreal, qreal, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of the layers in rect on the scene. Any cached content in layers is unconditionally invalidated and redrawn.

You can use this function overload to notify QGraphicsScene of changes to the background or the foreground of the scene. This function is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled CacheBackground.

Example:


  QRectF TileScene::rectForTile(int x, int y) const
  {
      // Return the rectangle for the tile at position (x, y).
      return QRectF(x * tileWidth, y * tileHeight, tileWidth, tileHeight);
  }

  void TileScene::setTile(int x, int y, const QPixmap &pixmap)
  {
      // Sets or replaces the tile at position (x, y) with pixmap.
      if (x >= 0 && x < numTilesH && y >= 0 && y < numTilesV) {
          tiles[y][x] = pixmap;
          invalidate(rectForTile(x, y), BackgroundLayer);
      }
  }

  void TileScene::drawBackground(QPainter *painter, const QRectF &exposed)
  {
      // Draws all tiles that intersect the exposed area.
      for (int y = 0; y < numTilesV; ++y) {
          for (int x = 0; x < numTilesH; ++x) {
              QRectF rect = rectForTile(x, y);
              if (exposed.intersects(rect))
                  painter->drawPixmap(rect.topLeft(), tiles[y][x]);
          }
      }
  }



Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but BackgroundLayer is passed.

See also QGraphicsView::resetCachedContent().
*/
func (this *QGraphicsScene) Invalidate__(x float64, y float64, w float64, h float64) {
	// arg: 4, QGraphicsScene::SceneLayers=Typedef, QGraphicsScene::SceneLayers=Typedef, QFlags<QGraphicsScene::SceneLayer>, Unexposed
	layers := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateEdddd6QFlagsINS_10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:258
// index:1
// Public Visibility=Default Availability=Available
// [-2] void invalidate(const QRectF &, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of the layers in rect on the scene. Any cached content in layers is unconditionally invalidated and redrawn.

You can use this function overload to notify QGraphicsScene of changes to the background or the foreground of the scene. This function is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled CacheBackground.

Example:


  QRectF TileScene::rectForTile(int x, int y) const
  {
      // Return the rectangle for the tile at position (x, y).
      return QRectF(x * tileWidth, y * tileHeight, tileWidth, tileHeight);
  }

  void TileScene::setTile(int x, int y, const QPixmap &pixmap)
  {
      // Sets or replaces the tile at position (x, y) with pixmap.
      if (x >= 0 && x < numTilesH && y >= 0 && y < numTilesV) {
          tiles[y][x] = pixmap;
          invalidate(rectForTile(x, y), BackgroundLayer);
      }
  }

  void TileScene::drawBackground(QPainter *painter, const QRectF &exposed)
  {
      // Draws all tiles that intersect the exposed area.
      for (int y = 0; y < numTilesV; ++y) {
          for (int x = 0; x < numTilesH; ++x) {
              QRectF rect = rectForTile(x, y);
              if (exposed.intersects(rect))
                  painter->drawPixmap(rect.topLeft(), tiles[y][x]);
          }
      }
  }



Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but BackgroundLayer is passed.

See also QGraphicsView::resetCachedContent().
*/
func (this *QGraphicsScene) Invalidate_1(rect qtcore.QRectF_ITF, layers int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateERK6QRectF6QFlagsINS_10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:258
// index:1
// Public Visibility=Default Availability=Available
// [-2] void invalidate(const QRectF &, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of the layers in rect on the scene. Any cached content in layers is unconditionally invalidated and redrawn.

You can use this function overload to notify QGraphicsScene of changes to the background or the foreground of the scene. This function is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled CacheBackground.

Example:


  QRectF TileScene::rectForTile(int x, int y) const
  {
      // Return the rectangle for the tile at position (x, y).
      return QRectF(x * tileWidth, y * tileHeight, tileWidth, tileHeight);
  }

  void TileScene::setTile(int x, int y, const QPixmap &pixmap)
  {
      // Sets or replaces the tile at position (x, y) with pixmap.
      if (x >= 0 && x < numTilesH && y >= 0 && y < numTilesV) {
          tiles[y][x] = pixmap;
          invalidate(rectForTile(x, y), BackgroundLayer);
      }
  }

  void TileScene::drawBackground(QPainter *painter, const QRectF &exposed)
  {
      // Draws all tiles that intersect the exposed area.
      for (int y = 0; y < numTilesV; ++y) {
          for (int x = 0; x < numTilesH; ++x) {
              QRectF rect = rectForTile(x, y);
              if (exposed.intersects(rect))
                  painter->drawPixmap(rect.topLeft(), tiles[y][x]);
          }
      }
  }



Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but BackgroundLayer is passed.

See also QGraphicsView::resetCachedContent().
*/
func (this *QGraphicsScene) Invalidate_1_() {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, QGraphicsScene::SceneLayers=Typedef, QGraphicsScene::SceneLayers=Typedef, QFlags<QGraphicsScene::SceneLayer>, Unexposed
	layers := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateERK6QRectF6QFlagsINS_10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:258
// index:1
// Public Visibility=Default Availability=Available
// [-2] void invalidate(const QRectF &, QGraphicsScene::SceneLayers)

/*
Invalidates and schedules a redraw of the layers in rect on the scene. Any cached content in layers is unconditionally invalidated and redrawn.

You can use this function overload to notify QGraphicsScene of changes to the background or the foreground of the scene. This function is commonly used for scenes with tile-based backgrounds to notify changes when QGraphicsView has enabled CacheBackground.

Example:


  QRectF TileScene::rectForTile(int x, int y) const
  {
      // Return the rectangle for the tile at position (x, y).
      return QRectF(x * tileWidth, y * tileHeight, tileWidth, tileHeight);
  }

  void TileScene::setTile(int x, int y, const QPixmap &pixmap)
  {
      // Sets or replaces the tile at position (x, y) with pixmap.
      if (x >= 0 && x < numTilesH && y >= 0 && y < numTilesV) {
          tiles[y][x] = pixmap;
          invalidate(rectForTile(x, y), BackgroundLayer);
      }
  }

  void TileScene::drawBackground(QPainter *painter, const QRectF &exposed)
  {
      // Draws all tiles that intersect the exposed area.
      for (int y = 0; y < numTilesV; ++y) {
          for (int x = 0; x < numTilesH; ++x) {
              QRectF rect = rectForTile(x, y);
              if (exposed.intersects(rect))
                  painter->drawPixmap(rect.topLeft(), tiles[y][x]);
          }
      }
  }



Note that QGraphicsView currently supports background caching only (see QGraphicsView::CacheBackground). This function is equivalent to calling update() if any layer but BackgroundLayer is passed.

See also QGraphicsView::resetCachedContent().
*/
func (this *QGraphicsScene) Invalidate_1_1(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, QGraphicsScene::SceneLayers=Typedef, QGraphicsScene::SceneLayers=Typedef, QFlags<QGraphicsScene::SceneLayer>, Unexposed
	layers := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateERK6QRectF6QFlagsINS_10SceneLayerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, layers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * style() const

/*
Returns the scene's style, or the same as QApplication::style() if the scene has not been explicitly assigned a style.

This function was introduced in  Qt 4.4.

See also setStyle().
*/
func (this *QGraphicsScene) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)

/*
Sets or replaces the style of the scene to style, and reparents the style to this scene. Any previously assigned style is deleted. The scene's style defaults to QApplication::style(), and serves as the default for all QGraphicsWidget items in the scene.

Changing the style, either directly by calling this function, or indirectly by calling QApplication::setStyle(), will automatically update the style for all widgets in the scene that do not have a style explicitly assigned to them.

If style is 0, QGraphicsScene will revert to QApplication::style().

This function was introduced in  Qt 4.4.

See also style().
*/
func (this *QGraphicsScene) SetStyle(style QStyle_ITF /*777 QStyle **/) {
	var convArg0 unsafe.Pointer
	if style != nil && style.QStyle_PTR() != nil {
		convArg0 = style.QStyle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene8setStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:239
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
func (this *QGraphicsScene) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QGraphicsScene) SetFont(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:242
// index:0
// Public Visibility=Default Availability=Available
// [16] QPalette palette() const

/*

 */
func (this *QGraphicsScene) Palette() *qtgui.QPalette /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene7paletteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &)

/*

 */
func (this *QGraphicsScene) SetPalette(palette qtgui.QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if palette != nil && palette.QPalette_PTR() != nil {
		convArg0 = palette.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:245
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Returns true if the scene is active (e.g., it's viewed by at least one QGraphicsView that is active); otherwise returns false.

This function was introduced in  Qt 4.6.

See also QGraphicsItem::isActive() and QWidget::isActiveWindow().
*/
func (this *QGraphicsScene) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:246
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * activePanel() const

/*
Returns the current active panel, or 0 if no panel is currently active.

This function was introduced in  Qt 4.6.

See also QGraphicsScene::setActivePanel().
*/
func (this *QGraphicsScene) ActivePanel() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene11activePanelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActivePanel(QGraphicsItem *)

/*
Activates item, which must be an item in this scene. You can also pass 0 for item, in which case QGraphicsScene will deactivate any currently active panel.

If the scene is currently inactive, item remains inactive until the scene becomes active (or, ir item is 0, no item will be activated).

This function was introduced in  Qt 4.6.

See also activePanel(), isActive(), and QGraphicsItem::isActive().
*/
func (this *QGraphicsScene) SetActivePanel(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:248
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * activeWindow() const

/*
Returns the current active window, or 0 if no window is currently active.

This function was introduced in  Qt 4.4.

See also QGraphicsScene::setActiveWindow().
*/
func (this *QGraphicsScene) ActiveWindow() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene12activeWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveWindow(QGraphicsWidget *)

/*
Activates widget, which must be a widget in this scene. You can also pass 0 for widget, in which case QGraphicsScene will deactivate any currently active window.

This function was introduced in  Qt 4.4.

See also activeWindow() and QGraphicsWidget::isActiveWindow().
*/
func (this *QGraphicsScene) SetActiveWindow(widget QGraphicsWidget_ITF /*777 QGraphicsWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QGraphicsWidget_PTR() != nil {
		convArg0 = widget.QGraphicsWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:251
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sendEvent(QGraphicsItem *, QEvent *)

/*
Sends event event to item item through possible event filters.

The event is sent only if the item is enabled.

Returns false if the event was filtered or if the item is disabled. Otherwise returns the value that was returned from the event handler.

This function was introduced in  Qt 4.6.

See also QGraphicsItem::sceneEvent() and QGraphicsItem::sceneEventFilter().
*/
func (this *QGraphicsScene) SendEvent(item QGraphicsItem_ITF /*777 QGraphicsItem **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:253
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minimumRenderSize() const

/*

 */
func (this *QGraphicsScene) MinimumRenderSize() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScene17minimumRenderSizeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:254
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumRenderSize(qreal)

/*

 */
func (this *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene20setMinimumRenderSizeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void advance()

/*
This slot advances the scene by one step, by calling QGraphicsItem::advance() for all items on the scene. This is done in two phases: in the first phase, all items are notified that the scene is about to change, and in the second phase all items are notified that they can move. In the first phase, QGraphicsItem::advance() is called passing a value of 0 as an argument, and 1 is passed in the second phase.

Note that you can also use the Animation Framework for animations.

See also QGraphicsItem::advance() and QTimeLine.
*/
func (this *QGraphicsScene) Advance() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene7advanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSelection()

/*
Clears the current selection.

See also setSelectionArea() and selectedItems().
*/
func (this *QGraphicsScene) ClearSelection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14clearSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes and deletes all items from the scene, but otherwise leaves the state of the scene unchanged.

This function was introduced in  Qt 4.4.

See also addItem().
*/
func (this *QGraphicsScene) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:264
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().

Processes the event event, and dispatches it to the respective event handlers.

In addition to calling the convenience event handlers, this function is responsible for converting mouse move events to hover events for when there is no mouse grabber item. Hover events are delivered directly to items; there is no convenience function for them.

Unlike QWidget, QGraphicsScene does not have the convenience functions enterEvent() and leaveEvent(). Use this function to obtain those events instead.

See also contextMenuEvent(), keyPressEvent(), keyReleaseEvent(), mousePressEvent(), mouseMoveEvent(), mouseReleaseEvent(), mouseDoubleClickEvent(), focusInEvent(), and focusOutEvent().
*/
func (this *QGraphicsScene) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:265
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().

QGraphicsScene filters QApplication's events to detect palette and font changes.
*/
func (this *QGraphicsScene) EventFilter(watched qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if watched != nil && watched.QObject_PTR() != nil {
		convArg0 = watched.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:266
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QGraphicsSceneContextMenuEvent *)

/*
This event handler, for event contextMenuEvent, can be reimplemented in a subclass to receive context menu events. The default implementation forwards the event to the topmost visible item that accepts context menu events at the position of the event. If no items accept context menu events at this position, the event is ignored.

Note: See items() for a definition of which items are considered visible by this function.

See also QGraphicsItem::contextMenuEvent().
*/
func (this *QGraphicsScene) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF /*777 QGraphicsSceneContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneContextMenuEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:267
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive drag enter events for the scene.

The default implementation accepts the event and prepares the scene to accept drag move events.

See also QGraphicsItem::dragEnterEvent(), dragMoveEvent(), dragLeaveEvent(), and dropEvent().
*/
func (this *QGraphicsScene) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:268
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive drag move events for the scene.

Note: See items() for a definition of which items are considered visible by this function.

See also QGraphicsItem::dragMoveEvent(), dragEnterEvent(), dragLeaveEvent(), and dropEvent().
*/
func (this *QGraphicsScene) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:269
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive drag leave events for the scene.

See also QGraphicsItem::dragLeaveEvent(), dragEnterEvent(), dragMoveEvent(), and dropEvent().
*/
func (this *QGraphicsScene) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:270
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive drop events for the scene.

See also QGraphicsItem::dropEvent(), dragEnterEvent(), dragMoveEvent(), and dragLeaveEvent().
*/
func (this *QGraphicsScene) DropEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:271
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
This event handler, for event focusEvent, can be reimplemented in a subclass to receive focus in events.

The default implementation sets focus on the scene, and then on the last focus item.

See also QGraphicsItem::focusOutEvent().
*/
func (this *QGraphicsScene) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:272
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
This event handler, for event focusEvent, can be reimplemented in a subclass to receive focus out events.

The default implementation removes focus from any focus item, then removes focus from the scene.

See also QGraphicsItem::focusInEvent().
*/
func (this *QGraphicsScene) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:273
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void helpEvent(QGraphicsSceneHelpEvent *)

/*
This event handler, for event helpEvent, can be reimplemented in a subclass to receive help events. The events are of type QEvent::ToolTip, which are created when a tooltip is requested.

The default implementation shows the tooltip of the topmost visible item, i.e., the item with the highest z-value, at the mouse cursor position. If no item has a tooltip set, this function does nothing.

Note: See items() for a definition of which items are considered visible by this function.

See also QGraphicsItem::toolTip() and QGraphicsSceneHelpEvent.
*/
func (this *QGraphicsScene) HelpEvent(event QGraphicsSceneHelpEvent_ITF /*777 QGraphicsSceneHelpEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHelpEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHelpEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:274
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
This event handler, for event keyEvent, can be reimplemented in a subclass to receive keypress events. The default implementation forwards the event to current focus item.

See also QGraphicsItem::keyPressEvent() and focusItem().
*/
func (this *QGraphicsScene) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:275
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
This event handler, for event keyEvent, can be reimplemented in a subclass to receive key release events. The default implementation forwards the event to current focus item.

See also QGraphicsItem::keyReleaseEvent() and focusItem().
*/
func (this *QGraphicsScene) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:276
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event mouseEvent, can be reimplemented in a subclass to receive mouse press events for the scene.

The default implementation depends on the state of the scene. If there is a mouse grabber item, then the event is sent to the mouse grabber. Otherwise, it is forwarded to the topmost visible item that accepts mouse events at the scene position from the event, and that item promptly becomes the mouse grabber item.

If there is no item at the given position on the scene, the selection area is reset, any focus item loses its input focus, and the event is then ignored.

Note: See items() for a definition of which items are considered visible by this function.

See also QGraphicsItem::mousePressEvent() and QGraphicsItem::setAcceptedMouseButtons().
*/
func (this *QGraphicsScene) MousePressEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:277
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event mouseEvent, can be reimplemented in a subclass to receive mouse move events for the scene.

The default implementation depends on the mouse grabber state. If there is a mouse grabber item, the event is sent to the mouse grabber. If there are any items that accept hover events at the current position, the event is translated into a hover event and accepted; otherwise it's ignored.

See also QGraphicsItem::mousePressEvent(), QGraphicsItem::mouseReleaseEvent(), QGraphicsItem::mouseDoubleClickEvent(), and QGraphicsItem::setAcceptedMouseButtons().
*/
func (this *QGraphicsScene) MouseMoveEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:278
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event mouseEvent, can be reimplemented in a subclass to receive mouse release events for the scene.

The default implementation depends on the mouse grabber state. If there is no mouse grabber, the event is ignored. Otherwise, if there is a mouse grabber item, the event is sent to the mouse grabber. If this mouse release represents the last pressed button on the mouse, the mouse grabber item then loses the mouse grab.

See also QGraphicsItem::mousePressEvent(), QGraphicsItem::mouseMoveEvent(), QGraphicsItem::mouseDoubleClickEvent(), and QGraphicsItem::setAcceptedMouseButtons().
*/
func (this *QGraphicsScene) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:279
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event mouseEvent, can be reimplemented in a subclass to receive mouse doubleclick events for the scene.

If someone doubleclicks on the scene, the scene will first receive a mouse press event, followed by a release event (i.e., a click), then a doubleclick event, and finally a release event. If the doubleclick event is delivered to a different item than the one that received the first press and release, it will be delivered as a press event. However, tripleclick events are not delivered as doubleclick events in this case.

The default implementation is similar to mousePressEvent().

Note: See items() for a definition of which items are considered visible by this function.

See also QGraphicsItem::mousePressEvent(), QGraphicsItem::mouseMoveEvent(), QGraphicsItem::mouseReleaseEvent(), and QGraphicsItem::setAcceptedMouseButtons().
*/
func (this *QGraphicsScene) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:280
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QGraphicsSceneWheelEvent *)

/*
This event handler, for event wheelEvent, can be reimplemented in a subclass to receive mouse wheel events for the scene.

By default, the event is delivered to the topmost visible item under the cursor. If ignored, the event propagates to the item beneath, and again until the event is accepted, or it reaches the scene. If no items accept the event, it is ignored.

Note: See items() for a definition of which items are considered visible by this function.

See also QGraphicsItem::wheelEvent().
*/
func (this *QGraphicsScene) WheelEvent(event QGraphicsSceneWheelEvent_ITF /*777 QGraphicsSceneWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneWheelEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:281
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive input method events for the scene.

The default implementation forwards the event to the focusItem(). If no item currently has focus or the current focus item does not accept input methods, this function does nothing.

See also QGraphicsItem::inputMethodEvent().
*/
func (this *QGraphicsScene) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QInputMethodEvent_PTR() != nil {
		convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:283
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawBackground(QPainter *, const QRectF &)

/*
Draws the background of the scene using painter, before any items and the foreground are drawn. Reimplement this function to provide a custom background for the scene.

All painting is done in scene coordinates. The rect parameter is the exposed rectangle.

If all you want is to define a color, texture, or gradient for the background, you can call setBackgroundBrush() instead.

See also drawForeground() and drawItems().
*/
func (this *QGraphicsScene) DrawBackground(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:284
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawForeground(QPainter *, const QRectF &)

/*
Draws the foreground of the scene using painter, after the background and all items have been drawn. Reimplement this function to provide a custom foreground for the scene.

All painting is done in scene coordinates. The rect parameter is the exposed rectangle.

If all you want is to define a color, texture or gradient for the foreground, you can call setForegroundBrush() instead.

See also drawBackground() and drawItems().
*/
func (this *QGraphicsScene) DrawForeground(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:285
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawItems(QPainter *, int, QGraphicsItem **, const QStyleOptionGraphicsItem *, QWidget *)

/*

 */
func (this *QGraphicsScene) DrawItems(painter qtgui.QPainter_ITF /*777 QPainter **/, numItems int, items []interface{}, options []interface{}, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg4 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numItems, items, options, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:285
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawItems(QPainter *, int, QGraphicsItem **, const QStyleOptionGraphicsItem *, QWidget *)

/*

 */
func (this *QGraphicsScene) DrawItems__(painter qtgui.QPainter_ITF /*777 QPainter **/, numItems int, items []interface{}, options []interface{}) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 4, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numItems, items, options, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:295
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Finds a new widget to give the keyboard focus to, as appropriate for Tab and Shift+Tab, and returns true if it can find a new widget, or false if it cannot. If next is true, this function searches forward; if next is false, it searches backward.

You can reimplement this function in a subclass of QGraphicsScene to provide fine-grained control over how tab focus passes inside your scene. The default implementation is based on the tab focus chain defined by QGraphicsWidget::setTabOrder().

This function was introduced in  Qt 4.4.
*/
func (this *QGraphicsScene) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneRectChanged(const QRectF &)

/*
This signal is emitted by QGraphicsScene whenever the scene rect changes. The rect parameter is the new scene rectangle.

See also QGraphicsView::updateSceneRect().
*/
func (this *QGraphicsScene) SceneRectChanged(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16sceneRectChangedERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()

/*
This signal is emitted by QGraphicsScene whenever the selection changes. You can call selectedItems() to get the new list of selected items.

The selection changes whenever an item is selected or unselected, a selection area is set, cleared or otherwise changed, if a preselected item is added to the scene, or if a selected item is removed from the scene.

QGraphicsScene emits this signal only once for group selection operations. For example, if you set a selection area, select or unselect a QGraphicsItemGroup, or if you add or remove from the scene a parent item that contains several selected items, selectionChanged() is emitted only once after the operation has completed (instead of once for each item).

This function was introduced in  Qt 4.3.

See also setSelectionArea(), selectedItems(), and QGraphicsItem::setSelected().
*/
func (this *QGraphicsScene) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusItemChanged(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)

/*
This signal is emitted by QGraphicsScene whenever focus changes in the scene (i.e., when an item gains or loses input focus, or when focus passes from one item to another). You can connect to this signal if you need to keep track of when other items gain input focus. It is particularly useful for implementing virtual keyboards, input methods, and cursor items.

oldFocusItem is a pointer to the item that previously had focus, or 0 if no item had focus before the signal was emitted. newFocusItem is a pointer to the item that gained input focus, or 0 if focus was lost. reason is the reason for the focus change (e.g., if the scene was deactivated while an input field had focus, oldFocusItem would point to the input field item, newFocusItem would be 0, and reason would be Qt::ActiveWindowFocusReason.
*/
func (this *QGraphicsScene) FocusItemChanged(newFocus QGraphicsItem_ITF /*777 QGraphicsItem **/, oldFocus QGraphicsItem_ITF /*777 QGraphicsItem **/, reason int) {
	var convArg0 unsafe.Pointer
	if newFocus != nil && newFocus.QGraphicsItem_PTR() != nil {
		convArg0 = newFocus.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if oldFocus != nil && oldFocus.QGraphicsItem_PTR() != nil {
		convArg1 = oldFocus.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScene16focusItemChangedEP13QGraphicsItemS1_N2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, reason)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the indexing algorithms QGraphicsScene provides for managing positional information about items on the scene.



See also setItemIndexMethod() and bspTreeDepth.

*/
type QGraphicsScene__ItemIndexMethod = int

// A Binary Space Partitioning tree is applied. All QGraphicsScene's item location algorithms are of an order close to logarithmic complexity, by making use of binary search. Adding, moving and removing items is logarithmic. This approach is best for static scenes (i.e., scenes where most items do not move).
const QGraphicsScene__BspTreeIndex QGraphicsScene__ItemIndexMethod = 0

//
const QGraphicsScene__NoIndex QGraphicsScene__ItemIndexMethod = -1

/*


 */
type QGraphicsScene__SceneLayer = int

//
const QGraphicsScene__ItemLayer QGraphicsScene__SceneLayer = 1

//
const QGraphicsScene__BackgroundLayer QGraphicsScene__SceneLayer = 2

//
const QGraphicsScene__ForegroundLayer QGraphicsScene__SceneLayer = 4

//
const QGraphicsScene__AllLayers QGraphicsScene__SceneLayer = 65535

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
