package main

import "fmt"

func maxProduct(words []string) int {
    if len(words) < 2 {
	return 0
    }

    //quick sort by word length
    quickSortByLength(words, 0, len(words) -1)

    //generate bit map array
    wordBitMapArr := generateWordBitMapArr(words)

    //iterate words from left to right
    maxProduct := 0
    for right := 1; right < len(words); right++ {
	for left := 0; left < right; left++ {
	    //find no-char shared word couple
	    if (wordBitMapArr[left] & wordBitMapArr[right] == 0) {
		current := len(words[left]) * len(words[right])
		if current > maxProduct {
		    maxProduct = current
		}
	    }
	}
    }

    return maxProduct
}

func quickSortByLength(words []string, start, end int) {
    if start >= end {
	return
    }

    mid := partion(words, start, end)
    quickSortByLength(words, start, mid-1)
    quickSortByLength(words, mid + 1, end)
}

func partion(words []string, start, end int) int{
    pivot := words[end]
    mid := start
    for i := start; i < end; i++ {
	if len(words[i]) > len(pivot) {
	    words[i], words[mid] = words[mid], words[i]
	    mid++
	}
    }

    words[end], words[mid] = words[mid], words[end]
    return mid
}

func generateWordBitMapArr(words []string) []int {
    result := []int{}
    for _, word := range words {
	bitWord := 0
	for _, char := range word {
	    bitWord = bitWord | (1 << uint((char - 'a')))
	}
	result = append(result, bitWord)
    }

    return result
}

func main() {
   // words := []string{"abcw","baz","foo","bar","xtfn","abcdef"}
   // fmt.Println(maxProduct(words))
   // words2 := []string{"a","ab","abc","d","cd","bcd","abcd"}
   // fmt.Println(maxProduct(words2))
   // words3 := []string{"a","aa","aaa","aaaa"}
   // fmt.Println(maxProduct(words3))
    word4 := []string{"ccd","accaceddeeeaefc","bcaffa","bbcfafbb","accacfebbabbeedfbfdb","beddecbffcdaededdaefdedfdea","cf","ddafdcbd","bbafacebacaefdaffccebddff","ebccffcddbeddccacceccaec","becfbfdccdfdeadfbfaddbcded","cbabeaaeabefedbaeaedc","dfadbbdbead","cafaefdcd","eccdbfceafeeeacfcddc","dbabbcdbb","abfbfbffcbebde","cfaadaa","fc","faebcabb","adbacebabcaaccbdeaffff","aeaefccf","dbacbeeabdbcdfccabebaecfef","ecdadeefcaddffaececffa","defcabf","abbcecbccbdaebaecaefabed","dfeeebcbaaefc","aecccbcbbdddb","dcfabacec","fccfbacbacddeaaea","dfdbfacbacbecb","cbfeebdbfecb","cffaacacbde","aafd","bdcebbbebd","afeffadcfcdacfba","dafeefbcdfaffcfacee","dcbbebfbedafedcdbab","cafaf","bcbcccfdebdd","efaaaacccff","cffbead","ebcfccfcddffdec","fffdfdcec","beeafefbdfa","cdfdbccfbaaeffcabab","ddadcbabbcb","decfaeabbecebaebeaddedae","cdcbfffbebae","aeccefcbcbbddfdc","ffefedaf","cddbabccafaffeafeedcbedbdad","eddeeccfedcefadfdfebfacb","aca","ffdcafaddcddf","ef","bbbbffe","ffccfebabaadcffacbbb","cbdeddfddffacbeeeebafebabda","ddeecb","cffdc","edcffcebadf","becbcadcafddcfbbeeddbfffcab","abcbaceeaeaddd","cfeffceebfaeefadaaccfa","eaccddb","caeafbfafecd","becaafdbaadbfecfdfde","ecabaaeafbfbcbadaac","bdcdffcfaeebeedfdfddfaf","dbbfbaeecbfcdebad","cceecddeeecdbde","beec","adbcfdbfdbccdcffffbcffbec","bbbbfe","cdaedaeaad","dadbfeafadd","fcacaaebcedfbfbcddfc","ceecfedceac","dada","ccfdaeffbcfcc","eadddbbbdfa","beb","fcaaedadabbbeacabefdabe","dfcddeeffbeec","defbdbeffebfceaedffbfee","cffadadfbaebfdbadebc","fbbadfccbeffbdeabecc","bdabbffeefeccb","bdeeddc","afcbacdeefbcecff","cfeaebbbadacbced","edfddfedbcfecfedb","faed","cbcdccfcbdebabc","efb","dbddadfcddbd","fbaefdfebeeacbdfbdcdddcbefc","cbbfaccdbffde","adbcabaffebdffad"}
    fmt.Println(maxProduct(word4))
}
