package qtquick

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h
// #include <qsgflatcolormaterial.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
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
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSGFlatColorMaterial struct {
	*QSGMaterial
}

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
	gopp.ErrPrint(err, rv)
	gothis := NewQSGFlatColorMaterialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGFlatColorMaterial)
	return gothis
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type()
func (this *QSGFlatColorMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader()
func (this *QSGFlatColorMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QSGFlatColorMaterial) SetColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGFlatColorMaterial8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:56
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QColor & color()
func (this *QSGFlatColorMaterial) Color() *qtgui.QColor {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtQuick/qsgflatcolormaterial.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int compare(const QSGMaterial *)
func (this *QSGFlatColorMaterial) Compare(other *QSGMaterial /*777 const QSGMaterial **/) int {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGFlatColorMaterial7compareEPK11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQSGFlatColorMaterial(this *QSGFlatColorMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGFlatColorMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
