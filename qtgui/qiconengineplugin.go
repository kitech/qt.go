package qtgui

// /usr/include/qt/QtGui/qiconengineplugin.h
// #include <qiconengineplugin.h>
// #include <QtGui>

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
type QIconEnginePlugin struct {
	*qtcore.QObject
}

func (this *QIconEnginePlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QIconEnginePlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQIconEnginePluginFromPointer(cthis unsafe.Pointer) *QIconEnginePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QIconEnginePlugin{bcthis0}
}
func (*QIconEnginePlugin) NewFromPointer(cthis unsafe.Pointer) *QIconEnginePlugin {
	return NewQIconEnginePluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qiconengineplugin.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QIconEnginePlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QIconEnginePlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qiconengineplugin.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIconEnginePlugin(QObject *)
func NewQIconEnginePlugin(parent *qtcore.QObject /*777 QObject **/) *QIconEnginePlugin {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIconEnginePluginC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconEnginePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qiconengineplugin.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIconEnginePlugin()
func DeleteQIconEnginePlugin(*QIconEnginePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIconEnginePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengineplugin.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIconEngine * create(const QString &)
func (this *QIconEnginePlugin) Create(filename *qtcore.QString) *QIconEngine /*777 QIconEngine **/ {
	var convArg0 = filename.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIconEnginePlugin6createERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIconEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
