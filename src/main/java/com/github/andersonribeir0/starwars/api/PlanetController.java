package com.github.andersonribeir0.starwars.api;

import com.github.andersonribeir0.starwars.commands.InsertPlanetCommand;
import com.github.andersonribeir0.starwars.handlers.CommandHandler;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController("/planet")
public class PlanetController {

    private final CommandHandler insertPlanetCommandHandlerImpl;

    @Autowired
    public PlanetController(CommandHandler insertPlanetCommandHandlerImpl) {
        this.insertPlanetCommandHandlerImpl = insertPlanetCommandHandlerImpl;
    }

    @PostMapping
    public ResponseEntity insertPlanet(@RequestBody InsertPlanetCommand insertPlanetCommand) {
        insertPlanetCommandHandlerImpl.handle(insertPlanetCommand);
        return ResponseEntity.ok().build();
    }
}
