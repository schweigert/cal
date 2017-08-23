# Gera os arquivos de entrada

quantia = [25000,50000,100000,1000000]

# na.in -> Crescente
# nb.in -> Decrescente
# nc.in -> Aleatorio
# nd.in -> Aleatório com desvio gigante

for i in quantia
	a = File.open "#{i}a.in", "w"
	b = File.open "#{i}b.in", "w"
	c = File.open "#{i}c.in", "w"
	d = File.open "#{i}d.in", "w"

	a.puts i
	b.puts i
	c.puts i
	d.puts i

	# Posicao do desvio
	pos = rand(i)

	i.times do |n|
		# Crescente
		a.puts(n+1)
		# Decrescente
		b.puts(i-n+1)
		# Aleatorio
		c.puts(rand(i))
		# Aleatório com Desvio
		if pos == n
			d.puts i*100
		else
			d.puts rand(i)
		end
	end
	a.close
	b.close
	c.close
	d.close
end
