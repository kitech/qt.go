package qtwidgets

// /usr/include/qt/QtWidgets/qstyleplugin.h
// #include <qstyleplugin.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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

type QStylePlugin struct {
	*qtcore.QObject
}

func (this *QStylePlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QStylePlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQStylePluginFromPointer(cthis unsafe.Pointer) *QStylePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStylePlugin{bcthis0}
}
func (*QStylePlugin) NewFromPointer(cthis unsafe.Pointer) *QStylePlugin {
	return NewQStylePluginFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStylePlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QStylePlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStylePlugin(QObject *)
func NewQStylePlugin(parent *qtcore.QObject /*777 QObject **/) *QStylePlugin {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStylePluginC1EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStylePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStylePlugin()
func DeleteQStylePlugin(this *QStylePlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStylePluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStyle * create(const QString &)
func (this *QStylePlugin) Create(key *qtcore.QString) *QStyle /*777 QStyle **/ {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStylePlugin6createERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
