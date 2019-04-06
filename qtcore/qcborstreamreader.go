package qtcore

// /usr/include/qt/QtCore/qcborstream.h
// #include <qcborstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QCborStreamReader struct {
	*qtrt.CObject
}
type QCborStreamReader_ITF interface {
	QCborStreamReader_PTR() *QCborStreamReader
}

func (ptr *QCborStreamReader) QCborStreamReader_PTR() *QCborStreamReader { return ptr }

func (this *QCborStreamReader) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborStreamReader) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborStreamReaderFromPointer(cthis unsafe.Pointer) *QCborStreamReader {
	return &QCborStreamReader{&qtrt.CObject{cthis}}
}
func (*QCborStreamReader) NewFromPointer(cthis unsafe.Pointer) *QCborStreamReader {
	return NewQCborStreamReaderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcborstream.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCborStreamReader()

/*

 */
func (*QCborStreamReader) NewForInherit() *QCborStreamReader {
	return NewQCborStreamReader()
}
func NewQCborStreamReader() *QCborStreamReader {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReaderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qcborstream.h:156
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCborStreamReader(const char *, qsizetype)

/*

 */
func (*QCborStreamReader) NewForInherit1(data string, len_ int64) *QCborStreamReader {
	return NewQCborStreamReader1(data, len_)
}
func NewQCborStreamReader1(data string, len_ int64) *QCborStreamReader {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReaderC2EPKcx", qtrt.FFI_TYPE_POINTER, convArg0, len_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qcborstream.h:157
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCborStreamReader(const quint8 *, qsizetype)

/*

 */
func (*QCborStreamReader) NewForInherit2(data unsafe.Pointer /*666*/, len_ int64) *QCborStreamReader {
	return NewQCborStreamReader2(data, len_)
}
func NewQCborStreamReader2(data unsafe.Pointer /*666*/, len_ int64) *QCborStreamReader {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReaderC2EPKhx", qtrt.FFI_TYPE_POINTER, data, len_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qcborstream.h:158
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCborStreamReader(const QByteArray &)

/*

 */
func (*QCborStreamReader) NewForInherit3(data QByteArray_ITF) *QCborStreamReader {
	return NewQCborStreamReader3(data)
}
func NewQCborStreamReader3(data QByteArray_ITF) *QCborStreamReader {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReaderC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qcborstream.h:159
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QCborStreamReader(QIODevice *)

/*

 */
func (*QCborStreamReader) NewForInherit4(device QIODevice_ITF /*777 QIODevice **/) *QCborStreamReader {
	return NewQCborStreamReader4(device)
}
func NewQCborStreamReader4(device QIODevice_ITF /*777 QIODevice **/) *QCborStreamReader {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReaderC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qcborstream.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCborStreamReader()

/*

 */
func DeleteQCborStreamReader(this *QCborStreamReader) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReaderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcborstream.h:161
// index:0
// Public inline Visibility=Default Availability=NotAvailable
// [24] QCborStreamReader & operator=(const QCborStreamReader &)

/*

 */
func (this *QCborStreamReader) Operator_equal(arg0 QCborStreamReader_ITF) *QCborStreamReader {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCborStreamReader_PTR() != nil {
		convArg0 = arg0.QCborStreamReader_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReaderaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborStreamReaderFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborStreamReader)
	return rv2
}

// /usr/include/qt/QtCore/qcborstream.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*

 */
func (this *QCborStreamReader) SetDevice(device QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*

 */
func (this *QCborStreamReader) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qcborstream.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addData(const QByteArray &)

/*

 */
func (this *QCborStreamReader) AddData(data QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader7addDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:166
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addData(const char *, qsizetype)

/*

 */
func (this *QCborStreamReader) AddData1(data string, len_ int64) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader7addDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:167
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void addData(const quint8 *, qsizetype)

/*

 */
func (this *QCborStreamReader) AddData2(data unsafe.Pointer /*666*/, len_ int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader7addDataEPKhx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), data, len_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reparse()

/*

 */
func (this *QCborStreamReader) Reparse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader7reparseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*

 */
func (this *QCborStreamReader) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*

 */
func (this *QCborStreamReader) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:173
// index:0
// Public Visibility=Default Availability=Available
// [4] QCborError lastError()

/*

 */
func (this *QCborStreamReader) LastError() *QCborError /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader9lastErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborErrorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborError)
	return rv2
}

// /usr/include/qt/QtCore/qcborstream.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 currentOffset() const

/*

 */
func (this *QCborStreamReader) CurrentOffset() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader13currentOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcborstream.h:177
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QCborStreamReader) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:179
// index:0
// Public Visibility=Default Availability=Available
// [4] int containerDepth() const

/*

 */
func (this *QCborStreamReader) ContainerDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader14containerDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcborstream.h:180
// index:0
// Public Visibility=Default Availability=Available
// [1] QCborStreamReader::Type parentContainerType() const

/*

 */
func (this *QCborStreamReader) ParentContainerType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader19parentContainerTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborstream.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNext() const

/*

 */
func (this *QCborStreamReader) HasNext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader7hasNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool next(int)

/*

 */
func (this *QCborStreamReader) Next(maxRecursion int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader4nextEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxRecursion)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool next(int)

/*

 */
func (this *QCborStreamReader) Nextp() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	maxRecursion := int(10000)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader4nextEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxRecursion)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QCborStreamReader::Type type() const

/*

 */
func (this *QCborStreamReader) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborstream.h:185
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUnsignedInteger() const

/*

 */
func (this *QCborStreamReader) IsUnsignedInteger() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader17isUnsignedIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNegativeInteger() const

/*

 */
func (this *QCborStreamReader) IsNegativeInteger() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader17isNegativeIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInteger() const

/*

 */
func (this *QCborStreamReader) IsInteger() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader9isIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isByteArray() const

/*

 */
func (this *QCborStreamReader) IsByteArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader11isByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isString() const

/*

 */
func (this *QCborStreamReader) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:190
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isArray() const

/*

 */
func (this *QCborStreamReader) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:191
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isMap() const

/*

 */
func (this *QCborStreamReader) IsMap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader5isMapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTag() const

/*

 */
func (this *QCborStreamReader) IsTag() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader5isTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:193
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSimpleType() const

/*

 */
func (this *QCborStreamReader) IsSimpleType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader12isSimpleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:199
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool isSimpleType(QCborSimpleType) const

/*

 */
func (this *QCborStreamReader) IsSimpleType1(st int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader12isSimpleTypeE15QCborSimpleType", qtrt.FFI_TYPE_POINTER, this.GetCthis(), st)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:194
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFloat16() const

/*

 */
func (this *QCborStreamReader) IsFloat16() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader9isFloat16Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:195
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFloat() const

/*

 */
func (this *QCborStreamReader) IsFloat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader7isFloatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDouble() const

/*

 */
func (this *QCborStreamReader) IsDouble() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader8isDoubleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:197
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInvalid() const

/*

 */
func (this *QCborStreamReader) IsInvalid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader9isInvalidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:200
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFalse() const

/*

 */
func (this *QCborStreamReader) IsFalse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader7isFalseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:201
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTrue() const

/*

 */
func (this *QCborStreamReader) IsTrue() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader6isTrueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBool() const

/*

 */
func (this *QCborStreamReader) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*

 */
func (this *QCborStreamReader) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:204
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndefined() const

/*

 */
func (this *QCborStreamReader) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:206
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLengthKnown() const

/*

 */
func (this *QCborStreamReader) IsLengthKnown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader13isLengthKnownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:207
// index:0
// Public Visibility=Default Availability=Available
// [8] quint64 length() const

/*

 */
func (this *QCborStreamReader) Length() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qcborstream.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isContainer() const

/*

 */
func (this *QCborStreamReader) IsContainer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader11isContainerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:210
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool enterContainer()

/*

 */
func (this *QCborStreamReader) EnterContainer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader14enterContainerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool leaveContainer()

/*

 */
func (this *QCborStreamReader) LeaveContainer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamReader14leaveContainerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qsizetype currentStringChunkSize() const

/*

 */
func (this *QCborStreamReader) CurrentStringChunkSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader22currentStringChunkSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcborstream.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toBool() const

/*

 */
func (this *QCborStreamReader) ToBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader6toBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QCborTag toTag() const

/*

 */
func (this *QCborStreamReader) ToTag() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader5toTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborstream.h:220
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quint64 toUnsignedInteger() const

/*

 */
func (this *QCborStreamReader) ToUnsignedInteger() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader17toUnsignedIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qcborstream.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QCborNegativeInteger toNegativeInteger() const

/*

 */
func (this *QCborStreamReader) ToNegativeInteger() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader17toNegativeIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborstream.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QCborSimpleType toSimpleType() const

/*

 */
func (this *QCborStreamReader) ToSimpleType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader12toSimpleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborstream.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [2] qfloat16 toFloat16() const

/*

 */
func (this *QCborStreamReader) ToFloat16() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader9toFloat16Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborstream.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float toFloat() const

/*

 */
func (this *QCborStreamReader) ToFloat() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader7toFloatEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qcborstream.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double toDouble() const

/*

 */
func (this *QCborStreamReader) ToDouble() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader8toDoubleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qcborstream.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 toInteger() const

/*

 */
func (this *QCborStreamReader) ToInteger() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamReader9toIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

/*


 */
type QCborStreamReader__Type = int

//
const QCborStreamReader__UnsignedInteger QCborStreamReader__Type = 0

//
const QCborStreamReader__NegativeInteger QCborStreamReader__Type = 32

//
const QCborStreamReader__ByteString QCborStreamReader__Type = 64

//
const QCborStreamReader__ByteArray QCborStreamReader__Type = 64

//
const QCborStreamReader__TextString QCborStreamReader__Type = 96

//
const QCborStreamReader__String QCborStreamReader__Type = 96

//
const QCborStreamReader__Array QCborStreamReader__Type = -128

//
const QCborStreamReader__Map QCborStreamReader__Type = -96

//
const QCborStreamReader__Tag QCborStreamReader__Type = -64

//
const QCborStreamReader__SimpleType QCborStreamReader__Type = -32

//
const QCborStreamReader__HalfFloat QCborStreamReader__Type = -7

//
const QCborStreamReader__Float16 QCborStreamReader__Type = -7

//
const QCborStreamReader__Float QCborStreamReader__Type = -6

//
const QCborStreamReader__Double QCborStreamReader__Type = -5

//
const QCborStreamReader__Invalid QCborStreamReader__Type = -1

func (this *QCborStreamReader) TypeItemName(val int) string {
	switch val {
	case QCborStreamReader__UnsignedInteger: // 0
		return "UnsignedInteger"
	case QCborStreamReader__NegativeInteger: // 32
		return "NegativeInteger"
	case QCborStreamReader__ByteString: // 64
		return "ByteString,ByteArray"
		// case QCborStreamReader__ByteArray: // 64
		// return ""
	case QCborStreamReader__TextString: // 96
		return "TextString,String"
		// case QCborStreamReader__String: // 96
		// return ""
	case QCborStreamReader__Array: // -128
		return "Array"
	case QCborStreamReader__Map: // -96
		return "Map"
	case QCborStreamReader__Tag: // -64
		return "Tag"
	case QCborStreamReader__SimpleType: // -32
		return "SimpleType"
	case QCborStreamReader__HalfFloat: // -7
		return "HalfFloat,Float16"
		// case QCborStreamReader__Float16: // -7
		// return ""
	case QCborStreamReader__Float: // -6
		return "Float"
	case QCborStreamReader__Double: // -5
		return "Double"
	case QCborStreamReader__Invalid: // -1
		return "Invalid"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCborStreamReader_TypeItemName(val int) string {
	var nilthis *QCborStreamReader
	return nilthis.TypeItemName(val)
}

/*


 */
type QCborStreamReader__StringResultCode = int

//
const QCborStreamReader__EndOfString QCborStreamReader__StringResultCode = 0

//
const QCborStreamReader__Ok QCborStreamReader__StringResultCode = 1

//
const QCborStreamReader__Error QCborStreamReader__StringResultCode = -1

func (this *QCborStreamReader) StringResultCodeItemName(val int) string {
	switch val {
	case QCborStreamReader__EndOfString: // 0
		return "EndOfString"
	case QCborStreamReader__Ok: // 1
		return "Ok"
	case QCborStreamReader__Error: // -1
		return "Error"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCborStreamReader_StringResultCodeItemName(val int) string {
	var nilthis *QCborStreamReader
	return nilthis.StringResultCodeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10361() {
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
