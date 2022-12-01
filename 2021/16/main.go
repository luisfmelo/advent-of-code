package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
)

const inputPath = "2021/16/input.txt"

const (
	TypeIDSumPacket         = 0
	TypeIDProductPacket     = 1
	TypeIDMinimumPacket     = 2
	TypeIDMaximumPacket     = 3
	TypeIDLiteralValue      = 4
	TypeIDGreaterThanPacket = 5
	TypeIDLessThanPacket    = 6
	TypeIDEqualToPacket     = 7

	LiteralValueNumberOfBits = 5
)

type Packet struct {
	version    int
	typeID     int
	value      int
	subPackets []Packet

	remaining string
}

func (p Packet) CalculateVersionSum() int {
	sum := p.version
	for _, subPacket := range p.subPackets {
		sum += subPacket.CalculateVersionSum()
	}

	return sum
}

var hexaToBinMap = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func hexaToBin(hexa string) string {
	var bin string
	for _, h := range hexa {
		bin += hexaToBinMap[h]
	}

	return bin
}

func binToDec(bin string) int {
	var i int
	for idx, b := range bin {
		if b == '1' {
			i += int(math.Pow(2, float64(len(bin)-1-idx)))
		}
	}

	return i
}

func ProcessLiteralValue(binRep string) (lv int, remaining string) {
	remaining = ""
	literalValueStr := ""

	// process literal packet
	for stIndex := 0; stIndex < len(binRep)-LiteralValueNumberOfBits+1; stIndex += LiteralValueNumberOfBits {
		group := binRep[stIndex : stIndex+LiteralValueNumberOfBits]
		literalValueStr += group[1:]
		if group[0] == '0' {
			remaining = binRep[stIndex+LiteralValueNumberOfBits:]

			break
		}
	}

	lv = binToDec(literalValueStr)

	return lv, remaining
}

func ProcessPacket(binRep string) Packet {
	packet := Packet{
		version: binToDec(binRep[:3]),
		typeID:  binToDec(binRep[3:6]),
	}

	switch packet.typeID {
	case TypeIDLiteralValue:
		lv, remaining := ProcessLiteralValue(binRep[6:])
		packet.value = lv
		packet.remaining = remaining

	default:
		switch binRep[6] {
		case '0': // total number of bits os the sub packets (read next 15)
			totalBits := binToDec(binRep[7:22])
			subPacketBinRep := binRep[22 : 22+totalBits]
			for len(subPacketBinRep) > 0 {
				p := ProcessPacket(subPacketBinRep)
				packet.subPackets = append(packet.subPackets, p)
				subPacketBinRep = p.remaining
			}
			packet.remaining = binRep[22+totalBits:]
		case '1': // total number of sub packets (read next 11 bits)
			numberOfSubPackets := binToDec(binRep[7:18])
			subPacketBinRep := binRep[18:]
			for len(packet.subPackets) < numberOfSubPackets {
				p := ProcessPacket(subPacketBinRep)
				packet.subPackets = append(packet.subPackets, p)
				subPacketBinRep = p.remaining
			}
			packet.remaining = subPacketBinRep
		}

		switch packet.typeID {
		case TypeIDSumPacket:
			for _, subPacket := range packet.subPackets {
				packet.value += subPacket.value
			}
		case TypeIDProductPacket:
			packet.value = 1
			for _, subPacket := range packet.subPackets {
				packet.value *= subPacket.value
			}
		case TypeIDMinimumPacket:
			packet.value = math.MaxInt
			for _, subPacket := range packet.subPackets {
				if subPacket.value < packet.value {
					packet.value = subPacket.value
				}
			}
		case TypeIDMaximumPacket:
			packet.value = -1
			for _, subPacket := range packet.subPackets {
				if subPacket.value > packet.value {
					packet.value = subPacket.value
				}
			}
		case TypeIDGreaterThanPacket:
			if packet.subPackets[0].value > packet.subPackets[1].value {
				packet.value = 1
			}
		case TypeIDLessThanPacket:
			if packet.subPackets[0].value < packet.subPackets[1].value {
				packet.value = 1
			}
		case TypeIDEqualToPacket:
			if packet.subPackets[0].value == packet.subPackets[1].value {
				packet.value = 1
			}
		}
	}

	return packet
}

func SumVersionsOfPackets(encodedTransmission string) int {
	return ProcessPacket(hexaToBin(encodedTransmission)).CalculateVersionSum()
}
func CalculateValueOfOutermostPacket(encodedTransmission string) int {
	return ProcessPacket(hexaToBin(encodedTransmission)).value
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error occurred: %v", err)
		}
	}()

	file, err := os.Open(inputPath)
	pkg.PanicErr(err)

	r := bufio.NewReader(file)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	encodedTransmission, err := pkg.ReadLine(scanner)
	pkg.PanicErr(err)

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", SumVersionsOfPackets(encodedTransmission)) },
		func() string {
			return fmt.Sprintf("%v", CalculateValueOfOutermostPacket(encodedTransmission))
		},
	)
}
