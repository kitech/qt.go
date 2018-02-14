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
// extern C begin: 9
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
// [8] const QMetaObject * metaObject()
func (this *QQmlComponent) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QObject *)
func NewQQmlComponent(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, QObject *)
func NewQQmlComponent_1(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	var convArg1 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:82
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QString &, QObject *)
func NewQQmlComponent_2(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:83
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QString &, enum QQmlComponent::CompilationMode, QObject *)
func NewQQmlComponent_3(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, fileName string, mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	var convArg3 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringNS_15CompilationModeEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:84
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QUrl &, QObject *)
func NewQQmlComponent_4(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	var convArg1 = url.QUrl_PTR().GetCthis()
	var convArg2 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:85
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QQmlComponent(QQmlEngine *, const QUrl &, enum QQmlComponent::CompilationMode, QObject *)
func NewQQmlComponent_5(arg0 QQmlEngine_ITF /*777 QQmlEngine **/, url qtcore.QUrl_ITF, mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QQmlComponent {
	var convArg0 = arg0.QQmlEngine_PTR().GetCthis()
	var convArg1 = url.QUrl_PTR().GetCthis()
	var convArg3 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlNS_15CompilationModeEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlComponent()
func DeleteQQmlComponent(this *QQmlComponent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QQmlComponent::Status status()
func (this *QQmlComponent) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QQmlComponent) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReady()
func (this *QQmlComponent) IsReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent7isReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isError()
func (this *QQmlComponent) IsError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent7isErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoading()
func (this *QQmlComponent) IsLoading() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent9isLoadingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
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
// [8] qreal progress()
func (this *QQmlComponent) Progress() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent8progressEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQml/qqmlcomponent.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
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
func (this *QQmlComponent) Create(context QQmlContext_ITF /*777 QQmlContext **/) *qtcore.QObject /*777 QObject **/ {
	var convArg0 = context.QQmlContext_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent6createEP11QQmlContext", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:108
// index:1
// Public Visibility=Default Availability=Available
// [-2] void create(QQmlIncubator &, QQmlContext *, QQmlContext *)
func (this *QQmlComponent) Create_1(arg0 QQmlIncubator_ITF, context QQmlContext_ITF /*777 QQmlContext **/, forContext QQmlContext_ITF /*777 QQmlContext **/) {
	var convArg0 = arg0.QQmlIncubator_PTR().GetCthis()
	var convArg1 = context.QQmlContext_PTR().GetCthis()
	var convArg2 = forContext.QQmlContext_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent6createER13QQmlIncubatorP11QQmlContextS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * beginCreate(QQmlContext *)
func (this *QQmlComponent) BeginCreate(arg0 QQmlContext_ITF /*777 QQmlContext **/) *qtcore.QObject /*777 QObject **/ {
	var convArg0 = arg0.QQmlContext_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent11beginCreateEP11QQmlContext", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void completeCreate()
func (this *QQmlComponent) CompleteCreate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent14completeCreateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * creationContext()
func (this *QQmlComponent) CreationContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlComponent15creationContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlcomponent.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [8] QQmlComponentAttached * qmlAttachedProperties(QObject *)
func (this *QQmlComponent) QmlAttachedProperties(arg0 qtcore.QObject_ITF /*777 QObject **/) unsafe.Pointer /*666*/ {
	var convArg0 = arg0.QObject_PTR().GetCthis()
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
func (this *QQmlComponent) LoadUrl(url qtcore.QUrl_ITF) {
	var convArg0 = url.QUrl_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent7loadUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:117
// index:1
// Public Visibility=Default Availability=Available
// [-2] void loadUrl(const QUrl &, enum QQmlComponent::CompilationMode)
func (this *QQmlComponent) LoadUrl_1(url qtcore.QUrl_ITF, mode int) {
	var convArg0 = url.QUrl_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent7loadUrlERK4QUrlNS_15CompilationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QByteArray &, const QUrl &)
func (this *QQmlComponent) SetData(arg0 qtcore.QByteArray_ITF, baseUrl qtcore.QUrl_ITF) {
	var convArg0 = arg0.QByteArray_PTR().GetCthis()
	var convArg1 = baseUrl.QUrl_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent7setDataERK10QByteArrayRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QQmlComponent::Status)
func (this *QQmlComponent) StatusChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent13statusChangedENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressChanged(qreal)
func (this *QQmlComponent) ProgressChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlComponent15progressChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

type QQmlComponent__CompilationMode = int

const QQmlComponent__PreferSynchronous QQmlComponent__CompilationMode = 0
const QQmlComponent__Asynchronous QQmlComponent__CompilationMode = 1

type QQmlComponent__Status = int

const QQmlComponent__Null QQmlComponent__Status = 0
const QQmlComponent__Ready QQmlComponent__Status = 1
const QQmlComponent__Loading QQmlComponent__Status = 2
const QQmlComponent__Error QQmlComponent__Status = 3

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
