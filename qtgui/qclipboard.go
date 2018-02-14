package qtgui

// /usr/include/qt/QtGui/qclipboard.h
// #include <qclipboard.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QClipboard struct {
	*qtcore.QObject
}
type QClipboard_ITF interface {
	qtcore.QObject_ITF
	QClipboard_PTR() *QClipboard
}

func (ptr *QClipboard) QClipboard_PTR() *QClipboard { return ptr }

func (this *QClipboard) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QClipboard) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQClipboardFromPointer(cthis unsafe.Pointer) *QClipboard {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QClipboard{bcthis0}
}
func (*QClipboard) NewFromPointer(cthis unsafe.Pointer) *QClipboard {
	return NewQClipboardFromPointer(cthis)
}

// /usr/include/qt/QtGui/qclipboard.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QClipboard) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qclipboard.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear(enum QClipboard::Mode)
func (this *QClipboard) Clear(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard5clearENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsSelection()
func (this *QClipboard) SupportsSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard17supportsSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsFindBuffer()
func (this *QClipboard) SupportsFindBuffer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard18supportsFindBufferEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsSelection()
func (this *QClipboard) OwnsSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard13ownsSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsClipboard()
func (this *QClipboard) OwnsClipboard() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard13ownsClipboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsFindBuffer()
func (this *QClipboard) OwnsFindBuffer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard14ownsFindBufferEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(enum QClipboard::Mode)
func (this *QClipboard) Text(mode int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard4textENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qclipboard.h:75
// index:1
// Public Visibility=Default Availability=Available
// [8] QString text(QString &, enum QClipboard::Mode)
func (this *QClipboard) Text_1(subtype string, mode int) string {
	var tmpArg0 = qtcore.NewQString_5(subtype)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard4textER7QStringNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qclipboard.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &, enum QClipboard::Mode)
func (this *QClipboard) SetText(arg0 string, mode int) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard7setTextERK7QStringNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMimeData * mimeData(enum QClipboard::Mode)
func (this *QClipboard) MimeData(mode int) *qtcore.QMimeData /*777 const QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard8mimeDataENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qclipboard.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMimeData(QMimeData *, enum QClipboard::Mode)
func (this *QClipboard) SetMimeData(data qtcore.QMimeData_ITF /*777 QMimeData **/, mode int) {
	var convArg0 = data.QMimeData_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard11setMimeDataEP9QMimeDataNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage image(enum QClipboard::Mode)
func (this *QClipboard) Image(mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard5imageENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:82
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(enum QClipboard::Mode)
func (this *QClipboard) Pixmap(mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard6pixmapENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImage(const QImage &, enum QClipboard::Mode)
func (this *QClipboard) SetImage(arg0 QImage_ITF, mode int) {
	var convArg0 = arg0.QImage_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard8setImageERK6QImageNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &, enum QClipboard::Mode)
func (this *QClipboard) SetPixmap(arg0 QPixmap_ITF, mode int) {
	var convArg0 = arg0.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard9setPixmapERK7QPixmapNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed(QClipboard::Mode)
func (this *QClipboard) Changed(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard7changedENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()
func (this *QClipboard) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void findBufferChanged()
func (this *QClipboard) FindBufferChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard17findBufferChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dataChanged()
func (this *QClipboard) DataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard11dataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQClipboard(this *QClipboard) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboardD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QClipboard__Mode = int

const QClipboard__Clipboard QClipboard__Mode = 0
const QClipboard__Selection QClipboard__Mode = 1
const QClipboard__FindBuffer QClipboard__Mode = 2
const QClipboard__LastMode QClipboard__Mode = 2

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
