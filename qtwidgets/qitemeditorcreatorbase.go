package qtwidgets

// /usr/include/qt/QtWidgets/qitemeditorfactory.h
// #include <qitemeditorfactory.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QItemEditorCreatorBase) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQItemEditorCreatorBaseFromPointer(cthis unsafe.Pointer) *QItemEditorCreatorBase {
	return &QItemEditorCreatorBase{&qtrt.CObject{cthis}}
}
func (*QItemEditorCreatorBase) NewFromPointer(cthis unsafe.Pointer) *QItemEditorCreatorBase {
	return NewQItemEditorCreatorBaseFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QItemEditorCreatorBase()
func DeleteQItemEditorCreatorBase(this *QItemEditorCreatorBase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QItemEditorCreatorBaseD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QWidget * createWidget(QWidget *)
func (this *QItemEditorCreatorBase) CreateWidget(parent *QWidget /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QItemEditorCreatorBase12createWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QByteArray valuePropertyName()
func (this *QItemEditorCreatorBase) ValuePropertyName() *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QItemEditorCreatorBase17valuePropertyNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

//  body block end
