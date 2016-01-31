package qt5

// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qaccessibleobject.h
// dst-file: /src/gui/qaccessibleobject.go
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
  // proto:  void QAccessibleApplication::QAccessibleApplication();
extern void C_ZN22QAccessibleApplicationC2Ev(void* qthis); // 3
  // proto:  QAccessibleInterface * QAccessibleApplication::parent();
extern void C_ZNK22QAccessibleApplication6parentEv(void* qthis); // 4
  // proto:  int QAccessibleApplication::indexOfChild(const QAccessibleInterface * );
extern void C_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface(void* qthis, void* arg0); // 4
  // proto:  QAccessible::State QAccessibleApplication::state();
extern void C_ZNK22QAccessibleApplication5stateEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleApplication::focusChild();
extern void C_ZNK22QAccessibleApplication10focusChildEv(void* qthis); // 4
  // proto:  QWindow * QAccessibleApplication::window();
extern void C_ZNK22QAccessibleApplication6windowEv(void* qthis); // 4
  // proto:  QAccessible::Role QAccessibleApplication::role();
extern void C_ZNK22QAccessibleApplication4roleEv(void* qthis); // 4
  // proto:  QAccessibleInterface * QAccessibleApplication::child(int index);
extern void C_ZNK22QAccessibleApplication5childEi(void* qthis, int32_t arg0); // 4
  // proto:  int QAccessibleApplication::childCount();
extern void C_ZNK22QAccessibleApplication10childCountEv(void* qthis); // 4
  // proto:  QObject * QAccessibleObject::object();
extern void C_ZNK17QAccessibleObject6objectEv(void* qthis); // 4
  // proto:  bool QAccessibleObject::isValid();
extern void C_ZNK17QAccessibleObject7isValidEv(void* qthis); // 4
  // proto:  void QAccessibleObject::QAccessibleObject(QObject * object);
extern void C_ZN17QAccessibleObjectC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  QAccessibleInterface * QAccessibleObject::childAt(int x, int y);
extern void C_ZNK17QAccessibleObject7childAtEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QRect QAccessibleObject::rect();
extern void C_ZNK17QAccessibleObject4rectEv(void* qthis); // 4
*/
import "C"

// } // <= ext block end

// body block begin =>
func init() {
	if false {
		qtrt.KeepMe()
	}
	if false {
		fmt.Println(123)
	}
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
}

// class sizeof(QAccessibleApplication)=16
type QAccessibleApplication struct {
	/*qbase*/ QAccessibleObject
	qclsinst unsafe.Pointer /* *C.void */
}

// class sizeof(QAccessibleObject)=16
type QAccessibleObject struct {
	/*qbase*/ QAccessibleInterface
	qclsinst unsafe.Pointer /* *C.void */
}

// QAccessibleApplication()
func NewQAccessibleApplication(args ...interface{}) QAccessibleApplication {
	// QAccessibleApplication()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZN22QAccessibleApplicationC1Ev
		// invoke: void QAccessibleApplication()
		var qthis = unsafe.Pointer(C.malloc(5))
		if false {
			reflect.TypeOf(qthis)
		}
		C.C_ZN22QAccessibleApplicationC2Ev(qthis)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "QAccessibleApplication", args)
	}

	return QAccessibleApplication{}
}

// parent()
func (this *QAccessibleApplication) parent(args ...interface{}) {
	// parent()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication6parentEv
		// invoke: QAccessibleInterface * parent()
		C.C_ZNK22QAccessibleApplication6parentEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "parent", args)
	}

}

// indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleApplication) indexOfChild(args ...interface{}) {
	// indexOfChild(const class QAccessibleInterface *)
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)
	vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "const QAccessibleInterface *"

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface
		// invoke: int indexOfChild(const class QAccessibleInterface *)
		var arg0 = args[0].(QAccessibleInterface).qclsinst
		if false {
			fmt.Println(arg0)
		}
		C.C_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface(this.qclsinst, arg0)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "indexOfChild", args)
	}

}

// state()
func (this *QAccessibleApplication) state(args ...interface{}) {
	// state()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication5stateEv
		// invoke: QAccessible::State state()
		C.C_ZNK22QAccessibleApplication5stateEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "state", args)
	}

}

