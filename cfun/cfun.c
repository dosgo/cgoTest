#include <stdio.h>
#include <stdlib.h>
#include "cfun.h"
int number_add_mod(int a, int b, int mod)
{
    return (a + b) % mod;
}