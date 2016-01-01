package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qevent.h
// dst-file: /src/gui/qevent.go
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
// class sizeof(QWhatsThisClickedEvent)=32
type QWhatsThisClickedEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QExposeEvent)=32
type QExposeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QInputMethodEvent)=1
type QInputMethodEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHelpEvent)=40
type QHelpEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QActionEvent)=40
type QActionEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMouseEvent)=1
type QMouseEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFileOpenEvent)=40
type QFileOpenEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QToolBarChangeEvent)=24
type QToolBarChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTabletEvent)=1
type QTabletEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTouchEvent)=1
type QTouchEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QScreenOrientationChangeEvent)=40
type QScreenOrientationChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QIconDragEvent)=24
type QIconDragEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QCloseEvent)=24
type QCloseEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDragEnterEvent)=1
type QDragEnterEvent struct {
  /*qbase*/ QDragMoveEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QWheelEvent)=1
type QWheelEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QScrollEvent)=64
type QScrollEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHoverEvent)=1
type QHoverEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDragMoveEvent)=1
type QDragMoveEvent struct {
  /*qbase*/ QDropEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QShowEvent)=24
type QShowEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPlatformSurfaceEvent)=24
type QPlatformSurfaceEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPaintEvent)=56
type QPaintEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFocusEvent)=24
type QFocusEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QNativeGestureEvent)=1
type QNativeGestureEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QResizeEvent)=40
type QResizeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStatusTipEvent)=32
type QStatusTipEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QEnterEvent)=72
type QEnterEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMoveEvent)=40
type QMoveEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHideEvent)=24
type QHideEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDragLeaveEvent)=24
type QDragLeaveEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDropEvent)=1
type QDropEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QInputEvent)=1
type QInputEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QApplicationStateChangeEvent)=24
type QApplicationStateChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QKeyEvent)=1
type QKeyEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QContextMenuEvent)=1
type QContextMenuEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QScrollPrepareEvent)=112
type QScrollPrepareEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QShortcutEvent)=40
type QShortcutEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QWindowStateChangeEvent)=1
type QWindowStateChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QInputMethodQueryEvent)=1
type QInputMethodQueryEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QWhatsThisClickedEvent) href(args ...interface{}) () {
  // href()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QWhatsThisClickedEvent4hrefEv
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "href", args)
  }

}


func (this *QWhatsThisClickedEvent) FreeQWhatsThisClickedEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "~QWhatsThisClickedEvent", args)
  }

}


func NewQWhatsThisClickedEvent(args ...interface{}) QWhatsThisClickedEvent {
  return QWhatsThisClickedEvent{}
}


func NewQExposeEvent(args ...interface{}) QExposeEvent {
  return QExposeEvent{}
}


func (this *QExposeEvent) region(args ...interface{}) () {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QExposeEvent6regionEv
  default:
    qtrt.ErrorResolve("QExposeEvent", "region", args)
  }

}


func (this *QExposeEvent) FreeQExposeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QExposeEvent", "~QExposeEvent", args)
  }

}


func (this *QInputMethodEvent) preeditString(args ...interface{}) () {
  // preeditString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent13preeditStringEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "preeditString", args)
  }

}


func NewQInputMethodEvent(args ...interface{}) QInputMethodEvent {
  return QInputMethodEvent{}
}


func (this *QInputMethodEvent) replacementStart(args ...interface{}) () {
  // replacementStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent16replacementStartEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementStart", args)
  }

}


func (this *QInputMethodEvent) commitString(args ...interface{}) () {
  // commitString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent12commitStringEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "commitString", args)
  }

}


func (this *QInputMethodEvent) setCommitString(args ...interface{}) () {
  // setCommitString(const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QInputMethodEvent15setCommitStringERK7QStringii
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "setCommitString", args)
  }

}


func (this *QInputMethodEvent) replacementLength(args ...interface{}) () {
  // replacementLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent17replacementLengthEv
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementLength", args)
  }

}


func (this *QHelpEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalPos", args)
  }

}


func (this *QHelpEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalXEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalX", args)
  }

}


