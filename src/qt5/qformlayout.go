package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qformlayout.h
// dst-file: /src/widgets/qformlayout.go
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
// class sizeof(QFormLayout)=1
type QFormLayout struct {
  /*qbase*/ QLayout;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QFormLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout17horizontalSpacingEv
  default:
    qtrt.ErrorResolve("QFormLayout", "horizontalSpacing", args)
  }

}


func (this *QFormLayout) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout8rowCountEv
  default:
    qtrt.ErrorResolve("QFormLayout", "rowCount", args)
  }

}


func (this *QFormLayout) labelForField(args ...interface{}) () {
  // labelForField(class QLayout *)
  // labelForField(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout13labelForFieldEP7QLayout
  case 1:
    // invoke: _ZNK11QFormLayout13labelForFieldEP7QWidget
  default:
    qtrt.ErrorResolve("QFormLayout", "labelForField", args)
  }

}


func (this *QFormLayout) addRow(args ...interface{}) () {
  // addRow(const class QString &, class QLayout *)
  // addRow(class QLayout *)
  // addRow(const class QString &, class QWidget *)
  // addRow(class QWidget *)
  // addRow(class QWidget *, class QLayout *)
  // addRow(class QWidget *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[5][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout6addRowERK7QStringP7QLayout
  case 1:
    // invoke: _ZN11QFormLayout6addRowEP7QLayout
  case 2:
    // invoke: _ZN11QFormLayout6addRowERK7QStringP7QWidget
  case 3:
    // invoke: _ZN11QFormLayout6addRowEP7QWidget
  case 4:
    // invoke: _ZN11QFormLayout6addRowEP7QWidgetP7QLayout
  case 5:
    // invoke: _ZN11QFormLayout6addRowEP7QWidgetS1_
  default:
    qtrt.ErrorResolve("QFormLayout", "addRow", args)
  }

}


func (this *QFormLayout) insertRow(args ...interface{}) () {
  // insertRow(int, const class QString &, class QLayout *)
  // insertRow(int, class QWidget *, class QLayout *)
  // insertRow(int, class QLayout *)
  // insertRow(int, const class QString &, class QWidget *)
  // insertRow(int, class QWidget *, class QWidget *)
  // insertRow(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][2] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[4][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout9insertRowEiRK7QStringP7QLayout
  case 1:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout
  case 2:
    // invoke: _ZN11QFormLayout9insertRowEiP7QLayout
  case 3:
    // invoke: _ZN11QFormLayout9insertRowEiRK7QStringP7QWidget
  case 4:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidgetS1_
  case 5:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidget
  default:
    qtrt.ErrorResolve("QFormLayout", "insertRow", args)
  }

}


func (this *QFormLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout5countEv
  default:
    qtrt.ErrorResolve("QFormLayout", "count", args)
  }

}


func (this *QFormLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout7spacingEv
  default:
    qtrt.ErrorResolve("QFormLayout", "spacing", args)
  }

}


func NewQFormLayout(args ...interface{}) QFormLayout {
  return QFormLayout{}
}


func (this *QFormLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout11setGeometryERK5QRect
  default:
    qtrt.ErrorResolve("QFormLayout", "setGeometry", args)
  }

}


func (this *QFormLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout18setVerticalSpacingEi
  default:
    qtrt.ErrorResolve("QFormLayout", "setVerticalSpacing", args)
  }

}


func (this *QFormLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout20setHorizontalSpacingEi
  default:
    qtrt.ErrorResolve("QFormLayout", "setHorizontalSpacing", args)
  }

}


func (this *QFormLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout10metaObjectEv
  default:
    qtrt.ErrorResolve("QFormLayout", "metaObject", args)
  }

}


func (this *QFormLayout) setSpacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout10setSpacingEi
  default:
    qtrt.ErrorResolve("QFormLayout", "setSpacing", args)
  }

}


func (this *QFormLayout) FreeQFormLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFormLayout", "~QFormLayout", args)
  }

}


func (this *QFormLayout) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout8sizeHintEv
  default:
    qtrt.ErrorResolve("QFormLayout", "sizeHint", args)
  }

}


func (this *QFormLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QFormLayout", "invalidate", args)
  }

}


func (this *QFormLayout) itemAt(args ...interface{}) () {
  // itemAt(int, enum QFormLayout::ItemRole)
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "QFormLayout::ItemRole"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout6itemAtEiNS_8ItemRoleE
  case 1:
    // invoke: _ZNK11QFormLayout6itemAtEi
  default:
    qtrt.ErrorResolve("QFormLayout", "itemAt", args)
  }

}


func (this *QFormLayout) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout6takeAtEi
  default:
    qtrt.ErrorResolve("QFormLayout", "takeAt", args)
  }

}


func (this *QFormLayout) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout11minimumSizeEv
  default:
    qtrt.ErrorResolve("QFormLayout", "minimumSize", args)
  }

}


func (this *QFormLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout15verticalSpacingEv
  default:
    qtrt.ErrorResolve("QFormLayout", "verticalSpacing", args)
  }

}


func (this *QFormLayout) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout14heightForWidthEi
  default:
    qtrt.ErrorResolve("QFormLayout", "heightForWidth", args)
  }

}


func (this *QFormLayout) addItem(args ...interface{}) () {
  // addItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout7addItemEP11QLayoutItem
  default:
    qtrt.ErrorResolve("QFormLayout", "addItem", args)
  }

}


func (this *QFormLayout) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout17hasHeightForWidthEv
  default:
    qtrt.ErrorResolve("QFormLayout", "hasHeightForWidth", args)
  }

}

// <= body block end

