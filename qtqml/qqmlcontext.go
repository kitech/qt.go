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
// extern C begin: 28
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
type QQmlContext struct {
	*qtcore.QObject
}
type QQmlContext_ITF interface {
	qtcore.QObject_ITF
	QQmlContext_PTR() *QQmlContext
}

func (ptr *QQmlContext) QQmlContext_PTR() *QQmlContext { return ptr }

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

// /usr/include/qt/QtQml/qqmlcontext.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQmlContext) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlContext(QQmlEngine *, QObject *)

/*
Create a new QQmlContext as a child of engine's root context, and the QObject parent.
*/
func (*QQmlContext) NewForInherit(parent QQmlEngine_ITF /*777 QQmlEngine **/, objParent qtcore.QObject_ITF /*777 QObject **/) *QQmlContext {
	return NewQQmlContext(parent, objParent)
}
func NewQQmlContext(parent QQmlEngine_ITF /*777 QQmlEngine **/, objParent qtcore.QObject_ITF /*777 QObject **/) *QQmlContext {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQmlEngine_PTR() != nil {
		convArg0 = parent.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if objParent != nil && objParent.QObject_PTR() != nil {
		convArg1 = objParent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextC2EP10QQmlEngineP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlContext")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlContext(QQmlEngine *, QObject *)

/*
Create a new QQmlContext as a child of engine's root context, and the QObject parent.
*/
func (*QQmlContext) NewForInheritp(parent QQmlEngine_ITF /*777 QQmlEngine **/) *QQmlContext {
	return NewQQmlContextp(parent)
}
func NewQQmlContextp(parent QQmlEngine_ITF /*777 QQmlEngine **/) *QQmlContext {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQmlEngine_PTR() != nil {
		convArg0 = parent.QQmlEngine_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextC2EP10QQmlEngineP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlContext")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlContext(QQmlContext *, QObject *)

/*
Create a new QQmlContext as a child of engine's root context, and the QObject parent.
*/
func (*QQmlContext) NewForInherit1(parent QQmlContext_ITF /*777 QQmlContext **/, objParent qtcore.QObject_ITF /*777 QObject **/) *QQmlContext {
	return NewQQmlContext1(parent, objParent)
}
func NewQQmlContext1(parent QQmlContext_ITF /*777 QQmlContext **/, objParent qtcore.QObject_ITF /*777 QObject **/) *QQmlContext {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQmlContext_PTR() != nil {
		convArg0 = parent.QQmlContext_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if objParent != nil && objParent.QObject_PTR() != nil {
		convArg1 = objParent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextC2EPS_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlContext")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlContext(QQmlContext *, QObject *)

/*
Create a new QQmlContext as a child of engine's root context, and the QObject parent.
*/
func (*QQmlContext) NewForInherit1p(parent QQmlContext_ITF /*777 QQmlContext **/) *QQmlContext {
	return NewQQmlContext1p(parent)
}
func NewQQmlContext1p(parent QQmlContext_ITF /*777 QQmlContext **/) *QQmlContext {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQmlContext_PTR() != nil {
		convArg0 = parent.QQmlContext_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextC2EPS_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlContext")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcontext.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlContext()

/*

 */
func DeleteQQmlContext(this *QQmlContext) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContextD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlcontext.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns whether the context is valid.

To be valid, a context must have a engine, and it's contextObject(), if any, must not have been deleted.
*/
func (this *QQmlContext) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcontext.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlEngine * engine() const

/*
Return the context's QQmlEngine, or 0 if the context has no QQmlEngine or the QQmlEngine was destroyed.
*/
func (this *QQmlContext) Engine() *QQmlEngine /*777 QQmlEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext6engineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * parentContext() const

/*
Return the context's parent QQmlContext, or 0 if this context has no parent or if the parent has been destroyed.
*/
func (this *QQmlContext) ParentContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext13parentContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * contextObject() const

/*
Return the context object, or 0 if there is no context object.

See also setContextObject().
*/
func (this *QQmlContext) ContextObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext13contextObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcontext.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContextObject(QObject *)

/*
Set the context object.

See also contextObject().
*/
func (this *QQmlContext) SetContextObject(arg0 qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext16setContextObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:81
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant contextProperty(const QString &) const

/*
Returns the value of the name property for this context as a QVariant.

See also setContextProperty().
*/
func (this *QQmlContext) ContextProperty(arg0 string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext15contextPropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContextProperty(const QString &, QObject *)

/*
Set the value of the name property on this context.

QQmlContext does not take ownership of value.

See also contextProperty().
*/
func (this *QQmlContext) SetContextProperty(arg0 string, arg1 qtcore.QObject_ITF /*777 QObject **/) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QObject_PTR() != nil {
		convArg1 = arg1.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext18setContextPropertyERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setContextProperty(const QString &, const QVariant &)

/*
Set the value of the name property on this context.

QQmlContext does not take ownership of value.

See also contextProperty().
*/
func (this *QQmlContext) SetContextProperty1(arg0 string, arg1 qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QVariant_PTR() != nil {
		convArg1 = arg1.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext18setContextPropertyERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString nameForObject(QObject *) const

/*
Returns the name of object in this context, or an empty string if object is not named in the context. Objects are named by setContextProperty(), or by ids in the case of QML created contexts.

If the object has multiple names, the first is returned.
*/
func (this *QQmlContext) NameForObject(arg0 qtcore.QObject_ITF /*777 QObject **/) string {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext13nameForObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlcontext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl resolvedUrl(const QUrl &)

/*
Resolves the URL src relative to the URL of the containing component.

See also QQmlEngine::baseUrl() and setBaseUrl().
*/
func (this *QQmlContext) ResolvedUrl(arg0 qtcore.QUrl_ITF) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext11resolvedUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlcontext.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseUrl(const QUrl &)

/*
Explicitly sets the url resolvedUrl() will use for relative references to baseUrl.

Calling this function will override the url of the containing component used by default.

See also baseUrl() and resolvedUrl().
*/
func (this *QQmlContext) SetBaseUrl(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QQmlContext10setBaseUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcontext.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl baseUrl() const

/*
Returns the base url of the component, or the containing component if none is set.

See also setBaseUrl().
*/
func (this *QQmlContext) BaseUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QQmlContext7baseUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_11517() {
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
