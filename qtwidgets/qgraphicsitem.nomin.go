//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

// /usr/include/qt/QtWidgets/qgraphicsitem.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEffect * graphicsEffect() const

/*
Returns a pointer to this item's effect if it has one; otherwise 0.

This function was introduced in  Qt 4.6.

See also setGraphicsEffect().
*/
func (this *QGraphicsItem) GraphicsEffect() *QGraphicsEffect /*777 QGraphicsEffect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14graphicsEffectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGraphicsEffect(QGraphicsEffect *)

/*
Sets effect as the item's effect. If there already is an effect installed on this item, QGraphicsItem will delete the existing effect before installing the new effect. You can delete an existing effect by calling setGraphicsEffect(0).

If effect is the installed effect on a different item, setGraphicsEffect() will remove the effect from the item and install it on this item.

QGraphicsItem takes ownership of effect.

Note: This function will apply the effect on itself and all its children.

This function was introduced in  Qt 4.6.

See also graphicsEffect().
*/
func (this *QGraphicsItem) SetGraphicsEffect(effect QGraphicsEffect_ITF /*777 QGraphicsEffect **/) {
	var convArg0 unsafe.Pointer
	if effect != nil && effect.QGraphicsEffect_PTR() != nil {
		convArg0 = effect.QGraphicsEffect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11166() {
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
