package math

func hidden(){
//	fmt.Println("is hidden")
}


//get minimum of slice of float64
func Min(xs []float64) (ret float64){
	if(len(xs) == 0){
		ret = 0
		return
	}
	ret=xs[0];

	for _, v := range xs{
		if( ret > v ){
			 ret = v
		}
	}
	return
}

//get maximum of slice of float64
func Max(xs []float64) (ret float64){
	if(len(xs) == 0){
		return
	}

	ret = xs[0]
	for _, v := range xs{
		if( ret < v ) {
			ret = v
		}
	}
	return
}
//get average of slice of float64
func Average(xs []float64) float64{
	if( len(xs) == 0 ) { return 0 }
	var r float64
	for _, v := range xs{
		r+=v
	}
	return r/float64(len(xs))
}
