package qtcore

// /usr/include/qt/QtCore/qlist.h
// #include <qlist.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QListData struct {
	*qtrt.CObject
}
type QListData_ITF interface {
	QListData_PTR() *QListData
}

func (ptr *QListData) QListData_PTR() *QListData { return ptr }

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
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData6detachEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alloc)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QListData::Data * detach_grow(int *, int)
func (this *QListData) Detach_grow(i unsafe.Pointer /*666*/, n int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData11detach_growEPii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, n)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void realloc(int)
func (this *QListData) Realloc(alloc int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData7reallocEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alloc)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void realloc_grow(int)
func (this *QListData) Realloc_grow(growth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData12realloc_growEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), growth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void dispose()
func (this *QListData) Dispose() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData7disposeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] void ** erase(void **)
func (this *QListData) Erase(xi unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData5eraseEPPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), xi)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(int)
func (this *QListData) Remove(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData6removeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:111
// index:1
// Public Visibility=Default Availability=Available
// [-2] void remove(int, int)
func (this *QListData) Remove_1(i int, n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData6removeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void move(int, int)
func (this *QListData) Move(from int, to int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListData4moveEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlist.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const
func (this *QListData) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListData4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlist.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QListData) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListData7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlist.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void ** at(int) const
func (this *QListData) At(i int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListData2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void ** begin() const
func (this *QListData) Begin() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListData5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlist.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void ** end() const
func (this *QListData) End() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListData3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

func DeleteQListData(this *QListData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QListData__ = int

const QListData__DataHeaderSize QListData__ = 16

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
