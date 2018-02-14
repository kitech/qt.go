package qtcore

// /usr/include/qt/QtCore/qhash.h
// #include <qhash.h>
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

type QHashData struct {
	*qtrt.CObject
}
type QHashData_ITF interface {
	QHashData_PTR() *QHashData
}

func (ptr *QHashData) QHashData_PTR() *QHashData { return ptr }

func (this *QHashData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHashData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQHashDataFromPointer(cthis unsafe.Pointer) *QHashData {
	return &QHashData{&qtrt.CObject{cthis}}
}
func (*QHashData) NewFromPointer(cthis unsafe.Pointer) *QHashData {
	return NewQHashDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qhash.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] void * allocateNode(int)
func (this *QHashData) AllocateNode(nodeAlign int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData12allocateNodeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nodeAlign)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qhash.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void freeNode(void *)
func (this *QHashData) FreeNode(node unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData8freeNodeEPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), node)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QHashData * detach_helper(void (*)(struct QHashData::Node *, void *), void (*)(struct QHashData::Node *), int, int)
func (this *QHashData) Detach_helper(node_duplicate unsafe.Pointer /*666*/, node_delete unsafe.Pointer /*666*/, nodeSize int, nodeAlign int) *QHashData /*777 QHashData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData13detach_helperEPFvPNS_4NodeEPvEPFvS1_Eii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), node_duplicate, node_delete, nodeSize, nodeAlign)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQHashDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qhash.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool willGrow()
func (this *QHashData) WillGrow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData8willGrowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qhash.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hasShrunk()
func (this *QHashData) HasShrunk() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData9hasShrunkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rehash(int)
func (this *QHashData) Rehash(hint int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData6rehashEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void free_helper(void (*)(struct QHashData::Node *))
func (this *QHashData) Free_helper(node_delete unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData11free_helperEPFvPNS_4NodeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), node_delete)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QHashData::Node * firstNode()
func (this *QHashData) FirstNode() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashData9firstNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

func DeleteQHashData(this *QHashData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHashDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
