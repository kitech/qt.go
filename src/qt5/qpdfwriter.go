package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtGui/qpdfwriter.h
// dst-file: /src/gui/qpdfwriter.go
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
  // proto:  bool QPdfWriter::newPage();
extern bool C_ZN10QPdfWriter7newPageEv(void* qthis); // 4
  // proto:  QString QPdfWriter::creator();
extern void* C_ZNK10QPdfWriter7creatorEv(void* qthis); // 4
  // proto:  void QPdfWriter::~QPdfWriter();
extern void C_ZN10QPdfWriterD2Ev(void* qthis); // 4
  // proto:  void QPdfWriter::setTitle(const QString & title);
extern void C_ZN10QPdfWriter8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QPdfWriter::title();
extern void* C_ZNK10QPdfWriter5titleEv(void* qthis); // 4
  // proto:  void QPdfWriter::QPdfWriter(const QString & filename);
extern void* C_ZN10QPdfWriterC2ERK7QString(void* arg0); // 3
  // proto:  void QPdfWriter::QPdfWriter(QIODevice * device);
extern void* C_ZN10QPdfWriterC2EP9QIODevice(void* arg0); // 3
  // proto:  void QPdfWriter::setPageSizeMM(const QSizeF & size);
extern void C_ZN10QPdfWriter13setPageSizeMMERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QPdfWriter::setResolution(int resolution);
extern void C_ZN10QPdfWriter13setResolutionEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QPdfWriter::metaObject();
extern void C_ZNK10QPdfWriter10metaObjectEv(void* qthis); // 4
  // proto:  void QPdfWriter::setCreator(const QString & creator);
extern void C_ZN10QPdfWriter10setCreatorERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QPdfWriter::resolution();
extern int32_t C_ZNK10QPdfWriter10resolutionEv(void* qthis); // 4
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

// class sizeof(QPdfWriter)=1
type QPdfWriter struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// newPage()
func (this *QPdfWriter) Newpage(args ...interface{}) (ret interface{}) {
  // newPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter7newPageEv
    // invoke: bool newPage()
    var ret0 = C.C_ZN10QPdfWriter7newPageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPdfWriter", "newPage", args)
  }

  return
}

// creator()
func (this *QPdfWriter) Creator(args ...interface{}) (ret interface{}) {
  // creator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter7creatorEv
    // invoke: QString creator()
    var ret0 = C.C_ZNK10QPdfWriter7creatorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPdfWriter", "creator", args)
  }

  return
}

// ~QPdfWriter()
func (this *QPdfWriter) Freeqpdfwriter(args ...interface{}) () {
  // ~QPdfWriter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriterD0Ev
    // invoke: void ~QPdfWriter()
    C.C_ZN10QPdfWriterD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "~QPdfWriter", args)
  }

  return
}

// setTitle(const class QString &)
func (this *QPdfWriter) Settitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter8setTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setTitle", args)
  }

  return
}

// title()
func (this *QPdfWriter) Title(args ...interface{}) (ret interface{}) {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter5titleEv
    // invoke: QString title()
    var ret0 = C.C_ZNK10QPdfWriter5titleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPdfWriter", "title", args)
  }

  return
}

// QPdfWriter(const class QString &)
func NewQPdfWriter(args ...interface{}) *QPdfWriter {
  // QPdfWriter(const class QString &)
  // QPdfWriter(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriterC1ERK7QString
    // invoke: void QPdfWriter(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QPdfWriterC2ERK7QString(arg0)
    return &QPdfWriter{Qclsinst:qthis}
  case 1:
    // invoke: _ZN10QPdfWriterC1EP9QIODevice
    // invoke: void QPdfWriter(class QIODevice *)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QPdfWriterC2EP9QIODevice(arg0)
    return &QPdfWriter{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPdfWriter", "QPdfWriter", args)
  }

  return nil // QPdfWriter{Qclsinst:qthis}
}

// setPageSizeMM(const class QSizeF &)
func (this *QPdfWriter) Setpagesizemm(args ...interface{}) () {
  // setPageSizeMM(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter13setPageSizeMMERK6QSizeF
    // invoke: void setPageSizeMM(const class QSizeF &)
    var arg0 = args[0].(*QSizeF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter13setPageSizeMMERK6QSizeF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setPageSizeMM", args)
  }

  return
}

// setResolution(int)
func (this *QPdfWriter) Setresolution(args ...interface{}) () {
  // setResolution(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter13setResolutionEi
    // invoke: void setResolution(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter13setResolutionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setResolution", args)
  }

  return
}

// metaObject()
func (this *QPdfWriter) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QPdfWriter10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPdfWriter", "metaObject", args)
  }

  return
}

// setCreator(const class QString &)
func (this *QPdfWriter) Setcreator(args ...interface{}) () {
  // setCreator(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPdfWriter10setCreatorERK7QString
    // invoke: void setCreator(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPdfWriter10setCreatorERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPdfWriter", "setCreator", args)
  }

  return
}

// resolution()
func (this *QPdfWriter) Resolution(args ...interface{}) (ret interface{}) {
  // resolution()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPdfWriter10resolutionEv
    // invoke: int resolution()
    var ret0 = C.C_ZNK10QPdfWriter10resolutionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPdfWriter", "resolution", args)
  }

  return
}

// <= body block end

