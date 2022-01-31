/*
Problem #3

Given a string, say S, print it in reverse manner eliminating the repeated characters and spaces.

Example 1:

Input: S = "GEEKS FOR GEEKS"
Output: "SKEGROF"

Example 2:

Input: S = "I AM AWESOME"
Output: "EMOSWAI"

-------------------
Problem#3 has been asked in the interview by samsung.
*/

package main

package main
import (
    "fmt"
    "strings")

func RunLengthEncode(input string){
    s:=string(input[0])
  for i:=0;i<(len(input));i++ {
      if strings.Contains(s,string(input[i]))==false{
          s+=string(input[i])
      }
    }
    fmt.Println(s)
}
                           //res1 := strings.Count(str1, "o")                          
func main() {
    var S string
    fmt.Scanln(&S)
    RunLengthEncode(S)
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}