func (this *QHelpEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent3posEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "pos", args)
  }

}


func (this *QHelpEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1yEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "y", args)
  }

}


func (this *QHelpEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalYEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalY", args)
  }

}


func (this *QHelpEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1xEv
  default:
    qtrt.ErrorResolve("QHelpEvent", "x", args)
  }

}


func (this *QHelpEvent) FreeQHelpEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHelpEvent", "~QHelpEvent", args)
  }

}


func NewQActionEvent(args ...interface{}) QActionEvent {
  return QActionEvent{}
}


func (this *QActionEvent) before(args ...interface{}) () {
  // before()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionEvent6beforeEv
  default:
    qtrt.ErrorResolve("QActionEvent", "before", args)
  }

}


func (this *QActionEvent) action(args ...interface{}) () {
  // action()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionEvent6actionEv
  default:
    qtrt.ErrorResolve("QActionEvent", "action", args)
  }

}


func (this *QActionEvent) FreeQActionEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QActionEvent", "~QActionEvent", args)
  }

}


func (this *QMouseEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalPos", args)
  }

}


func (this *QMouseEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1yEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "y", args)
  }

}


func (this *QMouseEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "screenPos", args)
  }

}


func (this *QMouseEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1xEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "x", args)
  }

}


func (this *QMouseEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent8localPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "localPos", args)
  }

}


func (this *QMouseEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalXEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalX", args)
  }

}


func (this *QMouseEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9windowPosEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "windowPos", args)
  }

}


func (this *QMouseEvent) FreeQMouseEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMouseEvent", "~QMouseEvent", args)
  }

}


func (this *QMouseEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalYEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalY", args)
  }

}


func (this *QMouseEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent3posEv
  default:
    qtrt.ErrorResolve("QMouseEvent", "pos", args)
  }

}


func NewQFileOpenEvent(args ...interface{}) QFileOpenEvent {
  return QFileOpenEvent{}
}


func (this *QFileOpenEvent) FreeQFileOpenEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "~QFileOpenEvent", args)
  }

}


func (this *QFileOpenEvent) file(args ...interface{}) () {
  // file()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent4fileEv
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "file", args)
  }

}


func (this *QFileOpenEvent) url(args ...interface{}) () {
  // url()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent3urlEv
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "url", args)
  }

}


func NewQToolBarChangeEvent(args ...interface{}) QToolBarChangeEvent {
  return QToolBarChangeEvent{}
}


func (this *QToolBarChangeEvent) FreeQToolBarChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "~QToolBarChangeEvent", args)
  }

}


func (this *QToolBarChangeEvent) toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QToolBarChangeEvent6toggleEv
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "toggle", args)
  }

}


func (this *QTabletEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1xEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "x", args)
  }

}


func (this *QTabletEvent) xTilt(args ...interface{}) () {
  // xTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5xTiltEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "xTilt", args)
  }

}


func (this *QTabletEvent) uniqueId(args ...interface{}) () {
  // uniqueId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8uniqueIdEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "uniqueId", args)
  }

}


func (this *QTabletEvent) globalPosF(args ...interface{}) () {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent10globalPosFEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPosF", args)
  }

}


func (this *QTabletEvent) z(args ...interface{}) () {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1zEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "z", args)
  }

}


func (this *QTabletEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1yEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "y", args)
  }

}


func (this *QTabletEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent3posEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "pos", args)
  }

}


func (this *QTabletEvent) rotation(args ...interface{}) () {
  // rotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8rotationEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "rotation", args)
  }

}


func (this *QTabletEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPos", args)
  }

}


func (this *QTabletEvent) FreeQTabletEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTabletEvent", "~QTabletEvent", args)
  }

}


func (this *QTabletEvent) tangentialPressure(args ...interface{}) () {
  // tangentialPressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent18tangentialPressureEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "tangentialPressure", args)
  }

}


func (this *QTabletEvent) hiResGlobalX(args ...interface{}) () {
  // hiResGlobalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalXEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalX", args)
  }

}


func (this *QTabletEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalYEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalY", args)
  }

}


