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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qurlquery.h:59
// index:0
// void QUrlQuery()
func NewQUrlQuery() *QUrlQuery {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QUrlQuery{cthis}
}

// /usr/include/qt/QtCore/qurlquery.h:60
// index:1
// void QUrlQuery(const class QUrl &)
func NewQUrlQuery_1(url unsafe.Pointer) *QUrlQuery {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, url)
	gopp.ErrPrint(err, rv)
	return &QUrlQuery{cthis}
}

// /usr/include/qt/QtCore/qurlquery.h:61
// index:2
// void QUrlQuery(const class QString &)
func NewQUrlQuery_2(queryString unsafe.Pointer) *QUrlQuery {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, queryString)
	gopp.ErrPrint(err, rv)
	return &QUrlQuery{cthis}
}

// /usr/include/qt/QtCore/qurlquery.h:67
// index:0
// void ~QUrlQuery()
func DeleteQUrlQuery(*QUrlQuery) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQueryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:73
// index:0
// inline
// void swap(class QUrlQuery &)
func (this *QUrlQuery) Swap(other unsafe.Pointer) {
	// 0: (, QUrlQuery & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:75
// index:0
// bool isEmpty()
func (this *QUrlQuery) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:76
// index:0
// bool isDetached()
func (this *QUrlQuery) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:77
// index:0
// void clear()
func (this *QUrlQuery) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:80
// index:0
// void setQuery(const class QString &)
func (this *QUrlQuery) SetQuery(queryString unsafe.Pointer) {
	// 0: (, const QString & queryString), (queryString)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery8setQueryERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, queryString)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:84
// index:0
// void setQueryDelimiters(class QChar, class QChar)
func (this *QUrlQuery) SetQueryDelimiters(valueDelimiter unsafe.Pointer, pairDelimiter unsafe.Pointer) {
	// 0: (, QChar valueDelimiter, QChar pairDelimiter), (valueDelimiter, pairDelimiter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery18setQueryDelimitersE5QCharS0_", ffiqt.FFI_TYPE_VOID, this.cthis, valueDelimiter, pairDelimiter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:85
// index:0
// QChar queryValueDelimiter()
func (this *QUrlQuery) QueryValueDelimiter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery19queryValueDelimiterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:86
// index:0
// QChar queryPairDelimiter()
func (this *QUrlQuery) QueryPairDelimiter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery18queryPairDelimiterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:91
// index:0
// bool hasQueryItem(const class QString &)
func (this *QUrlQuery) HasQueryItem(key unsafe.Pointer) {
	// 0: (, const QString & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QUrlQuery12hasQueryItemERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:92
// index:0
// void addQueryItem(const class QString &, const class QString &)
func (this *QUrlQuery) AddQueryItem(key unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, const QString & key, const QString & value), (key, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery12addQueryItemERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.cthis, key, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:93
// index:0
// void removeQueryItem(const class QString &)
func (this *QUrlQuery) RemoveQueryItem(key unsafe.Pointer) {
	// 0: (, const QString & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery15removeQueryItemERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:96
// index:0
// void removeAllQueryItems(const class QString &)
func (this *QUrlQuery) RemoveAllQueryItems(key unsafe.Pointer) {
	// 0: (, const QString & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery19removeAllQueryItemsERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:98
// index:0
// static inline
// QChar defaultQueryValueDelimiter()
func (this *QUrlQuery) DefaultQueryValueDelimiter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery26defaultQueryValueDelimiterEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrlQuery_DefaultQueryValueDelimiter() {
	// 0: (), ()
	var nilthis *QUrlQuery
	nilthis.DefaultQueryValueDelimiter()
}

// /usr/include/qt/QtCore/qurlquery.h:100
// index:0
// static inline
// QChar defaultQueryPairDelimiter()
func (this *QUrlQuery) DefaultQueryPairDelimiter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QUrlQuery25defaultQueryPairDelimiterEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrlQuery_DefaultQueryPairDelimiter() {
	// 0: (), ()
	var nilthis *QUrlQuery
	nilthis.DefaultQueryPairDelimiter()
}

//  body block end
