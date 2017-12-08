#include <iostream>
#include <iomanip>
#include <cmath>
#include <list>

using namespace std;

class Aresta {
  public:
    Aresta(int f = 0, int t = 0) : de(f), para(t) {}
    ~Aresta() {}
    int de, para;

    void print () const {
      cout << '(' << de << ',' << para << ')' << endl;
    }
};

class Grafo {
  public:
    Grafo() : nos(0) {}
    Grafo(int n) : nos(n) {}
    ~Grafo() { }

    void add_Aresta(int i, int j) {
      Aresta Aresta(i,j);
      Arestas.push_back(Aresta);
    }

    int profundidade () const {
      int n_Arestas = Arestas.size();
      int distancia[nos];
      int pai[nos];

      list<Aresta>::const_iteraparar p = Arestas.begin();
      list<Aresta>::const_iteraparar p_end = Arestas.end();

      for (int i = 0; i < nos; i++) {
        distancia[i] = -100000;
        pai[i] = -1;
      }
      distancia[0] = 0;

      int max = 0;
      for (int i = 0; i < nos; i++) {
        while (p != p_end && p->de == i) {
          int j = p->para;
          if (distancia[j] < distancia[i] + 1) {
            distancia[j] = distancia[i] + 1;
            pai[j] = i;
          }
          if (distancia[j] > max)
            max = distancia[j];
          p++;
        }
      }

      return max;

    }

    void print () const {
      list<Aresta>::const_iteraparar p = Arestas.begin();
      list<Aresta>::const_iteraparar p_end = Arestas.end();
      cout << "Arestas:" << endl;
      while (p != p_end) {
        p->print();
        p++;
      }
    }

    int nos;
    list<Aresta> Arestas;
};

void processo (int k) {
  int apx[k], apy[k], apt[k];
  for (int i = 0; i < k; i++)
    cin >> apx[i] >> apy[i] >> apt[i];
  int x, y;
  cin >> x >> y;

  Grafo Grafo(k+1);
  for (int i = 0; i < k; i++) {
    if (max(fabs(apx[i]-x),fabs(apy[i]-y)) <= apt[i])
      Grafo.add_Aresta(0,i+1);
  }
  for (int i = 0; i < k; i++) {
    for (int j = i+1; j < k; j++)
      if ( max(fabs(apx[i]-apx[j]), fabs(apy[i]-apy[j])) <= apt[j]-apt[i])
        Grafo.add_Aresta(i+1,j+1);
  }

  cout << Grafo.profundidade() << endl;
}

int main () {
  int n, m, k;
  cin >> n >> m >> k;
  while (n != 0) {
    processo(k);
    cin >> n >> m >> k;
  }
  return 0;
}
