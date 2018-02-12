package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h
// #include <qwinjumplist.h>
// #include <Qtwinextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QWinJumpList struct {
	*qtcore.QObject
}

func (this *QWinJumpList) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWinJumpList) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWinJumpListFromPointer(cthis unsafe.Pointer) *QWinJumpList {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWinJumpList{bcthis0}
}
func (*QWinJumpList) NewFromPointer(cthis unsafe.Pointer) *QWinJumpList {
	return NewQWinJumpListFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QWinJumpList) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QWinJumpList10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinJumpList(QObject *)
func NewQWinJumpList(parent *qtcore.QObject /*777 QObject **/) *QWinJumpList {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWinJumpListC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinJumpListFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinJumpList()
func DeleteQWinJumpList(this *QWinJumpList) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWinJumpListD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QString identifier()
func (this *QWinJumpList) Identifier() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QWinJumpList10identifierEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIdentifier(const QString &)
func (this *QWinJumpList) SetIdentifier(identifier string) {
	var tmpArg0 = qtcore.NewQString_5(identifier)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWinJumpList13setIdentifierERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QWinJumpListCategory * recent()
func (this *QWinJumpList) Recent() *QWinJumpListCategory /*777 QWinJumpListCategory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QWinJumpList6recentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWinJumpListCategoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QWinJumpListCategory * frequent()
func (this *QWinJumpList) Frequent() *QWinJumpListCategory /*777 QWinJumpListCategory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QWinJumpList8frequentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWinJumpListCategoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QWinJumpListCategory * tasks()
func (this *QWinJumpList) Tasks() *QWinJumpListCategory /*777 QWinJumpListCategory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QWinJumpList5tasksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWinJumpListCategoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCategory(QWinJumpListCategory *)
func (this *QWinJumpList) AddCategory(category *QWinJumpListCategory /*777 QWinJumpListCategory **/) {
	var convArg0 = category.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWinJumpList11addCategoryEP20QWinJumpListCategory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinjumplist.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QWinJumpList) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QWinJumpList5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
