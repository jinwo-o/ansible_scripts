# Use golang:1.13 when becomes available
FROM golang:1.12

WORKDIR /go/src/github.com/Bio-core/jtree
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Import gouuid executable to generate uuid
RUN go get github.com/nu7hatch/gouuid

RUN make build

EXPOSE 8000

# May have to add entrypoint
# 1 NOPE
# CMD /go/src/github.com/Bio-core/jtree/bin/jtree
# 2 NOPE
# CMD ./bin/jtree
# 3 NOPE
# RUN ./bin/jtree
# 4 NOPE
# RUN /go/src/github.com/Bio-core/jtree/bin/jtree
# 5 NOPE
# RUN ./go/src/github.com/Bio-core/jtree/bin/jtree
# 6 NOPE
# ENTRYPOINT [ "/go/src/github.com/Bio-core/jtree/bin/jtree" ]
# 7 NOPE
ENTRYPOINT [ "./bin/jtree" ]
# 8 NOPE
# ENTRYPOINT [ "/bin/jtree" ]