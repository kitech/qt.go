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
// extern C begin: 20
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
func (this *QGraphicsOpacityEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "draw", f)
}

/*

 */
type QGraphicsOpacityEffect struct {
	*QGraphicsEffect
}
type QGraphicsOpacityEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsOpacityEffect_PTR() *QGraphicsOpacityEffect
}

func (ptr *QGraphicsOpacityEffect) QGraphicsOpacityEffect_PTR() *QGraphicsOpacityEffect { return ptr }

func (this *QGraphicsOpacityEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsEffect.GetCthis()
	}
}
func (this *QGraphicsOpacityEffect) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsEffect = NewQGraphicsEffectFromPointer(cthis)
}
func NewQGraphicsOpacityEffectFromPointer(cthis unsafe.Pointer) *QGraphicsOpacityEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsOpacityEffect{bcthis0}
}
func (*QGraphicsOpacityEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsOpacityEffect {
	return NewQGraphicsOpacityEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:254
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsOpacityEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsOpacityEffect(QObject *)

/*

 */
func (*QGraphicsOpacityEffect) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsOpacityEffect {
	return NewQGraphicsOpacityEffect(parent)
}
func NewQGraphicsOpacityEffect(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsOpacityEffect {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsOpacityEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsOpacityEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsOpacityEffect(QObject *)

/*

 */
func (*QGraphicsOpacityEffect) NewForInheritp() *QGraphicsOpacityEffect {
	return NewQGraphicsOpacityEffectp()
}
func NewQGraphicsOpacityEffectp() *QGraphicsOpacityEffect {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsOpacityEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsOpacityEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:259
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsOpacityEffect()

/*

 */
func DeleteQGraphicsOpacityEffect(this *QGraphicsOpacityEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:261
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity() const

/*

 */
func (this *QGraphicsOpacityEffect) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:262
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush opacityMask() const

/*

 */
func (this *QGraphicsOpacityEffect) OpacityMask() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect11opacityMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)

/*

 */
func (this *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacityMask(const QBrush &)

/*

 */
func (this *QGraphicsOpacityEffect) SetOpacityMask(mask qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if mask != nil && mask.QBrush_PTR() != nil {
		convArg0 = mask.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityChanged(qreal)

/*

 */
func (this *QGraphicsOpacityEffect) OpacityChanged(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14opacityChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityMaskChanged(const QBrush &)

/*

 */
func (this *QGraphicsOpacityEffect) OpacityMaskChanged(mask qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if mask != nil && mask.QBrush_PTR() != nil {
		convArg0 = mask.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect18opacityMaskChangedERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:273
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
func (this *QGraphicsOpacityEffect) Draw(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect4drawEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
