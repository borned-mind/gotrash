#include<iostream>

template <typename T>
	auto EvenGenerator(){
		T i = 0;
		return [&i]{
			T ret = i;
			i += 2;
			return ret;
		};
	}

int main(void){
	auto generator = EvenGenerator<unsigned int>();
	generator();
	generator();
	std::cout << generator() << std::endl;
}
