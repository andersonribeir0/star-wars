package com.github.andersonribeir0.starwars.handlers;

import com.github.andersonribeir0.starwars.commands.Command;

public interface CommandHandler<T extends Command> {
    void handle(T command);
}
