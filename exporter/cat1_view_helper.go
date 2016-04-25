package exporter

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

func escape(i interface{}) string {
	switch str := i.(type) {
	case string:
		return escapeString(str)
	case int64:
		return escapeString(strconv.FormatInt(str, 10))
	case int:
		return escapeString(strconv.Itoa(str))
	}
	return ""
}

func escapeString(s string) string {
	b := new(bytes.Buffer)
	xml.Escape(b, []byte(s))
	return b.String()
}

// TimeToFormat parses time from a seconds since Epoch value, and spits out a string in the supplied format
func timeToFormat(t int64, f string) string {
	parsedTime := time.Unix(t, 0)
	return escapeString(parsedTime.Format(f))
}

// IdentifierFor MD5s a byte slice, and returns a String
func identifierFor(b []byte) string {
	md := md5.Sum(b)
	return escapeString(strings.ToUpper(hex.EncodeToString(md[:])))
}

// IdentifierForInt generates an MD5 representation of a set of int64s
func identifierForInt(objs ...int64) string {
	b := make([]byte, len(objs))
	for _, val := range objs {
		b = append(b, []byte(strconv.FormatInt(val, 10))...)
	}
	return identifierFor(b)
}

// IdentifierForString generates an MD5 representation of a set of string objects
func identifierForString(objs ...string) string {
	b := strings.Join(objs, ",")
	return identifierFor([]byte(b))
}
