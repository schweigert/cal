require 'unionf'
include Unionf

nodes = [
  :vinho,
  :vermelho,
  :preto,
  :laranja,
  :verde,
  :amarelo,
  :ciano,
  :azul,
  :roxo,
  :pink
]

floresta = UnionFind.new nodes

g = [
  [:vinho, :vermelho, 3],
  [:vinho, :laranja, 6],
  [:vinho, :amarelo, 9],
  [:vermelho, :preto, 2],
  [:vermelho, :laranja, 4],
  [:vermelho, :verde, 9],
  [:vermelho, :amarelo, 9],
  [:preto, :laranja, 2],
  [:preto, :verde, 8],
  [:preto, :ciano, 9],
  [:laranja, :ciano, 9],
  [:verde, :amarelo, 8],
  [:verde, :azul, 7],
  [:verde, :roxo, 9],
  [:verde, :pink, 10],
  [:amarelo, :pink, 18],
  [:ciano, :azul, 4],
  [:ciano, :roxo, 5],
  [:azul, :roxo, 1],
  [:azul, :pink, 4],
  [:roxo, :pink, 3]
]

arv_min = []

g.sort! do | a,b |
  a[2] <=> b[2]
end

for v in g
  next if floresta.connected? v[0], v[1]

  arv_min << v
  floresta.union v[0], v[1]
end

arv_min.each { |a| puts a.to_s }
