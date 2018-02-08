package qtwidgets

// /usr/include/qt/QtWidgets/qitemeditorfactory.h
// #include <qitemeditorfactory.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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

type QItemEditorFactory struct {
	*qtrt.CObject
}

func (this *QItemEditorFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QItemEditorFactory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQItemEditorFactoryFromPointer(cthis unsafe.Pointer) *QItemEditorFactory {
	return &QItemEditorFactory{&qtrt.CObject{cthis}}
}
func (*QItemEditorFactory) NewFromPointer(cthis unsafe.Pointer) *QItemEditorFactory {
	return NewQItemEditorFactoryFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QItemEditorFactory()
func NewQItemEditorFactory() *QItemEditorFactory {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactoryC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemEditorFactory)
	return gothis
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QItemEditorFactory()
func DeleteQItemEditorFactory(this *QItemEditorFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(int, QWidget *)
func (this *QItemEditorFactory) CreateEditor(userType int, parent *QWidget /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QItemEditorFactory12createEditorEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), userType, convArg1)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:102
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QByteArray valuePropertyName(int)
func (this *QItemEditorFactory) ValuePropertyName(userType int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QItemEditorFactory17valuePropertyNameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), userType)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void registerEditor(int, QItemEditorCreatorBase *)
func (this *QItemEditorFactory) RegisterEditor(userType int, creator *QItemEditorCreatorBase /*777 QItemEditorCreatorBase **/) {
	var convArg1 = creator.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase", qtrt.FFI_TYPE_POINTER, this.GetCthis(), userType, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [8] const QItemEditorFactory * defaultFactory()
func (this *QItemEditorFactory) DefaultFactory() *QItemEditorFactory /*777 const QItemEditorFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactory14defaultFactoryEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QItemEditorFactory_DefaultFactory() *QItemEditorFactory /*777 const QItemEditorFactory **/ {
	var nilthis *QItemEditorFactory
	rv := nilthis.DefaultFactory()
	return rv
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:107
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultFactory(QItemEditorFactory *)
func (this *QItemEditorFactory) SetDefaultFactory(factory *QItemEditorFactory /*777 QItemEditorFactory **/) {
	var convArg0 = factory.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactory17setDefaultFactoryEPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QItemEditorFactory_SetDefaultFactory(factory *QItemEditorFactory /*777 QItemEditorFactory **/) {
	var nilthis *QItemEditorFactory
	nilthis.SetDefaultFactory(factory)
}

//  body block end
