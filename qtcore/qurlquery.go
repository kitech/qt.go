package qtcore

// /usr/include/qt/QtCore/qurlquery.h
// #include <qurlquery.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QUrlQuery struct {
	*qtrt.CObject
}

func (this *QUrlQuery) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QUrlQuery) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQUrlQueryFromPointer(cthis unsafe.Pointer) *QUrlQuery {
	return &QUrlQuery{&qtrt.CObject{cthis}}
}
func (*QUrlQuery) NewFromPointer(cthis unsafe.Pointer) *QUrlQuery {
	return NewQUrlQueryFromPointer(cthis)
}

// /usr/include/qt/QtCore/qurlquery.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUrlQuery()
func NewQUrlQuery() *QUrlQuery {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrlQuery)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUrlQuery(const QUrl &)
func NewQUrlQuery_1(url *QUrl) *QUrlQuery {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK4QUrl", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrlQuery)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUrlQuery(const QString &)
func NewQUrlQuery_2(queryString *QString) *QUrlQuery {
	var convArg0 = queryString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrlQuery)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QUrlQuery()
func DeleteQUrlQuery(this *QUrlQuery) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qurlquery.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QUrlQuery &)
func (this *QUrlQuery) Swap(other *QUrlQuery) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QUrlQuery) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QUrlQuery) IsDetached() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QUrlQuery) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QString &)
func (this *QUrlQuery) SetQuery(queryString *QString) {
	var convArg0 = queryString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery8setQueryERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQueryDelimiters(QChar, QChar)
func (this *QUrlQuery) SetQueryDelimiters(valueDelimiter *QChar /*123*/, pairDelimiter *QChar /*123*/) {
	var convArg0 = valueDelimiter.GetCthis()
	var convArg1 = pairDelimiter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery18setQueryDelimitersE5QCharS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:85
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar queryValueDelimiter()
func (this *QUrlQuery) QueryValueDelimiter() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery19queryValueDelimiterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qurlquery.h:86
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar queryPairDelimiter()
func (this *QUrlQuery) QueryPairDelimiter() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery18queryPairDelimiterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qurlquery.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasQueryItem(const QString &)
func (this *QUrlQuery) HasQueryItem(key *QString) bool {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery12hasQueryItemERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addQueryItem(const QString &, const QString &)
func (this *QUrlQuery) AddQueryItem(key *QString, value *QString) {
	var convArg0 = key.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery12addQueryItemERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeQueryItem(const QString &)
func (this *QUrlQuery) RemoveQueryItem(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery15removeQueryItemERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAllQueryItems(const QString &)
func (this *QUrlQuery) RemoveAllQueryItems(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery19removeAllQueryItemsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:98
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] QChar defaultQueryValueDelimiter()
func (this *QUrlQuery) DefaultQueryValueDelimiter() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery26defaultQueryValueDelimiterEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}
func QUrlQuery_DefaultQueryValueDelimiter() *QChar /*123*/ {
	var nilthis *QUrlQuery
	rv := nilthis.DefaultQueryValueDelimiter()
	return rv
}

// /usr/include/qt/QtCore/qurlquery.h:100
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] QChar defaultQueryPairDelimiter()
func (this *QUrlQuery) DefaultQueryPairDelimiter() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery25defaultQueryPairDelimiterEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}
func QUrlQuery_DefaultQueryPairDelimiter() *QChar /*123*/ {
	var nilthis *QUrlQuery
	rv := nilthis.DefaultQueryPairDelimiter()
	return rv
}

//  body block end
