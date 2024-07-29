package com.halokid.eapi.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PingController {

  @GetMapping("/ping")
  public String hello() {
    return "Ping!";
  }

}
