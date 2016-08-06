package gobang

func NewGameContext(rule *GameRule) *GameContext {
  board := NewBoard(rule.BoardSize())

  return &GameContext {
    GameRule: rule,
    board: board,
    currentPlayer: nil,
    player: nil,
    npcPlayer: nil,
  }
}

type GameContext struct {
  *GameRule
  board *Board
  currentPlayer Player
  player *GamePlayer
  npcPlayer *NpcPlayer
}

func (g *GameContext) CurrentBoard() *Board {
  return g.board
}

func (g *GameContext) SelectBoardCell(point *Point) (*Cell, error) {
  if !g.board.HaveCell(point) {
    return nil, NewCellNotFoundError(point)
  }
  return g.board.SelectCell(point), nil
}

func (g *GameContext) CurrentPlayer() Player {
  return g.currentPlayer
}

func (g *GameContext) GamePlayer() *GamePlayer {
  return g.player
}

func (g *GameContext) NpcPlayer() *NpcPlayer {
  return g.npcPlayer
}

func (g *GameContext) ChangeToNextPlayer() {
  var player Player

  if (g.currentPlayer == g.npcPlayer) {
    player = g.player
  } else {
    player = g.npcPlayer
  }

  g.currentPlayer = player
}

func (g *GameContext) CheckBoard() PutStoneResult {
  player := g.CurrentPlayer()

  matcher := NewCellReachedMatcher(player.Stone(), g.ReachedStoneCount())
  result := matcher.Matches(g.board)

  if result.HasResult() {
    return Reached
  }

  if g.board.IsAllFilled() {
    return Filled
  }

  return Continue
}
