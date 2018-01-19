//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QGraphicsBlurEffect struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:157
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsBlurEffect) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:170
// index:0
// void QGraphicsBlurEffect(class QObject *)
func NewQGraphicsBlurEffect(parent unsafe.Pointer) *QGraphicsBlurEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsBlurEffect{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:171
// index:0
// virtual
// void ~QGraphicsBlurEffect()
func DeleteQGraphicsBlurEffect(*QGraphicsBlurEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:173
// index:0
// virtual
// QRectF boundingRectFor(const class QRectF &)
func (this *QGraphicsBlurEffect) BoundingRectFor(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:174
// index:0
// qreal blurRadius()
func (this *QGraphicsBlurEffect) BlurRadius() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10blurRadiusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:175
// index:0
// QGraphicsBlurEffect::BlurHints blurHints()
func (this *QGraphicsBlurEffect) BlurHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect9blurHintsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:178
// index:0
// void setBlurRadius(qreal)
func (this *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64) {
	// 0: (, qreal blurRadius), (&blurRadius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect13setBlurRadiusEd", ffiqt.FFI_TYPE_VOID, this.cthis, &blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:182
// index:0
// void blurRadiusChanged(qreal)
func (this *QGraphicsBlurEffect) BlurRadiusChanged(blurRadius float64) {
	// 0: (, qreal blurRadius), (&blurRadius)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect17blurRadiusChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &blurRadius)
	gopp.ErrPrint(err, rv)
}

//  body block end
