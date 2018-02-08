package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinmime.h
// #include <qwinmime.h>
// #include <Qtwinextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QWinMime struct {
	*qtrt.CObject
}

func (this *QWinMime) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWinMime) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWinMimeFromPointer(cthis unsafe.Pointer) *QWinMime {
	return &QWinMime{&qtrt.CObject{cthis}}
}
func (*QWinMime) NewFromPointer(cthis unsafe.Pointer) *QWinMime {
	return NewQWinMimeFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinmime.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinMime()
func NewQWinMime() *QWinMime {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QWinMimeC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinMimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWinMime)
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinmime.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinMime()
func DeleteQWinMime(this *QWinMime) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QWinMimeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinmime.h:71
// index:0
// Public static Visibility=Default Availability=Available
// [4] int registerMimeType(const QString &)
func (this *QWinMime) RegisterMimeType(mime string) int {
	var tmpArg0 = qtcore.NewQString_5(mime)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QWinMime16registerMimeTypeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QWinMime_RegisterMimeType(mime string) int {
	var nilthis *QWinMime
	rv := nilthis.RegisterMimeType(mime)
	return rv
}

//  body block end
