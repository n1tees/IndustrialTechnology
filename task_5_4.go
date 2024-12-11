func sort2Mass(mass1 []int, mass2 []int) []int {

    var resultmass []int = make([]int, len(mass1)+len(mass2))

    for i := 0; i < len(mass1); i++ {
        resultmass[i] = mass1[i]
    }

    for j := 0; j < len(mass2); j++ {
        resultmass[len(mass1)+j] = mass2[j]
    }

    for i := 0; i < len(resultmass); i++ {
        for j := i + 1; j < len(resultmass); j++ {
            if resultmass[i] > resultmass[j] {
                resultmass[i], resultmass[j] = resultmass[j], resultmass[i]
            }
        }
    }
    return resultmass
}

