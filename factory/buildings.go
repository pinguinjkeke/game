package factory

import (
	"game/archetype"
	"game/component"
	"github.com/yohamta/donburi/ecs"
	"math/rand"
)

const (
	minBuildingWidth = 40
	maxBuildingGrowX = 60
	windowMinSize    = 5
)

func CreateBuildings(ecs *ecs.ECS, levelWidth, levelHeight float64) {
	ground := levelHeight * 0.7
	minBuildingHeight := int(levelHeight / 4)
	maxBuildingGrowY := int(levelHeight / 4)
	buildingsEntry := archetype.Buildings.Spawn(ecs)
	buildingsLen := int(levelWidth / minBuildingWidth)
	buildings := make([]*component.BuildingData, buildingsLen)
	offset := 0

	for i := 0; i < buildingsLen; i++ {
		width, height := minBuildingWidth+rand.Intn(maxBuildingGrowX), minBuildingHeight+rand.Intn(maxBuildingGrowY)

		if rand.Intn(3) == 1 {
			width, height = height, width
		}

		x, y := offset+rand.Intn(50), int(ground)-height

		windowWidth, windowHeight := windowMinSize+rand.Intn(windowMinSize), windowMinSize+rand.Intn(windowMinSize)
		windowOffsetX, windowOffsetY := int(width-windowWidth)%int(windowWidth), int(height-windowHeight)%int(windowHeight)
		if windowOffsetX == 0 {
			windowOffsetX = 2
		}
		if windowOffsetY == 0 {
			windowOffsetY = 2
		}

		windowRows, windowColumns := int((width-windowOffsetX)/(windowWidth+windowOffsetX)), int((height-windowOffsetY)/(windowHeight+windowOffsetY))
		windowLights := make([]bool, windowRows*windowColumns)
		currentWindow := 0

		for i := 0; i < windowRows; i++ {
			for j := 0; j < windowColumns; j++ {
				windowLights[currentWindow] = rand.Intn(4) != 1

				currentWindow++
			}
		}

		buildings[i] = &component.BuildingData{
			X:             x,
			Y:             y,
			Width:         width,
			Height:        height,
			WindowWidth:   windowWidth,
			WindowHeight:  windowHeight,
			WindowRows:    windowRows,
			WindowColumns: windowColumns,
			WindowOffsetX: windowOffsetX,
			WindowOffsetY: windowOffsetY,
			WindowLights:  windowLights,
			Layer:         rand.Intn(2),
		}

		offset += minBuildingWidth
	}

	component.Buildings.Set(buildingsEntry, &component.BuildingsData{
		Buildings: buildings,
	})
}
