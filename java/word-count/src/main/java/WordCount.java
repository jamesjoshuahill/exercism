import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class WordCount {
    public Map<String,Integer> phrase(String phrase) {
        List<String> words = Arrays.asList(phrase.split("\\W+"));
        Map<String, Integer> counts = new HashMap<>();

        words.forEach(word -> {
            String key = word.toLowerCase();
            Integer count = counts.containsKey(key) ? counts.get(key) + 1 : 1;
            counts.put(key, count);
        });

        return counts;
    }
}
