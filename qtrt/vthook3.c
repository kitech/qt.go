// https://github.com/lvc/vtable-dumper/blob/master/dump-vtable.h
// dlsym(handle, "_ZTV6QLabel")

#define _GNU_SOURCE

#include <stdlib.h>
#include <stdio.h>
#include <dlfcn.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
// #include <libelf.h>
// #include <gelf.h>
#include <string.h>
#include <ctype.h>
#include <stddef.h>
#include <limits.h>

#define ERR -1

#if defined(__ia64__)
typedef struct {
    void *func;
    void *gpoffset;
} fptr;
#else
typedef unsigned long fptr;
#endif

struct cat1vtable_mem {
    unsigned long   baseoffset;
    const char  *typeinfo;
    fptr        virtfuncs[0];
};

struct cat2vtable_mem {
    unsigned long   vcalloffset;
    unsigned long   baseoffset;
    const char  *typeinfo;
    fptr        virtfuncs[0];
};

union classvtable_mem {
    struct cat1vtable_mem cat1;
    struct cat2vtable_mem cat2;
};

typedef struct
{
    char* name;
    unsigned long size;
} vtable_info;

/////////////
void *
fptr2ptrp(fptr * fptr)
{
#if defined(__ia64__)
    return fptr->func;
#else
    return (void *) *fptr;
#endif
}

// elfvtptr = dlsym(handle, "_ZTV6QLabel")
void dump_vtable_entry(void* elfvtptr) {
    union classvtable_mem* vtablep = elfvtptr;
    fptr *vtvirtfuncs = vtablep->cat1.virtfuncs;
    const char* vttypeinfo = vtablep->cat1.typeinfo;
    printf("baseoffset %d, typeinfo %p, virtfuncs %p\n",
           vtablep->cat1.baseoffset, vttypeinfo, vtvirtfuncs);

    int step = sizeof(ptrdiff_t);
    int offset = 0;
    void* vfuncp;
    Dl_info dlainfo;

    dladdr(vttypeinfo, &dlainfo);
    printf("typeinfo %s\n", dlainfo.dli_sname);
    offset += step;

    int haszti = 0;
    for (int i = 0; i < 128 ; i++) {
        offset += step;
        vfuncp = fptr2ptrp(&vtvirtfuncs[i]);

        memset(&dlainfo, 0, sizeof(Dl_info));
        int rv = dladdr(vfuncp, &dlainfo);
        const char* sname = dlainfo.dli_sname;
        if (rv == 0) {
            if (haszti) {
                break; // error break
            }
        }
        if (sname != NULL && strncmp(sname, "_ZTI", 4) == 0) {
            haszti = 1;
            //break;
        }
        printf("%d r%d %d %s\n", i, rv, offset, dlainfo.dli_sname);
    }
}

// arr [128]voidptr
int fillin_vtable_entry(void* elfvtptr, void** arr) {
    union classvtable_mem* vtablep = elfvtptr;
    fptr *vtvirtfuncs = vtablep->cat1.virtfuncs;
    const char* vttypeinfo = vtablep->cat1.typeinfo;
    printf("baseoffset %d, typeinfo %p, virtfuncs %p\n",
           vtablep->cat1.baseoffset, vttypeinfo, vtvirtfuncs);

    int step = sizeof(ptrdiff_t);
    int offset = 0;
    void* vfuncp;
    Dl_info dlainfo;

    dladdr(vttypeinfo, &dlainfo);
    printf("typeinfo %s\n", dlainfo.dli_sname);
    offset += step;

    int haszti = 0;
    for (int i = 0; i < 128 ; i++) {
        offset += step;
        vfuncp = fptr2ptrp(&vtvirtfuncs[i]);

        memset(&dlainfo, 0, sizeof(Dl_info));
        int rv = dladdr(vfuncp, &dlainfo);
        const char* sname = dlainfo.dli_sname;
        if (rv == 0) {
            if (haszti) {
                break; // error break
            }
        }
        if (sname != NULL && strncmp(sname, "_ZTI", 4) == 0) {
            haszti = 1;
            //break;
        }
        printf("%d r%d %d %s\n", i, rv, offset, dlainfo.dli_sname);
        arr[offset/8] = (char*)sname;
    }
    return offset/8;
}

typedef struct {
    int step;
    int offset;
} VTableIter;
