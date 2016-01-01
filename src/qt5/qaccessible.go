package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qaccessible.h
// dst-file: /src/gui/qaccessible.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QAccessible)=1
type QAccessible struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTableModelChangeEvent)=48
type QAccessibleTableModelChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextInterface)=8
type QAccessibleTextInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleEvent)=32
type QAccessibleEvent struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleActionInterface)=8
type QAccessibleActionInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleInterface)=8
type QAccessibleInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleEditableTextInterface)=8
type QAccessibleEditableTextInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTableCellInterface)=8
type QAccessibleTableCellInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTableInterface)=8
type QAccessibleTableInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextUpdateEvent)=56
type QAccessibleTextUpdateEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleStateChangeEvent)=40
type QAccessibleStateChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleImageInterface)=8
type QAccessibleImageInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextInsertEvent)=48
type QAccessibleTextInsertEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleValueInterface)=8
type QAccessibleValueInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextRemoveEvent)=48
type QAccessibleTextRemoveEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextSelectionEvent)=40
type QAccessibleTextSelectionEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextCursorEvent)=32
type QAccessibleTextCursorEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleValueChangeEvent)=48
type QAccessibleValueChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QAccessible) isActive_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "isActive", args)
 }

}


func (this *QAccessible) uniqueId_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "uniqueId", args)
 }

}


func (this *QAccessible) registerAccessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "registerAccessibleInterface", args)
 }

}


func (this *QAccessible) setActive_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "setActive", args)
 }

}


func (this *QAccessible) queryAccessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "queryAccessibleInterface", args)
 }

}


func (this *QAccessible) updateAccessibility_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "updateAccessibility", args)
 }

}


func (this *QAccessible) cleanup_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "cleanup", args)
 }

}


func (this *QAccessible) setRootObject_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "setRootObject", args)
 }

}


func (this *QAccessible) deleteAccessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "deleteAccessibleInterface", args)
 }

}


func (this *QAccessible) accessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "accessibleInterface", args)
 }

}


func NewQAccessible(args ...interface{})() {
}


func (this *QAccessibleTableModelChangeEvent) setFirstColumn(args ...interface{}) () {
  // setFirstColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstColumn", args)
 }

}


func (this *QAccessibleTableModelChangeEvent) setFirstRow(args ...interface{}) () {
  // setFirstRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent11setFirstRowEi
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstRow", args)
 }

}


func (this *QAccessibleTableModelChangeEvent) firstRow(args ...interface{}) () {
  // firstRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent8firstRowEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstRow", args)
 }

}


func (this *QAccessibleTableModelChangeEvent) setLastColumn(args ...interface{}) () {
  // setLastColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent13setLastColumnEi
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastColumn", args)
 }

}


func (this *QAccessibleTableModelChangeEvent) firstColumn(args ...interface{}) () {
  // firstColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent11firstColumnEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstColumn", args)
 }

}


func (this *QAccessibleTableModelChangeEvent) lastColumn(args ...interface{}) () {
  // lastColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent10lastColumnEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastColumn", args)
 }

}


func (this *QAccessibleTableModelChangeEvent) setLastRow(args ...interface{}) () {
  // setLastRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent10setLastRowEi
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastRow", args)
 }

}


func (this *QAccessibleTableModelChangeEvent) lastRow(args ...interface{}) () {
  // lastRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent7lastRowEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastRow", args)
 }

}


func (this *QAccessibleTextInterface) selection(args ...interface{}) () {
  // selection(int, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface9selectionEiPiS0_
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "selection", args)
 }

}


func (this *QAccessibleTextInterface) setCursorPosition(args ...interface{}) () {
  // setCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface17setCursorPositionEi
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "setCursorPosition", args)
 }

}


func (this *QAccessibleTextInterface) offsetAtPoint(args ...interface{}) () {
  // offsetAtPoint(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface13offsetAtPointERK6QPoint
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "offsetAtPoint", args)
 }

}


func (this *QAccessibleTextInterface) attributes(args ...interface{}) () {
  // attributes(int, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface10attributesEiPiS0_
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "attributes", args)
 }

}


func (this *QAccessibleTextInterface) selectionCount(args ...interface{}) () {
  // selectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface14selectionCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "selectionCount", args)
 }

}


