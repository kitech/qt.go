#include <stdio.h>
#include <stdint.h>
#include <stddef.h>

typedef struct {
    void** vtab;
} vtclass;

typedef void (*func_ptr_t)();
struct method_ptr_t {
    union {
        func_ptr_t function; // Pointer to the function.
        ptrdiff_t offset; // Byte offset into the vtable plus one.
    };
    ptrdiff_t delta; // Byte shift for the "this" pointer.
};
typedef struct {
    func_ptr_t func[1];
} vtable_t;
typedef struct {
    vtable_t * vtable_ptr;
} class_instance_t;
///

void* getmethod(void* cthis, int idx) {
    vtclass* t = (vtclass*)cthis;
    return t->vtab[idx];
}

void* setmethod(void* cthis, int idx, void* newfn) {
    void* oldmth = getmethod(cthis, idx);
    vtclass* t = (vtclass*)cthis;
    t->vtab[idx] = newfn;
    return oldmth;
}


/////// https://github.com/Thordin/vtable-hook/blob/master/vtablehook.h
#include <stdint.h>

#ifdef WIN32
#include <Windows.h>
#elif __linux__
#include <unistd.h>
#include <sys/mman.h>
//int vtablehook_pagesize = sysconf(_SC_PAGE_SIZE);
//int vtablehook_pagemask = ~(vtablehook_pagesize-1);
#endif

int vtablehook_unprotect(void* region) {
#ifdef WIN32
        MEMORY_BASIC_INFORMATION mbi;
        VirtualQuery((LPCVOID)region, &mbi, sizeof(mbi));
        VirtualProtect(mbi.BaseAddress, mbi.RegionSize, PAGE_READWRITE, &mbi.Protect);
        return mbi.Protect;
#elif __linux__
        int vtablehook_pagesize = sysconf(_SC_PAGE_SIZE);
        int vtablehook_pagemask = ~(vtablehook_pagesize-1);

        mprotect((void*) ((intptr_t)region & vtablehook_pagemask), vtablehook_pagesize, PROT_READ|PROT_WRITE|PROT_EXEC);
        return PROT_READ|PROT_EXEC;
#endif
}

void vtablehook_protect(void* region, int protection) {
#ifdef WIN32
        MEMORY_BASIC_INFORMATION mbi;
        VirtualQuery((LPCVOID)region, &mbi, sizeof(mbi));
        VirtualProtect(mbi.BaseAddress, mbi.RegionSize, protection, &mbi.Protect);
#elif __linux__
        int vtablehook_pagesize = sysconf(_SC_PAGE_SIZE);
        int vtablehook_pagemask = ~(vtablehook_pagesize-1);
        mprotect((void*) ((intptr_t)region & vtablehook_pagemask), vtablehook_pagesize, protection);
#endif
}

/*
* instance: pointer to an instance of a class
* hook: function to overwrite with
* offset: 0 = method 1, 1 = method 2 etc...
* return: original function
*/

void* vtablehook_hook(void* instance, void* hook, int offset) {
    intptr_t vtable = *((intptr_t*)instance);
    intptr_t entry = vtable + sizeof(intptr_t) * offset;
    intptr_t original = *((intptr_t*) entry);

    int original_protection = vtablehook_unprotect((void*)entry);
    *((intptr_t*)entry) = (intptr_t)hook;
    vtablehook_protect((void*)entry, original_protection);

    return (void*)original;
}

void myevent(void* e) {
    printf("%p \n", e);
}
void test_vthook(void* app) {
    vtablehook_hook(app, myevent, 200/8);
    // setmethod(app, )
}


