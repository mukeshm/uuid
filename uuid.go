package uuid

import (
	"crypto/rand"
	"fmt"
)

type UUID [16]byte

//Generates the UUID v4
func GenerateV4() (UUID, error) {
	var uuid UUID
	_, err := rand.Read(uuid[:])
	if err != nil {
		return uuid, nil
	}
	//13th character is "4" for UUID v4
	uuid[6] = (uuid[6] | 0x40) & 0x4f
	// 17th character should be "8", "9", "a", or "b"
	uuid[8] = (uuid[8] | 0x80) & 0xbf
	return uuid, nil
}

//returns the string representation of UUID
func (u *UUID) String() string {
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid
}
