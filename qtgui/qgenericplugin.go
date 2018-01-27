package qtgui

// /usr/include/qt/QtGui/qgenericplugin.h
// #include <qgenericplugin.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QGenericPlugin struct {
	*qtcore.QObject
}

func (this *QGenericPlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGenericPlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGenericPluginFromPointer(cthis unsafe.Pointer) *QGenericPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGenericPlugin{bcthis0}
}
func (*QGenericPlugin) NewFromPointer(cthis unsafe.Pointer) *QGenericPlugin {
	return NewQGenericPluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qgenericplugin.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGenericPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGenericPlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qgenericplugin.h:55
// index:0
// Public
// void QGenericPlugin(QObject *)
func NewQGenericPlugin(parent *qtcore.QObject /*777 QObject **/) *QGenericPlugin {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGenericPluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGenericPluginFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qgenericplugin.h:56
// index:0
// Public virtual
// void ~QGenericPlugin()
func DeleteQGenericPlugin(*QGenericPlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGenericPluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qgenericplugin.h:58
// index:0
// Public pure virtual
// QObject * create(const QString &, const QString &)
func (this *QGenericPlugin) Create(name *qtcore.QString, spec *qtcore.QString) *qtcore.QObject /*777 QObject **/ {
	var convArg0 = name.GetCthis()
	var convArg1 = spec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGenericPlugin6createERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
