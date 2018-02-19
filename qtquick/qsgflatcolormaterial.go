package qtquick

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h
// #include <qsgflatcolormaterial.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QSGFlatColorMaterial struct {
	*QSGMaterial
}
type QSGFlatColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial
}

func (ptr *QSGFlatColorMaterial) QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial { return ptr }

func (this *QSGFlatColorMaterial) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGMaterial.GetCthis()
	}
}
func (this *QSGFlatColorMaterial) SetCthis(cthis unsafe.Pointer) {
	this.QSGMaterial = NewQSGMaterialFromPointer(cthis)
}
func NewQSGFlatColorMaterialFromPointer(cthis unsafe.Pointer) *QSGFlatColorMaterial {
	bcthis0 := NewQSGMaterialFromPointer(cthis)
	return &QSGFlatColorMaterial{bcthis0}
}
func (*QSGFlatColorMaterial) NewFromPointer(cthis unsafe.Pointer) *QSGFlatColorMaterial {
	return NewQSGFlatColorMaterialFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:51
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGFlatColorMaterial()
func NewQSGFlatColorMaterial() *QSGFlatColorMaterial {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGFlatColorMaterialC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGFlatColorMaterialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGFlatColorMaterial)
	return gothis
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type() const
func (this *QSGFlatColorMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader() const
func (this *QSGFlatColorMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QSGFlatColorMaterial) SetColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGFlatColorMaterial8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:56
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color() const
func (this *QSGFlatColorMaterial) Color() *qtgui.QColor {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int compare(const QSGMaterial *) const
func (this *QSGFlatColorMaterial) Compare(other QSGMaterial_ITF /*777 const QSGMaterial **/) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSGMaterial_PTR() != nil {
		convArg0 = other.QSGMaterial_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial7compareEPK11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQSGFlatColorMaterial(this *QSGFlatColorMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGFlatColorMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
