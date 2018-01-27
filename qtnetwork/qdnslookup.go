package qtnetwork

// /usr/include/qt/QtNetwork/qdnslookup.h
// #include <qdnslookup.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QDnsLookup struct {
	*qtcore.QObject
}

func (this *QDnsLookup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QDnsLookup) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQDnsLookupFromPointer(cthis unsafe.Pointer) *QDnsLookup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDnsLookup{bcthis0}
}
func (*QDnsLookup) NewFromPointer(cthis unsafe.Pointer) *QDnsLookup {
	return NewQDnsLookupFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:186
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDnsLookup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDnsLookup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:221
// index:0
// Public
// void QDnsLookup(QObject *)
func NewQDnsLookup(parent *qtcore.QObject /*777 QObject **/) *QDnsLookup {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:222
// index:1
// Public
// void QDnsLookup(QDnsLookup::Type, const QString &, QObject *)
func NewQDnsLookup_1(type_ int, name *qtcore.QString, parent *qtcore.QObject /*777 QObject **/) *QDnsLookup {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = name.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookupC2ENS_4TypeERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, type_, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:223
// index:2
// Public
// void QDnsLookup(QDnsLookup::Type, const QString &, const QHostAddress &, QObject *)
func NewQDnsLookup_2(type_ int, name *qtcore.QString, nameserver *QHostAddress, parent *qtcore.QObject /*777 QObject **/) *QDnsLookup {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = name.GetCthis()
	var convArg2 = nameserver.GetCthis()
	var convArg3 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookupC2ENS_4TypeERK7QStringRK12QHostAddressP7QObject", ffiqt.FFI_TYPE_VOID, cthis, type_, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:224
// index:0
// Public virtual
// void ~QDnsLookup()
func DeleteQDnsLookup(*QDnsLookup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:226
// index:0
// Public
// QDnsLookup::Error error()
func (this *QDnsLookup) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDnsLookup5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:227
// index:0
// Public
// QString errorString()
func (this *QDnsLookup) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDnsLookup11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:228
// index:0
// Public
// bool isFinished()
func (this *QDnsLookup) IsFinished() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDnsLookup10isFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdnslookup.h:230
// index:0
// Public
// QString name()
func (this *QDnsLookup) Name() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDnsLookup4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:231
// index:0
// Public
// void setName(const QString &)
func (this *QDnsLookup) SetName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup7setNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:233
// index:0
// Public
// QDnsLookup::Type type()
func (this *QDnsLookup) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDnsLookup4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:234
// index:0
// Public
// void setType(QDnsLookup::Type)
func (this *QDnsLookup) SetType(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup7setTypeENS_4TypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:236
// index:0
// Public
// QHostAddress nameserver()
func (this *QDnsLookup) Nameserver() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDnsLookup10nameserverEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:237
// index:0
// Public
// void setNameserver(const QHostAddress &)
func (this *QDnsLookup) SetNameserver(nameserver *QHostAddress) {
	var convArg0 = nameserver.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup13setNameserverERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:249
// index:0
// Public
// void abort()
func (this *QDnsLookup) Abort() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup5abortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:250
// index:0
// Public
// void lookup()
func (this *QDnsLookup) Lookup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup6lookupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:253
// index:0
// Public
// void finished()
func (this *QDnsLookup) Finished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup8finishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:254
// index:0
// Public
// void nameChanged(const QString &)
func (this *QDnsLookup) NameChanged(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup11nameChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:255
// index:0
// Public
// void typeChanged(QDnsLookup::Type)
func (this *QDnsLookup) TypeChanged(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup11typeChangedENS_4TypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:256
// index:0
// Public
// void nameserverChanged(const QHostAddress &)
func (this *QDnsLookup) NameserverChanged(nameserver *QHostAddress) {
	var convArg0 = nameserver.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDnsLookup17nameserverChangedERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QDnsLookup__Error = int

const QDnsLookup__NoError QDnsLookup__Error = 0
const QDnsLookup__ResolverError QDnsLookup__Error = 1
const QDnsLookup__OperationCancelledError QDnsLookup__Error = 2
const QDnsLookup__InvalidRequestError QDnsLookup__Error = 3
const QDnsLookup__InvalidReplyError QDnsLookup__Error = 4
const QDnsLookup__ServerFailureError QDnsLookup__Error = 5
const QDnsLookup__ServerRefusedError QDnsLookup__Error = 6
const QDnsLookup__NotFoundError QDnsLookup__Error = 7

type QDnsLookup__Type = int

const QDnsLookup__A QDnsLookup__Type = 1
const QDnsLookup__AAAA QDnsLookup__Type = 28
const QDnsLookup__ANY QDnsLookup__Type = 255
const QDnsLookup__CNAME QDnsLookup__Type = 5
const QDnsLookup__MX QDnsLookup__Type = 15
const QDnsLookup__NS QDnsLookup__Type = 2
const QDnsLookup__PTR QDnsLookup__Type = 12
const QDnsLookup__SRV QDnsLookup__Type = 33
const QDnsLookup__TXT QDnsLookup__Type = 16

//  body block end
