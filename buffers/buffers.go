package buffers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func Run() {

	// we can pass below buffers anywhere io.Reader/io.Writer is required to read/write our given data.
	buff := new(bytes.Buffer)                         // empty buffer - no data given
	buff1 := bytes.NewBuffer([]byte("one two three")) // create rw buffer from given byte
	buff2 := bytes.NewBufferString("testdata")        // from string rw from given string
	readerBuff := strings.NewReader("testdata")       // readonly from given string

	if buff == nil || buff1 == nil || buff2 != nil || readerBuff != nil {
		fmt.Println("buffers nil")
	}

	// Fscan - read data from reader, separate by coma or newline , put values in argument pointers
	var str1, str2, str3 string
	n, err := fmt.Fscan(buff1, &str1, &str2, &str3)
	if err == nil && n > 0 {
		fmt.Printf("Fscan read from buffer: str1:%v, str2:%v, str3:%v", str1, str2, str3)
	} else {
		fmt.Printf("error in Fscan")
	}

	// write data to buffer and print it
	buff.Write([]byte("one two three four yeess"))
	buff.WriteString(" ***using wrtie string")
	fmt.Printf("\nprint data from buffer:%v", buff.String())

	// Fprint - write data to buffer using Fprint, by putting values in argument
	buff = new(bytes.Buffer)
	n, err = fmt.Fprint(buff, str1, str2, str3)
	if err == nil && n > 0 {
		fmt.Printf("\nFprint wrtie to buffer data:%v", buff.String())
	}

	//ReadAll, read from custom buffer
	src := "source data"
	srcBuff := strings.NewReader(src) // create a new reader from source string
	srcRead, err := io.ReadAll(srcBuff)
	fmt.Printf("\nReadAll from custom buffer, data:%v", string(srcRead))

	// compress using gzip buffer
	com_buff := new(bytes.Buffer)                       // create empty rw buffer
	gzip_writer := gzip.NewWriter(com_buff)             // create new gzip write buffer holding our buffer
	n, err = gzip_writer.Write([]byte("compress this")) //  write data to gzip buffer
	if err == nil && n > 0 {
		gzip_writer.Close() // after closing, our buffer have/flush the compressed data
	}
	gzip_reader, err := gzip.NewReader(com_buff) // com_buff contains the compressed data, create a new read buf from our buffer
	if err == nil {
		gzip_reader.Close()
		//io.Copy(os.Stdout, gzip_reader)
		data_orig, _ := io.ReadAll(gzip_reader) // read from gzip_reader to read uncompressed original data
		fmt.Printf("\ngzip uncompress, data:%v", string(data_orig))
	}

}
