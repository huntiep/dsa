package main


func main() {
    /*
    testArray()
    testDList()
    testBT()
    */
    testTrie()
}

func assert(pred bool, msg string) {
    if !pred {
        panic(msg)
    }
}
