package qtcore

// /usr/include/qt/QtCore/qmimedata.h
// #include <qmimedata.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/cffiqt"
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
type QMimeData struct {
	*QObject
}

func (this *QMimeData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMimeData) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQMimeDataFromPointer(cthis unsafe.Pointer) *QMimeData {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QMimeData{bcthis0}
}
func (*QMimeData) NewFromPointer(cthis unsafe.Pointer) *QMimeData {
	return NewQMimeDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmimedata.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMimeData) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:56
// index:0
// Public
// void QMimeData()
func NewQMimeData() *QMimeData {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeDataC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeDataFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmimedata.h:57
// index:0
// Public virtual
// void ~QMimeData()
func DeleteQMimeData(*QMimeData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:61
// index:0
// Public
// bool hasUrls()
func (this *QMimeData) HasUrls() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasUrlsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:63
// index:0
// Public
// QString text()
func (this *QMimeData) Text() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4textEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:64
// index:0
// Public
// void setText(const QString &)
func (this *QMimeData) SetText(text *QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:65
// index:0
// Public
// bool hasText()
func (this *QMimeData) HasText() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:67
// index:0
// Public
// QString html()
func (this *QMimeData) Html() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4htmlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:68
// index:0
// Public
// void setHtml(const QString &)
func (this *QMimeData) SetHtml(html *QString) {
	var convArg0 = html.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:69
// index:0
// Public
// bool hasHtml()
func (this *QMimeData) HasHtml() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasHtmlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:71
// index:0
// Public
// QVariant imageData()
func (this *QMimeData) ImageData() *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9imageDataEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:72
// index:0
// Public
// void setImageData(const QVariant &)
func (this *QMimeData) SetImageData(image *QVariant) {
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12setImageDataERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:73
// index:0
// Public
// bool hasImage()
func (this *QMimeData) HasImage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData8hasImageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:75
// index:0
// Public
// QVariant colorData()
func (this *QMimeData) ColorData() *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9colorDataEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:76
// index:0
// Public
// void setColorData(const QVariant &)
func (this *QMimeData) SetColorData(color *QVariant) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12setColorDataERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:77
// index:0
// Public
// bool hasColor()
func (this *QMimeData) HasColor() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData8hasColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:79
// index:0
// Public
// QByteArray data(const QString &)
func (this *QMimeData) Data(mimetype *QString) *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4dataERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:80
// index:0
// Public
// void setData(const QString &, const QByteArray &)
func (this *QMimeData) SetData(mimetype *QString, data *QByteArray) {
	var convArg0 = mimetype.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setDataERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:81
// index:0
// Public
// void removeFormat(const QString &)
func (this *QMimeData) RemoveFormat(mimetype *QString) {
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12removeFormatERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:83
// index:0
// Public virtual
// bool hasFormat(const QString &)
func (this *QMimeData) HasFormat(mimetype *QString) bool {
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9hasFormatERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:86
// index:0
// Public
// void clear()
func (this *QMimeData) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:88
// index:0
// Protected virtual
// QVariant retrieveData(const QString &, QVariant::Type)
func (this *QMimeData) RetrieveData(mimetype *QString, preferredType int) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, preferredType)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
