package main

import (
	"fmt"
	"regexp"
	"strings"
  "math"
	//"unicode"

	"bufio"
	"os"
)

func main() {
//-------------------------------Variables------------------------------
  var (
    mainText string
	  usrInput string
    usrWrdsContainer []string
	  kWrds uint = 0

    textAsSlice []string
    maxK float64
    op float64 = 0.15
  )

  in := bufio.NewScanner(os.Stdin)
//---------------------------Entering text------------------------------
  fmt.Println("Paste an article/essay on terminal and try to enter keywords of it...")

  for {
  	// Reads user input until \n by default
  	in.Scan()
  	mainText := in.Text()

    // Remove symbols from string
    re, err := regexp.Compile(`[^\w]`)
	  if err != nil {
      fmt.Println("Error:", err, ", try again.")
      mainText = ""; // To keep loop running
	  }
    mainText = re.ReplaceAllString(mainText, " ")

    mainText = strings.ToLower(mainText)
    
  	if len(mainText) != 0 {
      as := strings.Fields(mainText)
      // Copy elements of the inner scope slice as to a higher scope slice
      textAsSlice = append(textAsSlice, as...)
      break
	  }
    // Block for handling input error
    if in.Err() != nil {
    	fmt.Println("Error:", in.Err(), ", try again.")
    }
  }
//----------------------Processing text as slice------------------------
  maxK = float64(len(textAsSlice)) * op
  maxK = math.Ceil(maxK)
  
  for i := 0; i < len(textAsSlice); i++ {
    //Delete elements that are just numbers
    //r := []rune(textAsSlice[i])
    //if unicode.IsNumber(r[i]) {
      //textAsSlice[i] = ""
      // An alternative is to overwrite element to remove with the last index
      // and then return all slice without the last element
    //}
    //fmt.Println(textAsSlice[i])
  }
//-------------------------Entering keywords----------------------------
  fmt.Println("\n", `Now try to remember content of the text and enter keywords! (Enter "esc" to stop entering keywords...)`)
  fmt.Println("Needed keywords to approve:", maxK, "\n")

  for {
    // Reads user input until \n by default
  	in.Scan()
  	usrInput := in.Text()

    // Remove symbols from string
    re, err := regexp.Compile(`[^\w]`)
	  if err != nil {
      fmt.Println("Error:", err, ", try again.")
      usrInput = ""; // To keep loop running
	  }
    usrInput = re.ReplaceAllString(usrInput, " ")

    usrInput = strings.ToLower(usrInput)

    if len(usrInput) != 0 {
      // Allows multiple words input (separated by space)
      if usrInput != "esc"{
         ts := strings.Fields(usrInput)
      // Copy elements of the inner scope slice as to a higher scope slice
      usrWrdsContainer = append(usrWrdsContainer, ts...)
      }
      
	  }
    if usrInput == "esc" {
      break
    }
    
    // Block for handling input error
    if in.Err() != nil {
    	fmt.Println("Error:", in.Err(), ", try again.")
    }
  }
//---------------------------Compare slices-----------------------------
  // Remove dupes
  textAsSlice = rmvDupedString(textAsSlice)

  for i:=0; i<len(textAsSlice);i++{
    for j:=0; j<len(usrWrdsContainer);j++{
      if textAsSlice[i]==usrWrdsContainer[j] {
      //if strings.Contains(textAsSlice[i],usrWrdsContainer[j])==true {
        kWrds++
      }
    }
  }
//-------------------------------Results--------------------------------
  //fmt.Println(textAsSlice)
  fmt.Println(mainText)
  fmt.Println(usrInput)

  fmt.Println("Keywords entered by user:", len(usrWrdsContainer))
  fmt.Println("Correct Keywords (those inside the text):", kWrds)
  fmt.Println("Total Keywords:", len(textAsSlice))

  if float64(kWrds) >= maxK {
    fmt.Println("\nYou have good understanding of the subject :)")
  } else {
    fmt.Println("\nKeep going!, you need more practice")
  }
}
//---------------------------Helper methods-----------------------------
func rmvDupedString(strSlice []string) []string {
    ks := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := ks[item]; !value {
            ks[item] = true
            list = append(list, item)
        }
    }
    return list
}
