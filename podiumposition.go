package student

func PodiumPosition(podium [][]string) [][]string {
	result := make([][]string, len(podium))
	for i := 0; i < len(podium); i++ {
		result[i] = podium[len(podium)-1-i]
	}
	return result
}