func (this *QAccessibleTextInterface) characterCount(args ...interface{}) () {
  // characterCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface14characterCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "characterCount", args)
 }

}


func (this *QAccessibleTextInterface) FreeQAccessibleTextInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "~QAccessibleTextInterface", args)
 }

}


func (this *QAccessibleTextInterface) text(args ...interface{}) () {
  // text(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface4textEii
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "text", args)
 }

}


func (this *QAccessibleTextInterface) characterRect(args ...interface{}) () {
  // characterRect(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface13characterRectEi
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "characterRect", args)
 }

}


func (this *QAccessibleTextInterface) removeSelection(args ...interface{}) () {
  // removeSelection(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface15removeSelectionEi
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "removeSelection", args)
 }

}


func (this *QAccessibleTextInterface) addSelection(args ...interface{}) () {
  // addSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface12addSelectionEii
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "addSelection", args)
 }

}


func (this *QAccessibleTextInterface) scrollToSubstring(args ...interface{}) () {
  // scrollToSubstring(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface17scrollToSubstringEii
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "scrollToSubstring", args)
 }

}


func (this *QAccessibleTextInterface) cursorPosition(args ...interface{}) () {
  // cursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface14cursorPositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "cursorPosition", args)
 }

}


func (this *QAccessibleTextInterface) setSelection(args ...interface{}) () {
  // setSelection(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface12setSelectionEiii
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "setSelection", args)
 }

}


func (this *QAccessibleEvent) object(args ...interface{}) () {
  // object()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent6objectEv
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "object", args)
 }

}


func (this *QAccessibleEvent) setChild(args ...interface{}) () {
  // setChild(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QAccessibleEvent8setChildEi
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "setChild", args)
 }

}


func (this *QAccessibleEvent) accessibleInterface(args ...interface{}) () {
  // accessibleInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent19accessibleInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "accessibleInterface", args)
 }

}


func NewQAccessibleEvent(args ...interface{})() {
}


func (this *QAccessibleEvent) child(args ...interface{}) () {
  // child()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent5childEv
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "child", args)
 }

}


func (this *QAccessibleEvent) FreeQAccessibleEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "~QAccessibleEvent", args)
 }

}


func (this *QAccessibleActionInterface) scrollUpAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollUpAction", args)
 }

}


func (this *QAccessibleActionInterface) actionNames(args ...interface{}) () {
  // actionNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface11actionNamesEv
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "actionNames", args)
 }

}


func (this *QAccessibleActionInterface) decreaseAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "decreaseAction", args)
 }

}


func (this *QAccessibleActionInterface) toggleAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "toggleAction", args)
 }

}


func (this *QAccessibleActionInterface) localizedActionName(args ...interface{}) () {
  // localizedActionName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface19localizedActionNameERK7QString
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionName", args)
 }

}


func (this *QAccessibleActionInterface) localizedActionDescription(args ...interface{}) () {
  // localizedActionDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionDescription", args)
 }

}


func (this *QAccessibleActionInterface) scrollLeftAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollLeftAction", args)
 }

}


func (this *QAccessibleActionInterface) previousPageAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "previousPageAction", args)
 }

}


func (this *QAccessibleActionInterface) showMenuAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "showMenuAction", args)
 }

}


func (this *QAccessibleActionInterface) scrollRightAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollRightAction", args)
 }

}


func (this *QAccessibleActionInterface) setFocusAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "setFocusAction", args)
 }

}


func (this *QAccessibleActionInterface) nextPageAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "nextPageAction", args)
 }

}


func (this *QAccessibleActionInterface) FreeQAccessibleActionInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "~QAccessibleActionInterface", args)
 }

}


func (this *QAccessibleActionInterface) pressAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "pressAction", args)
 }

}


func (this *QAccessibleActionInterface) doAction(args ...interface{}) () {
  // doAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface8doActionERK7QString
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "doAction", args)
 }

}


func (this *QAccessibleActionInterface) increaseAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "increaseAction", args)
 }

}


func (this *QAccessibleActionInterface) keyBindingsForAction(args ...interface{}) () {
  // keyBindingsForAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface20keyBindingsForActionERK7QString
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "keyBindingsForAction", args)
 }

}


func (this *QAccessibleActionInterface) scrollDownAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollDownAction", args)
 }

}


func (this *QAccessibleInterface) imageInterface(args ...interface{}) () {
  // imageInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface14imageInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "imageInterface", args)
 }

}


