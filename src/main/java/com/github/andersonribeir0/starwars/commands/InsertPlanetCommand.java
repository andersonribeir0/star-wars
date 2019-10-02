package com.github.andersonribeir0.starwars.commands;

public class InsertPlanetCommand implements Command {
    private final String name;
    private final String climate;
    private final String terrain;

    public InsertPlanetCommand(String name, String climate, String terrain) {
        this.name = name;
        this.climate = climate;
        this.terrain = terrain;
    }

    public String getName() {
        return name;
    }

    public String getClimate() {
        return climate;
    }

    public String getTerrain() {
        return terrain;
    }
}
