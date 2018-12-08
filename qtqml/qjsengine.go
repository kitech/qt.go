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
// extern C begin: 46
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QJSEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QJSEngine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qjsengine.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSEngine()

/*
Constructs a QJSEngine object.

The globalObject() is initialized to have properties as described in ECMA-262, Section 15.1.
*/
func (*QJSEngine) NewForInherit() *QJSEngine {
	return NewQJSEngine()
}
func NewQJSEngine() *QJSEngine {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngineC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QJSEngine")
	return gothis
}

// /usr/include/qt/QtQml/qjsengine.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJSEngine(QObject *)

/*
Constructs a QJSEngine object.

The globalObject() is initialized to have properties as described in ECMA-262, Section 15.1.
*/
func (*QJSEngine) NewForInherit1(parent qtcore.QObject_ITF /*777 QObject **/) *QJSEngine {
	return NewQJSEngine1(parent)
}
func NewQJSEngine1(parent qtcore.QObject_ITF /*777 QObject **/) *QJSEngine {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QJSEngine")
	return gothis
}

// /usr/include/qt/QtQml/qjsengine.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QJSEngine()

/*

 */
func DeleteQJSEngine(this *QJSEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qjsengine.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue globalObject() const

/*
Returns this engine's Global Object.

By default, the Global Object contains the built-in objects that are part of ECMA-262, such as Math, Date and String. Additionally, you can set properties of the Global Object to make your own extensions available to all script code. Non-local variables in script code will be created as properties of the Global Object, as well as local variables in global code.
*/
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

/*
Evaluates program, using lineNumber as the base line number, and returns the result of the evaluation.

The script code will be evaluated in the context of the global object.

The evaluation of program can cause an exception in the engine; in this case the return value will be the exception that was thrown (typically an Error object; see QJSValue::isError()).

lineNumber is used to specify a starting line number for program; line number information reported by the engine that pertains to this evaluation will be based on this argument. For example, if program consists of two lines of code, and the statement on the second line causes a script exception, the exception line number would be lineNumber plus one. When no starting line number is specified, line numbers will be 1-based.

fileName is used for error reporting. For example, in error objects the file name is accessible through the "fileName" property if it is provided with this function.

Note: If an exception was thrown and the exception value is not an Error instance (i.e., QJSValue::isError() returns false), the exception value will still be returned, but there is currently no API for detecting that an exception did occur in this case.
*/
func (this *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue /*123*/ {
	var tmpArg0 = qtcore.NewQString5(program)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine8evaluateERK7QStringS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, lineNumber)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue evaluate(const QString &, const QString &, int)

/*
Evaluates program, using lineNumber as the base line number, and returns the result of the evaluation.

The script code will be evaluated in the context of the global object.

The evaluation of program can cause an exception in the engine; in this case the return value will be the exception that was thrown (typically an Error object; see QJSValue::isError()).

lineNumber is used to specify a starting line number for program; line number information reported by the engine that pertains to this evaluation will be based on this argument. For example, if program consists of two lines of code, and the statement on the second line causes a script exception, the exception line number would be lineNumber plus one. When no starting line number is specified, line numbers will be 1-based.

fileName is used for error reporting. For example, in error objects the file name is accessible through the "fileName" property if it is provided with this function.

Note: If an exception was thrown and the exception value is not an Error instance (i.e., QJSValue::isError() returns false), the exception value will still be returned, but there is currently no API for detecting that an exception did occur in this case.
*/
func (this *QJSEngine) Evaluatep(program string) *QJSValue /*123*/ {
	var tmpArg0 = qtcore.NewQString5(program)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, int=Int, =Invalid, , Invalid
	lineNumber := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine8evaluateERK7QStringS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, lineNumber)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue evaluate(const QString &, const QString &, int)

/*
Evaluates program, using lineNumber as the base line number, and returns the result of the evaluation.

The script code will be evaluated in the context of the global object.

The evaluation of program can cause an exception in the engine; in this case the return value will be the exception that was thrown (typically an Error object; see QJSValue::isError()).

lineNumber is used to specify a starting line number for program; line number information reported by the engine that pertains to this evaluation will be based on this argument. For example, if program consists of two lines of code, and the statement on the second line causes a script exception, the exception line number would be lineNumber plus one. When no starting line number is specified, line numbers will be 1-based.

fileName is used for error reporting. For example, in error objects the file name is accessible through the "fileName" property if it is provided with this function.

Note: If an exception was thrown and the exception value is not an Error instance (i.e., QJSValue::isError() returns false), the exception value will still be returned, but there is currently no API for detecting that an exception did occur in this case.
*/
func (this *QJSEngine) Evaluatep1(program string, fileName string) *QJSValue /*123*/ {
	var tmpArg0 = qtcore.NewQString5(program)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, int=Int, =Invalid, , Invalid
	lineNumber := int(1)
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

/*
Creates a JavaScript object of class Object.

The prototype of the created object will be the Object prototype object.

See also newArray() and QJSValue::setProperty().
*/
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

/*
Creates a JavaScript object of class Array with the given length.

See also newObject().
*/
func (this *QJSEngine) NewArray(length uint) *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine8newArrayEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), length)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsengine.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue newArray(uint)

/*
Creates a JavaScript object of class Array with the given length.

See also newObject().
*/
func (this *QJSEngine) NewArrayp() *QJSValue /*123*/ {
	// arg: 0, uint=Typedef, uint=Typedef, unsigned int, UInt
	length := uint(0)
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

/*
Creates a JavaScript object that wraps the given QObject object, using JavaScriptOwnership.

Signals and slots, properties and children of object are available as properties of the created QJSValue.

If object is a null pointer, this function returns a null value.

If a default prototype has been registered for the object's class (or its superclass, recursively), the prototype of the new script object will be set to be that default prototype.

If the given object is deleted outside of the engine's control, any attempt to access the deleted QObject's members through the JavaScript wrapper object (either by script code or C++) will result in a script exception.

See also QJSValue::toQObject().
*/
func (this *QJSEngine) NewQObject(object qtcore.QObject_ITF /*777 QObject **/) *QJSValue /*123*/ {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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

/*
Runs the garbage collector.

The garbage collector will attempt to reclaim memory by locating and disposing of objects that are no longer reachable in the script environment.

Normally you don't need to call this function; the garbage collector will automatically be invoked when the QJSEngine decides that it's wise to do so (i.e. when a certain number of new objects have been created). However, you can call this function to explicitly request that garbage collection should be performed as soon as possible.
*/
func (this *QJSEngine) CollectGarbage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine14collectGarbageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installTranslatorFunctions(const QJSValue &)

/*

 */
func (this *QJSEngine) InstallTranslatorFunctions(object QJSValue_ITF) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QJSValue_PTR() != nil {
		convArg0 = object.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine26installTranslatorFunctionsERK8QJSValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installTranslatorFunctions(const QJSValue &)

/*

 */
func (this *QJSEngine) InstallTranslatorFunctionsp() {
	// arg: 0, const QJSValue &=LValueReference, QJSValue=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine26installTranslatorFunctionsERK8QJSValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installExtensions(QJSEngine::Extensions, const QJSValue &)

/*
Installs JavaScript extensions to add functionality that is not available in a standard ECMAScript implementation.

The extensions are installed on the given object, or on the Global Object if no object is specified.

Several extensions can be installed at once by OR-ing the enum values:


  installExtensions(QJSEngine::TranslationExtension | QJSEngine::ConsoleExtension);



This function was introduced in  Qt 5.6.

See also Extension.
*/
func (this *QJSEngine) InstallExtensions(extensions int, object QJSValue_ITF) {
	var convArg1 unsafe.Pointer
	if object != nil && object.QJSValue_PTR() != nil {
		convArg1 = object.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine17installExtensionsE6QFlagsINS_9ExtensionEERK8QJSValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extensions, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installExtensions(QJSEngine::Extensions, const QJSValue &)

/*
Installs JavaScript extensions to add functionality that is not available in a standard ECMAScript implementation.

The extensions are installed on the given object, or on the Global Object if no object is specified.

Several extensions can be installed at once by OR-ing the enum values:


  installExtensions(QJSEngine::TranslationExtension | QJSEngine::ConsoleExtension);



This function was introduced in  Qt 5.6.

See also Extension.
*/
func (this *QJSEngine) InstallExtensionsp(extensions int) {
	// arg: 1, const QJSValue &=LValueReference, QJSValue=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QJSEngine17installExtensionsE6QFlagsINS_9ExtensionEERK8QJSValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extensions, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsengine.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QV8Engine * handle() const

/*

 */
func (this *QJSEngine) Handle() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QJSEngine6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

/*


 */
type QJSEngine__Extension = int

//
const QJSEngine__TranslationExtension QJSEngine__Extension = 1

//
const QJSEngine__ConsoleExtension QJSEngine__Extension = 2

//
const QJSEngine__GarbageCollectionExtension QJSEngine__Extension = 4

//
const QJSEngine__AllExtensions QJSEngine__Extension = -1

func (this *QJSEngine) ExtensionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QJSEngine_ExtensionItemName(val int) string {
	var nilthis *QJSEngine
	return nilthis.ExtensionItemName(val)
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
