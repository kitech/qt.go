//  header block begin
// /usr/include/qt/QtCore/qmimedata.h
// #include <qmimedata.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
type QMimeData struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qmimedata.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMimeData) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:56
// index:0
// void QMimeData()
func NewQMimeData() *QMimeData {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeDataC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMimeData{cthis}
}

// /usr/include/qt/QtCore/qmimedata.h:57
// index:0
// virtual
// void ~QMimeData()
func DeleteQMimeData(*QMimeData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:59
// index:0
// QList<QUrl> urls()
func (this *QMimeData) Urls() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4urlsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:61
// index:0
// bool hasUrls()
func (this *QMimeData) HasUrls() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasUrlsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:63
// index:0
// QString text()
func (this *QMimeData) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:64
// index:0
// void setText(const class QString &)
func (this *QMimeData) SetText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:65
// index:0
// bool hasText()
func (this *QMimeData) HasText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:67
// index:0
// QString html()
func (this *QMimeData) Html() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4htmlEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:68
// index:0
// void setHtml(const class QString &)
func (this *QMimeData) SetHtml(html unsafe.Pointer) {
	// 0: (, const QString & html), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, html)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:69
// index:0
// bool hasHtml()
func (this *QMimeData) HasHtml() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasHtmlEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:71
// index:0
// QVariant imageData()
func (this *QMimeData) ImageData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9imageDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:72
// index:0
// void setImageData(const class QVariant &)
func (this *QMimeData) SetImageData(image unsafe.Pointer) {
	// 0: (, const QVariant & image), (image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12setImageDataERK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:73
// index:0
// bool hasImage()
func (this *QMimeData) HasImage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData8hasImageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:75
// index:0
// QVariant colorData()
func (this *QMimeData) ColorData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9colorDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:76
// index:0
// void setColorData(const class QVariant &)
func (this *QMimeData) SetColorData(color unsafe.Pointer) {
	// 0: (, const QVariant & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12setColorDataERK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:77
// index:0
// bool hasColor()
func (this *QMimeData) HasColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData8hasColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:79
// index:0
// QByteArray data(const class QString &)
func (this *QMimeData) Data(mimetype unsafe.Pointer) {
	// 0: (, const QString & mimetype), (mimetype)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4dataERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, mimetype)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:80
// index:0
// void setData(const class QString &, const class QByteArray &)
func (this *QMimeData) SetData(mimetype unsafe.Pointer, data unsafe.Pointer) {
	// 0: (, const QString & mimetype, const QByteArray & data), (mimetype, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setDataERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, mimetype, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:81
// index:0
// void removeFormat(const class QString &)
func (this *QMimeData) RemoveFormat(mimetype unsafe.Pointer) {
	// 0: (, const QString & mimetype), (mimetype)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12removeFormatERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, mimetype)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:83
// index:0
// virtual
// bool hasFormat(const class QString &)
func (this *QMimeData) HasFormat(mimetype unsafe.Pointer) {
	// 0: (, const QString & mimetype), (mimetype)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9hasFormatERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, mimetype)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:84
// index:0
// virtual
// QStringList formats()
func (this *QMimeData) Formats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7formatsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:86
// index:0
// void clear()
func (this *QMimeData) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
