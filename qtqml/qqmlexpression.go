package qtqml

// /usr/include/qt/QtQml/qqmlexpression.h
// #include <qqmlexpression.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QQmlExpression struct {
	*qtcore.QObject
}

func (this *QQmlExpression) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQmlExpression) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQmlExpressionFromPointer(cthis unsafe.Pointer) *QQmlExpression {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQmlExpression{bcthis0}
}
func (*QQmlExpression) NewFromPointer(cthis unsafe.Pointer) *QQmlExpression {
	return NewQQmlExpressionFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlexpression.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQmlExpression) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:62
// index:0
// Public
// void QQmlExpression()
func NewQQmlExpression() *QQmlExpression {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpressionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:63
// index:1
// Public
// void QQmlExpression(class QQmlContext *, class QObject *, const class QString &, class QObject *)
func NewQQmlExpression_1(arg0 *QQmlContext /*777 QQmlContext **/, arg1 *qtcore.QObject /*777 QObject **/, arg2 *qtcore.QString, arg3 *qtcore.QObject /*777 QObject **/) *QQmlExpression {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	var convArg2 = arg2.GetCthis()
	var convArg3 = arg3.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpressionC2EP11QQmlContextP7QObjectRK7QStringS3_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:64
// index:2
// Public
// void QQmlExpression(const class QQmlScriptString &, class QQmlContext *, class QObject *, class QObject *)
func NewQQmlExpression_2(arg0 *QQmlScriptString, arg1 *QQmlContext /*777 QQmlContext **/, arg2 *qtcore.QObject /*777 QObject **/, arg3 *qtcore.QObject /*777 QObject **/) *QQmlExpression {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	var convArg2 = arg2.GetCthis()
	var convArg3 = arg3.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpressionC2ERK16QQmlScriptStringP11QQmlContextP7QObjectS6_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:65
// index:0
// Public virtual
// void ~QQmlExpression()
func DeleteQQmlExpression(*QQmlExpression) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpressionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:67
// index:0
// Public
// QQmlEngine * engine()
func (this *QQmlExpression) Engine() *QQmlEngine /*777 QQmlEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression6engineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:68
// index:0
// Public
// QQmlContext * context()
func (this *QQmlExpression) Context() *QQmlContext /*777 QQmlContext **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression7contextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:70
// index:0
// Public
// QString expression()
func (this *QQmlExpression) Expression() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression10expressionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:71
// index:0
// Public
// void setExpression(const class QString &)
func (this *QQmlExpression) SetExpression(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpression13setExpressionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:73
// index:0
// Public
// bool notifyOnValueChanged()
func (this *QQmlExpression) NotifyOnValueChanged() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression20notifyOnValueChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlexpression.h:74
// index:0
// Public
// void setNotifyOnValueChanged(_Bool)
func (this *QQmlExpression) SetNotifyOnValueChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpression23setNotifyOnValueChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:76
// index:0
// Public
// QString sourceFile()
func (this *QQmlExpression) SourceFile() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression10sourceFileEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:77
// index:0
// Public
// int lineNumber()
func (this *QQmlExpression) LineNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression10lineNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQml/qqmlexpression.h:78
// index:0
// Public
// int columnNumber()
func (this *QQmlExpression) ColumnNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression12columnNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQml/qqmlexpression.h:79
// index:0
// Public
// void setSourceLocation(const class QString &, int, int)
func (this *QQmlExpression) SetSourceLocation(fileName *qtcore.QString, line int, column int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpression17setSourceLocationERK7QStringii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, line, column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:81
// index:0
// Public
// QObject * scopeObject()
func (this *QQmlExpression) ScopeObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression11scopeObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:83
// index:0
// Public
// bool hasError()
func (this *QQmlExpression) HasError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression8hasErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlexpression.h:84
// index:0
// Public
// void clearError()
func (this *QQmlExpression) ClearError() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpression10clearErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:85
// index:0
// Public
// QQmlError error()
func (this *QQmlExpression) Error() *QQmlError /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QQmlExpression5errorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQQmlErrorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:87
// index:0
// Public
// QVariant evaluate(_Bool *)
func (this *QQmlExpression) Evaluate(valueIsUndefined unsafe.Pointer /*666*/) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpression8evaluateEPb", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), &valueIsUndefined)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:90
// index:0
// Public
// void valueChanged()
func (this *QQmlExpression) ValueChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QQmlExpression12valueChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
