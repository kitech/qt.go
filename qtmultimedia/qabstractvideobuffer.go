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
// extern C begin: 45
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
type QAbstractVideoBuffer struct {
	*qtrt.CObject
}
type QAbstractVideoBuffer_ITF interface {
	QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer
}

func (ptr *QAbstractVideoBuffer) QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer { return ptr }

func (this *QAbstractVideoBuffer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractVideoBuffer) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAbstractVideoBufferFromPointer(cthis unsafe.Pointer) *QAbstractVideoBuffer {
	return &QAbstractVideoBuffer{&qtrt.CObject{cthis}}
}
func (*QAbstractVideoBuffer) NewFromPointer(cthis unsafe.Pointer) *QAbstractVideoBuffer {
	return NewQAbstractVideoBufferFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractVideoBuffer(QAbstractVideoBuffer::HandleType)

/*
Constructs an abstract video buffer of the given type.
*/
func NewQAbstractVideoBuffer(type_ int) *QAbstractVideoBuffer {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoBufferC2ENS_10HandleTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractVideoBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAbstractVideoBuffer)
	return gothis
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractVideoBuffer()

/*

 */
func DeleteQAbstractVideoBuffer(this *QAbstractVideoBuffer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoBufferD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void release()

/*
Releases the video buffer.

QVideoFrame calls QAbstractVideoBuffer::release when the buffer is not used any more and can be destroyed or returned to the buffer pool.

The default implementation deletes the buffer instance.
*/
func (this *QAbstractVideoBuffer) Release() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoBuffer7releaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractVideoBuffer::HandleType handleType() const

/*
Returns the type of a video buffer's handle.

See also handle().
*/
func (this *QAbstractVideoBuffer) HandleType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractVideoBuffer10handleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:84
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAbstractVideoBuffer::MapMode mapMode() const

/*
Returns the mode a video buffer is mapped in.

See also map().
*/
func (this *QAbstractVideoBuffer) MapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractVideoBuffer7mapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
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
func (this *QAbstractVideoBuffer) Map(mode int, numBytes unsafe.Pointer /*666*/, bytesPerLine unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoBuffer3mapENS_7MapModeEPiS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, numBytes, bytesPerLine)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int mapPlanes(QAbstractVideoBuffer::MapMode, int *, int *, uchar **)

/*
Independently maps the planes of a video buffer to memory.

The map mode indicates whether the contents of the mapped memory should be read from and/or written to the buffer. If the map mode includes the QAbstractVideoBuffer::ReadOnly flag the mapped memory will be populated with the content of the buffer when initially mapped. If the map mode includes the QAbstractVideoBuffer::WriteOnly flag the content of the possibly modified mapped memory will be written back to the buffer when unmapped.

When access to the data is no longer needed be sure to call the unmap() function to release the mapped memory and possibly update the buffer contents.

Returns the number of planes in the mapped video data. For each plane the line stride of that plane will be returned in bytesPerLine, and a pointer to the plane data will be returned in data. The accumulative size of the mapped data is returned in numBytes.

Not all buffer implementations will map more than the first plane, if this returns a single plane for a planar format the additional planes will have to be calculated from the line stride of the first plane and the frame height. Mapping a buffer with QVideoFrame will do this for you.

To implement this function create a derivative of QAbstractPlanarVideoBuffer and implement its map function instance instead.

This function was introduced in  Qt 5.4.
*/
func (this *QAbstractVideoBuffer) MapPlanes(mode int, numBytes unsafe.Pointer /*666*/, bytesPerLine unsafe.Pointer, data unsafe.Pointer) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoBuffer9mapPlanesENS_7MapModeEPiS1_PPh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, numBytes, bytesPerLine, data)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void unmap()

/*
Releases the memory mapped by the map() function.

If the MapMode included the QAbstractVideoBuffer::WriteOnly flag this will write the current content of the mapped memory back to the video frame.

See also map().
*/
func (this *QAbstractVideoBuffer) Unmap() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAbstractVideoBuffer5unmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideobuffer.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant handle() const

/*
Returns a type specific handle to the data buffer.

The type of the handle is given by handleType() function.

See also handleType().
*/
func (this *QAbstractVideoBuffer) Handle() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAbstractVideoBuffer6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

/*
Identifies the type of a video buffers handle.



See also handleType().

*/
type QAbstractVideoBuffer__HandleType = int

// The buffer has no handle, its data can only be accessed by mapping the buffer.
const QAbstractVideoBuffer__NoHandle QAbstractVideoBuffer__HandleType = 0

// The handle of the buffer is an OpenGL texture ID.
const QAbstractVideoBuffer__GLTextureHandle QAbstractVideoBuffer__HandleType = 1

// The handle contains pointer to shared memory XVideo image.
const QAbstractVideoBuffer__XvShmImageHandle QAbstractVideoBuffer__HandleType = 2

// The handle contains pointer to macOS CIImage.
const QAbstractVideoBuffer__CoreImageHandle QAbstractVideoBuffer__HandleType = 3

// The handle of the buffer is a QPixmap.
const QAbstractVideoBuffer__QPixmapHandle QAbstractVideoBuffer__HandleType = 4

// The handle of the buffer is an EGLImageKHR.
const QAbstractVideoBuffer__EGLImageHandle QAbstractVideoBuffer__HandleType = 5

//
const QAbstractVideoBuffer__UserHandle QAbstractVideoBuffer__HandleType = 1000

/*
Enumerates how a video buffer's data is mapped to system memory.

QAbstractVideoBuffer::ReadWriteReadOnly | WriteOnlyThe mapped memory is populated with data from the video buffer, and the video buffer is repopulated with the content of the mapped memory when it is unmapped.


See also mapMode() and map().

*/
type QAbstractVideoBuffer__MapMode = int

//
const QAbstractVideoBuffer__NotMapped QAbstractVideoBuffer__MapMode = 0

//
const QAbstractVideoBuffer__ReadOnly QAbstractVideoBuffer__MapMode = 1

//
const QAbstractVideoBuffer__WriteOnly QAbstractVideoBuffer__MapMode = 2

//
const QAbstractVideoBuffer__ReadWrite QAbstractVideoBuffer__MapMode = 3

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
