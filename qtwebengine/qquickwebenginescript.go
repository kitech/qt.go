package qtwebengine

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h
// #include <qquickwebenginescript.h>
// #include <QtWebEngine>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

// void timerEvent(QTimerEvent *)
func (this *QQuickWebEngineScript) InheritTimerEvent(f func(e *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

/*

 */
type QQuickWebEngineScript struct {
	*qtcore.QObject
}
type QQuickWebEngineScript_ITF interface {
	qtcore.QObject_ITF
	QQuickWebEngineScript_PTR() *QQuickWebEngineScript
}

func (ptr *QQuickWebEngineScript) QQuickWebEngineScript_PTR() *QQuickWebEngineScript { return ptr }

func (this *QQuickWebEngineScript) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickWebEngineScript) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickWebEngineScriptFromPointer(cthis unsafe.Pointer) *QQuickWebEngineScript {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickWebEngineScript{bcthis0}
}
func (*QQuickWebEngineScript) NewFromPointer(cthis unsafe.Pointer) *QQuickWebEngineScript {
	return NewQQuickWebEngineScriptFromPointer(cthis)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickWebEngineScript) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWebEngineScript(QObject *)

/*
Constructs a new QQuickWebEngineScript with the parent parent.
*/
func (*QQuickWebEngineScript) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QQuickWebEngineScript {
	return NewQQuickWebEngineScript(parent)
}
func NewQQuickWebEngineScript(parent qtcore.QObject_ITF /*777 QObject **/) *QQuickWebEngineScript {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScriptC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWebEngineScriptFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWebEngineScript")
	return gothis
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWebEngineScript(QObject *)

/*
Constructs a new QQuickWebEngineScript with the parent parent.
*/
func (*QQuickWebEngineScript) NewForInheritp() *QQuickWebEngineScript {
	return NewQQuickWebEngineScriptp()
}
func NewQQuickWebEngineScriptp() *QQuickWebEngineScript {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScriptC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWebEngineScriptFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWebEngineScript")
	return gothis
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickWebEngineScript()

/*

 */
func DeleteQQuickWebEngineScript(this *QQuickWebEngineScript) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScriptD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the script object as string.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.
*/
func (this *QQuickWebEngineScript) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QQuickWebEngineScript) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl sourceUrl() const

/*

 */
func (this *QQuickWebEngineScript) SourceUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript9sourceUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sourceCode() const

/*

 */
func (this *QQuickWebEngineScript) SourceCode() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript10sourceCodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickWebEngineScript::InjectionPoint injectionPoint() const

/*

 */
func (this *QQuickWebEngineScript) InjectionPoint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript14injectionPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickWebEngineScript::ScriptWorldId worldId() const

/*

 */
func (this *QQuickWebEngineScript) WorldId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript7worldIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool runOnSubframes() const

/*

 */
func (this *QQuickWebEngineScript) RunOnSubframes() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQuickWebEngineScript14runOnSubframesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QString &)

/*

 */
func (this *QQuickWebEngineScript) SetName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript7setNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceUrl(const QUrl &)

/*

 */
func (this *QQuickWebEngineScript) SetSourceUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript12setSourceUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceCode(const QString &)

/*

 */
func (this *QQuickWebEngineScript) SetSourceCode(code string) {
	var tmpArg0 = qtcore.NewQString5(code)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript13setSourceCodeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInjectionPoint(QQuickWebEngineScript::InjectionPoint)

/*

 */
func (this *QQuickWebEngineScript) SetInjectionPoint(injectionPoint int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript17setInjectionPointENS_14InjectionPointE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), injectionPoint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorldId(QQuickWebEngineScript::ScriptWorldId)

/*

 */
func (this *QQuickWebEngineScript) SetWorldId(scriptWorldId int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript10setWorldIdENS_13ScriptWorldIdE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scriptWorldId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRunOnSubframes(bool)

/*

 */
func (this *QQuickWebEngineScript) SetRunOnSubframes(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript17setRunOnSubframesEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void nameChanged(const QString &)

/*

 */
func (this *QQuickWebEngineScript) NameChanged(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript11nameChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sourceUrlChanged(const QUrl &)

/*

 */
func (this *QQuickWebEngineScript) SourceUrlChanged(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript16sourceUrlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sourceCodeChanged(const QString &)

/*

 */
func (this *QQuickWebEngineScript) SourceCodeChanged(code string) {
	var tmpArg0 = qtcore.NewQString5(code)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript17sourceCodeChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void injectionPointChanged(QQuickWebEngineScript::InjectionPoint)

/*

 */
func (this *QQuickWebEngineScript) InjectionPointChanged(injectionPoint int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript21injectionPointChangedENS_14InjectionPointE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), injectionPoint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void worldIdChanged(QQuickWebEngineScript::ScriptWorldId)

/*

 */
func (this *QQuickWebEngineScript) WorldIdChanged(scriptWorldId int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript14worldIdChangedENS_13ScriptWorldIdE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scriptWorldId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void runOnSubframesChanged(bool)

/*

 */
func (this *QQuickWebEngineScript) RunOnSubframesChanged(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript21runOnSubframesChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebenginescript.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*

 */
func (this *QQuickWebEngineScript) TimerEvent(e qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QTimerEvent_PTR() != nil {
		convArg0 = e.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQuickWebEngineScript10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
The point in the loading process at which the script will be executed.


*/
type QQuickWebEngineScript__InjectionPoint = int

//
const QQuickWebEngineScript__Deferred QQuickWebEngineScript__InjectionPoint = 0

// The script will run as soon as the DOM is ready. This is equivalent to the DOMContentLoaded event firing in JavaScript.
const QQuickWebEngineScript__DocumentReady QQuickWebEngineScript__InjectionPoint = 1

// The script will be executed as soon as the document is created. This is not suitable for any DOM operation.
const QQuickWebEngineScript__DocumentCreation QQuickWebEngineScript__InjectionPoint = 2

func (this *QQuickWebEngineScript) InjectionPointItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWebEngineScript_InjectionPointItemName(val int) string {
	var nilthis *QQuickWebEngineScript
	return nilthis.InjectionPointItemName(val)
}

/*
The world ID defining which isolated world the script is executed in. Besides these predefined IDs custom IDs can be used, but must be integers between 0 and 256.


*/
type QQuickWebEngineScript__ScriptWorldId = int

// The world used by the page's web contents. It can be useful in order to expose custom functionality to web contents in certain scenarios.
const QQuickWebEngineScript__MainWorld QQuickWebEngineScript__ScriptWorldId = 0

// The default isolated world used for application level functionality implemented in JavaScript.
const QQuickWebEngineScript__ApplicationWorld QQuickWebEngineScript__ScriptWorldId = 1

// The first isolated world to be used by scripts set by users if the application is not making use of more worlds. As a rule of thumb, if that functionality is exposed to the application users, each individual script should probably get its own isolated world.
const QQuickWebEngineScript__UserWorld QQuickWebEngineScript__ScriptWorldId = 2

func (this *QQuickWebEngineScript) ScriptWorldIdItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWebEngineScript_ScriptWorldIdItemName(val int) string {
	var nilthis *QQuickWebEngineScript
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
		qtpositioning.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
