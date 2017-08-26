def countingfunc v
  s = v.size

  couting = []

  for i in v
    couting[i] ||= 0
    couting[i] = couting[i] + 1
  end

  r = []


  for i in 0...couting.size
    next if couting[i] == nil || couting[i] == 0
    r = r + ([i]*couting[i])
  end
  return r

end



source = []

size = gets.to_i

size.times do
	source << gets.to_i
end

puts(countingfunc source)
