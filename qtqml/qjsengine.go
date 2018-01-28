package qtqml

// /usr/include/qt/QtQml/qjsengine.h
// #include <qjsengine.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QJSEngine struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QJSEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QJSEngine10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:64
// index:0
// Public
// void QJSEngine()
func NewQJSEngine() *QJSEngine {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngineC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSEngineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsengine.h:65
// index:1
// Public
// void QJSEngine(QObject *)
func NewQJSEngine_1(parent *qtcore.QObject /*777 QObject **/) *QJSEngine {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngineC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSEngineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsengine.h:66
// index:0
// Public virtual
// void ~QJSEngine()
func DeleteQJSEngine(*QJSEngine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:68
// index:0
// Public
// QJSValue globalObject()
func (this *QJSEngine) GlobalObject() *QJSValue /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QJSEngine12globalObjectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:70
// index:0
// Public
// QJSValue evaluate(const QString &, const QString &, int)
func (this *QJSEngine) Evaluate(program *qtcore.QString, fileName *qtcore.QString, lineNumber int) *QJSValue /*123*/ {
	var convArg0 = program.GetCthis()
	var convArg1 = fileName.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngine8evaluateERK7QStringS2_i", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, lineNumber)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:72
// index:0
// Public
// QJSValue newObject()
func (this *QJSEngine) NewObject() *QJSValue /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngine9newObjectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:73
// index:0
// Public
// QJSValue newArray(uint)
func (this *QJSEngine) NewArray(length uint) *QJSValue /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngine8newArrayEj", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), length)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:75
// index:0
// Public
// QJSValue newQObject(QObject *)
func (this *QJSEngine) NewQObject(object *qtcore.QObject /*777 QObject **/) *QJSValue /*123*/ {
	var convArg0 = object.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngine10newQObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:96
// index:0
// Public
// void collectGarbage()
func (this *QJSEngine) CollectGarbage() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngine14collectGarbageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:99
// index:0
// Public
// void installTranslatorFunctions(const QJSValue &)
func (this *QJSEngine) InstallTranslatorFunctions(object *QJSValue) {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QJSEngine26installTranslatorFunctionsERK8QJSValue", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:112
// index:0
// Public inline
// QV8Engine * handle()
func (this *QJSEngine) Handle() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QJSEngine6handleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

type QJSEngine__Extension = int

const QJSEngine__TranslationExtension QJSEngine__Extension = 1
const QJSEngine__ConsoleExtension QJSEngine__Extension = 2
const QJSEngine__GarbageCollectionExtension QJSEngine__Extension = 4
const QJSEngine__AllExtensions QJSEngine__Extension = 4294967295

//  body block end
