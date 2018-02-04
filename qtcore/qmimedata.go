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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// QVariant retrieveData(const class QString &, class QVariant::Type)
func (this *QMimeData) InheritRetrieveData(f func(mimetype *QString, preferredType int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "retrieveData", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QMimeData) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMimeData()
func NewQMimeData() *QMimeData {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeDataC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qmimedata.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMimeData()
func DeleteQMimeData(this *QMimeData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmimedata.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasUrls()
func (this *QMimeData) HasUrls() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData7hasUrlsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QMimeData) Text() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QMimeData) SetText(text *QString) {
	var convArg0 = text.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasText()
func (this *QMimeData) HasText() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData7hasTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString html()
func (this *QMimeData) Html() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData4htmlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &)
func (this *QMimeData) SetHtml(html *QString) {
	var convArg0 = html.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData7setHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasHtml()
func (this *QMimeData) HasHtml() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData7hasHtmlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant imageData()
func (this *QMimeData) ImageData() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData9imageDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImageData(const QVariant &)
func (this *QMimeData) SetImageData(image *QVariant) {
	var convArg0 = image.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData12setImageDataERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasImage()
func (this *QMimeData) HasImage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData8hasImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:75
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant colorData()
func (this *QMimeData) ColorData() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData9colorDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorData(const QVariant &)
func (this *QMimeData) SetColorData(color *QVariant) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData12setColorDataERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasColor()
func (this *QMimeData) HasColor() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData8hasColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray data(const QString &)
func (this *QMimeData) Data(mimetype *QString) *QByteArray /*123*/ {
	var convArg0 = mimetype.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData4dataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QString &, const QByteArray &)
func (this *QMimeData) SetData(mimetype *QString, data *QByteArray) {
	var convArg0 = mimetype.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData7setDataERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeFormat(const QString &)
func (this *QMimeData) RemoveFormat(mimetype *QString) {
	var convArg0 = mimetype.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData12removeFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasFormat(const QString &)
func (this *QMimeData) HasFormat(mimetype *QString) bool {
	var convArg0 = mimetype.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData9hasFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QMimeData) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant retrieveData(const QString &, QVariant::Type)
func (this *QMimeData) RetrieveData(mimetype *QString, preferredType int) *QVariant /*123*/ {
	var convArg0 = mimetype.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, preferredType)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

//  body block end
