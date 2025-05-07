package navigation

var (
	validNavigation1 = `5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`

	validNavigation2 = `20 20
0 10 E
MMMMMMMMMMRMMMMMRMMMMMRMMMMMMMMMMMMMMM
20 10 w
MMMMMMMMMMRMMMMMRMMMMMRMMMMMMMMMMMMMMM
`

	validNavigationWithRoverFailInvalidPosition = `5 5
1 2 N
LMLMLMLMM
2 E
MMRMMRMRRM
5 0 W
MMM
`

	invalidNavigationAllRoversFailingToDeploy = `5 5
1 2 K
LMLMLMLMM
0.asdlkfj E
MMRMMRMRRM
`

	validNavigationWithRoverFailInvalidCoordinatesOutOfBounds = `5 5
1 2 N
LMLMLMLMM
5 6 E
MMRMMRMRRM
`

	validNavigationWithRoverFailInvalidCoordinatesNegative = `5 5
1 2 N
LMLMLMLMM
-1 5 E
MMRMMRMRRM
`

	validNavigationWithRoverFailInvalidOccupiedCoordinates = `5 5
1 2 N
LMLMLMLMM
1 3 S
MMRMMRMRRM
`

	validNavigationWithRoverGoingOutOfBounds = `5 5
1 2 N
LMLMLMLMM
3 3 E
MMMMMMMMMMMM
5 0 W
MMM
0 4 W
MMMMMRMRMMMML
`

	validNavigationWithRoverGoingOutOfBoundsNegative = `5 5
1 2 N
LMLMLMLMM
3 3 E
rmrMMMMMMMM
`

	validNavigationWithRoverGoingOutOfBoundsUpperOverflow = `5 18446744073709551615
1 2 N
LMLMLMLMM
1 18446744073709551613 s
LLMMMMMMM
`

	validNavigationWithRoverHittingOccupiedPosition = `5 5
1 2 N
LMLMLMLMM
5 0 W
MMM
0 0 E
MMMMMLMMMMM
`

	invalidPlateauCoordinates = `5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`

	validNavigationWithInvalidInstructions = `5 5
1 2 N
LMLMSSWEETLMLMM
3 3 E
MMRMMRABCMRRM
`

	validNavigationWithMissingInstructions = `5 5
1 2 N
LMLMSSWEETLMLMM
3 3 E
`
)
