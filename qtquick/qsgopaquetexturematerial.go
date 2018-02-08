package qtquick

// /usr/include/qt/QtQuick/qsgtexturematerial.h
// #include <qsgtexturematerial.h>
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

type QSGOpaqueTextureMaterial struct {
	*QSGMaterial
}

func (this *QSGOpaqueTextureMaterial) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGMaterial.GetCthis()
	}
}
func (this *QSGOpaqueTextureMaterial) SetCthis(cthis unsafe.Pointer) {
	this.QSGMaterial = NewQSGMaterialFromPointer(cthis)
}
func NewQSGOpaqueTextureMaterialFromPointer(cthis unsafe.Pointer) *QSGOpaqueTextureMaterial {
	bcthis0 := NewQSGMaterialFromPointer(cthis)
	return &QSGOpaqueTextureMaterial{bcthis0}
}
func (*QSGOpaqueTextureMaterial) NewFromPointer(cthis unsafe.Pointer) *QSGOpaqueTextureMaterial {
	return NewQSGOpaqueTextureMaterialFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:51
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGOpaqueTextureMaterial()
func NewQSGOpaqueTextureMaterial() *QSGOpaqueTextureMaterial {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterialC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGOpaqueTextureMaterialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGOpaqueTextureMaterial)
	return gothis
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type()
func (this *QSGOpaqueTextureMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader()
func (this *QSGOpaqueTextureMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int compare(const QSGMaterial *)
func (this *QSGOpaqueTextureMaterial) Compare(other *QSGMaterial /*777 const QSGMaterial **/) int {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial7compareEPK11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTexture(QSGTexture *)
func (this *QSGOpaqueTextureMaterial) SetTexture(texture *QSGTexture /*777 QSGTexture **/) {
	var convArg0 = texture.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterial10setTextureEP10QSGTexture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGTexture * texture()
func (this *QSGOpaqueTextureMaterial) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial7textureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setMipmapFiltering(QSGTexture::Filtering)
func (this *QSGOpaqueTextureMaterial) SetMipmapFiltering(filteringType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterial18setMipmapFilteringEN10QSGTexture9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filteringType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGTexture::Filtering mipmapFiltering()
func (this *QSGOpaqueTextureMaterial) MipmapFiltering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial15mipmapFilteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFiltering(QSGTexture::Filtering)
func (this *QSGOpaqueTextureMaterial) SetFiltering(filteringType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterial12setFilteringEN10QSGTexture9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filteringType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGTexture::Filtering filtering()
func (this *QSGOpaqueTextureMaterial) Filtering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial9filteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHorizontalWrapMode(QSGTexture::WrapMode)
func (this *QSGOpaqueTextureMaterial) SetHorizontalWrapMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterial21setHorizontalWrapModeEN10QSGTexture8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGTexture::WrapMode horizontalWrapMode()
func (this *QSGOpaqueTextureMaterial) HorizontalWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial18horizontalWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setVerticalWrapMode(QSGTexture::WrapMode)
func (this *QSGOpaqueTextureMaterial) SetVerticalWrapMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterial19setVerticalWrapModeEN10QSGTexture8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGTexture::WrapMode verticalWrapMode()
func (this *QSGOpaqueTextureMaterial) VerticalWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial16verticalWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAnisotropyLevel(QSGTexture::AnisotropyLevel)
func (this *QSGOpaqueTextureMaterial) SetAnisotropyLevel(level int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterial18setAnisotropyLevelEN10QSGTexture15AnisotropyLevelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexturematerial.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGTexture::AnisotropyLevel anisotropyLevel()
func (this *QSGOpaqueTextureMaterial) AnisotropyLevel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSGOpaqueTextureMaterial15anisotropyLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQSGOpaqueTextureMaterial(this *QSGOpaqueTextureMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSGOpaqueTextureMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
