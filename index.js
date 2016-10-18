var gpio = require('rpi-gpio');

gpio.setup(7, gpio.DID_IN, readInput);

function readInput(){
    gpio.read(7, function(err, value){
        console.log('The value is ' + value);
    });
}