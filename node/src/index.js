"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var OSMD = require("opensheetmusicdisplay");
var osmd = new OSMD.OpenSheetMusicDisplay("osmd-container");
osmd
    .load("./sheet-music.xml")
    .then(function () {
    osmd.render();
})
    .catch(function (error) {
    console.error(error);
});
