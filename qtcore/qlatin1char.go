package qtcore

// /usr/include/qt/QtCore/qchar.h
// #include <qchar.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QLatin1Char struct {
	*qtrt.CObject
}
type QLatin1Char_ITF interface {
	QLatin1Char_PTR() *QLatin1Char
}

func (ptr *QLatin1Char) QLatin1Char_PTR() *QLatin1Char { return ptr }

func (this *QLatin1Char) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLatin1Char) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLatin1CharFromPointer(cthis unsafe.Pointer) *QLatin1Char {
	return &QLatin1Char{&qtrt.CObject{cthis}}
}
func (*QLatin1Char) NewFromPointer(cthis unsafe.Pointer) *QLatin1Char {
	return NewQLatin1CharFromPointer(cthis)
}

// /usr/include/qt/QtCore/qchar.h:53
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1Char(char)
func NewQLatin1Char(c byte) *QLatin1Char {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLatin1CharC2Ec", qtrt.FFI_TYPE_POINTER, c)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1CharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1Char)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:54
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char toLatin1() const
func (this *QLatin1Char) ToLatin1() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLatin1Char8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qchar.h:55
// index:0
// Public inline Visibility=Default Availability=Available
// [2] ushort unicode() const
func (this *QLatin1Char) Unicode() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QLatin1Char7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

func DeleteQLatin1Char(this *QLatin1Char) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QLatin1CharD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
