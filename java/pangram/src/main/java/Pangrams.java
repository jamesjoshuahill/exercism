public class Pangrams {
    private static char[] alphabet = "abcdefghijklmnopqrstuvwxyz".toCharArray();

    public static boolean isPangram(String sentence) {
        String lowercaseSentence = sentence.toLowerCase();
        for (char c : alphabet) {
            if (lowercaseSentence.indexOf(c) == -1) return false;
        }
        return true;
    }
}
