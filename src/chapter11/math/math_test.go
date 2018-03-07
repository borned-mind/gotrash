package math
import "testing"

type pair_test struct{
	v []float64
	average float64
}


var test_average = []pair_test{
    { []float64{1,2}, 1.5 },
    { []float64{1,1,1,1,1,1}, 1 },
    { []float64{-1,1}, 0 },
}

var test_min = []pair_test{
	{[]float64{1,2,3,4,5,6,7}, 1},
	{[]float64{228,666,999}, 228},
}

var test_max = []pair_test{
        {[]float64{1,2,3,4,5,6,7}, 7},
        {[]float64{228,666,999}, 999},
}


func TestMin(t *testing.T){
        for _, v := range test_min{
                //tmp := Average(v.v)
                if( v.average != Min(v.v) ){
                        t.Error("Expected, not got ", v.v, v.average)
                }
        }

}

func TestMax(t *testing.T){
        for _, v := range test_max{
                //tmp := Average(v.v)  
                if( v.average != Max(v.v) ){
                        t.Error("Expected, not got ", v.v, v.average)
                }
        }


}




func TestAverage(t *testing.T){
//	t.Error("test")
	for _, v := range test_average{
		//tmp := Average(v.v)
		if( v.average != Average(v.v) ){
			t.Error("Expected, not got ", v.v, v.average)
		}
	}

}
