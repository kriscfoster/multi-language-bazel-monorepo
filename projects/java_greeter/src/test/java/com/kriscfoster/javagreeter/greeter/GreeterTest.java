package com.kriscfoster.javagreeter.greeter;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertTrue;


class GreeterTest {

    @Test
     void greet() {
        Greeter greeter = new Greeter();
        String greeting = greeter.greet();
        assertTrue(greeting.contains("Hello"));
    }
}
