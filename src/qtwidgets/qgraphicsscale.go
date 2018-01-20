//  header block begin
// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QGraphicsScale struct {
	*QGraphicsTransform
}

func (this *QGraphicsScale) GetCthis() unsafe.Pointer {
	return this.QGraphicsTransform.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:81
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsScale) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:88
// index:0
// void QGraphicsScale(class QObject *)
func NewQGraphicsScale(parent unsafe.Pointer) *QGraphicsScale {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScaleC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsScaleFromPointer(cthis)
	return gothis
}
func NewQGraphicsScaleFromPointer(cthis unsafe.Pointer) *QGraphicsScale {
	bcthis0 := NewQGraphicsTransformFromPointer(cthis)
	return &QGraphicsScale{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:89
// index:0
// virtual
// void ~QGraphicsScale()
func DeleteQGraphicsScale(*QGraphicsScale) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScaleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:91
// index:0
// QVector3D origin()
func (this *QGraphicsScale) Origin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6originEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:92
// index:0
// void setOrigin(const class QVector3D &)
func (this *QGraphicsScale) SetOrigin(point unsafe.Pointer) {
	// 0: (, point const QVector3D &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setOriginERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:94
// index:0
// qreal xScale()
func (this *QGraphicsScale) XScale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6xScaleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:95
// index:0
// void setXScale(qreal)
func (this *QGraphicsScale) SetXScale(arg0 float64) {
	// 0: (, qreal arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setXScaleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:97
// index:0
// qreal yScale()
func (this *QGraphicsScale) YScale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6yScaleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:98
// index:0
// void setYScale(qreal)
func (this *QGraphicsScale) SetYScale(arg0 float64) {
	// 0: (, qreal arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setYScaleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:100
// index:0
// qreal zScale()
func (this *QGraphicsScale) ZScale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6zScaleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:101
// index:0
// void setZScale(qreal)
func (this *QGraphicsScale) SetZScale(arg0 float64) {
	// 0: (, qreal arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setZScaleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:103
// index:0
// virtual
// void applyTo(class QMatrix4x4 *)
func (this *QGraphicsScale) ApplyTo(matrix unsafe.Pointer) {
	// 0: (, matrix QMatrix4x4 *), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale7applyToEP10QMatrix4x4", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:106
// index:0
// void originChanged()
func (this *QGraphicsScale) OriginChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13originChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:107
// index:0
// void xScaleChanged()
func (this *QGraphicsScale) XScaleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13xScaleChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:108
// index:0
// void yScaleChanged()
func (this *QGraphicsScale) YScaleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13yScaleChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:109
// index:0
// void zScaleChanged()
func (this *QGraphicsScale) ZScaleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13zScaleChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:110
// index:0
// void scaleChanged()
func (this *QGraphicsScale) ScaleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale12scaleChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
