package qtquick

// /usr/include/qt/QtQuick/qsgmaterial.h
// #include <qsgmaterial.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
type QSGMaterialShader struct {
	*qtrt.CObject
}

func (this *QSGMaterialShader) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGMaterialShader) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSGMaterialShaderFromPointer(cthis unsafe.Pointer) *QSGMaterialShader {
	return &QSGMaterialShader{&qtrt.CObject{cthis}}
}
func (*QSGMaterialShader) NewFromPointer(cthis unsafe.Pointer) *QSGMaterialShader {
	return NewQSGMaterialShaderFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:96
// index:0
// Public
// void QSGMaterialShader()
func NewQSGMaterialShader() *QSGMaterialShader {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGMaterialShaderC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGMaterialShaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgmaterial.h:97
// index:0
// Public virtual
// void ~QSGMaterialShader()
func DeleteQSGMaterialShader(*QSGMaterialShader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGMaterialShaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:99
// index:0
// Public virtual
// void activate()
func (this *QSGMaterialShader) Activate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGMaterialShader8activateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:100
// index:0
// Public virtual
// void deactivate()
func (this *QSGMaterialShader) Deactivate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGMaterialShader10deactivateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:103
// index:0
// Public pure virtual
// const char *const * attributeNames()
func (this *QSGMaterialShader) AttributeNames() []string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSGMaterialShader14attributeNamesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.CCharPPToStringSlice(unsafe.Pointer(uintptr(rv)))
}

// /usr/include/qt/QtQuick/qsgmaterial.h:117
// index:0
// Protected virtual
// void compile()
func (this *QSGMaterialShader) Compile() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGMaterialShader7compileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:119
// index:0
// Protected inline virtual
// void initialize()
func (this *QSGMaterialShader) Initialize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGMaterialShader10initializeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:121
// index:0
// Protected virtual
// const char * vertexShader()
func (this *QSGMaterialShader) VertexShader() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSGMaterialShader12vertexShaderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:122
// index:0
// Protected virtual
// const char * fragmentShader()
func (this *QSGMaterialShader) FragmentShader() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSGMaterialShader14fragmentShaderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

//  body block end
