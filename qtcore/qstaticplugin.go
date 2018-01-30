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
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
	this.CObject = &qtrt.CObject{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStaticPlugin8metaDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
