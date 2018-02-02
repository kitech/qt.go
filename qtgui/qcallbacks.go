//  package block begin
package qtgui

/*
#include <stdint.h>
#include <stdbool.h>
//  package block end

//  extern block begin
extern void callback_ZN12QPaintDeviceC1Ev(void* fnptr );
extern void callback_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE(void* fnptr , int metric);
extern void callback_ZNK12QPaintDevice11initPainterEP8QPainter(void* fnptr , void* painter);
extern void callback_ZNK12QPaintDevice10redirectedEP6QPoint(void* fnptr , void* offset);
extern void callback_ZNK12QPaintDevice13sharedPainterEv(void* fnptr );
extern void callback_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE(void* fnptr , int metric);
extern void callback_ZNK6QImage15mirrored_helperEbb(void* fnptr , bool horizontal, bool vertical);
extern void callback_ZNK6QImage17rgbSwapped_helperEv(void* fnptr );
extern void callback_ZN6QImage16mirrored_inplaceEbb(void* fnptr , bool horizontal, bool vertical);
extern void callback_ZN6QImage18rgbSwapped_inplaceEv(void* fnptr );
extern void callback_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE(void* fnptr , int format, int flags);
extern void callback_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE(void* fnptr , int format, int flags);
extern void callback_ZNK6QImage12smoothScaledEii(void* fnptr , int w, int h);
extern void callback_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE(void* fnptr , int arg0);
extern void callback_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE(void* fnptr , void* image, int flags);
extern void callback_ZN15QTextCharFormatC1ERK11QTextFormat(void* fnptr , void* fmt);
extern void callback_ZN16QTextBlockFormatC1ERK11QTextFormat(void* fnptr , void* fmt);
extern void callback_ZN15QTextListFormatC1ERK11QTextFormat(void* fnptr , void* fmt);
extern void callback_ZN16QTextImageFormatC1ERK11QTextFormat(void* fnptr , void* format);
extern void callback_ZN16QTextFrameFormatC1ERK11QTextFormat(void* fnptr , void* fmt);
extern void callback_ZN16QTextTableFormatC1ERK11QTextFormat(void* fnptr , void* fmt);
extern void callback_ZN20QTextTableCellFormatC1ERK11QTextFormat(void* fnptr , void* fmt);
extern void callback_ZN13QTextDocument12createObjectERK11QTextFormat(void* fnptr , void* f);
extern void callback_ZN13QTextDocument12loadResourceEiRK4QUrl(void* fnptr , int type_, void* name);
extern void callback_ZN27QAbstractTextDocumentLayout15documentChangedEiii(void* fnptr , int from, int charsRemoved, int charsAdded);
extern void callback_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat(void* fnptr , void* item, int posInDocument, void* format);
extern void callback_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat(void* fnptr , void* item, int posInDocument, void* format);
extern void callback_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat(void* fnptr , void* painter, void* rect, void* object, int posInDocument, void* format);
extern void callback_ZN27QAbstractTextDocumentLayout11formatIndexEi(void* fnptr , int pos);
extern void callback_ZN27QAbstractTextDocumentLayout6formatEi(void* fnptr , int pos);
extern void callback_ZN20QAccessibleInterfaceD1Ev(void* fnptr );
extern void callback_ZN8QSurfaceC1ENS_12SurfaceClassE(void* fnptr , int type_);
extern void callback_ZN7QWindow11exposeEventEP12QExposeEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow9moveEventEP10QMoveEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow12focusInEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow13focusOutEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow9hideEventEP10QHideEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow15keyReleaseEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow10wheelEventEP11QWheelEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow10touchEventEP11QTouchEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow11tabletEventEP12QTabletEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWindow11nativeEventERK10QByteArrayPvPl(void* fnptr , void* eventType, void* message, void* result);
extern void callback_ZN15QGuiApplication5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN24QAbstractOpenGLFunctionsC1Ev(void* fnptr );
extern void callback_ZNK24QAbstractOpenGLFunctions13isInitializedEv(void* fnptr );
extern void callback_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE(void* fnptr , int metric);
extern void callback_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent(void* fnptr , void* arg0);
extern void callback_ZN18QPaintDeviceWindow5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZNK17QPagedPaintDevice16devicePageLayoutEv(void* fnptr );
extern void callback_ZN17QPagedPaintDevice16devicePageLayoutEv(void* fnptr );
extern void callback_ZNK10QPdfWriter11paintEngineEv(void* fnptr );
extern void callback_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE(void* fnptr , int id);
extern void callback_ZNK8QPicture6metricEN12QPaintDevice17PaintDeviceMetricE(void* fnptr , int m);
extern void callback_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE(void* fnptr , int metric);
extern void callback_ZNK13QRasterWindow10redirectedEP6QPoint(void* fnptr , void* arg0);
extern void callback_ZN13QStandardItem15emitDataChangedEv(void* fnptr );
extern void callback_ZN11QTextObjectC1EP13QTextDocument(void* fnptr , void* doc);
extern void callback_ZN11QTextObjectD1Ev(void* fnptr );
extern void callback_ZN11QTextObject9setFormatERK11QTextFormat(void* fnptr , void* format);
extern void callback_ZN15QTextBlockGroupC1EP13QTextDocument(void* fnptr , void* doc);
extern void callback_ZN15QTextBlockGroupD1Ev(void* fnptr );
extern void callback_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock(void* fnptr , void* block);
extern void callback_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock(void* fnptr , void* block);
extern void callback_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock(void* fnptr , void* block);
extern void callback_ZN18QSyntaxHighlighter14highlightBlockERK7QString(void* fnptr , void* text);
extern void callback_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat(void* fnptr , int start, int count, void* format);
extern void callback_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor(void* fnptr , int start, int count, void* color);
extern void callback_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont(void* fnptr , int start, int count, void* font);
extern void callback_ZNK18QSyntaxHighlighter6formatEi(void* fnptr , int pos);
extern void callback_ZNK18QSyntaxHighlighter18previousBlockStateEv(void* fnptr );
extern void callback_ZNK18QSyntaxHighlighter17currentBlockStateEv(void* fnptr );
extern void callback_ZN18QSyntaxHighlighter20setCurrentBlockStateEi(void* fnptr , int newState);
extern void callback_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData(void* fnptr , void* data);
extern void callback_ZNK18QSyntaxHighlighter20currentBlockUserDataEv(void* fnptr );
extern void callback_ZNK18QSyntaxHighlighter12currentBlockEv(void* fnptr );
//  extern block end

//  header block begin
*/
import "C"
import "unsafe"
import "qt.go/cffiqt"
import "gopp"

