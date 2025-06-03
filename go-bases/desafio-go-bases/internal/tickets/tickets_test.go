/* Exercício 4:
Crie testes de unidade para cada um dos requisitos acima,
com um mínimo de 2 casos por requisito
*/

package tickets

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Change directory to the root of the project
	os.Chdir("../../../..")
}

func TestGetTicketsByDestination_EmptyDestination(t *testing.T) {
	expectedError := "destination cannot be empty"

	_, err := GetTicketsByDestination("")

	require.Error(t, err)
	require.EqualError(t, err, expectedError)
}

func TestGetTicketsByDestination_Brazil(t *testing.T) {
	expectedCount := 45

	count, err := GetTicketsByDestination("Brazil")

	require.NoError(t, err)
	require.Equal(t, expectedCount, count)
}

func TestGetTicketsByPeriod_InvalidOrEmptyPeriod(t *testing.T) {
	expectedEmptyError := "period cannot be empty"

	invalidPeriod := "dusk"
	expectedInvalidError := "invalid period: " + invalidPeriod

	_, err := GetTicketsByPeriod("")

	require.Error(t, err)
	require.EqualError(t, err, expectedEmptyError)

	_, err = GetTicketsByPeriod(invalidPeriod)

	require.Error(t, err)
	require.EqualError(t, err, expectedInvalidError)
}

func TestGetTicketsByPeriod_AllPeriods(t *testing.T) {
	tests := []struct {
		period        string
		expectedCount int
	}{
		{Dawn, 304},
		{Morning, 256},
		{Afternoon, 289},
		{Evening, 151},
	}
	expectedTotalCount := 1000
	totalCount := 0

	for _, tt := range tests {
		count, err := GetTicketsByPeriod(tt.period)
		totalCount += count
		require.NoError(t, err)
		require.Equal(t, tt.expectedCount, count)
	}
	require.Equal(t, expectedTotalCount, totalCount)
}

func TestDestinationPercentage_InvalidOrEmptyDestination(t *testing.T) {
	expectedEmptyError := "destination cannot be empty"
	invalidDestination := "Atlantis"
	expectedInvalidError := "no tickets found for destination: " + invalidDestination

	_, err := DestinationPercentage("")

	require.Error(t, err)
	require.EqualError(t, err, expectedEmptyError)

	_, err = DestinationPercentage(invalidDestination)

	require.Error(t, err)
	require.EqualError(t, err, expectedInvalidError)
}

func TestDestinationPercentage_Brazil(t *testing.T) {
	expectedCount := 0.045

	count, err := DestinationPercentage("Brazil")

	require.NoError(t, err)
	require.InDelta(t, expectedCount, count, 0.001, "Percentage should be approximately 4.5%")
}
