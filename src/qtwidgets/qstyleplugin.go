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
// Public virtual
// const QMetaObject * metaObject()
func (this *QStylePlugin) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStylePlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:58
// index:0
// Public
// void QStylePlugin(class QObject *)
func NewQStylePlugin(parent *qtcore.QObject /*444 QObject **/) *QStylePlugin {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStylePluginFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:59
// index:0
// Public virtual
// void ~QStylePlugin()
func DeleteQStylePlugin(*QStylePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:61
// index:0
// Public pure virtual
// QStyle * create(const class QString &)
func (this *QStylePlugin) Create(key *qtcore.QString) *QStyle /*444 QStyle **/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePlugin6createERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
