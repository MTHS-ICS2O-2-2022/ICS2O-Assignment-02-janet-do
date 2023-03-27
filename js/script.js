// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file generates the area of a triangle
"use strict"
function calculate() {
  // input
  const radius = parseInt(document.getElementById("radius").value)
  const height = parseInt(document.getElementById("height").value)

  // process
  const volume = 3.14 * radius**2 * height

  // output
  document.getElementById("volume").innerHTML = "volume: " + volume + " cm³"
}