// import "log"
//  header block end

//  body block begin
// void QPaintDevice()
//export callback_ZN12QPaintDeviceC1Ev
func callback_ZN12QPaintDeviceC1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPaintDevice.QPaintDevice")
	rvx := ffiqt.CallbackAllInherits(cthis, "QPaintDevice")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QPaintDeviceC1Ev", C.callback_ZN12QPaintDeviceC1Ev /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE
func callback_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE(cthis unsafe.Pointer, metric C.int) {
	// log.Println(cthis, "QPaintDevice.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", metric)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE", C.callback_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE /*nil*/)
}

// void initPainter(class QPainter *)
//export callback_ZNK12QPaintDevice11initPainterEP8QPainter
func callback_ZNK12QPaintDevice11initPainterEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPaintDevice.initPainter")
	rvx := ffiqt.CallbackAllInherits(cthis, "initPainter", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QPaintDevice11initPainterEP8QPainter", C.callback_ZNK12QPaintDevice11initPainterEP8QPainter /*nil*/)
}

// QPaintDevice * redirected(class QPoint *)
//export callback_ZNK12QPaintDevice10redirectedEP6QPoint
func callback_ZNK12QPaintDevice10redirectedEP6QPoint(cthis unsafe.Pointer, offset unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPaintDevice.redirected")
	rvx := ffiqt.CallbackAllInherits(cthis, "redirected", offset)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QPaintDevice10redirectedEP6QPoint", C.callback_ZNK12QPaintDevice10redirectedEP6QPoint /*nil*/)
}

// QPainter * sharedPainter()
//export callback_ZNK12QPaintDevice13sharedPainterEv
func callback_ZNK12QPaintDevice13sharedPainterEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPaintDevice.sharedPainter")
	rvx := ffiqt.CallbackAllInherits(cthis, "sharedPainter")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QPaintDevice13sharedPainterEv", C.callback_ZNK12QPaintDevice13sharedPainterEv /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE
func callback_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE(cthis unsafe.Pointer, metric C.int) {
	// log.Println(cthis, "QImage.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", metric)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE", C.callback_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE /*nil*/)
}

// QImage mirrored_helper(_Bool, _Bool)
//export callback_ZNK6QImage15mirrored_helperEbb
func callback_ZNK6QImage15mirrored_helperEbb(cthis unsafe.Pointer, horizontal C.bool, vertical C.bool) {
	// log.Println(cthis, "QImage.mirrored_helper")
	rvx := ffiqt.CallbackAllInherits(cthis, "mirrored_helper", horizontal, vertical)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK6QImage15mirrored_helperEbb", C.callback_ZNK6QImage15mirrored_helperEbb /*nil*/)
}

