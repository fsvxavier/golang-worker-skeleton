package uuid

import (
	"regexp"
	"testing"
)

const format = "^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$"

func TestSetVarient(t *testing.T) {

	ReservedNCS := byte(0x80)
	ReservedRFC4122 := byte(0x40)
	ReservedMicrosoft := byte(0x20)
	ReservedFuture := byte(0x00)

	u := new(UUID)
	u.setVariant(ReservedNCS)
	u.Variant()
	u.setVariant(ReservedRFC4122)
	u.Variant()
	u.setVariant(ReservedMicrosoft)
	u.Variant()
	u.setVariant(ReservedFuture)
	u.Variant()
}

func TestParse(t *testing.T) {
	_, err := Parse([]byte{1, 2, 3, 4, 5})
	if err == nil {
		t.Errorf("Expected error due to invalid UUID sequence")
	}
	base, _ := NewUUIDV4()
	u, err := Parse(base[:])
	if err != nil {
		t.Errorf("Expected to parse UUID sequence without problems")
		return
	}
	if u.String() != base.String() {
		t.Errorf("Expected parsed UUID to be the same as base, %s != %s", u.String(), base.String())
	}
}

func TestParseString(t *testing.T) {
	_, err := ParseHex("foo")
	if err == nil {
		t.Errorf("Expected error due to invalid UUID string")
	}
	base, _ := NewUUIDV4()
	u, err := ParseHex(base.String())
	if err != nil {
		t.Errorf("Expected to parse UUID sequence without problems")
		return
	}
	if u.String() != base.String() {
		t.Errorf("Expected parsed UUID to be the same as base, %s != %s", u.String(), base.String())
	}
}

func TestNewV3(t *testing.T) {
	u, err := NewUUIDV3(NamespaceURL, []byte("golang.org"))
	if err != nil {
		t.Errorf("Expected to generate UUID without problems, error thrown: %v", err.Error())
		return
	}
	if u.Version() != 3 {
		t.Errorf("Expected to generate UUIDv3, given %d", u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected to generate UUIDv3 RFC4122 variant, given %x", u.Variant())
	}
	re := regexp.MustCompile(format)
	if !re.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given %s", u.String())
	}
	u2, _ := NewUUIDV3(NamespaceURL, []byte("golang.org"))
	if u2.String() != u.String() {
		t.Errorf("Expected UUIDs generated of the same namespace and name to be the same")
	}
	u3, _ := NewUUIDV3(NamespaceDNS, []byte("golang.org"))
	if u3.String() == u.String() {
		t.Errorf("Expected UUIDs generated of different namespace and the same name to be different")
	}
	u4, _ := NewUUIDV3(NamespaceURL, []byte("code.google.com"))
	if u4.String() == u.String() {
		t.Errorf("Expected UUIDs generated of the same namespace and different names to be different")
	}
	u5, _ := NewUUIDV3(NamespaceX500, []byte("code.google.com"))
	if u5.String() == u.String() {
		t.Errorf("Expected UUIDs generated of the same namespace and different names to be different")
	}
	u6, _ := NewUUIDV3(nil, []byte(""))
	if u6.String() == u.String() {
		t.Errorf("Expected UUIDs generated of the same namespace and different names to be different")
	}
}

func TestNewV4(t *testing.T) {
	u, err := NewUUIDV4()
	if err != nil {
		t.Errorf("Expected to generate UUID without problems, error thrown: %s", err.Error())
		return
	}
	if u.Version() != 4 {
		t.Errorf("Expected to generate UUIDv4, given %d", u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected to generate UUIDv4 RFC4122 variant, given %x", u.Variant())
	}
	re := regexp.MustCompile(format)
	if !re.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given %s", u.String())
	}
}

func TestNewV5(t *testing.T) {
	u, err := NewUUIDV5(NamespaceURL, []byte("golang.org"))
	if err != nil {
		t.Errorf("Expected to generate UUID without problems, error thrown: %v", err.Error())
		return
	}
	if u.Version() != 5 {
		t.Errorf("Expected to generate UUIDv5, given %d", u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected to generate UUIDv5 RFC4122 variant, given %x", u.Variant())
	}
	re := regexp.MustCompile(format)
	if !re.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given %s", u.String())
	}
	u2, _ := NewUUIDV5(NamespaceURL, []byte("golang.org"))
	if u2.String() != u.String() {
		t.Errorf("Expected UUIDs generated of the same namespace and name to be the same")
	}
	u3, _ := NewUUIDV5(NamespaceDNS, []byte("golang.org"))
	if u3.String() == u.String() {
		t.Errorf("Expected UUIDs generated of different namespace and the same name to be different")
	}
	u4, _ := NewUUIDV5(NamespaceURL, []byte("code.google.com"))
	if u4.String() == u.String() {
		t.Errorf("Expected UUIDs generated of the same namespace and different names to be different")
	}
}

func BenchmarkParseHex(b *testing.B) {
	s := "f3593cff-ee92-40df-4086-87825b523f13"
	for i := 0; i < b.N; i++ {
		_, err := ParseHex(s)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}
