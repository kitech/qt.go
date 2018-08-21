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

/*

 */
func DeleteQSGRendererInterface(this *QSGRendererInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGRendererInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::GraphicsApi graphicsApi() const

/*
Returns the graphics API that is in use by the Qt Quick scenegraph.

Note: This function can be called on any thread.
*/
func (this *QSGRendererInterface) GraphicsApi() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface11graphicsApiEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] void * getResource(QQuickWindow *, QSGRendererInterface::Resource) const

/*
Queries a graphics resource in window. Returns null when the resource in question is not supported or not available.

When successful, the returned pointer is either a direct pointer to an interface (and can be cast, for example, to ID3D12Device *) or a pointer to an opaque handle that needs to be dereferenced first (for example, VkDevice dev = *static_cast<VkDevice *>(result)). The latter is necessary since such handles may have sizes different from a pointer.

Note: The ownership of the returned pointer is never transferred to the caller.

Note: This function must only be called on the render thread.
*/
func (this *QSGRendererInterface) GetResource(window QQuickWindow_ITF /*777 QQuickWindow **/, resource int) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QQuickWindow_PTR() != nil {
		convArg0 = window.QQuickWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface11getResourceEP12QQuickWindowNS_8ResourceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, resource)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:91
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] void * getResource(QQuickWindow *, const char *) const

