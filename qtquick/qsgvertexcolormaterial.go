package qtquick

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h
// #include <qsgvertexcolormaterial.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

// QSGMaterialType * type()
func (this *QSGVertexColorMaterial) InheritType(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "type", f)
}

// QSGMaterialShader * createShader()
func (this *QSGVertexColorMaterial) InheritCreateShader(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createShader", f)
}

/*

 */
type QSGVertexColorMaterial struct {
	*QSGMaterial
}
type QSGVertexColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial
}

func (ptr *QSGVertexColorMaterial) QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial { return ptr }

func (this *QSGVertexColorMaterial) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGMaterial.GetCthis()
	}
}
func (this *QSGVertexColorMaterial) SetCthis(cthis unsafe.Pointer) {
	this.QSGMaterial = NewQSGMaterialFromPointer(cthis)
}
func NewQSGVertexColorMaterialFromPointer(cthis unsafe.Pointer) *QSGVertexColorMaterial {
	bcthis0 := NewQSGMaterialFromPointer(cthis)
	return &QSGVertexColorMaterial{bcthis0}
}
func (*QSGVertexColorMaterial) NewFromPointer(cthis unsafe.Pointer) *QSGVertexColorMaterial {
	return NewQSGVertexColorMaterialFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h:50
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGVertexColorMaterial()

/*
Creates a new vertex color material.
*/
func (*QSGVertexColorMaterial) NewForInherit() *QSGVertexColorMaterial {
	return NewQSGVertexColorMaterial()
}
func NewQSGVertexColorMaterial() *QSGVertexColorMaterial {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QSGVertexColorMaterialC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGVertexColorMaterialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGVertexColorMaterial)
	return gothis
}

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int compare(const QSGMaterial *) const

/*

 */
func (this *QSGVertexColorMaterial) Compare(other QSGMaterial_ITF /*777 const QSGMaterial **/) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSGMaterial_PTR() != nil {
		convArg0 = other.QSGMaterial_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QSGVertexColorMaterial7compareEPK11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h:55
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type() const

/*

 */
func (this *QSGVertexColorMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QSGVertexColorMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h:56
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader() const

/*

 */
func (this *QSGVertexColorMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QSGVertexColorMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQSGVertexColorMaterial(this *QSGVertexColorMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QSGVertexColorMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11625() {
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
