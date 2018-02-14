package qtwidgets

// /usr/include/qt/QtWidgets/qitemeditorfactory.h
// #include <qitemeditorfactory.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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

type QItemEditorCreatorBase struct {
	*qtrt.CObject
}
type QItemEditorCreatorBase_ITF interface {
	QItemEditorCreatorBase_PTR() *QItemEditorCreatorBase
}

func (ptr *QItemEditorCreatorBase) QItemEditorCreatorBase_PTR() *QItemEditorCreatorBase { return ptr }

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
	rv, err := qtrt.InvokeQtFunc6("_ZN22QItemEditorCreatorBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QWidget * createWidget(QWidget *)
func (this *QItemEditorCreatorBase) CreateWidget(parent QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QItemEditorCreatorBase12createWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QByteArray valuePropertyName()
func (this *QItemEditorCreatorBase) ValuePropertyName() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QItemEditorCreatorBase17valuePropertyNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

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
