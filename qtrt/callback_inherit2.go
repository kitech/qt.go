package qtrt

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

*/
import "C"
import (
	"log"
	"reflect"

	qt "github.com/kitech/qt.go/qtqt"
)

// inherit using virtual method hook

func SetAllInheritCallback2(cbobj CObjectITF, name string, f interface{}) {
	if cbobj.GetCthis() != nil {
	} else { // take as global function call, just no this field
		if debugDynSlot {
			log.Println("obj is nil, try as global event", name)
		}
	}

	clsname := getClassNameByObject(cbobj)
	log.Println(clsname, reflect.TypeOf(cbobj).String(), cbobj.GetCthis())

	// check hookable
	vtidx := virtab.getVTMIndex(clsname, name)
	if vtidx == -1 {
		log.Println("Cannot inherit for not found", clsname, name)
		return
	}

	if signal := qt.LendSignal(cbobj.GetCthis(), name); signal != nil {
		log.Println("already exists:", name, reflect.TypeOf(cbobj))
		qt.ConnectSignal(cbobj.GetCthis(), name, func(args ...uint64) {
			signal.(func(...uint64))(args...)
			callbackInheritInvokeGo(f, args...)
		})
	} else {
		qt.ConnectSignal(cbobj.GetCthis(), name, f)
	}

	// check hook
	virtab.hookit(clsname, name, cbobj.GetCthis(), f)
}

// TODO unhookit
func UnsetAllInheritCallback2(cbobj CObjectITF, name string) {
	if signal := qt.LendSignal(cbobj.GetCthis(), name); signal != nil {
		qt.DisconnectSignal(cbobj.GetCthis(), name)
	} else {
		log.Println("not exists:", cbobj, name)
	}
}

// TODO unhookit
func UnsetAllInheritCallbackAll2(cbobj CObjectITF) {
	qt.DisconnectObject(cbobj.GetCthis())
}

// called by cppvm_ffi_closure_callback
func callbackHookInherits2(vmclos *CppvmClosure, argvals *Voidptr) {
	log.Println("here", vmclos, argvals)
	argv := *(*[1 << 8]Voidptr)(Voidptr(argvals)) // 128 enough
	cbobj := *(*Voidptr)(argv[0])
	argc := vmclos.argc - 1
	pargs := make([]uint64, argc)
	for i := 0; i < argc; i++ {
		t := *(*Voidptr)(argv[i+1]) // dereference
		pargs[i] = uint64(uintptr(t))
	}
	log.Println(cbobj, vmclos.mthname, argc, len(pargs))
	var handled int
	callbackAllInheritsGo(cbobj, vmclos.mthname, &handled, argc, pargs...)
	log.Println(cbobj, vmclos.mthname, handled, argc)
}
