package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h
// #include <qwebenginescript.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

/*

 */
type QWebEngineScript struct {
	*qtrt.CObject
}
type QWebEngineScript_ITF interface {
	QWebEngineScript_PTR() *QWebEngineScript
}

func (ptr *QWebEngineScript) QWebEngineScript_PTR() *QWebEngineScript { return ptr }

func (this *QWebEngineScript) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineScript) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineScriptFromPointer(cthis unsafe.Pointer) *QWebEngineScript {
	return &QWebEngineScript{&qtrt.CObject{cthis}}
}
func (*QWebEngineScript) NewFromPointer(cthis unsafe.Pointer) *QWebEngineScript {
	return NewQWebEngineScriptFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineScript()

/*

 */
func NewQWebEngineScript() *QWebEngineScript {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScriptC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineScriptFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineScript)
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineScript()

/*

 */
func DeleteQWebEngineScript(this *QWebEngineScript) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScriptD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineScript & operator=(const QWebEngineScript &)

/*

 */
func (this *QWebEngineScript) Operator_equal(other QWebEngineScript_ITF) *QWebEngineScript {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineScript_PTR() != nil {
		convArg0 = other.QWebEngineScript_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScriptaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineScriptFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineScript)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*

 */
func (this *QWebEngineScript) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScript6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QWebEngineScript) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScript4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QString &)

/*

 */
func (this *QWebEngineScript) SetName(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScript7setNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sourceCode() const

/*

 */
func (this *QWebEngineScript) SourceCode() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScript10sourceCodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceCode(const QString &)

/*

 */
func (this *QWebEngineScript) SetSourceCode(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScript13setSourceCodeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineScript::InjectionPoint injectionPoint() const

/*

 */
func (this *QWebEngineScript) InjectionPoint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScript14injectionPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInjectionPoint(QWebEngineScript::InjectionPoint)

/*

 */
func (this *QWebEngineScript) SetInjectionPoint(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScript17setInjectionPointENS_14InjectionPointE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 worldId() const

/*

 */
func (this *QWebEngineScript) WorldId() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScript7worldIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldId(quint32)

/*

 */
func (this *QWebEngineScript) SetWorldId(arg0 uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScript10setWorldIdEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool runsOnSubFrames() const

/*

 */
func (this *QWebEngineScript) RunsOnSubFrames() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScript15runsOnSubFramesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRunsOnSubFrames(bool)

/*

 */
func (this *QWebEngineScript) SetRunsOnSubFrames(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScript18setRunsOnSubFramesEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QWebEngineScript &) const

/*

 */
func (this *QWebEngineScript) Operator_equal_equal(other QWebEngineScript_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineScript_PTR() != nil {
		convArg0 = other.QWebEngineScript_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScripteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QWebEngineScript &) const

/*

 */
func (this *QWebEngineScript) Operator_not_equal(other QWebEngineScript_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineScript_PTR() != nil {
		convArg0 = other.QWebEngineScript_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QWebEngineScriptneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineScript &)

/*

 */
func (this *QWebEngineScript) Swap(other QWebEngineScript_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineScript_PTR() != nil {
		convArg0 = other.QWebEngineScript_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QWebEngineScript4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QWebEngineScript__InjectionPoint = int

//
const QWebEngineScript__Deferred QWebEngineScript__InjectionPoint = 0

//
const QWebEngineScript__DocumentReady QWebEngineScript__InjectionPoint = 1

//
const QWebEngineScript__DocumentCreation QWebEngineScript__InjectionPoint = 2

func (this *QWebEngineScript) InjectionPointItemName(val int) string {
	switch val {
	case QWebEngineScript__Deferred: // 0
		return "Deferred"
	case QWebEngineScript__DocumentReady: // 1
		return "DocumentReady"
	case QWebEngineScript__DocumentCreation: // 2
		return "DocumentCreation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineScript_InjectionPointItemName(val int) string {
	var nilthis *QWebEngineScript
	return nilthis.InjectionPointItemName(val)
}

/*


 */
type QWebEngineScript__ScriptWorldId = int

//
const QWebEngineScript__MainWorld QWebEngineScript__ScriptWorldId = 0

//
const QWebEngineScript__ApplicationWorld QWebEngineScript__ScriptWorldId = 1

//
const QWebEngineScript__UserWorld QWebEngineScript__ScriptWorldId = 2

func (this *QWebEngineScript) ScriptWorldIdItemName(val int) string {
	switch val {
	case QWebEngineScript__MainWorld: // 0
		return "MainWorld"
	case QWebEngineScript__ApplicationWorld: // 1
		return "ApplicationWorld"
	case QWebEngineScript__UserWorld: // 2
		return "UserWorld"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineScript_ScriptWorldIdItemName(val int) string {
	var nilthis *QWebEngineScript
	return nilthis.ScriptWorldIdItemName(val)
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
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtquickwidgets.KeepMe()
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
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
