package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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

// void draw(QPainter *)
func (this *QGraphicsDropShadowEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "draw", f)
}

/*

 */
type QGraphicsDropShadowEffect struct {
	*QGraphicsEffect
}
type QGraphicsDropShadowEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsDropShadowEffect_PTR() *QGraphicsDropShadowEffect
}

func (ptr *QGraphicsDropShadowEffect) QGraphicsDropShadowEffect_PTR() *QGraphicsDropShadowEffect {
	return ptr
}

func (this *QGraphicsDropShadowEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsEffect.GetCthis()
	}
}
func (this *QGraphicsDropShadowEffect) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsEffect = NewQGraphicsEffectFromPointer(cthis)
}
func NewQGraphicsDropShadowEffectFromPointer(cthis unsafe.Pointer) *QGraphicsDropShadowEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsDropShadowEffect{bcthis0}
}
func (*QGraphicsDropShadowEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsDropShadowEffect {
	return NewQGraphicsDropShadowEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:198
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsDropShadowEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsDropShadowEffect(QObject *)

/*

 */
func NewQGraphicsDropShadowEffect(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsDropShadowEffect {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsDropShadowEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsDropShadowEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsDropShadowEffect(QObject *)

/*

 */
func NewQGraphicsDropShadowEffect__() *QGraphicsDropShadowEffect {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsDropShadowEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsDropShadowEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:206
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsDropShadowEffect()

/*

 */
func DeleteQGraphicsDropShadowEffect(this *QGraphicsDropShadowEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:208
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRectFor(const QRectF &) const

/*
Returns the effective bounding rectangle for this effect, given the provided rect in the device coordinates. When writing you own custom effect, you must call updateBoundingRect() whenever any parameters are changed that may cause this this function to return a different value.

See also sourceBoundingRect().
*/
func (this *QGraphicsDropShadowEffect) BoundingRectFor(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:209
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF offset() const

/*

 */
func (this *QGraphicsDropShadowEffect) Offset() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect6offsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:211
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal xOffset() const

/*

 */
func (this *QGraphicsDropShadowEffect) XOffset() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect7xOffsetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal yOffset() const

/*

 */
func (this *QGraphicsDropShadowEffect) YOffset() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect7yOffsetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:217
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal blurRadius() const

/*

 */
func (this *QGraphicsDropShadowEffect) BlurRadius() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect10blurRadiusEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:218
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color() const

/*

 */
func (this *QGraphicsDropShadowEffect) Color() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QGraphicsDropShadowEffect5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffset(const QPointF &)

/*

 */
func (this *QGraphicsDropShadowEffect) SetOffset(ofs qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if ofs != nil && ofs.QPointF_PTR() != nil {
		convArg0 = ofs.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:223
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setOffset(qreal, qreal)

/*

 */
func (this *QGraphicsDropShadowEffect) SetOffset_1(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:226
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void setOffset(qreal)

/*

 */
func (this *QGraphicsDropShadowEffect) SetOffset_2(d float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect9setOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), d)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setXOffset(qreal)

/*

 */
func (this *QGraphicsDropShadowEffect) SetXOffset(dx float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect10setXOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setYOffset(qreal)

/*

 */
func (this *QGraphicsDropShadowEffect) SetYOffset(dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect10setYOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlurRadius(qreal)

/*

 */
func (this *QGraphicsDropShadowEffect) SetBlurRadius(blurRadius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect13setBlurRadiusEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blurRadius)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)

/*

 */
func (this *QGraphicsDropShadowEffect) SetColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void offsetChanged(const QPointF &)

/*

 */
func (this *QGraphicsDropShadowEffect) OffsetChanged(offset qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPointF_PTR() != nil {
		convArg0 = offset.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect13offsetChangedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void blurRadiusChanged(qreal)

/*

 */
func (this *QGraphicsDropShadowEffect) BlurRadiusChanged(blurRadius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect17blurRadiusChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blurRadius)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void colorChanged(const QColor &)

/*

 */
func (this *QGraphicsDropShadowEffect) ColorChanged(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect12colorChangedERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void draw(QPainter *)

/*
This pure virtual function draws the effect and is called whenever the source needs to be drawn.

Reimplement this function in a QGraphicsEffect subclass to provide the effect's drawing implementation, using painter.

For example:


  MyGraphicsEffect::draw(QPainter *painter)
  {
      ...
      QPoint offset;
      if (sourceIsPixmap()) {
          // No point in drawing in device coordinates (pixmap will be scaled anyways).
          const QPixmap pixmap = sourcePixmap(Qt::LogicalCoordinates, &offset);
          ...
          painter->drawPixmap(offset, pixmap);
      } else {
          // Draw pixmap in device coordinates to avoid pixmap scaling;
          const QPixmap pixmap = sourcePixmap(Qt::DeviceCoordinates, &offset);
          painter->setWorldTransform(QTransform());
          ...
          painter->drawPixmap(offset, pixmap);
      }
      ...
  }



This function should not be called explicitly by the user, since it is meant for reimplementation purposes only.
*/
func (this *QGraphicsDropShadowEffect) Draw(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QGraphicsDropShadowEffect4drawEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