// focusChild()
func (this *QAccessibleApplication) focusChild(args ...interface{}) {
	// focusChild()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication10focusChildEv
		// invoke: QAccessibleInterface * focusChild()
		C.C_ZNK22QAccessibleApplication10focusChildEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "focusChild", args)
	}

}

// window()
func (this *QAccessibleApplication) window(args ...interface{}) {
	// window()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication6windowEv
		// invoke: QWindow * window()
		C.C_ZNK22QAccessibleApplication6windowEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "window", args)
	}

}

// role()
func (this *QAccessibleApplication) role(args ...interface{}) {
	// role()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication4roleEv
		// invoke: QAccessible::Role role()
		C.C_ZNK22QAccessibleApplication4roleEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "role", args)
	}

}

// child(int)
func (this *QAccessibleApplication) child(args ...interface{}) {
	// child(int)
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)
	vtys[0][0] = qtrt.Int32Ty(false) // "int"

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication5childEi
		// invoke: QAccessibleInterface * child(int)
		var arg0 = C.int32_t(args[0].(int32))
		if false {
			fmt.Println(arg0)
		}
		C.C_ZNK22QAccessibleApplication5childEi(this.qclsinst, arg0)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "child", args)
	}

}

// childCount()
func (this *QAccessibleApplication) childCount(args ...interface{}) {
	// childCount()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK22QAccessibleApplication10childCountEv
		// invoke: int childCount()
		C.C_ZNK22QAccessibleApplication10childCountEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleApplication", "childCount", args)
	}

}

// object()
func (this *QAccessibleObject) object(args ...interface{}) {
	// object()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK17QAccessibleObject6objectEv
		// invoke: QObject * object()
		C.C_ZNK17QAccessibleObject6objectEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleObject", "object", args)
	}

}

// isValid()
func (this *QAccessibleObject) isValid(args ...interface{}) {
	// isValid()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK17QAccessibleObject7isValidEv
		// invoke: bool isValid()
		C.C_ZNK17QAccessibleObject7isValidEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleObject", "isValid", args)
	}

}

// QAccessibleObject(class QObject *)
func NewQAccessibleObject(args ...interface{}) QAccessibleObject {
	// QAccessibleObject(class QObject *)
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)
	vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZN17QAccessibleObjectC1EP7QObject
		// invoke: void QAccessibleObject(class QObject *)
		var arg0 = args[0].(QObject).qclsinst
		if false {
			fmt.Println(arg0)
		}
		var qthis = unsafe.Pointer(C.malloc(5))
		if false {
			reflect.TypeOf(qthis)
		}
		// C.C_ZN17QAccessibleObjectC2EP7QObject(qthis, arg0)
	default:
		qtrt.ErrorResolve("QAccessibleObject", "QAccessibleObject", args)
	}

	return QAccessibleObject{}
}

// childAt(int, int)
func (this *QAccessibleObject) childAt(args ...interface{}) {
	// childAt(int, int)
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)
	vtys[0][0] = qtrt.Int32Ty(false) // "int"
	vtys[0][1] = qtrt.Int32Ty(false) // "int"

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK17QAccessibleObject7childAtEii
		// invoke: QAccessibleInterface * childAt(int, int)
		var arg0 = C.int32_t(args[0].(int32))
		if false {
			fmt.Println(arg0)
		}
		var arg1 = C.int32_t(args[1].(int32))
		if false {
			fmt.Println(arg1)
		}
		C.C_ZNK17QAccessibleObject7childAtEii(this.qclsinst, arg0, arg1)
	default:
		qtrt.ErrorResolve("QAccessibleObject", "childAt", args)
	}

}

// rect()
func (this *QAccessibleObject) rect(args ...interface{}) {
	// rect()
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if false {
		fmt.Println(matched_index)
	}
	switch matched_index {
	case 0:
		// invoke: _ZNK17QAccessibleObject4rectEv
		// invoke: QRect rect()
		C.C_ZNK17QAccessibleObject4rectEv(this.qclsinst)
	default:
		qtrt.ErrorResolve("QAccessibleObject", "rect", args)
	}

}

// <= body block end
