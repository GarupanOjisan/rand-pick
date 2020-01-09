# rand-pick

A Package for picking some items randomly based on a probability table.
It utilizes for like loot-boxes.

## usage

```go
r := &RandChoice{
    Probabilities: &map[int]interface{}{
        40: "N",
        30: "R",
        20: "SR",
        10: "UR",
    },
    Algorithm: algorythm.NewWeightedPicker(),
}
numOfPick := 10
items, err := r.Pick(numOfPick)
```