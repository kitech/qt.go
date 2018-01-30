package qtquick

// /usr/include/qt/QtQuick/qsgtexture.h
// #include <qsgtexture.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QSGDynamicTexture struct {
	*QSGTexture
}

func (this *QSGDynamicTexture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGTexture.GetCthis()
	}
}
func (this *QSGDynamicTexture) SetCthis(cthis unsafe.Pointer) {
	this.QSGTexture = NewQSGTextureFromPointer(cthis)
}
func NewQSGDynamicTextureFromPointer(cthis unsafe.Pointer) *QSGDynamicTexture {
	bcthis0 := NewQSGTextureFromPointer(cthis)
	return &QSGDynamicTexture{bcthis0}
}
func (*QSGDynamicTexture) NewFromPointer(cthis unsafe.Pointer) *QSGDynamicTexture {
	return NewQSGDynamicTextureFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgtexture.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSGDynamicTexture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSGDynamicTexture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgtexture.h:133
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool updateTexture()
func (this *QSGDynamicTexture) UpdateTexture() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGDynamicTexture13updateTextureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
