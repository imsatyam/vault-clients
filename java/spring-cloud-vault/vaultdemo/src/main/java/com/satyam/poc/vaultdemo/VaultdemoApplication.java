package com.satyam.poc.vaultdemo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.core.env.Environment;

import javax.annotation.PostConstruct;

@SpringBootApplication
public class VaultdemoApplication {


	public static void main(String[] args) {
		SpringApplication.run(VaultdemoApplication.class, args);
	}


	@Autowired
	private Environment env;

	@Value("${username:}")
	private String user;

	@Value("${password:}")
	private String password;

	@PostConstruct
	private void postConstruct() {
		System.out.println("My Username is: " + user);
		System.out.println("My password is: " + password);
	}

}
