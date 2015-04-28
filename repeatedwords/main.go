package main

import "fmt"
import "strings"

type Text struct {
  value string
}

func (t *Text) splitString() []string {
  words := strings.Fields(t.value)
  
  return words
}

func (t *Text) countWords() int {
  return len(t.splitString())
}

func (t *Text) countRepeatedWords() map[string]int {
  words := t.splitString()
  
  dictionary := make(map[string]int)
  
  for i, _ := range words {
    for j, _ := range words {
      if i != j {
        if words[i] == words[j] {
          dictionary[words[i]] = dictionary[words[i]] + 1       
        }
      }
    }
  }

  return dictionary
}

func main() {
  var t Text  

  value := "go javascript php java c python ruby js php go c lisp"
  
  t.value = value

  fmt.Println(t.value)
  fmt.Println(t.countWords())
  fmt.Println(t.countRepeatedWords())
}