#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <math.h>
#include "../lab1/lab1.hpp"
double *get_arr(long size);
int compare_doubles(const void* a, const void* b);
static long count_op = 0;

int main(){
    srand(time(NULL));
    long size = 30000;
    double* arr = get_arr(size);

    clock_t start_time = clock();
    qsort(arr, size, sizeof(double), compare_doubles);
    clock_t end_time = clock();

    double elapsed_time = (double)(end_time - start_time) / CLOCKS_PER_SEC;

    printf("qsort took %f seconds and %ld compares to sort %ld elements.\n", elapsed_time, count_op, size);
    if (arr)
    {
        delete[] arr;
    }
    
}

double *get_arr(long size){
    double *arr = new double[size];
    lab1<double>::sinusoidal_seq(arr,size, 10, 20 , M_PI_2, 15);
    // lab1<double>::sorted(arr,size, 10,2, 0);
    // lab1<double>::random_seq(arr,size, 10, 100);

    return arr;
} 

int compare_doubles(const void* a, const void* b) {
    double arg1 = *(const double*)a;
    double arg2 = *(const double*)b;
    if (arg1 < arg2){
        count_op++;
         return -1;}
    if (arg1 > arg2){
        count_op+= 2;
        return 1;
    }
    count_op+=2;
    return 0;
}

