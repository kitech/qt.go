package qtcore

// /usr/include/qt/QtCore/qurlquery.h
// #include <qurlquery.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 48
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QUrlQuery struct {
	*qtrt.CObject
}
type QUrlQuery_ITF interface {
	QUrlQuery_PTR() *QUrlQuery
}

func (ptr *QUrlQuery) QUrlQuery_PTR() *QUrlQuery { return ptr }

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

/*
Constructs an empty QUrlQuery object. A query can be set afterwards by calling setQuery() or items can be added by using addQueryItem().

See also setQuery() and addQueryItem().
*/
func (*QUrlQuery) NewForInherit() *QUrlQuery {
	return NewQUrlQuery()
}
func NewQUrlQuery() *QUrlQuery {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQueryC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrlQuery)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUrlQuery(const QUrl &)

/*
Constructs an empty QUrlQuery object. A query can be set afterwards by calling setQuery() or items can be added by using addQueryItem().

See also setQuery() and addQueryItem().
*/
func (*QUrlQuery) NewForInherit1(url QUrl_ITF) *QUrlQuery {
	return NewQUrlQuery1(url)
}
func NewQUrlQuery1(url QUrl_ITF) *QUrlQuery {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrlQuery)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUrlQuery(const QString &)

/*
Constructs an empty QUrlQuery object. A query can be set afterwards by calling setQuery() or items can be added by using addQueryItem().

See also setQuery() and addQueryItem().
*/
func (*QUrlQuery) NewForInherit2(queryString string) *QUrlQuery {
	return NewQUrlQuery2(queryString)
}
func NewQUrlQuery2(queryString string) *QUrlQuery {
	var tmpArg0 = NewQString5(queryString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQueryC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrlQuery)
	return gothis
}

// /usr/include/qt/QtCore/qurlquery.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrlQuery & operator=(const QUrlQuery &)

/*

 */
func (this *QUrlQuery) Operator_equal(other QUrlQuery_ITF) *QUrlQuery {
	var convArg0 unsafe.Pointer
	if other != nil && other.QUrlQuery_PTR() != nil {
		convArg0 = other.QUrlQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQueryaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrlQuery)
	return rv2
}

// /usr/include/qt/QtCore/qurlquery.h:65
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QUrlQuery & operator=(QUrlQuery &&)

/*

 */
func (this *QUrlQuery) Operator_equal1(other unsafe.Pointer /*333*/) *QUrlQuery {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQueryaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlQueryFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrlQuery)
	return rv2
}

// /usr/include/qt/QtCore/qurlquery.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QUrlQuery()

/*

 */
func DeleteQUrlQuery(this *QUrlQuery) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQueryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qurlquery.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QUrlQuery &) const

/*

 */