// QImage rgbSwapped_helper()
//export callback_ZNK6QImage17rgbSwapped_helperEv
func callback_ZNK6QImage17rgbSwapped_helperEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QImage.rgbSwapped_helper")
	rvx := ffiqt.CallbackAllInherits(cthis, "rgbSwapped_helper")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK6QImage17rgbSwapped_helperEv", C.callback_ZNK6QImage17rgbSwapped_helperEv /*nil*/)
}

// void mirrored_inplace(_Bool, _Bool)
//export callback_ZN6QImage16mirrored_inplaceEbb
func callback_ZN6QImage16mirrored_inplaceEbb(cthis unsafe.Pointer, horizontal C.bool, vertical C.bool) {
	// log.Println(cthis, "QImage.mirrored_inplace")
	rvx := ffiqt.CallbackAllInherits(cthis, "mirrored_inplace", horizontal, vertical)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QImage16mirrored_inplaceEbb", C.callback_ZN6QImage16mirrored_inplaceEbb /*nil*/)
}

// void rgbSwapped_inplace()
//export callback_ZN6QImage18rgbSwapped_inplaceEv
func callback_ZN6QImage18rgbSwapped_inplaceEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QImage.rgbSwapped_inplace")
	rvx := ffiqt.CallbackAllInherits(cthis, "rgbSwapped_inplace")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QImage18rgbSwapped_inplaceEv", C.callback_ZN6QImage18rgbSwapped_inplaceEv /*nil*/)
}

// QImage convertToFormat_helper(enum QImage::Format, Qt::ImageConversionFlags)
//export callback_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE
func callback_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE(cthis unsafe.Pointer, format C.int, flags C.int) {
	// log.Println(cthis, "QImage.convertToFormat_helper")
	rvx := ffiqt.CallbackAllInherits(cthis, "convertToFormat_helper", format, flags)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", C.callback_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE /*nil*/)
}

// bool convertToFormat_inplace(enum QImage::Format, Qt::ImageConversionFlags)
//export callback_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE
func callback_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE(cthis unsafe.Pointer, format C.int, flags C.int) {
	// log.Println(cthis, "QImage.convertToFormat_inplace")
	rvx := ffiqt.CallbackAllInherits(cthis, "convertToFormat_inplace", format, flags)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", C.callback_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE /*nil*/)
}

// QImage smoothScaled(int, int)
//export callback_ZNK6QImage12smoothScaledEii
func callback_ZNK6QImage12smoothScaledEii(cthis unsafe.Pointer, w C.int, h C.int) {
	// log.Println(cthis, "QImage.smoothScaled")
	rvx := ffiqt.CallbackAllInherits(cthis, "smoothScaled", w, h)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK6QImage12smoothScaledEii", C.callback_ZNK6QImage12smoothScaledEii /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE
func callback_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QPixmap.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE", C.callback_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE /*nil*/)
}

// QPixmap fromImageInPlace(class QImage &, Qt::ImageConversionFlags)
//export callback_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE
func callback_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE(cthis unsafe.Pointer, image unsafe.Pointer /*555*/, flags C.int) {
	// log.Println(cthis, "QPixmap.fromImageInPlace")
	rvx := ffiqt.CallbackAllInherits(cthis, "fromImageInPlace", image, flags)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE", C.callback_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE /*nil*/)
}

// void QTextCharFormat(const class QTextFormat &)
//export callback_ZN15QTextCharFormatC1ERK11QTextFormat
func callback_ZN15QTextCharFormatC1ERK11QTextFormat(cthis unsafe.Pointer, fmt unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextCharFormat.QTextCharFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextCharFormat", fmt)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTextCharFormatC1ERK11QTextFormat", C.callback_ZN15QTextCharFormatC1ERK11QTextFormat /*nil*/)
}

// void QTextBlockFormat(const class QTextFormat &)
//export callback_ZN16QTextBlockFormatC1ERK11QTextFormat
func callback_ZN16QTextBlockFormatC1ERK11QTextFormat(cthis unsafe.Pointer, fmt unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextBlockFormat.QTextBlockFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextBlockFormat", fmt)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QTextBlockFormatC1ERK11QTextFormat", C.callback_ZN16QTextBlockFormatC1ERK11QTextFormat /*nil*/)
}

// void QTextListFormat(const class QTextFormat &)
//export callback_ZN15QTextListFormatC1ERK11QTextFormat
func callback_ZN15QTextListFormatC1ERK11QTextFormat(cthis unsafe.Pointer, fmt unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextListFormat.QTextListFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextListFormat", fmt)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTextListFormatC1ERK11QTextFormat", C.callback_ZN15QTextListFormatC1ERK11QTextFormat /*nil*/)
}

