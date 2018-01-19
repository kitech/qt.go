//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:254
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsOpacityEffect) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:258
// index:0
// void QGraphicsOpacityEffect(class QObject *)
func NewQGraphicsOpacityEffect(parent unsafe.Pointer) *QGraphicsOpacityEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsOpacityEffect{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:259
// index:0
// virtual
// void ~QGraphicsOpacityEffect()
func DeleteQGraphicsOpacityEffect(*QGraphicsOpacityEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:261
// index:0
// qreal opacity()
func (this *QGraphicsOpacityEffect) Opacity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect7opacityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:262
// index:0
// QBrush opacityMask()
func (this *QGraphicsOpacityEffect) OpacityMask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect11opacityMaskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:265
// index:0
// void setOpacity(qreal)
func (this *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	// 0: (, qreal opacity), (&opacity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect10setOpacityEd", ffiqt.FFI_TYPE_VOID, this.cthis, &opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:266
// index:0
// void setOpacityMask(const class QBrush &)
func (this *QGraphicsOpacityEffect) SetOpacityMask(mask unsafe.Pointer) {
	// 0: (, const QBrush & mask), (mask)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, mask)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:269
// index:0
// void opacityChanged(qreal)
func (this *QGraphicsOpacityEffect) OpacityChanged(opacity float64) {
	// 0: (, qreal opacity), (&opacity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14opacityChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:270
// index:0
// void opacityMaskChanged(const class QBrush &)
func (this *QGraphicsOpacityEffect) OpacityMaskChanged(mask unsafe.Pointer) {
	// 0: (, const QBrush & mask), (mask)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect18opacityMaskChangedERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, mask)
	gopp.ErrPrint(err, rv)
}

//  body block end
