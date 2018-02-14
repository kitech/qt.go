package qtquick

// /usr/include/qt/QtQuick/qsgrendererinterface.h
// #include <qsgrendererinterface.h>
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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QSGRendererInterface struct {
	*qtrt.CObject
}
type QSGRendererInterface_ITF interface {
	QSGRendererInterface_PTR() *QSGRendererInterface
}

func (ptr *QSGRendererInterface) QSGRendererInterface_PTR() *QSGRendererInterface { return ptr }

func (this *QSGRendererInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGRendererInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSGRendererInterfaceFromPointer(cthis unsafe.Pointer) *QSGRendererInterface {
	return &QSGRendererInterface{&qtrt.CObject{cthis}}
}
func (*QSGRendererInterface) NewFromPointer(cthis unsafe.Pointer) *QSGRendererInterface {
	return NewQSGRendererInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGRendererInterface()
func DeleteQSGRendererInterface(this *QSGRendererInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGRendererInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::GraphicsApi graphicsApi()
func (this *QSGRendererInterface) GraphicsApi() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface11graphicsApiEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] void * getResource(QQuickWindow *, enum QSGRendererInterface::Resource)
func (this *QSGRendererInterface) GetResource(window QQuickWindow_ITF /*777 QQuickWindow **/, resource int) unsafe.Pointer /*666*/ {
	var convArg0 = window.QQuickWindow_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface11getResourceEP12QQuickWindowNS_8ResourceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, resource)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:91
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] void * getResource(QQuickWindow *, const char *)
func (this *QSGRendererInterface) GetResource_1(window QQuickWindow_ITF /*777 QQuickWindow **/, resource string) unsafe.Pointer /*666*/ {
	var convArg0 = window.QQuickWindow_PTR().GetCthis()
	var convArg1 = qtrt.CString(resource)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface11getResourceEP12QQuickWindowPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:93
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::ShaderType shaderType()
func (this *QSGRendererInterface) ShaderType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface10shaderTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:94
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::ShaderCompilationTypes shaderCompilationType()
func (this *QSGRendererInterface) ShaderCompilationType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface21shaderCompilationTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:95
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::ShaderSourceTypes shaderSourceType()
func (this *QSGRendererInterface) ShaderSourceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface16shaderSourceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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

//  keep block begin

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
