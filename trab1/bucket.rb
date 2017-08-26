class Array
  def quicksort
    return [] if empty?

    pivot = delete_at(rand(size))
    left, right = partition(&pivot.method(:>))

    return *left.quicksort, pivot, *right.quicksort
  end
end


def bucket arr

  bucketsize = Math::log(arr.size).to_i

  min = arr.min
  max = arr.max

  range = max - min

  bucket = Array.new(bucketsize)

  for i in 0...bucketsize
    bucket[i] = []
  end

  for i in arr
      bucket[(((i-min)/max.to_f)*bucketsize.to_f).to_i] ||= []
    bucket[(((i-min)/max.to_f)*bucketsize.to_f).to_i] << i
  end


  bucketsize = bucket.size

  for i in 0...bucketsize
    bucket[i] = bucket[i].quicksort
  end

  r = []

  for i in bucket
    r = r+i
  end

  return r

end

source = []

n = gets.to_i

n.times do
  source << gets.to_i
end

t = Time.now
puts "Bucket: #{n}"
# puts "#{Time.now - t}"

bucket source

puts "#{Time.now - t}"
