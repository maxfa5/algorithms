#ifndef lab1_HPP
#define lab1_HPP

#include <stddef.h>
#include <cmath>
#include <random>
#include <iostream>
#include <ctime>

#define POWER_OF_CHAINSAW 3 // амлитуда с которой будут меняться числа в пилообразной последовательности

template <typename T>
class lab1 {
public:
    static void sorted(T* dist,  size_t size, T start, T step, short  is_revers );
    static void random_seq(T* dist,  size_t size, T min, T max);
    static void  chainsaw_seq(T* dist,  size_t size,T start, T step, size_t interval);
    static void  stepper_seq(T* dist, size_t size,T start, size_t interval, T step, T delta);
    static void  quazy_seq(T* dist, size_t size, T start, T delta);
    static void sinusoidal_seq(T* array, size_t size, double amplitude=5, int interval=10, double phase=M_PI_2,T start=10);
};

using namespace std;

//глобальные переменные для рандомных значений, что бы они создавались не в каждой функции
static random_device rd;
static mt19937 gen(rd());

template <typename T>
void lab1<T>::sorted(T* dist,  size_t size, T start, T step, short  is_revers )
{
    T add = start;
    for (size_t i = 0; i < size; i++)
    {
        dist[i] = add;
        add += is_revers ? -step:step;
    }
}

template <typename T>
void  lab1<T>::random_seq(T* dist,  size_t size, T min, T max){
    if constexpr(is_same_v<T, double>){ //проверяем тип int или double
        uniform_real_distribution<T> distrib(min, max);
        for (size_t i = 0; i < size; ++i) {
            dist[i] = distrib(gen);
        }
    }else if constexpr(is_same_v<T, int>) {
        uniform_int_distribution<T> distrib(min, max);
        for (size_t i = 0; i < size; ++i) {
            dist[i] = distrib(gen);
        }
    }
}

template <typename T>
void  lab1<T>::chainsaw_seq(T* dist,  size_t size,T start, T step, size_t interval){
    
    size_t count_chainsaw = size/interval;
    size_t size_of_down_chainsaw = interval / 10; 

    for (size_t i = 0; i < count_chainsaw; ++i) {
            if (i == 0){
                sorted(dist, interval,start, step, 0 );//вверх
                dist += interval;
            }
            else{
                sorted(dist, size_of_down_chainsaw,start + step*interval , step * POWER_OF_CHAINSAW, 0 ); //вниз
                sorted(dist+size_of_down_chainsaw, interval - size_of_down_chainsaw, start - (step * POWER_OF_CHAINSAW) * size_of_down_chainsaw , step, 0 );//вверх
                dist +=size_of_down_chainsaw+ interval - size_of_down_chainsaw;
            }
        
    }
}


template <typename T>
void  lab1<T>::stepper_seq(T* dist, size_t size, T start, size_t interval, T step, T delta){
    T min = 0, max = 0;
    size_t count_interval = size/interval;
    for (size_t i = 0; i < count_interval; ++i) {
        min = start - delta;
        max = start + delta;
        random_seq(dist, interval, min, max);
        dist+=interval;
        start += step;
    }
}

template <typename T>
void  lab1<T>::quazy_seq(T* dist, size_t size, T start, T delta){
    T min = 0, max = 0;
    double multy = 1.001;
    for (size_t i = 0; i < size; ++i, dist++) {
        min = (start *multy) - delta;
        max = (start *multy) + delta;
        random_seq(dist, 1, min, max);
        multy *=1.01423;
    }
}

template <typename T>
void lab1<T>::sinusoidal_seq(T* array, size_t size, double amplitude, int interval, double phase, T start) {

    double step = 2 * M_PI / interval;
    
    for (size_t i = 0; i < size; ++i) {
        double t = i * step; 
        array[i] =start + amplitude * std::sin(t + phase);
    }
}


#endif // LAB1_HPP