// void QTextImageFormat(const class QTextFormat &)
//export callback_ZN16QTextImageFormatC1ERK11QTextFormat
func callback_ZN16QTextImageFormatC1ERK11QTextFormat(cthis unsafe.Pointer, format unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextImageFormat.QTextImageFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextImageFormat", format)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QTextImageFormatC1ERK11QTextFormat", C.callback_ZN16QTextImageFormatC1ERK11QTextFormat /*nil*/)
}

// void QTextFrameFormat(const class QTextFormat &)
//export callback_ZN16QTextFrameFormatC1ERK11QTextFormat
func callback_ZN16QTextFrameFormatC1ERK11QTextFormat(cthis unsafe.Pointer, fmt unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextFrameFormat.QTextFrameFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextFrameFormat", fmt)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QTextFrameFormatC1ERK11QTextFormat", C.callback_ZN16QTextFrameFormatC1ERK11QTextFormat /*nil*/)
}

// void QTextTableFormat(const class QTextFormat &)
//export callback_ZN16QTextTableFormatC1ERK11QTextFormat
func callback_ZN16QTextTableFormatC1ERK11QTextFormat(cthis unsafe.Pointer, fmt unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextTableFormat.QTextTableFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextTableFormat", fmt)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QTextTableFormatC1ERK11QTextFormat", C.callback_ZN16QTextTableFormatC1ERK11QTextFormat /*nil*/)
}

// void QTextTableCellFormat(const class QTextFormat &)
//export callback_ZN20QTextTableCellFormatC1ERK11QTextFormat
func callback_ZN20QTextTableCellFormatC1ERK11QTextFormat(cthis unsafe.Pointer, fmt unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextTableCellFormat.QTextTableCellFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextTableCellFormat", fmt)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QTextTableCellFormatC1ERK11QTextFormat", C.callback_ZN20QTextTableCellFormatC1ERK11QTextFormat /*nil*/)
}

// QTextObject * createObject(const class QTextFormat &)
//export callback_ZN13QTextDocument12createObjectERK11QTextFormat
func callback_ZN13QTextDocument12createObjectERK11QTextFormat(cthis unsafe.Pointer, f unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextDocument.createObject")
	rvx := ffiqt.CallbackAllInherits(cthis, "createObject", f)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QTextDocument12createObjectERK11QTextFormat", C.callback_ZN13QTextDocument12createObjectERK11QTextFormat /*nil*/)
}

// QVariant loadResource(int, const class QUrl &)
//export callback_ZN13QTextDocument12loadResourceEiRK4QUrl
func callback_ZN13QTextDocument12loadResourceEiRK4QUrl(cthis unsafe.Pointer, type_ C.int, name unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextDocument.loadResource")
	rvx := ffiqt.CallbackAllInherits(cthis, "loadResource", type_, name)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QTextDocument12loadResourceEiRK4QUrl", C.callback_ZN13QTextDocument12loadResourceEiRK4QUrl /*nil*/)
}

// void documentChanged(int, int, int)
//export callback_ZN27QAbstractTextDocumentLayout15documentChangedEiii
func callback_ZN27QAbstractTextDocumentLayout15documentChangedEiii(cthis unsafe.Pointer, from C.int, charsRemoved C.int, charsAdded C.int) {
	// log.Println(cthis, "QAbstractTextDocumentLayout.documentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "documentChanged", from, charsRemoved, charsAdded)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN27QAbstractTextDocumentLayout15documentChangedEiii", C.callback_ZN27QAbstractTextDocumentLayout15documentChangedEiii /*nil*/)
}

// void resizeInlineObject(class QTextInlineObject, int, const class QTextFormat &)
//export callback_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat
func callback_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat(cthis unsafe.Pointer, item unsafe.Pointer /*444*/, posInDocument C.int, format unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractTextDocumentLayout.resizeInlineObject")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeInlineObject", item, posInDocument, format)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat", C.callback_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat /*nil*/)
}

// void positionInlineObject(class QTextInlineObject, int, const class QTextFormat &)
//export callback_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat
func callback_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat(cthis unsafe.Pointer, item unsafe.Pointer /*444*/, posInDocument C.int, format unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractTextDocumentLayout.positionInlineObject")
	rvx := ffiqt.CallbackAllInherits(cthis, "positionInlineObject", item, posInDocument, format)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat", C.callback_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat /*nil*/)
}

// void drawInlineObject(class QPainter *, const class QRectF &, class QTextInlineObject, int, const class QTextFormat &)
//export callback_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat
func callback_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/, object unsafe.Pointer /*444*/, posInDocument C.int, format unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractTextDocumentLayout.drawInlineObject")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawInlineObject", painter, rect, object, posInDocument, format)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat", C.callback_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat /*nil*/)
}

