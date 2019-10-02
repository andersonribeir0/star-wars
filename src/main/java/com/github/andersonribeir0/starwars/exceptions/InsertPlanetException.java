package com.github.andersonribeir0.starwars.exceptions;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.BAD_REQUEST)
public class InsertPlanetException extends RuntimeException {

}
