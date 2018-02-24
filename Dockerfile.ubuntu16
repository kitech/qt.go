FROM ubuntu:16.04

ADD run-ubuntu16.sh /root/
RUN chmod +x /root/run-ubuntu16.sh

ADD . /root/go/src/github.com/kitech/qt.go/
# ADD CMakeLists.txt /root/
# ADD src /root/src
# ADD tc-mingw.cmake /root/

ENTRYPOINT ["/root/run-ubuntu16.sh"]
