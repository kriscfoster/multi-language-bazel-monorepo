package com.kriscfoster.javagreeter.main;

import com.kriscfoster.javagreeter.greeter.Greeter;

class Main {
    public static void main(String[] args) {
        Greeter greeter = new Greeter();
        System.out.println(greeter.greet());
    }
}
