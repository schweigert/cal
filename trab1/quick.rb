# Implementação do Quick com elemento inicial

class Array
  def quicksort
    return [] if empty?

    pivot = delete_at(rand(size))
    left, right = partition(&pivot.method(:>))

    return *left.quicksort, pivot, *right.quicksort
  end
end


source = []

n = gets.to_i

n.times do
	source << gets.to_i
end

source.quicksort
