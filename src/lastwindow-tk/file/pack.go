package file

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readBytes(f *os.File, byteLen int) []byte {
	b1 := make([]byte, byteLen)
	_, err := f.Read(b1)
	check(err)
	return b1
}

func readLong(f *os.File) uint32 {
	return binary.BigEndian.Uint32(readBytes(f, 4))
}

func verifyPack(f *os.File) {
	invalid := "Invalid byte pattern in packfile"
	if readLong(f) != 0 {
		panic(invalid)
	}
	if !(readLong(f) > 0) {
		panic(invalid)
	}
	if !(readLong(f) > 0) {
		panic(invalid)
	}
}

type extractFile struct {
	size int
	name string
}

func readByte(f *os.File) uint64 {
	dat, err := binary.ReadUvarint(bytes.NewBuffer(readBytes(f, 1)))
	check(err)
	return dat
}

func readHeader(f *os.File) (int, []extractFile) {
	f.Seek(4, 0) // 0 resets the offset
	fileCount := readLong(f)
	startAddr := readLong(f) + 8
	readLong(f)
	var files []extractFile
	for i := 1; i <= int(fileCount); i++ {
		nameLen := readByte(f)
		nameBytes := readBytes(f, int(nameLen))
		fname := bytes.NewBuffer(nameBytes).String()
		zsize := readLong(f)
		files = append(files, extractFile{
			size: int(zsize),
			name: string(fname),
		})
	}
	return int(startAddr), files
}

func extractFiles(f *os.File, start int, files []extractFile, outputDir string) {
	f.Seek(int64(start), 0)
	for _, toExtract := range files {
		fileData := bytes.NewBuffer(readBytes(f, toExtract.size))
		if filepath.Ext(toExtract.name) != ".bra" { // .bra files are uncompressed for some reason
			reader, err := zlib.NewReader(fileData)
			check(err)
			io.Copy(fileData, reader)
			reader.Close()
		}
		exportPath := filepath.Join(outputDir, toExtract.name)
		err := os.WriteFile(exportPath, fileData.Bytes(), os.ModePerm)
		check(err)
		fmt.Printf("Extracted %s\n", exportPath)
	}
}

func UnpackPackfile(fpath string, outputDir string) {
	f, err := os.Open(fpath)
	check(err)
	verifyPack(f)
	start, files := readHeader(f)
	extractFiles(f, start, files, outputDir)
	defer f.Close()
}
