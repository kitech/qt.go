package qtcore

// /usr/include/qt/QtCore/qmimedata.h
// #include <qmimedata.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// QVariant retrieveData(const class QString &, class QVariant::Type)
func (this *QMimeData) InheritRetrieveData(f func(mimetype string, preferredType int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "retrieveData", f)
}

/*

 */
type QMimeData struct {
	*QObject
}
type QMimeData_ITF interface {
	QObject_ITF
	QMimeData_PTR() *QMimeData
}

func (ptr *QMimeData) QMimeData_PTR() *QMimeData { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMimeData) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qmimedata.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMimeData()

/*
Constructs a new MIME data object with no data in it.
*/
func NewQMimeData() *QMimeData {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeDataC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMimeData")
	return gothis
}

// /usr/include/qt/QtCore/qmimedata.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMimeData()

/*

 */
func DeleteQMimeData(this *QMimeData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmimedata.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasUrls() const

/*
Returns true if the object can return a list of urls; otherwise returns false.

URLs correspond to the MIME type text/uri-list.

See also setUrls(), urls(), and hasFormat().
*/
func (this *QMimeData) HasUrls() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData7hasUrlsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*
Returns a plain text (MIME type text/plain) representation of the data.

See also setText(), hasText(), html(), and data().
*/
func (this *QMimeData) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimedata.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*
Sets text as the plain text (MIME type text/plain) used to represent the data.

See also text(), hasText(), setHtml(), and setData().
*/
func (this *QMimeData) SetText(text string) {
	var tmpArg0 = NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasText() const

/*
Returns true if the object can return plain text (MIME type text/plain); otherwise returns false.

See also setText(), text(), hasHtml(), and hasFormat().
*/
func (this *QMimeData) HasText() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData7hasTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString html() const

/*
Returns a string if the data stored in the object is HTML (MIME type text/html); otherwise returns an empty string.

See also setHtml(), hasHtml(), and setData().
*/
func (this *QMimeData) Html() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData4htmlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimedata.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &)

/*
Sets html as the HTML (MIME type text/html) used to represent the data.

See also html(), hasHtml(), setText(), and setData().
*/
func (this *QMimeData) SetHtml(html string) {
	var tmpArg0 = NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData7setHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasHtml() const

/*
Returns true if the object can return HTML (MIME type text/html); otherwise returns false.

See also setHtml(), html(), and hasFormat().
*/
func (this *QMimeData) HasHtml() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData7hasHtmlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant imageData() const

/*
Returns a QVariant storing a QImage if the object can return an image; otherwise returns a null variant.

A QVariant is used because QMimeData belongs to the Qt Core module, whereas QImage belongs to Qt GUI. To convert the QVariant to a QImage, simply use qvariant_cast(). For example:


  if (event->mimeData()->hasImage()) {
      QImage image = qvariant_cast<QImage>(event->mimeData()->imageData());
      ...
  }



See also setImageData() and hasImage().
*/
func (this *QMimeData) ImageData() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData9imageDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImageData(const QVariant &)

/*
Sets the data in the object to the given image.

A QVariant is used because QMimeData belongs to the Qt Core module, whereas QImage belongs to Qt GUI. The conversion from QImage to QVariant is implicit. For example:


  mimeData->setImageData(QImage("beautifulfjord.png"));



See also imageData(), hasImage(), and setData().
*/
func (this *QMimeData) SetImageData(image QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if image != nil && image.QVariant_PTR() != nil {
		convArg0 = image.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData12setImageDataERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasImage() const

/*
Returns true if the object can return an image; otherwise returns false.

See also setImageData(), imageData(), and hasFormat().
*/
func (this *QMimeData) HasImage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData8hasImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:75
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant colorData() const

/*
Returns a color if the data stored in the object represents a color (MIME type application/x-color); otherwise returns a null variant.

A QVariant is used because QMimeData belongs to the Qt Core module, whereas QColor belongs to Qt GUI. To convert the QVariant to a QColor, simply use qvariant_cast(). For example:


  if (event->mimeData()->hasColor()) {
      QColor color = qvariant_cast<QColor>(event->mimeData()->colorData());
      ...
  }



See also hasColor(), setColorData(), and data().
*/
func (this *QMimeData) ColorData() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData9colorDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorData(const QVariant &)

/*
Sets the color data in the object to the given color.

Colors correspond to the MIME type application/x-color.

See also colorData(), hasColor(), and setData().
*/
func (this *QMimeData) SetColorData(color QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QVariant_PTR() != nil {
		convArg0 = color.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData12setColorDataERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasColor() const

/*
Returns true if the object can return a color (MIME type application/x-color); otherwise returns false.

See also setColorData(), colorData(), and hasFormat().
*/
func (this *QMimeData) HasColor() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData8hasColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray data(const QString &) const

/*
Returns the data stored in the object in the format described by the MIME type specified by mimeType.

See also setData().
*/
func (this *QMimeData) Data(mimetype string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(mimetype)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData4dataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QString &, const QByteArray &)

/*
Sets the data associated with the MIME type given by mimeType to the specified data.

For the most common types of data, you can call the higher-level functions setText(), setHtml(), setUrls(), setImageData(), and setColorData() instead.

Note that if you want to use a custom data type in an item view drag and drop operation, you must register it as a Qt meta type, using the Q_DECLARE_METATYPE() macro, and implement stream operators for it. The stream operators must then be registered with the qRegisterMetaTypeStreamOperators() function.

See also data(), hasFormat(), QMetaType, and qRegisterMetaTypeStreamOperators().
*/
func (this *QMimeData) SetData(mimetype string, data QByteArray_ITF) {
	var tmpArg0 = NewQString_5(mimetype)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg1 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData7setDataERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeFormat(const QString &)

/*
Removes the data entry for mimeType in the object.

This function was introduced in  Qt 4.4.
*/
func (this *QMimeData) RemoveFormat(mimetype string) {
	var tmpArg0 = NewQString_5(mimetype)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData12removeFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasFormat(const QString &) const

/*
Returns true if the object can return data for the MIME type specified by mimeType; otherwise returns false.

For the most common types of data, you can call the higher-level functions hasText(), hasHtml(), hasUrls(), hasImage(), and hasColor() instead.

See also formats(), setData(), and data().
*/
func (this *QMimeData) HasFormat(mimetype string) bool {
	var tmpArg0 = NewQString_5(mimetype)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData9hasFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimedata.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList formats() const

/*
Returns a list of formats supported by the object. This is a list of MIME types for which the object can return suitable data. The formats in the list are in a priority order.

For the most common types of data, you can call the higher-level functions hasText(), hasHtml(), hasUrls(), hasImage(), and hasColor() instead.

See also hasFormat(), setData(), and data().
*/
func (this *QMimeData) Formats() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData7formatsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qmimedata.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all the MIME type and data entries in the object.
*/
func (this *QMimeData) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeData5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant retrieveData(const QString &, QVariant::Type) const

/*
Returns a variant with the given type containing data for the MIME type specified by mimeType. If the object does not support the MIME type or variant type given, a null variant is returned instead.

This function is called by the general data() getter and by the convenience getters (text(), html(), urls(), imageData(), and colorData()). You can reimplement it if you want to store your data using a custom data structure (instead of a QByteArray, which is what setData() provides). You would then also need to reimplement hasFormat() and formats().

See also data().
*/
func (this *QMimeData) RetrieveData(mimetype string, preferredType int) *QVariant /*123*/ {
	var tmpArg0 = NewQString_5(mimetype)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, preferredType)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
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
}

//  keep block end
