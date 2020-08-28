#include <stdio.h>
#include "mytest.h"

int main(int argc, char** argv)
{
	(void)argc;
	(void)argv;
	mytest_t tests;

#ifdef MYTEST1_DEBUG_SUPPORT
	tests.a = 1;
	tests.b = 1;
	printf("MYTEST1_DEBUG_1");
#endif

#ifdef MYTESTDEBUG_1
	tests.a = 2;
	tests.b = 2;
	printf("MYTEST1_DEBUG_2");
#endif

	printf("mytest1 struct.a=%d, struct .b=%d\n", tests.a, tests.b);

	return 0;
}