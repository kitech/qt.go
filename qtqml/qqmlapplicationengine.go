package qtqml

// /usr/include/qt/QtQml/qqmlapplicationengine.h
// #include <qqmlapplicationengine.h>
// #include <QtQml>

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
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QQmlApplicationEngine struct {
	*QQmlEngine
}
type QQmlApplicationEngine_ITF interface {
	QQmlEngine_ITF
	QQmlApplicationEngine_PTR() *QQmlApplicationEngine
}

func (ptr *QQmlApplicationEngine) QQmlApplicationEngine_PTR() *QQmlApplicationEngine { return ptr }

func (this *QQmlApplicationEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QQmlEngine.GetCthis()
	}
}
func (this *QQmlApplicationEngine) SetCthis(cthis unsafe.Pointer) {
	this.QQmlEngine = NewQQmlEngineFromPointer(cthis)
}
func NewQQmlApplicationEngineFromPointer(cthis unsafe.Pointer) *QQmlApplicationEngine {
	bcthis0 := NewQQmlEngineFromPointer(cthis)
	return &QQmlApplicationEngine{bcthis0}
}
func (*QQmlApplicationEngine) NewFromPointer(cthis unsafe.Pointer) *QQmlApplicationEngine {
	return NewQQmlApplicationEngineFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQmlApplicationEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQmlApplicationEngine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(QObject *)

/*
Create a new QQmlApplicationEngine with the given parent. You will have to call load() later in order to load a QML file.
*/
func (*QQmlApplicationEngine) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlApplicationEngine {
	return NewQQmlApplicationEngine(parent)
}
func NewQQmlApplicationEngine(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlApplicationEngine {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlApplicationEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(QObject *)

/*
Create a new QQmlApplicationEngine with the given parent. You will have to call load() later in order to load a QML file.
*/
func (*QQmlApplicationEngine) NewForInheritp() *QQmlApplicationEngine {
	return NewQQmlApplicationEnginep()
}
func NewQQmlApplicationEnginep() *QQmlApplicationEngine {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlApplicationEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(const QUrl &, QObject *)

/*
Create a new QQmlApplicationEngine with the given parent. You will have to call load() later in order to load a QML file.
*/
func (*QQmlApplicationEngine) NewForInherit1(url qtcore.QUrl_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlApplicationEngine {
	return NewQQmlApplicationEngine1(url, parent)
}
func NewQQmlApplicationEngine1(url qtcore.QUrl_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlApplicationEngine {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2ERK4QUrlP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlApplicationEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(const QUrl &, QObject *)

/*
Create a new QQmlApplicationEngine with the given parent. You will have to call load() later in order to load a QML file.
*/
func (*QQmlApplicationEngine) NewForInherit1p(url qtcore.QUrl_ITF) *QQmlApplicationEngine {
	return NewQQmlApplicationEngine1p(url)
}
func NewQQmlApplicationEngine1p(url qtcore.QUrl_ITF) *QQmlApplicationEngine {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2ERK4QUrlP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlApplicationEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:58
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(const QString &, QObject *)

/*
Create a new QQmlApplicationEngine with the given parent. You will have to call load() later in order to load a QML file.
*/
func (*QQmlApplicationEngine) NewForInherit2(filePath string, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlApplicationEngine {
	return NewQQmlApplicationEngine2(filePath, parent)
}
func NewQQmlApplicationEngine2(filePath string, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlApplicationEngine {
	var tmpArg0 = qtcore.NewQString5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlApplicationEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:58
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlApplicationEngine(const QString &, QObject *)

/*
Create a new QQmlApplicationEngine with the given parent. You will have to call load() later in order to load a QML file.
*/
func (*QQmlApplicationEngine) NewForInherit2p(filePath string) *QQmlApplicationEngine {
	return NewQQmlApplicationEngine2p(filePath)
}
func NewQQmlApplicationEngine2p(filePath string) *QQmlApplicationEngine {
	var tmpArg0 = qtcore.NewQString5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlApplicationEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlApplicationEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlApplicationEngine()

/*

 */
func DeleteQQmlApplicationEngine(this *QQmlApplicationEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QObject *> rootObjects()

/*
Returns a list of all the root objects instantiated by the QQmlApplicationEngine. This will only contain objects loaded via load() or a convenience constructor.

Note: In Qt versions prior to 5.9, this function is marked as non-const.
*/
func (this *QQmlApplicationEngine) RootObjects() *qtcore.QObjectList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine11rootObjectsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQObjectListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:64
// index:1
// Public Visibility=Default Availability=Available
// [8] QList<QObject *> rootObjects() const

/*
Returns a list of all the root objects instantiated by the QQmlApplicationEngine. This will only contain objects loaded via load() or a convenience constructor.

Note: In Qt versions prior to 5.9, this function is marked as non-const.
*/
func (this *QQmlApplicationEngine) RootObjects1() *qtcore.QObjectList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QQmlApplicationEngine11rootObjectsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQObjectListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(const QUrl &)

/*
Loads the root QML file located at url. The object tree defined by the file is created immediately for local file urls. Remote urls are loaded asynchronously, listen to the objectCreated signal to determine when the object tree is ready.

If an error occurs, error messages are printed with qWarning.
*/
func (this *QQmlApplicationEngine) Load(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine4loadERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(const QString &)

/*
Loads the root QML file located at url. The object tree defined by the file is created immediately for local file urls. Remote urls are loaded asynchronously, listen to the objectCreated signal to determine when the object tree is ready.

If an error occurs, error messages are printed with qWarning.
*/
func (this *QQmlApplicationEngine) Load1(filePath string) {
	var tmpArg0 = qtcore.NewQString5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine4loadERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadData(const QByteArray &, const QUrl &)

/*
Loads the QML given in data. The object tree defined by data is instantiated immediately.

If a url is specified it is used as the base url of the component. This affects relative paths within the data and error messages.

If an error occurs, error messages are printed with qWarning.
*/
func (this *QQmlApplicationEngine) LoadData(data qtcore.QByteArray_ITF, url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine8loadDataERK10QByteArrayRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadData(const QByteArray &, const QUrl &)

/*
Loads the QML given in data. The object tree defined by data is instantiated immediately.

If a url is specified it is used as the base url of the component. This affects relative paths within the data and error messages.

If an error occurs, error messages are printed with qWarning.
*/
func (this *QQmlApplicationEngine) LoadDatap(data qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg1 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine8loadDataERK10QByteArrayRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlapplicationengine.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void objectCreated(QObject *, const QUrl &)

/*
This signal is emitted when an object finishes loading. If loading was successful, object contains a pointer to the loaded object, otherwise the pointer is NULL.

The url to the component the object came from is also provided.

Note: If the path to the component was provided as a QString containing a relative path, the url will contain a fully resolved path to the file.
*/
func (this *QQmlApplicationEngine) ObjectCreated(object qtcore.QObject_ITF /*777 QObject **/, url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QQmlApplicationEngine13objectCreatedEP7QObjectRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
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
}

//  keep block end
