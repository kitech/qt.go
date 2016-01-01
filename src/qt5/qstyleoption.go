package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qstyleoption.h
// dst-file: /src/widgets/qstyleoption.go
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
// class sizeof(QStyleOptionComboBox)=1
type QStyleOptionComboBox struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionMenuItem)=1
type QStyleOptionMenuItem struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleHintReturnVariant)=24
type QStyleHintReturnVariant struct {
  /*qbase*/ QStyleHintReturn;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionTitleBar)=1
type QStyleOptionTitleBar struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionGraphicsItem)=1
type QStyleOptionGraphicsItem struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOption)=1
type QStyleOption struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionDockWidget)=1
type QStyleOptionDockWidget struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionProgressBar)=1
type QStyleOptionProgressBar struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionSlider)=1
type QStyleOptionSlider struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionFrame)=1
type QStyleOptionFrame struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionComplex)=1
type QStyleOptionComplex struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleHintReturn)=8
type QStyleHintReturn struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionHeader)=1
type QStyleOptionHeader struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionToolBox)=1
type QStyleOptionToolBox struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionFocusRect)=1
type QStyleOptionFocusRect struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionGroupBox)=1
type QStyleOptionGroupBox struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionTab)=1
type QStyleOptionTab struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionTabBarBase)=1
type QStyleOptionTabBarBase struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionRubberBand)=1
type QStyleOptionRubberBand struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionButton)=1
type QStyleOptionButton struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleHintReturnMask)=16
type QStyleHintReturnMask struct {
  /*qbase*/ QStyleHintReturn;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionToolButton)=1
type QStyleOptionToolButton struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionSizeGrip)=1
type QStyleOptionSizeGrip struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionViewItem)=1
type QStyleOptionViewItem struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionSpinBox)=1
type QStyleOptionSpinBox struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionToolBar)=1
type QStyleOptionToolBar struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStyleOptionTabWidgetFrame)=1
type QStyleOptionTabWidgetFrame struct {
  /*qbase*/ QStyleOption;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQStyleOptionComboBox(args ...interface{})() {
}


func NewQStyleOptionMenuItem(args ...interface{})() {
}


func (this *QStyleHintReturnVariant) FreeQStyleHintReturnVariant(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleHintReturnVariant", "~QStyleHintReturnVariant", args)
 }

}


func NewQStyleHintReturnVariant(args ...interface{})() {
}


func NewQStyleOptionTitleBar(args ...interface{})() {
}


func NewQStyleOptionGraphicsItem(args ...interface{})() {
}


func (this *QStyleOptionGraphicsItem) levelOfDetailFromTransform_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleOptionGraphicsItem", "levelOfDetailFromTransform", args)
 }

}


func (this *QStyleOption) FreeQStyleOption(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleOption", "~QStyleOption", args)
 }

}


func (this *QStyleOption) init(args ...interface{}) () {
  // init(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QStyleOption4initEPK7QWidget
  default:
    qtrt.ErrorResolve("QStyleOption", "init", args)
 }

}


func NewQStyleOption(args ...interface{})() {
}


func (this *QStyleOption) initFrom(args ...interface{}) () {
  // initFrom(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QStyleOption8initFromEPK7QWidget
  default:
    qtrt.ErrorResolve("QStyleOption", "initFrom", args)
 }

}


func NewQStyleOptionDockWidget(args ...interface{})() {
}


func NewQStyleOptionProgressBar(args ...interface{})() {
}


func NewQStyleOptionSlider(args ...interface{})() {
}


func NewQStyleOptionFrame(args ...interface{})() {
}


func NewQStyleOptionComplex(args ...interface{})() {
}


func (this *QStyleHintReturn) FreeQStyleHintReturn(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleHintReturn", "~QStyleHintReturn", args)
 }

}


func NewQStyleHintReturn(args ...interface{})() {
}


func NewQStyleOptionHeader(args ...interface{})() {
}


func NewQStyleOptionToolBox(args ...interface{})() {
}


func NewQStyleOptionFocusRect(args ...interface{})() {
}


func NewQStyleOptionGroupBox(args ...interface{})() {
}


func NewQStyleOptionTab(args ...interface{})() {
}


func NewQStyleOptionTabBarBase(args ...interface{})() {
}


func NewQStyleOptionRubberBand(args ...interface{})() {
}


func NewQStyleOptionButton(args ...interface{})() {
}


func NewQStyleHintReturnMask(args ...interface{})() {
}


func (this *QStyleHintReturnMask) FreeQStyleHintReturnMask(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleHintReturnMask", "~QStyleHintReturnMask", args)
 }

}


func NewQStyleOptionToolButton(args ...interface{})() {
}


func NewQStyleOptionSizeGrip(args ...interface{})() {
}


func NewQStyleOptionViewItem(args ...interface{})() {
}


func NewQStyleOptionSpinBox(args ...interface{})() {
}


func NewQStyleOptionToolBar(args ...interface{})() {
}


func NewQStyleOptionTabWidgetFrame(args ...interface{})() {
}

// <= body block end