func (this *QUrlQuery) Operator_equal_equal(other QUrlQuery_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QUrlQuery_PTR() != nil {
		convArg0 = other.QUrlQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUrlQueryeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QUrlQuery &) const

/*

 */
func (this *QUrlQuery) Operator_not_equal(other QUrlQuery_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QUrlQuery_PTR() != nil {
		convArg0 = other.QUrlQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUrlQueryneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QUrlQuery &)

/*
Swaps this URL query instance with other. This function is very fast and never fails.
*/
func (this *QUrlQuery) Swap(other QUrlQuery_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QUrlQuery_PTR() != nil {
		convArg0 = other.QUrlQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if this QUrlQuery object contains no key-value pairs, such as after being default-constructed or after parsing an empty query string.

See also setQuery() and clear().
*/
func (this *QUrlQuery) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUrlQuery7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QUrlQuery) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUrlQuery10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears this QUrlQuery object by removing all of the key-value pairs currently stored. If the query delimiters have been changed, this function will leave them with their changed values.

See also isEmpty() and setQueryDelimiters().
*/
func (this *QUrlQuery) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QString &)

/*
Parses the query string in queryString and sets the internal items to the values found there. If any delimiters have been specified with setQueryDelimiters(), this function will use them instead of the default delimiters to parse the string.

See also query().
*/
func (this *QUrlQuery) SetQuery(queryString string) {
	var tmpArg0 = NewQString5(queryString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery8setQueryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQueryDelimiters(QChar, QChar)

/*
Sets the characters used for delimiting between keys and values, and between key-value pairs in the URL's query string. The default value delimiter is '=' and the default pair delimiter is '&'.



valueDelimiter will be used for separating keys from values, and pairDelimiter will be used to separate key-value pairs. Any occurrences of these delimiting characters in the encoded representation of the keys and values of the query string are percent encoded when returned in query().

If valueDelimiter is set to '(' and pairDelimiter is ')', the above query string would instead be represented like this:


  http://www.example.com/cgi-bin/drawgraph.cgi?type(pie)color(green)



Note: Non-standard delimiters should be chosen from among what RFC 3986 calls "sub-delimiters". They are:


  sub-delims    = "!" / "$" / "&" / "'" / "(" / ")"
                / "*" / "+" / "," / ";" / "="



Use of other characters is not supported and may result in unexpected behaviour. This method does not verify that you passed a valid delimiter.

See also queryValueDelimiter() and queryPairDelimiter().
*/
func (this *QUrlQuery) SetQueryDelimiters(valueDelimiter QChar_ITF /*123*/, pairDelimiter QChar_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if valueDelimiter != nil && valueDelimiter.QChar_PTR() != nil {
		convArg0 = valueDelimiter.QChar_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pairDelimiter != nil && pairDelimiter.QChar_PTR() != nil {
		convArg1 = pairDelimiter.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery18setQueryDelimitersE5QCharS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:85
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar queryValueDelimiter() const

/*
Returns the character used to delimit between keys and values when reconstructing the query string in query() or when parsing in setQuery().

See also setQueryDelimiters() and queryPairDelimiter().
*/
func (this *QUrlQuery) QueryValueDelimiter() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUrlQuery19queryValueDelimiterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qurlquery.h:86
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar queryPairDelimiter() const

/*
Returns the character used to delimit between keys-value pairs when reconstructing the query string in query() or when parsing in setQuery().

See also setQueryDelimiters() and queryValueDelimiter().
*/
func (this *QUrlQuery) QueryPairDelimiter() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUrlQuery18queryPairDelimiterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qurlquery.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasQueryItem(const QString &) const

/*
Returns true if there is a query string pair whose key is equal to key from the URL.

See also addQueryItem() and queryItemValue().
*/
func (this *QUrlQuery) HasQueryItem(key string) bool {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QUrlQuery12hasQueryItemERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurlquery.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addQueryItem(const QString &, const QString &)

/*
Appends the pair key = value to the end of the query string of the URL. This method does not overwrite existing items that might exist with the same key.

Note: This method does not treat spaces (ASCII 0x20) and plus ("+") signs as the same, like HTML forms do. If you need spaces to be represented as plus signs, use actual plus signs.

See also hasQueryItem() and queryItemValue().
*/
func (this *QUrlQuery) AddQueryItem(key string, value string) {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(value)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery12addQueryItemERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeQueryItem(const QString &)

/*
Removes the query string pair whose key is equal to key from the URL. If there are multiple items with a key equal to key, it removes the first item in the order they were present in the query string or added with addQueryItem().

See also removeAllQueryItems().
*/
func (this *QUrlQuery) RemoveQueryItem(key string) {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery15removeQueryItemERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAllQueryItems(const QString &)

/*
Removes all the query string pairs whose key is equal to key from the URL.

See also removeQueryItem().
*/
func (this *QUrlQuery) RemoveAllQueryItems(key string) {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery19removeAllQueryItemsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurlquery.h:98
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] QChar defaultQueryValueDelimiter()

/*
Returns the default character for separating keys from values in the query, an equal sign ("=").

See also setQueryDelimiters(), queryValueDelimiter(), and defaultQueryPairDelimiter().
*/
func (this *QUrlQuery) DefaultQueryValueDelimiter() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery26defaultQueryValueDelimiterEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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

/*
Returns the default character for separating keys-value pairs from each other, an ampersand ("&").

See also setQueryDelimiters(), queryPairDelimiter(), and defaultQueryValueDelimiter().
*/
func (this *QUrlQuery) DefaultQueryPairDelimiter() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QUrlQuery25defaultQueryPairDelimiterEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
}

//  keep block end
