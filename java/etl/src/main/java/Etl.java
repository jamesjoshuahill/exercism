import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Etl {
   public Map<String, Integer> transform(Map<Integer, List<String>> old) {
       Map<String, Integer> transformed = new HashMap<>();

       old.forEach((score, letters) -> {
           letters.forEach((letter) -> {
               transformed.put(letter.toString().toLowerCase(), score);
           });
       });

       return transformed;
   }
}
