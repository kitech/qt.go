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
// extern C begin: 10
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

// void draw(class QPainter *)
func (this *QGraphicsBlurEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "draw", f)
}

type QGraphicsBlurEffect struct {
	*QGraphicsEffect
}
type QGraphicsBlurEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsBlurEffect_PTR() *QGraphicsBlurEffect
}

func (ptr *QGraphicsBlurEffect) QGraphicsBlurEffect_PTR() *QGraphicsBlurEffect { return ptr }

func (this *QGraphicsBlurEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsEffect.GetCthis()
	}
}
func (this *QGraphicsBlurEffect) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsEffect = NewQGraphicsEffectFromPointer(cthis)
}
func NewQGraphicsBlurEffectFromPointer(cthis unsafe.Pointer) *QGraphicsBlurEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsBlurEffect{bcthis0}
}
func (*QGraphicsBlurEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsBlurEffect {
	return NewQGraphicsBlurEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsBlurEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsBlurEffect(QObject *)
func NewQGraphicsBlurEffect(parent *qtcore.QObject /*777 QObject **/) *QGraphicsBlurEffect {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsBlurEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:171
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsBlurEffect()
func DeleteQGraphicsBlurEffect(this *QGraphicsBlurEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:173
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRectFor(const QRectF &)
func (this *QGraphicsBlurEffect) BoundingRectFor(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:174
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal blurRadius()
func (this *QGraphicsBlurEffect) BlurRadius() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10blurRadiusEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:175
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsBlurEffect::BlurHints blurHints()
func (this *QGraphicsBlurEffect) BlurHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect9blurHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlurRadius(qreal)
func (this *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect13setBlurRadiusEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blurRadius)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlurHints(QGraphicsBlurEffect::BlurHints)
func (this *QGraphicsBlurEffect) SetBlurHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect12setBlurHintsE6QFlagsINS_8BlurHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void blurRadiusChanged(qreal)
func (this *QGraphicsBlurEffect) BlurRadiusChanged(blurRadius float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect17blurRadiusChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blurRadius)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void blurHintsChanged(QGraphicsBlurEffect::BlurHints)
func (this *QGraphicsBlurEffect) BlurHintsChanged(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect16blurHintsChangedE6QFlagsINS_8BlurHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:186
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void draw(QPainter *)
func (this *QGraphicsBlurEffect) Draw(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect4drawEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QGraphicsBlurEffect__BlurHint = int

const QGraphicsBlurEffect__PerformanceHint QGraphicsBlurEffect__BlurHint = 0
const QGraphicsBlurEffect__QualityHint QGraphicsBlurEffect__BlurHint = 1
const QGraphicsBlurEffect__AnimationHint QGraphicsBlurEffect__BlurHint = 2

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
