package com.github.andersonribeir0.starwars.api;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import static org.assertj.core.api.Assertions.assertThat;

@RunWith(SpringRunner.class)
@SpringBootTest
public class PlanetControllerTest {

    @Autowired
    private PlanetController planetController;

    @Test
    public void contextLoads() throws Exception {
        assertThat(planetController).isNotNull();
    }
}
