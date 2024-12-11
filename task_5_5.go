func findSubstring(mainString string, subString string) int {
    if len(subString) == 0 {
        return 0
    }
    for i := 0; i <= len(mainString)-len(subString); i++ {
        match := true
        for j := 0; j < len(subString); j++ {
            if mainString[i+j] != subString[j] {
                match = false
                break
            }
        }
        if match {
            return i
        }
    }
    return -1
}