// int formatIndex(int)
//export callback_ZN27QAbstractTextDocumentLayout11formatIndexEi
func callback_ZN27QAbstractTextDocumentLayout11formatIndexEi(cthis unsafe.Pointer, pos C.int) {
	// log.Println(cthis, "QAbstractTextDocumentLayout.formatIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "formatIndex", pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN27QAbstractTextDocumentLayout11formatIndexEi", C.callback_ZN27QAbstractTextDocumentLayout11formatIndexEi /*nil*/)
}

// QTextCharFormat format(int)
//export callback_ZN27QAbstractTextDocumentLayout6formatEi
func callback_ZN27QAbstractTextDocumentLayout6formatEi(cthis unsafe.Pointer, pos C.int) {
	// log.Println(cthis, "QAbstractTextDocumentLayout.format")
	rvx := ffiqt.CallbackAllInherits(cthis, "format", pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN27QAbstractTextDocumentLayout6formatEi", C.callback_ZN27QAbstractTextDocumentLayout6formatEi /*nil*/)
}

// void ~QAccessibleInterface()
//export callback_ZN20QAccessibleInterfaceD1Ev
func callback_ZN20QAccessibleInterfaceD1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAccessibleInterface.~QAccessibleInterface")
	rvx := ffiqt.CallbackAllInherits(cthis, "~QAccessibleInterface")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QAccessibleInterfaceD1Ev", C.callback_ZN20QAccessibleInterfaceD1Ev /*nil*/)
}

// void QSurface(enum QSurface::SurfaceClass)
//export callback_ZN8QSurfaceC1ENS_12SurfaceClassE
func callback_ZN8QSurfaceC1ENS_12SurfaceClassE(cthis unsafe.Pointer, type_ C.int) {
	// log.Println(cthis, "QSurface.QSurface")
	rvx := ffiqt.CallbackAllInherits(cthis, "QSurface", type_)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QSurfaceC1ENS_12SurfaceClassE", C.callback_ZN8QSurfaceC1ENS_12SurfaceClassE /*nil*/)
}

// void exposeEvent(class QExposeEvent *)
//export callback_ZN7QWindow11exposeEventEP12QExposeEvent
func callback_ZN7QWindow11exposeEventEP12QExposeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.exposeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "exposeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow11exposeEventEP12QExposeEvent", C.callback_ZN7QWindow11exposeEventEP12QExposeEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN7QWindow11resizeEventEP12QResizeEvent
func callback_ZN7QWindow11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow11resizeEventEP12QResizeEvent", C.callback_ZN7QWindow11resizeEventEP12QResizeEvent /*nil*/)
}

// void moveEvent(class QMoveEvent *)
//export callback_ZN7QWindow9moveEventEP10QMoveEvent
func callback_ZN7QWindow9moveEventEP10QMoveEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.moveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow9moveEventEP10QMoveEvent", C.callback_ZN7QWindow9moveEventEP10QMoveEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN7QWindow12focusInEventEP11QFocusEvent
func callback_ZN7QWindow12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow12focusInEventEP11QFocusEvent", C.callback_ZN7QWindow12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN7QWindow13focusOutEventEP11QFocusEvent
func callback_ZN7QWindow13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow13focusOutEventEP11QFocusEvent", C.callback_ZN7QWindow13focusOutEventEP11QFocusEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN7QWindow9showEventEP10QShowEvent
func callback_ZN7QWindow9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow9showEventEP10QShowEvent", C.callback_ZN7QWindow9showEventEP10QShowEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN7QWindow9hideEventEP10QHideEvent
func callback_ZN7QWindow9hideEventEP10QHideEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow9hideEventEP10QHideEvent", C.callback_ZN7QWindow9hideEventEP10QHideEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN7QWindow5eventEP6QEvent
func callback_ZN7QWindow5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow5eventEP6QEvent", C.callback_ZN7QWindow5eventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN7QWindow13keyPressEventEP9QKeyEvent
func callback_ZN7QWindow13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow13keyPressEventEP9QKeyEvent", C.callback_ZN7QWindow13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN7QWindow15keyReleaseEventEP9QKeyEvent
func callback_ZN7QWindow15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow15keyReleaseEventEP9QKeyEvent", C.callback_ZN7QWindow15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN7QWindow15mousePressEventEP11QMouseEvent
func callback_ZN7QWindow15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow15mousePressEventEP11QMouseEvent", C.callback_ZN7QWindow15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN7QWindow17mouseReleaseEventEP11QMouseEvent
func callback_ZN7QWindow17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow17mouseReleaseEventEP11QMouseEvent", C.callback_ZN7QWindow17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN7QWindow14mouseMoveEventEP11QMouseEvent
func callback_ZN7QWindow14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow14mouseMoveEventEP11QMouseEvent", C.callback_ZN7QWindow14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN7QWindow10wheelEventEP11QWheelEvent
func callback_ZN7QWindow10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow10wheelEventEP11QWheelEvent", C.callback_ZN7QWindow10wheelEventEP11QWheelEvent /*nil*/)
}

