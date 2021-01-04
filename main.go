package main 

import ("fmt"
		"os"
		"time"
		"strings"
		
		"github.com/iotaledger/iota.go/kerl" 
		"github.com/iotaledger/iota.go/trinary"
		"github.com/iotaledger/iota.go/address"
	   )



func main() {

// target TOYAEEDHMGOZTEST9MBE9FXCRXHALLO99CGLAYSNIUMDFMDQ9AEJRFVDWQLEUPYWOQACLCFYRUOIBNQTN
/*	
	var in input
	in.seed= "TOYAEEDHMGOZFOWUZMBE9FXCRXCGLAYSNIUMDFMDQ9AEJRFVDWQLEUPYWOQACLCFYRUOIBNQKTYMCMTTN"
	in.word1 = "TEST9"
	in.word2 = "HALLO99"
	in.checksum = "LPC"
	in.address = "WZIIHYGN9VMWWYVNFXTPILVTLZJRULETTTSPWKELGZFSFRKFDHMXYPPJLSKD9XRNWCOTRPFHFV9JIIHQW"
*/

	in := getInput()
	findSeedCombi(in)
}


func getInput () input{
	var in input
	for {
		fmt.Print("Enter base seed: ")
		fmt.Scanln(&in.seed)
		if hasLength (in.seed,81){
			break
		}
	}

	for {
		fmt.Print("Enter known address: ")
		fmt.Scanln(&in.address)
		if hasLength (in.address,81){
			break
		}
	}
	
	for {
		fmt.Print("Enter seed checksum: ")
		fmt.Scanln(&in.checksum)
		if hasLength (in.checksum,3){
			break
		}
	}
	fmt.Print("Enter word 1: ")
	fmt.Scanln(&in.word1)

	fmt.Print("Enter word 2: ")
	fmt.Scanln(&in.word2)
	
	
	
	return in
}




func findSeedCombi(in input){
	t0:= time.Now()
	removeCombis := getRemoveCombis(in.seed,in.word1,in.word2)
	count :=0
	total := len(getInsertCombis(removeCombis[0],in.word1,in.word2))*len(removeCombis)
	fmt.Printf("\n%d possible seed combinations to check\n", total)
	for _,dSeed := range(removeCombis){
		insertCombis := getInsertCombis(dSeed,in.word1,in.word2)
		for _,iSeed := range(insertCombis){
			count +=1
			checksum := getChecksum(iSeed)
			if count%10000==0 {
				mSecs := time.Now().Sub(t0).Milliseconds()
				fmt.Printf("\rChecked combinations: %d  process: %.1f%%  speed: %dk seeds/sec", count, float64(count)/float64(total)*100, count/int(mSecs))
			}			
			if checksum == in.checksum && addressMatch(iSeed, in.address){
				fmt.Println("\n\nFound seed: ", iSeed)
				exitAfterEnter()
			}
		}
	}
	fmt.Println("\n\nCould not find a match.")
	exitAfterEnter()
	
}

func addressMatch(seed,targetAdd string) bool{
	addCount := 10
	seedTrits := trinary.MustTrytesToTrits(seed)
	seedTrits[242]=0
	trytes := trinary.MustTritsToTrytes(seedTrits)
	adds,err := address.GenerateAddresses(trytes, 0, uint64(addCount), 2, false)
	if err != nil {
		fmt.Println(err)
	}
	for _,a := range (adds) {
		if a == targetAdd {
			return true
		}
	}
	return false
}




func getChecksum(seed string) string{

	k:= kerl.NewKerl()
	seedTrits := trinary.MustTrytesToTrits(seed)
	seedTrits[242]=0
	k.Absorb(seedTrits)
	trits,_:=k.Squeeze(243)
	trytes := trinary.MustTritsToTrytes(trits)
	return trytes[len(trytes)-3:]
}


func getRemoveCombis(seed, word1, word2 string) []string{
	var combis  []string
	for i:=0;i<len(seed)-len(word1)+1;i++{
		removed1 := removeAt(seed,i, len(word1))
		for j:=0;j<len(removed1)-len(word2)+1;j++{
			removed2 := removeAt(removed1,j, len(word2))
			if !contains(&combis,removed2){
				combis = append(combis,removed2)
			}
		}
	}
	return combis
}



func getInsertCombis(seed string, word1 string, word2 string) []string{
	var combis []string
	for i:=0;i<len(seed)-len(word1)+1;i++{
		insert1 := insertAt (seed, word1, i)
		for j:=0;j<len(insert1)-len(word2)+1;j++{
			insert2:= insertAt(insert1,word2,j)
			if strings.Contains(insert2,word1)&&strings.Contains(insert2,word2){
				if !contains(&combis,insert2){
					combis=append(combis,insert2)
				}
			}
		}
	}
	return combis
}

func replaceAt (base string,  replace string, pos int) string{
	return  base[:pos] + replace + base[pos+len(replace):]
}

func insertAt (base string, insert string, pos int) string {
	return base[:pos] + insert + base[pos:]
}

func removeAt (base string, pos, length int) string{
	return base[:pos] + base[pos+length:]
}

func contains (slice *[]string, lookup string)bool{
	for _, s := range (*slice) {
		if s==lookup {
			return true
		}
	}
	return false
}



func addUnique(slice *[]string, s string){
	if !contains (slice, s) {
		*slice = append(*slice,s)
	}
}


func hasLength(text string, l int)bool{
	if len(text)==l {
		return true
	} else {
		fmt.Println ("\nInput is not",l," characters! Try again.\n")
		return false
	}	
}


func exitAfterEnter(){
	fmt.Println("\n\nPress Enter to exit.")
	k := ""
	fmt.Scanln(&k)
	os.Exit(0)
}

type input struct {
	seed, word1, word2, checksum, address string	
}
