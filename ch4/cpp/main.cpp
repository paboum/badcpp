#include <string>
#include <cstdio>

using namespace std;

//run ulimit -s first
#define ELEMS 100000000

struct B {
    string s;
    virtual ~B() { }
} ;

struct D: B {
    string t;
} ;

int main ( )
{
    B* bp[ELEMS];
    int i;
    for (i = 0 ; i < ELEMS ; i++)
    {
        bp[i] = new D();
        bp[i]->s = 'A' + i % 26;
    }
    int32_t x = 0;
    for (i = 0 ; i < ELEMS ; i++)
        x += bp[i]->s[0] - 'A';
    printf("%d\n", x);
    //for (i = 0 ; i < ELEMS ; i++)
    //   delete bp[i];
}