// void touchEvent(class QTouchEvent *)
//export callback_ZN7QWindow10touchEventEP11QTouchEvent
func callback_ZN7QWindow10touchEventEP11QTouchEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.touchEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "touchEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow10touchEventEP11QTouchEvent", C.callback_ZN7QWindow10touchEventEP11QTouchEvent /*nil*/)
}

// void tabletEvent(class QTabletEvent *)
//export callback_ZN7QWindow11tabletEventEP12QTabletEvent
func callback_ZN7QWindow11tabletEventEP12QTabletEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.tabletEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabletEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow11tabletEventEP12QTabletEvent", C.callback_ZN7QWindow11tabletEventEP12QTabletEvent /*nil*/)
}

// bool nativeEvent(const class QByteArray &, void *, long *)
//export callback_ZN7QWindow11nativeEventERK10QByteArrayPvPl
func callback_ZN7QWindow11nativeEventERK10QByteArrayPvPl(cthis unsafe.Pointer, eventType unsafe.Pointer /*555*/, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWindow.nativeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "nativeEvent", eventType, message, result)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWindow11nativeEventERK10QByteArrayPvPl", C.callback_ZN7QWindow11nativeEventERK10QByteArrayPvPl /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QGuiApplication5eventEP6QEvent
func callback_ZN15QGuiApplication5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGuiApplication.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGuiApplication5eventEP6QEvent", C.callback_ZN15QGuiApplication5eventEP6QEvent /*nil*/)
}

// void QAbstractOpenGLFunctions()
//export callback_ZN24QAbstractOpenGLFunctionsC1Ev
func callback_ZN24QAbstractOpenGLFunctionsC1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractOpenGLFunctions.QAbstractOpenGLFunctions")
	rvx := ffiqt.CallbackAllInherits(cthis, "QAbstractOpenGLFunctions")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN24QAbstractOpenGLFunctionsC1Ev", C.callback_ZN24QAbstractOpenGLFunctionsC1Ev /*nil*/)
}

// bool isInitialized()
//export callback_ZNK24QAbstractOpenGLFunctions13isInitializedEv
func callback_ZNK24QAbstractOpenGLFunctions13isInitializedEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractOpenGLFunctions.isInitialized")
	rvx := ffiqt.CallbackAllInherits(cthis, "isInitialized")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK24QAbstractOpenGLFunctions13isInitializedEv", C.callback_ZNK24QAbstractOpenGLFunctions13isInitializedEv /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent
func callback_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPaintDeviceWindow.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent", C.callback_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE
func callback_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE(cthis unsafe.Pointer, metric C.int) {
	// log.Println(cthis, "QPaintDeviceWindow.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", metric)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE", C.callback_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE /*nil*/)
}

// void exposeEvent(class QExposeEvent *)
//export callback_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent
func callback_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPaintDeviceWindow.exposeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "exposeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent", C.callback_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN18QPaintDeviceWindow5eventEP6QEvent
func callback_ZN18QPaintDeviceWindow5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPaintDeviceWindow.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QPaintDeviceWindow5eventEP6QEvent", C.callback_ZN18QPaintDeviceWindow5eventEP6QEvent /*nil*/)
}

// QPageLayout devicePageLayout()
//export callback_ZNK17QPagedPaintDevice16devicePageLayoutEv
func callback_ZNK17QPagedPaintDevice16devicePageLayoutEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPagedPaintDevice.devicePageLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "devicePageLayout")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QPagedPaintDevice16devicePageLayoutEv", C.callback_ZNK17QPagedPaintDevice16devicePageLayoutEv /*nil*/)
}

// QPageLayout & devicePageLayout()
//export callback_ZN17QPagedPaintDevice16devicePageLayoutEv
func callback_ZN17QPagedPaintDevice16devicePageLayoutEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPagedPaintDevice.devicePageLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "devicePageLayout")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QPagedPaintDevice16devicePageLayoutEv", C.callback_ZN17QPagedPaintDevice16devicePageLayoutEv /*nil*/)
}

