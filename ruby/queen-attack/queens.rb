class Colour
  def self.white
    new("W")
  end

  def self.black
    new("B")
  end

  def initialize(short_name)
    @short_name = short_name
  end

  def to_s
    @short_name
  end
end

class NullPiece
  def to_s
    "_"
  end
end

class Piece
  attr_reader :colour, :position

  def initialize(colour, position)
    @colour = colour
    @position = position
  end

  def row
    position.first
  end

  def column
    position.last
  end

  def to_s
    colour.to_s
  end
end

class Queen < Piece
  def attack?(other)
    straight_line_attack?(other) || diagonal_attack?(other)
  end

  def straight_line_attack?(other)
    row == other.row || column == other.column
  end

  def diagonal_attack?(other)
    (row - other.row).abs == (column - other.column).abs
  end
end

class QueenMaker
  def initialize(colour_klass = Colour, queen_klass = Queen)
    @colour_klass = colour_klass
    @queen_klass = queen_klass
  end

  def white(position)
    @queen_klass.new(@colour_klass.white, position)
  end

  def black(position)
    @queen_klass.new(@colour_klass.black, position)
  end
end

class Board
  def initialize(null_klass = NullPiece)
    @board = Array.new(8) { Array.new(8) { null_klass.new } }
  end

  def place(piece)
    @board[piece.row][piece.column] = piece
  end

  def to_s
    rows = @board.reduce([]) do |rows, row|
      row.map { |piece| piece.to_s }
      rows << row.join(" ")
    end
    rows.join("\n")
  end
end

class Queens
  def initialize(board = Board.new, generator = QueenMaker.new, white: [0, 3], black: [7, 3])
    if white == black
      raise ArgumentError, "Queens cannot occupy the same square"
    end
    @white_queen = generator.white(white)
    @black_queen = generator.black(black)
    @board = board
    @board.place(@white_queen)
    @board.place(@black_queen)
  end

  def white
    @white_queen.position
  end

  def black
    @black_queen.position
  end

  def to_s
    @board.to_s
  end

  def attack?
    @white_queen.attack?(@black_queen)
  end
end
