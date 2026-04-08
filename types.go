package main

type Vec3 struct {
	X, Y, Z float64
}

type Command struct {
	PlayerUUID	string	`json: "playerUuid"`
	Animation	string  `json: "animation"`
	PlayerOrigin 	Vec3	`json: "playerOrigin"`
	PlayerRotation	Vec3	`json: "playerRotation"`
	PlayerScale	int	`json: "playerScale"`
}

type PlayerState struct {
	PlayerUUID	string  `json: "animation"`
	Origin		Vec3	`json: "origin"`
	Rotation	Vec3	`json: "rotation"`
	Scale		int	`json: "scale"`
}

type Event struct {
	PlayerUUID	string	`json: "playerUuid"`
	Animation	string	`json: "animation"`
	NewState	PlayerState `json: "newState"`
}
