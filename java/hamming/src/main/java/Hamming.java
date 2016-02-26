import java.util.stream.IntStream;

public class Hamming {
    public static Integer compute(String strand, String other) {
        if (strand.length() != other.length()) {
            throw new IllegalArgumentException("Strands do not have equal length");
        }

        Integer distance = IntStream.range(0, strand.length())
                .reduce(0, (count, index) -> {
                    if (strand.charAt(index) != other.charAt(index)) {
                        count++;
                    }
                    return count;
                });

        return distance;
    }
}
