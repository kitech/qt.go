package qtcore

// /usr/include/qt/QtCore/qlist.h
// #include <qlist.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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

type QListData struct {
	*qtrt.CObject
}

func (this *QListData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QListData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQListDataFromPointer(cthis unsafe.Pointer) *QListData {
	return &QListData{&qtrt.CObject{cthis}}
}
func (*QListData) NewFromPointer(cthis unsafe.Pointer) *QListData {
	return NewQListDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlist.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QListData::Data * detach(int)
func (this *QListData) Detach(alloc int) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData6detachEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alloc)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QListData::Data * detach_grow(int *, int)
func (this *QListData) Detach_grow(i unsafe.Pointer /*666*/, n int) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData11detach_growEPii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, n)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void realloc(int)
func (this *QListData) Realloc(alloc int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData7reallocEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alloc)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void realloc_grow(int)
func (this *QListData) Realloc_grow(growth int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData12realloc_growEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), growth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void dispose()
func (this *QListData) Dispose() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData7disposeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] void ** erase(void **)
func (this *QListData) Erase(xi unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData5eraseEPPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), xi)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(int)
func (this *QListData) Remove(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData6removeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:111
// index:1
// Public Visibility=Default Availability=Available
// [-2] void remove(int, int)
func (this *QListData) Remove_1(i int, n int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData6removeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i, n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void move(int, int)
func (this *QListData) Move(from int, to int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListData4moveEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size()
func (this *QListData) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListData4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlist.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QListData) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListData7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlist.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void ** at(int)
func (this *QListData) At(i int) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListData2atEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void ** begin()
func (this *QListData) Begin() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListData5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void ** end()
func (this *QListData) End() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListData3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

func DeleteQListData(this *QListData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListDataD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QListData__ = int

const QListData__DataHeaderSize QListData__ = 16

//  body block end
