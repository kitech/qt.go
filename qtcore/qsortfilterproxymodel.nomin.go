//  header block begin

// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qsortfilterproxymodel.h
// #include <qsortfilterproxymodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression filterRegularExpression() const

/*

 */
func (this *QSortFilterProxyModel) FilterRegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel23filterRegularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterRegularExpression(const QString &)

/*
Sets the regular expression used to filter the contents of the source model to pattern.

This method should be preferred for new code as it will use QRegularExpression internally.

This function was introduced in  Qt 5.12.

Note: Setter function for property filterRegularExpression.

See also setFilterCaseSensitivity(), setFilterWildcard(), setFilterFixedString(), and filterRegularExpression().
*/
func (this *QSortFilterProxyModel) SetFilterRegularExpression(pattern string) {
	var tmpArg0 = NewQString5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel26setFilterRegularExpressionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:128
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFilterRegularExpression(const QRegularExpression &)

/*
Sets the regular expression used to filter the contents of the source model to pattern.

This method should be preferred for new code as it will use QRegularExpression internally.

This function was introduced in  Qt 5.12.

Note: Setter function for property filterRegularExpression.

See also setFilterCaseSensitivity(), setFilterWildcard(), setFilterFixedString(), and filterRegularExpression().
*/
func (this *QSortFilterProxyModel) SetFilterRegularExpression1(regularExpression QRegularExpression_ITF) {
	var convArg0 unsafe.Pointer
	if regularExpression != nil && regularExpression.QRegularExpression_PTR() != nil {
		convArg0 = regularExpression.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel26setFilterRegularExpressionERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10548() {
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
