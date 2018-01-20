//  header block begin
// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
	*QGraphicsEffect
}

func (this *QGraphicsColorizeEffect) GetCthis() unsafe.Pointer {
	return this.QGraphicsEffect.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:128
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsColorizeEffect) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:132
// index:0
// void QGraphicsColorizeEffect(class QObject *)
func NewQGraphicsColorizeEffect(parent unsafe.Pointer) *QGraphicsColorizeEffect {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsColorizeEffectFromPointer(cthis)
	return gothis
}
func NewQGraphicsColorizeEffectFromPointer(cthis unsafe.Pointer) *QGraphicsColorizeEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsColorizeEffect{bcthis0}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect5colorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:136
// index:0
// qreal strength()
func (this *QGraphicsColorizeEffect) Strength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsColorizeEffect8strengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:139
// index:0
// void setColor(const class QColor &)
func (this *QGraphicsColorizeEffect) SetColor(c unsafe.Pointer) {
	// 0: (, c const QColor &), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect8setColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:140
// index:0
// void setStrength(qreal)
func (this *QGraphicsColorizeEffect) SetStrength(strength float64) {
	// 0: (, strength qreal), (&strength)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect11setStrengthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &strength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:143
// index:0
// void colorChanged(const class QColor &)
func (this *QGraphicsColorizeEffect) ColorChanged(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect12colorChangedERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:144
// index:0
// void strengthChanged(qreal)
func (this *QGraphicsColorizeEffect) StrengthChanged(strength float64) {
	// 0: (, strength qreal), (&strength)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect15strengthChangedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &strength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:147
// index:0
// virtual
// void draw(class QPainter *)
func (this *QGraphicsColorizeEffect) Draw(painter unsafe.Pointer) {
	// 0: (, painter QPainter *), (painter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsColorizeEffect4drawEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

//  body block end
