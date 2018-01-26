package qtquick

// /usr/include/qt/QtQuick/qsgrendererinterface.h
// #include <qsgrendererinterface.h>
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
import "mkuse/cffiqt"
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
type QSGRendererInterface struct {
	*qtrt.CObject
}

func (this *QSGRendererInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGRendererInterface) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSGRendererInterfaceFromPointer(cthis unsafe.Pointer) *QSGRendererInterface {
	return &QSGRendererInterface{&qtrt.CObject{cthis}}
}
func (*QSGRendererInterface) NewFromPointer(cthis unsafe.Pointer) *QSGRendererInterface {
	return NewQSGRendererInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:86
// index:0
// Public virtual
// void ~QSGRendererInterface()
func DeleteQSGRendererInterface(*QSGRendererInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGRendererInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:88
// index:0
// Public pure virtual
// QSGRendererInterface::GraphicsApi graphicsApi()
func (this *QSGRendererInterface) GraphicsApi() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGRendererInterface11graphicsApiEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:90
// index:0
// Public virtual
// void * getResource(class QQuickWindow *, enum QSGRendererInterface::Resource)
func (this *QSGRendererInterface) GetResource(window *QQuickWindow /*777 QQuickWindow **/, resource int) unsafe.Pointer /*666*/ {
	var convArg0 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGRendererInterface11getResourceEP12QQuickWindowNS_8ResourceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, resource)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:91
// index:1
// Public virtual
// void * getResource(class QQuickWindow *, const char *)
func (this *QSGRendererInterface) GetResource_1(window *QQuickWindow /*777 QQuickWindow **/, resource string) unsafe.Pointer /*666*/ {
	var convArg0 = window.GetCthis()
	var convArg1 = qtrt.CString(resource)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGRendererInterface11getResourceEP12QQuickWindowPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:93
// index:0
// Public pure virtual
// QSGRendererInterface::ShaderType shaderType()
func (this *QSGRendererInterface) ShaderType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGRendererInterface10shaderTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:94
// index:0
// Public pure virtual
// QSGRendererInterface::ShaderCompilationTypes shaderCompilationType()
func (this *QSGRendererInterface) ShaderCompilationType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGRendererInterface21shaderCompilationTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:95
// index:0
// Public pure virtual
// QSGRendererInterface::ShaderSourceTypes shaderSourceType()
func (this *QSGRendererInterface) ShaderSourceType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGRendererInterface16shaderSourceTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QSGRendererInterface__GraphicsApi = int

const QSGRendererInterface__Unknown QSGRendererInterface__GraphicsApi = 0
const QSGRendererInterface__Software QSGRendererInterface__GraphicsApi = 1
const QSGRendererInterface__OpenGL QSGRendererInterface__GraphicsApi = 2
const QSGRendererInterface__Direct3D12 QSGRendererInterface__GraphicsApi = 3
const QSGRendererInterface__OpenVG QSGRendererInterface__GraphicsApi = 4

type QSGRendererInterface__Resource = int

const QSGRendererInterface__DeviceResource QSGRendererInterface__Resource = 0
const QSGRendererInterface__CommandQueueResource QSGRendererInterface__Resource = 1
const QSGRendererInterface__CommandListResource QSGRendererInterface__Resource = 2
const QSGRendererInterface__PainterResource QSGRendererInterface__Resource = 3

type QSGRendererInterface__ShaderType = int

const QSGRendererInterface__UnknownShadingLanguage QSGRendererInterface__ShaderType = 0
const QSGRendererInterface__GLSL QSGRendererInterface__ShaderType = 1
const QSGRendererInterface__HLSL QSGRendererInterface__ShaderType = 2

type QSGRendererInterface__ShaderCompilationType = int

const QSGRendererInterface__RuntimeCompilation QSGRendererInterface__ShaderCompilationType = 1
const QSGRendererInterface__OfflineCompilation QSGRendererInterface__ShaderCompilationType = 2

type QSGRendererInterface__ShaderSourceType = int

const QSGRendererInterface__ShaderSourceString QSGRendererInterface__ShaderSourceType = 1
const QSGRendererInterface__ShaderSourceFile QSGRendererInterface__ShaderSourceType = 2
const QSGRendererInterface__ShaderByteCode QSGRendererInterface__ShaderSourceType = 4

//  body block end