func (this *QTabletEvent) hiResGlobalY(args ...interface{}) () {
  // hiResGlobalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalYEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalY", args)
  }

}


func (this *QTabletEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalXEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalX", args)
  }

}


func (this *QTabletEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent4posFEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "posF", args)
  }

}


func (this *QTabletEvent) pressure(args ...interface{}) () {
  // pressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8pressureEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "pressure", args)
  }

}


func (this *QTabletEvent) yTilt(args ...interface{}) () {
  // yTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5yTiltEv
  default:
    qtrt.ErrorResolve("QTabletEvent", "yTilt", args)
  }

}


func (this *QTouchEvent) setDevice(args ...interface{}) () {
  // setDevice(class QTouchDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTouchDevice{}) // "QTouchDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setDeviceEP12QTouchDevice
  default:
    qtrt.ErrorResolve("QTouchEvent", "setDevice", args)
  }

}


func (this *QTouchEvent) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6windowEv
  default:
    qtrt.ErrorResolve("QTouchEvent", "window", args)
  }

}


func (this *QTouchEvent) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6deviceEv
  default:
    qtrt.ErrorResolve("QTouchEvent", "device", args)
  }

}


func (this *QTouchEvent) target(args ...interface{}) () {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6targetEv
  default:
    qtrt.ErrorResolve("QTouchEvent", "target", args)
  }

}


func (this *QTouchEvent) FreeQTouchEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTouchEvent", "~QTouchEvent", args)
  }

}


func (this *QTouchEvent) setWindow(args ...interface{}) () {
  // setWindow(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setWindowEP7QWindow
  default:
    qtrt.ErrorResolve("QTouchEvent", "setWindow", args)
  }

}


func (this *QTouchEvent) setTarget(args ...interface{}) () {
  // setTarget(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setTargetEP7QObject
  default:
    qtrt.ErrorResolve("QTouchEvent", "setTarget", args)
  }

}


func (this *QScreenOrientationChangeEvent) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QScreenOrientationChangeEvent6screenEv
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "screen", args)
  }

}


func (this *QScreenOrientationChangeEvent) FreeQScreenOrientationChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "~QScreenOrientationChangeEvent", args)
  }

}


func (this *QIconDragEvent) FreeQIconDragEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIconDragEvent", "~QIconDragEvent", args)
  }

}


func NewQIconDragEvent(args ...interface{}) QIconDragEvent {
  return QIconDragEvent{}
}


func (this *QCloseEvent) FreeQCloseEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCloseEvent", "~QCloseEvent", args)
  }

}


func NewQCloseEvent(args ...interface{}) QCloseEvent {
  return QCloseEvent{}
}


func (this *QDragEnterEvent) FreeQDragEnterEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDragEnterEvent", "~QDragEnterEvent", args)
  }

}


func (this *QWheelEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1xEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "x", args)
  }

}


func (this *QWheelEvent) angleDelta(args ...interface{}) () {
  // angleDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10angleDeltaEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "angleDelta", args)
  }

}


func (this *QWheelEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent3posEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "pos", args)
  }

}


func (this *QWheelEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalYEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalY", args)
  }

}


func (this *QWheelEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent4posFEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "posF", args)
  }

}


func (this *QWheelEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalXEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalX", args)
  }

}


func (this *QWheelEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1yEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "y", args)
  }

}


func (this *QWheelEvent) FreeQWheelEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWheelEvent", "~QWheelEvent", args)
  }

}


func (this *QWheelEvent) pixelDelta(args ...interface{}) () {
  // pixelDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10pixelDeltaEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "pixelDelta", args)
  }

}


func (this *QWheelEvent) delta(args ...interface{}) () {
  // delta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent5deltaEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "delta", args)
  }

}


func (this *QWheelEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPos", args)
  }

}


func (this *QWheelEvent) globalPosF(args ...interface{}) () {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10globalPosFEv
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPosF", args)
  }

}


func (this *QScrollEvent) contentPos(args ...interface{}) () {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent10contentPosEv
  default:
    qtrt.ErrorResolve("QScrollEvent", "contentPos", args)
  }

}


