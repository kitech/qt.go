package qtquick

// /usr/include/qt/QtQuick/qsgmaterial.h
// #include <qsgmaterial.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
// void setShaderSourceFile(class QOpenGLShader::ShaderType, const class QString &)
func (this *QSGMaterialShader) InheritSetShaderSourceFile(f func(type_ int, sourceFile *qtcore.QString) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setShaderSourceFile", f)
}

// void setShaderSourceFiles(class QOpenGLShader::ShaderType, const class QStringList &)
func (this *QSGMaterialShader) InheritSetShaderSourceFiles(f func(type_ int, sourceFiles *qtcore.QStringList) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setShaderSourceFiles", f)
}

// void compile()
func (this *QSGMaterialShader) InheritCompile(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "compile", f)
}

// void initialize()
func (this *QSGMaterialShader) InheritInitialize(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "initialize", f)
}

// const char * vertexShader()
func (this *QSGMaterialShader) InheritVertexShader(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "vertexShader", f)
}

// const char * fragmentShader()
func (this *QSGMaterialShader) InheritFragmentShader(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "fragmentShader", f)
}

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSGMaterialShaderFromPointer(cthis unsafe.Pointer) *QSGMaterialShader {
	return &QSGMaterialShader{&qtrt.CObject{cthis}}
}
func (*QSGMaterialShader) NewFromPointer(cthis unsafe.Pointer) *QSGMaterialShader {
	return NewQSGMaterialShaderFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGMaterialShader()
func NewQSGMaterialShader() *QSGMaterialShader {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShaderC1Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGMaterialShader)
	return gothis
}

// /usr/include/qt/QtQuick/qsgmaterial.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGMaterialShader()
func DeleteQSGMaterialShader(this *QSGMaterialShader) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShaderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void activate()
func (this *QSGMaterialShader) Activate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader8activateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void deactivate()
func (this *QSGMaterialShader) Deactivate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader10deactivateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:103
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] const char *const * attributeNames()
func (this *QSGMaterialShader) AttributeNames() []string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGMaterialShader14attributeNamesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.CCharPPToStringSlice(unsafe.Pointer(uintptr(rv)))
}

// /usr/include/qt/QtQuick/qsgmaterial.h:114
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setShaderSourceFile(QOpenGLShader::ShaderType, const QString &)
func (this *QSGMaterialShader) SetShaderSourceFile(type_ int, sourceFile *qtcore.QString) {
	var convArg1 = sourceFile.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader19setShaderSourceFileE6QFlagsIN13QOpenGLShader13ShaderTypeBitEERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:115
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setShaderSourceFiles(QOpenGLShader::ShaderType, const QStringList &)
func (this *QSGMaterialShader) SetShaderSourceFiles(type_ int, sourceFiles *qtcore.QStringList) {
	var convArg1 = sourceFiles.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader20setShaderSourceFilesE6QFlagsIN13QOpenGLShader13ShaderTypeBitEERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void compile()
func (this *QSGMaterialShader) Compile() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader7compileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:119
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void initialize()
func (this *QSGMaterialShader) Initialize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader10initializeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] const char * vertexShader()
func (this *QSGMaterialShader) VertexShader() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGMaterialShader12vertexShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] const char * fragmentShader()
func (this *QSGMaterialShader) FragmentShader() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGMaterialShader14fragmentShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

//  body block end
