#include "lab1.hpp"

int main()
{
    long long size = 5000000;
    double *arr = new double[size];
    std::clock_t start = std::clock();

    lab1<double>::random_seq(arr,size,1,10);
    lab1<double>::random_seq(arr,size,1,10);
    lab1<double>::random_seq(arr,size,1,10);
    lab1<double>::random_seq(arr,size,1,10);
    lab1<double>::random_seq(arr,size,1,10);

    std::clock_t end = std::clock();

    double duration = static_cast<double>(end - start) / CLOCKS_PER_SEC;

    std::cout << "Время работы функции random_seq: " << duration/5 << " секунд" << std::endl;


start = std::clock();
    lab1<double>::sorted(arr,size,1,1,0);
    lab1<double>::sorted(arr,size,1,1,0);
    lab1<double>::sorted(arr,size,1,1,0);
    lab1<double>::sorted(arr,size,1,1,0);
    lab1<double>::sorted(arr,size,1,1,0);

    end = std::clock();
    duration = static_cast<double>(end - start) / CLOCKS_PER_SEC;

    std::cout << "Время работы функции sorted: " << duration/5 << " секунд" << std::endl;

    start = std::clock();
    lab1<double>::sorted(arr,size,1,1,1);
    lab1<double>::sorted(arr,size,1,1,1);
    lab1<double>::sorted(arr,size,1,1,1);
    lab1<double>::sorted(arr,size,1,1,1);
    lab1<double>::sorted(arr,size,1,1,1);

    end = std::clock();
    duration = static_cast<double>(end - start) / CLOCKS_PER_SEC;

    std::cout << "Время работы функции sorted(reversed): " << duration/5 << " секунд" << std::endl;

    start = std::clock();
    lab1<double>::chainsaw_seq(arr,size,5,1,5);
    lab1<double>::chainsaw_seq(arr,size,5,1,5);
    lab1<double>::chainsaw_seq(arr,size,5,1,5);
    lab1<double>::chainsaw_seq(arr,size,5,1,5);
    lab1<double>::chainsaw_seq(arr,size,5,1,5);
    end = std::clock();
    
    duration = static_cast<double>(end - start) / CLOCKS_PER_SEC;
    std::cout << "Время работы функции chainsaw_seq: " << duration/5 << " секунд" << std::endl;


    start = std::clock();
    lab1<double>::stepper_seq(arr,size, 10, 5,8, 3);
    lab1<double>::stepper_seq(arr,size, 10, 5,8, 3);
    lab1<double>::stepper_seq(arr,size, 10, 5,8, 3);
    lab1<double>::stepper_seq(arr,size, 10, 5,8, 3);
    lab1<double>::stepper_seq(arr,size, 10, 5,8, 3);
    end = std::clock();

    duration = static_cast<double>(end - start) / CLOCKS_PER_SEC;
    std::cout << "Время работы функции stepper_seq: " << duration/5 << " секунд" << std::endl;

        start = std::clock();
    lab1<double>::quazy_seq(arr,size, 10,4);
    lab1<double>::quazy_seq(arr,size, 10,4);
    lab1<double>::quazy_seq(arr,size, 10,4);
    lab1<double>::quazy_seq(arr,size, 10,4);
    lab1<double>::quazy_seq(arr,size, 10,4);

    end = std::clock();

    duration = static_cast<double>(end - start) / CLOCKS_PER_SEC;
    std::cout << "Время работы функции quazy_seq: " << duration/5 << " секунд" << std::endl;



    start = std::clock();
    lab1<double>::sinusoidal_seq(arr,size, 5, 10,M_PI_2, 15);
    lab1<double>::sinusoidal_seq(arr,size, 5, 10,M_PI_2, 15);
    lab1<double>::sinusoidal_seq(arr,size, 5, 10,M_PI_2, 15);
    lab1<double>::sinusoidal_seq(arr,size, 5, 10,M_PI_2, 15);
    lab1<double>::sinusoidal_seq(arr,size, 5, 10,M_PI_2, 15);

    end = std::clock();

    duration = static_cast<double>(end - start) / CLOCKS_PER_SEC;
    std::cout << "Время работы функции sinusoidal_seq: " << duration/5 << " секунд" << std::endl;


    delete[] arr;
}


    