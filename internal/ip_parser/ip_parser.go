package ipparser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_MASK    = 32
	SIZE_IP     = 4
	MAX_SIZE    = 0xFFFFFFFF
	QUARTET     = 8
	MAX_QUARTET = 0xFF
)

type IpParser interface {
	GetParentIp(in string) (string, error)
}

type ipParser struct {
	mask []uint64
}

func New(mask int8) *ipParser {
	return &ipParser{
		mask: getMask(mask),
	}
}

// функция для перевода хоста к маски подсети в 10 виде
func getMask(in int8) []uint64 {
	mask := (MAX_SIZE << (MAX_MASK - in)) & MAX_SIZE
	tempMask := MAX_MASK
	localmask := make([]uint64, 0, SIZE_IP)
	for i := 0; i < SIZE_IP; i++ {
		tmp := mask >> (tempMask - QUARTET) & MAX_QUARTET
		localmask = append(localmask, uint64(tmp))
		tempMask -= QUARTET
	}
	return localmask
}

// функция для парсинга ip адресса в слайс из 4 октетов
func parseIP(in string) ([]uint64, error) {
	arr := strings.Split(in, ".")
	ipArr := make([]uint64, SIZE_IP)
	for i, v := range arr {
		val, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("[ParseIP]:%v", err)
		}
		ipArr[i] = val
	}
	return ipArr, nil
}

// метод для определения подсети для ip адресса
func (ip *ipParser) GetParentIp(in string) (string, error) {
	ipArr, err := parseIP(in)

	if err != nil {
		return "", fmt.Errorf("[ParseIP]:%v", err)
	}

	if len(ip.mask) != SIZE_IP || len(ipArr) != SIZE_IP {
		return "", fmt.Errorf("[GetParentIpStr]:incorrect len slice ip")
	}

	parent := make([]uint64, SIZE_IP)

	for i := 0; i < SIZE_IP; i++ {
		parent[i] = ipArr[i] & ip.mask[i]
	}
	return fmt.Sprintf("%v.%v.%v.%v", parent[0], parent[1], parent[2], parent[3]), nil
}
