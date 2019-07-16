package f12019

/*
Each packet has the following header
 */
type PacketHeader struct {
	PacketFormat     uint16
	GameMajorVersion uint8
	GameMinorVersion uint8
	PacketVersion    uint8
	PacketId         uint8
	SessionUID       uint64
	SessionTime      float32
	FrameIdentifier  uint32
	PlayerCarIndex   uint8
}





/*
The motion packet gives physics data for all the cars being driven.
There is additional data for the car being driven with the goal of being able to drive a motion platform setup.
*/
type CarMotionData struct {
	WorldPositionX     float32 
	WorldPositionY     float32 
	WorldPositionZ     float32 
	WorldVelocityX     float32 
	WorldVelocityY     float32 
	WorldVelocityZ     float32 
	WorldForwardDirX   int16   
	WorldForwardDirY   int16   
	WorldForwardDirZ   int16   
	WorldRightDirX     int16   
	WorldRightDirY     int16   
	WorldRightDirZ     int16   
	GForceLateral      float32 
	GForceLongitudinal float32 
	GForceVertical     float32 
	Yaw                float32 
	Pitch              float32 
	Roll               float32 
}

type PacketMotionData struct {
	Header PacketHeader 

	CarMotionData [20]CarMotionData 

	
	SuspensionPosition     [4]float32 
	SuspensionVelocity     [4]float32 
	SuspensionAcceleration [4]float32 
	WheelSpeed             [4]float32 
	WheelSlip              [4]float32 
	LocalVelocityX         float32    
	LocalVelocityY         float32    
	LocalVelocityZ         float32    
	AngularVelocityX       float32    
	AngularVelocityY       float32    
	AngularVelocityZ       float32    
	AngularAccelerationX   float32    
	AngularAccelerationY   float32    
	AngularAccelerationZ   float32    
	FrontWheelsAngle       float32    
}

/*
The session packet includes details about the current session in progress.
*/
type MarshalZone struct {
	ZoneStart float32 
	ZoneFlag  int8    
}

type PacketSessionData struct {
	Header PacketHeader 

	Weather             uint8           
	TrackTemperature    int8            
	AirTemperature      int8            
	TotalLaps           uint8           
	TrackLength         uint16          
	SessionType         uint8           
	TrackId             int8            
	Formula             uint8
	SessionTimeLeft     uint16          
	SessionDuration     uint16          
	PitSpeedLimit       uint8           
	GamePaused          bool		`struct:"uint8"`
	IsSpectating        bool		`struct:"uint8"`
	SpectatorCarIndex   uint8           
	SliProNativeSupport bool		`struct:"uint8"`
	NumMarshalZones     uint8           
	MarshalZones        [21]MarshalZone 
	SafetyCarStatus     uint8           
	NetworkGame         bool		`struct:"uint8"`
}

/*
The lap data packet gives details of all the cars in the session.
*/
type LapData struct {
	LastLapTime       float32 
	CurrentLapTime    float32 
	BestLapTime       float32 
	Sector1Time       float32 
	Sector2Time       float32 
	LapDistance       float32 
	TotalDistance     float32 
	SafetyCarDelta    float32 
	CarPosition       uint8   
	CurrentLapNum     uint8   
	PitStatus         uint8   
	Sector            uint8   
	CurrentLapInvalid bool		`struct:"uint8"`
	Penalties         uint8   
	GridPosition      uint8   
	DriverStatus      uint8   
	ResultStatus      uint8   
}

type PacketLapData struct {
	Header PacketHeader 

	LapData [20]LapData 
}


/*
This packet gives details of events that happen during the course of a session.
*/
type PacketEventData struct {
	Header PacketHeader 

	EventStringCode string `struct:"[4]uint8"`
	EventData EventData
}

type EventData struct {
	VehicleId  uint8
	LapTime    float32
}


/*
This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name.
If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.
*/
type ParticipantData struct {
	AiControlled        bool	`struct:"uint8"`
	DriverId            uint8
	TeamId              uint8
	RaceNumber          uint8
	Nationality         uint8
	Name                string	`struct:"[48]byte"`
	RestrictedTelemetry uint8
}

type PacketParticipantsData struct {
	Header PacketHeader 

	NumCars      uint8 
	Participants [20]ParticipantData
}

/*
This packet details the car setups for each vehicle in the session.
Note that in multiplayer games, other player cars will appear as blank,
you will only be able to see your car setup and AI cars.
*/
type CarSetupData struct {
	FrontWing             uint8   
	RearWing              uint8   
	OnThrottle            uint8   
	OffThrottle           uint8   
	FrontCamber           float32 
	RearCamber            float32 
	FrontToe              float32 
	RearToe               float32 
	FrontSuspension       uint8   
	RearSuspension        uint8   
	FrontAntiRollBar      uint8   
	RearAntiRollBar       uint8   
	FrontSuspensionHeight uint8   
	RearSuspensionHeight  uint8   
	BrakePressure         uint8   
	BrakeBias             uint8   
	FrontTyrePressure     float32 
	RearTyrePressure      float32 
	Ballast               uint8   
	FuelLoad              float32 
}

type PacketCarSetupData struct {
	Header PacketHeader 

	CarSetups [20]CarSetupData
}

/*
This packet details telemetry for all the cars in the race.
It details various values that would be recorded on the car such as speed, throttle application, DRS etc.
*/
type CarTelemetryData struct {
	Speed                   uint16     
	Throttle                float32
	Steer                   float32
	Brake                   float32
	Clutch                  uint8      
	Gear                    int8       
	EngineRPM               uint16     
	DrsOpen                 bool		`struct:"uint8"`
	RevLightsPercent        uint8      
	BrakesTemperature       [4]uint16  
	TyresSurfaceTemperature [4]uint16  
	TyresInnerTemperature   [4]uint16  
	EngineTemperature       uint16     
	TyresPressure           [4]float32
	SurfaceType             [4]float32
}

type PacketCarTelemetryData struct {
	Header PacketHeader 

	CarTelemetryData [20]CarTelemetryData

	ButtonStatus uint32 
}

/*
This packet details car statuses for all the cars in the race. It includes values such as the damage readings on the car.
*/
type CarStatusData struct {
	TractionControl         uint8    
	AntiLockBrakesOn        bool		`struct:"uint8"`
	FuelMix                 uint8    
	FrontBrakeBias          uint8    
	PitLimiterOn        	bool		`struct:"uint8"`
	FuelInTank              float32  
	FuelCapacity            float32
	FuelRemainingLaps		float32
	MaxRPM                  uint16   
	IdleRPM                 uint16   
	MaxGears                uint8    
	DrsAllowed              uint8    
	TyresWear               [4]uint8 
	ActualTyreCompound      uint8
	VisualCompund           uint8
	TyresDamage             [4]uint8 
	FrontLeftWingDamage     uint8    
	FrontRightWingDamage    uint8    
	RearWingDamage          uint8    
	EngineDamage            uint8    
	GearBoxDamage           uint8
	VehicleFiaFlags         int8     
	ErsStoreEnergy          float32  
	ErsDeployMode           uint8    
	ErsHarvestedThisLapMGUK float32  
	ErsHarvestedThisLapMGUH float32  
	ErsDeployedThisLap      float32  
}

type PacketCarStatusData struct {
	Header PacketHeader 

	CarStatusData [20]CarStatusData
}