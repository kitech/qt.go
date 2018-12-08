package qtmultimedia

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h
// #include <qabstractvideobuffer.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QAbstractPlanarVideoBuffer struct {
	*QAbstractVideoBuffer
}
type QAbstractPlanarVideoBuffer_ITF interface {
	QAbstractVideoBuffer_ITF
	QAbstractPlanarVideoBuffer_PTR() *QAbstractPlanarVideoBuffer
}

func (ptr *QAbstractPlanarVideoBuffer) QAbstractPlanarVideoBuffer_PTR() *QAbstractPlanarVideoBuffer {
	return ptr
}

func (this *QAbstractPlanarVideoBuffer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractVideoBuffer.GetCthis()
	}
}
func (this *QAbstractPlanarVideoBuffer) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractVideoBuffer = NewQAbstractVideoBufferFromPointer(cthis)
}
func NewQAbstractPlanarVideoBufferFromPointer(cthis unsafe.Pointer) *QAbstractPlanarVideoBuffer {
	bcthis0 := NewQAbstractVideoBufferFromPointer(cthis)
	return &QAbstractPlanarVideoBuffer{bcthis0}
}
func (*QAbstractPlanarVideoBuffer) NewFromPointer(cthis unsafe.Pointer) *QAbstractPlanarVideoBuffer {
	return NewQAbstractPlanarVideoBufferFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractPlanarVideoBuffer(QAbstractVideoBuffer::HandleType)

/*

 */
func (*QAbstractPlanarVideoBuffer) NewForInherit(type_ int) *QAbstractPlanarVideoBuffer {
	return NewQAbstractPlanarVideoBuffer(type_)
}
func NewQAbstractPlanarVideoBuffer(type_ int) *QAbstractPlanarVideoBuffer {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractPlanarVideoBufferC2EN20QAbstractVideoBuffer10HandleTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractPlanarVideoBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAbstractPlanarVideoBuffer)
	return gothis
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:108
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractPlanarVideoBuffer()

/*

 */
func DeleteQAbstractPlanarVideoBuffer(this *QAbstractPlanarVideoBuffer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractPlanarVideoBufferD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:110
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] uchar * map(QAbstractVideoBuffer::MapMode, int *, int *)

/*
Maps the contents of a video buffer to memory.

In some cases the video buffer might be stored in video memory or otherwise inaccessible memory, so it is necessary to map the buffer before accessing the pixel data. This may involve copying the contents around, so avoid mapping and unmapping unless required.

The map mode indicates whether the contents of the mapped memory should be read from and/or written to the buffer. If the map mode includes the QAbstractVideoBuffer::ReadOnly flag the mapped memory will be populated with the content of the buffer when initially mapped. If the map mode includes the QAbstractVideoBuffer::WriteOnly flag the content of the possibly modified mapped memory will be written back to the buffer when unmapped.

When access to the data is no longer needed be sure to call the unmap() function to release the mapped memory and possibly update the buffer contents.

Returns a pointer to the mapped memory region, or a null pointer if the mapping failed. The size in bytes of the mapped memory region is returned in numBytes, and the line stride in bytesPerLine.

Note: Writing to memory that is mapped as read-only is undefined, and may result in changes to shared data or crashes.

See also unmap() and mapMode().
*/
func (this *QAbstractPlanarVideoBuffer) Map(mode int, numBytes unsafe.Pointer /*666*/, bytesPerLine unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractPlanarVideoBuffer3mapEN20QAbstractVideoBuffer7MapModeEPiS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, numBytes, bytesPerLine)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:111
// index:1
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int map(QAbstractVideoBuffer::MapMode, int *, int *, uchar **)

/*
Maps the contents of a video buffer to memory.

In some cases the video buffer might be stored in video memory or otherwise inaccessible memory, so it is necessary to map the buffer before accessing the pixel data. This may involve copying the contents around, so avoid mapping and unmapping unless required.

The map mode indicates whether the contents of the mapped memory should be read from and/or written to the buffer. If the map mode includes the QAbstractVideoBuffer::ReadOnly flag the mapped memory will be populated with the content of the buffer when initially mapped. If the map mode includes the QAbstractVideoBuffer::WriteOnly flag the content of the possibly modified mapped memory will be written back to the buffer when unmapped.

When access to the data is no longer needed be sure to call the unmap() function to release the mapped memory and possibly update the buffer contents.

Returns a pointer to the mapped memory region, or a null pointer if the mapping failed. The size in bytes of the mapped memory region is returned in numBytes, and the line stride in bytesPerLine.

Note: Writing to memory that is mapped as read-only is undefined, and may result in changes to shared data or crashes.

See also unmap() and mapMode().
*/
func (this *QAbstractPlanarVideoBuffer) Map1(mode int, numBytes unsafe.Pointer /*666*/, bytesPerLine unsafe.Pointer, data unsafe.Pointer) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractPlanarVideoBuffer3mapEN20QAbstractVideoBuffer7MapModeEPiS2_PPh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, numBytes, bytesPerLine, data)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
