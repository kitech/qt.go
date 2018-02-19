package qtqml

// /usr/include/qt/QtQml/qqmlexpression.h
// #include <qqmlexpression.h>
// #include <QtQml>

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

//  ext block end

//  body block begin

type QQmlExpression struct {
	*qtcore.QObject
}
type QQmlExpression_ITF interface {
	qtcore.QObject_ITF
	QQmlExpression_PTR() *QQmlExpression
}

func (ptr *QQmlExpression) QQmlExpression_PTR() *QQmlExpression { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QQmlExpression) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlexpression.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlExpression()
func NewQQmlExpression() *QQmlExpression {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExpression")
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlExpression(QQmlContext *, QObject *, const QString &, QObject *)
func NewQQmlExpression_1(arg0 QQmlContext_ITF /*777 QQmlContext **/, arg1 qtcore.QObject_ITF /*777 QObject **/, arg2 string, arg3 qtcore.QObject_ITF /*777 QObject **/) *QQmlExpression {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlContext_PTR() != nil {
		convArg0 = arg0.QQmlContext_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QObject_PTR() != nil {
		convArg1 = arg1.QObject_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(arg2)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if arg3 != nil && arg3.QObject_PTR() != nil {
		convArg3 = arg3.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionC2EP11QQmlContextP7QObjectRK7QStringS3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExpression")
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlExpression(QQmlContext *, QObject *, const QString &, QObject *)
func NewQQmlExpression_1_(arg0 QQmlContext_ITF /*777 QQmlContext **/, arg1 qtcore.QObject_ITF /*777 QObject **/, arg2 string) *QQmlExpression {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlContext_PTR() != nil {
		convArg0 = arg0.QQmlContext_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QObject_PTR() != nil {
		convArg1 = arg1.QObject_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(arg2)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QObject *=Pointer, QObject=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionC2EP11QQmlContextP7QObjectRK7QStringS3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExpression")
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:64
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlExpression(const QQmlScriptString &, QQmlContext *, QObject *, QObject *)
func NewQQmlExpression_2(arg0 QQmlScriptString_ITF, arg1 QQmlContext_ITF /*777 QQmlContext **/, arg2 qtcore.QObject_ITF /*777 QObject **/, arg3 qtcore.QObject_ITF /*777 QObject **/) *QQmlExpression {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlScriptString_PTR() != nil {
		convArg0 = arg0.QQmlScriptString_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QQmlContext_PTR() != nil {
		convArg1 = arg1.QQmlContext_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if arg2 != nil && arg2.QObject_PTR() != nil {
		convArg2 = arg2.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if arg3 != nil && arg3.QObject_PTR() != nil {
		convArg3 = arg3.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionC2ERK16QQmlScriptStringP11QQmlContextP7QObjectS6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExpression")
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:64
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlExpression(const QQmlScriptString &, QQmlContext *, QObject *, QObject *)
func NewQQmlExpression_2_(arg0 QQmlScriptString_ITF) *QQmlExpression {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlScriptString_PTR() != nil {
		convArg0 = arg0.QQmlScriptString_PTR().GetCthis()
	}
	// arg: 1, QQmlContext *=Pointer, QQmlContext=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, QObject *=Pointer, QObject=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, QObject *=Pointer, QObject=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionC2ERK16QQmlScriptStringP11QQmlContextP7QObjectS6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExpression")
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:64
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlExpression(const QQmlScriptString &, QQmlContext *, QObject *, QObject *)
func NewQQmlExpression_2_1(arg0 QQmlScriptString_ITF, arg1 QQmlContext_ITF /*777 QQmlContext **/) *QQmlExpression {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlScriptString_PTR() != nil {
		convArg0 = arg0.QQmlScriptString_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QQmlContext_PTR() != nil {
		convArg1 = arg1.QQmlContext_PTR().GetCthis()
	}
	// arg: 2, QObject *=Pointer, QObject=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, QObject *=Pointer, QObject=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionC2ERK16QQmlScriptStringP11QQmlContextP7QObjectS6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExpression")
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:64
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlExpression(const QQmlScriptString &, QQmlContext *, QObject *, QObject *)
func NewQQmlExpression_2_2(arg0 QQmlScriptString_ITF, arg1 QQmlContext_ITF /*777 QQmlContext **/, arg2 qtcore.QObject_ITF /*777 QObject **/) *QQmlExpression {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlScriptString_PTR() != nil {
		convArg0 = arg0.QQmlScriptString_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QQmlContext_PTR() != nil {
		convArg1 = arg1.QQmlContext_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if arg2 != nil && arg2.QObject_PTR() != nil {
		convArg2 = arg2.QObject_PTR().GetCthis()
	}
	// arg: 3, QObject *=Pointer, QObject=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionC2ERK16QQmlScriptStringP11QQmlContextP7QObjectS6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExpression")
	return gothis
}

// /usr/include/qt/QtQml/qqmlexpression.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlExpression()
func DeleteQQmlExpression(this *QQmlExpression) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpressionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlexpression.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlEngine * engine() const
func (this *QQmlExpression) Engine() *QQmlEngine /*777 QQmlEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression6engineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlexpression.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * context() const
func (this *QQmlExpression) Context() *QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression7contextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlexpression.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString expression() const
func (this *QQmlExpression) Expression() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression10expressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlexpression.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpression(const QString &)
func (this *QQmlExpression) SetExpression(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression13setExpressionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool notifyOnValueChanged() const
func (this *QQmlExpression) NotifyOnValueChanged() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression20notifyOnValueChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlexpression.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotifyOnValueChanged(_Bool)
func (this *QQmlExpression) SetNotifyOnValueChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression23setNotifyOnValueChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sourceFile() const
func (this *QQmlExpression) SourceFile() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression10sourceFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlexpression.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineNumber() const
func (this *QQmlExpression) LineNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression10lineNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlexpression.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnNumber() const
func (this *QQmlExpression) ColumnNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression12columnNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlexpression.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceLocation(const QString &, int, int)
func (this *QQmlExpression) SetSourceLocation(fileName string, line int, column int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression17setSourceLocationERK7QStringii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, line, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceLocation(const QString &, int, int)
func (this *QQmlExpression) SetSourceLocation__(fileName string, line int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression17setSourceLocationERK7QStringii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, line, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * scopeObject() const
func (this *QQmlExpression) ScopeObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression11scopeObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlexpression.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasError() const
func (this *QQmlExpression) HasError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression8hasErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlexpression.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearError()
func (this *QQmlExpression) ClearError() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression10clearErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlexpression.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlError error() const
func (this *QQmlExpression) Error() *QQmlError /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlExpression5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlErrorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlError)
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant evaluate(_Bool *)
func (this *QQmlExpression) Evaluate(valueIsUndefined *bool) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression8evaluateEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), valueIsUndefined)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant evaluate(_Bool *)
func (this *QQmlExpression) Evaluate__() *qtcore.QVariant /*123*/ {
	// arg: 0, bool *=Pointer, =Invalid,
	var valueIsUndefined unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression8evaluateEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), valueIsUndefined)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlexpression.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged()
func (this *QQmlExpression) ValueChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlExpression12valueChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
