package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QTextBlock struct {
	*qtrt.CObject
}
type QTextBlock_ITF interface {
	QTextBlock_PTR() *QTextBlock
}

func (ptr *QTextBlock) QTextBlock_PTR() *QTextBlock { return ptr }

func (this *QTextBlock) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextBlock) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextBlockFromPointer(cthis unsafe.Pointer) *QTextBlock {
	return &QTextBlock{&qtrt.CObject{cthis}}
}
func (*QTextBlock) NewFromPointer(cthis unsafe.Pointer) *QTextBlock {
	return NewQTextBlockFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextBlock()

/*

 */
func NewQTextBlock() *QTextBlock {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlockC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBlock)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextBlock & operator=(const QTextBlock &)

/*

 */
func (this *QTextBlock) Operator_equal(o QTextBlock_ITF) *QTextBlock {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextBlock_PTR() != nil {
		convArg0 = o.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlockaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:210
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QTextBlock) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:212
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QTextBlock &) const

/*

 */
func (this *QTextBlock) Operator_equal_equal(o QTextBlock_ITF) bool {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextBlock_PTR() != nil {
		convArg0 = o.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlockeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:213
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QTextBlock &) const

/*

 */
func (this *QTextBlock) Operator_not_equal(o QTextBlock_ITF) bool {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextBlock_PTR() != nil {
		convArg0 = o.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlockneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QTextBlock &) const

/*

 */
func (this *QTextBlock) Operator_less_than(o QTextBlock_ITF) bool {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextBlock_PTR() != nil {
		convArg0 = o.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlockltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:216
// index:0
// Public Visibility=Default Availability=Available
// [4] int position() const

/*

 */
func (this *QTextBlock) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:217
// index:0
// Public Visibility=Default Availability=Available
// [4] int length() const

/*

 */
func (this *QTextBlock) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:218
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(int) const

/*

 */
func (this *QTextBlock) Contains(position int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8containsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextLayout * layout() const

/*

 */
func (this *QTextBlock) Layout() *QTextLayout /*777 QTextLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearLayout()

/*

 */
func (this *QTextBlock) ClearLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock11clearLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:222
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlockFormat blockFormat() const

/*

 */
func (this *QTextBlock) BlockFormat() *QTextBlockFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock11blockFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlockFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockFormatIndex() const

/*

 */
func (this *QTextBlock) BlockFormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock16blockFormatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:224
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat charFormat() const

/*

 */
func (this *QTextBlock) CharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock10charFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] int charFormatIndex() const

/*

 */
func (this *QTextBlock) CharFormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock15charFormatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:227
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection textDirection() const

/*

 */
func (this *QTextBlock) TextDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock13textDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextobject.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QTextBlock) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextobject.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] const QTextDocument * document() const

/*
Returns the document this object belongs to.

See also format().
*/
func (this *QTextBlock) Document() *QTextDocument /*777 const QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextList * textList() const

/*

 */
func (this *QTextBlock) TextList() *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8textListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:237
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextBlockUserData * userData() const

/*

 */
func (this *QTextBlock) UserData() *QTextBlockUserData /*777 QTextBlockUserData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8userDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextBlockUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserData(QTextBlockUserData *)

/*

 */
func (this *QTextBlock) SetUserData(data QTextBlockUserData_ITF /*777 QTextBlockUserData **/) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QTextBlockUserData_PTR() != nil {
		convArg0 = data.QTextBlockUserData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock11setUserDataEP18QTextBlockUserData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:240
// index:0
// Public Visibility=Default Availability=Available
// [4] int userState() const

/*

 */
func (this *QTextBlock) UserState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock9userStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserState(int)

/*

 */
func (this *QTextBlock) SetUserState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock12setUserStateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int revision() const

/*

 */
func (this *QTextBlock) Revision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8revisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:244
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRevision(int)

/*

 */
func (this *QTextBlock) SetRevision(rev int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock11setRevisionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rev)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:246
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QTextBlock) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)

/*

 */
func (this *QTextBlock) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:249
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockNumber() const

/*

 */
func (this *QTextBlock) BlockNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock11blockNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:250
// index:0
// Public Visibility=Default Availability=Available
// [4] int firstLineNumber() const

/*

 */
func (this *QTextBlock) FirstLineNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock15firstLineNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineCount(int)

/*

 */
func (this *QTextBlock) SetLineCount(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlock12setLineCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineCount() const

/*

 */
func (this *QTextBlock) LineCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock9lineCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:283
// index:0
// Public Visibility=Default Availability=Available
// [24] QTextBlock::iterator begin() const

/*

 */
func (this *QTextBlock) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qtextobject.h:284
// index:0
// Public Visibility=Default Availability=Available
// [24] QTextBlock::iterator end() const

/*

 */
func (this *QTextBlock) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qtextobject.h:286
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock next() const

/*

 */
func (this *QTextBlock) Next() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:287
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock previous() const

/*

 */
func (this *QTextBlock) Previous() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock8previousEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:290
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int fragmentIndex() const

/*

 */
func (this *QTextBlock) FragmentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextBlock13fragmentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTextBlock(this *QTextBlock) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextBlockD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
}

//  keep block end
