//  header block begin

// +build !minimal

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
// extern C begin: 11
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

// /usr/include/qt/QtQuick/qsgmaterial.h:114
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setShaderSourceFile(QOpenGLShader::ShaderType, const QString &)

/*

 */
func (this *QSGMaterialShader) SetShaderSourceFile(type_ int, sourceFile string) {
	var tmpArg1 = qtcore.NewQString5(sourceFile)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader19setShaderSourceFileE6QFlagsIN13QOpenGLShader13ShaderTypeBitEERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:115
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setShaderSourceFiles(QOpenGLShader::ShaderType, const QStringList &)

/*

 */
func (this *QSGMaterialShader) SetShaderSourceFiles(type_ int, sourceFiles qtcore.QStringList_ITF) {
	var convArg1 unsafe.Pointer
	if sourceFiles != nil && sourceFiles.QStringList_PTR() != nil {
		convArg1 = sourceFiles.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader20setShaderSourceFilesE6QFlagsIN13QOpenGLShader13ShaderTypeBitEERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void compile()

/*

 */
func (this *QSGMaterialShader) Compile() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGMaterialShader7compileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] const char * vertexShader() const

/*

 */
func (this *QSGMaterialShader) VertexShader() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGMaterialShader12vertexShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] const char * fragmentShader() const

/*

 */
func (this *QSGMaterialShader) FragmentShader() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGMaterialShader14fragmentShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

//  body block end

//  keep block begin

func init_unused_11596() {
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
