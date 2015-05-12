module Acronym
  def self.abbreviate(full_name)
    full_name.scan(/^(\w)|\W(\w)|[a-z]([A-Z])/).flatten.join.upcase
  end
end
