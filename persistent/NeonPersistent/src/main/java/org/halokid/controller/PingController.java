package org.halokid.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
public class PingController {

  @GetMapping(path = "/ping", consumes = MediaType.ALL_VALUE)
  public String ping(@RequestParam String name) {
    log.info("-->>> get ping controller request");
    return  String.format("Hello %s", name) ;
  }
}


