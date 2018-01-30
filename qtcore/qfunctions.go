package qtcore

import "unsafe"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
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

//  header block end

//  body block begin
// /usr/include/qt/QtCore/qglobal.h:531
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qRound(double)
func QRound(d float64) int {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qRoundd", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:533
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qRound(float)
func QRound_1(d float32) int {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qRoundf", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:536
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qint64 qRound64(double)
func QRound64(d float64) int64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z8qRound64d", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qglobal.h:538
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] qint64 qRound64(float)
func QRound64_1(d float32) int64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z8qRound64f", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qglobal.h:681
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qSharedBuild()
func QSharedBuild() bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z12qSharedBuildv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:775
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qBadAlloc()
func QBadAlloc() {
	rv, err := ffiqt.InvokeQtFunc6("_Z9qBadAllocv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qglobal.h:796
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyCompare(double, double)
func QFuzzyCompare(p1 float64, p2 float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZL13qFuzzyComparedd", ffiqt.FFI_TYPE_POINTER, p1, p2)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:801
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyCompare(float, float)
func QFuzzyCompare_1(p1 float32, p2 float32) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZL13qFuzzyCompareff", ffiqt.FFI_TYPE_POINTER, p1, p2)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:806
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyIsNull(double)
func QFuzzyIsNull(d float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZL12qFuzzyIsNulld", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:811
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyIsNull(float)
func QFuzzyIsNull_1(f float32) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZL12qFuzzyIsNullf", ffiqt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:821
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qIsNull(double)
func QIsNull(d float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZL7qIsNulld", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:837
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool qIsNull(float)
func QIsNull_1(f float32) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZL7qIsNullf", ffiqt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:894
// index:0
// Invalid Visibility=Default Availability=Available
// [8] void * qMallocAligned(size_t, size_t)
func QMallocAligned(size uint, alignment uint) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z14qMallocAlignedmm", ffiqt.FFI_TYPE_POINTER, size, alignment)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qglobal.h:895
// index:0
// Invalid Visibility=Default Availability=Available
// [8] void * qReallocAligned(void *, size_t, size_t, size_t)
func QReallocAligned(ptr unsafe.Pointer /*666*/, size uint, oldsize uint, alignment uint) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z15qReallocAlignedPvmmm", ffiqt.FFI_TYPE_POINTER, ptr, size, oldsize, alignment)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qglobal.h:896
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qFreeAligned(void *)
func QFreeAligned(ptr unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_Z12qFreeAlignedPv", ffiqt.FFI_TYPE_POINTER, ptr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qglobal.h:1133
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qEnvironmentVariable(const char *)
func QEnvironmentVariable(varName string) *QString /*123*/ {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z20qEnvironmentVariablePKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qglobal.h:1134
// index:1
// Invalid Visibility=Default Availability=Available
// [8] QString qEnvironmentVariable(const char *, const QString &)
func QEnvironmentVariable_1(varName string, defaultValue *QString) *QString /*123*/ {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z20qEnvironmentVariablePKcRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qglobal.h:1139
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qEnvironmentVariableIsEmpty(const char *)
func QEnvironmentVariableIsEmpty(varName string) bool {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z27qEnvironmentVariableIsEmptyPKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:1140
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qEnvironmentVariableIsSet(const char *)
func QEnvironmentVariableIsSet(varName string) bool {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z25qEnvironmentVariableIsSetPKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:1141
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qEnvironmentVariableIntValue(const char *, _Bool *)
func QEnvironmentVariableIntValue(varName string, ok unsafe.Pointer /*666*/) int {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z28qEnvironmentVariableIntValuePKcPb", ffiqt.FFI_TYPE_POINTER, convArg0, &ok)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:1143
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qIntCast(double)
func QIntCast(f float64) int {
	rv, err := ffiqt.InvokeQtFunc6("_Z8qIntCastd", ffiqt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:1144
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qIntCast(float)
func QIntCast_1(f float32) int {
	rv, err := ffiqt.InvokeQtFunc6("_Z8qIntCastf", ffiqt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlogging.h:182
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qErrnoWarning(int, const char *, ...)
func QErrnoWarning(code int, msg string) {
	var convArg1 = qtrt.CString(msg)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_Z13qErrnoWarningiPKcz", ffiqt.FFI_TYPE_POINTER, code, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:183
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qErrnoWarning(const char *, ...)
func QErrnoWarning_1(msg string) {
	var convArg0 = qtrt.CString(msg)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z13qErrnoWarningPKcz", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:191
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QtMessageHandler qInstallMessageHandler(QtMessageHandler)
func QInstallMessageHandler(arg0 unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z22qInstallMessageHandlerPFv9QtMsgTypeRK18QMessageLogContextRK7QStringE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlogging.h:193
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qSetMessagePattern(const QString &)
func QSetMessagePattern(messagePattern *QString) {
	var convArg0 = messagePattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z18qSetMessagePatternRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:194
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qFormatLogMessage(enum QtMsgType, const QMessageLogContext &, const QString &)
func QFormatLogMessage(type_ int, context *QMessageLogContext, buf *QString) *QString /*123*/ {
	var convArg1 = context.GetCthis()
	var convArg2 = buf.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z17qFormatLogMessage9QtMsgTypeRK18QMessageLogContextRK7QString", ffiqt.FFI_TYPE_POINTER, type_, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qnumeric.h:48
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qIsInf(double)
func QIsInf(d float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qIsInfd", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:49
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qIsNaN(double)
func QIsNaN(d float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qIsNaNd", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:50
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qIsFinite(double)
func QIsFinite(d float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z9qIsFinited", ffiqt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:51
// index:1
// Invalid Visibility=Default Availability=Available
// [1] bool qIsInf(float)
func QIsInf_1(f float32) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qIsInff", ffiqt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:52
// index:1
// Invalid Visibility=Default Availability=Available
// [1] bool qIsNaN(float)
func QIsNaN_1(f float32) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qIsNaNf", ffiqt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:53
// index:1
// Invalid Visibility=Default Availability=Available
// [1] bool qIsFinite(float)
func QIsFinite_1(f float32) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z9qIsFinitef", ffiqt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:54
// index:0
// Invalid Visibility=Default Availability=Available
// [8] double qSNaN()
func QSNaN() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qSNaNv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:55
// index:0
// Invalid Visibility=Default Availability=Available
// [8] double qQNaN()
func QQNaN() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qQNaNv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:56
// index:0
// Invalid Visibility=Default Availability=Available
// [8] double qInf()
func QInf() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z4qInfv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:58
// index:0
// Invalid Visibility=Default Availability=Available
// [4] quint32 qFloatDistance(float, float)
func QFloatDistance(a float32, b float32) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z14qFloatDistanceff", ffiqt.FFI_TYPE_POINTER, a, b)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qnumeric.h:59
// index:1
// Invalid Visibility=Default Availability=Available
// [8] quint64 qFloatDistance(double, double)
func QFloatDistance_1(a float64, b float64) uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z14qFloatDistancedd", ffiqt.FFI_TYPE_POINTER, a, b)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qobjectdefs.h:261
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qFlagLocation(const char *)
func QFlagLocation(method string) string {
	var convArg0 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z13qFlagLocationPKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:108
// index:0
// Invalid Visibility=Default Availability=Available
// [2] quint16 qChecksum(const char *, uint)
func QChecksum(s string, len uint) uint16 {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z9qChecksumPKcj", ffiqt.FFI_TYPE_POINTER, convArg0, len)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:109
// index:1
// Invalid Visibility=Default Availability=Available
// [2] quint16 qChecksum(const char *, uint, Qt::ChecksumType)
func QChecksum_1(s string, len uint, standard int) uint16 {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_Z9qChecksumPKcjN2Qt12ChecksumTypeE", ffiqt.FFI_TYPE_POINTER, convArg0, len, standard)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:683
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QByteArray qCompress(const uchar *, int, int)
func QCompress(data unsafe.Pointer /*666*/, nbytes int, compressionLevel int) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z9qCompressPKhii", ffiqt.FFI_TYPE_POINTER, &data, nbytes, compressionLevel)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:684
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QByteArray qUncompress(const uchar *, int)
func QUncompress(data unsafe.Pointer /*666*/, nbytes int) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z11qUncompressPKhi", ffiqt.FFI_TYPE_POINTER, &data, nbytes)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:685
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] QByteArray qCompress(const QByteArray &, int)
func QCompress_1(data *QByteArray, compressionLevel int) *QByteArray /*123*/ {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z9qCompressRK10QByteArrayi", ffiqt.FFI_TYPE_POINTER, convArg0, compressionLevel)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:687
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] QByteArray qUncompress(const QByteArray &)
func QUncompress_1(data *QByteArray) *QByteArray /*123*/ {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z11qUncompressRK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qalgorithms.h:683
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint32)
func QPopulationCount(v uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z16qPopulationCountj", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:696
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint8)
func QPopulationCount_1(v byte) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z16qPopulationCounth", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:706
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint16)
func QPopulationCount_2(v uint16) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z16qPopulationCountt", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:717
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint64)
func QPopulationCount_3(v uint64) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z16qPopulationCounty", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:732
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(unsigned long)
func QPopulationCount_4(v uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z16qPopulationCountm", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:742
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint32)
func QCountTrailingZeroBits(v uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsj", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:760
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint8)
func QCountTrailingZeroBits_1(v byte) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsh", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:775
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint16)
func QCountTrailingZeroBits_2(v uint16) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z22qCountTrailingZeroBitst", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:791
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint64)
func QCountTrailingZeroBits_3(v uint64) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsy", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:802
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(unsigned long)
func QCountTrailingZeroBits_4(v uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsm", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:807
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint32)
func QCountLeadingZeroBits(v uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsj", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:822
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint8)
func QCountLeadingZeroBits_1(v byte) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsh", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:834
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint16)
func QCountLeadingZeroBits_2(v uint16) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z21qCountLeadingZeroBitst", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:847
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint64)
func QCountLeadingZeroBits_3(v uint64) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsy", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:862
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(unsigned long)
func QCountLeadingZeroBits_4(v uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsm", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:67
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qGlobalQHashSeed()
func QGlobalQHashSeed() int {
	rv, err := ffiqt.InvokeQtFunc6("_Z16qGlobalQHashSeedv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qhashfunctions.h:68
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qSetGlobalQHashSeed(int)
func QSetGlobalQHashSeed(newSeed int) {
	rv, err := ffiqt.InvokeQtFunc6("_Z19qSetGlobalQHashSeedi", ffiqt.FFI_TYPE_POINTER, newSeed)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhashfunctions.h:70
// index:0
// Invalid Visibility=Default Availability=Available
// [4] uint qHashBits(const void *, size_t, uint)
func QHashBits(p unsafe.Pointer /*666*/, size uint, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z9qHashBitsPKvmj", ffiqt.FFI_TYPE_POINTER, p, size, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:72
// index:7
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(char, uint)
func QHash_7(key byte, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashcj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:73
// index:8
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(uchar, uint)
func QHash_8(key byte, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashhj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:74
// index:9
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(signed char, uint)
func QHash_9(key byte, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashaj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:75
// index:10
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(ushort, uint)
func QHash_10(key uint16, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashtj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:76
// index:11
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(short, uint)
func QHash_11(key int16, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashsj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:77
// index:12
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(uint, uint)
func QHash_12(key uint, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashjj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:78
// index:13
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(int, uint)
func QHash_13(key int, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashij", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:79
// index:14
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(ulong, uint)
func QHash_14(key uint, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashmj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:85
// index:15
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(long, uint)
func QHash_15(key int, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashlj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:86
// index:16
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(quint64, uint)
func QHash_16(key uint64, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashyj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:90
// index:17
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(qint64, uint)
func QHash_17(key int64, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashxj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:91
// index:18
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(float, uint)
func QHash_18(key float32, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashfj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:92
// index:19
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(double, uint)
func QHash_19(key float64, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashdj", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:94
// index:20
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(long double, uint)
func QHash_20(key float64, seed uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashej", ffiqt.FFI_TYPE_POINTER, key, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:96
// index:21
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QChar, uint)
func QHash_21(key *QChar /*123*/, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHash5QCharj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:97
// index:22
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QByteArray &, uint)
func QHash_22(key *QByteArray, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK10QByteArrayj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:99
// index:23
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QString &, uint)
func QHash_23(key *QString, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK7QStringj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:100
// index:24
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QStringRef &, uint)
func QHash_24(key *QStringRef, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK10QStringRefj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:102
// index:25
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(QStringView, uint)
func QHash_25(key *QStringView /*123*/, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHash11QStringViewj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:103
// index:26
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QBitArray &, uint)
func QHash_26(key *QBitArray, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK9QBitArrayj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:104
// index:27
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(QLatin1String, uint)
func QHash_27(key *QLatin1String /*123*/, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHash13QLatin1Stringj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qregexp.h:56
// index:28
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QRegExp &, uint)
func QHash_28(key *QRegExp, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK7QRegExpj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qobject.h:93
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QObject * qt_qFindChild_helper(const QObject *, const QString &, const struct QMetaObject &, Qt::FindChildOptions)
func Qt_qFindChild_helper(parent *QObject /*777 const QObject **/, name *QString, mo *QMetaObject, options int) *QObject /*777 QObject **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = name.GetCthis()
	var convArg2 = mo.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z20qt_qFindChild_helperPK7QObjectRK7QStringRK11QMetaObject6QFlagsIN2Qt15FindChildOptionEE", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:102
// index:29
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QPersistentModelIndex &, uint)
func QHash_29(index *QPersistentModelIndex, seed uint) uint {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK21QPersistentModelIndexj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:149
// index:30
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QPersistentModelIndex &, uint)
func QHash_30(index *QPersistentModelIndex, seed uint) uint {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK21QPersistentModelIndexj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:437
// index:31
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QModelIndex &)
func QHash_31(index *QModelIndex) uint {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:62
// index:32
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QLocale &, uint)
func QHash_32(key *QLocale, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK7QLocalej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qcoreapplication.h:258
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qAddPreRoutine(QtStartUpFunction)
func QAddPreRoutine(arg0 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_Z14qAddPreRoutinePFvvE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:259
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qAddPostRoutine(QtCleanUpFunction)
func QAddPostRoutine(arg0 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_Z15qAddPostRoutinePFvvE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:260
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qRemovePostRoutine(QtCleanUpFunction)
func QRemovePostRoutine(arg0 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_Z18qRemovePostRoutinePFvvE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:261
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qAppName()
func QAppName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z8qAppNamev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:412
// index:33
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QDateTime &, uint)
func QHash_33(key *QDateTime, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK9QDateTimej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:413
// index:34
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QDate &, uint)
func QHash_34(key *QDate, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK5QDatej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:414
// index:35
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QTime &, uint)
func QHash_35(key *QTime, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK5QTimej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qtextstream.h:267
// index:0
// Invalid inline Visibility=Default Availability=Available
// [40] QTextStreamManipulator qSetFieldWidth(int)
func QSetFieldWidth(width int) *QTextStreamManipulator /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z14qSetFieldWidthi", ffiqt.FFI_TYPE_POINTER, width)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextStreamManipulatorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:273
// index:0
// Invalid inline Visibility=Default Availability=Available
// [40] QTextStreamManipulator qSetPadChar(QChar)
func QSetPadChar(ch *QChar /*123*/) *QTextStreamManipulator /*123*/ {
	var convArg0 = ch.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z11qSetPadChar5QChar", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextStreamManipulatorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:279
// index:0
// Invalid inline Visibility=Default Availability=Available
// [40] QTextStreamManipulator qSetRealNumberPrecision(int)
func QSetRealNumberPrecision(precision int) *QTextStreamManipulator /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_Z23qSetRealNumberPrecisioni", ffiqt.FFI_TYPE_POINTER, precision)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextStreamManipulatorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:227
// index:36
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QItemSelectionRange &)
func QHash_36(arg0 *QItemSelectionRange) uint {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK19QItemSelectionRange", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qversionnumber.h:54
// index:37
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QVersionNumber &, uint)
func QHash_37(key *QVersionNumber, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK14QVersionNumberj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:68
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qCeil(qreal)
func QCeil(v float64) int {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qCeild", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmath.h:74
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qFloor(qreal)
func QFloor(v float64) int {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qFloord", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmath.h:80
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qFabs(qreal)
func QFabs(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qFabsd", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:86
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qSin(qreal)
func QSin(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z4qSind", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:92
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qCos(qreal)
func QCos(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z4qCosd", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:98
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qTan(qreal)
func QTan(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z4qTand", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:104
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAcos(qreal)
func QAcos(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qAcosd", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:110
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAsin(qreal)
func QAsin(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qAsind", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:116
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAtan(qreal)
func QAtan(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qAtand", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:122
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAtan2(qreal, qreal)
func QAtan2(y float64, x float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z6qAtan2dd", ffiqt.FFI_TYPE_POINTER, y, x)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:128
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qSqrt(qreal)
func QSqrt(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z5qSqrtd", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:134
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qLn(qreal)
func QLn(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z3qLnd", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:140
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qExp(qreal)
func QExp(v float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z4qExpd", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:146
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qPow(qreal, qreal)
func QPow(x float64, y float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z4qPowdd", ffiqt.FFI_TYPE_POINTER, x, y)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:206
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qFastSin(qreal)
func QFastSin(x float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z8qFastSind", ffiqt.FFI_TYPE_POINTER, x)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:216
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qFastCos(qreal)
func QFastCos(x float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z8qFastCosd", ffiqt.FFI_TYPE_POINTER, x)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:226
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] float qDegreesToRadians(float)
func QDegreesToRadians(degrees float32) float32 {
	rv, err := ffiqt.InvokeQtFunc6("_Z17qDegreesToRadiansf", ffiqt.FFI_TYPE_POINTER, degrees)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qmath.h:231
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] double qDegreesToRadians(double)
func QDegreesToRadians_1(degrees float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z17qDegreesToRadiansd", ffiqt.FFI_TYPE_POINTER, degrees)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:236
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] float qRadiansToDegrees(float)
func QRadiansToDegrees(radians float32) float32 {
	rv, err := ffiqt.InvokeQtFunc6("_Z17qRadiansToDegreesf", ffiqt.FFI_TYPE_POINTER, radians)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qmath.h:241
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] double qRadiansToDegrees(double)
func QRadiansToDegrees_1(radians float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z17qRadiansToDegreesd", ffiqt.FFI_TYPE_POINTER, radians)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:247
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] quint32 qNextPowerOfTwo(quint32)
func QNextPowerOfTwo(v uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z15qNextPowerOfTwoj", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:264
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] quint64 qNextPowerOfTwo(quint64)
func QNextPowerOfTwo_1(v uint64) uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z15qNextPowerOfTwoy", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:282
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] quint32 qNextPowerOfTwo(qint32)
func QNextPowerOfTwo_2(v int) uint {
	rv, err := ffiqt.InvokeQtFunc6("_Z15qNextPowerOfTwoi", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:287
// index:3
// Invalid inline Visibility=Default Availability=Available
// [8] quint64 qNextPowerOfTwo(qint64)
func QNextPowerOfTwo_3(v int64) uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_Z15qNextPowerOfTwox", ffiqt.FFI_TYPE_POINTER, v)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qmimetype.h:58
// index:38
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QMimeType &, uint)
func QHash_38(key *QMimeType, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK9QMimeTypej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qplugin.h:78
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qRegisterStaticPluginFunction(struct QStaticPlugin)
func QRegisterStaticPluginFunction(staticPlugin *QStaticPlugin /*123*/) {
	var convArg0 = staticPlugin.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z29qRegisterStaticPluginFunction13QStaticPlugin", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:62
// index:39
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QRegularExpression &, uint)
func QHash_39(key *QRegularExpression, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK18QRegularExpressionj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qurl.h:122
// index:40
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QUrl &, uint)
func QHash_40(url *QUrl, seed uint) uint {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK4QUrlj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qurlquery.h:53
// index:41
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QUrlQuery &, uint)
func QHash_41(key *QUrlQuery, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK9QUrlQueryj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/quuid.h:235
// index:42
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QUuid &, uint)
func QHash_42(uuid *QUuid, seed uint) uint {
	var convArg0 = uuid.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK5QUuidj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /home/me/oss/qt.gen/headers/QtCore/extra_export.h:6
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qRegisterResourceData(int, const unsigned char *, const unsigned char *, const unsigned char *)
func QRegisterResourceData(arg0 int, arg1 unsafe.Pointer /*666*/, arg2 unsafe.Pointer /*666*/, arg3 unsafe.Pointer /*666*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z21qRegisterResourceDataiPKhS0_S0_", ffiqt.FFI_TYPE_POINTER, arg0, arg1, arg2, arg3)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

// /home/me/oss/qt.gen/headers/QtCore/extra_export.h:7
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qUnregisterResourceData(int, const unsigned char *, const unsigned char *, const unsigned char *)
func QUnregisterResourceData(arg0 int, arg1 unsafe.Pointer /*666*/, arg2 unsafe.Pointer /*666*/, arg3 unsafe.Pointer /*666*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_Z23qUnregisterResourceDataiPKhS0_S0_", ffiqt.FFI_TYPE_POINTER, arg0, arg1, arg2, arg3)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}

//  body block end
