package demo;

import com.sun.jna.Library;
 
public interface link extends Library {
 
    public String Ase256(String plaintext, String key, String iv, int blockSize);
    public String EncryptString(GoString.ByValue input);

    public Void Hello();
    
 
}