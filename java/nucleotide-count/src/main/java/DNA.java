import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

public class DNA {
    public static final List<Character> VALID_BASES = Arrays.asList('A', 'C', 'G', 'T');
    private final String bases;

    public DNA(String bases) {
        this.bases = bases;
    }

    public Integer count(Character base) {
        if (!VALID_BASES.contains(base)) {
            throw new IllegalArgumentException("Invalid base: " + base);
        }
        return (int) bases.chars().filter(character -> character == base).count();
    }

    public HashMap<Character, Integer> nucleotideCounts() {
        HashMap<Character, Integer> counts = new HashMap<>();
        VALID_BASES.forEach(character -> counts.put(character, count(character)));
        return counts;
    }
}
