def rand_name
  o = [('a'..'z'), ('A'..'Z')].map(&:to_a).flatten
  (0...5).map { o[rand(o.length)] }.join
end

100.times do
  puts "#{rand_name} #{rand_name}"
end
