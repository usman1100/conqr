BINARY_NAME=conqr

build:
	go build -o ${BINARY_NAME} -v

clean:
	go clean
	rm -f ${BINARY_NAME}


