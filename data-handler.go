package f12019

import (
	"encoding/binary"
	"gopkg.in/restruct.v1"
)

var MotionPackets = make(chan PacketMotionData, 1)
var SessionPackets = make(chan PacketSessionData, 1)
var LapPackets = make(chan PacketLapData, 1)
var EventPackets = make(chan PacketEventData, 1)
var ParticipantPackets = make(chan PacketParticipantsData, 1)
var CarSetupPackets = make(chan PacketCarSetupData, 1)
var CarTelemetryPackets = make(chan PacketCarTelemetryData, 1)
var CarStatusPackets = make(chan PacketCarStatusData, 1)

func handleData(data []byte) {
	headerBytes := data[0:24]

	var header PacketHeader
	unpackData(headerBytes, &header)

	switch header.PacketId {
	case 0:
		motionPacket := PacketMotionData{}
		unpackData(data, &motionPacket)

		if len(MotionPackets) != 1 {
			MotionPackets <- motionPacket
		}

		break
	case 1:
		sessionPacket := PacketSessionData{}
		unpackData(data, &sessionPacket)

		if len(SessionPackets) != 1 {
			SessionPackets <- sessionPacket
		}

		break
	case 2:
		lapPacket := PacketLapData{}
		unpackData(data, &lapPacket)

		if len(LapPackets) != 1 {
			LapPackets <- lapPacket
		}

		break
	case 3:
		eventPacket := PacketEventData{}
		unpackData(data, &eventPacket)

		if len(EventPackets) != 1 {
			EventPackets <- eventPacket
		}

		break
	case 4:
		participantPacket := PacketParticipantsData{}
		unpackData(data, &participantPacket)

		if len(ParticipantPackets) != 1 {
			ParticipantPackets <- participantPacket
		}

		break
	case 5:
		carSetupPacket := PacketCarSetupData{}
		unpackData(data, &carSetupPacket)

		if len(CarSetupPackets) != 1 {
			CarSetupPackets <- carSetupPacket
		}

		break
	case 6:
		carTelemetryPacket := PacketCarTelemetryData{}
		unpackData(data, &carTelemetryPacket)

		if len(CarTelemetryPackets) != 1 {
			CarTelemetryPackets <- carTelemetryPacket
		}

		break
	case 7:
		carStatusPacket := PacketCarStatusData{}
		unpackData(data, &carStatusPacket)

		if len(CarStatusPackets) != 1 {
			CarStatusPackets <- carStatusPacket
		}

		break
	}

}

func unpackData(data []byte, _struct interface{}) {
	restruct.Unpack(data, binary.LittleEndian, _struct)
}