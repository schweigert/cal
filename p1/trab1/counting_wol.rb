def countingfunc v
  s = v.size

  couting = []

  for i in v
    couting[i] ||= 0
    couting[i] = couting[i] + 1
  end

  r = [0]*v.size
  pointer = 0


  for i in 0...couting.size
    next if couting[i] == nil
    couting[i].times do
      r[pointer] = i
      pointer += 1
    end
  end

  return r

end



source = []

size = gets.to_i

size.times do
	source << gets.to_i
end
t = Time.now
puts "Counting: #{size}"
# puts "#{Time.now - t}"
countingfunc source
puts "#{Time.now - t}"
