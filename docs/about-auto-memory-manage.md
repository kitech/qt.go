the qt.go's way:

1. for QObject and subclass, the memory is control by Qt's life time

    when Qt emit destroyed() signal, qt.go will release too.
    
    that's say programmer need release top level QObject instance. 

2. for non QObject, the memory is controlled by qt.go, it's go's GC. 

    that's no need do anything.

3. for non QObject class, but instance's owner maybe move to Qt, like QTableItem.

    this is the problem.
    
    programmer need keep this instance's reference in a container, and 
    
    should carefully remove it's owner by SetCthis to nil when the owner move to Qt.
    
    suck as after QTableWidget::setItem().

4. for functions that return class record instance, qt.go will make a copy of that by new it.

    luckly, these are almost non QObject, so qt.go can automaticlly release it memory.
    
    if you want keep the instance alive, put into some container for own it's reference.

5. there is a possible case, more than one go instance own one qt instance's pointer.

    luckly, these are almost QObject, so there all almost work's fine.
    
    if there really happen, it's need special control.


