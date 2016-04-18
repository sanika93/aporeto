package main

import (
"fmt"
"flag"
"strings"
"regexp"
"io/ioutil"
"net/http"
"github.com/jaytaylor/html2text"
"os"
"strconv"
)

var getURLS= flag.String("-urls", "<url1>[,<url2>[,<url3>[,<url4>]]]]", "")

func check(e error){
   if(e != nil){
	panic(e)	
	}
}

func extract_words(text string) []string{
  words:= regexp.MustCompile("\\w+")
  return words.FindAllString(text, -1)
}

func count_words(words []string) map[string]int{
	word_counts := make(map[string]int)
	for _,word := range words{
		word_counts[word]++		
	}
return word_counts
}

func console_out(word_counts map[string]int, name string){
	f,err := os.Create(name + ".txt")
	urlname := string("url" + ":" + name)
	f.WriteString(urlname + "\n")
 for word,word_count:= range word_counts{
	t:= strconv.Itoa(word_count)
	writestring := word + ":" + t + "\n"
        f.WriteString(writestring)
	check(err)
}
}

func main() {
	flag.Parse()
	result := strings.Split(*getURLS,",")
	for i:=range result{
		re:=regexp.MustCompile("\\[|\\]")
		repl:=[]byte("")
		src:=[]byte(result[i])
		newByte:=re.ReplaceAll(src,repl)
		newUrl:= "http://" + string(newByte)	
		resp, err := http.Get(newUrl)
		if err != nil {
		fmt.Println("error")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		text ,err := html2text.FromString(string([]byte(body)))
		console_out(count_words(extract_words(text)), string(newByte))
		}

}



