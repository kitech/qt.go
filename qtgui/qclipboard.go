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
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QClipboard) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qclipboard.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear(QClipboard::Mode)

/*
Clear the clipboard contents.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, this function clears the global clipboard contents. If mode is QClipboard::Selection, this function clears the global mouse selection contents. If mode is QClipboard::FindBuffer, this function clears the search string buffer.

See also QClipboard::Mode and supportsSelection().
*/
func (this *QClipboard) Clear(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard5clearENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear(QClipboard::Mode)

/*
Clear the clipboard contents.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, this function clears the global clipboard contents. If mode is QClipboard::Selection, this function clears the global mouse selection contents. If mode is QClipboard::FindBuffer, this function clears the search string buffer.

See also QClipboard::Mode and supportsSelection().
*/
func (this *QClipboard) Clear__() {
	// arg: 0, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard5clearENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsSelection() const

/*
Returns true if the clipboard supports mouse selection; otherwise returns false.
*/
func (this *QClipboard) SupportsSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard17supportsSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsFindBuffer() const

/*
Returns true if the clipboard supports a separate search buffer; otherwise returns false.
*/
func (this *QClipboard) SupportsFindBuffer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard18supportsFindBufferEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsSelection() const

/*
Returns true if this clipboard object owns the mouse selection data; otherwise returns false.
*/
func (this *QClipboard) OwnsSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard13ownsSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsClipboard() const

/*
Returns true if this clipboard object owns the clipboard data; otherwise returns false.
*/
func (this *QClipboard) OwnsClipboard() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard13ownsClipboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsFindBuffer() const

/*
Returns true if this clipboard object owns the find buffer data; otherwise returns false.

This function was introduced in  Qt 4.2.
*/
func (this *QClipboard) OwnsFindBuffer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard14ownsFindBufferEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qclipboard.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(QClipboard::Mode) const

/*
Returns the clipboard text as plain text, or an empty string if the clipboard does not contain any text.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the text is retrieved from the global clipboard. If mode is QClipboard::Selection, the text is retrieved from the global mouse selection. If mode is QClipboard::FindBuffer, the text is retrieved from the search string buffer.

See also setText() and mimeData().
*/
func (this *QClipboard) Text(mode int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard4textENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qclipboard.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(QClipboard::Mode) const

/*
Returns the clipboard text as plain text, or an empty string if the clipboard does not contain any text.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the text is retrieved from the global clipboard. If mode is QClipboard::Selection, the text is retrieved from the global mouse selection. If mode is QClipboard::FindBuffer, the text is retrieved from the search string buffer.

See also setText() and mimeData().
*/
func (this *QClipboard) Text__() string {
	// arg: 0, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
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
// [8] QString text(QString &, QClipboard::Mode) const

/*
Returns the clipboard text as plain text, or an empty string if the clipboard does not contain any text.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the text is retrieved from the global clipboard. If mode is QClipboard::Selection, the text is retrieved from the global mouse selection. If mode is QClipboard::FindBuffer, the text is retrieved from the search string buffer.

See also setText() and mimeData().
*/
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

// /usr/include/qt/QtGui/qclipboard.h:75
// index:1
// Public Visibility=Default Availability=Available
// [8] QString text(QString &, QClipboard::Mode) const

/*
Returns the clipboard text as plain text, or an empty string if the clipboard does not contain any text.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the text is retrieved from the global clipboard. If mode is QClipboard::Selection, the text is retrieved from the global mouse selection. If mode is QClipboard::FindBuffer, the text is retrieved from the search string buffer.

See also setText() and mimeData().
*/
func (this *QClipboard) Text_1_(subtype string) string {
	var tmpArg0 = qtcore.NewQString_5(subtype)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
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
// [-2] void setText(const QString &, QClipboard::Mode)

/*
Copies text into the clipboard as plain text.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the text is stored in the global clipboard. If mode is QClipboard::Selection, the text is stored in the global mouse selection. If mode is QClipboard::FindBuffer, the text is stored in the search string buffer.

See also text() and setMimeData().
*/
func (this *QClipboard) SetText(arg0 string, mode int) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard7setTextERK7QStringNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &, QClipboard::Mode)

/*
Copies text into the clipboard as plain text.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the text is stored in the global clipboard. If mode is QClipboard::Selection, the text is stored in the global mouse selection. If mode is QClipboard::FindBuffer, the text is stored in the search string buffer.

See also text() and setMimeData().
*/
func (this *QClipboard) SetText__(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard7setTextERK7QStringNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMimeData * mimeData(QClipboard::Mode) const

/*
Returns a pointer to a QMimeData representation of the current clipboard data (can be NULL if the given mode is not supported by the platform).

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the data is retrieved from the global clipboard. If mode is QClipboard::Selection, the data is retrieved from the global mouse selection. If mode is QClipboard::FindBuffer, the data is retrieved from the search string buffer.

The text(), image(), and pixmap() functions are simpler wrappers for retrieving text, image, and pixmap data.

Note: The pointer returned might become invalidated when the contents of the clipboard changes; either by calling one of the setter functions or externally by the system clipboard changing.

See also setMimeData().
*/
func (this *QClipboard) MimeData(mode int) *qtcore.QMimeData /*777 const QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard8mimeDataENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qclipboard.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMimeData * mimeData(QClipboard::Mode) const

/*
Returns a pointer to a QMimeData representation of the current clipboard data (can be NULL if the given mode is not supported by the platform).

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the data is retrieved from the global clipboard. If mode is QClipboard::Selection, the data is retrieved from the global mouse selection. If mode is QClipboard::FindBuffer, the data is retrieved from the search string buffer.

The text(), image(), and pixmap() functions are simpler wrappers for retrieving text, image, and pixmap data.

Note: The pointer returned might become invalidated when the contents of the clipboard changes; either by calling one of the setter functions or externally by the system clipboard changing.

See also setMimeData().
*/
func (this *QClipboard) MimeData__() *qtcore.QMimeData /*777 const QMimeData **/ {
	// arg: 0, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard8mimeDataENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qclipboard.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMimeData(QMimeData *, QClipboard::Mode)

/*
Sets the clipboard data to src. Ownership of the data is transferred to the clipboard. If you want to remove the data either call clear() or call setMimeData() again with new data.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the data is stored in the global clipboard. If mode is QClipboard::Selection, the data is stored in the global mouse selection. If mode is QClipboard::FindBuffer, the data is stored in the search string buffer.

The setText(), setImage() and setPixmap() functions are simpler wrappers for setting text, image and pixmap data respectively.

See also mimeData().
*/
func (this *QClipboard) SetMimeData(data qtcore.QMimeData_ITF /*777 QMimeData **/, mode int) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard11setMimeDataEP9QMimeDataNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMimeData(QMimeData *, QClipboard::Mode)

/*
Sets the clipboard data to src. Ownership of the data is transferred to the clipboard. If you want to remove the data either call clear() or call setMimeData() again with new data.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the data is stored in the global clipboard. If mode is QClipboard::Selection, the data is stored in the global mouse selection. If mode is QClipboard::FindBuffer, the data is stored in the search string buffer.

The setText(), setImage() and setPixmap() functions are simpler wrappers for setting text, image and pixmap data respectively.

See also mimeData().
*/
func (this *QClipboard) SetMimeData__(data qtcore.QMimeData_ITF /*777 QMimeData **/) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	// arg: 1, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard11setMimeDataEP9QMimeDataNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage image(QClipboard::Mode) const

/*
Returns the clipboard image, or returns a null image if the clipboard does not contain an image or if it contains an image in an unsupported image format.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the image is retrieved from the global clipboard. If mode is QClipboard::Selection, the image is retrieved from the global mouse selection.

See also setImage(), pixmap(), mimeData(), and QImage::isNull().
*/
func (this *QClipboard) Image(mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard5imageENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage image(QClipboard::Mode) const

/*
Returns the clipboard image, or returns a null image if the clipboard does not contain an image or if it contains an image in an unsupported image format.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the image is retrieved from the global clipboard. If mode is QClipboard::Selection, the image is retrieved from the global mouse selection.

See also setImage(), pixmap(), mimeData(), and QImage::isNull().
*/
func (this *QClipboard) Image__() *QImage /*123*/ {
	// arg: 0, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard5imageENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:82
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QClipboard::Mode) const

/*
Returns the clipboard pixmap, or null if the clipboard does not contain a pixmap. Note that this can lose information. For example, if the image is 24-bit and the display is 8-bit, the result is converted to 8 bits, and if the image has an alpha channel, the result just has a mask.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the pixmap is retrieved from the global clipboard. If mode is QClipboard::Selection, the pixmap is retrieved from the global mouse selection.

See also setPixmap(), image(), mimeData(), and QPixmap::convertFromImage().
*/
func (this *QClipboard) Pixmap(mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard6pixmapENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:82
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QClipboard::Mode) const

/*
Returns the clipboard pixmap, or null if the clipboard does not contain a pixmap. Note that this can lose information. For example, if the image is 24-bit and the display is 8-bit, the result is converted to 8 bits, and if the image has an alpha channel, the result just has a mask.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the pixmap is retrieved from the global clipboard. If mode is QClipboard::Selection, the pixmap is retrieved from the global mouse selection.

See also setPixmap(), image(), mimeData(), and QPixmap::convertFromImage().
*/
func (this *QClipboard) Pixmap__() *QPixmap /*123*/ {
	// arg: 0, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QClipboard6pixmapENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qclipboard.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImage(const QImage &, QClipboard::Mode)

/*
Copies the image into the clipboard.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the image is stored in the global clipboard. If mode is QClipboard::Selection, the data is stored in the global mouse selection.

This is shorthand for:


  QMimeData *data = new QMimeData;
  data->setImageData(image);
  clipboard->setMimeData(data, mode);



See also image(), setPixmap(), and setMimeData().
*/
func (this *QClipboard) SetImage(arg0 QImage_ITF, mode int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard8setImageERK6QImageNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImage(const QImage &, QClipboard::Mode)

/*
Copies the image into the clipboard.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the image is stored in the global clipboard. If mode is QClipboard::Selection, the data is stored in the global mouse selection.

This is shorthand for:


  QMimeData *data = new QMimeData;
  data->setImageData(image);
  clipboard->setMimeData(data, mode);



See also image(), setPixmap(), and setMimeData().
*/
func (this *QClipboard) SetImage__(arg0 QImage_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	// arg: 1, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard8setImageERK6QImageNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &, QClipboard::Mode)

/*
Copies pixmap into the clipboard. Note that this is slower than setImage() because it needs to convert the QPixmap to a QImage first.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the pixmap is stored in the global clipboard. If mode is QClipboard::Selection, the pixmap is stored in the global mouse selection.

See also pixmap(), setImage(), and setMimeData().
*/
func (this *QClipboard) SetPixmap(arg0 QPixmap_ITF, mode int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard9setPixmapERK7QPixmapNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &, QClipboard::Mode)

/*
Copies pixmap into the clipboard. Note that this is slower than setImage() because it needs to convert the QPixmap to a QImage first.

The mode argument is used to control which part of the system clipboard is used. If mode is QClipboard::Clipboard, the pixmap is stored in the global clipboard. If mode is QClipboard::Selection, the pixmap is stored in the global mouse selection.

See also pixmap(), setImage(), and setMimeData().
*/
func (this *QClipboard) SetPixmap__(arg0 QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	// arg: 1, QClipboard::Mode=Enum, QClipboard::Mode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard9setPixmapERK7QPixmapNS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed(QClipboard::Mode)

/*
This signal is emitted when the data for the given clipboard mode is changed.

This function was introduced in  Qt 4.2.

See also dataChanged(), selectionChanged(), and findBufferChanged().
*/
func (this *QClipboard) Changed(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard7changedENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()

/*
This signal is emitted when the selection is changed. This only applies to windowing systems that support selections, e.g. X11. Windows and macOS don't support selections.

See also dataChanged(), findBufferChanged(), and changed().
*/
func (this *QClipboard) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void findBufferChanged()

/*
This signal is emitted when the find buffer is changed. This only applies to macOS.

With Qt version 4.3 or higher, clipboard changes made by other applications will only be detected when the application is activated.

This function was introduced in  Qt 4.2.

See also dataChanged(), selectionChanged(), and changed().
*/
func (this *QClipboard) FindBufferChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard17findBufferChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qclipboard.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dataChanged()

/*
This signal is emitted when the clipboard data is changed.

On macOS and with Qt version 4.3 or higher, clipboard changes made by other applications will only be detected when the application is activated.

See also findBufferChanged(), selectionChanged(), and changed().
*/
func (this *QClipboard) DataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboard11dataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQClipboard(this *QClipboard) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QClipboardD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum type is used to control which part of the system clipboard is used by QClipboard::mimeData(), QClipboard::setMimeData() and related functions.



See also QClipboard::supportsSelection().

*/
type QClipboard__Mode = int

// indicates that data should be stored and retrieved from the global clipboard.
const QClipboard__Clipboard QClipboard__Mode = 0

//
const QClipboard__Selection QClipboard__Mode = 1

// indicates that data should be stored and retrieved from the Find buffer. This mode is used for holding search strings on macOS.
const QClipboard__FindBuffer QClipboard__Mode = 2

//
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