func (this *QAccessibleInterface) tableInterface(args ...interface{}) () {
  // tableInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface14tableInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableInterface", args)
 }

}


func (this *QAccessibleInterface) editableTextInterface(args ...interface{}) () {
  // editableTextInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface21editableTextInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "editableTextInterface", args)
 }

}


func (this *QAccessibleInterface) valueInterface(args ...interface{}) () {
  // valueInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface14valueInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "valueInterface", args)
 }

}


func (this *QAccessibleInterface) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface4rectEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "rect", args)
 }

}


func (this *QAccessibleInterface) object(args ...interface{}) () {
  // object()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface6objectEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "object", args)
 }

}


func (this *QAccessibleInterface) actionInterface(args ...interface{}) () {
  // actionInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface15actionInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "actionInterface", args)
 }

}


func (this *QAccessibleInterface) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface6parentEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "parent", args)
 }

}


func (this *QAccessibleInterface) childAt(args ...interface{}) () {
  // childAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface7childAtEii
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "childAt", args)
 }

}


func (this *QAccessibleInterface) childCount(args ...interface{}) () {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface10childCountEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "childCount", args)
 }

}


func (this *QAccessibleInterface) tableCellInterface(args ...interface{}) () {
  // tableCellInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface18tableCellInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableCellInterface", args)
 }

}


func (this *QAccessibleInterface) indexOfChild(args ...interface{}) () {
  // indexOfChild(const class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "const QAccessibleInterface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface12indexOfChildEPKS_
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "indexOfChild", args)
 }

}


func (this *QAccessibleInterface) foregroundColor(args ...interface{}) () {
  // foregroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface15foregroundColorEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "foregroundColor", args)
 }

}


func (this *QAccessibleInterface) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface7isValidEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "isValid", args)
 }

}


func (this *QAccessibleInterface) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface6windowEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "window", args)
 }

}


func (this *QAccessibleInterface) virtual_hook(args ...interface{}) () {
  // virtual_hook(int, void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface12virtual_hookEiPv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "virtual_hook", args)
 }

}


func (this *QAccessibleInterface) focusChild(args ...interface{}) () {
  // focusChild()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface10focusChildEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "focusChild", args)
 }

}


func (this *QAccessibleInterface) child(args ...interface{}) () {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface5childEi
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "child", args)
 }

}


func (this *QAccessibleInterface) textInterface(args ...interface{}) () {
  // textInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface13textInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "textInterface", args)
 }

}


func (this *QAccessibleInterface) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface15backgroundColorEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "backgroundColor", args)
 }

}


func (this *QAccessibleInterface) FreeQAccessibleInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "~QAccessibleInterface", args)
 }

}


func (this *QAccessibleEditableTextInterface) insertText(args ...interface{}) () {
  // insertText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleEditableTextInterface10insertTextEiRK7QString
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "insertText", args)
 }

}


func (this *QAccessibleEditableTextInterface) replaceText(args ...interface{}) () {
  // replaceText(int, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleEditableTextInterface11replaceTextEiiRK7QString
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "replaceText", args)
 }

}


func (this *QAccessibleEditableTextInterface) deleteText(args ...interface{}) () {
  // deleteText(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleEditableTextInterface10deleteTextEii
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "deleteText", args)
 }

}


func (this *QAccessibleEditableTextInterface) FreeQAccessibleEditableTextInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "~QAccessibleEditableTextInterface", args)
 }

}


func (this *QAccessibleTableCellInterface) columnIndex(args ...interface{}) () {
  // columnIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface11columnIndexEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "columnIndex", args)
 }

}


func (this *QAccessibleTableCellInterface) FreeQAccessibleTableCellInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "~QAccessibleTableCellInterface", args)
 }

}


func (this *QAccessibleTableCellInterface) columnExtent(args ...interface{}) () {
  // columnExtent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface12columnExtentEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "columnExtent", args)
 }

}


func (this *QAccessibleTableCellInterface) rowIndex(args ...interface{}) () {
  // rowIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface8rowIndexEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "rowIndex", args)
 }

}


func (this *QAccessibleTableCellInterface) table(args ...interface{}) () {
  // table()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface5tableEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "table", args)
 }

}


