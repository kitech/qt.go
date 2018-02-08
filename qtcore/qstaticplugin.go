package qtcore

// /usr/include/qt/QtCore/qplugin.h
// #include <qplugin.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QStaticPlugin struct {
	*qtrt.CObject
}

func (this *QStaticPlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStaticPlugin) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStaticPluginFromPointer(cthis unsafe.Pointer) *QStaticPlugin {
	return &QStaticPlugin{&qtrt.CObject{cthis}}
}
func (*QStaticPlugin) NewFromPointer(cthis unsafe.Pointer) *QStaticPlugin {
	return NewQStaticPluginFromPointer(cthis)
}

// /usr/include/qt/QtCore/qplugin.h:74
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject metaData()
func (this *QStaticPlugin) MetaData() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStaticPlugin8metaDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

func DeleteQStaticPlugin(this *QStaticPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStaticPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
