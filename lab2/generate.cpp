#include "../lab1/lab1.hpp"
void gen_double_sinus_in_file(size_t size, double start, const char * file_name );


void gen_double_sorted_in_file(size_t size, double start, bool is_rev, const char * file_name );
void gen_double_random_in_file(size_t size, double start, const char * file_name );

int main()
{
    gen_double_sinus_in_file(30000,10,"si");
    gen_double_random_in_file(300,10,"ra");
    gen_double_sorted_in_file(300,10,"so", 0);
    gen_double_sorted_in_file(100000,10,true,"sr" );

    return 0;
}

void gen_double_sinus_in_file(size_t size, double start, const char * file_name )
{
    FILE * dst = fopen(file_name, "w+");
    if (dst)
    {
        double *arr = new double[size];
        lab1<double>::sinusoidal_seq(arr,size, 10, 20,M_PI_2, 15);
        for (size_t i = 0; i < size; i++)
        {
            fprintf(dst, "%lf\n", arr[i]);
        }
        fclose(dst);
        delete arr;
        
    }
    
}


void gen_double_random_in_file(size_t size, double start, const char * file_name )
{
    FILE * dst = fopen(file_name, "w+");
    if (dst)
    {
        double *arr = new double[size];
        lab1<double>::random_seq(arr,size, 10, 100);
        for (size_t i = 0; i < size; i++)
        {
            fprintf(dst, "%lf\n", arr[i]);
        }
        fclose(dst);
        delete arr;
        
    }
    
}



void gen_double_sorted_in_file(size_t size, double start, bool is_rev, const char * file_name )
{
    FILE * dst = fopen(file_name, "w+");
    if (dst)
    {
        double *arr = new double[size];
        lab1<double>::sorted(arr,size, 10,2, is_rev);
        for (size_t i = 0; i < size; i++)
        {
            fprintf(dst, "%lf\n", arr[i]);
        }
        fclose(dst);
        delete arr;
        
    }
    
}