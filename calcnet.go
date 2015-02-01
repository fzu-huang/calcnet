package calcnet

import (
	"encoding/binary"
	"net"
	"strconv"
	"strings"
)

func CalcNetAddr(ip1, ip2 []byte) string {
	ipint1 := binary.BigEndian.Uint32(ip1[0:])
	ipint2 := binary.BigEndian.Uint32(ip2[0:])

	calc1 := ipint1 ^ ipint2
	count := 0
	for {
		if calc1 == 0 {
			break
		}
		calc1 = calc1 >> 1
		count++
	}
	//intipint1 := int(ipint1)
	ipnetint := (ipint1 >> uint32(count)) << uint32(count)
	ipnet := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(ipnet[0:], ipnetint)
	return net.IPv4(ipnet[0], ipnet[1], ipnet[2], ipnet[3]).String() + "/" + strconv.Itoa(32-count)
}

func AtoIPByte(ip1, ip2 string) ([]byte, []byte) {
	ip1part := strings.SplitN(ip1, ".", 4)
	ip2part := strings.SplitN(ip2, ".", 4)
	ip1byte := make([]byte, 4, 4)
	ip2byte := make([]byte, 4, 4)
	for i := 0; i < 4; i++ {
		tint, _ := strconv.Atoi(ip1part[i])
		ip1byte[i] = uint8(tint)

		tint, _ = strconv.Atoi(ip2part[i])
		ip2byte[i] = uint8(tint)
	}
	return ip1byte, ip2byte
}
