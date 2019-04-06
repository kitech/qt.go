//  header block begin

// +build !minimal

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

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags) const

/*
Finds the next occurrence of the string, subString, in the document. The search starts at the position of the given cursor, and proceeds forwards through the document unless specified otherwise in the search options. The options control the type of search performed.

Returns a cursor with the match selected if subString was found; otherwise returns a null cursor.

If the given cursor has a selection, the search begins after the selection; otherwise it begins at the cursor's position.

By default the search is case insensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find4(expr qtcore.QRegularExpression_ITF, from int, options int) *QTextCursor /*123*/ {
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

By default the search is case insensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find4p(expr qtcore.QRegularExpression_ITF) *QTextCursor /*123*/ {
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

By default the search is case insensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find4p1(expr qtcore.QRegularExpression_ITF, from int) *QTextCursor /*123*/ {
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

By default the search is case insensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find5(expr qtcore.QRegularExpression_ITF, cursor QTextCursor_ITF, options int) *QTextCursor /*123*/ {
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

By default the search is case insensitive, and can match text anywhere in the document.
*/
func (this *QTextDocument) Find5p(expr qtcore.QRegularExpression_ITF, cursor QTextCursor_ITF) *QTextCursor /*123*/ {
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

//  body block end

//  keep block begin

func init_unused_10794() {
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
