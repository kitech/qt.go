//  header block begin
// /usr/include/qt/QtCore/qstringlist.h
// #include <qstringlist.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QStringList struct {
	*qtrt.CObject
}

func (this *QStringList) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQStringListFromPointer(cthis unsafe.Pointer) *QStringList {
	return &QStringList{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstringlist.h:105
// index:0
// Public inline
// void QStringList()
func NewQStringList() *QStringList {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringListC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:106
// index:1
// Public inline
// void QStringList(const class QString &)
func NewQStringList_1(i *QString) *QStringList {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = i.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringListC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:122
// index:0
// Public inline
// bool contains(const class QString &, Qt::CaseSensitivity)
func (this *QStringList) Contains(str *QString, cs int) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList8containsERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlist.h:123
// index:1
// Public inline
// bool contains(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringList) Contains_1(str *QLatin1String, cs int) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList8containsE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlist.h:135
// index:0
// Public inline
// int indexOf(const class QRegExp &, int)
func (this *QStringList) IndexOf(rx *QRegExp, from int) interface{} {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlist.h:137
// index:1
// Public inline
// int indexOf(class QRegExp &, int)
func (this *QStringList) IndexOf_1(rx *QRegExp, from int) interface{} {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList7indexOfER7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlist.h:143
// index:2
// Public inline
// int indexOf(const class QRegularExpression &, int)
func (this *QStringList) IndexOf_2(re *QRegularExpression, from int) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK18QRegularExpressioni", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlist.h:136
// index:0
// Public inline
// int lastIndexOf(const class QRegExp &, int)
func (this *QStringList) LastIndexOf(rx *QRegExp, from int) interface{} {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlist.h:138
// index:1
// Public inline
// int lastIndexOf(class QRegExp &, int)
func (this *QStringList) LastIndexOf_1(rx *QRegExp, from int) interface{} {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfER7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlist.h:144
// index:2
// Public inline
// int lastIndexOf(const class QRegularExpression &, int)
func (this *QStringList) LastIndexOf_2(re *QRegularExpression, from int) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
