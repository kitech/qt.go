package qtmacextras

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
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

//  header block end

//  body block begin
// /usr/include/qt/QtMacExtras/../../src/macextras/qmacfunctions.h:68
// index:0
// Invalid Visibility=Default Availability=Available
// [8] NSData * toNSData(const QByteArray &)

/*

 */
func ToNSData(data qtcore.QByteArray_ITF) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtMac8toNSDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacfunctions.h:75
// index:0
// Invalid Visibility=Default Availability=Available
// [8] CGContextRef currentCGContext()

/*

 */
func CurrentCGContext() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtMac16currentCGContextEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacfunctions.h:73
// index:0
// Invalid Visibility=Default Availability=Available
// [32] QPixmap fromCGImageRef(CGImageRef)

/*

 */
func FromCGImageRef(image unsafe.Pointer /*666*/) *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtMac14fromCGImageRefEP7CGImage", qtrt.FFI_TYPE_POINTER, image)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacfunctions.h:72
// index:0
// Invalid Visibility=Default Availability=Available
// [8] CGImageRef toCGImageRef(const QPixmap &)

/*

 */
func ToCGImageRef(pixmap qtgui.QPixmap_ITF) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtMac12toCGImageRefERK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacfunctions.h:69
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QByteArray fromNSData(const NSData *)

/*

 */
func FromNSData(data unsafe.Pointer /*666*/) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtMac10fromNSDataEPK11objc_object", qtrt.FFI_TYPE_POINTER, data)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:56
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qRegisterDraggedTypes(const QStringList &)

/*
Registers the given types as custom pasteboard types.

This function should be called to enable the Drag and Drop events for custom pasteboard types on Cocoa implementations. This is required in addition to a QMacPasteboardMime subclass implementation. By default drag and drop is enabled for all standard pasteboard types.

See also QMacPasteboardMime.
*/
func QRegisterDraggedTypes(types qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if types != nil && types.QStringList_PTR() != nil {
		convArg0 = types.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z21qRegisterDraggedTypesRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end
