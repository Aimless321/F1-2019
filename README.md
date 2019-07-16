# F1-2019
Package to interact with F1-2019 UDP data

## Usage

#### Starting the udp server
You have to start the udp server manually by running `f12019.Start()`.

So that would look something like this most of the time:
```
func main() {
    f12019.Start()

    //Do your things
}
```

#### Accessing data
Packets are stored in channels with a buffersize of 1, so it won't keep filling up your memory.

You can access all of these channels like so:
```
func main() {
    motionPacket := <-f12019.MotionPackets
    sessionPacket := <-f12019.SessionPackets
    lapPacket := <-f12019.LapPackets
    eventPacket := <-f12019.EventPackets
    participantPacket := <-f12019.ParticipantPackets
    carSetupPacket := <-f12019.CarSetupPackets
    carTelemetryPacket := <-f12019.CarTelemetryPackets
    carStatusPacket := <-f12019.CarStatusPackets
}
```

To see where you can find what data you can look in [structs.go](https://github.com/Aimless321/F1-2019/tree/master/structs.go) or check the official specification [here](https://forums.codemasters.com/topic/38920-f1-2019-udp-specification/)
