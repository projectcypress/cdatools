package exporter

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// TimeToFormat parses time from a seconds since Epoch value, and spits out a string in the supplied format
func timeToFormat(t int64, f string) string {
	parsedTime := time.Unix(t, 0)
	return parsedTime.Format(f)
}

// IdentifierFor MD5s a byte slice, and returns a String
func identifierFor(b []byte) string {
	md := md5.Sum(b)
	return strings.ToUpper(hex.EncodeToString(md[:]))
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
