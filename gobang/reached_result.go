package gobang

type NeighborCellNumber int

const (
	OneSide NeighborCellNumber = iota + 1
	BothSides
)

func (num NeighborCellNumber) Value() int {
	return int(num)
}

func NewReachedResult(cells, emptyNeighborCells []*Cell, continuousEmptyCellCount int) *ReachedResult {
	var reachedType NeighborCellNumber
	emptyCellCount := len(emptyNeighborCells)

	switch emptyCellCount {
	default:
		reachedType = 0
	case OneSide.Value():
		reachedType = OneSide
	case BothSides.Value():
		reachedType = BothSides
	}

	return &ReachedResult{
		reachedType:              reachedType,
		cells:                    cells,
		emptyNeighborCells:       emptyNeighborCells,
		continuousEmptyCellCount: continuousEmptyCellCount,
	}
}

type ReachedResult struct {
	reachedType              NeighborCellNumber
	cells                    []*Cell
	emptyNeighborCells       []*Cell
	continuousEmptyCellCount int
}

func (result *ReachedResult) IsEmpty() bool {
	return len(result.cells) <= 0
}

func (result *ReachedResult) IsNeighborEmpty(num NeighborCellNumber) bool {
	return result.reachedType == num
}

func (result *ReachedResult) HasEmptyNeighborCell() bool {
	return len(result.emptyNeighborCells) > 0
}

func (result *ReachedResult) EmptyNeighborCells() []*Cell {
	return result.emptyNeighborCells
}

func (result *ReachedResult) ContinuousEmptyCellCount() int {
	return result.continuousEmptyCellCount
}
