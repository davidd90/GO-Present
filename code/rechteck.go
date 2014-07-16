package main

import "fmt"


type Rechteck struct { 
   laenge, breite int
}

func (r Rechteck) flaeche() int {
   return r.laenge * r.breite
}

func main() {
   r := Rechteck{laenge:5, breite:3} 	
   fmt.Println("Werte des Rechtsecks: ",r)  
   fmt.Println("Flaeche des Rechtecks: ", r.flaeche())
}

