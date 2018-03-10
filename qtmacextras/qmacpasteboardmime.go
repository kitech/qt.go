package qtmacextras

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h
// #include <qmacpasteboardmime.h>
// #include <Qtmacextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QMacPasteboardMime struct {
	*qtrt.CObject
}
type QMacPasteboardMime_ITF interface {
	QMacPasteboardMime_PTR() *QMacPasteboardMime
}

func (ptr *QMacPasteboardMime) QMacPasteboardMime_PTR() *QMacPasteboardMime { return ptr }

func (this *QMacPasteboardMime) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMacPasteboardMime) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMacPasteboardMimeFromPointer(cthis unsafe.Pointer) *QMacPasteboardMime {
	return &QMacPasteboardMime{&qtrt.CObject{cthis}}
}
func (*QMacPasteboardMime) NewFromPointer(cthis unsafe.Pointer) *QMacPasteboardMime {
	return NewQMacPasteboardMimeFromPointer(cthis)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMacPasteboardMime(char)

/*
Constructs a new conversion object of type t, adding it to the globally accessed list of available converters.
*/
func NewQMacPasteboardMime(arg0 byte) *QMacPasteboardMime {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMimeC2Ec", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMacPasteboardMimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMacPasteboardMime)
	return gothis
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMacPasteboardMime()

/*

 */
func DeleteQMacPasteboardMime(this *QMacPasteboardMime) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMimeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString convertorName()

/*
Returns a name for the converter.

All subclasses must reimplement this pure virtual function.
*/
func (this *QMacPasteboardMime) ConvertorName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime13convertorNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool canConvert(const QString &, QString)

/*
Returns true if the converter can convert (both ways) between mime and flav; otherwise returns false.

All subclasses must reimplement this pure virtual function.
*/
func (this *QMacPasteboardMime) CanConvert(mime string, flav string) bool {
	var tmpArg0 = qtcore.NewQString_5(mime)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(flav)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime10canConvertERK7QStringS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString mimeFor(QString)

/*
Returns the MIME UTI used for Mac flavor flav, or 0 if this converter does not support flav.

All subclasses must reimplement this pure virtual function.
*/
func (this *QMacPasteboardMime) MimeFor(flav string) string {
	var tmpArg0 = qtcore.NewQString_5(flav)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime7mimeForE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString flavorFor(const QString &)

/*
Returns the Mac UTI used for MIME type mime, or 0 if this converter does not support mime.

All subclasses must reimplement this pure virtual function.
*/
func (this *QMacPasteboardMime) FlavorFor(mime string) string {
	var tmpArg0 = qtcore.NewQString_5(mime)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime9flavorForERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count(QMimeData *)

/*
Returns the item count for the given mimeData
*/
func (this *QMacPasteboardMime) Count(mimeData qtcore.QMimeData_ITF /*777 QMimeData **/) int {
	var convArg0 unsafe.Pointer
	if mimeData != nil && mimeData.QMimeData_PTR() != nil {
		convArg0 = mimeData.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime5countEP9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

/*


 */
type QMacPasteboardMime__QMacPasteboardMimeType = int

//
const QMacPasteboardMime__MIME_DND QMacPasteboardMime__QMacPasteboardMimeType = 1

//
const QMacPasteboardMime__MIME_CLIP QMacPasteboardMime__QMacPasteboardMimeType = 2

//
const QMacPasteboardMime__MIME_QT_CONVERTOR QMacPasteboardMime__QMacPasteboardMimeType = 4

//
const QMacPasteboardMime__MIME_QT3_CONVERTOR QMacPasteboardMime__QMacPasteboardMimeType = 8

//
const QMacPasteboardMime__MIME_ALL QMacPasteboardMime__QMacPasteboardMimeType = 3

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
}

//  keep block end
