package qtquick

// /usr/include/qt/QtQuick/qsgtexturematerial.h
// #include <qsgtexturematerial.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QSGTextureMaterial struct {
	*QSGOpaqueTextureMaterial
}
type QSGTextureMaterial_ITF interface {
	QSGOpaqueTextureMaterial_ITF
	QSGTextureMaterial_PTR() *QSGTextureMaterial
}

func (ptr *QSGTextureMaterial) QSGTextureMaterial_PTR() *QSGTextureMaterial { return ptr }

func (this *QSGTextureMaterial) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGOpaqueTextureMaterial.GetCthis()
	}
}
func (this *QSGTextureMaterial) SetCthis(cthis unsafe.Pointer) {
	this.QSGOpaqueTextureMaterial = NewQSGOpaqueTextureMaterialFromPointer(cthis)
}
func NewQSGTextureMaterialFromPointer(cthis unsafe.Pointer) *QSGTextureMaterial {
	bcthis0 := NewQSGOpaqueTextureMaterialFromPointer(cthis)
	return &QSGTextureMaterial{bcthis0}
}
func (*QSGTextureMaterial) NewFromPointer(cthis unsafe.Pointer) *QSGTextureMaterial {
	return NewQSGTextureMaterialFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type()
func (this *QSGTextureMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSGTextureMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader()
func (this *QSGTextureMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSGTextureMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQSGTextureMaterial(this *QSGTextureMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSGTextureMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
