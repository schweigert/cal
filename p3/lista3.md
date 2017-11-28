# Exercício 1
Reduzir a seguinte instância de SAT em uma instância 3-CNF-SAT: o=x1 ^(x2 v ~x3)

Feito no caderno

# Exercício 2

Reduzie a seguinte instância de 3-CNF-SAT em uma instância de CLIQUE. Mostre uma solução equivalente para os dois problemas:
o = (x1 v ~x2 v ~x3) ^ (~x1 v ~x2 v ~x3) ^ (~x1 v x2 v ~x3)

# Exercício 3
Reduzir a seguinte instância de VERTEX-COVER em uma instância de CLIQUE. Mostre uma solução equivalente para ambos os problemas:

A -> [C, B]
B -> [E, F]
C -> [D, E, A]
D -> [C, E]
E -> [D, E, B]
F -> [B, E, G]
G -> [F]

#Exercício 4
Como seria codificada a sequência de caracteres:
"Abracadabra pe de cabra" em codificação de HUFFMAN

## Etapa 1
```ruby
"A" -> 7
"B" -> 3
" " -> 3
"R" -> 3
"C" -> 2
"D" -> 2
"E" -> 2
"P" -> 1
```

## Etapa 2

```ruby
(A: 7) (B: 3) (' ': 3) (R: 3) (C: 2) (D: 2) (E: 2) (P: 2)

(A: 7) (B: 3) (' ': 3) (R: 3) (C: 2) (D: 2) ((E: 2)(P: 1): 3)
(A: 7) (B: 3) (' ': 3) (R: 3) ((E: 2)(P: 1): 3) (C: 2) (D: 2)

(A: 7) (B: 3) (' ': 3) (R: 3) ((E: 2)(P: 1): 3) ((C: 2)(D: 2): 4)
(A: 7) ((C: 2)(D: 2): 4) (B: 3) (' ': 3) (R: 3) ((E: 2)(P: 1): 3)

(A: 7) ((C: 2)(D: 2): 4) (B: 3) (' ': 3) ((R: 3)((E: 2)(P: 1): 3): 6)
(A: 7) ((R: 3)((E: 2)(P: 1): 3): 6) ((C: 2)(D: 2): 4) (B: 3) (' ': 3)

(A: 7) ((R: 3)((E: 2)(P: 1): 3): 6) ((C: 2)(D: 2): 4) ((B: 3)(' ': 3): 6)
(A: 7) ((R: 3)((E: 2)(P: 1): 3): 6) ((B: 3)(' ': 3): 6) ((C: 2)(D: 2): 4)

(A: 7) ((R: 3)((E: 2)(P: 1): 3): 6) (((B: 3)(' ': 3): 6)((C: 2)(D: 2): 4): 10)
(((B: 3)(' ': 3): 6)((C: 2)(D: 2): 4): 10) (A: 7) ((R: 3)((E: 2)(P: 1): 3): 6)

(((B: 3)(' ': 3): 6)((C: 2)(D: 2): 4): 10) ((A: 7)((R: 3)((E: 2)(P: 1): 3): 6): 13)
((A: 7)((R: 3)((E: 2)(P: 1): 3): 6): 13) (((B: 3)(' ': 3): 6)((C: 2)(D: 2): 4): 10)

(((A: 7)((R: 3)((E: 2)(P: 1): 3): 6): 13)(((B: 3)(' ': 3): 6)((C: 2)(D: 2): 4): 10): 23)

(             0
  (A: 7)      00
  (           01
    (R: 3)    010
    (         011
      (E: 2)  0110
      (P: 1)  0111
    )
  )
)
(             1
  (           10
    (B: 3)    100
    (' ': 3)  101
  )
  (           11
    (C: 2)    110
    (D: 2)    111
  )
)

```

## Etapa 3

```ruby
A:    '00'
R:    '010'
E:    '0110'
P:    '0111'
B:    '100'
C:    '110'
D:    '111'
' ':  '101'
```

# Etapa 4
```ruby
"
Abracadabra\s
pe\s
de\s
cabra\s
"

"
  01100001 01100010 01110010 01100001 01100011 01100001 01100100 01100001 01100010 01110010 01100001 00100000
  01110000 01100101 00100000
  01100100 01100101 00100000
  01100011 01100001 01100010 01110010 01100001
"

"
  00 100 010 00 110 00 111 00 100 010 00 101
  0111 0110 101
  111 0110 101
  110 00 100 010 00
"
```

# Exercício 5
Multiplicar as matrizes com a melhor eficiência:
```ruby
# M(N,M)
A1(2,5)
A2(5,4)
A3(4,3)
A4(3,7)
```
Realizar `A1*A2*A3*A4`

Sabemos que a complexidade de multiplicar `Ax(n,m)` e `Ay(m,p)` é `O(nmp)`.
Lembramos que `Ax * Ay = B(n,p)`.

- (((A1 A2) A3) A4): 106
- ((A1 (A2 A3)) A4): 132
- (A1 ((A2 A3) A4)): 235
- ((A1 A2) (A3 A4)): 180
- (A1 (A2 (A3 A4))): 294
