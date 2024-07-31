package com.halokid.eapi.controller;

import org.springframework.web.bind.annotation.GetMapping;

public class UserController {

  @GetMapping("/login")
  public String login() {
    return "Ping!";
  }

}
