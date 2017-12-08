#include <bits/stdc++.h>

#define MAXN 2000001
using namespace std;

bitset<MAXN> primo;
int divr[MAXN], resp[MAXN];

void calcular()
{
	primo.set();
	memset(divr, 0 , sizeof(divr));

	for (int i = 2; i < MAXN; i++)
	{
		if (primo[i])
		{
			divr[i] = 2;
			for (int j = i + i; j < MAXN; j+= i)
			{
				primo.reset(j);

				int divs = 0;
				int tam = j;

				while ( tam % i == 0) tam/= i, divs++;

				if (divr[j] == 0) divr[j] = divs + 1;
				else divr[j]*= (divs + 1);
			}
		}
	}

	resp[1] = 0;
	for (int i = 2; i <= MAXN; i++)
	{
		resp[i] = resp[i - 1];
		if (primo[divr[i]]) resp[i]++;
	}
}

int main()
{
	int n;

	calcular();
	while(cin >> n)
	{
		cout << resp[n] << '\n';
	}
}
