package com.github.andersonribeir0.starwars.handlers.impl;
import com.github.andersonribeir0.starwars.commands.InsertPlanetCommand;
import com.github.andersonribeir0.starwars.handlers.CommandHandler;
import org.springframework.stereotype.Service;

@Service
public class InsertPlanetCommandHandlerImpl implements CommandHandler<InsertPlanetCommand> {

    @Override
    public void handle(InsertPlanetCommand insertPlanetCommand) {
    }
}
