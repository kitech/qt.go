package qtmultimedia

// /usr/include/qt/QtMultimedia/qvideoframe.h
// #include <qvideoframe.h>
// #include <QtMultimedia>

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
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QVideoFrame struct {
	*qtrt.CObject
}
type QVideoFrame_ITF interface {
	QVideoFrame_PTR() *QVideoFrame
}

func (ptr *QVideoFrame) QVideoFrame_PTR() *QVideoFrame { return ptr }

func (this *QVideoFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVideoFrame) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVideoFrameFromPointer(cthis unsafe.Pointer) *QVideoFrame {
	return &QVideoFrame{&qtrt.CObject{cthis}}
}
func (*QVideoFrame) NewFromPointer(cthis unsafe.Pointer) *QVideoFrame {
	return NewQVideoFrameFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVideoFrame()

/*
Constructs a null video frame.
*/
func NewQVideoFrame() *QVideoFrame {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrameC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoFrame)
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:112
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVideoFrame(QAbstractVideoBuffer *, const QSize &, QVideoFrame::PixelFormat)

/*
Constructs a null video frame.
*/
func NewQVideoFrame_1(buffer QAbstractVideoBuffer_ITF /*777 QAbstractVideoBuffer **/, size qtcore.QSize_ITF, format int) *QVideoFrame {
	var convArg0 unsafe.Pointer
	if buffer != nil && buffer.QAbstractVideoBuffer_PTR() != nil {
		convArg0 = buffer.QAbstractVideoBuffer_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrameC2EP20QAbstractVideoBufferRK5QSizeNS_11PixelFormatE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoFrame)
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:113
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QVideoFrame(int, const QSize &, int, QVideoFrame::PixelFormat)

/*
Constructs a null video frame.
*/
func NewQVideoFrame_2(bytes int, size qtcore.QSize_ITF, bytesPerLine int, format int) *QVideoFrame {
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrameC2EiRK5QSizeiNS_11PixelFormatE", qtrt.FFI_TYPE_POINTER, bytes, convArg1, bytesPerLine, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoFrame)
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:114
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QVideoFrame(const QImage &)

/*
Constructs a null video frame.
*/
func NewQVideoFrame_3(image qtgui.QImage_ITF) *QVideoFrame {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrameC2ERK6QImage", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVideoFrame)
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QVideoFrame()

/*

 */
func DeleteQVideoFrame(this *QVideoFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QVideoFrame & operator=(const QVideoFrame &)

/*

 */
func (this *QVideoFrame) Operator_equal(other QVideoFrame_ITF) *QVideoFrame {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVideoFrame_PTR() != nil {
		convArg0 = other.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrameaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoFrame)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QVideoFrame &) const

/*

 */
func (this *QVideoFrame) Operator_equal_equal(other QVideoFrame_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVideoFrame_PTR() != nil {
		convArg0 = other.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrameeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QVideoFrame &) const

/*

 */
func (this *QVideoFrame) Operator_not_equal(other QVideoFrame_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVideoFrame_PTR() != nil {
		convArg0 = other.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrameneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Identifies whether a video frame is valid.

An invalid frame has no video buffer associated with it.

Returns true if the frame is valid, and false if it is not.
*/
func (this *QVideoFrame) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] QVideoFrame::PixelFormat pixelFormat() const

/*
Returns the color format of a video frame.
*/
func (this *QVideoFrame) PixelFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame11pixelFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractVideoBuffer::HandleType handleType() const

/*
Returns the type of a video frame's handle.
*/
func (this *QVideoFrame) HandleType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame10handleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*
Returns the dimensions of a video frame.
*/
func (this *QVideoFrame) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:129
// index:0
// Public Visibility=Default Availability=Available
// [4] int width() const

/*
Returns the width of a video frame.
*/
func (this *QVideoFrame) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] int height() const

/*
Returns the height of a video frame.
*/
func (this *QVideoFrame) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] QVideoFrame::FieldType fieldType() const

/*
Returns the field an interlaced video frame belongs to.

If the video is not interlaced this will return WholeFrame.

See also setFieldType().
*/
func (this *QVideoFrame) FieldType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame9fieldTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldType(QVideoFrame::FieldType)

/*
Sets the field an interlaced video frame belongs to.

See also fieldType().
*/
func (this *QVideoFrame) SetFieldType(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame12setFieldTypeENS_9FieldTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMapped() const

/*
Identifies if a video frame's contents are currently mapped to system memory.

This is a convenience function which checks that the MapMode of the frame is not equal to QAbstractVideoBuffer::NotMapped.

Returns true if the contents of the video frame are mapped to system memory, and false otherwise.

See also mapMode() and QAbstractVideoBuffer::MapMode.
*/
func (this *QVideoFrame) IsMapped() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame8isMappedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable() const

/*
Identifies if the mapped contents of a video frame were read from the frame when it was mapped.

This is a convenience function which checks if the MapMode contains the QAbstractVideoBuffer::WriteOnly flag.

Returns true if the contents of the mapped memory were read from the video frame, and false otherwise.

See also mapMode() and QAbstractVideoBuffer::MapMode.
*/
func (this *QVideoFrame) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable() const

/*
Identifies if the mapped contents of a video frame will be persisted when the frame is unmapped.

This is a convenience function which checks if the MapMode contains the QAbstractVideoBuffer::WriteOnly flag.

Returns true if the video frame will be updated when unmapped, and false otherwise.

Note: The result of altering the data of a frame that is mapped in read-only mode is undefined. Depending on the buffer implementation the changes may be persisted, or worse alter a shared buffer.

See also mapMode() and QAbstractVideoBuffer::MapMode.
*/
func (this *QVideoFrame) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractVideoBuffer::MapMode mapMode() const

/*
Returns the mode a video frame was mapped to system memory in.

See also map() and QAbstractVideoBuffer::MapMode.
*/
func (this *QVideoFrame) MapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame7mapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool map(QAbstractVideoBuffer::MapMode)

/*
Maps the contents of a video frame to system (CPU addressable) memory.

In some cases the video frame data might be stored in video memory or otherwise inaccessible memory, so it is necessary to map a frame before accessing the pixel data. This may involve copying the contents around, so avoid mapping and unmapping unless required.

The map mode indicates whether the contents of the mapped memory should be read from and/or written to the frame. If the map mode includes the QAbstractVideoBuffer::ReadOnly flag the mapped memory will be populated with the content of the video frame when initially mapped. If the map mode includes the QAbstractVideoBuffer::WriteOnly flag the content of the possibly modified mapped memory will be written back to the frame when unmapped.

While mapped the contents of a video frame can be accessed directly through the pointer returned by the bits() function.

When access to the data is no longer needed, be sure to call the unmap() function to release the mapped memory and possibly update the video frame contents.

If the video frame has been mapped in read only mode, it is permissible to map it multiple times in read only mode (and unmap it a corresponding number of times). In all other cases it is necessary to unmap the frame first before mapping a second time.

Note: Writing to memory that is mapped as read-only is undefined, and may result in changes to shared data or crashes.

Returns true if the frame was mapped to memory in the given mode and false otherwise.

See also unmap(), mapMode(), and bits().
*/
func (this *QVideoFrame) Map(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame3mapEN20QAbstractVideoBuffer7MapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unmap()

/*
Releases the memory mapped by the map() function.

If the MapMode included the QAbstractVideoBuffer::WriteOnly flag this will persist the current content of the mapped memory to the video frame.

unmap() should not be called if map() function failed.

See also map().
*/
func (this *QVideoFrame) Unmap() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame5unmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] int bytesPerLine() const

/*
Returns the number of bytes in a scan line.

Note: For planar formats this is the bytes per line of the first plane only. The bytes per line of subsequent planes should be calculated as per the frame pixel format.

This value is only valid while the frame data is mapped.

See also bits(), map(), and mappedBytes().
*/
func (this *QVideoFrame) BytesPerLine() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame12bytesPerLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:145
// index:1
// Public Visibility=Default Availability=Available
// [4] int bytesPerLine(int) const

/*
Returns the number of bytes in a scan line.

Note: For planar formats this is the bytes per line of the first plane only. The bytes per line of subsequent planes should be calculated as per the frame pixel format.

This value is only valid while the frame data is mapped.

See also bits(), map(), and mappedBytes().
*/
func (this *QVideoFrame) BytesPerLine_1(plane int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame12bytesPerLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), plane)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] uchar * bits()

/*
Returns a pointer to the start of the frame data buffer.

This value is only valid while the frame data is mapped.

Changes made to data accessed via this pointer (when mapped with write access) are only guaranteed to have been persisted when unmap() is called and when the buffer has been mapped for writing.

See also map(), mappedBytes(), and bytesPerLine().
*/
func (this *QVideoFrame) Bits() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame4bitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:148
// index:1
// Public Visibility=Default Availability=Available
// [8] uchar * bits(int)

/*
Returns a pointer to the start of the frame data buffer.

This value is only valid while the frame data is mapped.

Changes made to data accessed via this pointer (when mapped with write access) are only guaranteed to have been persisted when unmap() is called and when the buffer has been mapped for writing.

See also map(), mappedBytes(), and bytesPerLine().
*/
func (this *QVideoFrame) Bits_1(plane int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame4bitsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), plane)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:149
// index:2
// Public Visibility=Default Availability=Available
// [8] const uchar * bits() const

/*
Returns a pointer to the start of the frame data buffer.

This value is only valid while the frame data is mapped.

Changes made to data accessed via this pointer (when mapped with write access) are only guaranteed to have been persisted when unmap() is called and when the buffer has been mapped for writing.

See also map(), mappedBytes(), and bytesPerLine().
*/
func (this *QVideoFrame) Bits_2() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame4bitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:150
// index:3
// Public Visibility=Default Availability=Available
// [8] const uchar * bits(int) const

/*
Returns a pointer to the start of the frame data buffer.

This value is only valid while the frame data is mapped.

Changes made to data accessed via this pointer (when mapped with write access) are only guaranteed to have been persisted when unmap() is called and when the buffer has been mapped for writing.

See also map(), mappedBytes(), and bytesPerLine().
*/
func (this *QVideoFrame) Bits_3(plane int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame4bitsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), plane)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:151
// index:0
// Public Visibility=Default Availability=Available
// [4] int mappedBytes() const

/*
Returns the number of bytes occupied by the mapped frame data.

This value is only valid while the frame data is mapped.

See also map().
*/
func (this *QVideoFrame) MappedBytes() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame11mappedBytesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:152
// index:0
// Public Visibility=Default Availability=Available
// [4] int planeCount() const

/*
Returns the number of planes in the video frame.

This value is only valid while the frame data is mapped.

This function was introduced in  Qt 5.4.

See also map().
*/
func (this *QVideoFrame) PlaneCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame10planeCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:154
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant handle() const

/*
Returns a type specific handle to a video frame's buffer.

For an OpenGL texture this would be the texture ID.

See also QAbstractVideoBuffer::handle().
*/
func (this *QVideoFrame) Handle() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 startTime() const

/*
Returns the presentation time (in microseconds) when the frame should be displayed.

An invalid time is represented as -1.

See also setStartTime().
*/
func (this *QVideoFrame) StartTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame9startTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartTime(qint64)

/*
Sets the presentation time (in microseconds) when the frame should initially be displayed.

An invalid time is represented as -1.

See also startTime().
*/
func (this *QVideoFrame) SetStartTime(time int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame12setStartTimeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 endTime() const

/*
Returns the presentation time (in microseconds) when a frame should stop being displayed.

An invalid time is represented as -1.

See also setEndTime().
*/
func (this *QVideoFrame) EndTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame7endTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEndTime(qint64)

/*
Sets the presentation time (in microseconds) when a frame should stop being displayed.

An invalid time is represented as -1.

See also endTime().
*/
func (this *QVideoFrame) SetEndTime(time int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame10setEndTimeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:163
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant metaData(const QString &) const

/*
Returns any metadata for this frame for the given key.

This might include frame specific information from a camera, or subtitles from a decoded video stream.

See the documentation for the relevant video frame producer for further information about available metadata.

See also setMetaData().
*/
func (this *QVideoFrame) MetaData(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVideoFrame8metaDataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMetaData(const QString &, const QVariant &)

/*
Sets the metadata for the given key to value.

If value is a null variant, any metadata for this key will be removed.

The producer of the video frame might use this to associate certain data with this frame, or for an intermediate processor to add information for a consumer of this frame.

See also metaData().
*/
func (this *QVideoFrame) SetMetaData(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame11setMetaDataERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:166
// index:0
// Public static Visibility=Default Availability=Available
// [4] QVideoFrame::PixelFormat pixelFormatFromImageFormat(QImage::Format)

/*
Returns a video pixel format equivalent to an image format. If there is no equivalent format QVideoFrame::InvalidType is returned instead.

Note: In general QImage does not handle YUV formats.
*/
func (this *QVideoFrame) PixelFormatFromImageFormat(format int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame26pixelFormatFromImageFormatEN6QImage6FormatE", qtrt.FFI_TYPE_POINTER, format)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QVideoFrame_PixelFormatFromImageFormat(format int) int {
	var nilthis *QVideoFrame
	rv := nilthis.PixelFormatFromImageFormat(format)
	return rv
}

// /usr/include/qt/QtMultimedia/qvideoframe.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [4] QImage::Format imageFormatFromPixelFormat(QVideoFrame::PixelFormat)

/*
Returns an image format equivalent to a video frame pixel format. If there is no equivalent format QImage::Format_Invalid is returned instead.

Note: In general QImage does not handle YUV formats.
*/
func (this *QVideoFrame) ImageFormatFromPixelFormat(format int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVideoFrame26imageFormatFromPixelFormatENS_11PixelFormatE", qtrt.FFI_TYPE_POINTER, format)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QVideoFrame_ImageFormatFromPixelFormat(format int) int {
	var nilthis *QVideoFrame
	rv := nilthis.ImageFormatFromPixelFormat(format)
	return rv
}

/*
Specifies the field an interlaced video frame belongs to.


*/
type QVideoFrame__FieldType = int

// The frame is not interlaced.
const QVideoFrame__ProgressiveFrame QVideoFrame__FieldType = 0

// The frame contains a top field.
const QVideoFrame__TopField QVideoFrame__FieldType = 1

// The frame contains a bottom field.
const QVideoFrame__BottomField QVideoFrame__FieldType = 2

// The frame contains a merged top and bottom field.
const QVideoFrame__InterlacedFrame QVideoFrame__FieldType = 3

/*
Enumerates video data types.


*/
type QVideoFrame__PixelFormat = int

// The frame is invalid.
const QVideoFrame__Format_Invalid QVideoFrame__PixelFormat = 0

//
const QVideoFrame__Format_ARGB32 QVideoFrame__PixelFormat = 1

//
const QVideoFrame__Format_ARGB32_Premultiplied QVideoFrame__PixelFormat = 2

//
const QVideoFrame__Format_RGB32 QVideoFrame__PixelFormat = 3

//
const QVideoFrame__Format_RGB24 QVideoFrame__PixelFormat = 4

//
const QVideoFrame__Format_RGB565 QVideoFrame__PixelFormat = 5

//
const QVideoFrame__Format_RGB555 QVideoFrame__PixelFormat = 6

//
const QVideoFrame__Format_ARGB8565_Premultiplied QVideoFrame__PixelFormat = 7

//
const QVideoFrame__Format_BGRA32 QVideoFrame__PixelFormat = 8

//
const QVideoFrame__Format_BGRA32_Premultiplied QVideoFrame__PixelFormat = 9

//
const QVideoFrame__Format_BGR32 QVideoFrame__PixelFormat = 10

//
const QVideoFrame__Format_BGR24 QVideoFrame__PixelFormat = 11

//
const QVideoFrame__Format_BGR565 QVideoFrame__PixelFormat = 12

//
const QVideoFrame__Format_BGR555 QVideoFrame__PixelFormat = 13

//
const QVideoFrame__Format_BGRA5658_Premultiplied QVideoFrame__PixelFormat = 14

//
const QVideoFrame__Format_AYUV444 QVideoFrame__PixelFormat = 15

//
const QVideoFrame__Format_AYUV444_Premultiplied QVideoFrame__PixelFormat = 16

//
const QVideoFrame__Format_YUV444 QVideoFrame__PixelFormat = 17

//
const QVideoFrame__Format_YUV420P QVideoFrame__PixelFormat = 18

//
const QVideoFrame__Format_YV12 QVideoFrame__PixelFormat = 19

//
const QVideoFrame__Format_UYVY QVideoFrame__PixelFormat = 20

//
const QVideoFrame__Format_YUYV QVideoFrame__PixelFormat = 21

//
const QVideoFrame__Format_NV12 QVideoFrame__PixelFormat = 22

//
const QVideoFrame__Format_NV21 QVideoFrame__PixelFormat = 23

//
const QVideoFrame__Format_IMC1 QVideoFrame__PixelFormat = 24

//
const QVideoFrame__Format_IMC2 QVideoFrame__PixelFormat = 25

//
const QVideoFrame__Format_IMC3 QVideoFrame__PixelFormat = 26

//
const QVideoFrame__Format_IMC4 QVideoFrame__PixelFormat = 27

//
const QVideoFrame__Format_Y8 QVideoFrame__PixelFormat = 28

//
const QVideoFrame__Format_Y16 QVideoFrame__PixelFormat = 29

//
const QVideoFrame__Format_Jpeg QVideoFrame__PixelFormat = 30

//
const QVideoFrame__Format_CameraRaw QVideoFrame__PixelFormat = 31

//
const QVideoFrame__Format_AdobeDng QVideoFrame__PixelFormat = 32

//
const QVideoFrame__NPixelFormats QVideoFrame__PixelFormat = 33

//
const QVideoFrame__Format_User QVideoFrame__PixelFormat = 1000

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
