package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QTextFragment struct {
	*qtrt.CObject
}

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
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:311
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextFragment) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:317
// index:0
// Public Visibility=Default Availability=Available
// [4] int position()
func (this *QTextFragment) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:318
// index:0
// Public Visibility=Default Availability=Available
// [4] int length()
func (this *QTextFragment) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:319
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(int)
func (this *QTextFragment) Contains(position int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment8containsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:321
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat charFormat()
func (this *QTextFragment) CharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment10charFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
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
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:323
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QTextFragment) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextFragment4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

func DeleteQTextFragment(this *QTextFragment) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextFragmentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
