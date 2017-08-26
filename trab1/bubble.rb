
source = []

size = gets.to_i

size.times do
	source << gets.to_i
end

puts "Bubble: #{size}"
# puts "#{Time.now - t}"

t = Time.now
for i in 0...size
	break if(i/size.to_f >= 0.1)
	for j in i...size
		source[i], source[j] = source[j], source[i] if (source[i] > source[j])
	end
end

puts "#{Time.now - t}"
