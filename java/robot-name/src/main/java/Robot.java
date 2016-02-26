import java.util.Random;

public class Robot {

    private static final String LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    public static final Random RANDOM = new Random();
    private String name;

    public Robot() {
        reset();
    }

    public String getName() {
        return name;
    }

    public void reset() {
        this.name = generateName();
    }

    private String generateName() {
        return randomLetter() + randomLetter() + randomNumber();
    }

    private String randomLetter() {
        char randomCharacter = LETTERS.charAt(RANDOM.nextInt(LETTERS.length()));
        return String.valueOf(randomCharacter);
    }

    private String randomNumber() {
        return String.format("%03d", RANDOM.nextInt(1000));
    }
}
