package window

// TurnState 状态机
type TurnState int

const (
	// BeforePlayerAction is the state before the player takes an action
	BeforePlayerAction = iota
	// PlayerTurn is the state when the player is taking an action
	PlayerTurn
	// MonsterTurn is the state after the monster
	MonsterTurn
)

// GetNextState 获取下一个状态
func GetNextState(state TurnState) TurnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn
	case PlayerTurn:
		return MonsterTurn
	case MonsterTurn:
		return BeforePlayerAction
	default:
		return PlayerTurn
	}
}
