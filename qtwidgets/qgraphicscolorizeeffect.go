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
// extern C begin: 16
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
func (this *QGraphicsColorizeEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "draw", f)
}

/*

 */
type QGraphicsColorizeEffect struct {
	*QGraphicsEffect
}
type QGraphicsColorizeEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsColorizeEffect_PTR() *QGraphicsColorizeEffect
}

func (ptr *QGraphicsColorizeEffect) QGraphicsColorizeEffect_PTR() *QGraphicsColorizeEffect { return ptr }

func (this *QGraphicsColorizeEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsEffect.GetCthis()
	}
}
func (this *QGraphicsColorizeEffect) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsEffect = NewQGraphicsEffectFromPointer(cthis)
}
func NewQGraphicsColorizeEffectFromPointer(cthis unsafe.Pointer) *QGraphicsColorizeEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsColorizeEffect{bcthis0}
}
func (*QGraphicsColorizeEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsColorizeEffect {
	return NewQGraphicsColorizeEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsColorizeEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsColorizeEffect(QObject *)

/*

 */
func (*QGraphicsColorizeEffect) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsColorizeEffect {
	return NewQGraphicsColorizeEffect(parent)
}
func NewQGraphicsColorizeEffect(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsColorizeEffect {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsColorizeEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsColorizeEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsColorizeEffect(QObject *)

/*

 */
func (*QGraphicsColorizeEffect) NewForInheritp() *QGraphicsColorizeEffect {
	return NewQGraphicsColorizeEffectp()
}
func NewQGraphicsColorizeEffectp() *QGraphicsColorizeEffect {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsColorizeEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsColorizeEffect")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsColorizeEffect()

/*

 */
func DeleteQGraphicsColorizeEffect(this *QGraphicsColorizeEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:135
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color() const

/*

 */
func (this *QGraphicsColorizeEffect) Color() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal strength() const

/*

 */
func (this *QGraphicsColorizeEffect) Strength() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect8strengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)

/*

 */
func (this *QGraphicsColorizeEffect) SetColor(c qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg0 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStrength(qreal)

/*

 */
func (this *QGraphicsColorizeEffect) SetStrength(strength float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect11setStrengthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), strength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void colorChanged(const QColor &)

/*

 */
func (this *QGraphicsColorizeEffect) ColorChanged(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect12colorChangedERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void strengthChanged(qreal)

/*

 */
func (this *QGraphicsColorizeEffect) StrengthChanged(strength float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect15strengthChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), strength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:147
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
func (this *QGraphicsColorizeEffect) Draw(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect4drawEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
