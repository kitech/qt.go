//  header block begin
// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QRegularExpressionMatch struct {
	*qtrt.CObject
}

func (this *QRegularExpressionMatch) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qregularexpression.h:178
// index:0
// void QRegularExpressionMatch()
func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QRegularExpressionMatchC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchFromPointer(cthis)
	return gothis
}
func NewQRegularExpressionMatchFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatch {
	return &QRegularExpressionMatch{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qregularexpression.h:179
// index:0
// void ~QRegularExpressionMatch()
func DeleteQRegularExpressionMatch(*QRegularExpressionMatch) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QRegularExpressionMatchD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:187
// index:0
// inline
// void swap(class QRegularExpressionMatch &)
func (this *QRegularExpressionMatch) Swap(other unsafe.Pointer) {
	// 0: (, other QRegularExpressionMatch &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QRegularExpressionMatch4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:189
// index:0
// QRegularExpression regularExpression()
func (this *QRegularExpressionMatch) RegularExpression() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17regularExpressionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:190
// index:0
// QRegularExpression::MatchType matchType()
func (this *QRegularExpressionMatch) MatchType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch9matchTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:191
// index:0
// QRegularExpression::MatchOptions matchOptions()
func (this *QRegularExpressionMatch) MatchOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12matchOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:193
// index:0
// bool hasMatch()
func (this *QRegularExpressionMatch) HasMatch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8hasMatchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:194
// index:0
// bool hasPartialMatch()
func (this *QRegularExpressionMatch) HasPartialMatch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch15hasPartialMatchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:196
// index:0
// bool isValid()
func (this *QRegularExpressionMatch) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:198
// index:0
// int lastCapturedIndex()
func (this *QRegularExpressionMatch) LastCapturedIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17lastCapturedIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:200
// index:0
// QString captured(int)
func (this *QRegularExpressionMatch) Captured(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:205
// index:1
// QString captured(const class QString &)
func (this *QRegularExpressionMatch) Captured_1(name unsafe.Pointer) {
	// 1: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:209
// index:2
// QString captured(class QStringView)
func (this *QRegularExpressionMatch) Captured_2(name unsafe.Pointer) {
	// 2: (, name QStringView), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:201
// index:0
// QStringRef capturedRef(int)
func (this *QRegularExpressionMatch) CapturedRef(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:206
// index:1
// QStringRef capturedRef(const class QString &)
func (this *QRegularExpressionMatch) CapturedRef_1(name unsafe.Pointer) {
	// 1: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:210
// index:2
// QStringRef capturedRef(class QStringView)
func (this *QRegularExpressionMatch) CapturedRef_2(name unsafe.Pointer) {
	// 2: (, name QStringView), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:202
// index:0
// QStringView capturedView(int)
func (this *QRegularExpressionMatch) CapturedView(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:211
// index:1
// QStringView capturedView(class QStringView)
func (this *QRegularExpressionMatch) CapturedView_1(name unsafe.Pointer) {
	// 1: (, name QStringView), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:213
// index:0
// QStringList capturedTexts()
func (this *QRegularExpressionMatch) CapturedTexts() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedTextsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:215
// index:0
// int capturedStart(int)
func (this *QRegularExpressionMatch) CapturedStart(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:220
// index:1
// int capturedStart(const class QString &)
func (this *QRegularExpressionMatch) CapturedStart_1(name unsafe.Pointer) {
	// 1: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:225
// index:2
// int capturedStart(class QStringView)
func (this *QRegularExpressionMatch) CapturedStart_2(name unsafe.Pointer) {
	// 2: (, name QStringView), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:216
// index:0
// int capturedLength(int)
func (this *QRegularExpressionMatch) CapturedLength(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:221
// index:1
// int capturedLength(const class QString &)
func (this *QRegularExpressionMatch) CapturedLength_1(name unsafe.Pointer) {
	// 1: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:226
// index:2
// int capturedLength(class QStringView)
func (this *QRegularExpressionMatch) CapturedLength_2(name unsafe.Pointer) {
	// 2: (, name QStringView), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:217
// index:0
// int capturedEnd(int)
func (this *QRegularExpressionMatch) CapturedEnd(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:222
// index:1
// int capturedEnd(const class QString &)
func (this *QRegularExpressionMatch) CapturedEnd_1(name unsafe.Pointer) {
	// 1: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:227
// index:2
// int capturedEnd(class QStringView)
func (this *QRegularExpressionMatch) CapturedEnd_2(name unsafe.Pointer) {
	// 2: (, name QStringView), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

//  body block end
