package qtcore

// /usr/include/qt/QtCore/qstringlist.h
// #include <qstringlist.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QStringList struct {
	*qtrt.CObject
}

func (this *QStringList) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringList) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringListFromPointer(cthis unsafe.Pointer) *QStringList {
	return &QStringList{&qtrt.CObject{cthis}}
}
func (*QStringList) NewFromPointer(cthis unsafe.Pointer) *QStringList {
	return NewQStringListFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringlist.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStringList()
func NewQStringList() *QStringList {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringList)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:106
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QStringList(const QString &)
func NewQStringList_1(i string) *QStringList {
	var tmpArg0 = NewQString_5(i)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringList)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity)
func (this *QStringList) Contains(str string, cs int) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:123
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity)
func (this *QStringList) Contains_1(str *QLatin1String /*123*/, cs int) bool {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegExp &, int)
func (this *QStringList) IndexOf(rx *QRegExp, from int) int {
	var convArg0 = rx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:137
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(QRegExp &, int)
func (this *QStringList) IndexOf_1(rx *QRegExp, from int) int {
	var convArg0 = rx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:143
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int)
func (this *QStringList) IndexOf_2(re *QRegularExpression, from int) int {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegExp &, int)
func (this *QStringList) LastIndexOf(rx *QRegExp, from int) int {
	var convArg0 = rx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:138
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(QRegExp &, int)
func (this *QStringList) LastIndexOf_1(rx *QRegExp, from int) int {
	var convArg0 = rx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:144
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int)
func (this *QStringList) LastIndexOf_2(re *QRegularExpression, from int) int {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQStringList(this *QStringList) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
}

//  keep block end
