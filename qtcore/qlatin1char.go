package qtcore

// /usr/include/qt/QtCore/qchar.h
// #include <qchar.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QLatin1Char struct {
	*qtrt.CObject
}

func (this *QLatin1Char) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLatin1Char) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQLatin1CharFromPointer(cthis unsafe.Pointer) *QLatin1Char {
	return &QLatin1Char{&qtrt.CObject{cthis}}
}
func (*QLatin1Char) NewFromPointer(cthis unsafe.Pointer) *QLatin1Char {
	return NewQLatin1CharFromPointer(cthis)
}

// /usr/include/qt/QtCore/qchar.h:53
// index:0
// Public inline
// void QLatin1Char(char)
func NewQLatin1Char(c byte) *QLatin1Char {
	cthis := qtrt.Calloc(1, 256) // 1
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QLatin1CharC2Ec", ffiqt.FFI_TYPE_VOID, cthis, c)
	gopp.ErrPrint(err, rv)
	gothis := NewQLatin1CharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:54
// index:0
// Public inline
// char toLatin1()
func (this *QLatin1Char) ToLatin1() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLatin1Char8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 111
}

// /usr/include/qt/QtCore/qchar.h:55
// index:0
// Public inline
// ushort unicode()
func (this *QLatin1Char) Unicode() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QLatin1Char7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

//  body block end
