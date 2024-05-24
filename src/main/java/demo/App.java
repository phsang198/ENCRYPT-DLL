package demo;

import com.sun.jna.Native;


public class App 
{
    public static void main( String[] args )
    {
        
        link golang = (link) Native.loadLibrary("F:/OutSource/GOLANG/demo_1/src/main/java/demo/lib/mylib", link.class);
        
        GoString.ByValue tmp = new GoString.ByValue("Hello, World!");
        String encrypted = golang.EncryptString(tmp);
        System.out.println("encrypt string in java: " +encrypted);
    }
}