func (this *QScrollEvent) overshootDistance(args ...interface{}) () {
  // overshootDistance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent17overshootDistanceEv
  default:
    qtrt.ErrorResolve("QScrollEvent", "overshootDistance", args)
  }

}


func (this *QScrollEvent) FreeQScrollEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollEvent", "~QScrollEvent", args)
  }

}


func (this *QHoverEvent) FreeQHoverEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHoverEvent", "~QHoverEvent", args)
  }

}


func (this *QHoverEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent4posFEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "posF", args)
  }

}


func (this *QHoverEvent) oldPos(args ...interface{}) () {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent6oldPosEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPos", args)
  }

}


func (this *QHoverEvent) oldPosF(args ...interface{}) () {
  // oldPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent7oldPosFEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPosF", args)
  }

}


func (this *QHoverEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent3posEv
  default:
    qtrt.ErrorResolve("QHoverEvent", "pos", args)
  }

}


func (this *QDragMoveEvent) accept(args ...interface{}) () {
  // accept(const class QRect &)
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6acceptERK5QRect
  case 1:
    // invoke: _ZN14QDragMoveEvent6acceptEv
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "accept", args)
  }

}


func (this *QDragMoveEvent) answerRect(args ...interface{}) () {
  // answerRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDragMoveEvent10answerRectEv
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "answerRect", args)
  }

}


func (this *QDragMoveEvent) ignore(args ...interface{}) () {
  // ignore(const class QRect &)
  // ignore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6ignoreERK5QRect
  case 1:
    // invoke: _ZN14QDragMoveEvent6ignoreEv
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "ignore", args)
  }

}


func (this *QDragMoveEvent) FreeQDragMoveEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "~QDragMoveEvent", args)
  }

}


func (this *QShowEvent) FreeQShowEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QShowEvent", "~QShowEvent", args)
  }

}


func NewQShowEvent(args ...interface{}) QShowEvent {
  return QShowEvent{}
}


func (this *QPlatformSurfaceEvent) FreeQPlatformSurfaceEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPlatformSurfaceEvent", "~QPlatformSurfaceEvent", args)
  }

}


func (this *QPaintEvent) FreeQPaintEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPaintEvent", "~QPaintEvent", args)
  }

}


func (this *QPaintEvent) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent4rectEv
  default:
    qtrt.ErrorResolve("QPaintEvent", "rect", args)
  }

}


func NewQPaintEvent(args ...interface{}) QPaintEvent {
  return QPaintEvent{}
}


func (this *QPaintEvent) region(args ...interface{}) () {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent6regionEv
  default:
    qtrt.ErrorResolve("QPaintEvent", "region", args)
  }

}


func (this *QFocusEvent) lostFocus(args ...interface{}) () {
  // lostFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent9lostFocusEv
  default:
    qtrt.ErrorResolve("QFocusEvent", "lostFocus", args)
  }

}


func (this *QFocusEvent) gotFocus(args ...interface{}) () {
  // gotFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent8gotFocusEv
  default:
    qtrt.ErrorResolve("QFocusEvent", "gotFocus", args)
  }

}


func (this *QFocusEvent) FreeQFocusEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFocusEvent", "~QFocusEvent", args)
  }

}


func (this *QNativeGestureEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent8localPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "localPos", args)
  }

}


func (this *QNativeGestureEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "screenPos", args)
  }

}


func (this *QNativeGestureEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent3posEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "pos", args)
  }

}


func (this *QNativeGestureEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "globalPos", args)
  }

}


func (this *QNativeGestureEvent) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent5valueEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "value", args)
  }

}


func (this *QNativeGestureEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9windowPosEv
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "windowPos", args)
  }

}


func (this *QResizeEvent) oldSize(args ...interface{}) () {
  // oldSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent7oldSizeEv
  default:
    qtrt.ErrorResolve("QResizeEvent", "oldSize", args)
  }

}


func (this *QResizeEvent) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent4sizeEv
  default:
    qtrt.ErrorResolve("QResizeEvent", "size", args)
  }

}


func (this *QResizeEvent) FreeQResizeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QResizeEvent", "~QResizeEvent", args)
  }

}


