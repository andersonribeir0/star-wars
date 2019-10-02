package com.github.andersonribeir0.starwars.api;

import com.github.andersonribeir0.starwars.commands.InsertPlanetCommand;
import com.github.andersonribeir0.starwars.handlers.ICommandHandler;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController("/planet")
public class PlanetController {

    private final ICommandHandler planetHandler;

    @Autowired
    public PlanetController(ICommandHandler planetHandler) {
        this.planetHandler = planetHandler;
    }

    @PostMapping
    public ResponseEntity insertPlanet(@RequestBody InsertPlanetCommand anInsertPlanetCommand) {
        planetHandler.handle(anInsertPlanetCommand);
        return ResponseEntity.ok().build();
    }
}
