import java.util.regex.Pattern;

public class Bob {
    public String hey(String something) {
        Phrase phrase = new Phrase(something);
        if (phrase.isSilence()) return "Fine. Be that way!";
        if (phrase.isShout()) return "Whoa, chill out!";
        if (phrase.isQuestion()) return "Sure.";
        return "Whatever.";
    }

    public class Phrase {
        private final Pattern UPPERCASE_MATCHER = Pattern.compile(".*\\p{Upper}.*", Pattern.UNICODE_CHARACTER_CLASS);
        private final Pattern LOWERCASE_MATCHER = Pattern.compile(".*\\p{Lower}.*", Pattern.UNICODE_CHARACTER_CLASS);
        private final String phrase;

        public Phrase(String phrase) {
            this.phrase = phrase;
        }

        public boolean isSilence() {
            return phrase.replaceAll("\\s", "").isEmpty();
        }

        public boolean isShout() {
            return hasUppercaseLetters() && !hasLowercaseLetters();
        }

        public boolean isQuestion() {
            return phrase.endsWith("?");
        }

        private boolean hasLowercaseLetters() {
            return LOWERCASE_MATCHER.matcher(phrase).find();
        }

        private boolean hasUppercaseLetters() {
            return UPPERCASE_MATCHER.matcher(phrase).find();
        }
    }
}
