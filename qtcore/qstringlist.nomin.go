//  header block begin

// +build !minimal

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
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qstringlist.h:143
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int) const

/*
Returns the index position of the first exact match of rx in the list, searching forward from index position from. Returns -1 if no item matched.

See also lastIndexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) IndexOf2(re QRegularExpression_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:143
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int) const

/*
Returns the index position of the first exact match of rx in the list, searching forward from index position from. Returns -1 if no item matched.

See also lastIndexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) IndexOf2p(re QRegularExpression_ITF) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:144
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int) const

/*
Returns the index position of the last exact match of rx in the list, searching backward from index position from. If from is -1 (the default), the search starts at the last item. Returns -1 if no item matched.

See also indexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) LastIndexOf2(re QRegularExpression_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:144
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int) const

/*
Returns the index position of the last exact match of rx in the list, searching backward from index position from. If from is -1 (the default), the search starts at the last item. Returns -1 if no item matched.

See also indexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) LastIndexOf2p(re QRegularExpression_ITF) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

func init_unused_10210() {
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
