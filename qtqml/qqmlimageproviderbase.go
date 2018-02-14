package qtqml

// /usr/include/qt/QtQml/qqmlengine.h
// #include <qqmlengine.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlImageProviderBase struct {
	*qtrt.CObject
}
type QQmlImageProviderBase_ITF interface {
	QQmlImageProviderBase_PTR() *QQmlImageProviderBase
}

func (ptr *QQmlImageProviderBase) QQmlImageProviderBase_PTR() *QQmlImageProviderBase { return ptr }

func (this *QQmlImageProviderBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlImageProviderBase) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlImageProviderBaseFromPointer(cthis unsafe.Pointer) *QQmlImageProviderBase {
	return &QQmlImageProviderBase{&qtrt.CObject{cthis}}
}
func (*QQmlImageProviderBase) NewFromPointer(cthis unsafe.Pointer) *QQmlImageProviderBase {
	return NewQQmlImageProviderBaseFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlengine.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlImageProviderBase()
func DeleteQQmlImageProviderBase(this *QQmlImageProviderBase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlImageProviderBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlengine.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QQmlImageProviderBase::ImageType imageType()
func (this *QQmlImageProviderBase) ImageType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQmlImageProviderBase9imageTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:75
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QQmlImageProviderBase::Flags flags()
func (this *QQmlImageProviderBase) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQmlImageProviderBase5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
