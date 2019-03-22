FROM archlinux/base

ADD run-mingw.sh /root/
RUN chmod +x /root/run-mingw.sh

ADD . /root/go/src/github.com/kitech/qt.go/
# ADD CMakeLists.txt /root/
# ADD src /root/src
# ADD tc-mingw.cmake /root/

ENTRYPOINT ["/root/run-mingw.sh"]
