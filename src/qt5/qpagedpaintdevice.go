package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qpagedpaintdevice.h
// dst-file: /src/gui/qpagedpaintdevice.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QPagedPaintDevice::setPageMargins(const QMarginsF & margins);
extern void C_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF(void* qthis, void* arg0); // 4
  // proto:  QPageLayout QPagedPaintDevice::pageLayout();
extern void C_ZNK17QPagedPaintDevice10pageLayoutEv(void* qthis); // 4
  // proto:  bool QPagedPaintDevice::setPageSize(const QPageSize & pageSize);
extern void C_ZN17QPagedPaintDevice11setPageSizeERK9QPageSize(void* qthis, void* arg0); // 4
  // proto:  bool QPagedPaintDevice::setPageLayout(const QPageLayout & pageLayout);
extern void C_ZN17QPagedPaintDevice13setPageLayoutERK11QPageLayout(void* qthis, void* arg0); // 4
  // proto:  void QPagedPaintDevice::QPagedPaintDevice();
extern void C_ZN17QPagedPaintDeviceC2Ev(void* qthis); // 3
  // proto:  void QPagedPaintDevice::~QPagedPaintDevice();
extern void C_ZN17QPagedPaintDeviceD2Ev(void* qthis); // 4
  // proto:  QSizeF QPagedPaintDevice::pageSizeMM();
extern void C_ZNK17QPagedPaintDevice10pageSizeMMEv(void* qthis); // 4
  // proto:  void QPagedPaintDevice::setPageSizeMM(const QSizeF & size);
extern void C_ZN17QPagedPaintDevice13setPageSizeMMERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  QPagedPaintDevice::Margins QPagedPaintDevice::margins();
extern void C_ZNK17QPagedPaintDevice7marginsEv(void* qthis); // 4
  // proto:  QPagedPaintDevice::PageSize QPagedPaintDevice::pageSize();
extern void C_ZNK17QPagedPaintDevice8pageSizeEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPagedPaintDevice)=32
type QPagedPaintDevice struct {
  /*qbase*/ QPaintDevice;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setPageMargins(const class QMarginsF &)
func (this *QPagedPaintDevice) setPageMargins(args ...interface{}) () {
  // setPageMargins(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF
    // invoke: bool setPageMargins(const class QMarginsF &)
    var arg0 = args[0].(QMarginsF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageMargins", args)
  }

}

// pageLayout()
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
    // invoke: QPageLayout pageLayout()
    C.C_ZNK17QPagedPaintDevice10pageLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "pageLayout", args)
  }

}

// setPageSize(const class QPageSize &)
func (this *QPagedPaintDevice) setPageSize(args ...interface{}) () {
  // setPageSize(const class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDevice11setPageSizeERK9QPageSize
    // invoke: bool setPageSize(const class QPageSize &)
    var arg0 = args[0].(QPageSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QPagedPaintDevice11setPageSizeERK9QPageSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageSize", args)
  }

}

// setPageLayout(const class QPageLayout &)
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
    // invoke: bool setPageLayout(const class QPageLayout &)
    var arg0 = args[0].(QPageLayout).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QPagedPaintDevice13setPageLayoutERK11QPageLayout(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageLayout", args)
  }

}

// QPagedPaintDevice()
func NewQPagedPaintDevice(args ...interface{}) QPagedPaintDevice {
  // QPagedPaintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDeviceC1Ev
    // invoke: void QPagedPaintDevice()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN17QPagedPaintDeviceC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "QPagedPaintDevice", args)
  }

  return QPagedPaintDevice{}
}

// ~QPagedPaintDevice()
func (this *QPagedPaintDevice) FreeQPagedPaintDevice(args ...interface{}) () {
  // ~QPagedPaintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QPagedPaintDeviceD0Ev
    // invoke: void ~QPagedPaintDevice()
    C.C_ZN17QPagedPaintDeviceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "~QPagedPaintDevice", args)
  }

}

// pageSizeMM()
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
    // invoke: QSizeF pageSizeMM()
    C.C_ZNK17QPagedPaintDevice10pageSizeMMEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "pageSizeMM", args)
  }

}

// setPageSizeMM(const class QSizeF &)
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
    // invoke: void setPageSizeMM(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QPagedPaintDevice13setPageSizeMMERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "setPageSizeMM", args)
  }

}

// margins()
func (this *QPagedPaintDevice) margins(args ...interface{}) () {
  // margins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPagedPaintDevice7marginsEv
    // invoke: QPagedPaintDevice::Margins margins()
    C.C_ZNK17QPagedPaintDevice7marginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "margins", args)
  }

}

// pageSize()
func (this *QPagedPaintDevice) pageSize(args ...interface{}) () {
  // pageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QPagedPaintDevice8pageSizeEv
    // invoke: QPagedPaintDevice::PageSize pageSize()
    C.C_ZNK17QPagedPaintDevice8pageSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPagedPaintDevice", "pageSize", args)
  }

}

// <= body block end

