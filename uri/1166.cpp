#include <iostream>
#include <cmath>

using namespace std;

int m_bolas[51];

bool quadrado_perfeito(int n) {
   int root = (int)(sqrt(n) + 0.5);

   return root * root == n;
}

void calculatem_bolas() {
   int ultima_bola[51] = { 0 };

   int next = 1;
   int peg = 0;

   while(true) {
      bool continua_usando = true;
      for(int i = 1; i <= peg; ++i) {
         if (quadrado_perfeito(ultima_bola[i] + next)) {
            continua_usando = false;
            ultima_bola[i] = next;
         }
      }

      if (continua_usando) {
         m_bolas[peg] = next - 1;
         ultima_bola[++peg] = next;

         if (peg == 51)
            break;
      }

      ++next;
   }
}

int main() {
   calculatem_bolas();
   int t;
   cin >> t;

   for(int i = 0; i < t; ++i) {
      int n;
      cin >> n;
      cout << m_bolas[n] << '\n';
   }
}
