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
1:(16) 1:(13) 1:(10) 1:(7) 1:(5) 1:(3) 1: "E" (2)
                                       0: "P" (1)
                                 0:(2) "D" (2)
                           0:(2) "C" (2)
                     0:(3) "R" (3)
              0:(3) " " (3)
       0:(3) "B" (3)
0:(7) "A" (7)

```
