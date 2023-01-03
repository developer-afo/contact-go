package helper

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func UuidPointerToString(uid *uuid.UUID) (str string) {
	if uid == nil {
		return ""
	}
	return uid.String()
}

func StringToUuidPointer(str string) (uid *uuid.UUID) {
	if str == "00000000-0000-0000-0000-000000000000" || str == "" {
		return nil
	}
	nuid, _ := uuid.Parse(str)
	return &nuid
}

func GenerateAlphabets(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func GenerateNumeric(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func GenerateAlphaNumeric(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
