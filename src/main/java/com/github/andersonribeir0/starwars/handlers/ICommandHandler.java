package com.github.andersonribeir0.starwars.handlers;

import com.github.andersonribeir0.starwars.commands.Command;

public interface ICommandHandler<T extends Command> {
    void handle(T command);
}
