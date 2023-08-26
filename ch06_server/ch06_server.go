package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var contents = []string{
	"これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	"ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

func isGzipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","), "gzip",
	) != -1
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")

		// response := http.Response{
		// 	StatusCode: 200,
		// 	ProtoMajor: 1,
		// 	ProtoMinor: 1,
		// 	Header:     make(http.Header),
		// }
		// if isGzipAcceptable(request) {
		// 	content := "Hello, world (gzipped)\n"
		// 	var buffer bytes.Buffer
		// 	writer := gzip.NewWriter(&buffer)
		// 	io.WriteString(writer, content)
		// 	writer.Close()
		// 	response.Body = io.NopCloser(&buffer)
		// 	response.ContentLength = int64(buffer.Len())
		// 	response.Header.Set("Content-Encoding", "gzip")
		// } else {
		// 	content := "Hello, world\n"
		// 	response.Body = io.NopCloser(strings.NewReader(content))
		// 	response.ContentLength = int64(len(content))
		// }
		// response.Write(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
