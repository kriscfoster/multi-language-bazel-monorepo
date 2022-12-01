package com.kriscfoster.java_greeter.main;

import com.kriscfoster.java_greeter.greeter.Greeter;

class Main {
    public static void main(String[] args) {
        Greeter greeter = new Greeter();
        System.out.println(greeter.greet());
    }
}