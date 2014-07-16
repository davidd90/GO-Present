 func square (f float) float { return f*f }        // ein Rückgabewert

 func SumProd(i, j int) (int, int)  {              // mehrere Rückgabewerte
    sum := i+j
    prod := i*j
    return sum,prod
 }
 
 func SumProd(i, j int)(sum, prod int)  {          // bereits deklarierte Rückgabewerte                  
    sum = i+j
    prod = i*j
    return
 }