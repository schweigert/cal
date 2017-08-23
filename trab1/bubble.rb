
source = []

size = gets.to_i

size.times do
	source << gets.to_i
end


for i in 0..(size-1)
	for j in 0..(size-1)
		source[i], source[j] = source[j], source[i] if (source[i] > source[j])
	end
end

puts source.to_s

