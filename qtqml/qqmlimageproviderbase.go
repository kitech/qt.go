package qtqml

// /usr/include/qt/QtQml/qqmlengine.h
// #include <qqmlengine.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlImageProviderBase struct {
	*qtrt.CObject
}

func (this *QQmlImageProviderBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlImageProviderBase) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlImageProviderBaseFromPointer(cthis unsafe.Pointer) *QQmlImageProviderBase {
	return &QQmlImageProviderBase{&qtrt.CObject{cthis}}
}
func (*QQmlImageProviderBase) NewFromPointer(cthis unsafe.Pointer) *QQmlImageProviderBase {
	return NewQQmlImageProviderBaseFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlengine.h:72
// index:0
// Public virtual
// void ~QQmlImageProviderBase()
func DeleteQQmlImageProviderBase(*QQmlImageProviderBase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QQmlImageProviderBaseD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:74
// index:0
// Public pure virtual
// QQmlImageProviderBase::ImageType imageType()
func (this *QQmlImageProviderBase) ImageType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QQmlImageProviderBase9imageTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:75
// index:0
// Public pure virtual
// QQmlImageProviderBase::Flags flags()
func (this *QQmlImageProviderBase) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QQmlImageProviderBase5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QQmlImageProviderBase__ImageType = int

const QQmlImageProviderBase__Image QQmlImageProviderBase__ImageType = 0
const QQmlImageProviderBase__Pixmap QQmlImageProviderBase__ImageType = 1
const QQmlImageProviderBase__Texture QQmlImageProviderBase__ImageType = 2
const QQmlImageProviderBase__Invalid QQmlImageProviderBase__ImageType = 3
const QQmlImageProviderBase__ImageResponse QQmlImageProviderBase__ImageType = 4

type QQmlImageProviderBase__Flag = int

const QQmlImageProviderBase__ForceAsynchronousImageLoading QQmlImageProviderBase__Flag = 1

//  body block end
