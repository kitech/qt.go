package qtwidgets

// /usr/include/qt/QtWidgets/qmdisubwindow.h
// #include <qmdisubwindow.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end

//  body block begin
type QMdiSubWindowList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QMdiSubWindowList) Operator_equal0() *QMdiSubWindowList {
	// QMdiSubWindowList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QMdiSubWindowList) Operator_equal1() *QMdiSubWindowList {
	// QMdiSubWindowList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QMdiSubWindowList) Swap0() {
	// QMdiSubWindowList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QMdiSubWindowList) Operator_equal_equal0() bool {
	// QMdiSubWindowList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QMdiSubWindowList) Operator_not_equal0() bool {
	// QMdiSubWindowList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QMdiSubWindowList) Size0() int {
	// QMdiSubWindowList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QMdiSubWindowList) Detach0() {
	// QMdiSubWindowList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QMdiSubWindowList) DetachShared0() {
	// QMdiSubWindowList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QMdiSubWindowList) IsDetached0() bool {
	// QMdiSubWindowList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QMdiSubWindowList) SetSharable0() {
	// QMdiSubWindowList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QMdiSubWindowList) IsSharedWith0() bool {
	// QMdiSubWindowList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QMdiSubWindowList) IsEmpty0() bool {
	// QMdiSubWindowList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QMdiSubWindowList) Clear0() {
	// QMdiSubWindowList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QMdiSubWindowList) At0() *QMdiSubWindow {
	// QMdiSubWindowList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// const T & operator[](int)
func (this *QMdiSubWindowList) Operator_get_index0() *QMdiSubWindow {
	// QMdiSubWindowList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// T & operator[](int)
func (this *QMdiSubWindowList) Operator_get_index1() *QMdiSubWindow {
	// QMdiSubWindowList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// void reserve(int)
func (this *QMdiSubWindowList) Reserve0() {
	// QMdiSubWindowList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QMdiSubWindowList) Append0() {
	// QMdiSubWindowList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QMdiSubWindowList) Append1() {
	// QMdiSubWindowList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QMdiSubWindowList) Prepend0() {
	// QMdiSubWindowList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QMdiSubWindowList) Insert0() {
	// QMdiSubWindowList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QMdiSubWindowList) Replace0() {
	// QMdiSubWindowList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QMdiSubWindowList) RemoveAt0() {
	// QMdiSubWindowList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QMdiSubWindowList) RemoveAll0() int {
	// QMdiSubWindowList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QMdiSubWindowList) RemoveOne0() bool {
	// QMdiSubWindowList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QMdiSubWindowList) TakeAt0() *QMdiSubWindow {
	// QMdiSubWindowList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// T takeFirst()
func (this *QMdiSubWindowList) TakeFirst0() *QMdiSubWindow {
	// QMdiSubWindowList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// T takeLast()
func (this *QMdiSubWindowList) TakeLast0() *QMdiSubWindow {
	// QMdiSubWindowList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// void move(int, int)
func (this *QMdiSubWindowList) Move0() {
	// QMdiSubWindowList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QMdiSubWindowList) Swap1() {
	// QMdiSubWindowList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QMdiSubWindowList) IndexOf0() int {
	// QMdiSubWindowList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QMdiSubWindowList) LastIndexOf0() int {
	// QMdiSubWindowList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QMdiSubWindowList) Contains0() bool {
	// QMdiSubWindowList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QMdiSubWindowList) Count0() int {
	// QMdiSubWindowList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QMdiSubWindowList) Begin0() {
	// QMdiSubWindowList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QMdiSubWindowList) Begin1() {
	// QMdiSubWindowList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QMdiSubWindowList) Cbegin0() {
	// QMdiSubWindowList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QMdiSubWindowList) ConstBegin0() {
	// QMdiSubWindowList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QMdiSubWindowList) End0() {
	// QMdiSubWindowList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QMdiSubWindowList) End1() {
	// QMdiSubWindowList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QMdiSubWindowList) Cend0() {
	// QMdiSubWindowList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QMdiSubWindowList) ConstEnd0() {
	// QMdiSubWindowList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QMdiSubWindowList) Rbegin0() {
	// QMdiSubWindowList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QMdiSubWindowList) Rend0() {
	// QMdiSubWindowList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QMdiSubWindowList) Rbegin1() {
	// QMdiSubWindowList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QMdiSubWindowList) Rend1() {
	// QMdiSubWindowList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QMdiSubWindowList) Crbegin0() {
	// QMdiSubWindowList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QMdiSubWindowList) Crend0() {
	// QMdiSubWindowList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QMdiSubWindowList) Insert1() {
	// QMdiSubWindowList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QMdiSubWindowList) Erase0() {
	// QMdiSubWindowList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QMdiSubWindowList) Erase1() {
	// QMdiSubWindowList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QMdiSubWindowList) Count1() int {
	// QMdiSubWindowList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QMdiSubWindowList) Length0() int {
	// QMdiSubWindowList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QMdiSubWindowList) First0() *QMdiSubWindow {
	// QMdiSubWindowList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// const T & constFirst()
func (this *QMdiSubWindowList) ConstFirst0() *QMdiSubWindow {
	// QMdiSubWindowList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// const T & first()
func (this *QMdiSubWindowList) First1() *QMdiSubWindow {
	// QMdiSubWindowList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// T & last()
func (this *QMdiSubWindowList) Last0() *QMdiSubWindow {
	// QMdiSubWindowList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// const T & last()
func (this *QMdiSubWindowList) Last1() *QMdiSubWindow {
	// QMdiSubWindowList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// const T & constLast()
func (this *QMdiSubWindowList) ConstLast0() *QMdiSubWindow {
	// QMdiSubWindowList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// void removeFirst()
func (this *QMdiSubWindowList) RemoveFirst0() {
	// QMdiSubWindowList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QMdiSubWindowList) RemoveLast0() {
	// QMdiSubWindowList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QMdiSubWindowList) StartsWith0() bool {
	// QMdiSubWindowList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QMdiSubWindowList) EndsWith0() bool {
	// QMdiSubWindowList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QMdiSubWindowList) Mid0() *QMdiSubWindowList {
	// QMdiSubWindowList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QMdiSubWindowList) Value0() *QMdiSubWindow {
	// QMdiSubWindowList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// T value(int, const T &)
func (this *QMdiSubWindowList) Value1() *QMdiSubWindow {
	// QMdiSubWindowList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// void push_back(const T &)
func (this *QMdiSubWindowList) Push_back0() {
	// QMdiSubWindowList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QMdiSubWindowList) Push_front0() {
	// QMdiSubWindowList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QMdiSubWindowList) Front0() *QMdiSubWindow {
	// QMdiSubWindowList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// const T & front()
func (this *QMdiSubWindowList) Front1() *QMdiSubWindow {
	// QMdiSubWindowList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// T & back()
func (this *QMdiSubWindowList) Back0() *QMdiSubWindow {
	// QMdiSubWindowList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// const T & back()
func (this *QMdiSubWindowList) Back1() *QMdiSubWindow {
	// QMdiSubWindowList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMdiSubWindow{}
}

// void pop_front()
func (this *QMdiSubWindowList) Pop_front0() {
	// QMdiSubWindowList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QMdiSubWindowList) Pop_back0() {
	// QMdiSubWindowList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QMdiSubWindowList) Empty0() bool {
	// QMdiSubWindowList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QMdiSubWindowList) Operator_add_equal0() *QMdiSubWindowList {
	// QMdiSubWindowList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QMdiSubWindowList) Operator_add0() *QMdiSubWindowList {
	// QMdiSubWindowList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QMdiSubWindowList) Operator_add_equal1() *QMdiSubWindowList {
	// QMdiSubWindowList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QMdiSubWindowList) Operator_left_shift0() *QMdiSubWindowList {
	// QMdiSubWindowList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QMdiSubWindowList) Operator_left_shift1() *QMdiSubWindowList {
	// QMdiSubWindowList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QMdiSubWindowList) ToVector0() {
	// QMdiSubWindowList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QMdiSubWindowList) ToSet0() {
	// QMdiSubWindowList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QMdiSubWindowList) FromVector0() *QMdiSubWindowList {
	// QMdiSubWindowList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QMdiSubWindowList) FromSet0() *QMdiSubWindowList {
	// QMdiSubWindowList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QMdiSubWindowList) FromStdList0() *QMdiSubWindowList {
	// QMdiSubWindowList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QMdiSubWindowList) ToStdList0() {
	// QMdiSubWindowList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QMdiSubWindowList) Detach_helper_grow0() {
	// QMdiSubWindowList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QMdiSubWindowList) Detach_helper0() {
	// QMdiSubWindowList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QMdiSubWindowList) Detach_helper1() {
	// QMdiSubWindowList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QMdiSubWindowList) Dealloc0() {
	// QMdiSubWindowList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QMdiSubWindowList) Node_construct0() {
	// QMdiSubWindowList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QMdiSubWindowList) Node_destruct0() {
	// QMdiSubWindowList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QMdiSubWindowList) Node_copy0() {
	// QMdiSubWindowList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QMdiSubWindowList) Node_destruct1() {
	// QMdiSubWindowList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QMdiSubWindowList) IsValidIterator0() bool {
	// QMdiSubWindowList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QMdiSubWindowList) Op_eq_impl0() bool {
	// QMdiSubWindowList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QMdiSubWindowList) Op_eq_impl1() bool {
	// QMdiSubWindowList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QMdiSubWindowList) Contains_impl0() bool {
	// QMdiSubWindowList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QMdiSubWindowList) Contains_impl1() bool {
	// QMdiSubWindowList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QMdiSubWindowList) Count_impl0() int {
	// QMdiSubWindowList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QMdiSubWindowList) Count_impl1() int {
	// QMdiSubWindowList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMdiSubWindowList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
