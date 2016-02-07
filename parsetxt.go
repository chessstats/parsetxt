package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	"regexp"
	"strings"
)

var root="hist_download"
var croot="hist_smartconv"

var gettxt=regexp.MustCompile(`\.txt|\.TXT$`)

func main() {
	fmt.Printf("parsetxt @ golang\n\n")
	fmt.Printf("creating %s\n\n",croot)
	if os.Mkdir(croot,os.ModeDir)!=nil {
		//panic("io error")
	}
	files,err:=ioutil.ReadDir(root)
	if err==nil {
		for _,file:=range files {
			name:=file.Name()
			if gettxt.MatchString(name) {
				fmt.Printf("processing %s\n",name)
				txtfile,err:=os.Open(root+string(os.PathSeparator)+name)
				if err==nil {
					name:=strings.ToLower(name)
					fscanner:=bufio.NewScanner(txtfile)
					fscanner.Scan()
					head:=strings.ToLower(fscanner.Text())
					invalid:=regexp.MustCompile("aaberg, anton").MatchString(head)
					invalid=invalid||regexp.MustCompile("apr03").MatchString(name)
					invalid=invalid||regexp.MustCompile("apr04").MatchString(name)
					invalid=invalid||regexp.MustCompile("jan05").MatchString(name)
					invalid=invalid||regexp.MustCompile("jan06").MatchString(name)
					invalid=invalid||regexp.MustCompile("jul04").MatchString(name)
					invalid=invalid||regexp.MustCompile("jul05").MatchString(name)
					invalid=invalid||regexp.MustCompile("oct02").MatchString(name)
					invalid=invalid||regexp.MustCompile("oct04").MatchString(name)


					if !invalid {
						headm:=regexp.MustCompile(`id number`).ReplaceAllString(head,"id_number")
						headm=regexp.MustCompile(`^  id_number`).ReplaceAllString(headm,"id_number  ")
						headm=regexp.MustCompile(`^   code`).ReplaceAllString(headm,"code   ")
						headm=regexp.MustCompile(`titlfed`).ReplaceAllString(headm,"tit fed")
						headm=regexp.MustCompile(`gamesborn`).ReplaceAllString(headm,"gms  born")
						headm=regexp.MustCompile(` [a-z]{3}[0-9]{2} `).ReplaceAllString(headm," rtg   ")
						headm=regexp.MustCompile(` [a-z]{3}[0-9]{1} `).ReplaceAllString(headm," rtg  ")
						headm=regexp.MustCompile(` [a-z]{4}[0-9]{2} `).ReplaceAllString(headm," rtg    ")
						headm=regexp.MustCompile(` [a-z]{4}[0-9]{1} `).ReplaceAllString(headm," rtg   ")
						headm=regexp.MustCompile(`title`).ReplaceAllString(headm,"tit  ")
						headm=regexp.MustCompile(`country`).ReplaceAllString(headm,"fed    ")
						headm=regexp.MustCompile(`birthday`).ReplaceAllString(headm,"born    ")
						headm=regexp.MustCompile(`b-day`).ReplaceAllString(headm,"born ")
						token:=false
						begins:=[]int{}
						tokens:=[]string{}
						for i,c:=range headm {
							if c!=' ' {
								if !token {									
									begins=append(begins,i)
									token=true									
								} else {
									if (i==(len(headm)-1)) {
										tokens=append(tokens,headm[begins[len(begins)-1]:len(headm)])
									}
								}
							} else {
								if token {
									tokens=append(tokens,headm[begins[len(begins)-1]:i])
								}
								token=false
							}
						}
						begins=append(begins,len(headm))												
						fmt.Printf("%s\n%s\n%v\n%v\n",head,headm,begins,tokens)
						content:=""
						clen:=0
						for i,cname:=range tokens {
							if i>0 {
								content+=" "
							}
							content+=`"`+cname+`"`
						}
						content+="\r\n"
						outf,err:=os.OpenFile(croot+string(os.PathSeparator)+name,os.O_CREATE|os.O_WRONLY,0666)
						if err!=nil {
						    panic(err)
						}
						if _,err:=outf.WriteString(content); err!=nil {
							panic(err)
						}
						clen+=len(content)
						rownum:=1
						for fscanner.Scan() {
							line:=fscanner.Text()							
							buff:=""
							for i:=0; i<(len(begins)-1); i++ {
								if buff=="" {
									buff+=`"`+fmt.Sprintf("%d",rownum)+`" `
									rownum++
								} else {
									buff+=" "
								}
								var ll=begins[i+1]
								if ll>len(line) {
									ll=len(line)
								}
								if begins[i]>=len(line) {
									buff+=`""`
								} else {
									value:=line[begins[i]:ll]
									value=regexp.MustCompile(`^ +| +$`).ReplaceAllString(value,"")
									buff+=`"`+value+`"`
								}
							}
							content=buff+"\r\n"
							if _,err:=outf.WriteString(content); err!=nil {
								panic(err)
							}
							clen+=len(content)
						}
						outf.Close()
						fmt.Printf("lines %d content length %d\n",rownum-1,clen)
					}					
					txtfile.Close()					
				} else {
					panic("io error")
				}
			}
		}
	} else {
		panic("unable to open directory")
	}
}