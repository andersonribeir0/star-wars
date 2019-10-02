package com.github.andersonribeir0.starwars;

import com.github.andersonribeir0.starwars.commands.InsertPlanetCommand;
import com.github.andersonribeir0.starwars.exceptions.InsertPlanetException;
import com.github.andersonribeir0.starwars.handlers.CommandHandler;
import com.google.gson.Gson;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;

import static org.mockito.Mockito.doThrow;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@RunWith(SpringRunner.class)
@SpringBootTest
@AutoConfigureMockMvc
public class StartWarsApplicationTests {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    public CommandHandler planetHandlerMock;

    @Test
	public void should_return_bad_request_when_trying_insert_planet_with_empty_name() throws Exception {
        doThrow(InsertPlanetException.class).when(planetHandlerMock).handle(Mockito.any());
        InsertPlanetCommand invalidCommand = new InsertPlanetCommand("", "anyClimate", "anyTerrain");

        mockMvc.perform(
                    post("/planet")
                    .contentType(MediaType.APPLICATION_JSON)
                    .content(new Gson().toJson(invalidCommand)))
                .andExpect(status().isBadRequest())
                .andReturn();
    }

}
