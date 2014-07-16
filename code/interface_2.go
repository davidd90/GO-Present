 type Rechteck struct{
	x,y float64
 }

 func (r *Rechteck) flaeche float64{
	return r.x*r.y
 }