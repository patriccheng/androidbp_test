#include <stdio.h>
#include "mytest.h"

int main(int argc, char** argv)
{
	(void)argc;
	(void)argv;
	mytest_t tests;

#ifdef MYTEST2DEBUG_1
	tests.a = 1;
	tests.b = 1;
	printf("MYTEST2DEBUG_1");
#endif

#ifdef MYTEST2DEBUG_2
	tests.a = 2;
	tests.b = 2;
	printf("MYTEST2DEBUG_2");
#endif

	printf("mytest2 struct.a=%d, struct .b=%d\n", tests.a, tests.b);

	return 0;
}