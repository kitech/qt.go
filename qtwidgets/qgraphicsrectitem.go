// +build !minimal

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
// extern C begin: 15
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

// bool supportsExtension(QGraphicsItem::Extension)
func (this *QGraphicsRectItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsRectItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const QVariant &)
func (this *QGraphicsRectItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

/*

 */
type QGraphicsRectItem struct {
	*QAbstractGraphicsShapeItem
}
type QGraphicsRectItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsRectItem_PTR() *QGraphicsRectItem
}

func (ptr *QGraphicsRectItem) QGraphicsRectItem_PTR() *QGraphicsRectItem { return ptr }

func (this *QGraphicsRectItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractGraphicsShapeItem.GetCthis()
	}
}
func (this *QGraphicsRectItem) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractGraphicsShapeItem = NewQAbstractGraphicsShapeItemFromPointer(cthis)
}
func NewQGraphicsRectItemFromPointer(cthis unsafe.Pointer) *QGraphicsRectItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsRectItem{bcthis0}
}
func (*QGraphicsRectItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsRectItem {
	return NewQGraphicsRectItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:661
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(QGraphicsItem *)

/*

 */
func (*QGraphicsRectItem) NewForInherit(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	return NewQGraphicsRectItem(parent)
}
func NewQGraphicsRectItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:661
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(QGraphicsItem *)

/*

 */
func (*QGraphicsRectItem) NewForInheritp() *QGraphicsRectItem {
	return NewQGraphicsRectItemp()
}
func NewQGraphicsRectItemp() *QGraphicsRectItem {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:662
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(const QRectF &, QGraphicsItem *)

/*

 */
func (*QGraphicsRectItem) NewForInherit1(rect qtcore.QRectF_ITF, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	return NewQGraphicsRectItem1(rect, parent)
}
func NewQGraphicsRectItem1(rect qtcore.QRectF_ITF, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg1 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:662
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(const QRectF &, QGraphicsItem *)

/*

 */
func (*QGraphicsRectItem) NewForInherit1p(rect qtcore.QRectF_ITF) *QGraphicsRectItem {
	return NewQGraphicsRectItem1p(rect)
}
func NewQGraphicsRectItem1p(rect qtcore.QRectF_ITF) *QGraphicsRectItem {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:663
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(qreal, qreal, qreal, qreal, QGraphicsItem *)

/*

 */
func (*QGraphicsRectItem) NewForInherit2(x float64, y float64, w float64, h float64, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	return NewQGraphicsRectItem2(x, y, w, h, parent)
}
func NewQGraphicsRectItem2(x float64, y float64, w float64, h float64, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg4 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:663
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(qreal, qreal, qreal, qreal, QGraphicsItem *)

/*

 */
func (*QGraphicsRectItem) NewForInherit2p(x float64, y float64, w float64, h float64) *QGraphicsRectItem {
	return NewQGraphicsRectItem2p(x, y, w, h)
}
func NewQGraphicsRectItem2p(x float64, y float64, w float64, h float64) *QGraphicsRectItem {
	// arg: 4, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:664
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsRectItem()

/*

 */
func DeleteQGraphicsRectItem(this *QGraphicsRectItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:666
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect() const

/*

 */
func (this *QGraphicsRectItem) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:667
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)

/*

 */
func (this *QGraphicsRectItem) SetRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:668
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)

/*

 */
func (this *QGraphicsRectItem) SetRect1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:670
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
func (this *QGraphicsRectItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:671
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
func (this *QGraphicsRectItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:672
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &) const

/*
Returns true if this item contains point, which is in local coordinates; otherwise, false is returned. It is most often called from QGraphicsView to determine what item is under the cursor, and for that reason, the implementation of this function should be as light-weight as possible.

By default, this function calls shape(), but you can reimplement it in a subclass to provide a (perhaps more efficient) implementation.

See also shape(), boundingRect(), and collidesWithPath().
*/
func (this *QGraphicsRectItem) Contains(point qtcore.QPointF_ITF) bool {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:674
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
func (this *QGraphicsRectItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:674
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
func (this *QGraphicsRectItem) Paintp(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:676
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *) const

/*
Returns true if this item's bounding rect is completely obscured by the opaque shape of item.

The base implementation maps item's opaqueArea() to this item's coordinate system, and then checks if this item's boundingRect() is fully contained within the mapped shape.

You can reimplement this function to provide a custom algorithm for determining whether this item is obscured by item.

See also opaqueArea() and isObscured().
*/
func (this *QGraphicsRectItem) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:677
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea() const

/*
This virtual function returns a shape representing the area where this item is opaque. An area is opaque if it is filled using an opaque brush or color (i.e., not transparent).

This function is used by isObscuredBy(), which is called by underlying items to determine if they are obscured by this item.

The default implementation returns an empty QPainterPath, indicating that this item is completely transparent and does not obscure any other items.

See also isObscuredBy(), isObscured(), and shape().
*/
func (this *QGraphicsRectItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:680
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

     int type() const override
     {
         // Enable the use of qgraphicsitem_cast with this item.
         return Type;
     }
     ...
  };



See also UserType.
*/
func (this *QGraphicsRectItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:683
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(QGraphicsItem::Extension) const

/*

 */
func (this *QGraphicsRectItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:684
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(QGraphicsItem::Extension, const QVariant &)

/*

 */
func (this *QGraphicsRectItem) SetExtension(extension int, variant qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg1 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:685
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &) const

/*

 */
func (this *QGraphicsRectItem) Extension(variant qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

/*


 */
type QGraphicsRectItem__ = int

//
const QGraphicsRectItem__Type QGraphicsRectItem__ = 3

func (this *QGraphicsRectItem) ItemName(val int) string {
	switch val {
	case QGraphicsRectItem__Type: // 3
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsRectItem_ItemName(val int) string {
	var nilthis *QGraphicsRectItem
	return nilthis.ItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11171() {
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
