 func square (f float) float { return f*f }        // ein R�ckgabewert

 func SumProd(i, j int) (int, int)  {              // mehrere R�ckgabewerte
    sum := i+j
    prod := i*j
    return sum,prod
 }
 
 func SumProd(i, j int)(sum, prod int)  {          // bereits deklarierte R�ckgabewerte                  
    sum = i+j
    prod = i*j
    return
 }