func (this *QAccessibleTableCellInterface) rowExtent(args ...interface{}) () {
  // rowExtent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface9rowExtentEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "rowExtent", args)
 }

}


func (this *QAccessibleTableCellInterface) rowHeaderCells(args ...interface{}) () {
  // rowHeaderCells()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface14rowHeaderCellsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "rowHeaderCells", args)
 }

}


func (this *QAccessibleTableCellInterface) columnHeaderCells(args ...interface{}) () {
  // columnHeaderCells()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface17columnHeaderCellsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "columnHeaderCells", args)
 }

}


func (this *QAccessibleTableCellInterface) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface10isSelectedEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "isSelected", args)
 }

}


func (this *QAccessibleTableInterface) unselectColumn(args ...interface{}) () {
  // unselectColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface14unselectColumnEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "unselectColumn", args)
 }

}


func (this *QAccessibleTableInterface) columnDescription(args ...interface{}) () {
  // columnDescription(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface17columnDescriptionEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "columnDescription", args)
 }

}


func (this *QAccessibleTableInterface) selectedCellCount(args ...interface{}) () {
  // selectedCellCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface17selectedCellCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedCellCount", args)
 }

}


func (this *QAccessibleTableInterface) selectedCells(args ...interface{}) () {
  // selectedCells()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface13selectedCellsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedCells", args)
 }

}


func (this *QAccessibleTableInterface) selectRow(args ...interface{}) () {
  // selectRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface9selectRowEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectRow", args)
 }

}


func (this *QAccessibleTableInterface) selectedRowCount(args ...interface{}) () {
  // selectedRowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface16selectedRowCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedRowCount", args)
 }

}


func (this *QAccessibleTableInterface) FreeQAccessibleTableInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "~QAccessibleTableInterface", args)
 }

}


func (this *QAccessibleTableInterface) selectedColumns(args ...interface{}) () {
  // selectedColumns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface15selectedColumnsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedColumns", args)
 }

}


func (this *QAccessibleTableInterface) cellAt(args ...interface{}) () {
  // cellAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface6cellAtEii
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "cellAt", args)
 }

}


func (this *QAccessibleTableInterface) selectedRows(args ...interface{}) () {
  // selectedRows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface12selectedRowsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedRows", args)
 }

}


func (this *QAccessibleTableInterface) modelChange(args ...interface{}) () {
  // modelChange(class QAccessibleTableModelChangeEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleTableModelChangeEvent{}) // "QAccessibleTableModelChangeEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface11modelChangeEP32QAccessibleTableModelChangeEvent
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "modelChange", args)
 }

}


func (this *QAccessibleTableInterface) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface11columnCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "columnCount", args)
 }

}


func (this *QAccessibleTableInterface) selectColumn(args ...interface{}) () {
  // selectColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface12selectColumnEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectColumn", args)
 }

}


func (this *QAccessibleTableInterface) unselectRow(args ...interface{}) () {
  // unselectRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface11unselectRowEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "unselectRow", args)
 }

}


func (this *QAccessibleTableInterface) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface8rowCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "rowCount", args)
 }

}


func (this *QAccessibleTableInterface) rowDescription(args ...interface{}) () {
  // rowDescription(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface14rowDescriptionEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "rowDescription", args)
 }

}


func (this *QAccessibleTableInterface) summary(args ...interface{}) () {
  // summary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface7summaryEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "summary", args)
 }

}


func (this *QAccessibleTableInterface) isColumnSelected(args ...interface{}) () {
  // isColumnSelected(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface16isColumnSelectedEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "isColumnSelected", args)
 }

}


func (this *QAccessibleTableInterface) caption(args ...interface{}) () {
  // caption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface7captionEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "caption", args)
 }

}


func (this *QAccessibleTableInterface) isRowSelected(args ...interface{}) () {
  // isRowSelected(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface13isRowSelectedEi
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "isRowSelected", args)
 }

}


func (this *QAccessibleTableInterface) selectedColumnCount(args ...interface{}) () {
  // selectedColumnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface19selectedColumnCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedColumnCount", args)
 }

}


func (this *QAccessibleTextUpdateEvent) textInserted(args ...interface{}) () {
  // textInserted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextUpdateEvent12textInsertedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textInserted", args)
 }

}


func NewQAccessibleTextUpdateEvent(args ...interface{})() {
}