// QPaintEngine * paintEngine()
//export callback_ZNK10QPdfWriter11paintEngineEv
func callback_ZNK10QPdfWriter11paintEngineEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPdfWriter.paintEngine")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEngine")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QPdfWriter11paintEngineEv", C.callback_ZNK10QPdfWriter11paintEngineEv /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE
func callback_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE(cthis unsafe.Pointer, id C.int) {
	// log.Println(cthis, "QPdfWriter.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", id)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE", C.callback_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK8QPicture6metricEN12QPaintDevice17PaintDeviceMetricE
func callback_ZNK8QPicture6metricEN12QPaintDevice17PaintDeviceMetricE(cthis unsafe.Pointer, m C.int) {
	// log.Println(cthis, "QPicture.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", m)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QPicture6metricEN12QPaintDevice17PaintDeviceMetricE", C.callback_ZNK8QPicture6metricEN12QPaintDevice17PaintDeviceMetricE /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE
func callback_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE(cthis unsafe.Pointer, metric C.int) {
	// log.Println(cthis, "QRasterWindow.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", metric)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE", C.callback_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE /*nil*/)
}

// QPaintDevice * redirected(class QPoint *)
//export callback_ZNK13QRasterWindow10redirectedEP6QPoint
func callback_ZNK13QRasterWindow10redirectedEP6QPoint(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRasterWindow.redirected")
	rvx := ffiqt.CallbackAllInherits(cthis, "redirected", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QRasterWindow10redirectedEP6QPoint", C.callback_ZNK13QRasterWindow10redirectedEP6QPoint /*nil*/)
}

// void emitDataChanged()
//export callback_ZN13QStandardItem15emitDataChangedEv
func callback_ZN13QStandardItem15emitDataChangedEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QStandardItem.emitDataChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "emitDataChanged")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStandardItem15emitDataChangedEv", C.callback_ZN13QStandardItem15emitDataChangedEv /*nil*/)
}

// void QTextObject(class QTextDocument *)
//export callback_ZN11QTextObjectC1EP13QTextDocument
func callback_ZN11QTextObjectC1EP13QTextDocument(cthis unsafe.Pointer, doc unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextObject.QTextObject")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextObject", doc)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QTextObjectC1EP13QTextDocument", C.callback_ZN11QTextObjectC1EP13QTextDocument /*nil*/)
}

// void ~QTextObject()
//export callback_ZN11QTextObjectD1Ev
func callback_ZN11QTextObjectD1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTextObject.~QTextObject")
	rvx := ffiqt.CallbackAllInherits(cthis, "~QTextObject")
	gopp.ErrPrint(nil, rvx)
}
func init() { ffiqt.SetInheritCallback2c("_ZN11QTextObjectD1Ev", C.callback_ZN11QTextObjectD1Ev /*nil*/) }

// void setFormat(const class QTextFormat &)
//export callback_ZN11QTextObject9setFormatERK11QTextFormat
func callback_ZN11QTextObject9setFormatERK11QTextFormat(cthis unsafe.Pointer, format unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextObject.setFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "setFormat", format)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QTextObject9setFormatERK11QTextFormat", C.callback_ZN11QTextObject9setFormatERK11QTextFormat /*nil*/)
}

// void QTextBlockGroup(class QTextDocument *)
//export callback_ZN15QTextBlockGroupC1EP13QTextDocument
func callback_ZN15QTextBlockGroupC1EP13QTextDocument(cthis unsafe.Pointer, doc unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBlockGroup.QTextBlockGroup")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextBlockGroup", doc)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTextBlockGroupC1EP13QTextDocument", C.callback_ZN15QTextBlockGroupC1EP13QTextDocument /*nil*/)
}

// void ~QTextBlockGroup()
//export callback_ZN15QTextBlockGroupD1Ev
func callback_ZN15QTextBlockGroupD1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTextBlockGroup.~QTextBlockGroup")
	rvx := ffiqt.CallbackAllInherits(cthis, "~QTextBlockGroup")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTextBlockGroupD1Ev", C.callback_ZN15QTextBlockGroupD1Ev /*nil*/)
}

// void blockInserted(const class QTextBlock &)
//export callback_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock
func callback_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock(cthis unsafe.Pointer, block unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextBlockGroup.blockInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "blockInserted", block)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock", C.callback_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock /*nil*/)
}

// void blockRemoved(const class QTextBlock &)
//export callback_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock
func callback_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock(cthis unsafe.Pointer, block unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextBlockGroup.blockRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "blockRemoved", block)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock", C.callback_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock /*nil*/)
}

// void blockFormatChanged(const class QTextBlock &)
//export callback_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock
func callback_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock(cthis unsafe.Pointer, block unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextBlockGroup.blockFormatChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "blockFormatChanged", block)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock", C.callback_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock /*nil*/)
}

