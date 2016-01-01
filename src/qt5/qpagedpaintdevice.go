package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpagedpaintdevice.h
// dst-file: /src/gui/qpagedpaintdevice.go
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
// class sizeof(QPagedPaintDevice)=32
type QPagedPaintDevice struct {
  /*qbase*/ QPaintDevice;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPagedPaintDevice) pageSizeMM(args ...interface{}) () {
  // pageSizeMM()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPagedPaintDevice10pageSizeMMEv
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "pageSizeMM", args)
  }

}


func (this *QPagedPaintDevice) FreeQPagedPaintDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "~QPagedPaintDevice", args)
  }

}


func (this *QPagedPaintDevice) setPageMargins(args ...interface{}) () {
  // setPageMargins(const class QMarginsF &)
  // setPageMargins(const class QMarginsF &, class QPageLayout::Unit)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "QPageLayout::Unit"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF
  case 1:
    // invoke: _ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsFN11QPageLayout4UnitE
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageMargins", args)
  }

}


func (this *QPagedPaintDevice) pageLayout(args ...interface{}) () {
  // pageLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPagedPaintDevice10pageLayoutEv
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "pageLayout", args)
  }

}


func (this *QPagedPaintDevice) setPageSize(args ...interface{}) () {
  // setPageSize(enum QPagedPaintDevice::PageSize)
  // setPageSize(const class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QPagedPaintDevice::PageSize"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDevice11setPageSizeENS_8PageSizeE
  case 1:
    // invoke: _ZN17QPagedPaintDevice11setPageSizeERK9QPageSize
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageSize", args)
  }

}


func NewQPagedPaintDevice(args ...interface{}) QPagedPaintDevice {
  return QPagedPaintDevice{}
}


func (this *QPagedPaintDevice) setPageSizeMM(args ...interface{}) () {
  // setPageSizeMM(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDevice13setPageSizeMMERK6QSizeF
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageSizeMM", args)
  }

}


func (this *QPagedPaintDevice) setPageLayout(args ...interface{}) () {
  // setPageLayout(const class QPageLayout &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageLayout{}) // "const QPageLayout &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDevice13setPageLayoutERK11QPageLayout
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageLayout", args)
  }

}


func (this *QPagedPaintDevice) newPage(args ...interface{}) () {
  // newPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDevice7newPageEv
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "newPage", args)
  }

}

// <= body block end

