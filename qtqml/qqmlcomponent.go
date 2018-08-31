package qtqml

// /usr/include/qt/QtQml/qqmlcomponent.h
// #include <qqmlcomponent.h>
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

/*

 */
type QQmlComponent struct {
	*qtcore.QObject
}
type QQmlComponent_ITF interface {
	qtcore.QObject_ITF
	QQmlComponent_PTR() *QQmlComponent
}

func (ptr *QQmlComponent) QQmlComponent_PTR() *QQmlComponent { return ptr }

func (this *QQmlComponent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQmlComponent) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQmlComponentFromPointer(cthis unsafe.Pointer) *QQmlComponent {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQmlComponent{bcthis0}
}
func (*QQmlComponent) NewFromPointer(cthis unsafe.Pointer) *QQmlComponent {
	return NewQQmlComponentFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQmlComponent) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	return NewQQmlComponent(parent)
}
func NewQQmlComponent(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit__() *QQmlComponent {
	return NewQQmlComponent__()
}
func NewQQmlComponent__() *QQmlComponent {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_1(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	return NewQQmlComponent_1(arg0, parent)
}
func NewQQmlComponent_1(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_1_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/) *QQmlComponent {
	return NewQQmlComponent_1_(arg0)
}
func NewQQmlComponent_1_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:82
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QString &, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_2(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	return NewQQmlComponent_2(arg0, fileName, parent)
}
func NewQQmlComponent_2(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:82
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QString &, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_2_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string) *QQmlComponent {
	return NewQQmlComponent_2_(arg0, fileName)
}
func NewQQmlComponent_2_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:83
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QString &, QQmlComponent::CompilationMode, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_3(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	return NewQQmlComponent_3(arg0, fileName, mode, parent)
}
func NewQQmlComponent_3(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	var convArg3 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg3 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringNS_15CompilationModeEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:83
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QString &, QQmlComponent::CompilationMode, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_3_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, mode int) *QQmlComponent {
	return NewQQmlComponent_3_(arg0, fileName, mode)
}
func NewQQmlComponent_3_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, mode int) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, QObject *=Pointer, QObject=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringNS_15CompilationModeEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:84
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QUrl &, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_4(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	return NewQQmlComponent_4(arg0, url, parent)
}
func NewQQmlComponent_4(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:84
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QUrl &, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_4_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF) *QQmlComponent {
	return NewQQmlComponent_4_(arg0, url)
}
func NewQQmlComponent_4_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:85
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QUrl &, QQmlComponent::CompilationMode, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_5(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	return NewQQmlComponent_5(arg0, url, mode, parent)
}
func NewQQmlComponent_5(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg3 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlNS_15CompilationModeEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:85
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QUrl &, QQmlComponent::CompilationMode, QObject *)

/*
Create a QQmlComponent with no data and give it the specified engine and parent. Set the data with setData().
*/
func (*QQmlComponent) NewForInherit_5_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, mode int) *QQmlComponent {
	return NewQQmlComponent_5_(arg0, url, mode)
}
func NewQQmlComponent_5_(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, mode int) *QQmlComponent {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlEngine_PTR() != nil {
		convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
	// arg: 3, QObject *=Pointer, QObject=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlNS_15CompilationModeEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlComponent")
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlComponent()

/*

 */
func DeleteQQmlComponent(this *QQmlComponent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QQmlComponent::Status status() const

/*

 */
func (this *QQmlComponent) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if status() == QQmlComponent::Null.
*/
func (this *QQmlComponent) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReady() const

/*
Returns true if status() == QQmlComponent::Ready.
*/
func (this *QQmlComponent) IsReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent7isReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isError() const

/*
Returns true if status() == QQmlComponent::Error.
*/
func (this *QQmlComponent) IsError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent7isErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoading() const

/*
Returns true if status() == QQmlComponent::Loading.
*/
func (this *QQmlComponent) IsLoading() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent9isLoadingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*

 */
func (this *QQmlComponent) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlcomponent.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal progress() const

/*

 */
func (this *QQmlComponent) Progress() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent8progressEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQml/qqmlcomponent.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QQmlComponent) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlcomponent.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * create(QQmlContext *)

/*
Create an object instance from this component. Returns 0 if creation failed. context specifies the context within which to create the object instance.

If context is 0 (the default), it will create the instance in the engine' s root context.

The ownership of the returned object instance is transferred to the caller.

If the object being created from this component is a visual item, it must have a visual parent, which can be set by calling QQuickItem::setParentItem(). See Concepts - Visual Parent in Qt Quick for more details.

See also QQmlEngine::ObjectOwnership.
*/
func (this *QQmlComponent) Create(context QQmlContext_ITF /*777 QQmlContext **/) *qtcore.QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if context != nil && context.QQmlContext_PTR() != nil {
		convArg0 = context.QQmlContext_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent6createEP11QQmlContext", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * create(QQmlContext *)

/*
Create an object instance from this component. Returns 0 if creation failed. context specifies the context within which to create the object instance.

If context is 0 (the default), it will create the instance in the engine' s root context.

The ownership of the returned object instance is transferred to the caller.

If the object being created from this component is a visual item, it must have a visual parent, which can be set by calling QQuickItem::setParentItem(). See Concepts - Visual Parent in Qt Quick for more details.

See also QQmlEngine::ObjectOwnership.
*/
func (this *QQmlComponent) Create__() *qtcore.QObject /*777 QObject **/ {
	// arg: 0, QQmlContext *=Pointer, QQmlContext=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent6createEP11QQmlContext", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:108
// index:1
// Public Visibility=Default Availability=Available
// [-2] void create(QQmlIncubator &, QQmlContext *, QQmlContext *)

/*
Create an object instance from this component. Returns 0 if creation failed. context specifies the context within which to create the object instance.

If context is 0 (the default), it will create the instance in the engine' s root context.

The ownership of the returned object instance is transferred to the caller.

If the object being created from this component is a visual item, it must have a visual parent, which can be set by calling QQuickItem::setParentItem(). See Concepts - Visual Parent in Qt Quick for more details.

See also QQmlEngine::ObjectOwnership.
*/
func (this *QQmlComponent) Create_1(arg0 QQmlIncubator_ITF, context QQmlContext_ITF /*777 QQmlContext **/, forContext QQmlContext_ITF /*777 QQmlContext **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlIncubator_PTR() != nil {
		convArg0 = arg0.QQmlIncubator_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if context != nil && context.QQmlContext_PTR() != nil {
		convArg1 = context.QQmlContext_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if forContext != nil && forContext.QQmlContext_PTR() != nil {
		convArg2 = forContext.QQmlContext_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent6createER13QQmlIncubatorP11QQmlContextS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:108
// index:1
// Public Visibility=Default Availability=Available
// [-2] void create(QQmlIncubator &, QQmlContext *, QQmlContext *)

/*
Create an object instance from this component. Returns 0 if creation failed. context specifies the context within which to create the object instance.

If context is 0 (the default), it will create the instance in the engine' s root context.

The ownership of the returned object instance is transferred to the caller.

If the object being created from this component is a visual item, it must have a visual parent, which can be set by calling QQuickItem::setParentItem(). See Concepts - Visual Parent in Qt Quick for more details.

See also QQmlEngine::ObjectOwnership.
*/
func (this *QQmlComponent) Create_1_(arg0 QQmlIncubator_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlIncubator_PTR() != nil {
		convArg0 = arg0.QQmlIncubator_PTR().GetCthis()
	}
	// arg: 1, QQmlContext *=Pointer, QQmlContext=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QQmlContext *=Pointer, QQmlContext=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent6createER13QQmlIncubatorP11QQmlContextS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:108
// index:1
// Public Visibility=Default Availability=Available
// [-2] void create(QQmlIncubator &, QQmlContext *, QQmlContext *)

/*
Create an object instance from this component. Returns 0 if creation failed. context specifies the context within which to create the object instance.

If context is 0 (the default), it will create the instance in the engine' s root context.

The ownership of the returned object instance is transferred to the caller.

If the object being created from this component is a visual item, it must have a visual parent, which can be set by calling QQuickItem::setParentItem(). See Concepts - Visual Parent in Qt Quick for more details.

See also QQmlEngine::ObjectOwnership.
*/
func (this *QQmlComponent) Create_1_1(arg0 QQmlIncubator_ITF, context QQmlContext_ITF /*777 QQmlContext **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlIncubator_PTR() != nil {
		convArg0 = arg0.QQmlIncubator_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if context != nil && context.QQmlContext_PTR() != nil {
		convArg1 = context.QQmlContext_PTR().GetCthis()
	}
	// arg: 2, QQmlContext *=Pointer, QQmlContext=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent6createER13QQmlIncubatorP11QQmlContextS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * beginCreate(QQmlContext *)

/*
This method provides advanced control over component instance creation. In general, programmers should use QQmlComponent::create() to create object instances.

Create an object instance from this component. Returns 0 if creation failed. publicContext specifies the context within which to create the object instance.

When QQmlComponent constructs an instance, it occurs in three steps:
*/
func (this *QQmlComponent) BeginCreate(arg0 QQmlContext_ITF /*777 QQmlContext **/) *qtcore.QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlContext_PTR() != nil {
		convArg0 = arg0.QQmlContext_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent11beginCreateEP11QQmlContext", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void completeCreate()

/*
This method provides advanced control over component instance creation. In general, programmers should use QQmlComponent::create() to create a component.

This function completes the component creation begun with QQmlComponent::beginCreate() and must be called afterwards.

See also beginCreate().
*/
func (this *QQmlComponent) CompleteCreate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent14completeCreateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * creationContext() const

/*
Returns the QQmlContext the component was created in. This is only valid for components created directly from QML.
*/
func (this *QQmlComponent) CreationContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent15creationContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [8] QQmlComponentAttached * qmlAttachedProperties(QObject *)

/*

 */
func (this *QQmlComponent) QmlAttachedProperties(arg0 qtcore.QObject_ITF /*777 QObject **/) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent21qmlAttachedPropertiesEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QQmlComponent_QmlAttachedProperties(arg0 qtcore.QObject_ITF /*777 QObject **/) unsafe.Pointer /*666*/ {
	var nilthis *QQmlComponent
	rv := nilthis.QmlAttachedProperties(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlcomponent.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadUrl(const QUrl &)

/*
Load the QQmlComponent from the provided url.

Ensure that the URL provided is full and correct, in particular, use QUrl::fromLocalFile() when loading a file from the local filesystem.

Relative paths will be resolved against the engine's baseUrl(), which is the current working directory unless specified.
*/
func (this *QQmlComponent) LoadUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent7loadUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:117
// index:1
// Public Visibility=Default Availability=Available
// [-2] void loadUrl(const QUrl &, QQmlComponent::CompilationMode)

/*
Load the QQmlComponent from the provided url.

Ensure that the URL provided is full and correct, in particular, use QUrl::fromLocalFile() when loading a file from the local filesystem.

Relative paths will be resolved against the engine's baseUrl(), which is the current working directory unless specified.
*/
func (this *QQmlComponent) LoadUrl_1(url qtcore.QUrl_ITF, mode int) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent7loadUrlERK4QUrlNS_15CompilationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QByteArray &, const QUrl &)

/*
Sets the QQmlComponent to use the given QML data. If url is provided, it is used to set the component name and to provide a base path for items resolved by this component.
*/
func (this *QQmlComponent) SetData(arg0 qtcore.QByteArray_ITF, baseUrl qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if baseUrl != nil && baseUrl.QUrl_PTR() != nil {
		convArg1 = baseUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent7setDataERK10QByteArrayRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QQmlComponent::Status)

/*
Emitted whenever the component's status changes. status will be the new status.

Note: Notifier signal for property status.
*/
func (this *QQmlComponent) StatusChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent13statusChangedENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressChanged(qreal)

/*
Emitted whenever the component's loading progress changes. progress will be the current progress between 0.0 (nothing loaded) and 1.0 (finished).

Note: Notifier signal for property progress.
*/
func (this *QQmlComponent) ProgressChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent15progressChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

/*
Specifies whether the QQmlComponent should load the component immediately, or asynchonously.


*/
type QQmlComponent__CompilationMode = int

// Prefer loading/compiling the component immediately, blocking the thread. This is not always possible; for example, remote URLs will always load asynchronously.
const QQmlComponent__PreferSynchronous QQmlComponent__CompilationMode = 0

// Load/compile the component in a background thread.
const QQmlComponent__Asynchronous QQmlComponent__CompilationMode = 1

func (this *QQmlComponent) CompilationModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQmlComponent_CompilationModeItemName(val int) string {
	var nilthis *QQmlComponent
	return nilthis.CompilationModeItemName(val)
}

/*
Specifies the loading status of the QQmlComponent.


*/
type QQmlComponent__Status = int

// This QQmlComponent has no data. Call loadUrl() or setData() to add QML content.
const QQmlComponent__Null QQmlComponent__Status = 0

// This QQmlComponent is ready and create() may be called.
const QQmlComponent__Ready QQmlComponent__Status = 1

// This QQmlComponent is loading network data.
const QQmlComponent__Loading QQmlComponent__Status = 2

// An error has occurred. Call errors() to retrieve a list of errors.
const QQmlComponent__Error QQmlComponent__Status = 3

func (this *QQmlComponent) StatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQmlComponent_StatusItemName(val int) string {
	var nilthis *QQmlComponent
	return nilthis.StatusItemName(val)
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