func (this *QAccessibleTextUpdateEvent) textRemoved(args ...interface{}) () {
  // textRemoved()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextUpdateEvent11textRemovedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textRemoved", args)
 }

}


func (this *QAccessibleTextUpdateEvent) changePosition(args ...interface{}) () {
  // changePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextUpdateEvent14changePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "changePosition", args)
 }

}


func (this *QAccessibleImageInterface) imageDescription(args ...interface{}) () {
  // imageDescription()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleImageInterface16imageDescriptionEv
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "imageDescription", args)
 }

}


func (this *QAccessibleImageInterface) imagePosition(args ...interface{}) () {
  // imagePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleImageInterface13imagePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "imagePosition", args)
 }

}


func (this *QAccessibleImageInterface) FreeQAccessibleImageInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "~QAccessibleImageInterface", args)
 }

}


func (this *QAccessibleImageInterface) imageSize(args ...interface{}) () {
  // imageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleImageInterface9imageSizeEv
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "imageSize", args)
 }

}


func (this *QAccessibleTextInsertEvent) textInserted(args ...interface{}) () {
  // textInserted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextInsertEvent12textInsertedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "textInserted", args)
 }

}


func (this *QAccessibleTextInsertEvent) changePosition(args ...interface{}) () {
  // changePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextInsertEvent14changePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "changePosition", args)
 }

}


func NewQAccessibleTextInsertEvent(args ...interface{})() {
}


func (this *QAccessibleValueInterface) maximumValue(args ...interface{}) () {
  // maximumValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface12maximumValueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "maximumValue", args)
 }

}


func (this *QAccessibleValueInterface) minimumStepSize(args ...interface{}) () {
  // minimumStepSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface15minimumStepSizeEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "minimumStepSize", args)
 }

}


func (this *QAccessibleValueInterface) currentValue(args ...interface{}) () {
  // currentValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface12currentValueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "currentValue", args)
 }

}


func (this *QAccessibleValueInterface) minimumValue(args ...interface{}) () {
  // minimumValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface12minimumValueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "minimumValue", args)
 }

}


func (this *QAccessibleValueInterface) FreeQAccessibleValueInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "~QAccessibleValueInterface", args)
 }

}


func (this *QAccessibleValueInterface) setCurrentValue(args ...interface{}) () {
  // setCurrentValue(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleValueInterface15setCurrentValueERK8QVariant
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "setCurrentValue", args)
 }

}


func NewQAccessibleTextRemoveEvent(args ...interface{})() {
}


func (this *QAccessibleTextRemoveEvent) textRemoved(args ...interface{}) () {
  // textRemoved()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextRemoveEvent11textRemovedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "textRemoved", args)
 }

}


func (this *QAccessibleTextRemoveEvent) changePosition(args ...interface{}) () {
  // changePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextRemoveEvent14changePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "changePosition", args)
 }

}


func (this *QAccessibleTextSelectionEvent) selectionEnd(args ...interface{}) () {
  // selectionEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTextSelectionEvent12selectionEndEv
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionEnd", args)
 }

}


func NewQAccessibleTextSelectionEvent(args ...interface{})() {
}


func (this *QAccessibleTextSelectionEvent) selectionStart(args ...interface{}) () {
  // selectionStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTextSelectionEvent14selectionStartEv
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionStart", args)
 }

}


func (this *QAccessibleTextSelectionEvent) setSelection(args ...interface{}) () {
  // setSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN29QAccessibleTextSelectionEvent12setSelectionEii
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "setSelection", args)
 }

}


func NewQAccessibleTextCursorEvent(args ...interface{})() {
}


func (this *QAccessibleTextCursorEvent) setCursorPosition(args ...interface{}) () {
  // setCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextCursorEvent17setCursorPositionEi
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "setCursorPosition", args)
 }

}


func (this *QAccessibleTextCursorEvent) cursorPosition(args ...interface{}) () {
  // cursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextCursorEvent14cursorPositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "cursorPosition", args)
 }

}


func NewQAccessibleValueChangeEvent(args ...interface{})() {
}


func (this *QAccessibleValueChangeEvent) setValue(args ...interface{}) () {
  // setValue(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QAccessibleValueChangeEvent8setValueERK8QVariant
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "setValue", args)
 }

}


func (this *QAccessibleValueChangeEvent) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QAccessibleValueChangeEvent5valueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "value", args)
 }

}

// <= body block end

