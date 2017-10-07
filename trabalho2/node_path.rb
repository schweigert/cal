require 'unionf'
include Unionf

nodos = [1,2,3,4,5]
conjunto = UnionFind.new nodos
lista_adj = {
              1 => [2,4],
              2 => [1,3,5],
              3 => [2],
              4 => [1,5],
              5 => [2,4]
            }

lista_adj.each_key do |k|
  lista_adj[k].each { |n| conjunto.union k, n }
end

puts 'Existe caminho de 1 para 5?'
puts conjunto.connected? 1,5
