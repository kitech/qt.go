package qtmacextras

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h
// #include <qmacpasteboardmime.h>
// #include <Qtmacextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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

type QMacPasteboardMime struct {
	*qtrt.CObject
}

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
func NewQMacPasteboardMime(arg0 byte) *QMacPasteboardMime {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMimeC2Ec", qtrt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMacPasteboardMimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMacPasteboardMime)
	return gothis
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMacPasteboardMime()
func DeleteQMacPasteboardMime(this *QMacPasteboardMime) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMimeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString convertorName()
func (this *QMacPasteboardMime) ConvertorName() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime13convertorNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool canConvert(const QString &, QString)
func (this *QMacPasteboardMime) CanConvert(mime *qtcore.QString, flav *qtcore.QString /*123*/) bool {
	var convArg0 = mime.GetCthis()
	var convArg1 = flav.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime10canConvertERK7QStringS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString mimeFor(QString)
func (this *QMacPasteboardMime) MimeFor(flav *qtcore.QString /*123*/) *qtcore.QString /*123*/ {
	var convArg0 = flav.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime7mimeForE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString flavorFor(const QString &)
func (this *QMacPasteboardMime) FlavorFor(mime *qtcore.QString) *qtcore.QString /*123*/ {
	var convArg0 = mime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime9flavorForERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count(QMimeData *)
func (this *QMacPasteboardMime) Count(mimeData *qtcore.QMimeData /*777 QMimeData **/) int {
	var convArg0 = mimeData.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMacPasteboardMime5countEP9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QMacPasteboardMime__QMacPasteboardMimeType = int

const QMacPasteboardMime__MIME_DND QMacPasteboardMime__QMacPasteboardMimeType = 1
const QMacPasteboardMime__MIME_CLIP QMacPasteboardMime__QMacPasteboardMimeType = 2
const QMacPasteboardMime__MIME_QT_CONVERTOR QMacPasteboardMime__QMacPasteboardMimeType = 4
const QMacPasteboardMime__MIME_QT3_CONVERTOR QMacPasteboardMime__QMacPasteboardMimeType = 8
const QMacPasteboardMime__MIME_ALL QMacPasteboardMime__QMacPasteboardMimeType = 3

//  body block end
