import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Anagram {
    
    private final String word;
    private final String sortedWord;

    public Anagram(String word) {
        this.word = word;
        this.sortedWord = sortWord(word);
    }

    public List<String> match(List<String> candidates) {
        List<String> matches = new ArrayList<>();

        candidates.forEach(candidate -> {
            if (isDifferentWord(candidate) && isAnagram(candidate)) {
                matches.add(candidate);
            }
        });

        return matches;
    }

    private boolean isDifferentWord(String candidate) {
        return !candidate.toLowerCase().equals(word);
    }

    private boolean isAnagram(String candidate) {
        return sortWord(candidate).equals(sortedWord);
    }

    private String sortWord(String word) {
        char[] characters = word.toLowerCase().toCharArray();
        Arrays.sort(characters);
        return new String(characters);
    }
}
