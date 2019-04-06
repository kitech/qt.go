package qtquick

// /usr/include/qt/QtQuick/qsgtexture.h
// #include <qsgtexture.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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

/*

 */
type QSGDynamicTexture struct {
	*QSGTexture
}
type QSGDynamicTexture_ITF interface {
	QSGTexture_ITF
	QSGDynamicTexture_PTR() *QSGDynamicTexture
}

func (ptr *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSGDynamicTexture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGDynamicTexture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexture.h:133
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool updateTexture()

/*

 */
func (this *QSGDynamicTexture) UpdateTexture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGDynamicTexture13updateTextureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQSGDynamicTexture(this *QSGDynamicTexture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGDynamicTextureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11605() {
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
