package ipv4

import (
	"fmt"
	"strconv"
	"strings"
)

func Dec2ip(dec uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", dec>>24, (dec>>16)&0xff, (dec>>8)&0xff, dec&0xff)
}

func Ip2dec(ip string) (uint32, error) {
	oct := strings.Split(ip, ".")
	if len(oct) != 4 {
		return 0, fmt.Errorf("invalid ip address: %s", ip)
	}

	res := uint32(0)
	for i := 0; i < 4; i++ {
		v, err := strconv.ParseUint(oct[i], 10, 8)
		if err != nil {
			return 0, fmt.Errorf("invalid ip address: %s", ip)
		}
		res |= uint32(v) << (24 - 8*i)
	}
	return res, nil
}

func Mask2dec(mask uint32) uint32 {
	res := uint32(0)
	for i := 0; i < int(mask); i++ {
		res |= 1 << (31 - i)
	}
	return res
}
