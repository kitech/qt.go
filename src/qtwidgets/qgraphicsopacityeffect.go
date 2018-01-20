//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		ffiqt.KeepMe()
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
type QGraphicsOpacityEffect struct {
	*QGraphicsEffect
}

func (this *QGraphicsOpacityEffect) GetCthis() unsafe.Pointer {
	return this.QGraphicsEffect.GetCthis()
}
func NewQGraphicsOpacityEffectFromPointer(cthis unsafe.Pointer) *QGraphicsOpacityEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsOpacityEffect{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:254
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsOpacityEffect) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:258
// index:0
// Public
// void QGraphicsOpacityEffect(class QObject *)
func NewQGraphicsOpacityEffect(parent unsafe.Pointer) *QGraphicsOpacityEffect {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsOpacityEffectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:259
// index:0
// Public virtual
// void ~QGraphicsOpacityEffect()
func DeleteQGraphicsOpacityEffect(*QGraphicsOpacityEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:261
// index:0
// Public
// qreal opacity()
func (this *QGraphicsOpacityEffect) Opacity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect7opacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:262
// index:0
// Public
// QBrush opacityMask()
func (this *QGraphicsOpacityEffect) OpacityMask() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect11opacityMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:265
// index:0
// Public
// void setOpacity(qreal)
func (this *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect10setOpacityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:266
// index:0
// Public
// void setOpacityMask(const class QBrush &)
func (this *QGraphicsOpacityEffect) SetOpacityMask(mask *qtgui.QBrush) {
	var convArg0 = mask.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:269
// index:0
// Public
// void opacityChanged(qreal)
func (this *QGraphicsOpacityEffect) OpacityChanged(opacity float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14opacityChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:270
// index:0
// Public
// void opacityMaskChanged(const class QBrush &)
func (this *QGraphicsOpacityEffect) OpacityMaskChanged(mask *qtgui.QBrush) {
	var convArg0 = mask.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect18opacityMaskChangedERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:273
// index:0
// Protected virtual
// void draw(class QPainter *)
func (this *QGraphicsOpacityEffect) Draw(painter unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect4drawEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

//  body block end
