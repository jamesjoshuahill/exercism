require 'prime'

class Sieve
  def initialize(limit)
    @limit = limit
    @numbers = (0..@limit).to_a
  end

  def primes
    @primes ||= sieve
  end

  private

  def nil_zero_and_one
    @numbers[0] = @numbers[1] = nil
  end

  def nil_multiples_of(number)
    (number**2).step(by: number, to: @limit) do |multiple|
      @numbers[multiple] = nil
    end
  end

  def sieve
    nil_zero_and_one
    @numbers.each do |number|
      next if number.nil?
      break if number**2 > @limit
      nil_multiples_of(number)
    end
    @numbers.compact
  end
end