func NewQResizeEvent(args ...interface{}) QResizeEvent {
  return QResizeEvent{}
}


func (this *QStatusTipEvent) FreeQStatusTipEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "~QStatusTipEvent", args)
  }

}


func (this *QStatusTipEvent) tip(args ...interface{}) () {
  // tip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QStatusTipEvent3tipEv
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "tip", args)
  }

}


func NewQStatusTipEvent(args ...interface{}) QStatusTipEvent {
  return QStatusTipEvent{}
}


func (this *QEnterEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1yEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "y", args)
  }

}


func (this *QEnterEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent3posEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "pos", args)
  }

}


func (this *QEnterEvent) FreeQEnterEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEnterEvent", "~QEnterEvent", args)
  }

}


func (this *QEnterEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9screenPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "screenPos", args)
  }

}


func (this *QEnterEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent8localPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "localPos", args)
  }

}


func (this *QEnterEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9windowPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "windowPos", args)
  }

}


func (this *QEnterEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalXEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalX", args)
  }

}


func (this *QEnterEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1xEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "x", args)
  }

}


func (this *QEnterEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalPos", args)
  }

}


func (this *QEnterEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalYEv
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalY", args)
  }

}


func NewQEnterEvent(args ...interface{}) QEnterEvent {
  return QEnterEvent{}
}


func (this *QMoveEvent) FreeQMoveEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMoveEvent", "~QMoveEvent", args)
  }

}


func (this *QMoveEvent) oldPos(args ...interface{}) () {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent6oldPosEv
  default:
    qtrt.ErrorResolve("QMoveEvent", "oldPos", args)
  }

}


func NewQMoveEvent(args ...interface{}) QMoveEvent {
  return QMoveEvent{}
}


func (this *QMoveEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent3posEv
  default:
    qtrt.ErrorResolve("QMoveEvent", "pos", args)
  }

}


func NewQHideEvent(args ...interface{}) QHideEvent {
  return QHideEvent{}
}


func (this *QHideEvent) FreeQHideEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHideEvent", "~QHideEvent", args)
  }

}


func (this *QDragLeaveEvent) FreeQDragLeaveEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDragLeaveEvent", "~QDragLeaveEvent", args)
  }

}


func NewQDragLeaveEvent(args ...interface{}) QDragLeaveEvent {
  return QDragLeaveEvent{}
}


func (this *QDropEvent) FreeQDropEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDropEvent", "~QDropEvent", args)
  }

}


func (this *QDropEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent3posEv
  default:
    qtrt.ErrorResolve("QDropEvent", "pos", args)
  }

}


func (this *QDropEvent) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent6sourceEv
  default:
    qtrt.ErrorResolve("QDropEvent", "source", args)
  }

}


func (this *QDropEvent) mimeData(args ...interface{}) () {
  // mimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent8mimeDataEv
  default:
    qtrt.ErrorResolve("QDropEvent", "mimeData", args)
  }

}


func (this *QDropEvent) acceptProposedAction(args ...interface{}) () {
  // acceptProposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QDropEvent20acceptProposedActionEv
  default:
    qtrt.ErrorResolve("QDropEvent", "acceptProposedAction", args)
  }

}


func (this *QDropEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent4posFEv
  default:
    qtrt.ErrorResolve("QDropEvent", "posF", args)
  }

}


func (this *QInputEvent) setTimestamp(args ...interface{}) () {
  // setTimestamp(ulong)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "ulong"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QInputEvent12setTimestampEm
  default:
    qtrt.ErrorResolve("QInputEvent", "setTimestamp", args)
  }

}


func (this *QInputEvent) timestamp(args ...interface{}) () {
  // timestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QInputEvent9timestampEv
  default:
    qtrt.ErrorResolve("QInputEvent", "timestamp", args)
  }

}


func (this *QInputEvent) FreeQInputEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QInputEvent", "~QInputEvent", args)
  }

}


func (this *QKeyEvent) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent5countEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "count", args)
  }

}


func (this *QKeyEvent) FreeQKeyEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeyEvent", "~QKeyEvent", args)
  }

}


