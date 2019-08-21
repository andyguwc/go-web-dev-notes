/* Dials a TCP Network at the host address 

Which returns a TCP connection and uses the TCP connection to issue HTTP GET request to retrieve text and write to a local file


Because the net.Conn type implements the io.Reader and io.Writer, it can be used to
both send data and receive data using streaming IO semantics. In the preceding example,
conn.Write([]byte(httpRequest)) sends the HTTP request to the server.

*/



func main() {
	host, port := "www.gutenberg.org", "80"
	addr := net.JoinHostPort(host, port)
	httpRequest:="GET /cache/epub/16328/pg16328.txt HTTP/1.1\n" +"Host: " + host + "\n\n"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("beowulf.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	io.Copy(file, conn)
	fmt.Println("Text copied to file", file.Name())

}

