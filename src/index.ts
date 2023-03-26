import * as OSMD from "opensheetmusicdisplay";

const osmd = new OSMD.OpenSheetMusicDisplay("osmd-container");

osmd
    .load("./sheet-music.xml")
    .then(() => {
        osmd.render();
    })
    .catch((error) => {
        console.error(error);
        
    });
