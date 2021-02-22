package util

import (
	"bytes"

	crand "crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

const (
	// MaxUint64 is the max of uint64, used for check overflow
	MaxUint64 = 1<<64 - 1
)

// MakeFileAbs makes 'file' absolute relative to 'dir' if not already absolute
func MakeFileAbs(file, dir string) (string, error) {
	if file == "" {
		return "", nil
	}
	if filepath.IsAbs(file) {
		return file, nil
	}
	path, err := filepath.Abs(filepath.Join(dir, file))
	if err != nil {
		return "", fmt.Errorf("Failed making '%s' absolute based on '%s'", file, dir)
	}
	return path, nil
}

// FileExists checks to see if a file exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsEmpty check dir empty
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

// IsDir check if path is a dir, not a file
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// MakeDir create dir
func MakeDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	return err
}

// RmDir remove dir
func RmDir(path string) error {
	err := os.RemoveAll(path)
	return err
}

// RmFile remove file
func RmFile(path string) error {
	err := os.Remove(path)
	return err
}

// SafeAdd returns the result and whether overflow occurred,
// same with math.SafeAdd
func SafeAdd(x, y uint64) (v uint64, overflow bool) {
	return x + y, y > MaxUint64-x
}

// RandNum return int in [0, num)
func RandNum(num int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(num)
	return randNum
}

// RandomString return random string includes upper and lower chars
func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return randomString(length, letters)
}

// RandomHex return random string only contain 0-9abcde
func RandomHex(length int) string {
	var letters = []rune("0123456789abcde")
	return randomString(length, letters)
}

func randomString(length int, letters []rune) string {
	var letterLength = int64(len(letters))
	b := make([]rune, length)
	for i := range b {
		r, _ := crand.Int(crand.Reader, big.NewInt(letterLength))
		b[i] = letters[int(r.Int64())]
	}
	return string(b)
}

// RandUint64 return uint64
func RandUint64() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}

// RandUint32 return uint32
func RandUint32() uint32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint32()
}

// CopyBytes copy bytes
func CopyBytes(origin []byte) []byte {
	if origin == nil {
		return nil
	}
	var res = make([]byte, len(origin))
	for i := range origin {
		res[i] = origin[i]
	}
	return res
}

// Marshal json.Marshal with no escape html <>
func Marshal(obj interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(obj)
	if err != nil {
		return nil, err
	}
	data := buffer.Bytes()
	if l := len(data); l > 0 {
		data = data[:l-1]
	}
	return data, nil
}

// UnsafeCopyStruct copies *struct src to dst in a unsafe way(more check needed to avoid panic)
// it will not copy field in ignoreFields
// it will not copy non-exported fields
func UnsafeCopyStruct(src, dst interface{}, ignoreFields ...string) error {
	// check src and dst is the same struct
	if src == nil || dst == nil {
		return errors.New("src/dst should not be nil")
	}

	sType := reflect.TypeOf(src)
	dType := reflect.TypeOf(dst)

	if sType != dType {
		return fmt.Errorf("src type(%s) and dst type(%s) should be the same", sType, dType)
	}
	if sType.Kind() != reflect.Ptr {
		return fmt.Errorf("not ptr type, get: %s", reflect.ValueOf(src).Kind())
	}

	sVal := reflect.ValueOf(src).Elem()
	dVal := reflect.ValueOf(dst).Elem()

	if sVal.Kind() != reflect.Struct {
		return fmt.Errorf("not struct type, get: %s", sVal.Kind())
	}

	for i := 0; i < sVal.NumField(); i++ {
		value := sVal.Field(i)
		name := sVal.Type().Field(i).Name

		ignore := false
		for _, ignoreField := range ignoreFields {
			if name == ignoreField {
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}

		dValue := dVal.FieldByName(name)
		if !dValue.IsValid() {
			continue
		}
		if !dValue.CanSet() {
			continue
		}
		// todo: more check, to ensure dValue can be set
		dValue.Set(value)
	}
	return nil
}

//DelUint64Slice an element of the string slice
func DelUint64Slice(inSlice []uint64, del uint64) ([]uint64, error) {
	for i, v := range inSlice {
		if v == del {
			return append(inSlice[:i], inSlice[i+1:]...), nil
		}
	}
	return inSlice, fmt.Errorf("%d element not existence", del)
}

//DelStringSlice an element of the string slice
func DelStringSlice(inSlice []string, del string) ([]string, error) {
	for i, v := range inSlice {
		if v == del {
			return append(inSlice[:i], inSlice[i+1:]...), nil
		}
	}
	return inSlice, fmt.Errorf("%s element not existence", del)
}
