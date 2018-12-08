package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinmime.h
// #include <qwinmime.h>
// #include <Qtwinextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"

//  ext block end

//  body block begin

/*

 */
type QWinMime struct {
	*qtrt.CObject
}
type QWinMime_ITF interface {
	QWinMime_PTR() *QWinMime
}

func (ptr *QWinMime) QWinMime_PTR() *QWinMime { return ptr }

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

/*
Constructs a new conversion object, adding it to the globally accessed list of available converters.
*/
func (*QWinMime) NewForInherit() *QWinMime {
	return NewQWinMime()
}
func NewQWinMime() *QWinMime {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QWinMimeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinMimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWinMime)
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinmime.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinMime()

/*

 */
func DeleteQWinMime(this *QWinMime) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QWinMimeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinmime.h:71
// index:0
// Public static Visibility=Default Availability=Available
// [4] int registerMimeType(const QString &)

/*
Registers the MIME type mime, and returns an ID number identifying the format on Windows.
*/
func (this *QWinMime) RegisterMimeType(mime string) int {
	var tmpArg0 = qtcore.NewQString5(mime)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QWinMime16registerMimeTypeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QWinMime_RegisterMimeType(mime string) int {
	var nilthis *QWinMime
	rv := nilthis.RegisterMimeType(mime)
	return rv
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtwidgets.KeepMe()
	}
}

//  keep block end