func (this *QKeyEvent) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent4textEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "text", args)
  }

}


func (this *QKeyEvent) nativeVirtualKey(args ...interface{}) () {
  // nativeVirtualKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent16nativeVirtualKeyEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeVirtualKey", args)
  }

}


func (this *QKeyEvent) isAutoRepeat(args ...interface{}) () {
  // isAutoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent12isAutoRepeatEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "isAutoRepeat", args)
  }

}


func (this *QKeyEvent) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent3keyEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "key", args)
  }

}


func (this *QKeyEvent) nativeModifiers(args ...interface{}) () {
  // nativeModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent15nativeModifiersEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeModifiers", args)
  }

}


func (this *QKeyEvent) nativeScanCode(args ...interface{}) () {
  // nativeScanCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent14nativeScanCodeEv
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeScanCode", args)
  }

}


func (this *QContextMenuEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent9globalPosEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalPos", args)
  }

}


func (this *QContextMenuEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalYEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalY", args)
  }

}


func (this *QContextMenuEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalXEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalX", args)
  }

}


func (this *QContextMenuEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent3posEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "pos", args)
  }

}


func (this *QContextMenuEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1yEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "y", args)
  }

}


func (this *QContextMenuEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1xEv
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "x", args)
  }

}


func (this *QContextMenuEvent) FreeQContextMenuEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "~QContextMenuEvent", args)
  }

}


func (this *QScrollPrepareEvent) setContentPosRange(args ...interface{}) () {
  // setContentPosRange(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPosRange", args)
  }

}


func (this *QScrollPrepareEvent) setContentPos(args ...interface{}) () {
  // setContentPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent13setContentPosERK7QPointF
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPos", args)
  }

}


func (this *QScrollPrepareEvent) contentPosRange(args ...interface{}) () {
  // contentPosRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent15contentPosRangeEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPosRange", args)
  }

}


func (this *QScrollPrepareEvent) contentPos(args ...interface{}) () {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent10contentPosEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPos", args)
  }

}


func (this *QScrollPrepareEvent) setViewportSize(args ...interface{}) () {
  // setViewportSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setViewportSize", args)
  }

}


func NewQScrollPrepareEvent(args ...interface{}) QScrollPrepareEvent {
  return QScrollPrepareEvent{}
}


func (this *QScrollPrepareEvent) startPos(args ...interface{}) () {
  // startPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent8startPosEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "startPos", args)
  }

}


func (this *QScrollPrepareEvent) viewportSize(args ...interface{}) () {
  // viewportSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent12viewportSizeEv
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "viewportSize", args)
  }

}


func (this *QScrollPrepareEvent) FreeQScrollPrepareEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "~QScrollPrepareEvent", args)
  }

}


func (this *QShortcutEvent) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent3keyEv
  default:
    qtrt.ErrorResolve("QShortcutEvent", "key", args)
  }

}


func (this *QShortcutEvent) FreeQShortcutEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QShortcutEvent", "~QShortcutEvent", args)
  }

}


func (this *QShortcutEvent) isAmbiguous(args ...interface{}) () {
  // isAmbiguous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent11isAmbiguousEv
  default:
    qtrt.ErrorResolve("QShortcutEvent", "isAmbiguous", args)
  }

}


func NewQShortcutEvent(args ...interface{}) QShortcutEvent {
  return QShortcutEvent{}
}


func (this *QShortcutEvent) shortcutId(args ...interface{}) () {
  // shortcutId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent10shortcutIdEv
  default:
    qtrt.ErrorResolve("QShortcutEvent", "shortcutId", args)
  }

}


func (this *QWindowStateChangeEvent) isOverride(args ...interface{}) () {
  // isOverride()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QWindowStateChangeEvent10isOverrideEv
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "isOverride", args)
  }

}


func (this *QWindowStateChangeEvent) FreeQWindowStateChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "~QWindowStateChangeEvent", args)
  }

}


func (this *QInputMethodQueryEvent) FreeQInputMethodQueryEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QInputMethodQueryEvent", "~QInputMethodQueryEvent", args)
  }

}

// <= body block end

