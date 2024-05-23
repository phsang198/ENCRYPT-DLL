
import com.sun.jna.Library;
import com.sun.jna.Native;

public class GoEngineTransformerTest {
  static GoCqmTransformer GO_CQM_TRANSFORMER;
  static {
    GO_CQM_TRANSFORMER = (GoCqmTransformer) Native.loadLibrary("mylib.dll", GoCqmTransformer.class);
  }

  public interface GoCqmTransformer extends Library {
    String EncryptString(String input);
  }

  public static void main(String[] args) {
    System.out.println(GO_CQM_TRANSFORMER.EncryptString("HELLO WORLD!"));
  }
}