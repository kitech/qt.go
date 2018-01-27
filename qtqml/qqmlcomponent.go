package qtqml

// /usr/include/qt/QtQml/qqmlcomponent.h
// #include <qqmlcomponent.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QQmlComponent struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QQmlComponent) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcomponent.h:80
// index:0
// Public
// void QQmlComponent(QObject *)
func NewQQmlComponent(parent *qtcore.QObject /*777 QObject **/) *QQmlComponent {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponentC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:81
// index:1
// Public
// void QQmlComponent(QQmlEngine *, QObject *)
func NewQQmlComponent_1(arg0 *QQmlEngine /*777 QQmlEngine **/, parent *qtcore.QObject /*777 QObject **/) *QQmlComponent {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:82
// index:2
// Public
// void QQmlComponent(QQmlEngine *, const QString &, QObject *)
func NewQQmlComponent_2(arg0 *QQmlEngine /*777 QQmlEngine **/, fileName *qtcore.QString, parent *qtcore.QObject /*777 QObject **/) *QQmlComponent {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	var convArg1 = fileName.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:83
// index:3
// Public
// void QQmlComponent(QQmlEngine *, const QString &, QQmlComponent::CompilationMode, QObject *)
func NewQQmlComponent_3(arg0 *QQmlEngine /*777 QQmlEngine **/, fileName *qtcore.QString, mode int, parent *qtcore.QObject /*777 QObject **/) *QQmlComponent {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	var convArg1 = fileName.GetCthis()
	var convArg3 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK7QStringNS_15CompilationModeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, mode, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:84
// index:4
// Public
// void QQmlComponent(QQmlEngine *, const QUrl &, QObject *)
func NewQQmlComponent_4(arg0 *QQmlEngine /*777 QQmlEngine **/, url *qtcore.QUrl, parent *qtcore.QObject /*777 QObject **/) *QQmlComponent {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	var convArg1 = url.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:85
// index:5
// Public
// void QQmlComponent(QQmlEngine *, const QUrl &, QQmlComponent::CompilationMode, QObject *)
func NewQQmlComponent_5(arg0 *QQmlEngine /*777 QQmlEngine **/, url *qtcore.QUrl, mode int, parent *qtcore.QObject /*777 QObject **/) *QQmlComponent {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	var convArg1 = url.GetCthis()
	var convArg3 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponentC2EP10QQmlEngineRK4QUrlNS_15CompilationModeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, mode, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlComponentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlcomponent.h:86
// index:0
// Public virtual
// void ~QQmlComponent()
func DeleteQQmlComponent(*QQmlComponent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:90
// index:0
// Public
// QQmlComponent::Status status()
func (this *QQmlComponent) Status() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:92
// index:0
// Public
// bool isNull()
func (this *QQmlComponent) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:93
// index:0
// Public
// bool isReady()
func (this *QQmlComponent) IsReady() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent7isReadyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:94
// index:0
// Public
// bool isError()
func (this *QQmlComponent) IsError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent7isErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:95
// index:0
// Public
// bool isLoading()
func (this *QQmlComponent) IsLoading() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent9isLoadingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlcomponent.h:98
// index:0
// Public
// QString errorString()
func (this *QQmlComponent) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlcomponent.h:100
// index:0
// Public
// qreal progress()
func (this *QQmlComponent) Progress() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent8progressEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtQml/qqmlcomponent.h:102
// index:0
// Public
// QUrl url()
func (this *QQmlComponent) Url() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent3urlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlcomponent.h:104
// index:0
// Public virtual
// QObject * create(QQmlContext *)
func (this *QQmlComponent) Create(context *QQmlContext /*777 QQmlContext **/) *qtcore.QObject /*777 QObject **/ {
	var convArg0 = context.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent6createEP11QQmlContext", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcomponent.h:108
// index:1
// Public
// void create(QQmlIncubator &, QQmlContext *, QQmlContext *)
func (this *QQmlComponent) Create_1(arg0 *QQmlIncubator, context *QQmlContext /*777 QQmlContext **/, forContext *QQmlContext /*777 QQmlContext **/) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = context.GetCthis()
	var convArg2 = forContext.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent6createER13QQmlIncubatorP11QQmlContextS3_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:105
// index:0
// Public virtual
// QObject * beginCreate(QQmlContext *)
func (this *QQmlComponent) BeginCreate(arg0 *QQmlContext /*777 QQmlContext **/) *qtcore.QObject /*777 QObject **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent11beginCreateEP11QQmlContext", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcomponent.h:106
// index:0
// Public virtual
// void completeCreate()
func (this *QQmlComponent) CompleteCreate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent14completeCreateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:111
// index:0
// Public
// QQmlContext * creationContext()
func (this *QQmlComponent) CreationContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlComponent15creationContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlcomponent.h:113
// index:0
// Public static
// QQmlComponentAttached * qmlAttachedProperties(QObject *)
func (this *QQmlComponent) QmlAttachedProperties(arg0 *qtcore.QObject /*777 QObject **/) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent21qmlAttachedPropertiesEP7QObject", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}
func QQmlComponent_QmlAttachedProperties(arg0 *qtcore.QObject /*777 QObject **/) unsafe.Pointer /*666*/ {
	var nilthis *QQmlComponent
	rv := nilthis.QmlAttachedProperties(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlcomponent.h:116
// index:0
// Public
// void loadUrl(const QUrl &)
func (this *QQmlComponent) LoadUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent7loadUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:117
// index:1
// Public
// void loadUrl(const QUrl &, QQmlComponent::CompilationMode)
func (this *QQmlComponent) LoadUrl_1(url *qtcore.QUrl, mode int) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent7loadUrlERK4QUrlNS_15CompilationModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:118
// index:0
// Public
// void setData(const QByteArray &, const QUrl &)
func (this *QQmlComponent) SetData(arg0 *qtcore.QByteArray, baseUrl *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = baseUrl.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent7setDataERK10QByteArrayRK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:121
// index:0
// Public
// void statusChanged(QQmlComponent::Status)
func (this *QQmlComponent) StatusChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent13statusChangedENS_6StatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlcomponent.h:122
// index:0
// Public
// void progressChanged(qreal)
func (this *QQmlComponent) ProgressChanged(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlComponent15progressChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
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
