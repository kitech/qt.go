// https://blog.mozilla.org/nfroyd/2014/02/20/finding-addresses-of-virtual-functions/


// include
// #include
#include <stdlib.h>
#include <stdint.h>
#include <sys/types.h>
#include <stddef.h>
#include <stdio.h>
#include <string.h>

/* AbstractClass may have several different concrete implementations.  */
/*
class AbstractClass {
public:
  virtual int f() = 0;
  virtual int g() = 0;
};
*/

typedef  struct pointerToMember
  {
    /* This field has separate representations for non-virtual and virtual
       functions.  For non-virtual functions, this field is simply the
       address of the function.  For our case, virtual functions, this
       field is 1 plus the virtual table offset (in bytes) of the function
       in question.  The least-significant bit therefore discriminates
       between virtual and non-virtual functions.

       "Ah," you say, "what about architectures where function pointers do
       not necessarily have even addresses?"  (ARM, MIPS, and AArch64 are
       the major ones.)  Excellent point.  Please see below.  */
    size_t pointerOrOffset;

    /* This field is only interesting for calling the function; it
       describes the amount that the `this' pointer must be adjusted
       prior to the call.  However, on architectures where function
       pointers do not necessarily have even addresses, this field has the
       representation:

       2 * adjustment + (virtual_function_p ? 1 : 0)  */
    ptrdiff_t thisAdjustment;
  } pointerToMember;

/* Return the address of the `f' function of `aClass' that would be
   called for the expression:

   aClass->f();

   regardless of the concrete type of `aClass'.

   It is left as an exercise for the reader to templatize this function for
   arbitrary `f'.  */
void*
find_f_address(/*AbstractClass*/ void* aClass, void* memfnptr)
{
  /* The virtual function table is stored at the beginning of the object.  */
  void** vtable = *(void***)aClass;

  /* This structure is described in the cross-platform "Itanium" C++ ABI:

     http://mentorembedded.github.io/cxx-abi/abi.html

     The particular layout replicated here is described in:

     http://mentorembedded.github.io/cxx-abi/abi.html#member-pointers  */

  /* Translate from the opaque pointer-to-member type representation to
     the representation given above.  */
  pointerToMember p;
  // int ((AbstractClass::*m)()) = &AbstractClass::f;
  void* m = memfnptr; // 似乎和上一行的结果不一样吧？
  memcpy(&p, &m, sizeof(p));

  /* Compute the actual offset into the vtable.  Given the differing meaing
     of the fields between architectures, as described above, and that
     there's no convenient preprocessor macro, we have to do this
     ourselves.  */
#if defined(__arm__) || defined(__mips__) || defined(__aarch64__)
  /* No adjustment required to `pointerOrOffset'.  */
  static const size_t pfnAdjustment = 0;
#else
  /* Strip off the lowest bit of `pointerOrOffset'.  */
  static const size_t pfnAdjustment = 1;
#endif

  printf("ptroroff %d this?%d\n", p.pointerOrOffset, p.thisAdjustment);
  size_t offset = (p.pointerOrOffset - pfnAdjustment) / sizeof(void*);
  printf("vtable %p, offset=%d ptmsz %d\n", vtable, offset, sizeof(pointerToMember));
  /* Now grab the address out of the vtable and return it.  */
  return vtable[offset];
}


