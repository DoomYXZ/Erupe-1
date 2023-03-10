package channelserver

import (
	"erupe-ce/common/byteframe"
	ss "erupe-ce/common/stringsupport"
	"erupe-ce/network/mhfpacket"
	"go.uber.org/zap"
	"time"
)

func handleMsgMhfGetRyoudama(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetRyoudama)
	var data [][]byte
	switch pkt.Unk0 {
	case 0:
		switch pkt.Unk1 {
		case 4: // get_karidama_point
			temp := byteframe.NewByteFrame()
			temp.WriteUint32(177536)
			data = [][]byte{
				temp.Data(),
			}
		case 5: // get_luckymenber
			temp := byteframe.NewByteFrame()
			var i, j uint32
			for i = 0; i < 7; i++ {
				j++
				temp.WriteUint32(j)
				temp.WriteUint32(i + 1)
				temp.WriteBytes(ss.PaddedString("TestChar", 14, true))
				j++
				temp.WriteUint32(j)
				temp.WriteUint32(i + 1)
				temp.WriteBytes(ss.PaddedString("TestChar", 14, true))
			}
			data = append(data, temp.Data())
		case 6: // get_boost
			temp := byteframe.NewByteFrame()
			temp.WriteUint32(uint32(TimeAdjusted().Add(-1 * time.Hour).Unix()))
			temp.WriteUint32(uint32(TimeAdjusted().Add(1 * time.Hour).Unix()))
			data = [][]byte{
				temp.Data(),
			}
		}
	}
	if len(data) == 0 {
		s.logger.Warn("Unknown GetRyoudama request", zap.Uint8("Unk0", pkt.Unk0), zap.Uint8("Unk1", pkt.Unk1))
	}
	doAckEventEnum(s, pkt.AckHandle, data)
}

func handleMsgMhfPostRyoudama(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfGetTinyBin(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetTinyBin)
	// requested after conquest quests
	doAckBufSucceed(s, pkt.AckHandle, []byte{})
}

func handleMsgMhfPostTinyBin(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfPostTinyBin)
	doAckSimpleSucceed(s, pkt.AckHandle, make([]byte, 4))
}

func handleMsgMhfCaravanMyScore(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfCaravanRanking(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfCaravanMyRank(s *Session, p mhfpacket.MHFPacket) {}
