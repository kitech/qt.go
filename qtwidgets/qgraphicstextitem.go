package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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

// bool sceneEvent(QEvent *)
func (this *QGraphicsTextItem) InheritSceneEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEvent", f)
}

// void mousePressEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMousePressEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void contextMenuEvent(QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsTextItem) InheritContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QGraphicsTextItem) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsTextItem) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QGraphicsTextItem) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QGraphicsTextItem) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void dragEnterEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDropEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsTextItem) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void hoverEnterEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) InheritHoverEnterEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverEnterEvent", f)
}

// void hoverMoveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) InheritHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) InheritHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsTextItem) InheritInputMethodQuery(f func(query int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "inputMethodQuery", f)
}

// bool supportsExtension(QGraphicsItem::Extension)
func (this *QGraphicsTextItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsTextItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const QVariant &)
func (this *QGraphicsTextItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

/*

 */
type QGraphicsTextItem struct {
	*qtrt.CObject
}
type QGraphicsTextItem_ITF interface {
	QGraphicsTextItem_PTR() *QGraphicsTextItem
}

func (ptr *QGraphicsTextItem) QGraphicsTextItem_PTR() *QGraphicsTextItem { return ptr }

func (this *QGraphicsTextItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsTextItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGraphicsTextItemFromPointer(cthis unsafe.Pointer) *QGraphicsTextItem {
	return &QGraphicsTextItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsTextItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsTextItem {
	return NewQGraphicsTextItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:872
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsTextItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:877
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTextItem(QGraphicsItem *)

/*

 */
func (*QGraphicsTextItem) NewForInherit(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsTextItem {
	return NewQGraphicsTextItem(parent)
}
func NewQGraphicsTextItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsTextItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsTextItem")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:877
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTextItem(QGraphicsItem *)

/*

 */
func (*QGraphicsTextItem) NewForInherit__() *QGraphicsTextItem {
	return NewQGraphicsTextItem__()
}
func NewQGraphicsTextItem__() *QGraphicsTextItem {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsTextItem")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:878
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTextItem(const QString &, QGraphicsItem *)

/*

 */
func (*QGraphicsTextItem) NewForInherit_1(text string, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsTextItem {
	return NewQGraphicsTextItem_1(text, parent)
}
func NewQGraphicsTextItem_1(text string, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsTextItem {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg1 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsTextItem")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:878
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTextItem(const QString &, QGraphicsItem *)

/*

 */
func (*QGraphicsTextItem) NewForInherit_1_(text string) *QGraphicsTextItem {
	return NewQGraphicsTextItem_1_(text)
}
func NewQGraphicsTextItem_1_(text string) *QGraphicsTextItem {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsTextItem")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:879
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsTextItem()

/*

 */
func DeleteQGraphicsTextItem(this *QGraphicsTextItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:881
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml() const

/*

 */
func (this *QGraphicsTextItem) ToHtml() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem6toHtmlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:882
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &)

/*

 */
func (this *QGraphicsTextItem) SetHtml(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:884
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toPlainText() const

/*

 */
func (this *QGraphicsTextItem) ToPlainText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem11toPlainTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:885
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)

/*

 */
func (this *QGraphicsTextItem) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:887
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
func (this *QGraphicsTextItem) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:888
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QGraphicsTextItem) SetFont(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:890
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTextColor(const QColor &)

/*

 */
func (this *QGraphicsTextItem) SetDefaultTextColor(c qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg0 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:891
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor defaultTextColor() const

/*

 */
func (this *QGraphicsTextItem) DefaultTextColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem16defaultTextColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:893
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect() const

/*
This pure virtual function defines the outer bounds of the item as a rectangle; all painting must be restricted to inside an item's bounding rect. QGraphicsView uses this to determine whether the item requires redrawing.

Although the item's shape can be arbitrary, the bounding rect is always rectangular, and it is unaffected by the items' transformation.

If you want to change the item's bounding rectangle, you must first call prepareGeometryChange(). This notifies the scene of the imminent change, so that it can update its item geometry index; otherwise, the scene will be unaware of the item's new geometry, and the results are undefined (typically, rendering artifacts are left within the view).

Reimplement this function to let QGraphicsView determine what parts of the widget, if any, need to be redrawn.

Note: For shapes that paint an outline / stroke, it is important to include half the pen width in the bounding rect. It is not necessary to compensate for antialiasing, though.

Example:


  QRectF CircleItem::boundingRect() const
  {
      qreal penWidth = 1;
      return QRectF(-radius - penWidth / 2, -radius - penWidth / 2,
                    diameter + penWidth, diameter + penWidth);
  }



See also boundingRegion(), shape(), contains(), The Graphics View Coordinate System, and prepareGeometryChange().
*/
func (this *QGraphicsTextItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:894
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape() const

/*
Returns the shape of this item as a QPainterPath in local coordinates. The shape is used for many things, including collision detection, hit tests, and for the QGraphicsScene::items() functions.

The default implementation calls boundingRect() to return a simple rectangular shape, but subclasses can reimplement this function to return a more accurate shape for non-rectangular items. For example, a round item may choose to return an elliptic shape for better collision detection. For example:


  QPainterPath RoundItem::shape() const
  {
      QPainterPath path;
      path.addEllipse(boundingRect());
      return path;
  }



The outline of a shape can vary depending on the width and style of the pen used when drawing. If you want to include this outline in the item's shape, you can create a shape from the stroke using QPainterPathStroker.

This function is called by the default implementations of contains() and collidesWithPath().

See also boundingRect(), contains(), prepareGeometryChange(), and QPainterPathStroker.
*/
func (this *QGraphicsTextItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:895
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &) const

/*
Returns true if this item contains point, which is in local coordinates; otherwise, false is returned. It is most often called from QGraphicsView to determine what item is under the cursor, and for that reason, the implementation of this function should be as light-weight as possible.

By default, this function calls shape(), but you can reimplement it in a subclass to provide a (perhaps more efficient) implementation.

See also shape(), boundingRect(), and collidesWithPath().
*/
func (this *QGraphicsTextItem) Contains(point qtcore.QPointF_ITF) bool {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:897
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
This function, which is usually called by QGraphicsView, paints the contents of an item in local coordinates.

Reimplement this function in a QGraphicsItem subclass to provide the item's painting implementation, using painter. The option parameter provides style options for the item, such as its state, exposed area and its level-of-detail hints. The widget argument is optional. If provided, it points to the widget that is being painted on; otherwise, it is 0. For cached painting, widget is always 0.


  void RoundRectItem::paint(QPainter *painter,
                            const QStyleOptionGraphicsItem *option,
                            QWidget *widget)
  {
      painter->drawRoundedRect(-10, -10, 20, 20, 5, 5);
  }



The painter's pen is 0-width by default, and its pen is initialized to the QPalette::Text brush from the paint device's palette. The brush is initialized to QPalette::Window.

Make sure to constrain all painting inside the boundaries of boundingRect() to avoid rendering artifacts (as QGraphicsView does not clip the painter for you). In particular, when QPainter renders the outline of a shape using an assigned QPen, half of the outline will be drawn outside, and half inside, the shape you're rendering (e.g., with a pen width of 2 units, you must draw outlines 1 unit inside boundingRect()). QGraphicsItem does not support use of cosmetic pens with a non-zero width.

All painting is done in local coordinates.

Note: It is mandatory that an item will always redraw itself in the exact same way, unless update() was called; otherwise visual artifacts may occur. In other words, two subsequent calls to paint() must always produce the same output, unless update() was called between them.

Note: Enabling caching for an item does not guarantee that paint() will be invoked only once by the Graphics View framework, even without any explicit call to update(). See the documentation of setCacheMode() for more details.

See also setCacheMode(), QPen::width(), Item Coordinates, and ItemUsesExtendedStyleOption.
*/
func (this *QGraphicsTextItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:899
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *) const

/*
Returns true if this item's bounding rect is completely obscured by the opaque shape of item.

The base implementation maps item's opaqueArea() to this item's coordinate system, and then checks if this item's boundingRect() is fully contained within the mapped shape.

You can reimplement this function to provide a custom algorithm for determining whether this item is obscured by item.

See also opaqueArea() and isObscured().
*/
func (this *QGraphicsTextItem) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:900
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea() const

/*
This virtual function returns a shape representing the area where this item is opaque. An area is opaque if it is filled using an opaque brush or color (i.e., not transparent).

This function is used by isObscuredBy(), which is called by underlying items to determine if they are obscured by this item.

The default implementation returns an empty QPainterPath, indicating that this item is completely transparent and does not obscure any other items.

See also isObscuredBy(), isObscured(), and shape().
*/
func (this *QGraphicsTextItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:903
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const

/*
Returns the type of an item as an int. All standard graphicsitem classes are associated with a unique value; see QGraphicsItem::Type. This type information is used by qgraphicsitem_cast() to distinguish between types.

The default implementation (in QGraphicsItem) returns UserType.

To enable use of qgraphicsitem_cast() with a custom item, reimplement this function and declare a Type enum value equal to your custom item's type. Custom items must return a value larger than or equal to UserType (65536).

For example:


  class CustomItem : public QGraphicsItem
  {
  public:
     enum { Type = UserType + 1 };

     int type() const
     {
         // Enable the use of qgraphicsitem_cast with this item.
         return Type;
     }
     ...
  };



See also UserType.
*/
func (this *QGraphicsTextItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:905
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextWidth(qreal)

/*

 */
func (this *QGraphicsTextItem) SetTextWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setTextWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:906
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal textWidth() const

/*

 */
func (this *QGraphicsTextItem) TextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem9textWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:908
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()

/*

 */
func (this *QGraphicsTextItem) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:910
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocument(QTextDocument *)

/*

 */
func (this *QGraphicsTextItem) SetDocument(document qtgui.QTextDocument_ITF /*777 QTextDocument **/) {
	var convArg0 unsafe.Pointer
	if document != nil && document.QTextDocument_PTR() != nil {
		convArg0 = document.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:911
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document() const

/*

 */
func (this *QGraphicsTextItem) Document() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:913
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)

/*

 */
func (this *QGraphicsTextItem) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:914
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextInteractionFlags textInteractionFlags() const

/*

 */
func (this *QGraphicsTextItem) TextInteractionFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem20textInteractionFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:916
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabChangesFocus(bool)

/*

 */
func (this *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem18setTabChangesFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:917
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabChangesFocus() const

/*

 */
func (this *QGraphicsTextItem) TabChangesFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem15tabChangesFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:919
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenExternalLinks(bool)

/*

 */
func (this *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem20setOpenExternalLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:920
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openExternalLinks() const

/*

 */
func (this *QGraphicsTextItem) OpenExternalLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem17openExternalLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:922
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextCursor(const QTextCursor &)

/*

 */
func (this *QGraphicsTextItem) SetTextCursor(cursor qtgui.QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:923
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor textCursor() const

/*

 */
func (this *QGraphicsTextItem) TextCursor() *qtgui.QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10textCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:926
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkActivated(const QString &)

/*

 */
func (this *QGraphicsTextItem) LinkActivated(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13linkActivatedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:927
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkHovered(const QString &)

/*

 */
func (this *QGraphicsTextItem) LinkHovered(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem11linkHoveredERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:930
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool sceneEvent(QEvent *)

/*
This virtual function receives events to this item. Reimplement this function to intercept events before they are dispatched to the specialized event handlers contextMenuEvent(), focusInEvent(), focusOutEvent(), hoverEnterEvent(), hoverMoveEvent(), hoverLeaveEvent(), keyPressEvent(), keyReleaseEvent(), mousePressEvent(), mouseReleaseEvent(), mouseMoveEvent(), and mouseDoubleClickEvent().

Returns true if the event was recognized and handled; otherwise, (e.g., if the event type was not recognized,) false is returned.

event is the intercepted event.
*/
func (this *QGraphicsTextItem) SceneEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem10sceneEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:931
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse press events for this item. Mouse press events are only delivered to items that accept the mouse button that is pressed. By default, an item accepts all mouse buttons, but you can change this by calling setAcceptedMouseButtons().

The mouse press event decides which item should become the mouse grabber (see QGraphicsScene::mouseGrabberItem()). If you do not reimplement this function, the press event will propagate to any topmost item beneath this item, and no other mouse events will be delivered to this item.

If you do reimplement this function, event will by default be accepted (see QEvent::accept()), and this item is then the mouse grabber. This allows the item to receive future move, release and doubleclick events. If you call QEvent::ignore() on event, this item will lose the mouse grab, and event will propagate to any topmost item beneath. No further mouse events will be delivered to this item unless a new mouse press event is received.

The default implementation handles basic item interaction, such as selection and moving. If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mousePressEvent() in your reimplementation.

The event is QEvent::ignore()d for items that are neither movable nor selectable.

See also mouseMoveEvent(), mouseReleaseEvent(), mouseDoubleClickEvent(), and sceneEvent().
*/
func (this *QGraphicsTextItem) MousePressEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:932
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse move events for this item. If you do receive this event, you can be certain that this item also received a mouse press event, and that this item is the current mouse grabber.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

The default implementation handles basic item interaction, such as selection and moving. If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mouseMoveEvent() in your reimplementation.

Please note that mousePressEvent() decides which graphics item it is that receives mouse events. See the mousePressEvent() description for details.

See also mousePressEvent(), mouseReleaseEvent(), mouseDoubleClickEvent(), and sceneEvent().
*/
func (this *QGraphicsTextItem) MouseMoveEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:933
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse release events for this item.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

The default implementation handles basic item interaction, such as selection and moving. If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mouseReleaseEvent() in your reimplementation.

Please note that mousePressEvent() decides which graphics item it is that receives mouse events. See the mousePressEvent() description for details.

See also mousePressEvent(), mouseMoveEvent(), mouseDoubleClickEvent(), and sceneEvent().
*/
func (this *QGraphicsTextItem) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:934
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse doubleclick events for this item.

When doubleclicking an item, the item will first receive a mouse press event, followed by a release event (i.e., a click), then a doubleclick event, and finally a release event.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

The default implementation calls mousePressEvent(). If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mouseDoubleClickEvent() in your reimplementation.

Note that an item will not receive double click events if it is neither selectable nor movable (single mouse clicks are ignored in this case, and that stops the generation of double clicks).

See also mousePressEvent(), mouseMoveEvent(), mouseReleaseEvent(), and sceneEvent().
*/
func (this *QGraphicsTextItem) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:935
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QGraphicsSceneContextMenuEvent *)

/*
This event handler can be reimplemented in a subclass to process context menu events. The event parameter contains details about the event to be handled.

If you ignore the event (i.e., by calling QEvent::ignore()), event will propagate to any item beneath this item. If no items accept the event, it will be ignored by the scene and propagate to the view.

It's common to open a QMenu in response to receiving a context menu event. Example:


  void CustomItem::contextMenuEvent(QGraphicsSceneContextMenuEvent *event)
  {
      QMenu menu;
      QAction *removeAction = menu.addAction("Remove");
      QAction *markAction = menu.addAction("Mark");
      QAction *selectedAction = menu.exec(event->screenPos());
      // ...
  }



The default implementation ignores the event.

See also sceneEvent().
*/
func (this *QGraphicsTextItem) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF /*777 QGraphicsSceneContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneContextMenuEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:936
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
This event handler, for event event, can be reimplemented to receive key press events for this item. The default implementation ignores the event. If you reimplement this handler, the event will by default be accepted.

Note that key events are only received for items that set the ItemIsFocusable flag, and that have keyboard input focus.

See also keyReleaseEvent(), setFocus(), QGraphicsScene::setFocusItem(), and sceneEvent().
*/
func (this *QGraphicsTextItem) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:937
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
This event handler, for event event, can be reimplemented to receive key release events for this item. The default implementation ignores the event. If you reimplement this handler, the event will by default be accepted.

Note that key events are only received for items that set the ItemIsFocusable flag, and that have keyboard input focus.

See also keyPressEvent(), setFocus(), QGraphicsScene::setFocusItem(), and sceneEvent().
*/
func (this *QGraphicsTextItem) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:938
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
This event handler, for event event, can be reimplemented to receive focus in events for this item. The default implementation calls ensureVisible().

See also focusOutEvent(), sceneEvent(), and setFocus().
*/
func (this *QGraphicsTextItem) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:939
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
This event handler, for event event, can be reimplemented to receive focus out events for this item. The default implementation does nothing.

See also focusInEvent(), sceneEvent(), and setFocus().
*/
func (this *QGraphicsTextItem) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:940
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drag enter events for this item. Drag enter events are generated as the cursor enters the item's area.

By accepting the event (i.e., by calling QEvent::accept()), the item will accept drop events, in addition to receiving drag move and drag leave. Otherwise, the event will be ignored and propagate to the item beneath. If the event is accepted, the item will receive a drag move event before control goes back to the event loop.

A common implementation of dragEnterEvent accepts or ignores event depending on the associated mime data in event. Example:


  CustomItem::CustomItem()
  {
      setAcceptDrops(true);
      ...
  }

  void CustomItem::dragEnterEvent(QGraphicsSceneDragDropEvent *event)
  {
      event->setAccepted(event->mimeData()->hasFormat("text/plain"));
  }



Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dropEvent(), dragMoveEvent(), and dragLeaveEvent().
*/
func (this *QGraphicsTextItem) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:941
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drag leave events for this item. Drag leave events are generated as the cursor leaves the item's area. Most often you will not need to reimplement this function, but it can be useful for resetting state in your item (e.g., highlighting).

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dragEnterEvent(), dropEvent(), and dragMoveEvent().
*/
func (this *QGraphicsTextItem) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:942
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drag move events for this item. Drag move events are generated as the cursor moves around inside the item's area. Most often you will not need to reimplement this function; it is used to indicate that only parts of the item can accept drops.

Calling QEvent::ignore() or QEvent::accept() on event toggles whether or not the item will accept drops at the position from the event. By default, event is accepted, indicating that the item allows drops at the specified position.

Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dropEvent(), dragEnterEvent(), and dragLeaveEvent().
*/
func (this *QGraphicsTextItem) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:943
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drop events for this item. Items can only receive drop events if the last drag move event was accepted.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dragEnterEvent(), dragMoveEvent(), and dragLeaveEvent().
*/
func (this *QGraphicsTextItem) DropEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:944
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
This event handler, for event event, can be reimplemented to receive input method events for this item. The default implementation ignores the event.

See also inputMethodQuery() and sceneEvent().
*/
func (this *QGraphicsTextItem) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QInputMethodEvent_PTR() != nil {
		convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:945
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverEnterEvent(QGraphicsSceneHoverEvent *)

/*
This event handler, for event event, can be reimplemented to receive hover enter events for this item. The default implementation calls update(); otherwise it does nothing.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

See also hoverMoveEvent(), hoverLeaveEvent(), sceneEvent(), and setAcceptHoverEvents().
*/
func (this *QGraphicsTextItem) HoverEnterEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:946
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QGraphicsSceneHoverEvent *)

/*
This event handler, for event event, can be reimplemented to receive hover move events for this item. The default implementation does nothing.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

See also hoverEnterEvent(), hoverLeaveEvent(), sceneEvent(), and setAcceptHoverEvents().
*/
func (this *QGraphicsTextItem) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:947
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QGraphicsSceneHoverEvent *)

/*
This event handler, for event event, can be reimplemented to receive hover leave events for this item. The default implementation calls update(); otherwise it does nothing.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

See also hoverEnterEvent(), hoverMoveEvent(), sceneEvent(), and setAcceptHoverEvents().
*/
func (this *QGraphicsTextItem) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:949
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
This method is only relevant for input items. It is used by the input method to query a set of properties of the item to be able to support complex input method operations, such as support for surrounding text and reconversions. query specifies which property is queried.

See also inputMethodEvent() and QInputMethodEvent.
*/
func (this *QGraphicsTextItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:951
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(QGraphicsItem::Extension) const

/*

 */
func (this *QGraphicsTextItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:952
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(QGraphicsItem::Extension, const QVariant &)

/*

 */
func (this *QGraphicsTextItem) SetExtension(extension int, variant qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg1 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:953
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &) const

/*

 */
func (this *QGraphicsTextItem) Extension(variant qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

/*


 */
type QGraphicsTextItem__ = int

//
const QGraphicsTextItem__Type QGraphicsTextItem__ = 8

func (this *QGraphicsTextItem) ItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsTextItem_ItemName(val int) string {
	var nilthis *QGraphicsTextItem
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
