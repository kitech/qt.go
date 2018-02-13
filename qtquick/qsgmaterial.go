package qtquick

// /usr/include/qt/QtQuick/qsgmaterial.h
// #include <qsgmaterial.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QSGMaterial struct {
	*qtrt.CObject
}
type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (ptr *QSGMaterial) QSGMaterial_PTR() *QSGMaterial { return ptr }

func (this *QSGMaterial) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGMaterial) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSGMaterialFromPointer(cthis unsafe.Pointer) *QSGMaterial {
	return &QSGMaterial{&qtrt.CObject{cthis}}
}
func (*QSGMaterial) NewFromPointer(cthis unsafe.Pointer) *QSGMaterial {
	return NewQSGMaterialFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGMaterial()
func NewQSGMaterial() *QSGMaterial {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGMaterialC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGMaterial)
	return gothis
}

// /usr/include/qt/QtQuick/qsgmaterial.h:147
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGMaterial()
func DeleteQSGMaterial(this *QSGMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:149
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type()
func (this *QSGMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgmaterial.h:150
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader()
func (this *QSGMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgmaterial.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int compare(const QSGMaterial *)
func (this *QSGMaterial) Compare(other *QSGMaterial /*777 const QSGMaterial **/) int {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial7compareEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgmaterial.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGMaterial::Flags flags()
func (this *QSGMaterial) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QSGMaterial::Flags, _Bool)
func (this *QSGMaterial) SetFlag(flags int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGMaterial7setFlagE6QFlagsINS_4FlagEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, on)
	qtrt.ErrPrint(err, rv)
}

type QSGMaterial__Flag = int

const QSGMaterial__Blending QSGMaterial__Flag = 1
const QSGMaterial__RequiresDeterminant QSGMaterial__Flag = 2
const QSGMaterial__RequiresFullMatrixExceptTranslate QSGMaterial__Flag = 6
const QSGMaterial__RequiresFullMatrix QSGMaterial__Flag = 14
const QSGMaterial__CustomCompileStep QSGMaterial__Flag = 16

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
