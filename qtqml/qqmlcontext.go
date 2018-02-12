package qtqml

// /usr/include/qt/QtQml/qqmlcontext.h
// #include <qqmlcontext.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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

type QQmlContext struct {
	*qtcore.QObject
}

func (this *QQmlContext) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQmlContext) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQmlContextFromPointer(cthis unsafe.Pointer) *QQmlContext {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQmlContext{bcthis0}
}
func (*QQmlContext) NewFromPointer(cthis unsafe.Pointer) *QQmlContext {
	return NewQQmlContextFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlcontext.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQmlContext) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlContext(QQmlEngine *, QObject *)
func NewQQmlContext(parent *QQmlEngine /*777 QQmlEngine **/, objParent *qtcore.QObject /*777 QObject **/) *QQmlContext {
	var convArg0 = parent.GetCthis()
	var convArg1 = objParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextC2EP10QQmlEngineP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlContext(QQmlContext *, QObject *)
func NewQQmlContext_1(parent *QQmlContext /*777 QQmlContext **/, objParent *qtcore.QObject /*777 QObject **/) *QQmlContext {
	var convArg0 = parent.GetCthis()
	var convArg1 = objParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextC2EPS_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlContext()
func DeleteQQmlContext(this *QQmlContext) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlcontext.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QQmlContext) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcontext.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlEngine * engine()
func (this *QQmlContext) Engine() *QQmlEngine /*777 QQmlEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext6engineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * parentContext()
func (this *QQmlContext) ParentContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext13parentContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * contextObject()
func (this *QQmlContext) ContextObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext13contextObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContextObject(QObject *)
func (this *QQmlContext) SetContextObject(arg0 *qtcore.QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext16setContextObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:77
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant contextProperty(const QString &)
func (this *QQmlContext) ContextProperty(arg0 string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext15contextPropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContextProperty(const QString &, QObject *)
func (this *QQmlContext) SetContextProperty(arg0 string, arg1 *qtcore.QObject /*777 QObject **/) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext18setContextPropertyERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:79
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setContextProperty(const QString &, const QVariant &)
func (this *QQmlContext) SetContextProperty_1(arg0 string, arg1 *qtcore.QVariant) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext18setContextPropertyERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString nameForObject(QObject *)
func (this *QQmlContext) NameForObject(arg0 *qtcore.QObject /*777 QObject **/) string {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext13nameForObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlcontext.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl resolvedUrl(const QUrl &)
func (this *QQmlContext) ResolvedUrl(arg0 *qtcore.QUrl) *qtcore.QUrl /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext11resolvedUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseUrl(const QUrl &)
func (this *QQmlContext) SetBaseUrl(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext10setBaseUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl baseUrl()
func (this *QQmlContext) BaseUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext7baseUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
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
