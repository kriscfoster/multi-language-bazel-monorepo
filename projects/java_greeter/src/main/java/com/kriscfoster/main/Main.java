package com.kriscfoster.main;

import com.kriscfoster.greeter.Greeter;

class Main {
    public static void main(String[] args) {
        Greeter greeter = new Greeter();
        System.out.println(greeter.greet());
    }
}
