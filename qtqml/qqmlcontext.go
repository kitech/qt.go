package qtqml

// /usr/include/qt/QtQml/qqmlcontext.h
// #include <qqmlcontext.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QQmlContext) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:65
// index:0
// Public
// void QQmlContext(QQmlEngine *, QObject *)
func NewQQmlContext(parent *QQmlEngine /*777 QQmlEngine **/, objParent *qtcore.QObject /*777 QObject **/) *QQmlContext {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	var convArg1 = objParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContextC2EP10QQmlEngineP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:66
// index:1
// Public
// void QQmlContext(QQmlContext *, QObject *)
func NewQQmlContext_1(parent *QQmlContext /*777 QQmlContext **/, objParent *qtcore.QObject /*777 QObject **/) *QQmlContext {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	var convArg1 = objParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContextC2EPS_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:67
// index:0
// Public virtual
// void ~QQmlContext()
func DeleteQQmlContext(*QQmlContext) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContextD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:69
// index:0
// Public
// bool isValid()
func (this *QQmlContext) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcontext.h:71
// index:0
// Public
// QQmlEngine * engine()
func (this *QQmlContext) Engine() *QQmlEngine /*777 QQmlEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext6engineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:72
// index:0
// Public
// QQmlContext * parentContext()
func (this *QQmlContext) ParentContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext13parentContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:74
// index:0
// Public
// QObject * contextObject()
func (this *QQmlContext) ContextObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext13contextObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:75
// index:0
// Public
// void setContextObject(QObject *)
func (this *QQmlContext) SetContextObject(arg0 *qtcore.QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContext16setContextObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:77
// index:0
// Public
// QVariant contextProperty(const QString &)
func (this *QQmlContext) ContextProperty(arg0 *qtcore.QString) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext15contextPropertyERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:78
// index:0
// Public
// void setContextProperty(const QString &, QObject *)
func (this *QQmlContext) SetContextProperty(arg0 *qtcore.QString, arg1 *qtcore.QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContext18setContextPropertyERK7QStringP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:79
// index:1
// Public
// void setContextProperty(const QString &, const QVariant &)
func (this *QQmlContext) SetContextProperty_1(arg0 *qtcore.QString, arg1 *qtcore.QVariant) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContext18setContextPropertyERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:82
// index:0
// Public
// QString nameForObject(QObject *)
func (this *QQmlContext) NameForObject(arg0 *qtcore.QObject /*777 QObject **/) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext13nameForObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:84
// index:0
// Public
// QUrl resolvedUrl(const QUrl &)
func (this *QQmlContext) ResolvedUrl(arg0 *qtcore.QUrl) *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContext11resolvedUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:86
// index:0
// Public
// void setBaseUrl(const QUrl &)
func (this *QQmlContext) SetBaseUrl(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQmlContext10setBaseUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:87
// index:0
// Public
// QUrl baseUrl()
func (this *QQmlContext) BaseUrl() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQmlContext7baseUrlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
