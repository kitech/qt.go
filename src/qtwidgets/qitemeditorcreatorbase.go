//  header block begin
// /usr/include/qt/QtWidgets/qitemeditorfactory.h
// #include <qitemeditorfactory.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QItemEditorCreatorBase struct {
	*qtrt.CObject
}

func (this *QItemEditorCreatorBase) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQItemEditorCreatorBaseFromPointer(cthis unsafe.Pointer) *QItemEditorCreatorBase {
	return &QItemEditorCreatorBase{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:58
// index:0
// Public virtual
// void ~QItemEditorCreatorBase()
func DeleteQItemEditorCreatorBase(*QItemEditorCreatorBase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QItemEditorCreatorBaseD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:60
// index:0
// Public pure virtual
// QWidget * createWidget(class QWidget *)
func (this *QItemEditorCreatorBase) CreateWidget(parent unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QItemEditorCreatorBase12createWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:61
// index:0
// Public pure virtual
// QByteArray valuePropertyName()
func (this *QItemEditorCreatorBase) ValuePropertyName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QItemEditorCreatorBase17valuePropertyNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
