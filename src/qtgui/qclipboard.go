//  header block begin
// /usr/include/qt/QtGui/qclipboard.h
// #include <qclipboard.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qclipboard.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QClipboard) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:65
// index:0
// void clear(enum QClipboard::Mode)
func (this *QClipboard) Clear(mode int) {
	// 0: (, mode QClipboard::Mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard5clearENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:67
// index:0
// bool supportsSelection()
func (this *QClipboard) SupportsSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard17supportsSelectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:68
// index:0
// bool supportsFindBuffer()
func (this *QClipboard) SupportsFindBuffer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard18supportsFindBufferEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:70
// index:0
// bool ownsSelection()
func (this *QClipboard) OwnsSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard13ownsSelectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:71
// index:0
// bool ownsClipboard()
func (this *QClipboard) OwnsClipboard() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard13ownsClipboardEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:72
// index:0
// bool ownsFindBuffer()
func (this *QClipboard) OwnsFindBuffer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard14ownsFindBufferEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:74
// index:0
// QString text(enum QClipboard::Mode)
func (this *QClipboard) Text(mode int) {
	// 0: (, mode QClipboard::Mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard4textENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:75
// index:1
// QString text(class QString &, enum QClipboard::Mode)
func (this *QClipboard) Text_1(subtype unsafe.Pointer, mode int) {
	// 1: (, subtype QString &, mode QClipboard::Mode), (subtype, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard4textER7QStringNS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), subtype, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:76
// index:0
// void setText(const class QString &, enum QClipboard::Mode)
func (this *QClipboard) SetText(arg0 unsafe.Pointer, mode int) {
	// 0: (, const QString & arg0, mode QClipboard::Mode), (arg0, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard7setTextERK7QStringNS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:78
// index:0
// const QMimeData * mimeData(enum QClipboard::Mode)
func (this *QClipboard) MimeData(mode int) {
	// 0: (, mode QClipboard::Mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard8mimeDataENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:79
// index:0
// void setMimeData(class QMimeData *, enum QClipboard::Mode)
func (this *QClipboard) SetMimeData(data unsafe.Pointer, mode int) {
	// 0: (, data QMimeData *, mode QClipboard::Mode), (data, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard11setMimeDataEP9QMimeDataNS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:81
// index:0
// QImage image(enum QClipboard::Mode)
func (this *QClipboard) Image(mode int) {
	// 0: (, mode QClipboard::Mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard5imageENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:82
// index:0
// QPixmap pixmap(enum QClipboard::Mode)
func (this *QClipboard) Pixmap(mode int) {
	// 0: (, mode QClipboard::Mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QClipboard6pixmapENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:83
// index:0
// void setImage(const class QImage &, enum QClipboard::Mode)
func (this *QClipboard) SetImage(arg0 unsafe.Pointer, mode int) {
	// 0: (, const QImage & arg0, mode QClipboard::Mode), (arg0, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard8setImageERK6QImageNS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:84
// index:0
// void setPixmap(const class QPixmap &, enum QClipboard::Mode)
func (this *QClipboard) SetPixmap(arg0 unsafe.Pointer, mode int) {
	// 0: (, const QPixmap & arg0, mode QClipboard::Mode), (arg0, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard9setPixmapERK7QPixmapNS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:87
// index:0
// void changed(class QClipboard::Mode)
func (this *QClipboard) Changed(mode int) {
	// 0: (, mode QClipboard::Mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard7changedENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:88
// index:0
// void selectionChanged()
func (this *QClipboard) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:89
// index:0
// void findBufferChanged()
func (this *QClipboard) FindBufferChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard17findBufferChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:90
// index:0
// void dataChanged()
func (this *QClipboard) DataChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QClipboard11dataChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
