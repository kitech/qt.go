package qtgui

// /usr/include/qt/QtGui/qtextdocument.h
// #include <qtextdocument.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

// QTextObject * createObject(const QTextFormat &)
func (this *QTextDocument) InheritCreateObject(f func(f *QTextFormat) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createObject", f)
}

// QVariant loadResource(int, const QUrl &)
func (this *QTextDocument) InheritLoadResource(f func(type_ int, name *qtcore.QUrl) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "loadResource", f)
}

/*

 */
type QTextDocument struct {
	*qtcore.QObject
}
type QTextDocument_ITF interface {
	qtcore.QObject_ITF
	QTextDocument_PTR() *QTextDocument
}

func (ptr *QTextDocument) QTextDocument_PTR() *QTextDocument { return ptr }

func (this *QTextDocument) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTextDocument) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQTextDocumentFromPointer(cthis unsafe.Pointer) *QTextDocument {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QTextDocument{bcthis0}
}
func (*QTextDocument) NewFromPointer(cthis unsafe.Pointer) *QTextDocument {
	return NewQTextDocumentFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextdocument.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTextDocument) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(QObject *)

/*
Constructs an empty QTextDocument with the given parent.
*/
func (*QTextDocument) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument {
	return NewQTextDocument(parent)
}
func NewQTextDocument(parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(QObject *)

/*
Constructs an empty QTextDocument with the given parent.
*/
func (*QTextDocument) NewForInherit__() *QTextDocument {
	return NewQTextDocument__()
}
func NewQTextDocument__() *QTextDocument {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(const QString &, QObject *)

/*
Constructs an empty QTextDocument with the given parent.
*/
func (*QTextDocument) NewForInherit_1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument {
	return NewQTextDocument_1(text, parent)
}
func NewQTextDocument_1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(const QString &, QObject *)

/*
Constructs an empty QTextDocument with the given parent.
*/
func (*QTextDocument) NewForInherit_1_(text string) *QTextDocument {
	return NewQTextDocument_1_(text)
}
func NewQTextDocument_1_(text string) *QTextDocument {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:121
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextDocument()

/*

 */
func DeleteQTextDocument(this *QTextDocument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextdocument.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * clone(QObject *) const

/*
Creates a new QTextDocument that is a copy of this text document. parent is the parent of the returned text document.
*/
func (this *QTextDocument) Clone(parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument /*777 QTextDocument **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5cloneEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * clone(QObject *) const

/*
Creates a new QTextDocument that is a copy of this text document. parent is the parent of the returned text document.
*/
func (this *QTextDocument) Clone__() *QTextDocument /*777 QTextDocument **/ {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5cloneEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the document is empty; otherwise returns false.
*/
func (this *QTextDocument) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the document.
*/
func (this *QTextDocument) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUndoRedoEnabled(bool)

/*

 */
func (this *QTextDocument) SetUndoRedoEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument18setUndoRedoEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndoRedoEnabled() const

/*

 */
func (this *QTextDocument) IsUndoRedoEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17isUndoRedoEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndoAvailable() const

/*
Returns true if undo is available; otherwise returns false.

See also isRedoAvailable() and availableUndoSteps().
*/
func (this *QTextDocument) IsUndoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15isUndoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:132
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRedoAvailable() const

/*
Returns true if redo is available; otherwise returns false.

See also isUndoAvailable() and availableRedoSteps().
*/
func (this *QTextDocument) IsRedoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15isRedoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int availableUndoSteps() const

/*
Returns the number of available undo steps.

This function was introduced in  Qt 4.6.

See also isUndoAvailable().
*/
func (this *QTextDocument) AvailableUndoSteps() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument18availableUndoStepsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] int availableRedoSteps() const

/*
Returns the number of available redo steps.

This function was introduced in  Qt 4.6.

See also isRedoAvailable().
*/
func (this *QTextDocument) AvailableRedoSteps() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument18availableRedoStepsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int revision() const

/*
Returns the document's revision (if undo is enabled).

The revision is guaranteed to increase when a document that is not modified is edited.

This function was introduced in  Qt 4.4.

See also QTextBlock::revision() and isModified().
*/
func (this *QTextDocument) Revision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument8revisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentLayout(QAbstractTextDocumentLayout *)

/*
Sets the document to use the given layout. The previous layout is deleted.

See also documentLayoutChanged().
*/
func (this *QTextDocument) SetDocumentLayout(layout QAbstractTextDocumentLayout_ITF /*777 QAbstractTextDocumentLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QAbstractTextDocumentLayout_PTR() != nil {
		convArg0 = layout.QAbstractTextDocumentLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractTextDocumentLayout * documentLayout() const

/*
Returns the document layout for this document.

See also setDocumentLayout().
*/
func (this *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout /*777 QAbstractTextDocumentLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument14documentLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractTextDocumentLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMetaInformation(QTextDocument::MetaInformation, const QString &)

/*
Sets the document's meta information of the type specified by info to the given string.

See also metaInformation().
*/
func (this *QTextDocument) SetMetaInformation(info int, arg1 string) {
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument18setMetaInformationENS_15MetaInformationERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), info, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] QString metaInformation(QTextDocument::MetaInformation) const

/*
Returns meta information about the document of the type specified by info.

See also setMetaInformation().
*/
func (this *QTextDocument) MetaInformation(info int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15metaInformationENS_15MetaInformationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), info)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml(const QByteArray &) const

/*
Returns a string containing an HTML representation of the document.

The encoding parameter specifies the value for the charset attribute in the html header. For example if 'utf-8' is specified then the beginning of the generated html will look like this:


  <html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"></head><body>...



If no encoding is specified then no such meta information is generated.

If you later on convert the returned html string into a byte array for transmission over a network or when saving to disk you should specify the encoding you're going to use for the conversion to a byte array here.

See also Supported HTML Subset.
*/
func (this *QTextDocument) ToHtml(encoding qtcore.QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if encoding != nil && encoding.QByteArray_PTR() != nil {
		convArg0 = encoding.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument6toHtmlERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml(const QByteArray &) const

/*
Returns a string containing an HTML representation of the document.

The encoding parameter specifies the value for the charset attribute in the html header. For example if 'utf-8' is specified then the beginning of the generated html will look like this:


  <html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"></head><body>...



If no encoding is specified then no such meta information is generated.

If you later on convert the returned html string into a byte array for transmission over a network or when saving to disk you should specify the encoding you're going to use for the conversion to a byte array here.

See also Supported HTML Subset.
*/
func (this *QTextDocument) ToHtml__() string {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument6toHtmlERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &)

/*
Replaces the entire contents of the document with the given HTML-formatted text in the html string. The undo/redo history is reset when this function is called.

The HTML formatting is respected as much as possible; for example, "<b>bold</b> text" will produce text where the first word has a font weight that gives it a bold appearance: "bold text".

Note: It is the responsibility of the caller to make sure that the text is correctly decoded when a QString containing HTML is created and passed to setHtml().

See also setPlainText() and Supported HTML Subset.
*/
func (this *QTextDocument) SetHtml(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument7setHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toRawText() const

/*
Returns the raw text contained in the document without any formatting information. If you want formatting information use a QTextCursor instead.

This function was introduced in  Qt 5.9.

See also toPlainText().
*/
func (this *QTextDocument) ToRawText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9toRawTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toPlainText() const

/*
Returns the plain text contained in the document. If you want formatting information use a QTextCursor instead.

This function returns the same as toRawText(), but will replace some unicode characters with ASCII alternatives. In particular, no-break space (U+00A0) is replaced by a regular space (U+0020), and both paragraph (U+2029) and line (U+2028) separators are replaced by line feed (U+000A). If you need the precise contents of the document, use toRawText() instead.

Note: Embedded objects, such as images, are represented by a Unicode value U+FFFC (OBJECT REPLACEMENT CHARACTER).

See also toHtml().
*/
func (this *QTextDocument) ToPlainText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11toPlainTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)

/*
Replaces the entire contents of the document with the given plain text. The undo/redo history is reset when this function is called.

See also setHtml().
*/
func (this *QTextDocument) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12setPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:158
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar characterAt(int) const

/*
Returns the character at position pos, or a null character if the position is out of range.

This function was introduced in  Qt 4.5.

See also characterCount().
*/
func (this *QTextDocument) CharacterAt(pos int) *qtcore.QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11characterAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQChar)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find(subString string, from int, options int) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find__(subString string) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find__1(subString string, from int) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:169
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, const QTextCursor &, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_1(subString string, cursor QTextCursor_ITF, options int) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:169
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, const QTextCursor &, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_1_(subString string, cursor QTextCursor_ITF) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:172
// index:2
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_2(expr qtcore.QRegExp_ITF, from int, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:172
// index:2
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_2_(expr qtcore.QRegExp_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:172
// index:2
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_2_1(expr qtcore.QRegExp_ITF, from int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:173
// index:3
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, const QTextCursor &, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_3(expr qtcore.QRegExp_ITF, cursor QTextCursor_ITF, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:173
// index:3
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, const QTextCursor &, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_3_(expr qtcore.QRegExp_ITF, cursor QTextCursor_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_4(expr qtcore.QRegularExpression_ITF, from int, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressioni6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_4_(expr qtcore.QRegularExpression_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressioni6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_4_1(expr qtcore.QRegularExpression_ITF, from int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressioni6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:178
// index:5
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, const QTextCursor &, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_5(expr qtcore.QRegularExpression_ITF, cursor QTextCursor_ITF, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressionRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:178
// index:5
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, const QTextCursor &, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case-sensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find_5_(expr qtcore.QRegularExpression_ITF, cursor QTextCursor_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressionRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * frameAt(int) const

/*

 */
func (this *QTextDocument) FrameAt(pos int) *QTextFrame /*777 QTextFrame **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument7frameAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * rootFrame() const

/*
Returns the document's root frame.
*/
func (this *QTextDocument) RootFrame() *QTextFrame /*777 QTextFrame **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9rootFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextObject * object(int) const

/*
Returns the text object associated with the given objectIndex.
*/
func (this *QTextDocument) Object(objectIndex int) *QTextObject /*777 QTextObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument6objectEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), objectIndex)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextObject * objectForFormat(const QTextFormat &) const

/*
Returns the text object associated with the format f.
*/
func (this *QTextDocument) ObjectForFormat(arg0 QTextFormat_ITF) *QTextObject /*777 QTextObject **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTextFormat_PTR() != nil {
		convArg0 = arg0.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15objectForFormatERK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:187
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock findBlock(int) const

/*
Returns the text block that contains the pos-th character.
*/
func (this *QTextDocument) FindBlock(pos int) *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9findBlockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:188
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock findBlockByNumber(int) const

/*
Returns the text block with the specified blockNumber.

This function was introduced in  Qt 4.4.

See also QTextBlock::blockNumber().
*/
func (this *QTextDocument) FindBlockByNumber(blockNumber int) *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17findBlockByNumberEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blockNumber)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:189
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock findBlockByLineNumber(int) const

/*
Returns the text block that contains the specified lineNumber.

This function was introduced in  Qt 4.5.

See also QTextBlock::firstLineNumber().
*/
func (this *QTextDocument) FindBlockByLineNumber(blockNumber int) *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument21findBlockByLineNumberEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blockNumber)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:190
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock begin() const

/*
Returns the document's first text block.

See also firstBlock().
*/
func (this *QTextDocument) Begin() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:191
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock end() const

/*
This function returns a block to test for the end of the document while iterating over it.


      for (QTextBlock it = doc->begin(); it != doc->end(); it = it.next())
          cout << it.text().toStdString() << endl;



The block returned is invalid and represents the block after the last block in the document. You can use lastBlock() to retrieve the last valid block of the document.

See also lastBlock().
*/
func (this *QTextDocument) End() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:193
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock firstBlock() const

/*
Returns the document's first text block.

This function was introduced in  Qt 4.4.
*/
func (this *QTextDocument) FirstBlock() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10firstBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:194
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock lastBlock() const

/*
Returns the document's last (valid) text block.

This function was introduced in  Qt 4.4.
*/
func (this *QTextDocument) LastBlock() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9lastBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPageSize(const QSizeF &)

/*

 */
func (this *QTextDocument) SetPageSize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11setPageSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:197
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF pageSize() const

/*

 */
func (this *QTextDocument) PageSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument8pageSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultFont(const QFont &)

/*
Sets the default font to use in the document layout.

Note: Setter function for property defaultFont.

See also defaultFont().
*/
func (this *QTextDocument) SetDefaultFont(font QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14setDefaultFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:200
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont defaultFont() const

/*
Returns the default font to be used in the document layout.

Note: Getter function for property defaultFont.

See also setDefaultFont().
*/
func (this *QTextDocument) DefaultFont() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11defaultFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:202
// index:0
// Public Visibility=Default Availability=Available
// [4] int pageCount() const

/*
returns the number of pages in this document.
*/
func (this *QTextDocument) PageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9pageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:204
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModified() const

/*

 */
func (this *QTextDocument) IsModified() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10isModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void print(QPagedPaintDevice *) const

/*
Prints the document to the given printer. The QPageablePaintDevice must be set up before being used with this function.

This is only a convenience method to print the whole document to the printer.

If the document is already paginated through a specified height in the pageSize() property it is printed as-is.

If the document is not paginated, like for example a document used in a QTextEdit, then a temporary copy of the document is created and the copy is broken into multiple pages according to the size of the paint device's paperRect(). By default a 2 cm margin is set around the document contents. In addition the current page number is printed at the bottom of each page.

See also QTextEdit::print().
*/
func (this *QTextDocument) Print(printer QPagedPaintDevice_ITF /*777 QPagedPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if printer != nil && printer.QPagedPaintDevice_PTR() != nil {
		convArg0 = printer.QPagedPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5printEP17QPagedPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:216
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant resource(int, const QUrl &) const

/*
Returns data of the specified type from the resource with the given name.

This function is called by the rich text engine to request data that isn't directly stored by QTextDocument, but still associated with it. For example, images are referenced indirectly by the name attribute of a QTextImageFormat object.

Resources are cached internally in the document. If a resource can not be found in the cache, loadResource is called to try to load the resource. loadResource should then use addResource to add the resource to the cache.

See also QTextDocument::ResourceType.
*/
func (this *QTextDocument) Resource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument8resourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addResource(int, const QUrl &, const QVariant &)

/*
Adds the resource resource to the resource cache, using type and name as identifiers. type should be a value from QTextDocument::ResourceType.

For example, you can add an image as a resource in order to reference it from within the document:


      document->addResource(QTextDocument::ImageResource,
          QUrl("mydata://image.png"), QVariant(image));



The image can be inserted into the document using the QTextCursor API:


      QTextImageFormat imageFormat;
      imageFormat.setName("mydata://image.png");
      cursor.insertImage(imageFormat);



Alternatively, you can insert images using the HTML img tag:


      editor->append("<img src=\"mydata://image.png\" />");
*/
func (this *QTextDocument) AddResource(type_ int, name qtcore.QUrl_ITF, resource qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if resource != nil && resource.QVariant_PTR() != nil {
		convArg2 = resource.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markContentsDirty(int, int)

/*
Marks the contents specified by the given position and length as "dirty", informing the document that it needs to be laid out again.
*/
func (this *QTextDocument) MarkContentsDirty(from int, length int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17markContentsDirtyEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, length)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUseDesignMetrics(bool)

/*

 */
func (this *QTextDocument) SetUseDesignMetrics(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19setUseDesignMetricsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:224
// index:0
// Public Visibility=Default Availability=Available
// [1] bool useDesignMetrics() const

/*

 */
func (this *QTextDocument) UseDesignMetrics() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument16useDesignMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawContents(QPainter *, const QRectF &)

/*
Draws the content of the document with painter p, clipped to rect. If rect is a null rectangle (default) then the document is painted unclipped.

This function was introduced in  Qt 4.2.
*/
func (this *QTextDocument) DrawContents(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawContents(QPainter *, const QRectF &)

/*
Draws the content of the document with painter p, clipped to rect. If rect is a null rectangle (default) then the document is painted unclipped.

This function was introduced in  Qt 4.2.
*/
func (this *QTextDocument) DrawContents__(painter QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 1, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextWidth(qreal)

/*

 */
func (this *QTextDocument) SetTextWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12setTextWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal textWidth() const

/*

 */
func (this *QTextDocument) TextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9textWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:231
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal idealWidth() const

/*
Returns the ideal width of the text document. The ideal width is the actually used width of the document without optional alignments taken into account. It is always <= size().width().

This function was introduced in  Qt 4.2.

See also adjustSize() and textWidth.
*/
func (this *QTextDocument) IdealWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10idealWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal indentWidth() const

/*

 */
func (this *QTextDocument) IndentWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11indentWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndentWidth(qreal)

/*
Sets the width used for text list and text block indenting.

The indent properties of QTextListFormat and QTextBlockFormat specify multiples of this value. The default indent width is 40 .

This function was introduced in  Qt 4.4.

Note: Setter function for property indentWidth.

See also indentWidth().
*/
func (this *QTextDocument) SetIndentWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14setIndentWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal documentMargin() const

/*

 */
func (this *QTextDocument) DocumentMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument14documentMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMargin(qreal)

/*

 */
func (this *QTextDocument) SetDocumentMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()

/*
Adjusts the document to a reasonable size.

This function was introduced in  Qt 4.2.

See also idealWidth(), textWidth, and size.
*/
func (this *QTextDocument) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:240
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size() const

/*

 */
func (this *QTextDocument) Size() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:242
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockCount() const

/*

 */
func (this *QTextDocument) BlockCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10blockCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineCount() const

/*
Returns the number of lines of this document (if the layout supports this). Otherwise, this is identical to the number of blocks.

This function was introduced in  Qt 4.5.

See also blockCount() and characterCount().
*/
func (this *QTextDocument) LineCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9lineCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] int characterCount() const

/*
Returns the number of characters of this document.

This function was introduced in  Qt 4.5.

See also blockCount() and characterAt().
*/
func (this *QTextDocument) CharacterCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument14characterCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultStyleSheet(const QString &)

/*

 */
func (this *QTextDocument) SetDefaultStyleSheet(sheet string) {
	var tmpArg0 = qtcore.NewQString_5(sheet)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultStyleSheetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:248
// index:0
// Public Visibility=Default Availability=Available
// [8] QString defaultStyleSheet() const

/*

 */
func (this *QTextDocument) DefaultStyleSheet() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17defaultStyleSheetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo(QTextCursor *)

/*
Undoes the last editing operation on the document if undo is available. The provided cursor is positioned at the end of the location where the edition operation was undone.

See the Qt Undo Framework documentation for details.

This function was introduced in  Qt 4.2.

See also undoAvailable() and isUndoRedoEnabled().
*/
func (this *QTextDocument) Undo(cursor QTextCursor_ITF /*777 QTextCursor **/) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4undoEP11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:286
// index:1
// Public Visibility=Default Availability=Available
// [-2] void undo()

/*
Undoes the last editing operation on the document if undo is available. The provided cursor is positioned at the end of the location where the edition operation was undone.

See the Qt Undo Framework documentation for details.

This function was introduced in  Qt 4.2.

See also undoAvailable() and isUndoRedoEnabled().
*/
func (this *QTextDocument) Undo_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo(QTextCursor *)

/*
Redoes the last editing operation on the document if redo is available.

The provided cursor is positioned at the end of the location where the edition operation was redone.

This function was introduced in  Qt 4.2.
*/
func (this *QTextDocument) Redo(cursor QTextCursor_ITF /*777 QTextCursor **/) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4redoEP11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:287
// index:1
// Public Visibility=Default Availability=Available
// [-2] void redo()

/*
Redoes the last editing operation on the document if redo is available.

The provided cursor is positioned at the end of the location where the edition operation was redone.

This function was introduced in  Qt 4.2.
*/
func (this *QTextDocument) Redo_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearUndoRedoStacks(QTextDocument::Stacks)

/*
Clears the stacks specified by stacksToClear.

This method clears any commands on the undo stack, the redo stack, or both (the default). If commands are cleared, the appropriate signals are emitted, QTextDocument::undoAvailable() or QTextDocument::redoAvailable().

This function was introduced in  Qt 4.7.

See also QTextDocument::undoAvailable() and QTextDocument::redoAvailable().
*/
func (this *QTextDocument) ClearUndoRedoStacks(historyToClear int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19clearUndoRedoStacksENS_6StacksE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), historyToClear)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearUndoRedoStacks(QTextDocument::Stacks)

/*
Clears the stacks specified by stacksToClear.

This method clears any commands on the undo stack, the redo stack, or both (the default). If commands are cleared, the appropriate signals are emitted, QTextDocument::undoAvailable() or QTextDocument::redoAvailable().

This function was introduced in  Qt 4.7.

See also QTextDocument::undoAvailable() and QTextDocument::redoAvailable().
*/
func (this *QTextDocument) ClearUndoRedoStacks__() {
	// arg: 0, QTextDocument::Stacks=Enum, QTextDocument::Stacks=Enum, , Invalid
	historyToClear := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19clearUndoRedoStacksENS_6StacksE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), historyToClear)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:261
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumBlockCount() const

/*

 */
func (this *QTextDocument) MaximumBlockCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17maximumBlockCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumBlockCount(int)

/*

 */
func (this *QTextDocument) SetMaximumBlockCount(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument20setMaximumBlockCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:264
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextOption defaultTextOption() const

/*
The default text option is used on all QTextLayout objects in the document. This allows setting global properties for the document such as the default word wrap mode.

This function was introduced in  Qt 4.3.

Note: Getter function for property defaultTextOption.

See also setDefaultTextOption().
*/
func (this *QTextDocument) DefaultTextOption() *QTextOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17defaultTextOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTextOption(const QTextOption &)

/*
Sets the default text option to option.

This function was introduced in  Qt 4.3.

Note: Setter function for property defaultTextOption.

See also defaultTextOption().
*/
func (this *QTextDocument) SetDefaultTextOption(option QTextOption_ITF) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QTextOption_PTR() != nil {
		convArg0 = option.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultTextOptionERK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:267
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl baseUrl() const

/*

 */
func (this *QTextDocument) BaseUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument7baseUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseUrl(const QUrl &)

/*

 */
func (this *QTextDocument) SetBaseUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument10setBaseUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:270
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorMoveStyle defaultCursorMoveStyle() const

/*
The default cursor movement style is used by all QTextCursor objects created from the document. The default is Qt::LogicalMoveStyle.

This function was introduced in  Qt 4.8.

See also setDefaultCursorMoveStyle().
*/
func (this *QTextDocument) DefaultCursorMoveStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument22defaultCursorMoveStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultCursorMoveStyle(Qt::CursorMoveStyle)

/*
Sets the default cursor movement style to the given style.

This function was introduced in  Qt 4.8.

See also defaultCursorMoveStyle().
*/
func (this *QTextDocument) SetDefaultCursorMoveStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument25setDefaultCursorMoveStyleEN2Qt15CursorMoveStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentsChange(int, int, int)

/*
This signal is emitted whenever the document's content changes; for example, when text is inserted or deleted, or when formatting is applied.

Information is provided about the position of the character in the document where the change occurred, the number of characters removed (charsRemoved), and the number of characters added (charsAdded).

The signal is emitted before the document's layout manager is notified about the change. This hook allows you to implement syntax highlighting for the document.

See also QAbstractTextDocumentLayout::documentChanged() and contentsChanged().
*/
func (this *QTextDocument) ContentsChange(from int, charsRemoved int, charsAdded int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14contentsChangeEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, charsRemoved, charsAdded)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentsChanged()

/*
This signal is emitted whenever the document's content changes; for example, when text is inserted or deleted, or when formatting is applied.

See also contentsChange().
*/
func (this *QTextDocument) ContentsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument15contentsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoAvailable(bool)

/*
This signal is emitted whenever undo operations become available (available is true) or unavailable (available is false).

See the Qt Undo Framework documentation for details.

See also undo() and isUndoRedoEnabled().
*/
func (this *QTextDocument) UndoAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument13undoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:277
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoAvailable(bool)

/*
This signal is emitted whenever redo operations become available (available is true) or unavailable (available is false).
*/
func (this *QTextDocument) RedoAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument13redoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoCommandAdded()

/*
This signal is emitted every time a new level of undo is added to the QTextDocument.

This function was introduced in  Qt 4.4.
*/
func (this *QTextDocument) UndoCommandAdded() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument16undoCommandAddedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:279
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modificationChanged(bool)

/*
This signal is emitted whenever the content of the document changes in a way that affects the modification state. If changed is true, the document has been modified; otherwise it is false.

For example, calling setModified(false) on a document and then inserting text causes the signal to get emitted. If you undo that operation, causing the document to return to its original unmodified state, the signal will get emitted again.
*/
func (this *QTextDocument) ModificationChanged(m bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19modificationChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:280
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorPositionChanged(const QTextCursor &)

/*
This signal is emitted whenever the position of a cursor changed due to an editing operation. The cursor that changed is passed in cursor. If the document is used with the QTextEdit class and you need a signal when the cursor is moved with the arrow keys you can use the cursorPositionChanged() signal in QTextEdit.
*/
func (this *QTextDocument) CursorPositionChanged(cursor QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument21cursorPositionChangedERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void blockCountChanged(int)

/*
This signal is emitted when the total number of text blocks in the document changes. The value passed in newBlockCount is the new total.

This function was introduced in  Qt 4.3.
*/
func (this *QTextDocument) BlockCountChanged(newBlockCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17blockCountChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newBlockCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void baseUrlChanged(const QUrl &)

/*

 */
func (this *QTextDocument) BaseUrlChanged(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14baseUrlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void documentLayoutChanged()

/*
This signal is emitted when a new document layout is set.

This function was introduced in  Qt 4.4.

See also setDocumentLayout().
*/
func (this *QTextDocument) DocumentLayoutChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument21documentLayoutChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendUndoItem(QAbstractUndoItem *)

/*

 */
func (this *QTextDocument) AppendUndoItem(arg0 QAbstractUndoItem_ITF /*777 QAbstractUndoItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAbstractUndoItem_PTR() != nil {
		convArg0 = arg0.QAbstractUndoItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModified(bool)

/*

 */
func (this *QTextDocument) SetModified(m bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11setModifiedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModified(bool)

/*

 */
func (this *QTextDocument) SetModified__() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	m := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11setModifiedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:292
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QTextObject * createObject(const QTextFormat &)

/*
Creates and returns a new document object (a QTextObject), based on the given format.

QTextObjects will always get created through this method, so you must reimplement it if you use custom text objects inside your document.
*/
func (this *QTextDocument) CreateObject(f QTextFormat_ITF) *QTextObject /*777 QTextObject **/ {
	var convArg0 unsafe.Pointer
	if f != nil && f.QTextFormat_PTR() != nil {
		convArg0 = f.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12createObjectERK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:293
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant loadResource(int, const QUrl &)

/*
Loads data of the specified type from the resource with the given name.

This function is called by the rich text engine to request data that isn't directly stored by QTextDocument, but still associated with it. For example, images are referenced indirectly by the name attribute of a QTextImageFormat object.

When called by Qt, type is one of the values of QTextDocument::ResourceType.

If the QTextDocument is a child object of a QObject that has an invokable loadResource method such as QTextEdit, QTextBrowser or a QTextDocument itself then the default implementation tries to retrieve the data from the parent.
*/
func (this *QTextDocument) LoadResource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12loadResourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

/*
This enum describes the different types of meta information that can be added to a document.



See also metaInformation() and setMetaInformation().

*/
type QTextDocument__MetaInformation = int

// The title of the document.
const QTextDocument__DocumentTitle QTextDocument__MetaInformation = 0

// The url of the document. The loadResource() function uses this url as the base when loading relative resources.
const QTextDocument__DocumentUrl QTextDocument__MetaInformation = 1

func (this *QTextDocument) MetaInformationItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTextDocument_MetaInformationItemName(val int) string {
	var nilthis *QTextDocument
	return nilthis.MetaInformationItemName(val)
}

/*


 */
type QTextDocument__FindFlag = int

//
const QTextDocument__FindBackward QTextDocument__FindFlag = 1

//
const QTextDocument__FindCaseSensitively QTextDocument__FindFlag = 2

//
const QTextDocument__FindWholeWords QTextDocument__FindFlag = 4

func (this *QTextDocument) FindFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTextDocument_FindFlagItemName(val int) string {
	var nilthis *QTextDocument
	return nilthis.FindFlagItemName(val)
}

/*
This enum describes the types of resources that can be loaded by QTextDocument's loadResource() function.



See also loadResource().

*/
type QTextDocument__ResourceType = int

// The resource contains HTML.
const QTextDocument__HtmlResource QTextDocument__ResourceType = 1

// The resource contains image data. Currently supported data types are QVariant::Pixmap and QVariant::Image. If the corresponding variant is of type QVariant::ByteArray then Qt attempts to load the image using QImage::loadFromData. QVariant::Icon is currently not supported. The icon needs to be converted to one of the supported types first, for example using QIcon::pixmap.
const QTextDocument__ImageResource QTextDocument__ResourceType = 2

// The resource contains CSS.
const QTextDocument__StyleSheetResource QTextDocument__ResourceType = 3

//
const QTextDocument__UserResource QTextDocument__ResourceType = 100

func (this *QTextDocument) ResourceTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTextDocument_ResourceTypeItemName(val int) string {
	var nilthis *QTextDocument
	return nilthis.ResourceTypeItemName(val)
}

/*
QTextDocument::UndoAndRedoStacksUndoStack | RedoStackBoth the undo and redo stacks.

*/
type QTextDocument__Stacks = int

//
const QTextDocument__UndoStack QTextDocument__Stacks = 1

//
const QTextDocument__RedoStack QTextDocument__Stacks = 2

//
const QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = 3

func (this *QTextDocument) StacksItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTextDocument_StacksItemName(val int) string {
	var nilthis *QTextDocument
	return nilthis.StacksItemName(val)
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
