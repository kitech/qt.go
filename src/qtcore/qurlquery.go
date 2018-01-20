//  header block begin
// /usr/include/qt/QtCore/qurlquery.h
// #include <qurlquery.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
	return this.Cthis
}
func NewQUrlQueryFromPointer(cthis unsafe.Pointer) *QUrlQuery {
	return &QUrlQuery{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qurlquery.h:59
// index:0
// Public
// void QUrlQuery()
func NewQUrlQuery() *QUrlQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:60
// index:1
// Public
// void QUrlQuery(const class QUrl &)
func NewQUrlQuery_1(url *QUrl) *QUrlQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:61
// index:2
// Public
// void QUrlQuery(const class QString &)
func NewQUrlQuery_2(queryString *QString) *QUrlQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = queryString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:67
// index:0
// Public
// void ~QUrlQuery()
func DeleteQUrlQuery(*QUrlQuery) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:73
// index:0
// Public inline
// void swap(class QUrlQuery &)
func (this *QUrlQuery) Swap(other *QUrlQuery) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:75
// index:0
// Public
// bool isEmpty()
func (this *QUrlQuery) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qurlquery.h:76
// index:0
// Public
// bool isDetached()
func (this *QUrlQuery) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qurlquery.h:77
// index:0
// Public
// void clear()
func (this *QUrlQuery) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:80
// index:0
// Public
// void setQuery(const class QString &)
func (this *QUrlQuery) SetQuery(queryString *QString) {
	var convArg0 = queryString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery8setQueryERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:84
// index:0
// Public
// void setQueryDelimiters(class QChar, class QChar)
func (this *QUrlQuery) SetQueryDelimiters(valueDelimiter *QChar, pairDelimiter *QChar) {
	var convArg0 = valueDelimiter.GetCthis()
	var convArg1 = pairDelimiter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery18setQueryDelimitersE5QCharS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:85
// index:0
// Public
// QChar queryValueDelimiter()
func (this *QUrlQuery) QueryValueDelimiter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery19queryValueDelimiterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qurlquery.h:86
// index:0
// Public
// QChar queryPairDelimiter()
func (this *QUrlQuery) QueryPairDelimiter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery18queryPairDelimiterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qurlquery.h:91
// index:0
// Public
// bool hasQueryItem(const class QString &)
func (this *QUrlQuery) HasQueryItem(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery12hasQueryItemERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qurlquery.h:92
// index:0
// Public
// void addQueryItem(const class QString &, const class QString &)
func (this *QUrlQuery) AddQueryItem(key *QString, value *QString) {
	var convArg0 = key.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery12addQueryItemERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:93
// index:0
// Public
// void removeQueryItem(const class QString &)
func (this *QUrlQuery) RemoveQueryItem(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery15removeQueryItemERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:96
// index:0
// Public
// void removeAllQueryItems(const class QString &)
func (this *QUrlQuery) RemoveAllQueryItems(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery19removeAllQueryItemsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:98
// index:0
// Public static inline
// QChar defaultQueryValueDelimiter()
func (this *QUrlQuery) DefaultQueryValueDelimiter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery26defaultQueryValueDelimiterEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUrlQuery_DefaultQueryValueDelimiter() {
	var nilthis *QUrlQuery
	nilthis.DefaultQueryValueDelimiter()
}

// /usr/include/qt/QtCore/qurlquery.h:100
// index:0
// Public static inline
// QChar defaultQueryPairDelimiter()
func (this *QUrlQuery) DefaultQueryPairDelimiter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery25defaultQueryPairDelimiterEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUrlQuery_DefaultQueryPairDelimiter() {
	var nilthis *QUrlQuery
	nilthis.DefaultQueryPairDelimiter()
}

//  body block end
