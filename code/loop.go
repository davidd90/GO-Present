//for:

 for i := 0; i < 10; i++ {}
 

 //while:

 for <condition> {}
  


 //range:

 for key := range m {        //Iteration ueber ein Array
    if key.expired() {
        delete(m, key)
    }
 }