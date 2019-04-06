//  header block begin

// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qdir.h
// #include <qdir.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 60
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qdir.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool match(const QStringList &, const QString &)

/*
Returns true if the fileName matches the wildcard (glob) pattern filter; otherwise returns false. The filter may contain multiple patterns separated by spaces or semicolons. The matching is case insensitive.

See also QRegularExpression Wildcard Matching, entryList(), and entryInfoList().
*/
func (this *QDir) Match(filters QStringList_ITF, fileName string) bool {
	var convArg0 unsafe.Pointer
	if filters != nil && filters.QStringList_PTR() != nil {
		convArg0 = filters.QStringList_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir5matchERK11QStringListRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_Match(filters QStringList_ITF, fileName string) bool {
	var nilthis *QDir
	rv := nilthis.Match(filters, fileName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:211
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool match(const QString &, const QString &)

/*
Returns true if the fileName matches the wildcard (glob) pattern filter; otherwise returns false. The filter may contain multiple patterns separated by spaces or semicolons. The matching is case insensitive.

See also QRegularExpression Wildcard Matching, entryList(), and entryInfoList().
*/
func (this *QDir) Match1(filter string, fileName string) bool {
	var tmpArg0 = NewQString5(filter)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir5matchERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_Match1(filter string, fileName string) bool {
	var nilthis *QDir
	rv := nilthis.Match1(filter, fileName)
	return rv
}

//  body block end

//  keep block begin

func init_unused_10398() {
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
