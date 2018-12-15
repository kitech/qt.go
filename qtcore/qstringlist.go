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

/*

 */
type QStringList struct {
	*qtrt.CObject
}
type QStringList_ITF interface {
	QStringList_PTR() *QStringList
}

func (ptr *QStringList) QStringList_PTR() *QStringList { return ptr }

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

// /usr/include/qt/QtCore/qstringlist.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStringList()

/*
Constructs an empty string list.
*/
func (*QStringList) NewForInherit() *QStringList {
	return NewQStringList()
}
func NewQStringList() *QStringList {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringList)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:104
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QStringList(const QString &)

/*
Constructs an empty string list.
*/
func (*QStringList) NewForInherit1(i string) *QStringList {
	return NewQStringList1(i)
}
func NewQStringList1(i string) *QStringList {
	var tmpArg0 = NewQString5(i)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringList)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const

/*
Returns true if the list contains the string str; otherwise returns false. The search is case insensitive if cs is Qt::CaseInsensitive; the search is case sensitive by default.

See also indexOf(), lastIndexOf(), and QString::contains().
*/
func (this *QStringList) Contains(str string, cs int) bool {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const

/*
Returns true if the list contains the string str; otherwise returns false. The search is case insensitive if cs is Qt::CaseInsensitive; the search is case sensitive by default.

See also indexOf(), lastIndexOf(), and QString::contains().
*/
func (this *QStringList) Containsp(str string) bool {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:123
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if the list contains the string str; otherwise returns false. The search is case insensitive if cs is Qt::CaseInsensitive; the search is case sensitive by default.

See also indexOf(), lastIndexOf(), and QString::contains().
*/
func (this *QStringList) Contains1(str QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:123
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if the list contains the string str; otherwise returns false. The search is case insensitive if cs is Qt::CaseInsensitive; the search is case sensitive by default.

See also indexOf(), lastIndexOf(), and QString::contains().
*/
func (this *QStringList) Contains1p(str QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:124
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QStringView, Qt::CaseSensitivity) const

/*
Returns true if the list contains the string str; otherwise returns false. The search is case insensitive if cs is Qt::CaseInsensitive; the search is case sensitive by default.

See also indexOf(), lastIndexOf(), and QString::contains().
*/
func (this *QStringList) Contains2(str QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringView_PTR() != nil {
		convArg0 = str.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:124
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QStringView, Qt::CaseSensitivity) const

/*
Returns true if the list contains the string str; otherwise returns false. The search is case insensitive if cs is Qt::CaseInsensitive; the search is case sensitive by default.

See also indexOf(), lastIndexOf(), and QString::contains().
*/
func (this *QStringList) Contains2p(str QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringView_PTR() != nil {
		convArg0 = str.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringList operator+(const QStringList &) const

/*

 */
func (this *QStringList) Operator_add(other QStringList_ITF) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStringList_PTR() != nil {
		convArg0 = other.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringListplERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstringlist.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringList & operator<<(const QString &)

/*

 */
func (this *QStringList) Operator_left_shift(str string) *QStringList {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListlsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstringlist.h:130
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QStringList & operator<<(const QStringList &)

/*

 */
func (this *QStringList) Operator_left_shift1(l QStringList_ITF) *QStringList {
	var convArg0 unsafe.Pointer
	if l != nil && l.QStringList_PTR() != nil {
		convArg0 = l.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListlsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstringlist.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegExp &, int) const

/*
Returns the index position of the first exact match of rx in the list, searching forward from index position from. Returns -1 if no item matched.

See also lastIndexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) IndexOf(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegExp &, int) const

/*
Returns the index position of the first exact match of rx in the list, searching forward from index position from. Returns -1 if no item matched.

See also lastIndexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) IndexOfp(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:138
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(QRegExp &, int) const

/*
Returns the index position of the first exact match of rx in the list, searching forward from index position from. Returns -1 if no item matched.

See also lastIndexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) IndexOf1(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:138
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(QRegExp &, int) const

/*
Returns the index position of the first exact match of rx in the list, searching forward from index position from. Returns -1 if no item matched.

See also lastIndexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) IndexOf1p(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
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

// /usr/include/qt/QtCore/qstringlist.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegExp &, int) const

/*
Returns the index position of the last exact match of rx in the list, searching backward from index position from. If from is -1 (the default), the search starts at the last item. Returns -1 if no item matched.

See also indexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) LastIndexOf(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegExp &, int) const

/*
Returns the index position of the last exact match of rx in the list, searching backward from index position from. If from is -1 (the default), the search starts at the last item. Returns -1 if no item matched.

See also indexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) LastIndexOfp(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:139
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(QRegExp &, int) const

/*
Returns the index position of the last exact match of rx in the list, searching backward from index position from. If from is -1 (the default), the search starts at the last item. Returns -1 if no item matched.

See also indexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) LastIndexOf1(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:139
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(QRegExp &, int) const

/*
Returns the index position of the last exact match of rx in the list, searching backward from index position from. If from is -1 (the default), the search starts at the last item. Returns -1 if no item matched.

See also indexOf(), contains(), and QRegExp::exactMatch().
*/
func (this *QStringList) LastIndexOf1p(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