// void highlightBlock(const class QString &)
//export callback_ZN18QSyntaxHighlighter14highlightBlockERK7QString
func callback_ZN18QSyntaxHighlighter14highlightBlockERK7QString(cthis unsafe.Pointer, text unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSyntaxHighlighter.highlightBlock")
	rvx := ffiqt.CallbackAllInherits(cthis, "highlightBlock", text)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QSyntaxHighlighter14highlightBlockERK7QString", C.callback_ZN18QSyntaxHighlighter14highlightBlockERK7QString /*nil*/)
}

// void setFormat(int, int, const class QTextCharFormat &)
//export callback_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat
func callback_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat(cthis unsafe.Pointer, start C.int, count C.int, format unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSyntaxHighlighter.setFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "setFormat", start, count, format)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat", C.callback_ZN18QSyntaxHighlighter9setFormatEiiRK15QTextCharFormat /*nil*/)
}

// void setFormat(int, int, const class QColor &)
//export callback_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor
func callback_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor(cthis unsafe.Pointer, start C.int, count C.int, color unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSyntaxHighlighter.setFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "setFormat", start, count, color)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor", C.callback_ZN18QSyntaxHighlighter9setFormatEiiRK6QColor /*nil*/)
}

// void setFormat(int, int, const class QFont &)
//export callback_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont
func callback_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont(cthis unsafe.Pointer, start C.int, count C.int, font unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSyntaxHighlighter.setFormat")
	rvx := ffiqt.CallbackAllInherits(cthis, "setFormat", start, count, font)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont", C.callback_ZN18QSyntaxHighlighter9setFormatEiiRK5QFont /*nil*/)
}

// QTextCharFormat format(int)
//export callback_ZNK18QSyntaxHighlighter6formatEi
func callback_ZNK18QSyntaxHighlighter6formatEi(cthis unsafe.Pointer, pos C.int) {
	// log.Println(cthis, "QSyntaxHighlighter.format")
	rvx := ffiqt.CallbackAllInherits(cthis, "format", pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QSyntaxHighlighter6formatEi", C.callback_ZNK18QSyntaxHighlighter6formatEi /*nil*/)
}

// int previousBlockState()
//export callback_ZNK18QSyntaxHighlighter18previousBlockStateEv
func callback_ZNK18QSyntaxHighlighter18previousBlockStateEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QSyntaxHighlighter.previousBlockState")
	rvx := ffiqt.CallbackAllInherits(cthis, "previousBlockState")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QSyntaxHighlighter18previousBlockStateEv", C.callback_ZNK18QSyntaxHighlighter18previousBlockStateEv /*nil*/)
}

// int currentBlockState()
//export callback_ZNK18QSyntaxHighlighter17currentBlockStateEv
func callback_ZNK18QSyntaxHighlighter17currentBlockStateEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QSyntaxHighlighter.currentBlockState")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentBlockState")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QSyntaxHighlighter17currentBlockStateEv", C.callback_ZNK18QSyntaxHighlighter17currentBlockStateEv /*nil*/)
}

// void setCurrentBlockState(int)
//export callback_ZN18QSyntaxHighlighter20setCurrentBlockStateEi
func callback_ZN18QSyntaxHighlighter20setCurrentBlockStateEi(cthis unsafe.Pointer, newState C.int) {
	// log.Println(cthis, "QSyntaxHighlighter.setCurrentBlockState")
	rvx := ffiqt.CallbackAllInherits(cthis, "setCurrentBlockState", newState)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QSyntaxHighlighter20setCurrentBlockStateEi", C.callback_ZN18QSyntaxHighlighter20setCurrentBlockStateEi /*nil*/)
}

// void setCurrentBlockUserData(class QTextBlockUserData *)
//export callback_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData
func callback_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData(cthis unsafe.Pointer, data unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSyntaxHighlighter.setCurrentBlockUserData")
	rvx := ffiqt.CallbackAllInherits(cthis, "setCurrentBlockUserData", data)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData", C.callback_ZN18QSyntaxHighlighter23setCurrentBlockUserDataEP18QTextBlockUserData /*nil*/)
}

// QTextBlockUserData * currentBlockUserData()
//export callback_ZNK18QSyntaxHighlighter20currentBlockUserDataEv
func callback_ZNK18QSyntaxHighlighter20currentBlockUserDataEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QSyntaxHighlighter.currentBlockUserData")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentBlockUserData")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QSyntaxHighlighter20currentBlockUserDataEv", C.callback_ZNK18QSyntaxHighlighter20currentBlockUserDataEv /*nil*/)
}

// QTextBlock currentBlock()
//export callback_ZNK18QSyntaxHighlighter12currentBlockEv
func callback_ZNK18QSyntaxHighlighter12currentBlockEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QSyntaxHighlighter.currentBlock")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentBlock")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QSyntaxHighlighter12currentBlockEv", C.callback_ZNK18QSyntaxHighlighter12currentBlockEv /*nil*/)
}

//  body block end
