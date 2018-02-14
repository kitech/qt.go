package qtqml

// /usr/include/qt/QtQml/qjsengine.h
// #include <qjsengine.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QJSEngine struct {
	*qtcore.QObject
}
type QJSEngine_ITF interface {
	qtcore.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine { return ptr }

func (this *QJSEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QJSEngine) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQJSEngineFromPointer(cthis unsafe.Pointer) *QJSEngine {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QJSEngine{bcthis0}
}
func (*QJSEngine) NewFromPointer(cthis unsafe.Pointer) *QJSEngine {
	return NewQJSEngineFromPointer(cthis)
}

// /usr/include/qt/QtQml/qjsengine.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QJSEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QJSEngine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qjsengine.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSEngine()
func NewQJSEngine() *QJSEngine {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngineC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qjsengine.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJSEngine(QObject *)
func NewQJSEngine_1(parent qtcore.QObject_ITF /*777 QObject **/) *QJSEngine {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qjsengine.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QJSEngine()
func DeleteQJSEngine(this *QJSEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qjsengine.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue globalObject()
func (this *QJSEngine) GlobalObject() *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QJSEngine12globalObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue evaluate(const QString &, const QString &, int)
func (this *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine8evaluateERK7QStringS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, lineNumber)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue newObject()
func (this *QJSEngine) NewObject() *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine9newObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue newArray(uint)
func (this *QJSEngine) NewArray(length uint) *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine8newArrayEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), length)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue newQObject(QObject *)
func (this *QJSEngine) NewQObject(object qtcore.QObject_ITF /*777 QObject **/) *QJSValue /*123*/ {
	var convArg0 = object.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine10newQObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collectGarbage()
func (this *QJSEngine) CollectGarbage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine14collectGarbageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installTranslatorFunctions(const QJSValue &)
func (this *QJSEngine) InstallTranslatorFunctions(object QJSValue_ITF) {
	var convArg0 = object.QJSValue_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine26installTranslatorFunctionsERK8QJSValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installExtensions(QJSEngine::Extensions, const QJSValue &)
func (this *QJSEngine) InstallExtensions(extensions int, object QJSValue_ITF) {
	var convArg1 = object.QJSValue_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine17installExtensionsE6QFlagsINS_9ExtensionEERK8QJSValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extensions, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QV8Engine * handle()
func (this *QJSEngine) Handle() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QJSEngine6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

type QJSEngine__Extension = int

const QJSEngine__TranslationExtension QJSEngine__Extension = 1
const QJSEngine__ConsoleExtension QJSEngine__Extension = 2
const QJSEngine__GarbageCollectionExtension QJSEngine__Extension = 4
const QJSEngine__AllExtensions QJSEngine__Extension = -1

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
}

//  keep block end
