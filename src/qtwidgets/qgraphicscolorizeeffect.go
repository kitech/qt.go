//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QGraphicsColorizeEffect struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:128
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsColorizeEffect) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:132
// index:0
// void QGraphicsColorizeEffect(class QObject *)
func NewQGraphicsColorizeEffect(parent unsafe.Pointer) *QGraphicsColorizeEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsColorizeEffect{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:133
// index:0
// virtual
// void ~QGraphicsColorizeEffect()
func DeleteQGraphicsColorizeEffect(*QGraphicsColorizeEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:135
// index:0
// QColor color()
func (this *QGraphicsColorizeEffect) Color() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect5colorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:136
// index:0
// qreal strength()
func (this *QGraphicsColorizeEffect) Strength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect8strengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:139
// index:0
// void setColor(const class QColor &)
func (this *QGraphicsColorizeEffect) SetColor(c unsafe.Pointer) {
	// 0: (, const QColor & c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect8setColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:140
// index:0
// void setStrength(qreal)
func (this *QGraphicsColorizeEffect) SetStrength(strength float64) {
	// 0: (, qreal strength), (&strength)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect11setStrengthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &strength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:143
// index:0
// void colorChanged(const class QColor &)
func (this *QGraphicsColorizeEffect) ColorChanged(color unsafe.Pointer) {
	// 0: (, const QColor & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect12colorChangedERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:144
// index:0
// void strengthChanged(qreal)
func (this *QGraphicsColorizeEffect) StrengthChanged(strength float64) {
	// 0: (, qreal strength), (&strength)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect15strengthChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &strength)
	gopp.ErrPrint(err, rv)
}

//  body block end
