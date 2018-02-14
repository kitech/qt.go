package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextFragment struct {
	*qtrt.CObject
}
type QTextFragment_ITF interface {
	QTextFragment_PTR() *QTextFragment
}

func (ptr *QTextFragment) QTextFragment_PTR() *QTextFragment { return ptr }

func (this *QTextFragment) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextFragment) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextFragmentFromPointer(cthis unsafe.Pointer) *QTextFragment {
	return &QTextFragment{&qtrt.CObject{cthis}}
}
func (*QTextFragment) NewFromPointer(cthis unsafe.Pointer) *QTextFragment {
	return NewQTextFragmentFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:307
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextFragment()
func NewQTextFragment() *QTextFragment {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextFragmentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:309
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextFragment & operator=(const QTextFragment &)
func (this *QTextFragment) Operator_equal(o QTextFragment_ITF) *QTextFragment {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextFragment_PTR() != nil {
		convArg0 = o.QTextFragment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextFragmentaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFragment)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:311
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextFragment) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:313
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QTextFragment &)
func (this *QTextFragment) Operator_equal_equal(o QTextFragment_ITF) bool {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextFragment_PTR() != nil {
		convArg0 = o.QTextFragment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragmenteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:314
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QTextFragment &)
func (this *QTextFragment) Operator_not_equal(o QTextFragment_ITF) bool {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextFragment_PTR() != nil {
		convArg0 = o.QTextFragment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragmentneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:315
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QTextFragment &)
func (this *QTextFragment) Operator_less_than(o QTextFragment_ITF) bool {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextFragment_PTR() != nil {
		convArg0 = o.QTextFragment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragmentltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:317
// index:0
// Public Visibility=Default Availability=Available
// [4] int position()
func (this *QTextFragment) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:318
// index:0
// Public Visibility=Default Availability=Available
// [4] int length()
func (this *QTextFragment) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:319
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(int)
func (this *QTextFragment) Contains(position int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment8containsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:321
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat charFormat()
func (this *QTextFragment) CharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment10charFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:322
// index:0
// Public Visibility=Default Availability=Available
// [4] int charFormatIndex()
func (this *QTextFragment) CharFormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment15charFormatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:323
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QTextFragment) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

func DeleteQTextFragment(this *QTextFragment) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextFragmentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
