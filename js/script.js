// Copyright (c) 2023 Lamees Hemdan All rights reserved
//
// Created by: Lamees Hemdan
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

function myButtonClicked () {
  // this functions lets you enter age and day of week and give you an output

  // input
  const age = parseInt(document.getElementById('age').value)
  const day = document.getElementById('day').value

  // process
  if ((day == "Tuesday" || day == "Thursday") || (age >= 13 && age <= 18)) {
    // output
    document.getElementById('output').innerHTML = "You can get in for free!"
  } else {
    // output
    document.getElementById('output').innerHTML = "You have to pay."
  }
}