/*
Queries a graphics resource in window. Returns null when the resource in question is not supported or not available.

When successful, the returned pointer is either a direct pointer to an interface (and can be cast, for example, to ID3D12Device *) or a pointer to an opaque handle that needs to be dereferenced first (for example, VkDevice dev = *static_cast<VkDevice *>(result)). The latter is necessary since such handles may have sizes different from a pointer.

Note: The ownership of the returned pointer is never transferred to the caller.

Note: This function must only be called on the render thread.
*/
func (this *QSGRendererInterface) GetResource_1(window QQuickWindow_ITF /*777 QQuickWindow **/, resource string) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QQuickWindow_PTR() != nil {
		convArg0 = window.QQuickWindow_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(resource)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface11getResourceEP12QQuickWindowPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:93
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::ShaderType shaderType() const

/*
Returns the shading language supported by the Qt Quick backend the application is using.

Note: This function can be called on any thread.

See also QtQuick::GraphicsInfo.
*/
func (this *QSGRendererInterface) ShaderType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface10shaderTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:94
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::ShaderCompilationTypes shaderCompilationType() const

/*
Returns a bitmask of the shader compilation approaches supported by the Qt Quick backend the application is using.

Note: This function can be called on any thread.

See also QtQuick::GraphicsInfo.
*/
func (this *QSGRendererInterface) ShaderCompilationType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface21shaderCompilationTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:95
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSGRendererInterface::ShaderSourceTypes shaderSourceType() const

/*
Returns a bitmask of the supported ways of providing shader sources in ShaderEffect items.

Note: This function can be called on any thread.

See also QtQuick::GraphicsInfo.
*/
func (this *QSGRendererInterface) ShaderSourceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGRendererInterface16shaderSourceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*

 */
type QSGRendererInterface__GraphicsApi = int

// An unknown graphics API is in use
const QSGRendererInterface__Unknown QSGRendererInterface__GraphicsApi = 0

//
const QSGRendererInterface__Software QSGRendererInterface__GraphicsApi = 1

//
const QSGRendererInterface__OpenGL QSGRendererInterface__GraphicsApi = 2

//
const QSGRendererInterface__Direct3D12 QSGRendererInterface__GraphicsApi = 3

// OpenVG via EGL
const QSGRendererInterface__OpenVG QSGRendererInterface__GraphicsApi = 4

func (this *QSGRendererInterface) GraphicsApiItemName(val int) string {
	switch val {
	case QSGRendererInterface__Unknown: // 0
		return "Unknown"
	case QSGRendererInterface__Software: // 1
		return "Software"
	case QSGRendererInterface__OpenGL: // 2
		return "OpenGL"
	case QSGRendererInterface__Direct3D12: // 3
		return "Direct3D12"
	case QSGRendererInterface__OpenVG: // 4
		return "OpenVG"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGRendererInterface_GraphicsApiItemName(val int) string {
	var nilthis *QSGRendererInterface
	return nilthis.GraphicsApiItemName(val)
}

/*

 */
type QSGRendererInterface__Resource = int

// The graphics device, when applicable.
const QSGRendererInterface__DeviceResource QSGRendererInterface__Resource = 0

// The graphics command queue used by the scenegraph, when applicable.
const QSGRendererInterface__CommandQueueResource QSGRendererInterface__Resource = 1

// The command list or buffer used by the scenegraph, when applicable.
const QSGRendererInterface__CommandListResource QSGRendererInterface__Resource = 2

// The active QPainter used by the scenegraph, when running with the software backend.
const QSGRendererInterface__PainterResource QSGRendererInterface__Resource = 3

func (this *QSGRendererInterface) ResourceItemName(val int) string {
	switch val {
	case QSGRendererInterface__DeviceResource: // 0
		return "DeviceResource"
	case QSGRendererInterface__CommandQueueResource: // 1
		return "CommandQueueResource"
	case QSGRendererInterface__CommandListResource: // 2
		return "CommandListResource"
	case QSGRendererInterface__PainterResource: // 3
		return "PainterResource"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGRendererInterface_ResourceItemName(val int) string {
	var nilthis *QSGRendererInterface
	return nilthis.ResourceItemName(val)
}

/*

 */
type QSGRendererInterface__ShaderType = int

// Not yet known due to no window and scenegraph associated
const QSGRendererInterface__UnknownShadingLanguage QSGRendererInterface__ShaderType = 0

// GLSL or GLSL ES
const QSGRendererInterface__GLSL QSGRendererInterface__ShaderType = 1

// HLSL
const QSGRendererInterface__HLSL QSGRendererInterface__ShaderType = 2

func (this *QSGRendererInterface) ShaderTypeItemName(val int) string {
	switch val {
	case QSGRendererInterface__UnknownShadingLanguage: // 0
		return "UnknownShadingLanguage"
	case QSGRendererInterface__GLSL: // 1
		return "GLSL"
	case QSGRendererInterface__HLSL: // 2
		return "HLSL"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGRendererInterface_ShaderTypeItemName(val int) string {
	var nilthis *QSGRendererInterface
	return nilthis.ShaderTypeItemName(val)
}

/*


 */
type QSGRendererInterface__ShaderCompilationType = int

//
const QSGRendererInterface__RuntimeCompilation QSGRendererInterface__ShaderCompilationType = 1

//
const QSGRendererInterface__OfflineCompilation QSGRendererInterface__ShaderCompilationType = 2

func (this *QSGRendererInterface) ShaderCompilationTypeItemName(val int) string {
	switch val {
	case QSGRendererInterface__RuntimeCompilation: // 1
		return "RuntimeCompilation"
	case QSGRendererInterface__OfflineCompilation: // 2
		return "OfflineCompilation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGRendererInterface_ShaderCompilationTypeItemName(val int) string {
	var nilthis *QSGRendererInterface
	return nilthis.ShaderCompilationTypeItemName(val)
}

/*


 */
type QSGRendererInterface__ShaderSourceType = int

//
const QSGRendererInterface__ShaderSourceString QSGRendererInterface__ShaderSourceType = 1

//
const QSGRendererInterface__ShaderSourceFile QSGRendererInterface__ShaderSourceType = 2

//
const QSGRendererInterface__ShaderByteCode QSGRendererInterface__ShaderSourceType = 4

func (this *QSGRendererInterface) ShaderSourceTypeItemName(val int) string {
	switch val {
	case QSGRendererInterface__ShaderSourceString: // 1
		return "ShaderSourceString"
	case QSGRendererInterface__ShaderSourceFile: // 2
		return "ShaderSourceFile"
	case QSGRendererInterface__ShaderByteCode: // 4
		return "ShaderByteCode"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGRendererInterface_ShaderSourceTypeItemName(val int) string {
	var nilthis *QSGRendererInterface
	return nilthis.ShaderSourceTypeItemName(val)
}

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
