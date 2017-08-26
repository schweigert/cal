def insertion_sort(arr)
  for i in (1...(arr.size))
    break if i/ arr.size.to_f > 0.1
    if arr[i-1] > arr[i]
      i.downto(1) do |el|
        if arr[el] < arr[el-1]
          arr[el-1], arr[el] = arr[el], arr[el-1]
        end
      end
    end
  end
  return arr
end

size = gets.to_i
source = []

size.times do
  source << gets.to_i
end
t = Time.now
puts "Insert: #{size}"
# puts "#{Time.now - t}"
insertion_sort(source)
puts "#{Time.now - t}"
