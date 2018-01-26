package qtgui

// /usr/include/qt/QtGui/qclipboard.h
// #include <qclipboard.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
		ffiqt.KeepMe()
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
type QClipboard struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QClipboard) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:65
// index:0
// Public
// void clear(enum QClipboard::Mode)
func (this *QClipboard) Clear(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard5clearENS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:67
// index:0
// Public
// bool supportsSelection()
func (this *QClipboard) SupportsSelection() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard17supportsSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:68
// index:0
// Public
// bool supportsFindBuffer()
func (this *QClipboard) SupportsFindBuffer() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard18supportsFindBufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:70
// index:0
// Public
// bool ownsSelection()
func (this *QClipboard) OwnsSelection() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard13ownsSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:71
// index:0
// Public
// bool ownsClipboard()
func (this *QClipboard) OwnsClipboard() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard13ownsClipboardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:72
// index:0
// Public
// bool ownsFindBuffer()
func (this *QClipboard) OwnsFindBuffer() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard14ownsFindBufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:74
// index:0
// Public
// QString text(enum QClipboard::Mode)
func (this *QClipboard) Text(mode int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard4textENS_4ModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:75
// index:1
// Public
// QString text(class QString &, enum QClipboard::Mode)
func (this *QClipboard) Text_1(subtype *qtcore.QString, mode int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = subtype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard4textER7QStringNS_4ModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:76
// index:0
// Public
// void setText(const class QString &, enum QClipboard::Mode)
func (this *QClipboard) SetText(arg0 *qtcore.QString, mode int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard7setTextERK7QStringNS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:78
// index:0
// Public
// const QMimeData * mimeData(enum QClipboard::Mode)
func (this *QClipboard) MimeData(mode int) *qtcore.QMimeData /*777 const QMimeData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard8mimeDataENS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:79
// index:0
// Public
// void setMimeData(class QMimeData *, enum QClipboard::Mode)
func (this *QClipboard) SetMimeData(data *qtcore.QMimeData /*777 QMimeData **/, mode int) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard11setMimeDataEP9QMimeDataNS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:81
// index:0
// Public
// QImage image(enum QClipboard::Mode)
func (this *QClipboard) Image(mode int) *QImage /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard5imageENS_4ModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:82
// index:0
// Public
// QPixmap pixmap(enum QClipboard::Mode)
func (this *QClipboard) Pixmap(mode int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard6pixmapENS_4ModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:83
// index:0
// Public
// void setImage(const class QImage &, enum QClipboard::Mode)
func (this *QClipboard) SetImage(arg0 *QImage, mode int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard8setImageERK6QImageNS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:84
// index:0
// Public
// void setPixmap(const class QPixmap &, enum QClipboard::Mode)
func (this *QClipboard) SetPixmap(arg0 *QPixmap, mode int) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard9setPixmapERK7QPixmapNS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:87
// index:0
// Public
// void changed(class QClipboard::Mode)
func (this *QClipboard) Changed(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard7changedENS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:88
// index:0
// Public
// void selectionChanged()
func (this *QClipboard) SelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard16selectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:89
// index:0
// Public
// void findBufferChanged()
func (this *QClipboard) FindBufferChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard17findBufferChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:90
// index:0
// Public
// void dataChanged()
func (this *QClipboard) DataChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard11dataChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QClipboard__Mode = int

const QClipboard__Clipboard QClipboard__Mode = 0
const QClipboard__Selection QClipboard__Mode = 1
const QClipboard__FindBuffer QClipboard__Mode = 2
const QClipboard__LastMode QClipboard__Mode = 2

//  body block end
