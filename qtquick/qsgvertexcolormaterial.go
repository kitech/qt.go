package qtquick

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h
// #include <qsgvertexcolormaterial.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
import "qt.go/qtqml"

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

type QSGVertexColorMaterial struct {
	*QSGMaterial
}

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
func NewQSGVertexColorMaterial() *QSGVertexColorMaterial {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QSGVertexColorMaterialC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGVertexColorMaterialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGVertexColorMaterial)
	return gothis
}

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int compare(const QSGMaterial *)
func (this *QSGVertexColorMaterial) Compare(other *QSGMaterial /*777 const QSGMaterial **/) int {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QSGVertexColorMaterial7compareEPK11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h:55
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type()
func (this *QSGVertexColorMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QSGVertexColorMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgvertexcolormaterial.h:56
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader()
func (this *QSGVertexColorMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QSGVertexColorMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

func DeleteQSGVertexColorMaterial(this *QSGVertexColorMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QSGVertexColorMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
